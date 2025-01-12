package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	"go.uber.org/zap"
)

type MMemberLogService struct{}

// CreateMMemberLog 创建mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) CreateMMemberLog(mMemberLog *member.MMemberLog) (err error) {
	err = global.GVA_DB.Create(mMemberLog).Error
	return err
}

// DeleteMMemberLog 删除mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) DeleteMMemberLog(ID string) (err error) {
	err = global.GVA_DB.Delete(&member.MMemberLog{}, "id = ?", ID).Error
	return err
}

// DeleteMMemberLogByIds 批量删除mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) DeleteMMemberLogByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]member.MMemberLog{}, "id in ?", IDs).Error
	return err
}

// UpdateMMemberLog 更新mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) UpdateMMemberLog(mMemberLog member.MMemberLog) (err error) {
	err = global.GVA_DB.Model(&member.MMemberLog{}).Where("id = ?", mMemberLog.ID).Updates(&mMemberLog).Error
	return err
}

// GetMMemberLog 根据ID获取mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) GetMMemberLog(ID string) (mMemberLog member.MMemberLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mMemberLog).Error
	return
}

// GetMMemberLogInfoList 分页获取mMemberLog表记录
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) GetMMemberLogInfoList(info memberReq.MMemberLogSearch) (list []member.MMemberLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	global.GVA_LOG.Info("info.Status", zap.Any("info", info))

	// 创建db
	db := global.GVA_DB.Model(&member.MMemberLog{})
	var mMemberLogs []member.MMemberLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MemberUid != nil {
		db = db.Where("member_uid = ?", *info.MemberUid)
	}
	if info.OpType != nil && *info.OpType != "" {
		db = db.Where("op_type = ?", *info.OpType)
	}
	if info.PayWays != nil && *info.PayWays != "" {
		db = db.Where("pay_ways = ?", *info.PayWays)
	}
	global.GVA_LOG.Info("info.Status", zap.Any("info", info))
	if info.Status != nil {
		if *info.Status == 1 {
			db = db.Where("status = ?", *info.Status)
		} else {
			db = db.Where("status IS NULL or status <> ?", 1)
		}
	}
	if info.OperatorId != nil {
		db = db.Where("operator_id = ?", *info.OperatorId)
	}
	if info.OperatorName != nil && *info.OperatorName != "" {
		db = db.Where("operator_name LIKE ?", "%"+*info.OperatorName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mMemberLogs).Error
	return mMemberLogs, total, err
}
func (mMemberLogService *MMemberLogService) GetMMemberLogPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// UpdateCheckStatus 对账状态修改接口
// Author [yourname](https://github.com/yourname)
func (mMemberLogService *MMemberLogService) UpdateCheckStatus() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&member.MMemberLog{})
	return db.Error
}
