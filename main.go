package main

import (
	"github.com/han-songyang/information-management-serve/infra"
	"github.com/han-songyang/information-management-serve/infra/logger"
)

func main() {
	infra.Init()
	logger.WarnLog("main", "main", "main")
}
