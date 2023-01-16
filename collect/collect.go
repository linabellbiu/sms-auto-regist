package collect

import (
	"errors"
	"github.com/linabellbiu/sms-auto-regist/conf"
	"github.com/robfig/cron"
)

type FindSMSTel struct {
	Tel  string
	Code string
}

var (
	// 发送未接收到短信的手机号
	SendFindTel = make(chan string, 10000)
	// 发送接收到短信的手机号
	SendFindSMSTel = make(chan *FindSMSTel, 1000)
)

type CollerJob interface {
	cron.Job
	GetConfig() conf.DefaultCollectConfig
}

func WriteFindTel(a string) error {
	select {
	case SendFindTel <- a:

	default:
		return errors.New("failed to send phone.chan:SendFindTel")
	}
	return nil
}

func WriteFindSMSTel(a *FindSMSTel) error {
	select {
	case SendFindSMSTel <- a:

	default:
		return errors.New("failed to send phone.chan:SendFindSMSTel")
	}
	return nil
}
