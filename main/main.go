/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:09 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"FastStaticWeb/config"
	"FastStaticWeb/faststaticweb"
)

func main() {
	inst := faststaticweb.NewInstance(config.DefaultConfig())
	inst.Start().StartDaemon()
}
