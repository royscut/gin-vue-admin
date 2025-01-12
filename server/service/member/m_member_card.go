package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
)

type MMemberCardService struct{}

// CreateMMemberCard 创建mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) CreateMMemberCard(mMemberCard *member.MMemberCard) (err error) {
	err = global.GVA_DB.Create(mMemberCard).Error
	return err
}

// DeleteMMemberCard 删除mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) DeleteMMemberCard(ID string) (err error) {
	err = global.GVA_DB.Delete(&member.MMemberCard{}, "id = ?", ID).Error
	return err
}

// DeleteMMemberCardByIds 批量删除mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) DeleteMMemberCardByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]member.MMemberCard{}, "id in ?", IDs).Error
	return err
}

// UpdateMMemberCard 更新mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) UpdateMMemberCard(mMemberCard member.MMemberCard) (err error) {
	err = global.GVA_DB.Model(&member.MMemberCard{}).Where("id = ?", mMemberCard.ID).Updates(&mMemberCard).Error
	return err
}

// GetMMemberCard 根据ID获取mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) GetMMemberCard(ID string) (mMemberCard member.MMemberCard, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mMemberCard).Error
	return
}

func (mMemberCardService *MMemberCardService) GetMMemberCardByMemberId(memberid uint) (mMemberCard member.MMemberCard, err error) {
	err = global.GVA_DB.Where("member_id = ?", memberid).First(&mMemberCard).Error
	return
}

// GetMMemberCardInfoList 分页获取mMemberCard表记录
// Author [yourname](https://github.com/yourname)
func (mMemberCardService *MMemberCardService) GetMMemberCardInfoList(info memberReq.MMemberCardSearch) (list []member.MMemberCard, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&member.MMemberCard{})
	var mMemberCards []member.MMemberCard
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MemberId != nil {
		db = db.Where("member_id = ?", *info.MemberId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mMemberCards).Error
	return mMemberCards, total, err
}

func (mMemberCardService *MMemberCardService) GetMMemberCardInfoListByIDs(ids []*uint) (list []member.MMemberCard, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&member.MMemberCard{})
	var mMemberCards []member.MMemberCard

	// 查询 member_id 在 ids.MemberIds 数组中的数据
	if len(ids) > 0 {
		db = db.Where("member_id IN ?", ids)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = db.Limit(200)

	err = db.Find(&mMemberCards).Error
	return mMemberCards, total, err
}

func (mMemberCardService *MMemberCardService) GetMMemberCardPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
