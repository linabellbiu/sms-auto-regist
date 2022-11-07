package register

import (
	"fmt"
	"time"
)

/*
500px 网站
*/

type Px500 struct {
	Tel  string
	Code string
}

var PX500Channel = make(chan *Px500, 100)

func receive() {
	tick := time.NewTicker(5 * time.Second)
	for {
		data := <-PX500Channel
		fmt.Println("手机号:", data.Tel)
		fmt.Println("注册码:", data.Code)
		register(data)
		<-tick.C
	}
}

func register(px *Px500) {
	// 调用验证码api

	// 调用注册接口

}

func PX500() {
	go receive()
}
