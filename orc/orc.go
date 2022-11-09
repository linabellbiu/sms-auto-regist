package orc

import (
	"log"
	"os/exec"
)

func ImageCaptcha(path string) string {
	cmd := exec.Command("python3", "./image_captcha.py", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	return string(output)
}
