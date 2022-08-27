package service

import (
	"go-architecture/config"
	"go-architecture/usecase"
)

type HttpHandler struct {
	Usecase usecase.UsecaseInteractor
	Config *config.ShareConfig
}

func NewHttpHandler(usecase usecase.UsecaseInteractor, config config.ShareConfig) HttpHandler {
	return HttpHandler{
		Usecase: usecase,
		Config: &config,
	}
}
