package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/wangxudong123/sms-auto-regist/collect"
	"github.com/wangxudong123/sms-auto-regist/collect/www_yunjiema_top"
	"github.com/wangxudong123/sms-auto-regist/conf"
	"github.com/wangxudong123/sms-auto-regist/register"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var GlobalConfig = &conf.GlobalConfig{}
var Signal = make(chan int, 0)

func main() {
	parseConfig()

	register.Run(&register.Px500{})

	// 定时任务
	job(
		www_yunjiema_top.NewCollect(GlobalConfig.CollectSourceHtml.WwwYunjiemaTop),
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
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, GlobalConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(GlobalConfig)
}
