package v1

import (
	"my-server/api/v1/Customize"
	"my-server/api/v1/example"
	"my-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	CustomizeApiGroup Customize.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
