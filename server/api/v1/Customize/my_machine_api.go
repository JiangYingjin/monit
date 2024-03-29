package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type MyMachineApi struct {
}

func (m *MyMachineApi) Login(c *gin.Context) {
	response.Ok(c)
}
