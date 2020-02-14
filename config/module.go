/*******************************************************
 *  File        :   module.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 9:27 下午
 *  Notes       :
 *******************************************************/
package config

type config struct {
	Webconfig *WebconfigConfig `toml:"webconfig"`
	Web *WebConfig	`toml:"web"`

}

type WebconfigConfig struct {
	Enable bool	`toml:"enable"`
}

type HttpsConfig struct {
	Enable bool	`toml:"enable"`
	Filepath string `toml:"filepath"`
}

type WebConfig struct {
	Root string	`toml:"root"`
	Https *HttpsConfig	`toml:"htts"`
	Static *WebStaticConfiig `toml:"static"`
}

type WebStaticConfiig struct {
	Enable bool	`toml:"enable"`
	Path string `toml:"path"`
}

func (cfg *config)SetWebCfg(config *WebConfig) (*config) {
	cfg.Web = config
	return cfg
}

func (cfg *config)SetWebconfigCfg(config *WebconfigConfig) (*config) {
	cfg.Webconfig = config
	return cfg
}

func (cfg *WebConfig)SetRoot(path string) (*WebConfig) {
	cfg.Root=path
	return cfg
}

func (cfg *WebConfig)SetHttps(config *HttpsConfig) (*WebConfig) {
	cfg.Https=config
	return cfg
}

func (cfg *WebConfig)SetStatic(config *WebStaticConfiig) (*WebConfig) {
	cfg.Static=config
	return cfg
}