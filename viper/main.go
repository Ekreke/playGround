package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// https://www.liwenzhou.com/posts/Go/viper/
func main() {

	viper.SetConfigFile("./config.json")
	// 读取config
	err := viper.ReadInConfig()
	host := viper.GetString("datastore.metric.host")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host)
	fmt.Println("success")
}
