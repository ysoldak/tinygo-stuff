package main

import (
	"testing"
)

func TestInc(t *testing.T) {
	b := &Banana{}
	b.Inc()
	println(b.count)
	if b.count != 1 {
		t.Errorf("Expect 1 banana")
	}
}
