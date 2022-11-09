package orc

import (
	"fmt"
	"os/exec"
)

func ImageCaptcha(path string) (string, error) {
	cmd := exec.Command("python3", "/Users/xudong/code/my/githup/sms-auto-regist/orc/image_captcha.py", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("执行python失败:%v", err)
		return "", err
	}
	return string(output), nil
}
