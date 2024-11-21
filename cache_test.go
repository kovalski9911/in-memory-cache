package inmemorycache

import (
	"testing"
)

func TestCacheSet(t *testing.T) {
	tests := []struct {
		k        string
		v        any
		expected any
	}{
		{"name", "Anton", "Anton"},
		{"name", 1, 2},
	}

	cache := New[any]()

	for _, test := range tests {
		cache.Set(test.k, test.v)
		result, ok := cache.storage[test.k]

		if !ok || result != test.expected {
			t.Errorf("Set of key %s and value %s is not expected %s", test.k, test.v, test.expected)
		}
	}
}
