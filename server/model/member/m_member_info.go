
// 自动生成模板MMemberInfo
package member
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mMemberInfo表 结构体  MMemberInfo
type MMemberInfo struct {
    global.GVA_MODEL
    Username  *string `json:"username" form:"username" gorm:"column:username;comment:用户名;size:50;" binding:"required"`  //姓名
    Area  *string `json:"area" form:"area" gorm:"default:hk;column:area;comment:地区;size:10;" binding:"required"`  //地区
    Mobile  *string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:20;" binding:"required"`  //手机号
    Address  *string `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`  //地址
    Remark  *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`  //备注
}


// TableName mMemberInfo表 MMemberInfo自定义表名 m_member_info
func (MMemberInfo) TableName() string {
    return "m_member_info"
}





