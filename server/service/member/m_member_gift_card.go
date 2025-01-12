package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
)

type MMemberGiftCardService struct{}

// CreateMMemberGiftCard 创建mMemberGiftCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) CreateMMemberGiftCard(mMemberGiftCard *member.MMemberGiftCard) (err error) {
	err = global.GVA_DB.Create(mMemberGiftCard).Error
	return err
}

// DeleteMMemberGiftCard 删除mMemberGiftCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) DeleteMMemberGiftCard(ID string) (err error) {
	err = global.GVA_DB.Delete(&member.MMemberGiftCard{}, "id = ?", ID).Error
	return err
}

// DeleteMMemberGiftCardByIds 批量删除mMemberGiftCard表未核销过的记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) DeleteMMemberGiftCardByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]member.MMemberGiftCard{}, "id in ? AND used_times = 0", IDs).Error
	return err
}

// UpdateMMemberGiftCard 更新mMemberGiftCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) UpdateMMemberGiftCard(mMemberGiftCard member.MMemberGiftCard) (err error) {
	err = global.GVA_DB.Model(&member.MMemberGiftCard{}).Where("id = ?", mMemberGiftCard.ID).Updates(&mMemberGiftCard).Error
	return err
}

// GetMMemberGiftCard 根据ID获取mMemberGiftCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) GetMMemberGiftCard(ID string) (mMemberGiftCard member.MMemberGiftCard, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mMemberGiftCard).Error
	return
}

func (mMemberGiftCardService *MMemberGiftCardService) GetMMemberGiftCardByCardKey(cardkey string) (mMemberGiftCard member.MMemberGiftCard, err error) {
	err = global.GVA_DB.Where("card_key = ?", cardkey).First(&mMemberGiftCard).Error
	return
}

// GetMMemberGiftCardInfoList 分页获取mMemberGiftCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberGiftCardService *MMemberGiftCardService) GetMMemberGiftCardInfoList(info memberReq.MMemberGiftCardSearch) (list []member.MMemberGiftCard, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&member.MMemberGiftCard{})
	var mMemberGiftCards []member.MMemberGiftCard
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CardKey != nil && *info.CardKey != "" {
		db = db.Where("card_key = ?", *info.CardKey)
	}
	if info.RemainTimes != nil {
		db = db.Where("remain_times >= ?", *info.RemainTimes)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mMemberGiftCards).Error
	return mMemberGiftCards, total, err
}
func (mMemberGiftCardService *MMemberGiftCardService) GetMMemberGiftCardPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
