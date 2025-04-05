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

package cache

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gzylg/kits/errs"
)

// DefaultEvery sets a timer for how often to recycle the expired cache items in memory (in seconds)
var DefaultEvery = 60 // 1 minute

// MemoryItem stores memory cache item.
type MemoryItem struct {
	val         any
	createdTime time.Time
	lifespan    time.Duration
}

func (mi *MemoryItem) isExpire() bool {
	// 0 means forever
	if mi.lifespan == 0 {
		return false
	}
	return time.Since(mi.createdTime) > mi.lifespan
}

// MemoryCache is a memory cache adapter.
// Contains a RW locker for safe map storage.
type MemoryCache struct {
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
	Every int // run an expiration check Every clock time
}

// NewMemoryCache returns a new MemoryCache.
func NewMemoryCache() Cache {
	cache := MemoryCache{items: make(map[string]*MemoryItem)}
	return &cache
}

// Get returns cache from memory.
// If non-existent or expired, return nil.
func (bc *MemoryCache) Get(key string) (any, error) {
	bc.RLock()
	defer bc.RUnlock()
	if itm, ok := bc.items[key]; ok {
		if itm.isExpire() {
			return nil, errs.New("the key is expired")
		}
		return itm.val, nil
	}
	return nil, errs.New("the key isn't exist")
}

func (bc *MemoryCache) GetString(key string) (string, error) {
	v, err := bc.Get(key)
	if err != nil {
		return "", err
	}

	b, err := GetBytes(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (bc *MemoryCache) GetBytes(key string) ([]byte, error) {
	v, err := bc.Get(key)
	if err != nil {
		return nil, err
	}

	return GetBytes(v)
}

func (bc *MemoryCache) GetStruct(key string, s any) error {
	// vv := reflect.ValueOf(s)
	// if vv.Kind() != reflect.Ptr {
	// 	return errs.New(`parameter 's' must be a pointer.`)
	// } else {
	// 	vvv := reflect.ValueOf(vv.Elem().Interface())
	// 	if vvv.Kind() != reflect.Struct {
	// 		return errs.New(`parameter 's' must be a struct.`)
	// 	}
	// }

	v, err := bc.Get(key)
	if err != nil {
		return err
	}

	b, err := GetBytes(v)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, s)
}

// GetMulti gets caches from memory.
// If non-existent or expired, return nil.
func (bc *MemoryCache) GetMulti(keys []string) ([]any, error) {
	rc := make([]any, len(keys))
	keysErr := make([]string, 0)

	for i, ki := range keys {
		val, err := bc.Get(ki)
		if err != nil {
			keysErr = append(keysErr, fmt.Sprintf("key [%s] error: %s", ki, err.Error()))
			continue
		}
		rc[i] = val
	}

	if len(keysErr) == 0 {
		return rc, nil
	}
	return rc, errs.New(strings.Join(keysErr, "; "))
}

// Put puts cache into memory.
// If lifespan is 0, it will never overwrite this value unless restarted
func (bc *MemoryCache) Put(key string, val any, timeout ...any) error {
	bc.Lock()
	defer bc.Unlock()

	bc.items[key] = &MemoryItem{
		val:         val,
		createdTime: time.Now(),
		// lifespan:    timeout,
	}
	switch len(timeout) {
	case 1:
		bc.items[key].lifespan = timeout[0].(time.Duration)
	default:
		bc.items[key].lifespan = 0
	}
	return nil
}

func (bc *MemoryCache) PutJSON(key string, val any, timeout ...any) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return bc.Put(key, b, timeout...)
}

// Delete cache in memory.
// If the key is not found, it will not return error
func (bc *MemoryCache) Delete(key string) error {
	bc.Lock()
	defer bc.Unlock()
	delete(bc.items, key)
	return nil
}

// Incr increases cache counter in memory.
// Supports int,int32,int64,uint,uint32,uint64.
func (bc *MemoryCache) Incr(key string) error {
	bc.Lock()
	defer bc.Unlock()
	itm, ok := bc.items[key]
	if !ok {
		return errs.New("the key isn't exist")
	}

	val, err := incr(itm.val)
	if err != nil {
		return err
	}
	itm.val = val
	return nil
}

