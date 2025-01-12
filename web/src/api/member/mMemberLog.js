import service from '@/utils/request'
// @Tags MMemberLog
// @Summary 创建mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberLog true "创建mMemberLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mMemberLog/createMMemberLog [post]
export const createMMemberLog = (data) => {
  data['renewPeriod'] = parseInt(data['renewPeriod'], 10)
  return service({
    url: '/mMemberLog/createMMemberLog',
    method: 'post',
    data
  })
}

// @Tags MMemberLog
// @Summary 删除mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberLog true "删除mMemberLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberLog/deleteMMemberLog [delete]
export const deleteMMemberLog = (params) => {
  return service({
    url: '/mMemberLog/deleteMMemberLog',
    method: 'delete',
    params
  })
}

// @Tags MMemberLog
// @Summary 批量删除mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mMemberLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberLog/deleteMMemberLog [delete]
export const deleteMMemberLogByIds = (params) => {
  return service({
    url: '/mMemberLog/deleteMMemberLogByIds',
    method: 'delete',
    params
  })
}

export const updateMMemberLogList = (data) => {
  return service({
    url: '/mMemberLog/updateMMemberLogList',
    method: 'post',
    data
  })
}

// @Tags MMemberLog
// @Summary 更新mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberLog true "更新mMemberLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mMemberLog/updateMMemberLog [put]
export const updateMMemberLog = (data) => {
  return service({
    url: '/mMemberLog/updateMMemberLog',
    method: 'put',
    data
  })
}

// @Tags MMemberLog
// @Summary 用id查询mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MMemberLog true "用id查询mMemberLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mMemberLog/findMMemberLog [get]
export const findMMemberLog = (params) => {
  return service({
    url: '/mMemberLog/findMMemberLog',
    method: 'get',
    params
  })
}

// @Tags MMemberLog
// @Summary 分页获取mMemberLog表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mMemberLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mMemberLog/getMMemberLogList [get]
export const getMMemberLogList = (params) => {
  return service({
    url: '/mMemberLog/getMMemberLogList',
    method: 'get',
    params
  })
}

// @Tags MMemberLog
// @Summary 不需要鉴权的mMemberLog表接口
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberLogSearch true "分页获取mMemberLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberLog/getMMemberLogPublic [get]
export const getMMemberLogPublic = () => {
  return service({
    url: '/mMemberLog/getMMemberLogPublic',
    method: 'get',
  })
}
// UpdateCheckStatus 对账状态修改接口
// @Tags MMemberLog
// @Summary 对账状态修改接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberLog/updateCheckStatus [POST]
export const updateCheckStatus = () => {
  return service({
    url: '/mMemberLog/updateCheckStatus',
    method: 'POST'
  })
}
