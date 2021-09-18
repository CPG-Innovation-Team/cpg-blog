package common

import (
	"cpg-blog/pkg/snowflake"
	"github.com/spf13/viper"
)

var (
	Snowflake *snowflake.SnowFlake
)
func init(){
	viper.SetDefault("uuid.start_time", "2021-09-18 00:00:00 +0000")
	Snowflake = snowflake.NewWith(viper.GetTime("uuid.start_time"))
}
