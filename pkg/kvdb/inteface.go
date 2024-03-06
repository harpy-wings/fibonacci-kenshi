package kvdb

import (
	"context"

	gocache "github.com/patrickmn/go-cache"
)

type KVDB interface {
	// Get returns the value of key, Error if not found.
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
}

func hi() {

	gcClient := gocache.New(gocache.NoExpiration, gocache.NoExpiration)

}
