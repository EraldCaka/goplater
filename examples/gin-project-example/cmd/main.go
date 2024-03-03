package main

import (
	"github.com/EraldCaka/gin-project-example/router"
	"github.com/EraldCaka/gin-project-example/util"
)

func main() {
	util.InitEnvironmentVariables()
	router.InitRouter()
	router.Start("0.0.0.0:5000")
}
