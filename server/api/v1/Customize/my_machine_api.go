package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type MyMachineApi struct {
}

// MachineLogin
// @Tags     Base
// @Summary  机器登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *MyMachineApi) MachineLogin(c *gin.Context) {
	var l systemReq.MachineLoginReq
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.Rules{"MachineID": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	machine, err := machineService.GetMachine(l.MachineID)
	if err != nil || machine.Password != l.Password || strconv.FormatUint(uint64(machine.ID), 10) != l.MachineID {
		global.GVA_LOG.Error("登陆失败! 机器ID不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("机器ID不存在或者密码错误", c)
		return
	}
	b.TokenNext(c, machine)
	return
}

type MachineClaim struct {
	MachineID  string
	BufferTime int64
	jwt.RegisteredClaims
}

func (machine MachineClaim) Valid() error {
	return nil
}

// TokenNext 登录以后签发jwt
func (b *MyMachineApi) TokenNext(c *gin.Context, machine Customize.Machine) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名

	bf, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := MachineClaim{
		MachineID:  strconv.FormatUint(uint64(machine.ID), 10),
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,              // 签名的发行者
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	machineToken, err := token.SignedString(j.SigningKey)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}

	utils.SetToken(c, machineToken, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))

	rsp := make(map[string]interface{})
	rsp["Machine"] = machine
	rsp["Token"] = machineToken                                        // x-token
	rsp["ExpiresAt"] = claims.RegisteredClaims.ExpiresAt.Unix() * 1000 // in ms
	response.OkWithDetailed(rsp, "登录成功", c)
}

// GetData 获取数据
// @Tags Machine
// @Summary 创建Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body CustomizeReq.GetDataReq true "创建Machine"
// @Router /machine/createMachine [post]
func (m *MyMachineApi) GetData(c *gin.Context) {
	var req CustomizeReq.GetDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	for _, machineID := range req.MachineIDs {
		_, err := machineService.GetMachine(machineID)
		if err != nil {
			response.FailWithMessage("machine not found", c)
			return
		}
	}

	result := make(map[string][]Customize.Data)
	for _, machineID := range req.MachineIDs {
		tmp := make([]Customize.Data, 0)
		//global.GVA_DB.Where("machine_id in ?", req.MachineIDs).Where("data_type_id in ?", req.DataTypeID).Find(&Customize.Data{})
		global.GVA_DB.Where("machine_id = ?", machineID).Where("data_type_id = ?", req.DataTypeID).Where("created_at between ? and ?", req.StartTime, req.EndTime).Find(&tmp)
		result[machineID] = tmp
	}

	response.OkWithData(CustomizeReq.GetDataRsp{Data: result}, c)
}
