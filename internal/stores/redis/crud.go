package redisstore

import "github.com/redis/go-redis/v9"

// Get ...
func (r *RedisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// Set ...
func (r *RedisCache) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, r.expiration).Err()
}

// Delete ...
func (r *RedisCache) Delete(key string) (int, error) {
	n, err := r.client.Del(r.ctx, key).Result()
	return int(n), err
}
