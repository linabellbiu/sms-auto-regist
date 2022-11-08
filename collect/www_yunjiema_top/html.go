package www_yunjiema_top

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/wangxudong123/sms-auto-regist/conf"
	"github.com/wangxudong123/sms-auto-regist/register"
	"regexp"
	"strings"
)

type collect struct {
	config conf.DefaultCollectConfig
}

func NewCollect(config conf.DefaultCollectConfig) *collect {
	return &collect{
		config: config,
	}
}

func (c *collect) Run() {
	//创建内存队列，大小10000，goroutine数量 5
	q, err := queue.New(5, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		return
	}

	co := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(10),
	)

	co.OnHTML("div[class='number-boxes']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='number-boxes-item d-flex flex-column ']", func(i int, element *colly.HTMLElement) {
			tel := element.ChildText("h4")
			// 尝试对这个手机号的接受短信爬取内容
			c2 := co.Clone()
			c2.OnHTML("div[class='container']", func(element *colly.HTMLElement) {
				element.ForEach("div", func(i int, element *colly.HTMLElement) {
					text := element.ChildText("div[class='col-xs-12 col-md-8']") // [视觉中国]Your Visual China Group verification code is:713169
					register.PX500Channel <- &register.Px500{
						Tel:  tel,
						Code: "",
					}
					var isExist bool
					for _, word := range c.config.Keywords {
						if strings.Index(text, word) != -1 {
							isExist = true
							break
						}
					}

					if isExist {
						// 提取验证码
						compileRegex := regexp.MustCompile(c.config.CompileRegex)
						matchArr := compileRegex.FindStringSubmatch(text)
						if len(matchArr) < 0 {
							fmt.Println("500px没有匹配到code")
							return
						}

						register.PX500Channel <- &register.Px500{
							Tel:  tel,
							Code: matchArr[len(matchArr)-1],
						}

						fmt.Println("=======================")
						fmt.Println("====获取到了============")
						fmt.Println("=======================")
					}
				})
			})

			c2.OnRequest(func(r *colly.Request) {
				fmt.Println("爬取页面2：", c.config.Host+r.URL.Path)
			})

			// 获取列表的手机号
			err := c2.Visit(fmt.Sprintf("%s.html", c.config.Host+"/zh/message/"+strings.ReplaceAll(tel, "+", "")))
			if err != nil {
				fmt.Println(err)
				return
			}
		})
	})

	co.OnRequest(func(r *colly.Request) {
		fmt.Println("爬取页面1：", r.URL)
	})

	if err := q.AddURL(c.config.Host + "/zh/phone/1.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/2.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/3.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/4.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/5.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/6.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/7.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/8.html"); err != nil {
		return
	}
	if err := q.AddURL(c.config.Host + "/zh/phone/9.html"); err != nil {
		return
	}

	if err := q.Run(co); err != nil {
		return
	}
}

func (c *collect) GetConfig() conf.DefaultCollectConfig {
	return c.config
}
