package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func GetSession(sessionID string) (map[string]interface{}, error) {
	//session, _ := store.Get(c.Request, sessionName)
	//sessionID := session.ID
	//fmt.Println(sessionID)
	values, err := redisClient.Get(sessionID).Bytes()
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
	err = redisClient.Set(sessionID, jsonValues, expiration).Err()
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
