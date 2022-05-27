package main

import "testing"

// GCounter represent a G-counter in CRDT, which is
// a state-based grow-only counter that only supports
// increments.
type GCounter struct {
	// ident provides a unique identity to each replica.
	ident string

	// counter maps identity of each replica to their
	// entry values i.e. the counter value they individually
	// have.
	counter map[string]int
}

// NewGCounter returns a *GCounter by pre-assigning a unique
// identity to it.
func NewGCounter() *GCounter {
	return &GCounter{
		ident:   uuid.NewV4().String(),
		counter: make(map[string]int),
	}
}

// Inc increments the GCounter by the value of 1 everytime it
// is called.
func (g *GCounter) Inc() {
	g.IncVal(1)
}

// IncVal allows passing in an arbitrary delta to increment the
// current value of counter by. Only positive values are accepted.
// If a negative value is provided the implementation will panic.
func (g *GCounter) IncVal(incr int) {
	if incr < 0 {
		panic("cannot decrement a gcounter")
	}
	g.counter[g.ident] += incr
}

// Count returns the total count of this counter across all the
// present replicas.
func (g *GCounter) Count() (total int) {
	for _, val := range g.counter {
		total += val
	}
	return
}

// Merge combines the counter values across multiple replicas.
// The property of idempotency is preserved here across
// multiple merges as when no state is changed across any replicas,
// the result should be exactly the same everytime.
func (g *GCounter) Merge(c *GCounter) {
	for ident, val := range c.counter {
		if v, ok := g.counter[ident]; !ok || v < val {
			g.counter[ident] = val
		}
	}
}

// PNCounter represents a state-based PN-Counter. It is
// implemented as sets of two G-Counters, one that tracks
// increments while the other decrements.
type PNCounter struct {
	pCounter *GCounter
	nCounter *GCounter
}

// NewPNCounter returns a new *PNCounter with both its
// G-Counters initialized.
func NewPNCounter() *PNCounter {
	return &PNCounter{
		pCounter: NewGCounter(),
		nCounter: NewGCounter(),
	}
}

// Inc monotonically increments the current value of the
// PN-Counter by one.
func (pn *PNCounter) Inc() {
	pn.IncVal(1)
}

// IncVal increments the current value of the PN-Counter
// by the delta incr that is provided. The value of delta
// has to be >= 0. If the value of delta is < 0, then this
// implementation panics.
func (pn *PNCounter) IncVal(incr int) {
	pn.pCounter.IncVal(incr)
}

// Dec monotonically decrements the current value of the
// PN-Counter by one.
func (pn *PNCounter) Dec() {
	pn.DecVal(1)
}

// DecVal adds a decrement to the current value of
// PN-Counter by the value of delta decr. Similar to
// IncVal, the value of decr cannot be less than 0.
func (pn *PNCounter) DecVal(decr int) {
	pn.nCounter.IncVal(decr)
}

// Count returns the current value of the counter. It
// subtracts the value of negative G-Counter from the
// positive grow-only counter and returns the result.
// Because this counter can grow in either direction,
// negative integers as results are possible.
func (pn *PNCounter) Count() int {
	return pn.pCounter.Count() - pn.nCounter.Count()
}

// Merge combines both the PN-Counters and saves the result
// in the invoking counter. Respective G-Counters are merged
// i.e. +ve with +ve and -ve with -ve, but not computation
// is actually performed.
func (pn *PNCounter) Merge(pnpn *PNCounter) {
	pn.pCounter.Merge(pnpn.pCounter)
	pn.nCounter.Merge(pnpn.nCounter)
}

func TestPNCounter(t *testing.T) {
	for _, tt := range []struct {
		incOne, decOne int
		incTwo, decTwo int

		result int
	}{
		{5, 5, 6, 6, 0},
		{5, 6, 7, 8, -2},
		{8, 7, 6, 5, 2},
		{5, 0, 6, 0, 11},
		{0, 5, 0, 6, -11},
	} {
		pOne, pTwo := NewPNCounter(), NewPNCounter()

		for i := 0; i < tt.incOne; i++ {
			pOne.Inc()
		}

		for i := 0; i < tt.decOne; i++ {
			pOne.Dec()
		}

		for i := 0; i < tt.incTwo; i++ {
			pTwo.Inc()
		}

		for i := 0; i < tt.decTwo; i++ {
			pTwo.Dec()
		}

		pOne.Merge(pTwo)

		if pOne.Count() != tt.result {
			t.Errorf("expected the total count to be: %d, actual: %d",
				tt.result,
				pOne.Count())
		}

		pTwo.Merge(pOne)

		if pTwo.Count() != tt.result {
			t.Errorf("expected the total count to be: %d, actual: %d",
				tt.result,
				pTwo.Count())
		}
	}
}

func TestPNCounterInvalidP(t *testing.T) {
	pn := NewPNCounter()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("panic expected here")
		}
	}()

	pn.IncVal(-5)
}

func TestPNCounterInvalidN(t *testing.T) {
	pn := NewPNCounter()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("panic expected here")
		}
	}()

	pn.DecVal(-5)
}
