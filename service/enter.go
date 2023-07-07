package service

import (
	"Library_Project/service/example"
	"Library_Project/service/system"
)

type Service struct {
	SystemServiceGroup  system.SysGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceApp = new(Service)
