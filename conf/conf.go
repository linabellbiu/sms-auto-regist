package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func ParseConfig() {
	yamlFile, err := os.ReadFile("../config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, Global)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println(Global)
}

// ParseConfigPath 指定配置文件解析
func ParseConfigPath(path string) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, Global)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println(Global)
}
