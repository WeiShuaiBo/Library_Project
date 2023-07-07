package router

import "Library_Project/router/system"

type RouterGroup struct {
	System system.RouterGroup
	//Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
