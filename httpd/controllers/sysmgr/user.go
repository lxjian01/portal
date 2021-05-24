package sysmgr

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/httpd/middlewares"
	"portal/httpd/models"
	"portal/httpd/services/sysmgr"
	"portal/httpd/utils"
	"strconv"
)

func AddUser(c *gin.Context){
	var resp utils.Response
	var m models.User
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := sysmgr.AddUser(&m)
	if err != nil {
		log.Errorf("Add system user error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateUser(c *gin.Context){
	var resp utils.Response
	var m models.User
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	err := sysmgr.UpdateUser(&m)
	if err != nil {
		log.Errorf("Update system user id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteUser(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	err = sysmgr.DeleteUser(id)
	if err != nil {
		log.Errorf("Delete system user id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetUserDetail(c *gin.Context){
	resp := &utils.Response{}
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	data, err := sysmgr.GetUserDetail(id)
	if err != nil {
		log.Errorf("Get system user id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetUserPage(c *gin.Context){
	resp := &utils.Response{}
	obj, isExist := c.GetQuery("pageIndex")
	if isExist != true {
		resp.ToMsgBadRequest(c, "参数pageIndex不能为空")
		return
	}
	pageIndex, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数pageIndex必须是整数")
		return
	}
	obj, isExist = c.GetQuery("pageSize")
	if isExist != true {
		resp.ToMsgBadRequest(c, "参数pageSize不能为空")
		return
	}
	pageSize, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数pageSize必须是整数")
		return
	}
	keywords, _ := c.GetQuery("keywords")
	data, err := sysmgr.GetUserPage(pageIndex, pageSize, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
