package cache

// New ...
func New(namespace string, co Options) *UnifiedCache {
	uc := &UnifiedCache{
		namespace: namespace,
		L1:        co.L1,
		L2:        co.L2,
	}
	return uc
}
