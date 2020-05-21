package data

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

//Add redis加文本
func Add(a string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "text", a)
	return err
}

//Add2 第二选择后文本
func Add2(a string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "text2", a)
	return err
}

//Add3 第三选择后文本
func Add3(a string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "text3", a)
	return err
}

//Htext 把全部文本输出
func Htext() (string, error) {
	conn := pool.Get()
	defer conn.Close()
	ss, err := redis.Strings(conn.Do("mget", "text", "text2", "text3"))
	var text string
	for _, tr := range ss {
		text = text + tr
	}
	return text, err
}
