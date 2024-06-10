import service from '@/utils/request'

// @Tags MachineWarning
// @Summary 创建机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarning true "创建机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineWarning/createMachineWarning [post]
export const createMachineWarning = (data) => {
  return service({
    url: '/machineWarning/createMachineWarning',
    method: 'post',
    data
  })
}

// @Tags MachineWarning
// @Summary 删除机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarning true "删除机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarning/deleteMachineWarning [delete]
export const deleteMachineWarning = (params) => {
  return service({
    url: '/machineWarning/deleteMachineWarning',
    method: 'delete',
    params
  })
}

// @Tags MachineWarning
// @Summary 批量删除机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarning/deleteMachineWarning [delete]
export const deleteMachineWarningByIds = (params) => {
  return service({
    url: '/machineWarning/deleteMachineWarningByIds',
    method: 'delete',
    params
  })
}

// @Tags MachineWarning
// @Summary 更新机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineWarning true "更新机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineWarning/updateMachineWarning [put]
export const updateMachineWarning = (data) => {
  return service({
    url: '/machineWarning/updateMachineWarning',
    method: 'put',
    data
  })
}

// @Tags MachineWarning
// @Summary 用id查询机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MachineWarning true "用id查询机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineWarning/findMachineWarning [get]
export const findMachineWarning = (params) => {
  return service({
    url: '/machineWarning/findMachineWarning',
    method: 'get',
    params
  })
}

// @Tags MachineWarning
// @Summary 分页获取机器告警列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取机器告警列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineWarning/getMachineWarningList [get]
export const getMachineWarningList = (params) => {
  return service({
    url: '/machineWarning/getMachineWarningList',
    method: 'get',
    params
  })
}
