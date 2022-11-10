package register

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
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

					// 发送验证码
					if err := client.sendPhoneCode(code, z, tel); err != nil {
						continue
					}
					break
				}
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
	client := resty.New()
	resp, err := client.R().SetCookies(p.cookies).Get(fmt.Sprintf("https://500px.com.cn/user/v2/imgcode?dc=%d", time.Now().UnixMilli()))
	if err != nil {
		return "", err
	}

	decode, _, err := image.Decode(bytes.NewReader(resp.Body()))
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
	p.cookies = resp.Cookies()
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

type SendCodeResp struct {
	Flag    bool   `json:"flag"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (p *px500Client) sendPhoneCode(code, z, tel string) error {
	r := map[string]string{
		"imgCode":     code,
		"countryCode": z,
		"phone":       tel,
	}
	resp, err := resty.New().R().SetFormData(r).SetCookies(p.cookies).Post("https://500px.com.cn/user/v2/sendPhoneCode")
	if err != nil {
		fmt.Errorf("发送短信验证码失败:%v", err)
		return err
	}
	var res = &SendCodeResp{}
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		return err
	}

	if res.Status != "200" {
		fmt.Errorf("发送短信验证码失败:%s", resp.String())
		return errors.New("发送失败")
	}
	return nil
}
