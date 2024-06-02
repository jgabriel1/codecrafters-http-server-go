package main

import (
	"fmt"
	"testing"
)

func TestWorks(t *testing.T) {
	count := 0
	for range 2 {
		fmt.Println("ran")
		count++
	}
	if count != 2 {
		t.Error("didnt work")
	}
}
