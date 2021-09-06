package main

import (
	"fmt"
	"webqq/config"
	"github.com/spf13/viper"
)

func main(){
  
	//读取配置文件
	var conf = config.Config{
		"config/conf.yaml",
	}
	if err := conf.InitConfig(); err != nil{
		fmt.Println("读取配置文件出错: ", err)
		return
	}

	

}