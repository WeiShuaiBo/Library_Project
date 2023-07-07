package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// GinLogger 接收gin框架默认的日志
// @author: [ziop](https://github.com/wangpi26)
// @function: GinLogger
// @description: 接收gin框架默认的日志
// @param: c *gin.Context
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start := time.Now()
		//path := c.Request.URL.Path
		//requestId, _ := gonanoid.Nanoid()
		//c.Set("requestId", requestId)
		//bodyParams, queryParams := utils.OperateRequestLog(c)
		//cost := time.Since(start)
		//status := c.Writer.Status()
		//method := c.Request.Method
		//ip := c.ClientIP()
		//errString := c.Errors.ByType(gin.ErrorTypePrivate).String()
		//userAgent := c.Request.UserAgent()
		//global.FAST_LOG.Info(path,
		//	zap.Int("status", status),
		//	zap.String("method", method),
		//	zap.String("path", path),
		//	zap.String("query", string(queryParams)),
		//	zap.String("body", string(bodyParams)),
		//	zap.String("ip", ip),
		//	zap.String("requestId", requestId),
		//	zap.String("user-agent", userAgent),
		//	zap.String("errors", errString),
		//	zap.Duration("cost", cost))

		type response struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}
		resp := GinBodyLogMiddleware(c)
		var respStruct response
		_ = json.Unmarshal([]byte(resp), &respStruct)
		//global.FAST_LOG.Info(path,
		//	zap.Int("code", respStruct.Code),
		//	zap.String("msg", respStruct.Msg),
		//)
	}
}

// GinBodyLogMiddleware  使用 responseBodyWriter 替换 gin 中的 responseWriter, 替换的目的是把 response 返回值缓存起来
func GinBodyLogMiddleware(c *gin.Context) string {
	w := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	return w.body.String()
}

// Write
// @param: b []byte  b 即为 response
// @description: Write 方法把 response 缓存到 responseBodyWriter 结构体的 body 属性中
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
