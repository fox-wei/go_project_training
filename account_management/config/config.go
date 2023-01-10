package config

import (
	"log"
	"path"
	"runtime"

	mylog "github.com/fox-wei/go_project_training/life_revive/MyLog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(name string) error {
	c := Config{
		Name: name,
	}

	//*初始化配置文件
	if err := c.initConit(); err != nil {
		mylog.Log.Error(err)
		return err
	}
	c.watchConfig() //&监控配置文件
	return nil
}

func (c *Config) initConit() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //*指定了配置文件
	} else {
		// log.Printf("else-%s", c.Name)
		//?./不是表示当前代码文件所在目录；而是代表执行程序可执行文件的目录
		// viper.AddConfigPath("./conf") //*默认解析conf/config.yaml
		// viper.SetConfigName("config")

		//?获取项目根目录
		_, fileName, _, _ := runtime.Caller(0)
		root := path.Dir(path.Dir(fileName))

		viper.AddConfigPath(root + "/conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml") //*设置配置文件格式yaml
	//^解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//!监控配置文件并热加载部署
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed:%s", in.Name)
	})
}
