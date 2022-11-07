package test

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"regexp"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("C:\\Users\\Administrator\\Downloads\\imgcode.jpg")
	client.SetLanguage("chi_sim")
	text, _ := client.Text()
	//fmt.Println(text)
	reg := regexp.MustCompile("验证码: [^ ]*\\s")
	yzmStr := reg.FindAllString(text, 1)
	fmt.Println(yzmStr)
	yzmSlice := strings.Split(text, "验证码: ")
	fmt.Println(yzmSlice[1])
}
