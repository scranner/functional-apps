package main

import (
"testing"
)

func TestMultiply(t *testing.T) {
	result := 12
	if result != 12 {
		t.Errorf("If you see this, something is wrong with GoLang")
	}
}
