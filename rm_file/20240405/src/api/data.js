import service from '@/utils/request'

// @Tags Data
// @Summary 创建Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Data true "创建Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /data/createData [post]
export const createData = (data) => {
  return service({
    url: '/data/createData',
    method: 'post',
    data
  })
}

// @Tags Data
// @Summary 删除Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Data true "删除Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /data/deleteData [delete]
export const deleteData = (params) => {
  return service({
    url: '/data/deleteData',
    method: 'delete',
    params
  })
}

// @Tags Data
// @Summary 批量删除Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /data/deleteData [delete]
export const deleteDataByIds = (params) => {
  return service({
    url: '/data/deleteDataByIds',
    method: 'delete',
    params
  })
}

// @Tags Data
// @Summary 更新Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Data true "更新Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /data/updateData [put]
export const updateData = (data) => {
  return service({
    url: '/data/updateData',
    method: 'put',
    data
  })
}

// @Tags Data
// @Summary 用id查询Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Data true "用id查询Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /data/findData [get]
export const findData = (params) => {
  return service({
    url: '/data/findData',
    method: 'get',
    params
  })
}

// @Tags Data
// @Summary 分页获取Data列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Data列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /data/getDataList [get]
export const getDataList = (params) => {
  return service({
    url: '/data/getDataList',
    method: 'get',
    params
  })
}
