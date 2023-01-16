package collect

import (
	"errors"
	"github.com/linabellbiu/sms-auto-regist/collect/source"
	"github.com/linabellbiu/sms-auto-regist/conf"
	"github.com/robfig/cron"
)

type FindSMSTel struct {
	Tel  string
	Code string
}

type Job interface {
	cron.Job
	GetConfig() conf.DefaultCollectConfig
}

// 设置项
type (
	option struct {
		configPath string
	}

	Setting func(*option)
)

// SetConfigPath 设置配置文件路径
func SetConfigPath(path string) Setting {
	return func(o *option) {
		o.configPath = path
	}
}

func NewCollect(s ...Setting) {
	o := &option{}
	for _, f := range s {
		f(o)
	}

	// 解析配置
	conf.ParseConfigPath(o.configPath)

	//加载数据
	source.ParseCountryCode()
}

var (
	// SendFindTel 发送未接收到短信的手机号
	SendFindTel = make(chan string, 10000)
	// SendFindSMSTel 发送接收到短信的手机号
	SendFindSMSTel = make(chan *FindSMSTel, 1000)
)

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
