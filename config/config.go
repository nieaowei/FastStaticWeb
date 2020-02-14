/*******************************************************
 *  File        :   config.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:10 下午
 *  Notes       :
 *******************************************************/
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gogf/gf/os/gcfg"
	"os"
)

type Cfg interface {
	WriteConfig() error
	GetValue(key string) interface{}
}

type ConfigWriter struct {
	config
}

type ConfigReader struct {
	*gcfg.Config
}

type Config struct {
	filePath string
	*ConfigReader
	*ConfigWriter
	mux chan byte
}

const defaultConfigPath = "config/config.toml"

var (
	defaultCfg = NewConfig(defaultConfigPath)
)

func NewConfig(filepath string) Cfg {
	return &Config{
		filePath:     filepath,
		ConfigReader: NewConfigReader(defaultConfigPath),
		ConfigWriter: NewConfigWriter(defaultConfigPath),
		mux:          make(chan byte, 1),
	}
}

func init() {

}

func WriteDefaultConfig() error {
	return defaultCfg.WriteConfig()
}

func NewConfigReader(filepath string) *ConfigReader {
	return &ConfigReader{Config: gcfg.Instance(filepath)}
}

func NewConfigWriter(filepath string) *ConfigWriter {
	return &ConfigWriter{
		config: config{},
	}
}

func (cfgw *Config) WriteConfig() error {
	cfgw.mux <- 1
	file, err := os.OpenFile(cfgw.filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("read config file")
		return err
	}
	defer func() {
		defer func() {
			r := recover()
			if r != nil {
				<-cfgw.mux
			}
		}()
		err := file.Close()
		if err != nil {
			fmt.Println()
			return
		}
		<-cfgw.mux
	}()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(cfgw.config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (cfg *Config) GetValue(key string) interface{} {
	cfg.mux <- 1
	res := cfg.Get(key)
	<-cfg.mux
	return res
}

func Get(key string) interface{} {
	return defaultCfg.GetValue(key)
}

func DefaultConfig() Cfg {
	return defaultCfg
}
