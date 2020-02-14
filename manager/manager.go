/*******************************************************
 *  File        :   manager.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:10 下午
 *  Notes       :
 *******************************************************/
package manager

import (
	"FastStaticWeb/config"
	"FastStaticWeb/server"
	"context"
	"html/template"
	"net/http"
	"time"
)

type Mgr struct {
	server       *server.Instance
	configWriter config.ConfigWriter
	ctx          context.Context
	cancel       context.CancelFunc
}

var defaultMgr = &Mgr{
	server:       server.NewServer("0.0.0.0:80", config.Get("web.root").(string), &http.ServeMux{}),
	configWriter: config.ConfigWriter{},
	ctx:          nil,
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

	default: //get请求默认返回页面
		enable := config.Get("webconfig.enable").(bool)
		if !enable {
			return
		}
		file, _ := template.ParseFiles("config.html")
		file.Execute(w, nil)
	}
}

func ScanConfig() {

}

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	defaultMgr.ctx = ctx
	defaultMgr.cancel = cancel
}

func Start() {
	defaultMgr.server.RegisterFunc("/config", ConfigHandler).SetContext(defaultMgr.ctx).Start()
	time.Sleep(time.Second * 3)
	Restart()
}

func Restart() {
	defaultMgr.server.Stop()
	defaultMgr.server.SetContext(defaultMgr.ctx).Start()

}

func Cancel() {
	defaultMgr.cancel()
}
