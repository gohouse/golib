package random

import "testing"

func TestRandVariable(t *testing.T) {
	for i:=0;i<8;i++{
		t.Log(RandVariable(RandBetween(3, 11)))
	}
}

func TestRand(t *testing.T) {
	t.Log(Rand(5, TypeLOWERCASE))
}

func TestRandBetween(t *testing.T) {
	t.Log(RandBetween(7, 11))
}
