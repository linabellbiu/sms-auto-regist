package register

import (
	"bufio"
	"fmt"
	"github.com/go-resty/resty/v2"
	"image"
	"image/jpeg"
	"net/http"
	"os"
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

	r := map[string]interface{}{
		"imgCode":     "vbv1",
		"countryCode": "86",
		"phone":       "13265520262",
	}

	req := client.R().SetBody(r)
	req.Cookies = append(req.Cookies, &http.Cookie{
		Name:     "_uab_collina",
		Value:    "166783346481023169313068",
		HttpOnly: true,
	},
		&http.Cookie{
			Name:     "SESSION",
			Value:    "5afe8d9d-6d42-4f28-bf8b-5aad07b90c16",
			HttpOnly: true,
		},
		&http.Cookie{
			Name:     "acw_tc",
			Value:    "2760828716679131587487463e66f65a2dbc78fac3b3d63144ac59bf7a2e3b",
			HttpOnly: true},
	)

	resp, err := req.SetHeaders(map[string]string{
		"Accept":       "application/json;charset=utf-8",
		"Content-Type": "application/json;charset=utf-8",
		"Referer":      "https://500px.com.cn/user/registerMe?redirect=https://500px.com.cn/community/index.html",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
		//"Cookie":              "_uab_collina=166783346481023169313068; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2218452959c4dadd-0bc4641f13e72-302f7f45-2073600-18452959c4e14e3%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E8%87%AA%E7%84%B6%E6%90%9C%E7%B4%A2%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Fwww.google.com%2F%22%7D%2C%22%24device_id%22%3A%2218452959c4dadd-0bc4641f13e72-302f7f45-2073600-18452959c4e14e3%22%7D; testapp=s%3AhFQKCoeeU_XaljzjGM3jUVgSvo5UQaiz.cHE7DDyFgkUMjagsbvQ04miXpGSrTJGxhUvKUVhc%2FFg; SESSION=5afe8d9d-6d42-4f28-bf8b-5aad07b90c16; Hm_lvt_3eea10d35cb3423b367886fc968de15a=1667832851,1667911371; acw_tc=2760828716679131587487463e66f65a2dbc78fac3b3d63144ac59bf7a2e3b; Hm_lpvt_3eea10d35cb3423b367886fc968de15a=1667913175",
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
