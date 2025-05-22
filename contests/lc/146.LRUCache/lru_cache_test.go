package _146_LRUCache

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if got := cache.Get(1); got != 1 {
		t.Errorf("expected Get(1) = 1, got %d", got)
	}

	cache.Put(3, 3) // вытеснит key=2

	if got := cache.Get(2); got != -1 {
		t.Errorf("expected Get(2) = -1 (evicted), got %d", got)
	}

	cache.Put(4, 4) // вытеснит key=1

	if got := cache.Get(1); got != -1 {
		t.Errorf("expected Get(1) = -1 (evicted), got %d", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("expected Get(3) = 3, got %d", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("expected Get(4) = 4, got %d", got)
	}
}
