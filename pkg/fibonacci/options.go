package fibonacci

type Option func(*Fibonacci) error

func OptionWithCaching(v bool) Option {
	return func(f *Fibonacci) error {
		f.config.caching = v
		return nil
	}
}

func OptionWithMaxBitSize(v int) Option {
	return func(f *Fibonacci) error {
		if v < 0 {
			return ErrInvalidBitSize
		}
		f.config.maxBitSize = v
		return nil
	}
}
