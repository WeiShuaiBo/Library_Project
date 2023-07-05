package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"time"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("香香编程喵喵喵"))
var sessionName = "session-name"

func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}

func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id
	return session.Save(c.Request, c.Writer)
}

func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = 0
	return session.Save(c.Request, c.Writer)
}

// NewPool 创建 Redis 连接池
func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
