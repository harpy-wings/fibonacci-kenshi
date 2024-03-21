package fibonacci

import (
	"crypto/md5"
	"math/big"
	"strconv"

	"github.com/patrickmn/go-cache"
)

type Fibonacci struct {
	db     *cache.Cache
	config struct {
		a0         *big.Int
		a1         *big.Int
		caching    bool
		maxBitSize int
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

// f : calculate the nth element of the fibonacci series.
// Internal use only, no validation here.
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

	// TODO, maxBitSize validation skipped here.
	// f(int) *big.Int => f(int) (*big.Int,error)

	// store in cache
	if f.config.caching { // ok = false since it skipped that step
		f.db.Set(f.nthKey(n), res, cache.NoExpiration)
		f.db.Set(f.valKey(res), n, cache.NoExpiration)
	}

	return res
}

// F : calculate the nth element of the fibonacci series.
func (f *Fibonacci) F(n int) (*big.Int, error) {
	if n < 0 {
		return nil, ErrInvalidIndex
	}
	if n == 0 {
		return f.config.a0, nil
	} else if n == 1 {
		return f.config.a1, nil
	}
	return f.f(n), nil
}

// Next returns the next value of the fibonacci serries. F(IndexOf(n)+1)
// returns error
// ErrNotfound : if the n is not a valid fibonacci number
// ErrTooBig : if the n bit size is bigger than configuration.
// ErrInvalidNumber : if the n is less than zero.
func (f *Fibonacci) Next(n *big.Int) (*big.Int, error) {
	if n.Sign() < 0 {
		return nil, ErrInvalidNumber
	}
	if n.BitLen() > f.config.maxBitSize {
		return nil, ErrTooBig
	}
	index, err := f.indexOf(n)
	if err != nil {
		return nil, err
	}
	return f.F(index + 1)

}

// indexOf : returns the index of the given fibonacci number.
func (f *Fibonacci) indexOf(n *big.Int) (int, error) {
	if !f.config.caching {
		//todo, find a alternative approach
		return 0, ErrUnimplemented
	}
	v, found := f.db.Get(f.valKey(n))
	if !found {
		return 0, ErrNotfound
	}
	res, ok := v.(int)
	if !ok {
		return 0, ErrTypeAssertionFailed
	}

	return res, nil
}

// loadDefault : loads the default values.
func (f *Fibonacci) loadDefaults() {
	f.config.a0 = defaultA0
	f.config.a1 = defaultA1
	f.config.maxBitSize = defaultMaxBitSize
	f.config.caching = defaultCaching
}

// init : initialize the structure
func (f *Fibonacci) init() error {
	if f.config.caching {
		f.db = cache.New(cache.NoExpiration, cache.NoExpiration)
		f.db.Set(f.nthKey(0), f.config.a0, cache.NoExpiration)
		f.db.Set(f.nthKey(1), f.config.a1, cache.NoExpiration)
		f.db.Set(f.valKey(f.config.a0), int(0), cache.NoExpiration)
		f.db.Set(f.valKey(f.config.a1), int(1), cache.NoExpiration)
	}
	if f.config.maxBitSize < 0 {
		return ErrInvalidBitSize
	}

	// PreLoad
	index := 0
	for {
		fv, err := f.F(index)
		if err != nil {
			return err
		}
		if fv.BitLen() >= f.config.maxBitSize {
			break
		}
		index += 1
	}
	return nil
}

// nthKey convert the n to a unique key. n is int.
func (f *Fibonacci) nthKey(n int) string {
	return keyPrefix + strconv.Itoa(n)
}

// valKey : convert the "n" to a unique key. n is a big number.
func (f *Fibonacci) valKey(n *big.Int) string {
	// Hashing Applied to reduce the memory size and improve retrieval performance
	h := md5.New()
	_, err := h.Write(n.Bytes())
	if err != nil {
		panic(err)
	}
	return string(h.Sum(nil))
}
