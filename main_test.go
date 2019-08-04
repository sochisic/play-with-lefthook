package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	msg := hello()

	if msg != "Hello world v0.4.1" {
		t.Error("TestHello - string incorrect")
	}
}
