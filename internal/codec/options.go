package codec

type Option func(*codec) error

func OptionDescriber(v string) Option {
	return func(c *codec) error {

		return nil
	}
}
