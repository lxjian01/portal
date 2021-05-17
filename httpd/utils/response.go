package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"portal/global/log"
)

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg interface{} `json:"msg"`
}

func (resp *Response) ToJson(c *gin.Context){
	if resp.Code == 2000 {
		if resp.Msg == nil {
			resp.Msg = "success"
		}
	}else{
		log.Error(resp.Msg)
	}
	c.JSON(200, resp)
}

func (resp *Response) ToSuccess(c *gin.Context){
	resp.Code = 2000
	resp.ToJson(c)
}

func (resp *Response) ToError(c *gin.Context, err error){
	resp.Code = 4000
	//获取validator.ValidationErrors类型的errors
	errs,ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator类型错误直接返回
		resp.Msg = err.Error()
	}else{
		// valodator.ValidationErrors类型错误进行翻译
		errMsg := RemoveStructName(errs.Translate(Trans))
		resp.Msg = errMsg
	}
	resp.ToJson(c)
}


// 400
func (resp *Response) ToMsgBadRequest(c *gin.Context, msg interface{}){
	resp.Code = 4000
	resp.Msg = msg
	resp.ToJson(c)
}

// 401
func (resp *Response) ToMsgUnauthorized(c *gin.Context, msg interface{}){
	resp.Code = 4010
	resp.Msg = msg
	c.JSON(http.StatusUnauthorized, resp)
}

// 403
func (resp *Response) ToMsgForbidden(c *gin.Context, msg interface{}){
	resp.Code = 4030
	resp.Msg = msg
	c.JSON(http.StatusForbidden, resp)
}
