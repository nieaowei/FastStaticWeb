/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:09 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"FastStaticWeb/config"
	"FastStaticWeb/server"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {
	//flag.Parse()
	//if *cpuprofile != "" {
	//	f, err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	_ = pprof.StartCPUProfile(f)
	//	defer pprof.StopCPUProfile()
	//}
	inst := server.NewInstance(config.DefaultConfig())
	inst.Start().StartDaemon()
}
