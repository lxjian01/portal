package alarm

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/web/middlewares"
	"portal/web/models"
	"portal/web/services/alarm"
	"portal/web/utils"
	"strconv"
)

func AddAlertingMetric(c *gin.Context){
	var resp utils.Response
	var m models.AlertingMetric
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	service := alarm.AlertingMetricService{}
	_, err := service.AddAlertingMetric(&m)
	if err != nil {
		log.Errorf("Add alarm metric error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateAlertingMetric(c *gin.Context){
	var resp utils.Response
	var m models.AlertingMetric
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	service := alarm.AlertingMetricService{}
	err := service.UpdateAlertingMetric(&m)
	if err != nil {
		log.Errorf("Update alarm metric id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteAlertingMetric(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	service := alarm.AlertingMetricService{}
	_, err = service.DeleteAlertingMetric(id)
	if err != nil {
		log.Errorf("Delete alarm metric id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetAlertingMetricDetail(c *gin.Context){
	resp := &utils.Response{}
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	service := alarm.AlertingMetricService{}
	data, err := service.GetAlertingMetricDetail(id)
	if err != nil {
		log.Errorf("Get alerting metric detail id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetAlertingMetricList(c *gin.Context){
	resp := &utils.Response{}
	exporter := c.Query("exporter")
	service := alarm.AlertingMetricService{}
	data, err := service.GetAlertingMetricList(exporter)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetAlertingMetricPage(c *gin.Context){
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
	service := alarm.AlertingMetricService{}
	data, err := service.GetAlertingMetricPage(pageIndex, pageSize, exporter, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
