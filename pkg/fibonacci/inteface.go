package fibonacci

import "math/big"

type IFibonacci interface {

	// IsValid validates the given number is a valid fibonacci number.
	IsValid(n big.Int) error

	// Next returns the next value of the fibonacci serries. F(IndexOf(n)+1)
	Next(n big.Int) (big.Int, error)

	// Previous returns the previous value of fibonacci series. F(IndexOf(n)-1)
	Previous(n big.Int) (big.Int, error)

	// IndexOf returns the index of n in the fibonacci series.
	IndexOf(n big.Int) (big.Int, error)

	// F calculates the nth element of the fibonacci series.
	F(n big.Int) (big.Int, error)
}
