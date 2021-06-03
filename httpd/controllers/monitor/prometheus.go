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

func AddMonitorPrometheus(c *gin.Context){
	var resp utils.Response
	var m models.MonitorPrometheus
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := monitor.AddMonitorPrometheus(&m)
	if err != nil {
		log.Errorf("Add monitor prometheus error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateMonitorPrometheus(c *gin.Context){
	var resp utils.Response
	var m models.MonitorPrometheus
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	err := monitor.UpdateMonitorPrometheus(&m)
	if err != nil {
		log.Errorf("Update monitor prometheus id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteMonitorPrometheus(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	_, err = monitor.DeleteMonitorPrometheus(id)
	if err != nil {
		log.Errorf("Delete monitor prometheus id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetMonitorPrometheusList(c *gin.Context){
	resp := &utils.Response{}
	data, err := monitor.GetMonitorPrometheusList()
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetMonitorPrometheusPage(c *gin.Context){
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
	monitorClusterId := 0
	obj, isExist = c.GetQuery("monitorClusterId")
	if isExist == true {
		monitorClusterId, err = strconv.Atoi(obj)
		if err != nil {
			resp.ToMsgBadRequest(c, "参数monitorClusterId必须是整数")
			return
		}
	}
	keywords := c.Query("keywords")
	data, err := monitor.GetMonitorPrometheusPage(pageIndex, pageSize, monitorClusterId, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
