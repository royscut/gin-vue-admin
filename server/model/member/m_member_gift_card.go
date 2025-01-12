// 自动生成模板MMemberGiftCard
package member

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mMemberGiftCard表 结构体  MMemberGiftCard
type MMemberGiftCard struct {
	global.GVA_MODEL
	CardKey     *string    `json:"cardKey" form:"cardKey" gorm:"column:card_key;comment:卡券Key;size:255;"`           //卡券Key
	TotalTimes  *int       `json:"totalTimes" form:"totalTimes" gorm:"column:total_times;comment:总次数;size:19;"`     //总次数
	RemainTimes *int       `json:"remainTimes" form:"remainTimes" gorm:"column:remain_times;comment:剩余次数;size:19;"` //剩余次数
	UsedTimes   *int       `json:"usedTimes" form:"usedTimes" gorm:"column:used_times;comment:核销次数;size:19;"`       //核销次数
	CardPic     *string    `json:"cardPic" form:"cardPic" gorm:"column:card_pic;comment:卡券图片;size:255;"`            //卡券图片
	ExpireTime  *time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:过期时间;"`            //过期时间
	Remark      *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                  //备注
	SignOffLog  *string    `json:"signOffLog" form:"signOffLog" gorm:"column:sign_off_log;comment:核销日志;size:255;"`  //备注
}

// TableName mMemberGiftCard表 MMemberGiftCard自定义表名 m_member_gift_card
func (MMemberGiftCard) TableName() string {
	return "m_member_gift_card"
}
