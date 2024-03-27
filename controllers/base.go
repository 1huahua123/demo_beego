package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) ParseInput(v interface{}) (err error) {
	if "application/json" == c.Ctx.Request.Header.Get("Content-Type") && len(c.Ctx.Input.RequestBody) > 0 {
		err = json.Unmarshal(c.Ctx.Input.RequestBody, v)
	} else {
		err = c.ParseForm(v)
	}
	return
}

func (c *BaseController) SetMessage(t string, v string) {
	temp, ok := c.Data["pageMessage"]
	if !ok {
		return
	}
	message, ok := temp.(map[string]string)
	if !ok {
		return
	}
	message[t] = v
	c.Data["pageMessage"] = message
}

func (c *BaseController) WithFlashError(msg string) {
	flash := beego.NewFlash()
	flash.Error(msg)
	flash.Store(&c.Controller)
}

func (c *BaseController) WithFlashSuccess(msg string) {
	flash := beego.NewFlash()
	flash.Success(msg)
	flash.Store(&c.Controller)
}

func (c *BaseController) WithFlashWarning(msg string) {
	flash := beego.NewFlash()
	flash.Warning(msg)
	flash.Store(&c.Controller)
}

func (c *BaseController) WithFlashNotice(msg string) {
	flash := beego.NewFlash()
	flash.Notice(msg)
	flash.Store(&c.Controller)
}
