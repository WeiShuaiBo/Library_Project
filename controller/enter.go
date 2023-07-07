package controller

import "Library_Project/controller/system"

type ApiGroup struct {
	SystemApiGroup system.SystemControllerGroup
}

var ApiGroupApp = new(ApiGroup)
