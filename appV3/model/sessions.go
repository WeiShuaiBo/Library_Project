package model

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

// var store = cookie.NewStore([]byte("晴天阿良永远不阴天"))
var sessionName = "session-name"
var store sessions.Store

var redisClient *redis.Client

func init() {
	//store, _ = sessionRedis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	redisClient = CreateRedisClient()
}

/*
session, err := store.Get(c.Request, sessionName)

	if err != nil {
	    // 处理获取会话对象时的错误
	    log.Println("无法获取会话对象:", err)
	    // 可以选择执行其他恢复操作或返回特定错误提示给用户
	    return
	}

// 使用会话对象进行读取和写入会话数据
session.Values["id"] = 123
session.Values["name"] = "John"

// 保存会话数据
err = sessions.Save(c.Request, c.Writer)

	if err != nil {
	    // 处理保存会话数据时的错误
	    log.Println("无法保存会话数据:", err)
	    // 可以选择执行其他恢复操作或返回特定错误提示给用户
	    return
	}
*/
func GetSession(sessionID string) (map[string]interface{}, error) {
	//session, _ := store.Get(c.Request, sessionName)
	//sessionID := session.ID
	//fmt.Println(sessionID)
	values, err := redisClient.Get(context.Background(), sessionID).Bytes()
	if err != nil {
		fmt.Println("getsession1")
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(values, &data)
	if err != nil {
		fmt.Println("getsession2")

		return nil, err
	}

	return data, nil
}

func GenerateSessionID() (string, error) {
	// 生成一些随机字节
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 使用base64编码将字节转换为字符串
	sessionID := base64.URLEncoding.EncodeToString(randomBytes)
	return sessionID, nil
}
func SetSession(sessionID string, values map[string]interface{}, expiration time.Duration) error {
	//ctx := context.TODO()
	//fmt.Println(ctx)
	jsonValues, err := json.Marshal(values) //将values转为json
	if err != nil {
		return err
	}
	fmt.Println(sessionID)
	err = redisClient.Set(context.Background(), sessionID, jsonValues, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = 0
	return session.Save(c.Request, c.Writer)
}
