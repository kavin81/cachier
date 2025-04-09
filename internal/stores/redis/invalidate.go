package redisstore

import "github.com/redis/go-redis/v9"

// DeleteMany ...
func (r *RedisCache) DeleteMany(keys []string) (int, error) {
	n, err := r.client.Del(r.ctx, keys...).Result()
	return int(n), err
}

// Exists ...
func (r *RedisCache) Exists(key string) bool {
	n, err := r.client.Exists(r.ctx, key).Result()
	return err == nil && n > 0
}

// Pop ...
func (r *RedisCache) Pop(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	_, err = r.client.Del(r.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Flush ...
func (r *RedisCache) Flush() error {
	return r.client.FlushDB(r.ctx).Err()
}
