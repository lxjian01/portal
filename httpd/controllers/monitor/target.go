package monitor

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/httpd/middlewares"
	"portal/httpd/models"
	"portal/httpd/services/monitor"
	"portal/httpd/utils"
	"strconv"
)

func AddMonitorTarget(c *gin.Context){
	var resp utils.Response
	var m models.MonitorTarget
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := monitor.AddMonitorTarget(&m)
	if err != nil {
		log.Errorf("Add monitor target error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func DeleteMonitorTarget(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	err = monitor.DeleteMonitorTarget(id)
	if err != nil {
		log.Errorf("Delete monitor target id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetMonitorTargetPage(c *gin.Context){
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
	monitorClusterId := c.GetInt("monitorClusterId")
	monitorComponentId := c.GetInt("monitorComponentId")
	keywords := c.GetString("keywords")
	data, err := monitor.GetMonitorTargetPage(pageIndex, pageSize, monitorClusterId, monitorComponentId, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
