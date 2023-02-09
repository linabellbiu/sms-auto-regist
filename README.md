# sms-auto-regist 
## 爬取虚拟手机号的短信验证码，从而登录需要手机号验证码注册的网站，可以刷阅读，刷赞，刷关注之类的操作
<h>

### How to use

#### config.yml
```yml
collect_source_html:
  www_yunjiema_top:
    # host
    host: "https://www.yunjiema.top"
    # timed task
    cron: "2 * * * *"
    # Locate the keywords of the general text message content, choose one of multiple keywords
    keywords:
      - "500px"
      - "[视觉中国]"
    # Regular expression to extract captcha
    compile_regex: "\\d{6,}"	
```	
	
	
	
#### reference 

run

```go
    // sms-auto-regist\example\main.go
	
    // initialization 
	collect.NewCollect(
		collect.SetConfigPath("../config.yml"),
	)

	// start app
	app.Run(&app.Example{})

	// Start the crawler scheduled task
	job(
		www_yunjiema_top.NewCollect(conf.Global.CollectSourceHtml.WwwYunjiemaTop),
	)
```

receive
```go
	for {
		select {
		case tel := <-collect.SendFindTel:
			fmt.Printf("Find the crawled mobile phone number %s:\n", tel)
		case tel := <-collect.SendFindSMSTel:
			fmt.Printf("Find the phone number that sent the text message%s:\n", tel)
		}
	}
```


### Customize crawling SMS website

#### mplement the interface
```go
// sms-auto-regist\collect\collect.go
type Job interface
```

#### reference
```go
// sms-auto-regist\collect\origin\www_yunjiema_top\html.go
type Collect struct {
	config conf.DefaultCollectConfig
}

func NewCollect(config conf.DefaultCollectConfig) *Collect {
	return &Collect{
		config: config,
	}
}

func (c *Collect) Run() {
	...
}

func (c *Collect) GetConfig() conf.DefaultCollectConfig {
	return c.config
}
```

#### tools
```go
// Send the crawled mobile phone number to the pipeline
collect.WriteFindTel(tel)

// Send the SMS received by crawling the mobile phone number to the pipeline
collect.WriteFindSMSTel(*collect.FindSMSTel)
```


### Result
#### 接收到的虚拟短信  

<image src="https://user-images.githubusercontent.com/20228139/200343258-7c6696c1-79c6-4b31-99dd-5f85b8e8bb91.png" width=60% height=40%>

#### 爬到的 
  
<image src="https://user-images.githubusercontent.com/20228139/200343508-05558328-6469-4d47-9894-7a50bdff2afb.png">

#### 手动输入验证码注册成功 (自动的还没搞)

<image src="https://user-images.githubusercontent.com/20228139/200343210-e8481c0e-551e-4e89-a73d-5edd9f34ab8b.png">
