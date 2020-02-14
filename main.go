/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:09 下午
 *  Notes       :
 *******************************************************/
package main

import (
	_ "FastStaticWeb/config"
	"net/http"
)

//func config(w http.ResponseWriter,r *http.Request) {
//	fmt.Println(r.Method)
//	switch r.Method {
//	case "POST":
//		enable := r.FormValue("enable")
//		fmt.Println(enable)
//		file , fileHead ,err := r.FormFile("file")
//		if err != nil {
//			fmt.Println("read file failed")
//			return
//		}
//		fmt.Println("fileInfo:",fileHead.Filename,fileHead.Size)
//		var data []byte=make([]byte,fileHead.Size)
//		_,err = file.Read(data)
//		ioutil.WriteFile(fileHead.Filename,data,os.ModePerm)
//	default:
//		file,_ := template.ParseFiles("config.html")
//		file.Execute(w,nil)
//	}
//}

func main() {
	//config.config{}.SetWebconfigCfg().SetWebCfg(config.web)
	//http.HandleFunc("/config",config)
	http.ListenAndServe("0.0.0.0:80",nil)
}
