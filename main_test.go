package main

import (
	"testing"

	. "github.com/chai2010/assert"
)

func TestFoo(t *testing.T) {
	Assert(t, true)
}

func TestFailed(t *testing.T) {
	Assert(t, false)
}
