package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	result, err := Multiply("5", "5")
	if err != nil && result != 25 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", err, 25)
	}
}

func TestMultiply1(t *testing.T) {
	_, err := Multiply("vv", "aa")

	if err == nil {
		t.Errorf("No Error Produced, %d", err)
	}
}

func TestMultiply2(t *testing.T) {
	result, err := Multiply("1e4", "5")

	if err != nil && result != 50000 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", err, 25)
	}
}

func TestMultiply3(t *testing.T) {
	_, err := Multiply("vv", "aa")

	if err == nil {
		t.Errorf("No Error Produced, %d", err)
	}
}

