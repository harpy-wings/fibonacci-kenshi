package fibonacci

import "math/big"

const (
	keyPrefix = "a"
)

var (
	defaultA0 *big.Int = big.NewInt(0)
	defaultA1 *big.Int = big.NewInt(1)
)
