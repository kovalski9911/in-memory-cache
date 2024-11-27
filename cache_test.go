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
		{"name", 1, 1},
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

func TestCacheGet(t *testing.T) {
	tests := []struct {
		k        string
		v        any
		expected any
	}{
		{"name", "Anton", "Anton"},
		{"name", 1, 1},
		{"", "", ""},
		{"nonexistentKey", nil, nil},
	}

	cache := New[any]()

	for _, test := range tests {
		cache.storage[test.k] = test.v

		result := cache.Get(test.k)

		if result != test.expected {
			t.Errorf("Set of key %s and value %s is not expected %s", test.k, test.v, test.expected)
		}
	}
}

func TestCacheDelete(t *testing.T) {
	tests := []struct {
		k        string
		v        any
		expected any
	}{
		{"name", "Anton", nil},
		{"name", 1, nil},
	}

	cache := New[any]()

	for _, test := range tests {
		cache.storage[test.k] = test.v
		cache.Delete(test.k)

		_, ok := cache.storage[test.k]

		if ok {
			t.Errorf("Key %s was not deleted as expected", test.k)
		}
	}
}
