import service from '@/utils/request'
// @Tags MMemberCard
// @Summary 创建mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberCard true "创建mMemberCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mMemberCard/createMMemberCard [post]
export const createMMemberCard = (data) => {
  return service({
    url: '/mMemberCard/createMMemberCard',
    method: 'post',
    data
  })
}

// @Tags MMemberCard
// @Summary 删除mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberCard true "删除mMemberCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberCard/deleteMMemberCard [delete]
export const deleteMMemberCard = (params) => {
  return service({
    url: '/mMemberCard/deleteMMemberCard',
    method: 'delete',
    params
  })
}

// @Tags MMemberCard
// @Summary 批量删除mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mMemberCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberCard/deleteMMemberCard [delete]
export const deleteMMemberCardByIds = (params) => {
  return service({
    url: '/mMemberCard/deleteMMemberCardByIds',
    method: 'delete',
    params
  })
}

// @Tags MMemberCard
// @Summary 更新mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberCard true "更新mMemberCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mMemberCard/updateMMemberCard [put]
export const updateMMemberCard = (data) => {
  return service({
    url: '/mMemberCard/updateMMemberCard',
    method: 'put',
    data
  })
}

// @Tags MMemberCard
// @Summary 用id查询mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MMemberCard true "用id查询mMemberCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mMemberCard/findMMemberCard [get]
export const findMMemberCard = (params) => {
  return service({
    url: '/mMemberCard/findMMemberCard',
    method: 'get',
    params
  })
}

// @Tags MMemberCard
// @Summary 分页获取mMemberCard表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mMemberCard表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mMemberCard/getMMemberCardList [get]
export const getMMemberCardList = (params) => {
  return service({
    url: '/mMemberCard/getMMemberCardList',
    method: 'get',
    params
  })
}

// @Tags MMemberCard
// @Summary 不需要鉴权的mMemberCard表接口
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberCardSearch true "分页获取mMemberCard表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberCard/getMMemberCardPublic [get]
export const getMMemberCardPublic = () => {
  return service({
    url: '/mMemberCard/getMMemberCardPublic',
    method: 'get',
  })
}
