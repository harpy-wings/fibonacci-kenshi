package fibonacci

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		f, err := New(func(*Fibonacci) error { return nil })
		require.NoError(t, err)
		require.NotNil(t, f)
	})
	t.Run("Failure", func(t *testing.T) {
		t.Run("Option", func(t *testing.T) {
			f, err := New(func(*Fibonacci) error { return errors.New("any") })
			require.Error(t, err)
			require.Nil(t, f)
		})

		t.Run("Init", func(t *testing.T) {
			f, err := New(func(f *Fibonacci) error {
				f.config.maxBitSize = -1
				return nil
			})
			require.Error(t, err)
			require.Nil(t, f)
		})
	})
}

func TestF(t *testing.T) {
	f, err := New(OptionWithCaching(true))
	require.NoError(t, err)
	require.NotNil(t, f)

	t.Run("Success", func(t *testing.T) {
		var testcases = []int{0, 1, 2}
		for i := 0; i < 10; i++ {
			testcases = append(testcases, rand.Intn(50000))
		}
		for _, v := range testcases {
			t.Run(fmt.Sprintf("case : %d", v), func(t *testing.T) {
				fv, err := f.F(v)
				require.NoError(t, err)
				require.NotNil(t, fv)
				// t.Log(fv.String())
			})
		}
	})
	t.Run("Failure", func(t *testing.T) {
		t.Run("InvalidNumber", func(t *testing.T) {
			fv, err := f.F(-1)
			require.Error(t, err)
			require.Nil(t, fv)
		})
	})
}

func TestNext(t *testing.T) {
	f, err := New(OptionWithCaching(true), OptionWithMaxBitSize(300))
	require.NoError(t, err)
	require.NotNil(t, f)

	t.Run("Success", func(t *testing.T) {
		var testCases = map[string]string{
			"0":  "1",
			"1":  "2", // referring to our conversation.
			"13": "21",
			"32423247527351544763402471792982538876233052554697128188128597": "52461916524905785334311649958648296484733611329035169538240802",
		}

		for k, v := range testCases {
			t.Run(fmt.Sprintf("Case %s", k), func(t *testing.T) {
				TC := new(big.Int)
				err = TC.UnmarshalText([]byte(k))
				require.NoError(t, err)
				RC, err := f.Next(TC)
				require.NoError(t, err)
				require.Equal(t, v, RC.String())
			})
		}
	})

	t.Run("Failure", func(t *testing.T) {
		t.Run("Invalid Number", func(t *testing.T) {
			TC := new(big.Int)
			err := TC.UnmarshalText([]byte("324232475273515447634024717929825388762330525546971281881285979"))
			require.NoError(t, err)
			RC, err := f.Next(TC)
			require.Error(t, err)
			require.Nil(t, RC)
		})
		t.Run("Negative Number", func(t *testing.T) {
			TC := new(big.Int)
			err := TC.UnmarshalText([]byte("-3242324752735"))
			require.NoError(t, err)
			RC, err := f.Next(TC)
			require.Error(t, err)
			require.Nil(t, RC)
		})

		t.Run("Out Of Range Number", func(t *testing.T) {
			TC := new(big.Int)
			err := TC.UnmarshalText([]byte("86168291600238450732788312165664788095941068326060883324529903470149056115823592713458328176574447204501"))
			require.NoError(t, err)
			RC, err := f.Next(TC)
			require.Error(t, err)
			require.Nil(t, RC)
		})

	})
}
