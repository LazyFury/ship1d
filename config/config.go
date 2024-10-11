package config

import (
	"fmt"
	"os"

	"yoomall/core/constants"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	config_file_path := os.Getenv(constants.YOO_CONFIG)
	if config_file_path != "" {
		return GetConfig(config_file_path)
	}
	return GetConfig("./config.yaml")
}

func GetConfig(file string) *viper.Viper {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigFile(file)
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file, err: ", err)
		// gen default config.yaml
		config.WriteConfigAs("./config.yaml")
		os.Exit(1)
	}
	return config
}

var _viper *viper.Viper = NewConfig()

// 常用配置： viper 的用法很难收集配置，记录一些常用的配置，方便以后使用
var Config = struct {
	Port     int
	DEBUG    bool
	MysqlDsn string
}{
	Port:     _viper.GetInt(constants.PORT),
	DEBUG:    _viper.GetBool(constants.DEBUG),
	MysqlDsn: _viper.GetString(constants.MYSQL_DSN),
}
