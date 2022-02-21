package main

import (
	"douban/cmd"
	"douban/tool"
	"log"
)

func main() {
	err := tool.InitMySQL()
	if err != nil {
		log.Println(err)
		return
	}
	err = tool.InitRedis()
	if err != nil {
		return
	}
	cmd.URL()
}
