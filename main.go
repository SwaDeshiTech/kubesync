package main

import (
	"fmt"

	"github.com/SwaDeshiTech/kubesync/api"
	"github.com/SwaDeshiTech/kubesync/config"
	"github.com/SwaDeshiTech/kubesync/databases"
)

func main() {

	//load config.yml
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	if err := databases.InitializeMongoConnection(); err != nil {
		panic(err)
	}

	router := api.ServerV1()
	if err := router.Run(fmt.Sprintf(":%d", config.GetConfig().Port)); err != nil {
		panic(err)
	}
}
