import service from '@/utils/request'

// @Tags DataType
// @Summary 创建DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataType true "创建DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataType/createDataType [post]
export const createDataType = (data) => {
  return service({
    url: '/dataType/createDataType',
    method: 'post',
    data
  })
}

// @Tags DataType
// @Summary 删除DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataType true "删除DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataType/deleteDataType [delete]
export const deleteDataType = (params) => {
  return service({
    url: '/dataType/deleteDataType',
    method: 'delete',
    params
  })
}

// @Tags DataType
// @Summary 批量删除DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataType/deleteDataType [delete]
export const deleteDataTypeByIds = (params) => {
  return service({
    url: '/dataType/deleteDataTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags DataType
// @Summary 更新DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataType true "更新DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataType/updateDataType [put]
export const updateDataType = (data) => {
  return service({
    url: '/dataType/updateDataType',
    method: 'put',
    data
  })
}

// @Tags DataType
// @Summary 用id查询DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataType true "用id查询DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataType/findDataType [get]
export const findDataType = (params) => {
  return service({
    url: '/dataType/findDataType',
    method: 'get',
    params
  })
}

// @Tags DataType
// @Summary 分页获取DataType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DataType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataType/getDataTypeList [get]
export const getDataTypeList = (params) => {
  return service({
    url: '/dataType/getDataTypeList',
    method: 'get',
    params
  })
}
