package cache

import (
	"errors"
)

// Get ...
func (u *UnifiedCache) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	if l1Value, err := u.L1.Get(key); err != nil {
		return l1Value, err
	}
	if l2Value, err := u.L2.Get(u.nsKey(key)); err != nil {
		return l2Value, err
	}
	return "", errors.New("key not found in both caches")
}

// Set ...
func (u *UnifiedCache) Set(key string, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if value == "" {
		return errors.New("value cannot be empty")
	}

	if err := u.L1.Set(key, value); err != nil {
		return err
	}
	return u.L2.Set(u.nsKey(key), value)
}

// Delete ...
func (u *UnifiedCache) Delete(key string) (int, error) {
	if key == "" {
		return 0, errors.New("key cannot be empty")
	}

	l1deleted, err := u.L1.Delete(key)
	if err != nil {
		return l1deleted, err
	}
	l2Deleted, err := u.L2.Delete(u.nsKey(key))
	if err != nil {
		return l2Deleted, err
	}
	if l1deleted == 0 && l2Deleted == 0 {
		return 0, errors.New("key not found in both caches")
	}

	return max(l1deleted, l2Deleted), nil
}

// DeleteMany ...
func (u *UnifiedCache) DeleteMany(keys []string) (int, error) {
	if len(keys) == 0 {
		return 0, errors.New("keys cannot be empty")
	}

	l1deleted, err := u.L1.DeleteMany(keys)
	if err != nil {
		return l1deleted, err
	}
	l2Deleted, err := u.L2.DeleteMany(u.nsKeys(keys))
	if err != nil {
		return l2Deleted, err
	}
	if l1deleted == 0 && l2Deleted == 0 {
		return 0, errors.New("keys not found in both caches")
	}

	return max(l1deleted, l2Deleted), nil
}

// Exists ...
func (u *UnifiedCache) Exists(key string) bool {
	if key == "" {
		return false
	}

	if u.L1.Exists(key) {
		return true
	}
	if u.L2.Exists(u.nsKey(key)) {
		return true
	}
	return false
}

// Pop ...
func (u *UnifiedCache) Pop(key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	l1Value, err := u.L1.Pop(key)
	if err != nil {
		return l1Value, err
	}
	l2Value, err := u.L2.Pop(u.nsKey(key))
	if err != nil {
		return l2Value, err
	}
	return "", errors.New("key not found in both caches")
}

// Flush ...
func (u *UnifiedCache) Flush() error {
	if err := u.L1.Flush(); err != nil {
		return err
	}
	if err := u.L2.Flush(); err != nil {
		return err
	}
	return nil
}

// Disconnect ...
func (u *UnifiedCache) Disconnect() error {
	if err := u.L2.Disconnect(); err != nil {
		return err
	}
	return nil
}

// Ping ...
func (u *UnifiedCache) Ping() error {
	if err := u.L2.Ping(); err != nil {
		return err
	}
	return nil
}

// GetOrSet ...
func (u *UnifiedCache) GetOrSet(key string, valueFunc func() (string, error)) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	if l1Value, err := u.L1.Get(key); err == nil {
		return l1Value, nil
	}

	if l2Value, err := u.L2.Get(u.nsKey(key)); err == nil {
		return l2Value, nil
	}


	value, err := valueFunc()
	if err != nil {
		return "", err
	}

	if err := u.Set(key, value); err != nil {
		return "", err
	}
	return value, nil
}
