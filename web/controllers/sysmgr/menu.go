package sysmgr

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/web/middlewares"
	"portal/web/models"
	"portal/web/services/sysmgr"
	"portal/web/utils"
	"strconv"
)

func AddMenu(c *gin.Context){
	var resp utils.Response
	var m models.Menu
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := sysmgr.AddMenu(&m)
	if err != nil {
		log.Errorf("Add system menu error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateMenu(c *gin.Context){
	var resp utils.Response
	var m models.Menu
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	err := sysmgr.UpdateMenu(&m)
	if err != nil {
		log.Errorf("Update system menu id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteMenu(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	_, err = sysmgr.DeleteMenu(id)
	if err != nil {
		log.Errorf("Delete system menu id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetMenuDetail(c *gin.Context){
	resp := &utils.Response{}
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	data, err := sysmgr.GetMenuDetail(id)
	if err != nil {
		log.Errorf("Get system menu detail id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetMenuList(c *gin.Context){
	resp := &utils.Response{}
	data, err := sysmgr.GetMenuList()
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetParentMenuList(c *gin.Context){
	resp := &utils.Response{}
	data, err := sysmgr.GetParentMenuList()
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetMenuPage(c *gin.Context){
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
	title := c.Query("title")
	data, err := sysmgr.GetMenuPage(pageIndex, pageSize, title)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
