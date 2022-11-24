package main

import (
	"https://github.com/Yuki-J-Klang/udemy-go-lesson/tree/master/ToDo_app_golang/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.Logfile)
}
