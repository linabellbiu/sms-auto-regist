package orc

import (
	"log"
	"os/exec"
)

func ImageCaptcha() string {
	cmd := exec.Command("python3", "./image_captcha.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	return string(output)
}
