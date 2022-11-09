package register

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/wangxudong123/sms-auto-regist/data"
	"github.com/wangxudong123/sms-auto-regist/orc"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strings"
	"time"
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
		case d := <-PX500Channel:
			fmt.Println("手机号:", d.Tel)
			fmt.Println("注册码:", d.Code)
			exist, err := p.isExist(d.Tel)
			if err != nil {
				fmt.Errorf("检查手机号注册失败:%v", err)
				return
			}
			if exist {
				fmt.Println("手机号已存在" + d.Tel)
				break
			}

			{
				for _, code := range data.CountryCodeData {
					tel := strings.ReplaceAll(d.Tel, "+", "")
					if !strings.HasPrefix(tel, code.Code) {
						continue
					}
					tel = tel[len(code.Code):]

					// 获取验证码
					p.GetOrc(code.Code, tel)
				}
			}
		}
	}
}

// 获取验证码
func (p *Px500) GetOrc(z string, tel string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://500px.com.cn/user/v2/imgcode?dc=%d", time.Now().UnixMilli()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	decode, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}
	f, err := os.Create("E:\\CODE_WORK\\my\\sms-auto-regist\\register\\image.jpg")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = jpeg.Encode(b, decode, nil)
	if err != nil {
		return "", err
	}
	b.Flush()
	return orc.ImageCaptcha(""), nil
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
