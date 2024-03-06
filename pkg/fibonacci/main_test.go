package fibonacci

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	f, err := New(OptionWithCaching(true))
	require.NoError(t, err)
	tn := time.Now()
	v := f.F(100000)
	log.Println(v.String(), "in ns", time.Since(tn).Nanoseconds(), v.BitLen())

	tn = time.Now()
	v = f.F(500000)
	f.db.SaveFile("./data.txt")
	log.Println(v.String(), "in ns", time.Since(tn).Nanoseconds(), v.BitLen())

	_ = f
}

// 222232244629420445529739893461909967206666939096499764990979600 OK
// 222232244629422584762545790887617788825944041106761707346198528
