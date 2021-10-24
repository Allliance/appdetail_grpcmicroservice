package main

import (
	"microservice-grpc/pkg/cache"
	"testing"
)

type data []byte

func (data1 data) isEqual(data2 []byte) bool {
	if len(data1) != len(data2) {
		return false
	}
	for i, val := range data1 {
		if val != data2[i] {
			return false
		}
	}
	return true
}

func TestCache(t *testing.T) {
	testCache, _ := cache.NewCache()
	if nonCachedData, _ := testCache.Retrieve("Some Cache"); nonCachedData != nil {
		t.Errorf("Cache not clean, Expected nil , found non-nil")
	}
	testCache.CacheApp("Some Cache", []byte("Some Data"))
	cacheData, err := testCache.Retrieve("Some Cache")
	if err != nil {
		t.Errorf("Error occured while retrieving cached data from cache")
	}
	if !data(cacheData).isEqual([]byte("Some Data")) {
		t.Errorf("Cached data has changed")
	}
}
