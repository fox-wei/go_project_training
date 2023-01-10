package repository

import (
	"github.com/fox-wei/go_project_training/account_management/model"
)

type AccountModeRepo struct {
	DB model.DataBase
}

//?新建account
func (m *AccountModeRepo) CreatedAccount(account model.Account) error {
	return m.DB.MyDB.Create(&account).Error
}

//?删除一个acount
func (m *AccountModeRepo) DeleteAccount(accountId string) error {
	//?软删除
	return m.DB.MyDB.Where("account_id=?", accountId).Delete(&model.Account{}).Error
}

//?硬删除
func (m *AccountModeRepo) DeleteAccountById(id string) error {
	err := m.DB.MyDB.Unscoped().Where("id=?", id).Delete(&model.Account{}).Error
	return err
}

//*根据id查询
func (m *AccountModeRepo) GetAccountById(id string) (model.Account, error) {
	var account model.Account
	err := m.DB.MyDB.Where("account_id=?", id).Find(&account).Error

	if err != nil {
		return account, err
	}
	return account, nil
}

//*账户更新
func (m *AccountModeRepo) UpdateAccount(account model.Account) error {
	err := m.DB.MyDB.Model(model.Account{}).Where("account_id=?", account.AccountId).
		Updates(&account).Error

	return err
}

//*账户列表
func (m *AccountModeRepo) ListAccount(offset, limit int) ([]*model.Account, uint64, error) {

	accounts := make([]*model.Account, 0)
	var count uint64

	if err := m.DB.MyDB.Model(&model.Account{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := m.DB.MyDB.Model(&model.Account{}).Limit(limit).Offset(offset).
		Order("id desc").Find(&accounts).Error
	if err != nil {
		return nil, count, err
	}

	return accounts, count, nil
}

//?根据名字返回账户信息
func (m *AccountModeRepo) GetAccount(name string) (model.Account, error) {
	var ms model.Account
	err := m.DB.MyDB.Where("account_name=?", name).First(&ms).Error
	if err != nil {
		return ms, err
	}
	return ms, nil
}
