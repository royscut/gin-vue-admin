package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	memberRes "github.com/flipped-aurora/gin-vue-admin/server/model/member/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/poster"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MMemberInfoApi struct{}

// CreateMMemberInfo 创建mMemberInfo表
// @Tags MMemberInfo
// @Summary 创建mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberInfo true "创建mMemberInfo表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mMemberInfo/createMMemberInfo [post]
func (mMemberInfoApi *MMemberInfoApi) CreateMMemberInfo(c *gin.Context) {
	var mMemberInfo member.MMemberInfo
	err := c.ShouldBindJSON(&mMemberInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	member_file, err := poster.CreateMemberPoster("https://wwww.baidu.com", "测试会员ID:10000")
	if err != nil {
		global.GVA_LOG.Error("绘制图片失败!", zap.Error(err))
		response.FailWithMessage("绘制图片失败:"+err.Error(), c)
		return
	}
	_ = member_file

	err = mMemberInfoService.CreateMMemberInfo(&mMemberInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteMMemberInfo 删除mMemberInfo表
// @Tags MMemberInfo
// @Summary 删除mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberInfo true "删除mMemberInfo表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mMemberInfo/deleteMMemberInfo [delete]
func (mMemberInfoApi *MMemberInfoApi) DeleteMMemberInfo(c *gin.Context) {
	ID := c.Query("ID")
	err := mMemberInfoService.DeleteMMemberInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMMemberInfoByIds 批量删除mMemberInfo表
// @Tags MMemberInfo
// @Summary 批量删除mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mMemberInfo/deleteMMemberInfoByIds [delete]
func (mMemberInfoApi *MMemberInfoApi) DeleteMMemberInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mMemberInfoService.DeleteMMemberInfoByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMMemberInfo 更新mMemberInfo表
// @Tags MMemberInfo
// @Summary 更新mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberInfo true "更新mMemberInfo表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mMemberInfo/updateMMemberInfo [put]
func (mMemberInfoApi *MMemberInfoApi) UpdateMMemberInfo(c *gin.Context) {
	var mMemberInfo member.MMemberInfo
	err := c.ShouldBindJSON(&mMemberInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mMemberInfoService.UpdateMMemberInfo(mMemberInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMMemberInfo 用id查询mMemberInfo表
// @Tags MMemberInfo
// @Summary 用id查询mMemberInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mMemberInfo表"
// @Success 200 {object} response.Response{data=member.MMemberInfo,msg=string} "查询成功"
// @Router /mMemberInfo/findMMemberInfo [get]
func (mMemberInfoApi *MMemberInfoApi) FindMMemberInfo(c *gin.Context) {
	ID := c.Query("ID")
	remMemberInfo, err := mMemberInfoService.GetMMemberInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remMemberInfo, c)
}

// GetMMemberInfoList 分页获取mMemberInfo表列表
// @Tags MMemberInfo
// @Summary 分页获取mMemberInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberInfoSearch true "分页获取mMemberInfo表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mMemberInfo/getMMemberInfoList [get]
func (mMemberInfoApi *MMemberInfoApi) GetMMemberInfoList(c *gin.Context) {
	var pageInfo memberReq.MMemberInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := mMemberInfoService.GetMMemberInfoInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	var ids []*uint
	for _, item := range list {
		ids = append(ids, &item.ID)
	}

	cardList, _, err := mMemberCardService.GetMMemberCardInfoListByIDs(ids)
	if err != nil {
		global.GVA_LOG.Error("获取会员卡信息失败!", zap.Error(err))
		response.FailWithMessage("获取会员卡信息失败:"+err.Error(), c)
		return
	}

	cardMap := make(map[uint]member.MMemberCard)
	for _, card := range cardList {
		cardMap[*card.MemberId] = card
	}

	var reslist []memberRes.MemberExtInfoResponse
	for _, member := range list {
		mebRes := memberRes.MemberExtInfoResponse{
			MMemberInfo: member,
		}
		card, exists := cardMap[member.ID]
		if exists {
			mebRes.MemberId = card.MemberId
			mebRes.TotalTimes = card.TotalTimes
			mebRes.RemainTimes = card.RemainTimes
			mebRes.UsedTimes = card.UsedTimes
			mebRes.CardPic = card.CardPic
			mebRes.ExpireTime = card.ExpireTime
		}
		reslist = append(reslist, mebRes)
	}

	response.OkWithDetailed(response.PageResult{
		List:     reslist,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetMMemberInfoPublic 不需要鉴权的mMemberInfo表接口
// @Tags MMemberInfo
// @Summary 不需要鉴权的mMemberInfo表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberInfo/getMMemberInfoPublic [get]
func (mMemberInfoApi *MMemberInfoApi) GetMMemberInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mMemberInfoService.GetMMemberInfoPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mMemberInfo表接口信息",
	}, "获取成功", c)
}

// AddCardTimes 充值卡券次数
// @Tags MMemberInfo
// @Summary 充值卡券次数
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberInfoSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberInfo/addCardTimes [POST]
func (mMemberInfoApi *MMemberInfoApi) AddCardTimes(c *gin.Context) {
	// 请添加自己的业务逻辑
	err := mMemberInfoService.AddCardTimes()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// CutCardTimes 扣减卡券次数
// @Tags MMemberInfo
// @Summary 扣减卡券次数
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberInfoSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberInfo/cutCardTimes [POST]
func (mMemberInfoApi *MMemberInfoApi) CutCardTimes(c *gin.Context) {
	// 请添加自己的业务逻辑
	err := mMemberInfoService.CutCardTimes()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}
