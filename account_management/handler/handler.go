package handler

import (
	"fmt"
	"log"

	"github.com/fox-wei/go_project_training/account_management/conf"
	"github.com/fox-wei/go_project_training/account_management/config"
	"github.com/fox-wei/go_project_training/account_management/model"
	"github.com/fox-wei/go_project_training/account_management/repository"
	"github.com/fox-wei/go_project_training/account_management/service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB            *gorm.DB
	AccountHandle AccountHandler
)

//*viper读取配置文件
func InItVipe() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func InitDB() {
	conf := &conf.DBConf{
		Host:     viper.GetString("database.Host"),
		User:     viper.GetString("database.username"),
		Passwrod: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	DB = connectDataBase(conf.User, conf.Passwrod, conf.Host, conf.DbName)
}

func InitHanlder() {
	AccountHandle = AccountHandler{
		Srv: service.AccountService{
			Repo: &repository.AccountModeRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}
}

//?连接数据库
func connectDataBase(username, password, addr, name string) *gorm.DB {
	//?用户名，密码，addr,数据库名
	d := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, name)
	db, err := gorm.Open("mysql", d)
	if err != nil {
		// panic("connect database failed")
		log.Fatalf("connect error:%v\n", err)
	}
	db.SingularTable(true) //*gorm默认表名复数，取消默认复数
	return db
}
