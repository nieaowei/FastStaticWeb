/*******************************************************
 *  File        :   default.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/17 5:24 上午
 *  Notes       :
 *******************************************************/
package config

const defaultConfigPath = "config/config.toml"

var (
	defaultCfg = NewInstance(defaultConfigPath)
)

func WriteDefaultConfig() error {
	return defaultCfg.WriteConfig()
}

func Get(key string) interface{} {
	return defaultCfg.GetValue(key)
}

func DefaultConfig() Cfg {
	return defaultCfg
}
