package cache

func (u *UnifiedCache) nsKey(key string) string {
	return u.namespace + ":" + key
}

func (u *UnifiedCache) nsKeys(keys []string) []string {
	namespacedKeys := make([]string, len(keys))
	for i, key := range keys {
		namespacedKeys[i] = u.nsKey(key)
	}
	return namespacedKeys
}
