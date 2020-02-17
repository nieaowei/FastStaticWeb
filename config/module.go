/*******************************************************
 *  File        :   module.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 9:27 下午
 *  Notes       :
 *******************************************************/
package config

type config struct {
	Webconfig      WebconfigConfig      `toml:"webconfig"`
	Web            WebConfig            `toml:"web"`
	FilterRedirect FilterRedirectConfig `toml:"filterredirect"`
}

type FilterRedirectConfig struct {
	Enable   bool     `toml:"enable"`
	Keywords []string `toml:"keywords"`
	ReUrl    string   `toml:"reurl"`
}

type WebconfigConfig struct {
	Enable bool `toml:"enable"`
}

type HttpsConfig struct {
	Enable   bool   `toml:"enable"`
	Filepath string `toml:"filepath"`
}

type WebConfig struct {
	Root   string           `toml:"root"`
	Https  HttpsConfig      `toml:"https"`
	Static WebStaticConfiig `toml:"static"`
}

type WebStaticConfiig struct {
	Enable bool   `toml:"enable"`
	Path   string `toml:"path"`
}

func (cfg *config) SetFilterRedirectCfg(config WebConfig) *config {
	cfg.Web = config
	return cfg
}

func (cfg *config) SetFilterEnable(ok bool) *config {
	cfg.FilterRedirect.Enable = ok
	return cfg
}

func (cfg *config) SetFilterKeyWords(keywords []string) *config {
	cfg.FilterRedirect.Keywords = keywords
	return cfg
}

func (cfg *config) SetFilterReUrl(url string) *config {
	cfg.FilterRedirect.ReUrl = url
	return cfg
}

func (cfg *config) SetWebCfg(config WebConfig) *config {
	cfg.Web = config
	return cfg
}

func (cfg *config) SetWebconfigCfg(config WebconfigConfig) *config {
	cfg.Webconfig = config
	return cfg
}

func (cfg *config) SetHttps(config HttpsConfig) *config {
	cfg.Web.Https = config
	return cfg
}

func (cfg *config) SetStatic(config WebStaticConfiig) *config {
	cfg.Web.Static = config
	return cfg
}

func (cfg *config) SetWebCfgRoot(path string) *config {
	cfg.Web.Root = path
	return cfg
}

func (cfg *config) SetHttpsEnable(ok bool) *config {
	cfg.Web.Https.Enable = ok
	return cfg
}

func (cfg *config) SetHttpsFilePath(path string) *config {
	cfg.Web.Https.Filepath = path
	return cfg
}

func (cfg *config) SetStaticEnable(ok bool) *config {
	cfg.Web.Static.Enable = ok
	return cfg
}

func (cfg *config) SetStaticPath(path string) *config {
	cfg.Web.Static.Path = path
	return cfg
}
