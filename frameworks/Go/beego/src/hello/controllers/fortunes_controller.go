package controllers

import "hello/models"

type Fortunes struct {
	Base
}

func (c *Fortunes) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	c.Data["fortunes"] = models.GetFortunes()
	c.Render()
}
