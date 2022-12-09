package main

import (
	"fmt"
	"log"

	"github.com/fox-wei/go_project_training/life_revive/conf"
	"github.com/fox-wei/go_project_training/life_revive/config"
	"github.com/fox-wei/go_project_training/life_revive/handler"
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	dbservice "github.com/fox-wei/go_project_training/life_revive/service/db_service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

//?解析配置文件
func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("database init")

	conf := &conf.DBConf{
		Host:     viper.GetString("database.addr"),
		User:     viper.GetString("database.username"),
		Passwrod: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	DB = connectDataBase(conf.User, conf.Passwrod, conf.Host, conf.DbName)

	fmt.Println("database init end...")

}

//*程序逻辑初始化
func initHandler() {
	//?折扣
	DiscountHandler = handler.DiscountHandler{
		Srv: &dbservice.DiscountService{
			Repo: &repository.Discount{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}

	//?餐馆导航
	RestaurantNavHandler = handler.RestaurantNavHandler{
		Srv: &dbservice.RestaurantService{
			Repo: &repository.RestauranNavRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}

	//!餐馆信息
	HotelDetailHandler = handler.HotelDetailHandler{
		Srv: &dbservice.HotelService{
			Repo: &repository.HotelRepo{
				DB: model.DataBase{MyDB: DB},
			},

			TeamRepo: &repository.TeamRepo{
				DB: model.DataBase{MyDB: DB},
			},

			SuggestFoodRepo: &repository.SuggestFoodRepo{
				DB: model.DataBase{MyDB: DB},
			},

			CommentTagRepo: &repository.CommentTagRepo{
				DB: model.DataBase{MyDB: DB},
			},

			CommentRepo: &repository.CommentRepo{
				DB: model.DataBase{MyDB: DB},
			},

			MarketRepo: &repository.MarketRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	//?团购信息
	TeamDetailHandler = handler.TeamDetailHandler{
		Srv: &dbservice.TeamService{
			Repo: &repository.TeamRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	//?订单查询
	OrderSeatHandler = handler.OrderSeatHandler{
		Srv: &dbservice.OrderSeatService{
			Repo: &repository.OrderSeatRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	//*外卖
	TakeOutHandler = handler.TakeOutHandler{
		Srv: &dbservice.TakeOutService{
			Repo: &repository.TakeOutItemRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	//?我的设置项
	MeHandler = handler.MeHandler{
		Srv: &dbservice.MeService{
			Repo: repository.MeItemRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	SuggestFoodHandler = handler.SuggestFoodHandler{
		Srv: &dbservice.SuggestFoodService{
			Repo: &repository.SuggestFoodRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	SuggestHandler = handler.SuggestHandler{
		Srv: &dbservice.SuggestService{
			Repo: &repository.SuggestRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	//?猜你喜欢
	GuessHandler = handler.GuessHandler{
		Srv: &dbservice.GuessService{
			Repo: &repository.GuessRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	NavHandler = handler.NavHandler{
		Srv: &dbservice.ListNavItemService{
			Repo: &repository.ListNavItemRepo{
				DB: model.DataBase{MyDB: DB},
			},
		},
	}

	PostTeamOrderHandler = handler.PostTeamOrderHandler{
		Srv: &dbservice.PostTeamOrderService{
			Repo: &repository.TeamPostOrderRepo{
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
