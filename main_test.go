// main_test.go
package main

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("Add() = %v, want %v", got, 3)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
