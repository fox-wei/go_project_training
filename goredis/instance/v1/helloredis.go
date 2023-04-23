package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client //!全局客户端
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
}

func main() {
	r := gin.New()

	//?读取数据
	r.GET("/:key", getData)
	r.POST("/", writeData)

	r.Run()

}

// *获取数据
func getData(ctx *gin.Context) {

	//?路由参数获取key
	key := ctx.Param("key")
	//*redis获取数据
	result, err := client.Get(key).Result()

	if err == redis.Nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	} else if err != nil {
		//*日志记录
		log.Printf("redis get error:%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "redis fail to get data"})
		return
	}

	//*json数据反序列化
	var acount Account
	err = json.Unmarshal([]byte(result), &acount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fail to parse data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": acount})
}

// ?数据缓存在redis
func writeData(ctx *gin.Context) {
	var acount Account

	//*数据绑定
	if err := ctx.ShouldBind(&acount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	//*序列化为json
	dataJson, err := json.Marshal(acount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fail to serialize data"})
		return
	}

	//?缓存信息
	err = client.Set(acount.Id, dataJson, 0).Err()
	if err != nil {
		log.Printf("set data to redis error:%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fail to cache data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": acount.Id})
}

type Account struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
