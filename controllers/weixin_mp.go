package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
	"weupay/models/weixinmp"
)

// operations for WeixinMp
type WeixinMpController struct {
	beego.Controller
}

func (c *WeixinMpController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// @router / [post]
func (c *WeixinMpController) Post() {
	var req weixinmp.Request
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, &req); err == nil {
		if res, err := weixinmp.Reply(&req); err == nil {
			c.Data["xml"] = res
		} else {
			c.Data["xml"] = err.Error()
		}
	} else {
		c.Data["xml"] = err.Error()
	}
	c.ServeXML()
}

// @router / [get]
func (c *WeixinMpController) GetAll() {
	signature := c.GetString("signature")
	beego.Info("signature:" + signature)
	timestamp := c.GetString("timestamp")
	beego.Info("timestamp:" + timestamp)
	nonce := c.GetString("nonce")
	beego.Info("nonce:" + nonce)
	echostr := c.GetString("echostr")
	beego.Info("echostr:" + echostr)

	localSign := weixinmp.Signature(timestamp, nonce)
	beego.Info("local sign:" + localSign)
	if localSign == signature {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("")
	}
}
