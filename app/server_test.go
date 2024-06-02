package main

import (
	"slices"
	"strings"
	"testing"
)

func TestWorks(t *testing.T) {
	if slices.Equal(strings.Split("/echo/asdf", "/"), []string{"echo", "asdf"}) {
		t.Errorf("did not work")
	}
}
