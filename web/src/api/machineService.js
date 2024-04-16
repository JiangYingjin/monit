import service from '@/utils/request'

// @Tags MachineService
// @Summary 创建数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineService true "创建数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineService/createMachineService [post]
export const createMachineService = (data) => {
  return service({
    url: '/machineService/createMachineService',
    method: 'post',
    data
  })
}

// @Tags MachineService
// @Summary 删除数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineService true "删除数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineService/deleteMachineService [delete]
export const deleteMachineService = (params) => {
  return service({
    url: '/machineService/deleteMachineService',
    method: 'delete',
    params
  })
}

// @Tags MachineService
// @Summary 批量删除数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineService/deleteMachineService [delete]
export const deleteMachineServiceByIds = (params) => {
  return service({
    url: '/machineService/deleteMachineServiceByIds',
    method: 'delete',
    params
  })
}

// @Tags MachineService
// @Summary 更新数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MachineService true "更新数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineService/updateMachineService [put]
export const updateMachineService = (data) => {
  return service({
    url: '/machineService/updateMachineService',
    method: 'put',
    data
  })
}

// @Tags MachineService
// @Summary 用id查询数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MachineService true "用id查询数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineService/findMachineService [get]
export const findMachineService = (params) => {
  return service({
    url: '/machineService/findMachineService',
    method: 'get',
    params
  })
}

// @Tags MachineService
// @Summary 分页获取数据类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineService/getMachineServiceList [get]
export const getMachineServiceList = (params) => {
  return service({
    url: '/machineService/getMachineServiceList',
    method: 'get',
    params
  })
}
