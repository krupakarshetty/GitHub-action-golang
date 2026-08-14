package main

import "testing"

func TestAddition(t *testing.T) {
	result := 2 + 2
	expected := 4
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}