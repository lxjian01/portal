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

func AddMenu(c *gin.Context){
	var resp utils.Response
	var u models.Menu
	if err := c.ShouldBindJSON(&u);err != nil{
		resp.ToError(c, err)
		return
	}
	u.CreateUser = middlewares.GetLoginUser().UserCode
	u.UpdateUser = middlewares.GetLoginUser().UserCode
	var myTime utils.MyTime
	u.CreateTime = myTime.Now()
	u.UpdateTime = myTime.Now()
	_, err := sysmgr.AddMenu(&u)
	if err != nil {
		log.Errorf("Add system menu error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": u.Id}
	resp.ToSuccess(c)
}

func UpdateMenu(c *gin.Context){
	var resp utils.Response
	var u models.Menu
	if err := c.ShouldBindJSON(&u);err != nil{
		resp.ToError(c, err)
		return
	}
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	u.Id = id
	u.UpdateUser = middlewares.GetLoginUser().UserCode
	var myTime utils.MyTime
	u.UpdateTime = myTime.Now()
	err = sysmgr.UpdateMenu(&u)
	if err != nil {
		log.Errorf("Update system menu id=%d error %s", u.Id, err.Error())
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
		log.Errorf("Get system menu id=%d error %s", id, err.Error())
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

func GetMenuPage(c *gin.Context){
	resp := &utils.Response{}
	obj, isExist := c.GetQuery("page")
	if isExist != true {
		resp.ToMsgBadRequest(c, "参数page不能为空")
		return
	}
	pageIndex, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数page必须是整数")
		return
	}
	obj, isExist = c.GetQuery("page_size")
	if isExist != true {
		resp.ToMsgBadRequest(c, "参数page_size不能为空")
		return
	}
	pageSize, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数page_size必须是整数")
		return
	}
	keywords, _ := c.GetQuery("keywords")
	data, err := sysmgr.GetMenuPage(pageIndex, pageSize, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}
