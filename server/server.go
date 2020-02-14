/*******************************************************
 *  File        :   server.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 11:48 下午
 *  Notes       :
 *******************************************************/
package server

import (
	"context"
	"net/http"
)

type Instance struct {
	server http.Server
	root   string
	ctx    context.Context
}

func NewServer(addr string, root string, handler http.Handler) *Instance {
	return &Instance{
		server: http.Server{
			Addr:              "",
			Handler:           handler,
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

func (s *Instance) SetListenAddr(addr string) *Instance {
	s.server.Addr = addr
	return s
}

func (s *Instance) SetRoot(addr string) *Instance {
	s.root = addr
	return s
}

func (s *Instance) SetHandler(handler http.Handler) *Instance {
	s.server.Handler = handler
	return s
}

func (s *Instance) RegisterFunc(url string, handler func(http.ResponseWriter, *http.Request)) *Instance {
	s.server.Handler.(*http.ServeMux).HandleFunc(url, handler)
	return s
}

func (s *Instance) Start() {
	//s.server.Handler.(*http.ServeMux).Handle("/",http.FileServer(http.Dir(s.root)))
	go s.server.ListenAndServe()
}

func (s *Instance) Stop() {
	s.server.Shutdown(s.ctx)
}

func (s *Instance) SetContext(ctx context.Context) *Instance {
	s.ctx = ctx
	return s
}
