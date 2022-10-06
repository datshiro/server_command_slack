package main

import (
	"server/internal/infras/server"
	"server/internal/infras/server/config"
)

func main() {
	cfg := config.NewConfig()
	server := server.NewServer(cfg)
	server.RegisterHandlers()
	server.ConfigMiddlewares()
	server.Start()
}
