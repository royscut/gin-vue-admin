// 自动生成模板MMemberCard
package member

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mMemberCard表 结构体  MMemberCard
type MMemberCard struct {
	global.GVA_MODEL
	MemberId    *uint      `json:"memberId" form:"memberId" gorm:"column:member_id;comment:会员ID;size:19;" binding:"required"` //会员ID
	TotalTimes  *int       `json:"totalTimes" form:"totalTimes" gorm:"column:total_times;comment:总次数;"`                       //总次数
	RemainTimes *int       `json:"remainTimes" form:"remainTimes" gorm:"column:remain_times;comment:剩余次数;"`                   //剩余次数
	UsedTimes   *int       `json:"usedTimes" form:"usedTimes" gorm:"column:used_times;comment:已消费次数;"`                        //消费次数
	CardPic     *string    `json:"cardPic" form:"cardPic" gorm:"column:card_pic;comment:会员卡图片;size:255;"`                     //会员卡图片
	ExpireTime  *time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:过期时间;"`                      //过期时间
	Remark      *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                            //备注
}

// TableName mMemberCard表 MMemberCard自定义表名 m_member_card
func (MMemberCard) TableName() string {
	return "m_member_card"
}
