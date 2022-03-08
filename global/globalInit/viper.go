package globalInit

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"sync"
	"time"
)

var (
	once = new(sync.Once)
	config = flag.String("config", "config", "配置文件名称，默认 config")
	LocalDebug = flag.Bool("local", false, "本地启动，默认 false")
	Base    int
)

func init()  {
	ViperInit()
}
func ViperInit() {
	once.Do(func() {
		if !flag.Parsed() {
			flag.Parse()
		}

		rand.Seed(time.Now().UnixNano())

		//配置文件名称
		fmt.Print(*config)
		viper.SetConfigName(*config)
		viper.SetConfigType("toml")
		//配置文件查找路径
		viper.AddConfigPath("/etc/cpg-blog")
		viper.AddConfigPath("$HOME/.cpg-blog")
		viper.AddConfigPath(App.RootDir + "/config")
		//读取配置文件
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		//监控配置文件
		viper.WatchConfig()

	})
	//viper.SetDefault("",16)
	Base = viper.GetInt("uuid.base")
}
