package controllers

import "github.com/Coop25/the-meme-index-api/config"

type Controller struct {
	config *config.Config
}

func New(config *config.Config) *Controller {
	return &Controller{
		config: config,
	}
}
