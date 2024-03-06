package fibonacci

import (
	"math/big"
	"strconv"

	"github.com/patrickmn/go-cache"
)

type Fibonacci struct {
	db     *cache.Cache
	config struct {
		a0          *big.Int
		a1          *big.Int
		caching     bool
		maxMemoryKB uint64
	}
}

func New(ops ...Option) (*Fibonacci, error) {
	f := new(Fibonacci)
	f.loadDefaults()

	for _, fn := range ops {
		err := fn(f)
		if err != nil {
			return nil, err
		}
	}

	err := f.init()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *Fibonacci) f(n int) *big.Int {
	var res *big.Int
	if f.config.caching {
		if v, ok := f.db.Get(f.nthKey(n)); ok {
			// exist in cache
			return v.(*big.Int)
		} // else is not exist in cache

	}

	// f(n) = f(n-1) + f(n-2)
	res = new(big.Int).Add(f.f(n-1), f.f(n-2))

	// store in cache
	if f.config.caching { // ok = false since it skipped that step
		f.db.Set(f.nthKey(n), res, cache.NoExpiration)
	}

	return res
}

func (f *Fibonacci) F(n int) *big.Int {
	if n == 0 {
		return f.config.a0
	} else if n == 1 {
		return f.config.a1
	}
	return f.f(n)
}

func (f *Fibonacci) loadDefaults() {
	f.config.a0 = defaultA0
	f.config.a1 = defaultA1
}

func (f *Fibonacci) init() error {
	if f.config.caching {
		f.db = cache.New(cache.NoExpiration, cache.NoExpiration)
		f.db.Set(f.nthKey(0), f.config.a0, cache.NoExpiration)
		f.db.Set(f.nthKey(1), f.config.a1, cache.NoExpiration)
	}
	return nil
}

func (f *Fibonacci) nthKey(n int) string {
	return keyPrefix + strconv.Itoa(n)
}