// Decr decreases counter in memory.
func (bc *MemoryCache) Decr(key string) error {
	bc.Lock()
	defer bc.Unlock()
	itm, ok := bc.items[key]
	if !ok {
		return errs.New("the key isn't exist")
	}

	val, err := decr(itm.val)
	if err != nil {
		return err
	}
	itm.val = val
	return nil
}

// IsExist checks if cache exists in memory.
func (bc *MemoryCache) IsExist(key string) (bool, error) {
	bc.RLock()
	defer bc.RUnlock()
	if v, ok := bc.items[key]; ok {
		return !v.isExpire(), nil
	}
	return false, nil
}

// TTL 查询key到期剩余时间
// 如果不存在或查询错误返回-1以及error
// 如果长期有效则返回-1以及error=nil
// 如果非长期有效则返回到期剩余时间，单位秒
func (bc *MemoryCache) TTL(key string) (int64, error) {
	//* 先检查是否存在
	isExist, err := bc.IsExist(key)
	if err != nil || !isExist { // 如果查询出错，或不存在
		return -1, errs.New("key is not exist") // 统一返回不存在错误
	}

	//* 查询距离过期剩余时间
	lifespan := int64(bc.items[key].lifespan / 1000000000) // 有效期
	if lifespan == 0 {
		return -1, nil // 如果长期有效则返回-1以及error=nil
	}
	createTime := bc.items[key].createdTime.Unix() // 创建时间
	nowTime := time.Now().Unix()                   // 当前时间
	t := lifespan - (nowTime - createTime)
	if t < 0 {
		return -1, errs.New("key is not exist") // 如果剩余时间小于0，返回不存在
	}
	return t, nil
}

func (bc *MemoryCache) Expire(key string, timeout time.Duration) error {
	//* 先检查是否存在
	isExist, err := bc.IsExist(key)
	if err != nil || !isExist { // 如果查询出错，或不存在
		return errs.New("key is not exist") // 统一返回不存在错误
	}

	bc.items[key].lifespan = timeout
	bc.items[key].createdTime = time.Now()
	return nil
}

// ClearAll deletes all cache in memory.
func (bc *MemoryCache) ClearAll() error {
	bc.Lock()
	defer bc.Unlock()
	bc.items = make(map[string]*MemoryItem)
	return nil
}

func (bc *MemoryCache) Scan(pattern string) (keys []string, err error) {
	return []string{}, nil
}

// StartAndGC starts memory cache. Checks expiration in every clock time.
func (bc *MemoryCache) StartAndGC(config string) error {
	var cf map[string]int
	if err := json.Unmarshal([]byte(config), &cf); err != nil {
		return errs.New("invalid config, please check your input: " + config)
	}
	if _, ok := cf["interval"]; !ok {
		cf = make(map[string]int)
		cf["interval"] = DefaultEvery
	}
	dur := time.Duration(cf["interval"]) * time.Second
	bc.Every = cf["interval"]
	bc.dur = dur
	go bc.vacuum()
	return nil
}

func (bc *MemoryCache) Close() error {
	return nil
}

// check expiration.
func (bc *MemoryCache) vacuum() {
	bc.RLock()
	every := bc.Every
	bc.RUnlock()

	if every < 1 {
		return
	}
	for {
		<-time.After(bc.dur)
		bc.RLock()
		if bc.items == nil {
			bc.RUnlock()
			return
		}
		bc.RUnlock()
		if keys := bc.expiredKeys(); len(keys) != 0 {
			bc.clearItems(keys)
		}
	}
}

// expiredKeys returns keys list which are expired.
func (bc *MemoryCache) expiredKeys() (keys []string) {
	bc.RLock()
	defer bc.RUnlock()
	for key, itm := range bc.items {
		if itm.isExpire() {
			keys = append(keys, key)
		}
	}
	return
}

// ClearItems removes all items who's key is in keys
func (bc *MemoryCache) clearItems(keys []string) {
	bc.Lock()
	defer bc.Unlock()
	for _, key := range keys {
		delete(bc.items, key)
	}
}

func init() {
	Register("memory", NewMemoryCache)
}
