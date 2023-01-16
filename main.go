package main

import (
	"fmt"
	"github.com/linabellbiu/sms-auto-regist/app"
	"github.com/linabellbiu/sms-auto-regist/collect"
	"github.com/linabellbiu/sms-auto-regist/collect/www_yunjiema_top"
	"github.com/linabellbiu/sms-auto-regist/conf"
	"github.com/linabellbiu/sms-auto-regist/data"
	"github.com/robfig/cron"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Signal = make(chan int, 0)

func main() {
	// 解析配置
	parseConfig()

	//加载数据
	data.ParseCountryCode()

	// 需要自动注册的api
	app.Run(&app.Example{})

	// 启动爬虫定时任务
	job(
		www_yunjiema_top.NewCollect(conf.Global.CollectSourceHtml.WwwYunjiemaTop),
	)
}

func job(jobs ...collect.CollerJob) {
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

func parseConfig() {
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, conf.Global)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf.Global)
}
