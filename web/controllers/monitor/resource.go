package monitor

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/web/middlewares"
	"portal/web/models"
	"portal/web/services/monitor"
	"portal/web/utils"
	"strconv"
)

func AddMonitorResource(c *gin.Context){
	var resp utils.Response
	var m models.MonitorResource
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := monitor.AddMonitorResource(&m)
	if err != nil {
		log.Errorf("Add monitor resource error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateMonitorResource(c *gin.Context){
	var resp utils.Response
	var m models.MonitorResource
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	err := monitor.UpdateMonitorResource(&m)
	if err != nil {
		log.Errorf("Update monitor resource id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteMonitorResource(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	_, err = monitor.DeleteMonitorResource(id)
	if err != nil {
		log.Errorf("Delete monitor resource id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetMonitorResourceList(c *gin.Context){
	resp := &utils.Response{}
	data, err := monitor.GetMonitorResourceList()
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetMonitorResourcePage(c *gin.Context){
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
	exporter := c.Query("exporter")
	keywords := c.Query("keywords")
	data, err := monitor.GetMonitorResourcePage(pageIndex, pageSize, exporter, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
