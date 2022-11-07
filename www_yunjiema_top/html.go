package www_yunjiema_top

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/wangxudong123/sms-auto-regist/register"
	"regexp"
	"strings"
)

func On() {
	//创建内存队列，大小10000，goroutine数量 5
	q, err := queue.New(5, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		return
	}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(10),
	)

	c.OnHTML("div[class='number-boxes']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='number-boxes-item d-flex flex-column ']", func(i int, element *colly.HTMLElement) {
			c2 := c.Clone()

			tel := element.ChildText("h4")

			// 尝试对这个手机号的接受短信爬取内容
			c2.OnHTML("div[class='container']", func(element *colly.HTMLElement) {
				element.ForEach("div", func(i int, element *colly.HTMLElement) {
					text := element.ChildText("div[class='col-xs-12 col-md-8']")
					if strings.Index(text, "500px") != -1 || strings.Index(text, "[视觉中国]") != -1 {
						compileRegex := regexp.MustCompile("\\d{6,}")
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
				fmt.Println("爬取页面2：", "https://www.yunjiema.top"+r.URL.Path)
			})

			// 获取列表的手机号
			err := c2.Visit(fmt.Sprintf("%s.html", "https://www.yunjiema.top/zh/message/"+strings.ReplaceAll(tel, "+", "")))
			if err != nil {
				fmt.Println(err)
				return
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("爬取页面1：", r.URL)
	})

	if err := q.AddURL("https://www.yunjiema.top/zh/phone/1.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/2.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/3.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/4.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/5.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/6.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/7.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/8.html"); err != nil {
		return
	}
	if err := q.AddURL("https://www.yunjiema.top/zh/phone/9.html"); err != nil {
		return
	}

	if err := q.Run(c); err != nil {
		return
	}
}
