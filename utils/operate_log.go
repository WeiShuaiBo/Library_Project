package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// @author: [ziop](https://github.com/wangpi26)
// @function: OperateRequestLog
// @description: 处理请求数据,获取请求体里的数据和解析请求行里面的数据，注意非GET请求的数据只能放在请求体里面才能解析
// @param: c *gin.Context
// @return: mBody []byte post方法参数, queryParams []byte get方法参数
func OperateRequestLog(c *gin.Context) (mBody []byte, queryParams []byte) {
	if c.Request.Method != http.MethodGet {
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		json.Unmarshal(body, &mBody)
	} else {
		query := c.Request.URL.RawQuery
		//        解码URL编码
		query, _ = url.QueryUnescape(query)
		//        分割URL参数
		querys := strings.Split(query, "&")
		mGet := make(map[string]string)
		//        将URL参数转成map
		for _, v := range querys {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				mGet[kv[0]] = kv[1]
			}
		}
		queryParams, _ = json.Marshal(mGet)
	}

	return mBody, queryParams
}
