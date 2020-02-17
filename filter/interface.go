/*******************************************************
 *  File        :   filter.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/17 12:15 上午
 *  Notes       :
 *******************************************************/
package filter

import (
	"net/http"
)

type Filter interface {
	WhetherFilterBySource(source interface{}) bool
	//WhetherFilterByInfo(info interface{}) bool
	GetInfo() interface{} //规定必须要有信息存储
	SetInfo(interface{}) Filter
}

// 过滤规则
type FilterHandleFunc func(h http.Handler, filter Filter) http.HandlerFunc
