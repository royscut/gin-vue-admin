package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(member.MMemberInfo{}, member.MMemberCard{}, member.MMemberLog{}, member.MMemberGiftCard{})
	if err != nil {
		return err
	}
	return nil
}
