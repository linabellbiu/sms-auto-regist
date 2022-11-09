package register

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/wangxudong123/sms-auto-regist/conf"
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

type px500Client struct {
	cookies []*http.Cookie
}

var PX500Channel = make(chan *Px500, 1000)

func (p *Px500) Register() {
	for {
		select {
		case d := <-PX500Channel:
			fmt.Println("手机号:", d.Tel)
			fmt.Println("注册码:", d.Code)

			client := &px500Client{}

			var (
				z, tel string
			)

			// 检查手机的区号
			tel = strings.ReplaceAll(d.Tel, "+", "")
			for _, code := range data.CountryCodeData {
				if !strings.HasPrefix(tel, code.Code) {
					continue
				}
				tel = tel[len(code.Code):]
				z = code.Code
			}

			exist, err := client.isExist(z, d.Tel)
			if err != nil {
				fmt.Errorf("检查手机号注册失败:%v", err)
				return
			}
			if exist {
				fmt.Println("手机号已存在" + d.Tel)
				break
			}

			// 获取验证码
			{
				var code string
				// 错误的话尝试多次
				for i := 0; i < 10; i++ {
					code, err = client.GetOrc(z, tel)
					if err != nil {
						fmt.Errorf("获取验证码失败:%v", err)
					}
					if code == "" {
						time.Sleep(1 * time.Second)
						continue
					}
					break
				}
				// 请求注册
			}
		}
	}
}

func (p *px500Client) reg(code, z, tel string) error {
	client := resty.New()
	r := map[string]string{
		"countryCode": "86",
		"userName":    "13265520202",
		"displayName": "sssss",
		"password":    "aWha234234",
		"code":        "ssss",
	}
	req := client.R().SetFormData(r)

	for _, v := range p.cookies {
		fmt.Println(v.Name)
		fmt.Println(v.Value)
	}
	req.SetCookies(p.cookies)
	resp, err := req.Post("https://500px.com.cn/user/v2/toRegisterMe")
	if err != nil {
		return err
	}
	fmt.Println(resp.String())
	return nil
}

// 获取验证码
func (p *px500Client) GetOrc(z string, tel string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://500px.com.cn/user/v2/imgcode?dc=%d", time.Now().UnixMilli()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// cookie
	p.cookies = resp.Cookies()
	fmt.Println("获取验证码的cookie", p.cookies)
	decode, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}

	path := conf.Global.Orc.Image + fmt.Sprintf("/%s%s.jpg", z, tel)
	//path := "./orc/image/" + z + tel + ".jpg"
	f, err := os.Create(path)
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
	return orc.ImageCaptcha(path)
}

func (p *px500Client) isExist(z string, tel string) (bool, error) {
	client := resty.New()
	r := map[string]string{
		"countryCode": z,
		"userName":    tel,
	}
	resp, err := client.R().SetFormData(r).Post("https://500px.com.cn/user/v2/userIsExist")
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

func (p *px500Client) sendPhoneCode(code, z, tel string) {
	r := map[string]string{
		"imgCode":     code,
		"countryCode": z,
		"phone":       tel,
	}
	var cookie []*http.Cookie
	for _, v := range p.cookies {
		if v.Name == "SESSION" {
			cookie = append(cookie, &http.Cookie{
				Name:  v.Name,
				Value: v.Value,
			})
			break
		}
	}
	fmt.Println("发送验证码的cookie", cookie)
	resp, err := resty.New().R().SetFormData(r).SetCookies(cookie).Post("https://500px.com.cn/user/v2/sendPhoneCode")
	if err != nil {
		fmt.Errorf("发送短信验证码失败:%v", err)
		return
	}

	fmt.Println(resp)
}
