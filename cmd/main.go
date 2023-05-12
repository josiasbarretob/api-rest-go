package main

import (
	"fmt"

	"github.com/josiasbarretob/api-rest-go/config"
	"github.com/josiasbarretob/api-rest-go/router"
)
var(
	logger config.Logger
)
func main(){
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil{
		logger.Errorf("config inicialization error: %v", v... err)
		return
	}
	router.Initialize()
}