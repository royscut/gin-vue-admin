// 自动生成模板MMemberLog
package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mMemberLog表 结构体  MMemberLog
type MMemberLog struct {
	global.GVA_MODEL
	MemberUid    *uint   `json:"memberUid" form:"memberUid" gorm:"column:member_uid;comment:会员ID;size:19;" binding:"required"`     //会员ID
	Username     *string `json:"username" form:"username" gorm:"column:username;comment:用户名;size:50;"`                             //用户名
	OpType       *string `json:"opType" form:"opType" gorm:"column:op_type;comment:操作类型;size:10;" binding:"required"`              //操作类型
	OpTimes      *int    `json:"opTimes" form:"opTimes" gorm:"column:op_times;comment:操作次数;size:10;"`                              //操作次数
	PayWays      *string `json:"payWays" form:"payWays" gorm:"column:pay_ways;comment:付款方式;size:20;" binding:"required"`           //付款方式
	PayAmount    *string `json:"payAmount" form:"payAmount" gorm:"column:pay_amount;comment:付款金额;size:20;"`                        //付款金额
	RenewPeriod  *int    `json:"renewPeriod" form:"renewPeriod" gorm:"column:renew_period;comment:续期时间;size:10;"`                  //续期时间
	Status       *int    `json:"status" form:"status" gorm:"column:status;comment:对账状态;size:10;"`                                  //对账状态
	OperatorId   *int    `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:操作人ID;size:19;" binding:"required"` //操作人ID
	OperatorName *string `json:"operatorName" form:"operatorName" gorm:"column:operator_name;comment:操作人姓名;size:50;"`              //操作人姓名
	Remark       *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                   //备注
}

// TableName mMemberLog表 MMemberLog自定义表名 m_member_log
func (MMemberLog) TableName() string {
	return "m_member_log"
}
