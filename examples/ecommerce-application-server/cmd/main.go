package main

import (
	"github.com/EraldCaka/ecommerce-application-server/router"
	"github.com/EraldCaka/ecommerce-application-server/util"
)

func main() {
	util.InitEnvironmentVariables()
	router.InitRouter()
	router.Start("0.0.0.0:5000")
}
