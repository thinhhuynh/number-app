
package util

import (
	"math/big"
	"testing"
)

func TestNumber_IsANumber(t *testing.T) {
	number := "123456789012345678901234567890"
	gotOk, got := IsANumber(number)

	want := new(big.Int)
    want, wantOk := want.SetString(number, 10)
	
	if want.Cmp(got) != 0 || !gotOk || !wantOk  {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestNumber_IsANumber_WithError(t *testing.T) {
	number := "a"
	got, bigNumber := IsANumber(number)
	want := false

	if got {
		t.Errorf("handler returned unexpected body: got %v want %v bigNumber %bigNumber", got, want, bigNumber)
	}
}


func TestNumber_IsPrimeNumber(t *testing.T) {
	got := IsPrimeNumber(big.NewInt(53))
	
	want := true

	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestNumber_IsPrimeNumber_WithError(t *testing.T) {
	got := IsPrimeNumber(big.NewInt(55))	
	want := false

	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}


func TestNumber_FindHighestPrime(t *testing.T) {
	got := FindHighestPrime(big.NewInt(55))
	want := big.NewInt(53)

	if want.Cmp(got) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestNumber_FindHighestPrime_WithNotFound(t *testing.T) {
	got := FindHighestPrime(big.NewInt(2))
	want := big.NewInt(-1)

	if want.Cmp(got) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

