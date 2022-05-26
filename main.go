package main

import (
	"github.com/shubhamjagdhane/simple-load-balancer/config"
	"github.com/shubhamjagdhane/simple-load-balancer/entity"
	"github.com/shubhamjagdhane/simple-load-balancer/logger"
	"github.com/shubhamjagdhane/simple-load-balancer/server"
)

func main() {
	paths := []string{"./config", "."}
	cfg, err := config.LoadConfig(paths, "config", entity.Config{})
	if err != nil {
		panic(err)
	}
	configuration := cfg.(entity.Config)
	logger := logger.New(configuration.Tracer.TracerName, configuration.ENV, configuration.LogLevel)
	sv := server.New(&configuration, logger)
	sv.Start()
}
