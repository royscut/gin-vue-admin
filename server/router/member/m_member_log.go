package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MMemberLogRouter struct{}

func (s *MMemberLogRouter) InitMMemberLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mMemberLogRouter := Router.Group("mMemberLog").Use(middleware.OperationRecord())
	mMemberLogRouterWithoutRecord := Router.Group("mMemberLog")
	mMemberLogRouterWithoutAuth := PublicRouter.Group("mMemberLog")
	{
		mMemberLogRouter.POST("createMMemberLog", mMemberLogApi.CreateMMemberLog)
		mMemberLogRouter.DELETE("deleteMMemberLog", mMemberLogApi.DeleteMMemberLog)
		mMemberLogRouter.DELETE("deleteMMemberLogByIds", mMemberLogApi.DeleteMMemberLogByIds)
		mMemberLogRouter.PUT("updateMMemberLog", mMemberLogApi.UpdateMMemberLog)
		mMemberLogRouter.POST("updateCheckStatus", mMemberLogApi.UpdateCheckStatus)
	}
	{
		mMemberLogRouterWithoutRecord.GET("findMMemberLog", mMemberLogApi.FindMMemberLog)
		mMemberLogRouterWithoutRecord.GET("getMMemberLogList", mMemberLogApi.GetMMemberLogList)
	}
	{
		mMemberLogRouterWithoutAuth.GET("getMMemberLogPublic", mMemberLogApi.GetMMemberLogPublic)
	}
}
