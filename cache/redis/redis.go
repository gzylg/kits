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

// Package redis for cache provider
//
// depend on github.com/gomodule/redigo/redis
//
// go install github.com/gomodule/redigo/redis
//
// Usage:
// import(
//   _ "github.com/beego/beego/v2/client/cache/redis"
//   "github.com/beego/beego/v2/client/cache"
// )
//
//  bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:11211"}`)
//
//  more docs http://beego.vip/docs/module/cache.md
package redis

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gzylg/kits/cache"
	"github.com/gzylg/kits/errs"
)

// DefaultKey defines the collection name of redis for the cache adapter.
var DefaultKey = "Y-Server"

// Cache is Redis cache adapter.
type Cache struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
	maxIdle  int

	// Timeout value (less than the redis server's timeout value)
	timeout time.Duration
}

// NewRedisCache creates a new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Cache{key: DefaultKey}
}

// Execute the redis commands. args[0] must be the key name
func (rc *Cache) do(commandName string, args ...any) (any, error) {
	args[0] = rc.associate(args[0])
	c := rc.p.Get()
	defer func() {
		_ = c.Close()
	}()

	reply, err := c.Do(commandName, args...)
	if err != nil {
		return nil, errs.New("could not execute this command: " + commandName)
	}

	return reply, nil
}

// associate with config key.
func (rc *Cache) associate(originKey any) string {
	return fmt.Sprintf("%s:%s", rc.key, originKey)
}

// Get cache from redis.
func (rc *Cache) Get(key string) (any, error) {
	if v, err := rc.do("GET", key); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}

func (rc *Cache) GetString(key string) (string, error) {
	v, err := rc.Get(key)
	if err != nil {
		return "", err
	}

	b, err := cache.GetBytes(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (rc *Cache) GetBytes(key string) ([]byte, error) {
	v, err := rc.Get(key)
	if err != nil {
		return nil, err
	}

	return cache.GetBytes(v)
}

func (rc *Cache) GetStruct(key string, s any) error {
	vv := reflect.ValueOf(s)
	if vv.Kind() != reflect.Ptr {
		return errs.New(`parameter 's' must be a pointer.`)
	} else {
		vvv := reflect.ValueOf(vv.Elem().Interface())
		if vvv.Kind() != reflect.Struct {
			return errs.New(`parameter 's' must be a struct.`)
		}
	}

	v, err := rc.Get(key)
	if err != nil {
		return err
	}

	b, err := cache.GetBytes(v)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, s)
}

// GetMulti gets cache from redis.
func (rc *Cache) GetMulti(keys []string) ([]any, error) {
	c := rc.p.Get()
	defer func() {
		_ = c.Close()
	}()
	var args []any
	for _, key := range keys {
		args = append(args, rc.associate(key))
	}
	return redis.Values(c.Do("MGET", args...))
}

// Put puts cache into redis.
func (rc *Cache) Put(key string, val any, timeout ...any) error {
	var err error
	switch len(timeout) {
	case 1:
		_, err = rc.do("SETEX", key, int64(timeout[0].(time.Duration)/time.Second), val)
	default:
		_, err = rc.do("SET", key, val)
	}
	return err
}

func (rc *Cache) PutJSON(key string, val any, timeout ...any) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return rc.Put(key, b, timeout...)
}

// Delete deletes a key's cache in redis.
func (rc *Cache) Delete(key string) error {
	_, err := rc.do("DEL", key)
	return err
}

// IsExist checks cache's existence in redis.
func (rc *Cache) IsExist(key string) (bool, error) {
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return v, nil
}

// TTL 查询key到期剩余时间
// 如果不存在或查询错误返回-1以及error
// 如果长期有效则返回-1以及error=nil
// 如果非长期有效则返回到期剩余时间，单位秒
func (rc *Cache) TTL(key string) (int64, error) {
	//* 先检查是否存在
	isExist, err := rc.IsExist(key)
	if err != nil || !isExist { // 如果查询出错，或不存在
		return -1, errs.New("key is not exist") // 统一返回不存在错误
	}

	//* 查询距离过期剩余时间
	s, err := redis.Int64(rc.do("TTL", key))
	if err != nil {
		err = errs.New("key is not exist") // 如果出错，返回不存在
	}

	return s, err
}

