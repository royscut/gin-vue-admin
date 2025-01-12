package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MMemberInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	MemberId       *uint      `json:"memberId" form:"memberId" `
	Username       *string    `json:"username" form:"username" `
	Mobile         *string    `json:"mobile" form:"mobile" `
	request.PageInfo
}
