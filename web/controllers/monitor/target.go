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

func AddMonitorTarget(c *gin.Context){
	var resp utils.Response
	var m models.MonitorTargetAdd
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

func UpdateMonitorTarget(c *gin.Context){
	var resp utils.Response
	var m models.MonitorTargetAdd
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := monitor.UpdateMonitorTarget(&m)
	if err != nil {
		log.Errorf("Update monitor target id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
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
	obj = c.DefaultQuery("prometheusId", "0")
	prometheusId, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数prometheusId必须是整数")
		return
	}
	obj = c.DefaultQuery("monitorResourceId", "0")
	monitorResourceId, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数monitorResourceId必须是整数")
		return
	}
	keywords := c.Query("keywords")
	data, err := monitor.GetMonitorTargetPage(pageIndex, pageSize, prometheusId, monitorResourceId, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
