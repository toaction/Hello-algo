package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	if result != "Hello, Algo!" {
		t.Errorf("Hello() = %v, expected %v", result, "Hello, Algo!")
	}
}
