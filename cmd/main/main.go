package main

import (
	"fmt"

	"github.com/online_marketplace/helper/core/service"
	"github.com/online_marketplace/helper/rest"

	"github.com/online_marketplace/internal/config"
	"github.com/online_marketplace/internal/handler"
	"github.com/online_marketplace/internal/registry"
)

func main() {
	c := config.Init()
	appCtx := registry.NewServiceContext(*c)
	router := handler.NewRouter(appCtx)

	svcGroup := service.NewServiceGroup()
	srv := rest.MustNewServer(c.Server, router)

	svcGroup.Add(srv)
	defer svcGroup.Stop()

	fmt.Printf("Starting server at %s:%d...\n", c.Server.Http.Host, c.Server.Http.Port)
	svcGroup.Start()
}
