import service from '@/utils/request'
// @Tags MMemberGiftCard
// @Summary 创建mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberGiftCard true "创建mMemberGiftCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mMemberGiftCard/createMMemberGiftCard [post]
export const createMMemberGiftCard = (data) => {
  return service({
    url: '/mMemberGiftCard/createMMemberGiftCard',
    method: 'post',
    data
  })
}

// @Tags MMemberGiftCard
// @Summary 删除mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberGiftCard true "删除mMemberGiftCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberGiftCard/deleteMMemberGiftCard [delete]
export const deleteMMemberGiftCard = (params) => {
  return service({
    url: '/mMemberGiftCard/deleteMMemberGiftCard',
    method: 'delete',
    params
  })
}

// @Tags MMemberGiftCard
// @Summary 批量删除mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mMemberGiftCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mMemberGiftCard/deleteMMemberGiftCard [delete]
export const deleteMMemberGiftCardByIds = (params) => {
  return service({
    url: '/mMemberGiftCard/deleteMMemberGiftCardByIds',
    method: 'delete',
    params
  })
}

// @Tags MMemberGiftCard
// @Summary 更新mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MMemberGiftCard true "更新mMemberGiftCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mMemberGiftCard/updateMMemberGiftCard [put]
export const updateMMemberGiftCard = (data) => {
  return service({
    url: '/mMemberGiftCard/updateMMemberGiftCard',
    method: 'put',
    data
  })
}

// @Tags MMemberGiftCard
// @Summary 用id查询mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MMemberGiftCard true "用id查询mMemberGiftCard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mMemberGiftCard/findMMemberGiftCard [get]
export const findMMemberGiftCard = (params) => {
  return service({
    url: '/mMemberGiftCard/findMMemberGiftCard',
    method: 'get',
    params
  })
}

// @Tags MMemberGiftCard
// @Summary 分页获取mMemberGiftCard表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mMemberGiftCard表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mMemberGiftCard/getMMemberGiftCardList [get]
export const getMMemberGiftCardList = (params) => {
  return service({
    url: '/mMemberGiftCard/getMMemberGiftCardList',
    method: 'get',
    params
  })
}

// @Tags MMemberGiftCard
// @Summary 不需要鉴权的mMemberGiftCard表接口
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberGiftCardSearch true "分页获取mMemberGiftCard表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberGiftCard/getMMemberGiftCardPublic [get]
export const getMMemberGiftCardPublic = () => {
  return service({
    url: '/mMemberGiftCard/getMMemberGiftCardPublic',
    method: 'get',
  })
}
