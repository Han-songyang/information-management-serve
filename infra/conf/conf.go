package conf

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var Config = &ini.File{}

func InitConf() {
	var err error
	path := "./conf/config.ini"
	Config, err = ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println("conf init success")
}
