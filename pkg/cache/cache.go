package cache

import (
	"time"
)

type Cache[T any] interface {
	Set(key string, value T, expiration time.Duration) error
	Get(key string) (T, bool, error)
	Delete(key string) error
}
