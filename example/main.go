package main

import (
	"github.com/linabellbiu/sms-auto-regist/collect"
	"github.com/linabellbiu/sms-auto-regist/collect/origin/www_yunjiema_top"
	"github.com/linabellbiu/sms-auto-regist/conf"
	app2 "github.com/linabellbiu/sms-auto-regist/example/app"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Signal = make(chan int, 0)

func main() {
	collect.NewCollect(
		collect.SetConfigPath("../config.yml"),
	)

	// 启动应用
	app2.Run(&app2.Example{})

	// 启动爬虫定时任务
	job(
		www_yunjiema_top.NewCollect(conf.Global.CollectSourceHtml.WwwYunjiemaTop),
	)
}

func job(jobs ...collect.Job) {
	c := cron.New()
	for _, j := range jobs {
		if err := c.AddJob(j.GetConfig().Cron, j); err != nil {
			log.Fatalln(err)
			return
		}
	}
	c.Start()
	defer c.Stop()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-s
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			close(Signal)
			log.Println("server quit !!!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
