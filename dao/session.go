package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("session-secret-key"))
var sessionName = "session-name"

func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = ""
	session.Values["id"] = 0
	return session.Save(c.Request, c.Writer)
}
