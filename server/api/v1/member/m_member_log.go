package member

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MMemberLogApi struct{}

// CreateMMemberLog 创建mMemberLog表
// @Tags MMemberLog
// @Summary 创建mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberLog true "创建mMemberLog表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mMemberLog/createMMemberLog [post]
func (mMemberLogApi *MMemberLogApi) CreateMMemberLog(c *gin.Context) {
	var mMemberLog member.MMemberLog

	userid := int(utils.GetUserID(c))
	username := utils.GetUserName(c)
	mMemberLog.OperatorId = &userid
	mMemberLog.OperatorName = &username

	err := c.ShouldBindJSON(&mMemberLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	global.GVA_LOG.Info("mMemberLog", zap.Any("mMemberLog", mMemberLog))
	//先记录日志
	err = mMemberLogService.CreateMMemberLog(&mMemberLog)
	if err != nil {
		global.GVA_LOG.Error("创建日志失败!", zap.Error(err))
		response.FailWithMessage("创建日志失败:"+err.Error(), c)
		return
	}

	//查询是否有会员卡
	mMemberCard, err := mMemberCardService.GetMMemberCardByMemberId(*mMemberLog.MemberUid)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.GVA_LOG.Error("查询会员卡失败!", zap.Error(err))
		response.FailWithMessage("查询会员卡失败:"+err.Error(), c)
		return
	}

	global.GVA_LOG.Info("mMemberCard", zap.Any("mMemberCard", mMemberCard))
	var opErr error
	if err == gorm.ErrRecordNotFound { //没有会员卡
		//mMemberCard = member.MMemberCard{}
		mMemberCard.MemberId = mMemberLog.MemberUid
		mMemberCard.TotalTimes = mMemberLog.OpTimes
		mMemberCard.RemainTimes = mMemberLog.OpTimes

		usedTimes := 0
		mMemberCard.UsedTimes = &usedTimes
		//mMemberCard.CardPic =  ""

		expireTime := time.Unix(int64(*mMemberLog.RenewPeriod*24*3600)+time.Now().Unix(), 0)
		mMemberCard.ExpireTime = &expireTime

		//mMemberCard.Remark = ""

		//创建会员卡
		opErr = mMemberCardService.CreateMMemberCard(&mMemberCard)

	} else { //有会员卡
		if *mMemberLog.OpType == "add" { //充值，次数增加
			*mMemberCard.TotalTimes += *mMemberLog.OpTimes
			*mMemberCard.RemainTimes += *mMemberLog.OpTimes
		} else { //消费，次数减少
			*mMemberCard.RemainTimes -= *mMemberLog.OpTimes
			if mMemberCard.UsedTimes == nil {
				usedTimes := 0
				mMemberCard.UsedTimes = &usedTimes
			}
			*mMemberCard.UsedTimes += *mMemberLog.OpTimes
		}

		//更新会员卡
		opErr = mMemberCardService.UpdateMMemberCard(mMemberCard)
	}

	if opErr != nil {
		global.GVA_LOG.Error("创建/更新失败!", zap.Error(err), zap.Any("mMemberLog", mMemberLog))
		response.FailWithMessage("创建/更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("操作成功", c)
}

// DeleteMMemberLog 删除mMemberLog表
// @Tags MMemberLog
// @Summary 删除mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberLog true "删除mMemberLog表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mMemberLog/deleteMMemberLog [delete]
func (mMemberLogApi *MMemberLogApi) DeleteMMemberLog(c *gin.Context) {
	ID := c.Query("ID")
	err := mMemberLogService.DeleteMMemberLog(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMMemberLogByIds 批量删除mMemberLog表
// @Tags MMemberLog
// @Summary 批量删除mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mMemberLog/deleteMMemberLogByIds [delete]
func (mMemberLogApi *MMemberLogApi) DeleteMMemberLogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mMemberLogService.DeleteMMemberLogByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMMemberLog 更新mMemberLog表
// @Tags MMemberLog
// @Summary 更新mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberLog true "更新mMemberLog表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mMemberLog/updateMMemberLog [put]
func (mMemberLogApi *MMemberLogApi) UpdateMMemberLog(c *gin.Context) {
	var mMemberLog member.MMemberLog
	err := c.ShouldBindJSON(&mMemberLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mMemberLogService.UpdateMMemberLog(mMemberLog)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMMemberLog 用id查询mMemberLog表
// @Tags MMemberLog
// @Summary 用id查询mMemberLog表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mMemberLog表"
// @Success 200 {object} response.Response{data=member.MMemberLog,msg=string} "查询成功"
// @Router /mMemberLog/findMMemberLog [get]
func (mMemberLogApi *MMemberLogApi) FindMMemberLog(c *gin.Context) {
	ID := c.Query("ID")
	remMemberLog, err := mMemberLogService.GetMMemberLog(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remMemberLog, c)
}

// GetMMemberLogList 分页获取mMemberLog表列表
// @Tags MMemberLog
// @Summary 分页获取mMemberLog表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberLogSearch true "分页获取mMemberLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mMemberLog/getMMemberLogList [get]
func (mMemberLogApi *MMemberLogApi) GetMMemberLogList(c *gin.Context) {
	var pageInfo memberReq.MMemberLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Info("pageInfo Req", zap.Any("pageInfo", pageInfo))
	list, total, err := mMemberLogService.GetMMemberLogInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetMMemberLogPublic 不需要鉴权的mMemberLog表接口
// @Tags MMemberLog
// @Summary 不需要鉴权的mMemberLog表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberLog/getMMemberLogPublic [get]
func (mMemberLogApi *MMemberLogApi) GetMMemberLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mMemberLogService.GetMMemberLogPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mMemberLog表接口信息",
	}, "获取成功", c)
}

// UpdateCheckStatus 对账状态修改接口
// @Tags MMemberLog
// @Summary 对账状态修改接口
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberLogSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /mMemberLog/updateCheckStatus [POST]
func (mMemberLogApi *MMemberLogApi) UpdateCheckStatus(c *gin.Context) {
	// 请添加自己的业务逻辑
	err := mMemberLogService.UpdateCheckStatus()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}
