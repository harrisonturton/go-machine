package testutil

import (
	"fmt"
	"testing"
)

const defaultMsg = "Expected %s \"%s\", recieved \"%s\""

// Compare integer slices
func CompareSlice(a []int, b []int) {
	// a == nil xor b == nil
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for elemA := range a {
		for elemB := range b {
			if elemA != elemB {
				return false
			}
		}
	}
	return true
}
