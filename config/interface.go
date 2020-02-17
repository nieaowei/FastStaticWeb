/*******************************************************
 *  File        :   config.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/14 6:10 下午
 *  Notes       :
 *******************************************************/
package config

type Cfg interface {
	WriteConfig() error
	GetValue(key string) interface{}
	GetStrings(key string) []string
	GetWriter() interface{}
}
