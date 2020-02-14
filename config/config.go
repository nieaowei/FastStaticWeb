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

type ConfigWriter struct {
	filePath string
	*config
}

const defaultConfigPath = "config/config.toml"

var (
	defaultCfgRead = gcfg.Instance(defaultConfigPath)
	defaultCfgWriter *ConfigWriter = &ConfigWriter{}
)

func init() {
	defaultCfgWriter.SetFilePath(defaultConfigPath).
		SetWebCfg((&WebConfig{}).
			SetRoot("/").
			SetHttps(&HttpsConfig{
				Enable:   false,
				Filepath: "",
			}).
			SetStatic(&WebStaticConfiig{
				Enable: false,
				Path:   "",
			})).
		SetWebconfigCfg(&WebconfigConfig{Enable:false})
}

func WriteDefaultConfig()(error) {
	return defaultCfgWriter.WriteConfig()
}

func (cfgw *ConfigWriter)WriteConfig()( error) {
	file, err := os.OpenFile(cfgw.filePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("read config file")
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println()
			return
		}
	}()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(cfgw.config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (cfgw *ConfigWriter)SetFilePath(path string) (*ConfigWriter) {
	cfgw.filePath=path
	return cfgw
}

