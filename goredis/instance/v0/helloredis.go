package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	r := gin.New()

	//?读取数据
	r.GET("/:key", getData)
	r.POST("/", writerData)

	r.Run()

}

// *获取数据
func getData(ctx *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//?路由参数获取key
	key := ctx.Param("key")
	//*redis获取数据
	result, err := client.Get(key).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"redisGetErr": err})
		return
	}

	//*json数据反序列化
	var acount Account
	err = json.Unmarshal([]byte(result), &acount)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"unmarshallErr": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": acount})
}

// ?数据缓存在redis
func writerData(ctx *gin.Context) {
	var acount Account

	//*数据绑定
	if err := ctx.ShouldBind(&acount); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"bindErr": err})
		return
	}

	//*信息序列化为json字符串
	dataJson, err := json.Marshal(acount)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"marshallErr": err})
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//?缓存信息
	err = client.Set(acount.Id, dataJson, 0).Err()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"setErr": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": acount.Id})
}

type Account struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
