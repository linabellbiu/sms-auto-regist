package main

import (
	"github.com/wangxudong123/sms-auto-regist/register"
	"github.com/wangxudong123/sms-auto-regist/www_yunjiema_top"
)

func main() {
	// 需要丢到任务队列里面
	www_yunjiema_top.On()

	register.Run(&register.Px500{})
}
