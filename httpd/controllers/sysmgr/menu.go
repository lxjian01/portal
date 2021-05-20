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
	var myTime models.MyTime
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
	u.UpdateUser = middlewares.GetLoginUser().UserCode
	var myTime models.MyTime
	u.UpdateTime = myTime.Now()
	err := sysmgr.UpdateMenu(&u)
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
	title, _ := c.GetQuery("title")
	data, err := sysmgr.GetMenuPage(pageIndex, pageSize, title)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
