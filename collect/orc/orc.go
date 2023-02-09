package orc

import (
	"fmt"
	"github.com/linabellbiu/sms-auto-regist/conf"
	"os/exec"
	"strings"
)

func ImageCaptcha(path string) (string, error) {
	cmd := exec.Command(conf.Global.Orc.Cmd, conf.Global.Orc.Source+"/image_captcha.py", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("执行python失败:%v", err)
		return "", err
	}
	fmt.Println("执行命令", strings.TrimSpace(string(output)))
	return strings.TrimSpace(string(output)), nil
}
