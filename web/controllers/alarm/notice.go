package alarm

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"portal/web/models"
	"portal/web/services/alarm"
	"portal/web/utils"
	"strconv"
)

func AddAlarmNotice(c *gin.Context){
	var resp utils.Response
	var notice models.AlertManagerNotice
	if err := c.ShouldBindJSON(&notice);err != nil{
		resp.ToError(c, err)
		return
	}
	//输出序列化后的结果
	data, err := json.Marshal(&notice)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	fmt.Printf("monster 序列化后=%v\n", string(data))

	higAlert := notice.Alerts[0]
	for _,alert := range notice.Alerts {
		if alert.Labels["status"] == "critical" {
			thisAlert := alert
			higAlert = thisAlert
			break
		}
		if higAlert.Status != "critical" && alert.Labels["status"] == "warning" {
			thisAlert := alert
			higAlert = thisAlert
			break
		}
	}
	var m models.AlarmNotice
	m.PrometheusCode = notice.CommonLabels["pcode"]
	m.Fingerprint = higAlert.Fingerprint
	m.AlertName = notice.CommonLabels["alertname"]
	m.Instance = notice.CommonLabels["instance"]
	labels := make(map[string]string)
	for key,value := range notice.CommonLabels {
		if key != "pcode" && key != "alertname" && key != "instance" && key != "severity" {
			labels[key] = value
		}
	}
	labelsBytes,_ :=json.Marshal(labels)
	m.Labels = datatypes.JSON(labelsBytes)
	m.Summary = higAlert.Annotations["summary"]
	m.Description = higAlert.Annotations["description"]
	m.Status = higAlert.Status
	m.Severity = higAlert.Labels["status"]
	m.StartAt = models.MyTime{ Time: higAlert.StartsAt }
	m.AlarmNumber = 1
	if higAlert.Status == "resolved" {
		m.EndAt = models.MyTime{ Time: higAlert.EndsAt }
	}

	err = alarm.AddOrUpdateAlarmNotice(&m)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = data
	resp.ToSuccess(c)
}

func GetAlarmNoticePage(c *gin.Context){
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
	prometheusCode := c.DefaultQuery("prometheusCode", "")
	monitorResourceCode := c.DefaultQuery("monitorResourceCode", "")
	severity := c.DefaultQuery("severity", "")
	status := c.DefaultQuery("status", "")
	keywords := c.DefaultQuery("keywords", "")
	data, err := alarm.GetAlarmNoticePage(pageIndex, pageSize, prometheusCode, monitorResourceCode, severity, status, keywords)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}

	resp.Data = data
	resp.ToSuccess(c)
}
