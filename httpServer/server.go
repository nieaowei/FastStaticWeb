/*******************************************************
 *  File        :   server.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 11:48 下午
 *  Notes       :
 *******************************************************/
package httpServer

import (
	"FastStaticWeb/controller"
	"FastStaticWeb/filter"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Router struct {
	router http.Handler
}

type Server struct {
	server                http.Server //服务器
	tlsEnable             bool        //是否启用https
	tlsFilePath           string      //证书所在目录
	root                  string      //网站根目录
	count                 chan int    //记录重启次数
	controller.Controller             //通知-等待机制，用于服务器退出时通知上层
}

//NewRouter is to return a *ServerMux.
func NewRouter() *Router {
	return &Router{router: http.NewServeMux()}
}

//NewInstance is to return a Server instance.
func NewInstance(addr string, root string) *Server {
	return &Server{
		server: http.Server{
			Addr:              addr,
			Handler:           NewRouter().router,
			TLSConfig:         nil,
			ReadTimeout:       0,
			ReadHeaderTimeout: 0,
			WriteTimeout:      0,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          nil,
			BaseContext:       nil,
			ConnContext:       nil,
		},
		root: root,
	}
}

//SetRouter is to set up a Router or ServerMux for the server.
func (s *Server) SetRouter(router Router) *Server {
	//s.router = router.router
	s.server.Handler = router.router
	return s
}

//SetTilsEnable is to set up the enabled status of the TLS and set up
//the certificate files path.
func (s *Server) SetTlsEnable(ok bool, filepath string) *Server {
	s.tlsEnable = ok
	s.tlsFilePath = filepath
	return s
}

//SetListenAddr is to set up address for server.
func (s *Server) SetListenAddr(addr string) *Server {
	s.server.Addr = addr
	return s
}

//SetRoot is to set up path to a website path or other static file for the sever.
func (s *Server) SetRoot(addr string) *Server {
	s.root = addr
	return s
}

//SetHandler equivalent to SetRouter.
func (s *Server) SetHandler(handler http.Handler) *Server {
	s.server.Handler = handler
	return s
}

//
func (s *Server) RegisterFunc(url string, handler func(http.ResponseWriter, *http.Request)) *Server {
	s.server.Handler.(*http.ServeMux).HandleFunc(url, handler)
	return s
}

func (s *Server) RegisterHandle(url string, handler http.Handler) *Server {
	s.server.Handler.(*http.ServeMux).Handle(url, handler)
	return s
}

func (s *Server) RegisterFilterHandleFunc(url string, h http.Handler, filter filter.Filter, handleFunc filter.FilterHandleFunc) *Server {
	s.server.Handler.(*http.ServeMux).HandleFunc(url, handleFunc(h, filter))
	return s
}

func scanTlsFile(filePath string) (certPath, keyPath string) {
	files, err := os.Open(filePath)
	if err != nil {
		fmt.Println("readfilefalied")
		return
	}
	defer func() {
		defer func() {
			if ok := recover(); ok != nil {
				//todo 日志异常，程序退出
				return
			}
		}()
		files.Close()
	}()
	fileName, err := files.Readdir(0)
	keyPath = filepath.Join(filePath)
	certPath = filepath.Join(filePath)
	for _, fileInfo := range fileName {
		if strings.HasSuffix(fileInfo.Name(), ".pem") {
			certPath = filepath.Join(filePath, fileInfo.Name())
		} else if strings.HasSuffix(fileInfo.Name(), ".key") {
			keyPath = filepath.Join(filePath, fileInfo.Name())
		}
	}
	return certPath, keyPath
}

func (s *Server) Start() {
	if s.tlsEnable {
		err := s.server.ListenAndServeTLS(scanTlsFile(s.tlsFilePath))
		if err != nil { //异常或手动关闭
			//尝试重启
			fmt.Println(err)
			s.Notify()
		}

	} else {
		err := s.server.ListenAndServe()
		if err != nil {
			//todo 程序退出日志
			fmt.Println(err)

			s.Notify()
		}
	}
	return
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) ReStart(ctx context.Context) error {
	err := s.Stop(ctx)
	if err != nil {
		return err
	}
	s.Start()
	return nil
}

func (s *Server) SetController(controller controller.Controller) *Server {
	s.Controller = controller
	return s
}
