package service

import (
	"errors"

	"github.com/fox-wei/go_project_training/account_management/model"
	"github.com/fox-wei/go_project_training/account_management/repository"
)

type AccountService struct {
	Repo *repository.AccountModeRepo
}

func (as *AccountService) CreatedAccount(account model.Account) error {
	return as.Repo.CreatedAccount(account)
}

func (as *AccountService) DeleteAccount(id string) error {
	return as.Repo.DeleteAccount(id)
}

func (as *AccountService) UpdateAccount(account model.Account) error {
	//?查询是否存在
	accountInfo, err := as.Repo.GetAccountById(account.AccountId)

	if err != nil {
		return err
	}

	if accountInfo.AccountName == "" {
		return errors.New("accountName isn't exist")
	}
	return as.Repo.UpdateAccount(account)
}

func (as *AccountService) ListAccount(offset, limit int) ([]*model.AccountInfo, uint64, error) {
	infos := make([]*model.AccountInfo, 0)

	accounts, count, err := as.Repo.ListAccount(offset, limit)

	if err != nil {
		return nil, count, err
	}

	for _, items := range accounts {
		info := &model.AccountInfo{
			AccountId:   items.AccountId,
			AccountName: items.AccountName,
			Password:    items.Password,
		}

		infos = append(infos, info)
	}

	return infos, count, nil
}

func (as *AccountService) GetAccount(name string) (model.Account, error) {
	return as.Repo.GetAccount(name)
}
