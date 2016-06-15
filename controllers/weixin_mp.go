package controllers

import (
	"github.com/astaxie/beego"
	"weupay/models/weixin/mp"
)

// operations for WeixinMp
type WeixinMpController struct {
	beego.Controller
}

func (c *WeixinMpController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create WeixinMp
// @Param	body		body 	models.WeixinMp	true		"body for WeixinMp content"
// @Success 201 {object} models.WeixinMp
// @Failure 403 body is empty
// @router / [post]
func (c *WeixinMpController) Post() {

}

// @Title Get All
// @Description get WeixinMp
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.WeixinMp
// @Failure 403
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

	localSign := mp.Signature(timestamp, nonce)
	beego.Info("local sign:" + localSign)
	if localSign == signature {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("")
	}
}
