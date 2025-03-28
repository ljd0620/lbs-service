package middlewares

import (
	"context"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type RedisMiddleware struct {
	client *redis.Client
}

func NewRedisMiddleware(addr string, password string, db int) *RedisMiddleware {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisMiddleware{client: rdb}
}

func (rm *RedisMiddleware) Get(ctx context.Context, key string) (string, error) {
	val, err := rm.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (rm *RedisMiddleware) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return rm.client.Set(ctx, key, value, expiration).Err()
}

func (rm *RedisMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}