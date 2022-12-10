package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	for i := 0; i < 1000; i++ {
		Run()
	}
}
