package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	msg := hello()

	if msg != "Hello world v0.9.0" {
		t.Error("TestHello - string incorrect")
	}
}
