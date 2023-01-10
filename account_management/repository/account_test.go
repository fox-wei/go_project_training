package repository

import (
	"fmt"
	"log"
	"testing"

	"github.com/fox-wei/go_project_training/account_management/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestAccountModeRepo(t *testing.T) {
	DB := connectDataBase()
	defer DB.Close()

	repo := AccountModeRepo{
		DB: model.DataBase{MyDB: DB},
	}

	t.Run("create", func(t *testing.T) {
		account := model.Account{
			AccountId:   "22",
			AccountName: "hy",
			Password:    "222324",
		}

		err := repo.CreatedAccount(account)

		if err != nil {
			t.Errorf("wnat:nil got:%v", err)
		}
	})

	t.Run("list", func(t *testing.T) {

		accounts, count, err := repo.ListAccount(0, 2)

		if err != nil {
			t.Errorf("err:%v", err)
			return
		}

		if count != 3 {
			t.Errorf("got:%v", accounts)
		}
	})

	t.Run("get", func(t *testing.T) {
		got, err := repo.GetAccountById("1")

		if err != nil {
			t.Errorf("err:%v\n", err)
		}

		want := "1"

		if got.AccountId != want {
			t.Errorf("got:%v\n want:%v", got, want)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := repo.DeleteAccount("1")

		if err != nil {
			t.Errorf("error:%v", err)
		}
	})

	t.Run("delete2", func(t *testing.T) {
		err := repo.DeleteAccountById("12")

		if err != nil {
			t.Errorf("error:%v", err)
		}
	})

	t.Run("update", func(t *testing.T) {
		account := model.Account{
			AccountId:   "2",
			AccountName: "yi-li",
		}

		err := repo.UpdateAccount(account)

		if err != nil {
			t.Errorf("err:%v", err)
		}
	})

	t.Run("account", func(t *testing.T) {
		got, err := repo.GetAccount("yi-li")
		if err != nil {
			t.Errorf("err:%v", err)
		}

		if got.AccountId != "2" {
			t.Errorf("got:%s want:2", got.AccountId)
		}
	})
}

//?连接数据库
func connectDataBase() *gorm.DB {
	//?用户名，密码，addr,数据库名
	d := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "130126*", "127.0.0.1:3306", "life")
	db, err := gorm.Open("mysql", d)
	if err != nil {
		// panic("connect database failed")
		log.Fatalf("connect error:%v\n", err)
	}
	db.SingularTable(true) //*gorm默认表名复数，取消默认复数
	return db
}
