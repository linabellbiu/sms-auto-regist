package orc

import (
	"log"
	"os/exec"
	"testing"
)

func TestImageCaptcha(t *testing.T) {
	cmd := exec.Command("python3", "githup/sms-auto-regist/orc/image_captcha.py", "githup/sms-auto-regist/orc/image/imgcode.jpeg")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("result =>", string(output))
}
