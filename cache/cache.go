// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cache provide a Cache interface and some implement engine
// Usage:
//
// import(
//   "github.com/beego/beego/v2/client/cache"
// )
//
// bm, err := cache.NewCache("memory", `{"interval":60}`)
//
// Use it like this:
//
//	bm.Put("astaxie", 1, 10 * time.Second)
//	bm.Get("astaxie")
//	bm.IsExist("astaxie")
//	bm.Delete("astaxie")
//
//  more docs http://beego.vip/docs/module/cache.md
package cache

import (
	"time"

	"github.com/gzylg/kits/errs"
)

// Cache interface contains all behaviors for cache adapter.
// usage:
//	cache.Register("file",cache.NewFileCache) // this operation is run in init method of file.go.
//	c,err := cache.NewCache("file","{....}")
//	c.Put("key",value, 3600 * time.Second)
//	v := c.Get("key")
//
//	c.Incr("counter")  // now is 1
//	c.Incr("counter")  // now is 2
//	count := c.Get("counter").(int)
type Cache interface {
	Getter
	Putter

	// Delete cached value by key.
	// Should not return error if key not found
	Delete(key string) error
	// Increment a cached int value by key, as a counter.
	Incr(key string) error
	// Decrement a cached int value by key, as a counter.
	Decr(key string) error
	// Check if a cached value exists or not.
	// if key is expired, return (false, nil)
	IsExist(key string) (bool, error)

	// 获取剩余到期时间，单位秒
	TTL(key string) (int64, error)
	// 更新过期时间
	Expire(key string, timeout time.Duration) error
	Scan(pattern string) ([]string, error)

	// Clear all cache.
	ClearAll() error
	// Start gc routine based on config string settings.
	StartAndGC(config string) error

	Close() error
}

type Getter interface {
	// Get a cached value by key.
	Get(key string) (any, error)
	GetString(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	GetStruct(key string, s any) (err error)

	// GetMulti is a batch version of Get.
	GetMulti(keys []string) ([]any, error)
}

type Putter interface {
	// Set a cached value with key and expire time.
	Put(key string, val any, timeout ...any) error
	PutJSON(key string, val any, timeout ...any) error
}

// Instance is a function create a new Cache Instance
type Instance func() Cache

var adapters = make(map[string]Instance)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter Instance) {
	if adapter == nil {
		panic(errs.New("cache: Register adapter is nil"))
	}
	if _, ok := adapters[name]; ok {
		panic(errs.New("cache: Register called twice for adapter " + name))
	}
	adapters[name] = adapter
}

// NewCache creates a new cache driver by adapter name and config string.
// config: must be in JSON format such as {"interval":360}.
// Starts gc automatically.
func NewCache(adapterName, config string) (adapter Cache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = errs.New("cache: unknown adapter name " + adapterName + " (forgot to import?)")
		return
	}
	adapter = instanceFunc()
	err = adapter.StartAndGC(config)
	if err != nil {
		adapter = nil
	}
	return
}
