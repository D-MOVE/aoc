//This libraly provides test helper functions
package aoc

import "testing"

//Report checks whether or not got == want
func Report[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
