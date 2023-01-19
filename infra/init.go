package infra

import (
	"github.com/han-songyang/information-management-serve/infra/conf"
	"github.com/han-songyang/information-management-serve/infra/logger"
	"github.com/han-songyang/information-management-serve/infra/mysql"
	"github.com/han-songyang/information-management-serve/router"
)

func Init() {
	conf.InitConf()
	logger.InitLog()
	mysql.InitMysql()
	router.InitRouter()
}
