package controllers

import (
	"github.com/Coop25/the-meme-index-api/config"
	"github.com/Coop25/the-meme-index-api/managers"
)

type Controller struct {
	config   *config.Config
	managers managers.Managers
}

func New(config *config.Config, managers managers.Managers) *Controller {
	return &Controller{
		config:   config,
		managers: managers,
	}
}
