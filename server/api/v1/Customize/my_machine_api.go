package Customize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	Customize2 "github.com/flipped-aurora/gin-vue-admin/server/service/Customize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type MyMachineApi struct {
}

func init() {
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	for {
	//		time.Sleep(1 * time.Second)
	//
	//		dataTypeID := 1
	//		value := (float64(utils.RandomInt(1, 1000))) / 200.0
	//		machineID := 1
	//		data := Customize.Data{
	//			DataTypeID: &dataTypeID,
	//			Value:      &value,
	//			MachineID:  &machineID,
	//			CreatedBy:  0,
	//			UpdatedBy:  0,
	//			DeletedBy:  0,
	//		}
	//
	//		var dataService Customize2.DataService
	//		if err := dataService.CreateData(&data); err != nil {
	//		} else {
	//			var myMachineApi MyMachineApi
	//			myMachineApi.UploadDataHook(data)
	//		}
	//	}
	//}()
}

// MachineLogin
// @Tags Machine
// @Summary 机器登录
// @Produce  application/json
// @Param MachineID body string true "MachineID"
// @Param Password body string true "Password"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /machine/machineLogin [post]
//
//	{
//		"MachineID": "1",
//		"Password": "123456"
//	}
func (m *MyMachineApi) MachineLogin(c *gin.Context) {
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
	m.TokenNext(c, machine)
	return
}

// GetData 获取数据
// @Tags Machine
// @Summary 创建Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body CustomizeReq.GetDataReq true "创建Machine"
// @Router /machine/createMachine [post]
//
//	{
//		"data_type_id": "1",
//		"machine_ids": [
//		"1"
//		],
//		"start_time": "1980-03-18 07:13:05",
//		"end_time": "1989-07-28 18:34:43"
//	}
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
		global.GVA_DB.Model(&Customize.Data{}).
			Where("machine_i_d = ? and data_type_i_d = ? and created_at between ? and ?", machineID, req.DataTypeID, req.StartTime, req.EndTime). // "2024-04-13 14:25:00"
			Find(&tmp)
		result[machineID] = tmp
	}

	response.OkWithData(result, c)
}

type PropStatus struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type SetMachineServiceReq struct {
	MachineID uint                    `json:"machineID" binding:"required"`
	Services  map[string][]PropStatus `json:"services" binding:"required"`
}

func (s *SetMachineServiceReq) ToParamString() {
	var result []string
	for name, propStatus := range s.Services {
		for _, prop := range propStatus {
			result = append(result, fmt.Sprintf("--%s-%s=%s", name, prop.Name, prop.Value))
		}
	}
}

// SetMachineService 开启/关闭某个进程的监听服务
// @Tags Machine
// @Summary 设置机器监控状态
// @Produce  application/json
// @Param machineID body uint true "MachineID"
// @Param services body SetMachineServiceReq true "services"
// @Router /machine/setMachineService [post]
func (m *MyMachineApi) SetMachineService(c *gin.Context) {
	//var newServiceReq SetMachineServiceReq
	//err := c.ShouldBindJSON(&newServiceReq)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//
	//err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	//	var machine Customize.Machine
	//	if err := tx.Model(&Customize.Machine{}).Where("id =", newServiceReq.MachineID).Find(&machine).Error; err != nil {
	//		return err
	//	}
	//
	//	var oldServices SetMachineServiceReq
	//	err = json.Unmarshal([]byte(machine.Service), &oldServices)
	//	if err != nil {
	//		return err
	//	}
	//
	//	for name, propStatus := range oldServices.Services {
	//		var serviceInDB Customize.ServiceTemplate
	//		if err := tx.Model(&Customize.ServiceTemplate{}).Where("service =", name).Find(&serviceInDB).Error; err != nil {
	//			return err
	//		}
	//
	//		newService, ok := newServiceReq.Services[name]
	//		if !ok {
	//			continue
	//		}
	//
	//		for i, prop := range propStatus {
	//			if prop.Name != newService[i].Name {
	//				return errors.New("name not match")
	//			}
	//
	//			switch prop.Type {
	//			case "bool":
	//				_, err = strconv.ParseBool(prop.Value)
	//			case "int":
	//				_, err = strconv.ParseInt(prop.Value, 10, 64)
	//			case "str":
	//				break
	//			default:
	//				return errors.New("unknown type")
	//			}
	//			if err != nil {
	//				return err
	//			}
	//
	//			oldServices.Services[name][i].Value = newService[i].Value
	//		}
	//	}
	//
	//	machine.Service = cast.ToString(oldServices)
	//	err = tx.Save(&machine).Error
	//	return nil
	//})
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
}

func (m *MyMachineApi) UpdateMachineService(c *gin.Context) {
	var machine struct {
		MachineID uint `json:"machineID" binding:"required"`
		Services  []struct {
			Service string `json:"service" binding:"required"`
			Enable  bool   `json:"enable" binding:"required"`
		} `json:"services" binding:"required"`
	}
	err := c.ShouldBindJSON(&machine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = global.GVA_DB.Model(&Customize.Machine{}).Where("id = ?", machine.MachineID).Update("service", machine.Services).Error
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type UploadDataMultiReq struct {
	Data []CreateDataReq `json:"data,omitempty"`
}

func (dataApi *DataApi) CreateDataMulti(c *gin.Context) {
	var dataArray UploadDataMultiReq
	err := c.ShouldBindJSON(&dataArray)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	for _, machineDataReq := range dataArray.Data {
		machineData, err := machineDataReq.ToData()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}

		machineData.CreatedBy = uint(*machineData.MachineID)
		if err := dataService.CreateData(&machineData); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
			myMachineApi := MyMachineApi{}
			myMachineApi.UploadDataHook(machineData)
		}
	}
}

// 检测数据是否异常，如果异常则发送告警
func (m *MyMachineApi) UploadDataHook(data Customize.Data) {

	var warning Customize.MachineWarning
	global.GVA_DB.Model(&Customize.MachineWarning{}).
		Where("machine_i_d = ? and data_type_i_d = ?", data.MachineID, data.DataTypeID).
		Find(&warning)
	if warning != (Customize.MachineWarning{}) && *warning.Limit > *data.Value {
		userService := system.UserService{}
		userID, err := strconv.ParseInt(warning.ReporterID, 10, 64)
		user, err := userService.FindUserById(int(userID))
		if err != nil {
			global.GVA_LOG.Error("获取用户信息失败", zap.Error(err))
			return
		}
		warningMessage := "机器ID: " + strconv.FormatUint(uint64(*data.MachineID), 10) + " 数据类型ID: " + strconv.FormatUint(uint64(*data.DataTypeID), 10) + " 数据异常"
		myMachineService := Customize2.MyMachineService{}
		err = myMachineService.SendEmail(user.Email, warningMessage)
		if err != nil {
			global.GVA_LOG.Error("发送邮件失败", zap.Error(err))
			return
		}
	}
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
func (m *MyMachineApi) TokenNext(c *gin.Context, machine Customize.Machine) {
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
