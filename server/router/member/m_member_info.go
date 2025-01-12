package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MMemberInfoRouter struct{}

func (s *MMemberInfoRouter) InitMMemberInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mMemberInfoRouter := Router.Group("mMemberInfo").Use(middleware.OperationRecord())
	mMemberInfoRouterWithoutRecord := Router.Group("mMemberInfo")
	mMemberInfoRouterWithoutAuth := PublicRouter.Group("mMemberInfo")
	{
		mMemberInfoRouter.POST("createMMemberInfo", mMemberInfoApi.CreateMMemberInfo)
		mMemberInfoRouter.DELETE("deleteMMemberInfo", mMemberInfoApi.DeleteMMemberInfo)
		mMemberInfoRouter.DELETE("deleteMMemberInfoByIds", mMemberInfoApi.DeleteMMemberInfoByIds)
		mMemberInfoRouter.PUT("updateMMemberInfo", mMemberInfoApi.UpdateMMemberInfo)
		mMemberInfoRouter.POST("addCardTimes", mMemberInfoApi.AddCardTimes)
		mMemberInfoRouter.POST("cutCardTimes", mMemberInfoApi.CutCardTimes)
	}
	{
		mMemberInfoRouterWithoutRecord.GET("findMMemberInfo", mMemberInfoApi.FindMMemberInfo)
		mMemberInfoRouterWithoutRecord.GET("getMMemberInfoList", mMemberInfoApi.GetMMemberInfoList)
	}
	{
		mMemberInfoRouterWithoutAuth.GET("getMMemberInfoPublic", mMemberInfoApi.GetMMemberInfoPublic)
	}
}
