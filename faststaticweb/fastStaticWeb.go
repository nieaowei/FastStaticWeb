/*******************************************************
 *  File        :   manager.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:10 下午
 *  Notes       :
 *******************************************************/
package faststaticweb

import (
	"FastStaticWeb/config"
	"FastStaticWeb/controller"
	"FastStaticWeb/filter"
	"FastStaticWeb/server"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Instance struct {
	config.Cfg
	httpServer  *server.Server
	httpsServer *server.Server
	*config.ConfigWriter
	controlles map[*server.Server]controller.Controller
}

//var defaultMgr = &Instance{
//	httpsServer: server.NewInstance("0.0.0.0:443", config.Get("web.root").(string)).SetTlsEnable(true),
//	httpServer:  server.NewInstance("0.0.0.0:80", config.Get("web.root").(string)).SetTlsEnable(false),
//}

func NewInstance(cfg config.Cfg) *Instance {
	//:= control.NewController(context.Background(), time.Second*3)
	inst := &Instance{
		Cfg:          cfg,
		httpServer:   server.NewInstance("0.0.0.0:80", cfg.GetValue("web.root").(string)),
		ConfigWriter: cfg.GetWriter().(*config.ConfigWriter),
		controlles:   make(map[*server.Server]controller.Controller),
	}
	tempCC := controller.NewController(context.Background(), time.Hour)
	inst.controlles[inst.httpServer] = tempCC
	inst.httpServer.SetController(tempCC)
	if cfg.GetValue("web.https.enable").(bool) { //如果配置文件中启用了,没启用就暂时不实例化
		inst.httpsServer = server.NewInstance("0.0.0.0:443", cfg.GetValue("web.root").(string)).
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
	commonFilter := (&filter.RequestFilter{}).SetInfo(s.GetStrings("redirect.keywords"))
	//启动非安全服务器
	if webconfig {
		s.httpServer.RegisterFunc("/config", ConfigHandler)
	}
	if s.GetValue("redirect.enable").(bool) {
		s.httpServer.RegisterFilterHandleFunc("/", http.FileServer(http.Dir(root)),
			commonFilter, filter.FilterHandler)
	} else {
		s.httpServer.RegisterHandle("/", http.FileServer(http.Dir(root)))
	}
	go s.httpServer.Start()
	//启动安全服务器
	if s.GetValue("web.https.enable").(bool) {
		if s.httpsServer == nil {
			s.httpsServer = server.NewInstance("0.0.0.0:443", root).
				SetTlsEnable(true, s.GetValue("web.https.filepath").(string))
		}
		if webconfig {
			s.httpsServer.RegisterFunc("/config", ConfigHandler)
		}
		if s.GetValue("redirect.enable").(bool) {
			s.httpsServer.RegisterFilterHandleFunc("/", http.FileServer(http.Dir(root)),
				commonFilter, filter.FilterHandler)
		} else {
			s.httpsServer.RegisterHandle("/", http.FileServer(http.Dir(root)))
		}
		go s.httpsServer.Start()
	}

	return s
}

func (s *Instance) StartDaemon() {
	s.controlles[s.httpServer].Wait(func() {
		fmt.Println("http server exit.")
	})
	s.controlles[s.httpsServer].Wait(func() {
		fmt.Println("https server exit.")
	})
	for {
		option := ' '
		fmt.Scan("%c", &option)
		if option == 'q' {
			return
		}
	}
}

func (s *Instance) Stop(ctx context.Context) {
	go s.httpServer.Stop(ctx)
	go s.httpsServer.Stop(ctx)

}

func (inst *Instance) Restart() {

}
