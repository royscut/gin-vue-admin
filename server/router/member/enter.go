package member

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	MMemberInfoRouter
	MMemberCardRouter
	MMemberLogRouter
	MMemberGiftCardRouter
}

var (
	mMemberInfoApi     = api.ApiGroupApp.MemberApiGroup.MMemberInfoApi
	mMemberCardApi     = api.ApiGroupApp.MemberApiGroup.MMemberCardApi
	mMemberLogApi      = api.ApiGroupApp.MemberApiGroup.MMemberLogApi
	mMemberGiftCardApi = api.ApiGroupApp.MemberApiGroup.MMemberGiftCardApi
)
