import service from '@/utils/request'
// @Tags MMemberInfo
// @Summary 创建mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberInfo true "创建mMemberInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mMemberInfo/createMMemberInfo [post]
export const createMMemberInfo = (data) => {
  return service({
    url: '/mMemberInfo/createMMemberInfo',
    method: 'post',
    data
  })
}

// @Tags MMemberInfo
// @Summary 删除mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberInfo true "删除mMemberInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberInfo/deleteMMemberInfo [delete]
export const deleteMMemberInfo = (params) => {
  return service({
    url: '/mMemberInfo/deleteMMemberInfo',
    method: 'delete',
    params
  })
}

// @Tags MMemberInfo
// @Summary 批量删除mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mMemberInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberInfo/deleteMMemberInfo [delete]
export const deleteMMemberInfoByIds = (params) => {
  return service({
    url: '/mMemberInfo/deleteMMemberInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags MMemberInfo
// @Summary 更新mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberInfo true "更新mMemberInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mMemberInfo/updateMMemberInfo [put]
export const updateMMemberInfo = (data) => {
  return service({
    url: '/mMemberInfo/updateMMemberInfo',
    method: 'put',
    data
  })
}

// @Tags MMemberInfo
// @Summary 用id查询mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MMemberInfo true "用id查询mMemberInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mMemberInfo/findMMemberInfo [get]
export const findMMemberInfo = (params) => {
  return service({
    url: '/mMemberInfo/findMMemberInfo',
    method: 'get',
    params
  })
}

// @Tags MMemberInfo
// @Summary 分页获取mMemberInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mMemberInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mMemberInfo/getMMemberInfoList [get]
export const getMMemberInfoList = (params) => {
  return service({
    url: '/mMemberInfo/getMMemberInfoList',
    method: 'get',
    params
  })
}

// @Tags MMemberInfo
// @Summary 不需要鉴权的mMemberInfo表接口
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberInfoSearch true "分页获取mMemberInfo表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberInfo/getMMemberInfoPublic [get]
export const getMMemberInfoPublic = () => {
  return service({
    url: '/mMemberInfo/getMMemberInfoPublic',
    method: 'get',
  })
}
// AddCardTimes 充值卡券次数
// @Tags MMemberInfo
// @Summary 充值卡券次数
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberInfo/addCardTimes [POST]
export const addCardTimes = () => {
  return service({
    url: '/mMemberInfo/addCardTimes',
    method: 'POST'
  })
}
// CutCardTimes 扣减卡券次数
// @Tags MMemberInfo
// @Summary 扣减卡券次数
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberInfo/cutCardTimes [POST]
export const cutCardTimes = () => {
  return service({
    url: '/mMemberInfo/cutCardTimes',
    method: 'POST'
  })
}
