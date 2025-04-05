package redis

import (
	"log"
	"testing"
	"time"

	"github.com/gzylg/kits/cache"
)

func TestRedis(t *testing.T) {
	//* 新建 redis 缓存
	rds, err := cache.NewCache("redis", `{"key":"ySrv-test","conn":":6379","dbNum":"0","password":""}`)
	if err != nil {
		t.Fatal(err)
		return
	}

	//* 新建 memory 缓存
	m, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		t.Fatal(err)
		return
	}

	a := &struct{ Id int }{Id: 1234566789}

	//* put
	log.Println("redis put:", rds.PutJSON("testKey", a, 30*time.Second))
	log.Println("memory put:", m.PutJSON("testKey", a, 30*time.Second))
	log.Println("=============================")

	//* get
	b := &struct{ Id int }{}
	log.Println("redis get:", rds.GetStruct("testKey", b), b)
	c := &struct{ Id int }{}
	log.Println("memory get:", m.GetStruct("testKey", c), c)
	log.Println("=============================")

	//* 获取剩余到期时间
	time.Sleep(3 * time.Second)
	s, err := rds.TTL("testKey")
	log.Println("redis TTL:", s, err)
	s, err = m.TTL("testKey")
	log.Println("memory TTL:", s, err)
	log.Println("=============================")

	//* 设置到期时间
	log.Println("redis Expire:", rds.Expire("testKey", 60*time.Second))
	log.Println("memory Expire:", m.Expire("testKey", 60*time.Second))

	//* 重新获取剩余到期时间
	s, err = rds.TTL("testKey")
	log.Println("redis TTL:", s, err)
	s, err = m.TTL("testKey")
	log.Println("memory TTL:", s, err)

	//* 三秒后再次获取剩余到期时间
	time.Sleep(3 * time.Second)
	s, err = rds.TTL("testKey")
	log.Println("redis TTL:", s, err)
	s, err = m.TTL("testKey")
	log.Println("memory TTL:", s, err)
}
