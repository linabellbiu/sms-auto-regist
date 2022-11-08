package collect

import (
	"github.com/robfig/cron"
	"github.com/wangxudong123/sms-auto-regist/conf"
)

type CollerJob interface {
	cron.Job
	GetConfig() conf.DefaultCollectConfig
}
