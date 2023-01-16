package app

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestGetOrc(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("https://500px.com.cn/user/v2/imgcode?dc=%d", time.Now().UnixMilli()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	decode, _, err := image.Decode(resp.Body)
	if err != nil {
		return
	}
	f, err := os.Create("E:\\CODE_WORK\\my\\sms-auto-regist\\register\\image.jpg")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = jpeg.Encode(b, decode, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	b.Flush()
}

func TestPx500_Get(t *testing.T) {
	client := resty.New()

	r := map[string]string{
		"imgCode":     "6",
		"countryCode": "86",
		"phone":       "13265520262",
	}

	req := client.R().SetFormData(r)
	req.Cookies = append(req.Cookies,
		&http.Cookie{
			Name:     "SESSION",
			Value:    "a6426d6b-a36c-46da-8e3b-097f0db7913b",
			HttpOnly: true,
		},
	)

	resp, err := req.SetHeaders(map[string]string{
		"Accept":              "application/json;charset=utf-8",
		"Content-Type":        "application/json;charset=utf-8",
		"Referer":             "https://500px.com.cn/user/registerMe?redirect=https://500px.com.cn/community/index.html",
		"User-Agent":          "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
		"x-500px-client-info": "eyJhbm9ueW1vdXNJZCI6IjE4NDUyOTU5YzRkYWRkLTBiYzQ2NDFmMTNlNzItMzAyZjdmNDUtMjA3MzYwMC0xODQ1Mjk1OWM0ZTE0ZTMiLCJsYXRpdHVkZSI6bnVsbCwibG9uZ2l0dWRlIjpudWxsLCJtYW51ZmFjdHVyZXIiOiIiLCJwcm92aW5jZSI6IiIsImFyZWEiOiIiLCJjaXR5IjoiIn0=",
	}).Post("https://500px.com.cn/user/v2/sendPhoneCode")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.String())
}

func TestPx500_Exist(t *testing.T) {
	client := resty.New()

	r := map[string]string{
		"countryCode": "86",
		"userName":    "13265520262",
	}
	req := client.R().SetFormData(r)
	resp, err := req.Post("https://500px.com.cn/user/v2/userIsExist")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.String())
}

func TestPx500_SendCode(t *testing.T) {
	client := &px500Client{}
	code, err := client.GetOrc("86", "xx")
	fmt.Println(err)
	fmt.Println(fmt.Sprintf("---%s---", strings.TrimSpace(code)))
	client.sendPhoneCode(strings.ToLower(strings.TrimSpace(code)), "86", "xx")
}

func TestPx500_Register(t *testing.T) {
	client := resty.New()
	r := map[string]string{
		"countryCode": "86",
		"userName":    "xxx",
	}
	req := client.R().SetFormData(r)
	resp, err := req.Post("https://500px.com.cn/user/v2/userIsExist")
	if err != nil {
		t.Fatal(err)
	}
	cookie := resp.Cookies()

	app := &px500Client{
		cookie,
	}

	resp, err = client.R().SetCookies(cookie).Get(fmt.Sprintf("https://500px.com.cn/user/v2/imgcode?dc=%d", time.Now().UnixMilli()))
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
	decode, _, err := image.Decode(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	path := fmt.Sprintf("E:\\CODE_WORK\\my\\sms-auto-regist\\data\\image\\86.jpg")
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	b := bufio.NewWriter(f)
	err = jpeg.Encode(b, decode, nil)
	if err != nil {
		return
	}
	b.Flush()

	cmd := exec.Command("python", "E:\\CODE_WORK\\my\\sms-auto-regist\\orc\\image_captcha.py", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("执行python失败:%v", err)
		return
	}
	fmt.Println(strings.TrimSpace(string(output)))
	app.sendPhoneCode(strings.ToLower(strings.TrimSpace(string(output))), "86", "xxx")
}
