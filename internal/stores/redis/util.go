package redisstore

func (r *RedisCache) Ping() error {
	_, err := r.client.Ping(r.ctx).Result()
	return err
}

func (r *RedisCache) Disconnect() error {
	return r.client.Close()
}
