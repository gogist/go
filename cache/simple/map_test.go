package cache

import (
	"testing"
	"time"
)

func TestExpiration(t *testing.T) {
	cache := NewMap(0)
	if cache.Expiration() != 0 {
		t.Error("Expect 0 but get", cache.Expiration())
	}
}

func TestContains(t *testing.T) {
	cache := NewMap(0)

	cache.Set("1", []byte("0"))
	if !cache.Contains("1") {
		t.Error("Expect containing 1 but missed")
	}

	cache.Set("2", []byte("0"))
	if !cache.Contains("2") {
		t.Error("Expect containing 2 but missed")
	}

	cache.Set("3", []byte("0"))
	if !cache.Contains("3") {
		t.Error("Expect containing 3 but missed")
	}

	if cache.Contains("4") {
		t.Error("Expect not containing 4 but got")
	}

	if cache.Contains("5") {
		t.Error("Expect not containing 5 but got")
	}
}

func TestSet(t *testing.T) {
	cache := NewMap(1 * time.Second)
	cache.Set("1", []byte("0"))
	if !cache.Contains("1") {
		t.Error("Expect containing 1 but missed")
	}
	bytes, _ := cache.Get("1")
	if string(bytes) != "0" {
		t.Errorf("Expect %s but get %s\n", "0", string(bytes))
	}
	time.Sleep(2 * time.Second)
	_, err := cache.Get("1")
	if err == nil {
		t.Errorf("Expect expired but not")
	}
}

func TestSetGet(t *testing.T) {
	cache := NewMap(60 * time.Second)
	cache.Set("1", []byte("1"))
	cache.Set("2", []byte("2"))
	cache.Set("3", []byte("3"))
	bytes, _ := cache.Get("1")
	if string(bytes) != "1" {
		t.Errorf("Expect %s but get %s\n", "1", string(bytes))
	}
	bytes, _ = cache.Get("2")
	if string(bytes) != "2" {
		t.Errorf("Expect %s but get %s\n", "2", string(bytes))
	}
	bytes, _ = cache.Get("3")
	if string(bytes) != "3" {
		t.Errorf("Expect %s but get %s\n", "3", string(bytes))
	}
}
