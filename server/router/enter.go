package router

import (
	"my-server/router/Customize"
	"my-server/router/example"
	"my-server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Customize Customize.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
