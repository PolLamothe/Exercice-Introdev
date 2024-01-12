package pgcd

import (
	"fmt"
	"testing"
)

func TestEgal(t *testing.T) {
	if pgcd(1, 1) != 1 {
		t.Fail()
	}
}

func TestPerso(t *testing.T) {
	fmt.Println(pgcd(21, 15))
}

func TestPremier(t *testing.T) {
	fmt.Println(pgcd(3, 5))
	if pgcd(3, 5) != 1 {
		t.Fail()
	}
}

func TestDifferent(t *testing.T) {
	if pgcd(220, 315) != 5 {
		t.Fail()
	}
}
