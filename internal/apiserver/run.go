package apiserver

import (
	"github.com/RaulZuo/deep/internal/apiserver/config"
)

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}