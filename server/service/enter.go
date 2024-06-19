package service

import (
	"my-server/service/Customize"
	"my-server/service/example"
	"my-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	CustomizeServiceGroup Customize.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
