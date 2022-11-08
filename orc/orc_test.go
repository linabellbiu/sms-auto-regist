package orc

import (
	"log"
	"os/exec"
	"testing"
)

func TestImageCaptcha(t *testing.T) {
	cmd := exec.Command("python3", "/Users/xudong/code/my/githup/sms-auto-regist/orc/image_captcha.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("result =>", string(output))
}