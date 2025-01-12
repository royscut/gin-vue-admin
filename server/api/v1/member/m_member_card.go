package member

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/member"
    memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MMemberCardApi struct {}



// CreateMMemberCard 创建mMemberCard表
// @Tags MMemberCard
// @Summary 创建mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberCard true "创建mMemberCard表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mMemberCard/createMMemberCard [post]
func (mMemberCardApi *MMemberCardApi) CreateMMemberCard(c *gin.Context) {
	var mMemberCard member.MMemberCard
	err := c.ShouldBindJSON(&mMemberCard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mMemberCardService.CreateMMemberCard(&mMemberCard)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMMemberCard 删除mMemberCard表
// @Tags MMemberCard
// @Summary 删除mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberCard true "删除mMemberCard表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mMemberCard/deleteMMemberCard [delete]
func (mMemberCardApi *MMemberCardApi) DeleteMMemberCard(c *gin.Context) {
	ID := c.Query("ID")
	err := mMemberCardService.DeleteMMemberCard(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMMemberCardByIds 批量删除mMemberCard表
// @Tags MMemberCard
// @Summary 批量删除mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mMemberCard/deleteMMemberCardByIds [delete]
func (mMemberCardApi *MMemberCardApi) DeleteMMemberCardByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mMemberCardService.DeleteMMemberCardByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMMemberCard 更新mMemberCard表
// @Tags MMemberCard
// @Summary 更新mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body member.MMemberCard true "更新mMemberCard表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mMemberCard/updateMMemberCard [put]
func (mMemberCardApi *MMemberCardApi) UpdateMMemberCard(c *gin.Context) {
	var mMemberCard member.MMemberCard
	err := c.ShouldBindJSON(&mMemberCard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mMemberCardService.UpdateMMemberCard(mMemberCard)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMMemberCard 用id查询mMemberCard表
// @Tags MMemberCard
// @Summary 用id查询mMemberCard表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mMemberCard表"
// @Success 200 {object} response.Response{data=member.MMemberCard,msg=string} "查询成功"
// @Router /mMemberCard/findMMemberCard [get]
func (mMemberCardApi *MMemberCardApi) FindMMemberCard(c *gin.Context) {
	ID := c.Query("ID")
	remMemberCard, err := mMemberCardService.GetMMemberCard(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remMemberCard, c)
}
// GetMMemberCardList 分页获取mMemberCard表列表
// @Tags MMemberCard
// @Summary 分页获取mMemberCard表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query memberReq.MMemberCardSearch true "分页获取mMemberCard表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mMemberCard/getMMemberCardList [get]
func (mMemberCardApi *MMemberCardApi) GetMMemberCardList(c *gin.Context) {
	var pageInfo memberReq.MMemberCardSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mMemberCardService.GetMMemberCardInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetMMemberCardPublic 不需要鉴权的mMemberCard表接口
// @Tags MMemberCard
// @Summary 不需要鉴权的mMemberCard表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mMemberCard/getMMemberCardPublic [get]
func (mMemberCardApi *MMemberCardApi) GetMMemberCardPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mMemberCardService.GetMMemberCardPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mMemberCard表接口信息",
    }, "获取成功", c)
}
