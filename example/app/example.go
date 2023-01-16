package app

import (
	"fmt"
	"github.com/linabellbiu/sms-auto-regist/collect"
)

type Example struct {
	Tel  string
	Code string
}

func (p *Example) Register() {
	for {
		select {
		case tel := <-collect.SendFindTel:
			fmt.Printf("未发送的手机号%s:\n", tel)
		case tel := <-collect.SendFindSMSTel:
			fmt.Printf("找到发送短信的手机号%s:\n", tel)
		}
	}
}
