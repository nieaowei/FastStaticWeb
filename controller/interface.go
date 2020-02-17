/*******************************************************
 *  File        :   control.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/16 5:47 上午
 *  Notes       :	控制
 *******************************************************/
package controller

type Controller interface {
	Wait(func()) error
	Notify()
}
