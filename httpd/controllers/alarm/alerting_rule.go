package alarm

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"portal/httpd/middlewares"
	"portal/httpd/models"
	"portal/httpd/services/alarm"
	"portal/httpd/utils"
	"strconv"
)

func AddAlertingRule(c *gin.Context){
	var resp utils.Response
	var m models.AlertingRuleAdd
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.CreateUser = middlewares.GetLoginUser().UserCode
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	_, err := alarm.AddAlertingRule(&m)
	if err != nil {
		log.Errorf("Add alarm rule error %s",err.Error())
		resp.ToError(c, err)
		return
	}
	resp.Data = gin.H{"id": m.Id}
	resp.ToSuccess(c)
}

func UpdateAlertingRule(c *gin.Context){
	var resp utils.Response
	var m models.AlertingRuleAdd
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
	m.UpdateUser = middlewares.GetLoginUser().UserCode
	err := alarm.UpdateAlertingRule(&m)
	if err != nil {
		log.Errorf("Update alerting rule id=%d error %s", m.Id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func DeleteAlertingRule(c *gin.Context){
	var resp utils.Response
	obj := c.Param("id")
	id, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数id必须是整数")
		return
	}
	err = alarm.DeleteAlertingRule(id)
	if err != nil {
		log.Errorf("Delete alerting rule id=%d error %s", id, err.Error())
		resp.ToError(c, err)
		return
	}
	resp.ToSuccess(c)
}

func GetAlertingRulePage(c *gin.Context){
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
	exporter := c.Query("exporter")
	obj = c.DefaultQuery("alertingMetricId", "0")
	alertingMetricId, err := strconv.Atoi(obj)
	if err != nil {
		resp.ToMsgBadRequest(c, "参数alertingMetricId必须是整数")
		return
	}
	keywords := c.Query("keywords")
	data, err := alarm.GetAlertingRulePage(pageIndex, pageSize, prometheusId, exporter, alertingMetricId, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
