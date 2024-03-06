package fibonacci

type Option func(*Fibonacci) error

func OptionWithCaching(v bool) Option {
	return func(f *Fibonacci) error {
		f.config.caching = v
		return nil
	}
}
