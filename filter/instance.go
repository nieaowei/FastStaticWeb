/*******************************************************
 *  File        :   filter.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/17 1:35 上午
 *  Notes       :
 *******************************************************/
package filter

import (
	"FastStaticWeb/config"
	"fmt"
	"github.com/gogf/gf/encoding/gcharset"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RequestFilter struct {
	info []string
}

func (rf *RequestFilter) GetInfo() interface{} {
	return rf.info
}

func (rf *RequestFilter) SetInfo(info interface{}) Filter {
	rf.info = info.([]string)
	return rf
}

func (rf *RequestFilter) WhetherFilterBySource(source interface{}) bool {
	userIp := source.(string)
	fmt.Println(userIp)
	resStrs := strings.Split(userIp, ":")
	if len(resStrs) == 0 {

	}
	//userIp = "47.100.123.43"
	http.DefaultClient.Timeout = time.Second * 3
	res, err := http.Get("http://whois.pconline.com.cn/ip.jsp?ip=" + resStrs[0])
	if err != nil {
		fmt.Println("filter timeout.")
		return false
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}
	tar, err := gcharset.ToUTF8("GBK", string(buf))
	for _, val := range rf.info {
		if strings.Contains(tar, val) {
			return true
		}
	}
	return false
}

func FilterHandler(h http.Handler, filter Filter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("filter handler ......")
		if filter.WhetherFilterBySource(request.RemoteAddr) {
			fmt.Println("host:", request.Host)
			fmt.Println("requestUri:", request.RequestURI)
			fmt.Println("referer:", request.Referer())
			fmt.Println("path:", request.URL.Path)
			http.Redirect(writer, request, config.DefaultConfig().GetValue("redirect.reurl").(string)+request.URL.Path, 302)
		} else {
			h.ServeHTTP(writer, request)
		}
	}
}
