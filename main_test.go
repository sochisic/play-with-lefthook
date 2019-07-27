package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	msg := hello()

	if msg != "Hello world" {
		t.Error("TestHello - string incorrect")
	}
}
