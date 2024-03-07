package fibonacci

import "math/big"

type IFibonacci interface {

	// Next returns the next value of the fibonacci serries. F(IndexOf(n)+1)
	// returns error
	// ErrNotfound : if the n is not a valid fibonacci number
	// ErrTooBig : if the n bit size is bigger than configuration.
	Next(n *big.Int) (*big.Int, error)

	// F calculates the nth element of the fibonacci series.
	F(n int) (*big.Int, error)
}

// DB schema
// Fn => Value
// Value => Fn
