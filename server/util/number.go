
package util

import (
	"math/big"
)


// Check if input number is a prime with returns:
//
//   false if number is not a prime
//   true  if number is a prime
//
func IsPrimeNumber(number *big.Int) bool {
	if number.ProbablyPrime(10) {
		return true
	} 
	return false
}

func IsANumber(number string) (bool, *big.Int) {
	n := new(big.Int)
    n, ok := n.SetString(number, 10)
    if !ok {
        return false, big.NewInt(-1)
    }
	return true, n
}

// Find highest prime number that less than input number x returns:
//
//   -1 			if There is no prime number that less than input number
//   highest prime  if Exist a prime less than input number
//
func FindHighestPrime(number *big.Int) (*big.Int) {
	highestPrime := big.NewInt(-1)

	lowLimit := big.NewInt(2)

	if (number.Cmp(lowLimit) <=0) {
		return highestPrime
	}

	i := number.Sub(number, big.NewInt(1))

	for i.Cmp(lowLimit) >= 0 {
		if (IsPrimeNumber(i)) {
			return i
		} 
		i = i.Sub(i, big.NewInt(1))
	}

	return big.NewInt(-1)
}