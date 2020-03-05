/*******************************************************
 *  File        :   manager.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:10 下午
 *  Notes       :
 *******************************************************/
package server

import (
	"FastStaticWeb/config"
	"FastStaticWeb/controller"
	"FastStaticWeb/httpServer"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Instance struct {
	config.Cfg
	httpServer  *httpServer.Server
	httpsServer *httpServer.Server
	*config.ConfigWriter
	controlles map[*httpServer.Server]controller.Controller
}

func NewInstance(cfg config.Cfg) *Instance {
	//:= control.NewController(context.Background(), time.Second*3)
	inst := &Instance{
		Cfg:          cfg,
		httpServer:   httpServer.NewInstance("0.0.0.0:80", cfg.GetValue("web.root").(string)),
		ConfigWriter: cfg.GetWriter().(*config.ConfigWriter),
		controlles:   make(map[*httpServer.Server]controller.Controller),
	}
	tempCC := controller.NewController(context.Background(), time.Hour)
	inst.controlles[inst.httpServer] = tempCC
	inst.httpServer.SetController(tempCC)
	if cfg.GetValue("web.https.enable").(bool) { //如果配置文件中启用了,没启用就暂时不实例化
		inst.httpsServer = httpServer.NewInstance("0.0.0.0:443", cfg.GetValue("web.root").(string)).
			SetTlsEnable(true, cfg.GetValue("web.https.filepath").(string))
		inst.controlles[inst.httpsServer] = controller.NewController(context.Background(), time.Hour)
	}
	return inst
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

func (s *Instance) Start() *Instance {
	webconfig := s.GetValue("webconfig.enable").(bool)
	root := s.GetValue("web.root").(string)
	commonFilter := (&requestFilter{}).SetInfo(s.GetStrings("redirect.keywords"))
	//启动非安全服务器
	if webconfig {
		s.httpServer.RegisterFunc("/config", ConfigHandler)
	}
	if s.GetValue("redirect.enable").(bool) {
		s.httpServer.RegisterFilterHandleFunc("/", http.FileServer(http.Dir(root)),
			commonFilter, FilterHandler)
	} else {
		s.httpServer.RegisterHandle("/", http.FileServer(http.Dir(root)))
	}
	go s.httpServer.Start()
	//启动安全服务器
	if s.GetValue("web.https.enable").(bool) {
		if s.httpsServer == nil {
			s.httpsServer = httpServer.NewInstance("0.0.0.0:443", root).
				SetTlsEnable(true, s.GetValue("web.https.filepath").(string))
		}
		if webconfig {
			s.httpsServer.RegisterFunc("/config", ConfigHandler)
		}
		if s.GetValue("redirect.enable").(bool) {
			s.httpsServer.RegisterFilterHandleFunc("/", http.FileServer(http.Dir(root)),
				commonFilter, FilterHandler)
		} else {
			s.httpsServer.RegisterHandle("/", http.FileServer(http.Dir(root)))
		}
		go s.httpsServer.Start()
	}

	return s
}

func (s *Instance) StartDaemon() {
	go s.controlles[s.httpServer].Wait(func() {
		fmt.Println("http server exit.")
	})
	go s.controlles[s.httpsServer].Wait(func() {
		fmt.Println("https server exit.")
	})
	fmt.Println("wait for opration")
	option := ' '
	fmt.Scanf("%c", &option)
	if option == 'q' {
		return
	}
}

func (s *Instance) Stop(ctx context.Context) {
	go s.httpServer.Stop(ctx)
	go s.httpsServer.Stop(ctx)

}

func (inst *Instance) Restart() {

}
