package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MMemberGiftCardSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CardKey        *string    `json:"cardKey" form:"cardKey" `
	RemainTimes    *int       `json:"remainTimes" form:"remainTimes" `
	request.PageInfo
}
