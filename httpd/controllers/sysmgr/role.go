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

func AddRole(c *gin.Context){
	var resp utils.Response
	var m models.Role
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	var myTime models.MyTime
	m.CreateTime = myTime.Now()
	m.UpdateTime = myTime.Now()
	_, err := sysmgr.AddRole(&m)
	if err != nil {
		log.Errorf("Add system role error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateRole(c *gin.Context){
	var resp utils.Response
	var m models.Role
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	var myTime models.MyTime
	m.UpdateTime = myTime.Now()
	err := sysmgr.UpdateRole(&m)
	if err != nil {
		log.Errorf("Update system role id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteRole(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	_, err = sysmgr.DeleteRole(id)
	if err != nil {
		log.Errorf("Delete system role id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetRoleDetail(c *gin.Context){
	resp := &utils.Response{}
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	data, err := sysmgr.GetRoleDetail(id)
	if err != nil {
		log.Errorf("Get system role id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetRoleList(c *gin.Context){
	resp := &utils.Response{}
	data, err := sysmgr.GetRoleList()
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetRolePage(c *gin.Context){
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
	title, _ := c.GetQuery("title")
	data, err := sysmgr.GetRolePage(pageIndex, pageSize, title)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