// Expire 更新过期时间
func (rc *Cache) Expire(key string, timeout time.Duration) error {
	//* 先检查是否存在
	isExist, err := rc.IsExist(key)
	if err != nil || !isExist { // 如果查询出错，或不存在
		return errs.New("key is not exist") // 统一返回不存在错误
	}

	//* 更新过期时间
	_, err = rc.do("EXPIRE", key, int64(timeout/time.Second))

	return err
}

// Incr increases a key's counter in redis.
func (rc *Cache) Incr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

// Decr decreases a key's counter in redis.
func (rc *Cache) Decr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

// ClearAll deletes all cache in the redis collection
// Be careful about this method, because it scans all keys and the delete them one by one
func (rc *Cache) ClearAll() error {
	cachedKeys, err := rc.Scan(rc.key + ":*")
	if err != nil {
		return err
	}
	c := rc.p.Get()
	defer func() {
		_ = c.Close()
	}()
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// Scan scans all keys matching a given pattern.
func (rc *Cache) Scan(pattern string) (keys []string, err error) {
	c := rc.p.Get()
	defer func() {
		_ = c.Close()
	}()
	var (
		cursor uint64 = 0 // start
		result []any
		list   []string
	)
	for {
		result, err = redis.Values(c.Do("SCAN", cursor, "MATCH", pattern, "COUNT", 1024))
		if err != nil {
			return
		}
		list, err = redis.Strings(result[1], nil)
		if err != nil {
			return
		}
		keys = append(keys, list...)
		cursor, err = redis.Uint64(result[0], nil)
		if err != nil {
			return
		}
		if cursor == 0 { // over
			return
		}
	}
}

// StartAndGC starts the redis cache adapter.
// config: must be in this format {"key":"collection key","conn":"connection info","dbNum":"0"}
// Cached items in redis are stored forever, no garbage collection happens
func (rc *Cache) StartAndGC(config string) error {
	var cf map[string]string
	err := json.Unmarshal([]byte(config), &cf)
	if err != nil {
		return errs.New("could not unmarshal the config: " + config)
	}

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errs.New("config missing conn field: " + config)
	}

	// Format redis://<password>@<host>:<port>
	cf["conn"] = strings.Replace(cf["conn"], "redis://", "", 1)
	if i := strings.Index(cf["conn"], "@"); i > -1 {
		cf["password"] = cf["conn"][0:i]
		cf["conn"] = cf["conn"][i+1:]
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	if _, ok := cf["maxIdle"]; !ok {
		cf["maxIdle"] = "3"
	}
	if _, ok := cf["timeout"]; !ok {
		cf["timeout"] = "180s"
	}
	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]
	rc.maxIdle, _ = strconv.Atoi(cf["maxIdle"])

	if v, err := time.ParseDuration(cf["timeout"]); err == nil {
		rc.timeout = v
	} else {
		rc.timeout = 180 * time.Second
	}

	rc.connectInit()

	c := rc.p.Get()
	defer func() {
		_ = c.Close()
	}()

	// test connection
	if err = c.Err(); err != nil {
		return errs.New(
			"can not connect to remote redis server, please check the connection info and network state: " + config)
	}
	return nil
}

func (rc *Cache) Close() error {
	return rc.p.Close()
}

// connect to redis.
func (rc *Cache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.conninfo)
		if err != nil {
			return nil, errs.New(
				"could not dial to remote server: " + rc.conninfo)
		}

		if rc.password != "" {
			if _, err = c.Do("AUTH", rc.password); err != nil {
				_ = c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			_ = c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     rc.maxIdle,
		IdleTimeout: rc.timeout,
		Dial:        dialFunc,
	}
}

func init() {
	cache.Register("redis", NewRedisCache)
}
