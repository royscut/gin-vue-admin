package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MMemberLogSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	MemberUid      *uint      `json:"memberUid" form:"memberUid" `
	OpType         *string    `json:"opType" form:"opType" `
	PayWays        *string    `json:"payWays" form:"payWays" `
	Status         *int       `json:"status" form:"status" `
	OperatorId     *int       `json:"operatorId" form:"operatorId" `
	OperatorName   *string    `json:"operatorName" form:"operatorName" `
	request.PageInfo
}
