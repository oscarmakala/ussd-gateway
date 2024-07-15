package ports

type RedisPort[T any] interface {
	Get(Key string) (T, error)
	Set(Key string, value T) error
}
