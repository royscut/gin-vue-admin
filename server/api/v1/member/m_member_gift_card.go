package member

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/poster"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MMemberGiftCardApi struct{}

// CreateMMemberGiftCard 创建mMemberGiftCard表
// @Tags MMemberGiftCard
// @Summary 创建mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberGiftCard true "创建mMemberGiftCard表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mMemberGiftCard/createMMemberGiftCard [post]
func (mMemberGiftCardApi *MMemberGiftCardApi) CreateMMemberGiftCard(c *gin.Context) {
	var mMemberGiftCard member.MMemberGiftCard
	err := c.ShouldBindJSON(&mMemberGiftCard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//这三项从前端传过来
	// mMemberGiftCard.ExpireTime = member.GetExpireTime()
	// mMemberGiftCard.TotalTimes = 1
	// mMemberGiftCard.Remark = "备注"

	//后台自动生成
	randInt := rand.Intn(10000000) // Generate a random integer
	currentUnixTime := time.Now().Unix()
	tmpKey := strconv.Itoa(randInt) + strconv.FormatInt(currentUnixTime, 10)
	md5Key := utils.MD5V([]byte(tmpKey))
	mMemberGiftCard.CardKey = &md5Key

	tmpUsedTimes := int(0)
	mMemberGiftCard.UsedTimes = &tmpUsedTimes
	mMemberGiftCard.RemainTimes = mMemberGiftCard.TotalTimes

	member_gift_file, err := poster.CreateMemberGiftCardPoster("https://wwww.baidu.com", "测试会员ID:10000")
	if err != nil {
		global.GVA_LOG.Error("绘制礼券图片失败!", zap.Error(err))
		response.FailWithMessage("绘制礼券图片失败:"+err.Error(), c)
		return
	}
	_ = member_gift_file

	tmpPic := "https://img.kancloud.cn/d8/d0/d8d0c3fc612766ae81605268059d786c_1692x2000.png"
	mMemberGiftCard.CardPic = &tmpPic

	err = mMemberGiftCardService.CreateMMemberGiftCard(&mMemberGiftCard)
	if err != nil {
		global.GVA_LOG.Error("派发礼券失败!", zap.Error(err))
		response.FailWithMessage("派发礼券失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("派发礼券成功", c)
}

// DeleteMMemberGiftCard 删除mMemberGiftCard表
// @Tags MMemberGiftCard
// @Summary 删除mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberGiftCard true "删除mMemberGiftCard表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mMemberGiftCard/deleteMMemberGiftCard [delete]
func (mMemberGiftCardApi *MMemberGiftCardApi) DeleteMMemberGiftCard(c *gin.Context) {
	ID := c.Query("ID")
	err := mMemberGiftCardService.DeleteMMemberGiftCard(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMMemberGiftCardByIds 批量删除mMemberGiftCard表未核销过的记录
// @Tags MMemberGiftCard
// @Summary 批量删除mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mMemberGiftCard/deleteMMemberGiftCardByIds [delete]
func (mMemberGiftCardApi *MMemberGiftCardApi) DeleteMMemberGiftCardByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mMemberGiftCardService.DeleteMMemberGiftCardByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMMemberGiftCard 更新mMemberGiftCard表
// @Tags MMemberGiftCard
// @Summary 更新mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberGiftCard true "更新mMemberGiftCard表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mMemberGiftCard/updateMMemberGiftCard [put]
func (mMemberGiftCardApi *MMemberGiftCardApi) UpdateMMemberGiftCard(c *gin.Context) {
	var mMemberGiftCard member.MMemberGiftCard
	err := c.ShouldBindJSON(&mMemberGiftCard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if mMemberGiftCard.CardKey == nil {
		response.FailWithMessage("卡密不能为空", c)
		return
	}

	remMemberGiftCard, err := mMemberGiftCardService.GetMMemberGiftCardByCardKey(*mMemberGiftCard.CardKey)
	if err != nil {
		global.GVA_LOG.Error("查询礼券失败!", zap.Error(err))
		response.FailWithMessage("查询礼券失败:"+err.Error(), c)
		return
	}

	//累加本次核销次数
	usedTimes := *remMemberGiftCard.UsedTimes + *mMemberGiftCard.UsedTimes
	remMemberGiftCard.UsedTimes = &usedTimes
	remainTimes := *remMemberGiftCard.TotalTimes - usedTimes

	if remainTimes < 0 {
		response.FailWithMessage("超过了该礼券卡可核销次数", c)
		return
	}
	remMemberGiftCard.RemainTimes = &remainTimes

	userid := utils.GetUserID(c)
	username := utils.GetUserName(c)
	signoff := fmt.Sprintf("【%s】核销ID:%d,核销人:%s,核销次数:%d  \n", time.Now().Format("2006-01-02 15:04:05"), userid, username, *mMemberGiftCard.UsedTimes)

	if remMemberGiftCard.SignOffLog != nil {
		signoff = signoff + *remMemberGiftCard.SignOffLog
	}
	remMemberGiftCard.SignOffLog = &signoff

	err = mMemberGiftCardService.UpdateMMemberGiftCard(remMemberGiftCard)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMMemberGiftCard 用id查询mMemberGiftCard表
// @Tags MMemberGiftCard
// @Summary 用id查询mMemberGiftCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mMemberGiftCard表"
// @Success 200 {object} response.Response{data=member.MMemberGiftCard,msg=string} "查询成功"
// @Router /mMemberGiftCard/findMMemberGiftCard [get]
func (mMemberGiftCardApi *MMemberGiftCardApi) FindMMemberGiftCard(c *gin.Context) {
	ID := c.Query("ID")
	remMemberGiftCard, err := mMemberGiftCardService.GetMMemberGiftCard(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remMemberGiftCard, c)
}

// GetMMemberGiftCardList 分页获取mMemberGiftCard表列表
// @Tags MMemberGiftCard
// @Summary 分页获取mMemberGiftCard表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberGiftCardSearch true "分页获取mMemberGiftCard表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mMemberGiftCard/getMMemberGiftCardList [get]
func (mMemberGiftCardApi *MMemberGiftCardApi) GetMMemberGiftCardList(c *gin.Context) {
	var pageInfo memberReq.MMemberGiftCardSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mMemberGiftCardService.GetMMemberGiftCardInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	for i := 0; i < len(list); i++ {
		if len(*list[i].CardKey) >= 20 {
			tmpCardKey := *list[i].CardKey

			tmpCardKey = tmpCardKey[:8] + "**********" + tmpCardKey[18:]
			list[i].CardKey = &tmpCardKey
		}
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetMMemberGiftCardPublic 不需要鉴权的mMemberGiftCard表接口
// @Tags MMemberGiftCard
// @Summary 不需要鉴权的mMemberGiftCard表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberGiftCard/getMMemberGiftCardPublic [get]
func (mMemberGiftCardApi *MMemberGiftCardApi) GetMMemberGiftCardPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mMemberGiftCardService.GetMMemberGiftCardPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mMemberGiftCard表接口信息",
	}, "获取成功", c)
}
