package cache

type Cache interface {
	Set(key string, val string) bool
	Get(key string) (string, error)
}
