package random

import "testing"

func TestShuffle(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Shuffle("abcd"))
	}
}
