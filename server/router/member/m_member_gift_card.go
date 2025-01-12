package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MMemberGiftCardRouter struct {}

// InitMMemberGiftCardRouter 初始化 mMemberGiftCard表 路由信息
func (s *MMemberGiftCardRouter) InitMMemberGiftCardRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mMemberGiftCardRouter := Router.Group("mMemberGiftCard").Use(middleware.OperationRecord())
	mMemberGiftCardRouterWithoutRecord := Router.Group("mMemberGiftCard")
	mMemberGiftCardRouterWithoutAuth := PublicRouter.Group("mMemberGiftCard")
	{
		mMemberGiftCardRouter.POST("createMMemberGiftCard", mMemberGiftCardApi.CreateMMemberGiftCard)   // 新建mMemberGiftCard表
		mMemberGiftCardRouter.DELETE("deleteMMemberGiftCard", mMemberGiftCardApi.DeleteMMemberGiftCard) // 删除mMemberGiftCard表
		mMemberGiftCardRouter.DELETE("deleteMMemberGiftCardByIds", mMemberGiftCardApi.DeleteMMemberGiftCardByIds) // 批量删除mMemberGiftCard表
		mMemberGiftCardRouter.PUT("updateMMemberGiftCard", mMemberGiftCardApi.UpdateMMemberGiftCard)    // 更新mMemberGiftCard表
	}
	{
		mMemberGiftCardRouterWithoutRecord.GET("findMMemberGiftCard", mMemberGiftCardApi.FindMMemberGiftCard)        // 根据ID获取mMemberGiftCard表
		mMemberGiftCardRouterWithoutRecord.GET("getMMemberGiftCardList", mMemberGiftCardApi.GetMMemberGiftCardList)  // 获取mMemberGiftCard表列表
	}
	{
	    mMemberGiftCardRouterWithoutAuth.GET("getMMemberGiftCardPublic", mMemberGiftCardApi.GetMMemberGiftCardPublic)  // mMemberGiftCard表开放接口
	}
}
