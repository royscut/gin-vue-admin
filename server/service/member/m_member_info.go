package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
)

type MMemberInfoService struct{}

// CreateMMemberInfo 创建mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) CreateMMemberInfo(mMemberInfo *member.MMemberInfo) (err error) {
	err = global.GVA_DB.Create(mMemberInfo).Error
	return err
}

// DeleteMMemberInfo 删除mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) DeleteMMemberInfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&member.MMemberInfo{}, "id = ?", ID).Error
	return err
}

// DeleteMMemberInfoByIds 批量删除mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) DeleteMMemberInfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]member.MMemberInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateMMemberInfo 更新mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) UpdateMMemberInfo(mMemberInfo member.MMemberInfo) (err error) {
	err = global.GVA_DB.Model(&member.MMemberInfo{}).Where("id = ?", mMemberInfo.ID).Updates(&mMemberInfo).Error
	return err
}

// GetMMemberInfo 根据ID获取mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) GetMMemberInfo(ID string) (mMemberInfo member.MMemberInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mMemberInfo).Error
	return
}

// GetMMemberInfoInfoList 分页获取mMemberInfo表记录
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) GetMMemberInfoInfoList(info memberReq.MMemberInfoSearch) (list []member.MMemberInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&member.MMemberInfo{})
	var mMemberInfos []member.MMemberInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != nil && *info.Username != "" {
		db = db.Where("username LIKE ?", "%"+*info.Username+"%")
	}
	if info.Mobile != nil && *info.Mobile != "" {
		db = db.Where("mobile = ?", *info.Mobile)
	}

	if info.MemberId != nil {
		db = db.Where("id = ?", *info.MemberId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mMemberInfos).Error
	return mMemberInfos, total, err
}
func (mMemberInfoService *MMemberInfoService) GetMMemberInfoPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AddCardTimes 充值卡券次数
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) AddCardTimes() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&member.MMemberInfo{})
	return db.Error
}

// CutCardTimes 扣减卡券次数
// Author [yourname](https://github.com/yourname)
func (mMemberInfoService *MMemberInfoService) CutCardTimes() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&member.MMemberInfo{})
	return db.Error
}
