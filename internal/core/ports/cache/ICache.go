package cache

import "time"

type ICache interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, duration time.Duration) error
	Delete(key string) error
}
