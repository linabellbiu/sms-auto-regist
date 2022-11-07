package www_yunjiema_top

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
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
			c2.OnHTML("div[class='container']", func(element *colly.HTMLElement) {
				element.ForEach("div", func(i int, element *colly.HTMLElement) {
					fmt.Println(element.ChildText("div[class='col-xs-12 col-md-8']"))
				})
			})

			c2.OnRequest(func(r *colly.Request) {
				fmt.Println("爬取页面2：", "https://www.yunjiema.top"+r.URL.Path)
			})

			err := c2.Visit("https://www.yunjiema.top" + element.ChildAttr("a[href]", "href"))
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
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/2.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/3.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/4.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/5.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/6.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/7.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/8.html"); err != nil {
	//	return
	//}
	//if err := q.AddURL("https://www.yunjiema.top/zh/phone/9.html"); err != nil {
	//	return
	//}

	if err := q.Run(c); err != nil {
		return
	}
}
