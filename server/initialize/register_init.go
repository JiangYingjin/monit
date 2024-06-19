package initialize

import (
	_ "my-server/source/example"
	_ "my-server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
