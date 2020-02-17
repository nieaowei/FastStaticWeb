/*******************************************************
 *  File        :   instance.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/17 5:30 上午
 *  Notes       :
 *******************************************************/
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gogf/gf/os/gcfg"
	"os"
)

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

func (cfg *Config) GetWriter() interface{} {
	return cfg.ConfigWriter
}

func NewInstance(filepath string) Cfg {
	return &Config{
		filePath:     filepath,
		ConfigReader: NewConfigReader(filepath),
		ConfigWriter: NewConfigWriter(filepath),
		mux:          make(chan byte, 1),
	}
}

func init() {

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
	err = encoder.Encode(cfgw.ConfigWriter)
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

func (cfg *Config) GetStrings(key string) []string {
	cfg.mux <- 1
	res := cfg.ConfigReader.GetStrings(key)
	<-cfg.mux
	return res
}
