package service

import "testing"

func TestPublish(t *testing.T) {
	Publish("topic", "test publish")
}
