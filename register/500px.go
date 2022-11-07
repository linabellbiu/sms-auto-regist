package register

import "fmt"

/*
500px 网站
*/

type Px500 struct {
	Tel  string
	Code string
}

var PX500Channel = make(chan *Px500, 100)

func Register() {
	for {
		data := <-PX500Channel
		fmt.Println("手机号:", data.Tel)
		fmt.Println("注册码:", data.Code)
	}
}