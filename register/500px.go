package register

import (
	"fmt"
)

/*
500px 网站
*/

type Px500 struct {
	Tel  string
	Code string
}

var PX500Channel = make(chan *Px500, 100)

func (p *Px500) Register() {
	data := <-PX500Channel
	fmt.Println("手机号:", data.Tel)
	fmt.Println("注册码:", data.Code)

	// 调用验证码api

	// 调用注册接口

}

// 获取验证码
func (p *Px500) GetOrc() {
	//resp, err := http.Get("http://httpbin.org/get")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
}
