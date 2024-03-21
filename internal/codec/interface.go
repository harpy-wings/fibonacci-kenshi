package codec

import (
	"encoding/base64"
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

	switch c.kind {
	case codecKindJSON:
		bs, err := json.Marshal(struct {
			N string `json:"n"`
		}{N: n.String()})
		if err != nil {
			return nil, err
		}
		return bs, nil
	case codecKindBinary:
		return []byte(base64.RawStdEncoding.EncodeToString(n.Bytes())), nil
	default:
		return nil, ErrInvalidCodec
	}

}
func (c *codec) Decode(bs []byte) (*big.Int, error) {
	switch c.kind {
	case codecKindJSON:
		var data struct {
			N string `json:"n"`
		}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return nil, err
		}
		v := new(big.Int)
		err = v.UnmarshalJSON([]byte(data.N))
		if err != nil {
			return nil, err
		}
		return v, nil
	case codecKindBinary:
		bss, err := base64.RawStdEncoding.DecodeString(string(bs))
		if err != nil {
			return nil, err
		}
		V := new(big.Int).SetBytes(bss)
		return V, nil
	default:
		return nil, ErrInvalidCodec
	}

}

type codecKind int

const (
	codecKindInvalid = iota
	codecKindBinary
	codecKindJSON
)

func New(ops ...Option) (Codec, error) {
	c := new(codec)
	for _, fn := range ops {
		err := fn(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
