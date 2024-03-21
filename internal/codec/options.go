package codec

type Option func(*codec) error

func OptionDescriber(v string) Option {
	return func(c *codec) error {
		switch v {
		case "application/json":
			c.kind = codecKindJSON
		case "application/binary":
			c.kind = codecKindBinary
		default:
			c.kind = codecKindInvalid
		}
		return nil
	}
}
