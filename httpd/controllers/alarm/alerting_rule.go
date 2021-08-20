package alarm

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/services/alarm"
	"portal/httpd/utils"
	"strconv"
)

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
