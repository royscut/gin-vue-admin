package member

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MMemberInfoApi
	MMemberCardApi
	MMemberLogApi
	MMemberGiftCardApi
}

var (
	mMemberInfoService     = service.ServiceGroupApp.MemberServiceGroup.MMemberInfoService
	mMemberCardService     = service.ServiceGroupApp.MemberServiceGroup.MMemberCardService
	mMemberLogService      = service.ServiceGroupApp.MemberServiceGroup.MMemberLogService
	mMemberGiftCardService = service.ServiceGroupApp.MemberServiceGroup.MMemberGiftCardService
)
