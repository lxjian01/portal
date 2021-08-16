package alarm

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"portal/httpd/models"
	"portal/httpd/utils"
	"time"
)

func toDeci(t time.Time) int64 {
	return t.UnixNano() / 1e8
}

func fromDeci(t int64) time.Time {
	return time.Unix(0, t*1e8)
}

func AddAlarmSender(c *gin.Context){
	var resp utils.Response
	var m models.Notice
	if err := c.ShouldBindJSON(&m);err != nil{
		resp.ToError(c, err)
		return
	}
    //tt := m.Alerts[0].EndsAt
	//year, mon, day := tt.Date()
	//hour, min, sec := tt.Clock()
	//fmt.Printf("本地时间是 %d-%d-%d %02d:%02d:%02d \n",
	//	year, mon, day, hour, min, sec)
	//将 monster 序列化
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster 序列化后=%v\n", string(data))
	//_, err := alarm.AddAlarmSender(&m)
	//if err != nil {
	//	log.Errorf("Add system menu error %s",err.Error())
	//	resp.ToError(c, err)
	//	return
	//}
	resp.Data = gin.H{"id": 1}
	resp.ToSuccess(c)
}
