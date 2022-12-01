package mode

import (
	"fmt"

	mylog "github.com/fox-wei/go_project_training/life_revive/MyLog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DataBase struct {
	MyDB *gorm.DB
}

var DB *DataBase

//?初始化DB
func (db *DataBase) Init() {
	DB = &DataBase{
		MyDB: GetMySqlDB(),
	}
}

func GetMySqlDB() *gorm.DB {
	return InitSeltDB()
}

func InitSeltDB() *gorm.DB {
	db := openDB(viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.addr"),
		viper.GetString("database.name"))

	return db
}

func (db *DataBase) Close() {
	DB.MyDB.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	//?用户名，密码，addr,数据库名
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, name)
	db, err := gorm.Open("mysql", config)

	if err != nil {
		mylog.Log.Errorf("Database connection failed. Database name:%s error:%s", name, err.Error())
		return nil
	}
	mylog.Log.Info("连接数据成功")
	setupDB(db)
	db.SingularTable(true) //*取消表名复数
	return db

}

func setupDB(db *gorm.DB) {
	//*设置闲置连接数
	db.DB().SetMaxIdleConns(5)
}
