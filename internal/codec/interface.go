package codec

import (
	"encoding/json"
	"math/big"
)

type Codec interface {
	Encode(n *big.Int) ([]byte, error)
	Decode([]byte) (*big.Int, error)
}

type codec struct {
	kind codecKind
}

func (c *codec) Encode(n *big.Int) ([]byte, error) {
	bs, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
func (c *codec) Decode(bs []byte) (*big.Int, error) {
	v := new(big.Int)
	err := v.UnmarshalJSON(bs)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type codecKind int

const (
	codecKindInvalid = iota
	codecKindBinary
	codecKindJSON
)

func New(ops ...Option) (Codec, error) {
	return &codec{kind: codecKindJSON}, nil
}
