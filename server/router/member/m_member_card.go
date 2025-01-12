package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MMemberCardRouter struct {}

// InitMMemberCardRouter 初始化 mMemberCard表 路由信息
func (s *MMemberCardRouter) InitMMemberCardRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mMemberCardRouter := Router.Group("mMemberCard").Use(middleware.OperationRecord())
	mMemberCardRouterWithoutRecord := Router.Group("mMemberCard")
	mMemberCardRouterWithoutAuth := PublicRouter.Group("mMemberCard")
	{
		mMemberCardRouter.POST("createMMemberCard", mMemberCardApi.CreateMMemberCard)   // 新建mMemberCard表
		mMemberCardRouter.DELETE("deleteMMemberCard", mMemberCardApi.DeleteMMemberCard) // 删除mMemberCard表
		mMemberCardRouter.DELETE("deleteMMemberCardByIds", mMemberCardApi.DeleteMMemberCardByIds) // 批量删除mMemberCard表
		mMemberCardRouter.PUT("updateMMemberCard", mMemberCardApi.UpdateMMemberCard)    // 更新mMemberCard表
	}
	{
		mMemberCardRouterWithoutRecord.GET("findMMemberCard", mMemberCardApi.FindMMemberCard)        // 根据ID获取mMemberCard表
		mMemberCardRouterWithoutRecord.GET("getMMemberCardList", mMemberCardApi.GetMMemberCardList)  // 获取mMemberCard表列表
	}
	{
	    mMemberCardRouterWithoutAuth.GET("getMMemberCardPublic", mMemberCardApi.GetMMemberCardPublic)  // mMemberCard表开放接口
	}
}
