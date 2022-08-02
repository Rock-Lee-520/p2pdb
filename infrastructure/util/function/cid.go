package function

import (
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

const Version = 1
const Codec = mc.Raw
const MhType = mh.SHA2_256
const MhLength = -1

func GetCid(str string) (cid.Cid, error) {

	data := []byte(str)

	// Construct c1
	format := cid.V0Builder{}
	c1, err := format.Sum(data)

	if err != nil {
		return c1, err
	}
	return c1, nil
}

func GetCidString(str string) (string, error) {

	data := []byte(str)

	// Construct c1
	format := cid.V0Builder{}
	c1, err := format.Sum(data)

	if err != nil {
		return "", err
	}
	return c1.String(), nil
}

// Create a cid manually by specifying the 'prefix' parameters
//   pref := cid.Prefix{
// 	  Version: 1,
// 	  Codec: mc.Raw,
// 	  MhType: mh.SHA2_256,
// 	  MhLength: -1, // default length
//   }

//   // And then feed it some data
//   c, err := pref.Sum([]byte("Hello World!"))
//   if err != nil {...}

//   fmt.Println("Created CID: ", c)
