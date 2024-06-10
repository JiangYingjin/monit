import service from '@/utils/request'

// @Tags MachineWarningLog
// @Summary 创建machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarningLog true "创建machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineWarningLog/createMachineWarningLog [post]
export const createMachineWarningLog = (data) => {
  return service({
    url: '/machineWarningLog/createMachineWarningLog',
    method: 'post',
    data
  })
}

// @Tags MachineWarningLog
// @Summary 删除machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarningLog true "删除machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarningLog/deleteMachineWarningLog [delete]
export const deleteMachineWarningLog = (params) => {
  return service({
    url: '/machineWarningLog/deleteMachineWarningLog',
    method: 'delete',
    params
  })
}

// @Tags MachineWarningLog
// @Summary 批量删除machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarningLog/deleteMachineWarningLog [delete]
export const deleteMachineWarningLogByIds = (params) => {
  return service({
    url: '/machineWarningLog/deleteMachineWarningLogByIds',
    method: 'delete',
    params
  })
}

// @Tags MachineWarningLog
// @Summary 更新machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarningLog true "更新machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineWarningLog/updateMachineWarningLog [put]
export const updateMachineWarningLog = (data) => {
  return service({
    url: '/machineWarningLog/updateMachineWarningLog',
    method: 'put',
    data
  })
}

// @Tags MachineWarningLog
// @Summary 用id查询machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MachineWarningLog true "用id查询machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineWarningLog/findMachineWarningLog [get]
export const findMachineWarningLog = (params) => {
  return service({
    url: '/machineWarningLog/findMachineWarningLog',
    method: 'get',
    params
  })
}

// @Tags MachineWarningLog
// @Summary 分页获取machineWarningLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取machineWarningLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineWarningLog/getMachineWarningLogList [get]
export const getMachineWarningLogList = (params) => {
  return service({
    url: '/machineWarningLog/getMachineWarningLogList',
    method: 'get',
    params
  })
}
