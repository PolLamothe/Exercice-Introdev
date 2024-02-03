package fibonacci

import (
	"fmt"
	"math/big"
)

/*
On rappelle la définition de la suite de Fibonacci :
U(0) = 0
U(1) = 1
U(n) = U(n-1) + U(n-2)

Les premiers termes de cette suite sont :
U(0) = 0
U(1) = 1
U(2) = 1
U(3) = 2
U(4) = 3
U(5) = 5
U(6) = 8
U(7) = 13
Le premier terme qui a deux chiffres est donc le terme numéro 7.

On souhaite savoir, pour un nombre k de chiffres, quel est le numéro (n) du premier terme de la suite de Fibonacci qui a au moins k chiffres. C'est le rôle de la fonction fibonacci.

Attention, la fonction doit pouvoir gérer des nombres k de chiffres très grands (plus de 1000). Pour cela, vous pouvez utiliser la paquet math/big qui permet de représenter des nombres avec un nombre arbitraire de chiffres.

# Entrée
- k : un nombre entier de chiffres

# Sortie
- n : le numéro du premier terme de la suite de fibonacci qui a au moins k chiffres

# Info
2023-2024, test 3, exercice 9
*/

func fibo(k int) (big.Int, big.Int) {
	if k == 0 {
		return *big.NewInt(0), *big.NewInt(0)
	}
	if k == 1 {
		return *big.NewInt(0), *big.NewInt(1)
	}
	before, before2 := fibo(k - 1)
	result := new(big.Int)
	result.Add(&before, &before2)
	return before2, *result
}

func fibonacci(k int) (n int) {
	z := 0
	for {
		var term2, term = fibo(z)
		if false {
			fmt.Println(term, term2)
		}
		var stringTerm = term.String()
		if len(stringTerm) >= k {
			return z
		}
		z++
	}
}
