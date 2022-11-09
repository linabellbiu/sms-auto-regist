package register

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

/*
500px 网站
*/

type Px500 struct {
	Tel  string
	Code string
}

var PX500Channel = make(chan *Px500, 1000)

func (p *Px500) Register() {
	for {
		select {
		case data := <-PX500Channel:
			fmt.Println("手机号:", data.Tel)
			fmt.Println("注册码:", data.Code)
			exist, err := p.isExist(data.Tel)
			if err != nil {
				fmt.Errorf("检查手机号注册失败:%v", err)
				return
			}
			if exist {
				fmt.Println("手机号已存在" + data.Tel)
				break
			}

			{

				//p.GetOrc()

			}
		}
	}
}

// 获取验证码
func (p *Px500) GetOrc(z string, tel string) {
	//resp, err := http.Get("http://httpbin.org/get")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
}

func (p *Px500) isExist(tel string) (bool, error) {
	client := resty.New()
	r := map[string]string{
		"countryCode": "86",
		"userName":    "13265520262",
	}
	req := client.R().SetFormData(r)
	resp, err := req.Post("https://500px.com.cn/user/v2/userIsExist")
	if err != nil {
		return false, err
	}
	var exist = &IsExist{}
	if err := json.Unmarshal(resp.Body(), exist); err != nil {
		return false, err
	}
	return exist.IsExist, nil
}

type IsExist struct {
	Message  string `json:"message"`
	UserName string `json:"userName"`
	IsExist  bool   `json:"isExist"`
	Status   string `json:"status"`
}
