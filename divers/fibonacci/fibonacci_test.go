package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	fmt.Println(fibo(7))
}
func Test(t *testing.T) {
	if fibonacci(1) != 0 {
		t.Fatal(1)
	}

	if fibonacci(2) != 7 {
		t.Fatal(2)
	}

	if fibonacci(3) != 12 {
		t.Fatal(3)
	}

	// Ajouté après le test
	if fibonacci(4) != 17 {
		t.Fatal(4)
	}

	if fibonacci(5) != 21 {
		t.Fatal(5)
	}

	if fibonacci(10) != 45 {
		t.Fatal(10)
	}

	if fibonacci(100) != 476 {
		t.Fatal(100)
	}

	if fibonacci(1000) != 4782 {
		t.Fatal(1000)
	}
}
