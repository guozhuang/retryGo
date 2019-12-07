package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	MogodbUrl string `yaml:"mongodbUrl"`
	MgoDb     string `yaml:"mongodb"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

/*
获取配置信息的基本方法
*/
func GetConf() *conf {
	var c conf
	cfg := c.getConf()
	return cfg
}
