package fibonacci

import (
	"log"
	"math/big"
	"time"
)

type Fibonacci2 struct {
	Phi  *big.Float
	Psi  *big.Float
	sqr5 *big.Float
}

func New2() (*Fibonacci2, error) {
	f := new(Fibonacci2)
	f.Phi = big.NewFloat(0).Quo(big.NewFloat(0).Add(big.NewFloat(1), big.NewFloat(0).Sqrt(big.NewFloat(5))), big.NewFloat(2))
	f.Psi = big.NewFloat(1).Sub(big.NewFloat(1), f.Phi)
	f.sqr5 = new(big.Float).Sqrt(big.NewFloat(5))
	log.Println(f.Phi.String())
	log.Println(f.Psi.String())
	return f, nil
}

func (f *Fibonacci2) N2F(n uint64) *big.Float {
	t := time.Now()
	res := new(big.Float).Quo(f.power(f.Phi, n), f.sqr5)
	ires, acc := res.Int(new(big.Int))
	m := time.Since(t)

	log.Println("Res", res.Text('g', 20))
	log.Println("iRes", ires.String())
	log.Println("acc", acc.String())
	log.Println("took ns", m.Nanoseconds())
	return res
}

func (f *Fibonacci2) power(v *big.Float, n uint64) *big.Float {
	result := new(big.Float).Copy(v)
	for i := uint64(0); i < n-1; i++ {
		result = new(big.Float).Mul(result, v)
	}
	return result
}
