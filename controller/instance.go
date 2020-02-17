/*******************************************************
 *  File        :   instance.go
 *  Author      :   nieaowei
 *  Date        :   2020/2/17 5:33 上午
 *  Notes       :
 *******************************************************/
package controller

import (
	"context"
	"errors"
	"time"
)

type control struct {
	signal context.Context
	context.CancelFunc
	wait func(func()) error
}

func (c *control) Notify() {
	c.CancelFunc()
}

func (c *control) Wait(f func()) error {
	return c.wait(f)
}

func NewController(ctx context.Context, timeout time.Duration) Controller {
	_, cancel := context.WithCancel(ctx)
	return &control{
		signal:     ctx,
		CancelFunc: cancel,
		wait: func(f func()) error {
			timer := time.NewTicker(timeout)
			for true {
				select {
				case <-ctx.Done():
					f()
					return nil
				case <-timer.C:
					return errors.New("timeout")
				}
			}
			return nil
		},
	}
}
