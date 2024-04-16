import service from '@/utils/request'

// @Tags ServiceTemplate
// @Summary 创建命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceTemplate true "创建命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serviceTemplate/createServiceTemplate [post]
export const createServiceTemplate = (data) => {
  return service({
    url: '/serviceTemplate/createServiceTemplate',
    method: 'post',
    data
  })
}

// @Tags ServiceTemplate
// @Summary 删除命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceTemplate true "删除命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceTemplate/deleteServiceTemplate [delete]
export const deleteServiceTemplate = (params) => {
  return service({
    url: '/serviceTemplate/deleteServiceTemplate',
    method: 'delete',
    params
  })
}

// @Tags ServiceTemplate
// @Summary 批量删除命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceTemplate/deleteServiceTemplate [delete]
export const deleteServiceTemplateByIds = (params) => {
  return service({
    url: '/serviceTemplate/deleteServiceTemplateByIds',
    method: 'delete',
    params
  })
}

// @Tags ServiceTemplate
// @Summary 更新命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceTemplate true "更新命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serviceTemplate/updateServiceTemplate [put]
export const updateServiceTemplate = (data) => {
  return service({
    url: '/serviceTemplate/updateServiceTemplate',
    method: 'put',
    data
  })
}

// @Tags ServiceTemplate
// @Summary 用id查询命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ServiceTemplate true "用id查询命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serviceTemplate/findServiceTemplate [get]
export const findServiceTemplate = (params) => {
  return service({
    url: '/serviceTemplate/findServiceTemplate',
    method: 'get',
    params
  })
}

// @Tags ServiceTemplate
// @Summary 分页获取命令模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取命令模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceTemplate/getServiceTemplateList [get]
export const getServiceTemplateList = (params) => {
  return service({
    url: '/serviceTemplate/getServiceTemplateList',
    method: 'get',
    params
  })
}
