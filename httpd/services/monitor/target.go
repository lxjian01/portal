package monitor

import (
	consulapi "github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"portal/global/consul"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
	"strconv"
	"strings"
)

func AddMonitorTarget(m *models.MonitorTargetAdd) (int, error) {
	prometheusUrl := ""
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		prometheus := models.Prometheus{}
		err := tx.Table("prometheus").Where("id = ?", m.PrometheusId).Find(&prometheus).Error
		if err != nil {
			return err
		}
		prometheusUrl = prometheus.Url
		// find monitor cluster
		monitorCluster := models.MonitorCluster{}
		err = tx.Table("monitor_cluster").Where("id = ?", prometheus.MonitorClusterId).Find(&monitorCluster).Error
		if err != nil {
			return err
		}
		// find monitor resource
		monitorResource := models.MonitorResource{}
		err = tx.Table("monitor_resource").Where("id = ?", m.MonitorResourceId).Find(&monitorResource).Error
		if err != nil {
			return err
		}
		// add monitor target
		err = tx.Table("monitor_target").Create(m).Error
		if err != nil {
			return err
		}
		// add monitor target alarm group
		if len(m.AlarmGroupIds) > 0 {
			var mtag []models.MonitorTargetAlarmGroup
			for _, item := range m.AlarmGroupIds{
				var alarmGroupUser models.MonitorTargetAlarmGroup
				alarmGroupUser.MonitorTargetId = m.Id
				alarmGroupUser.AlarmGroupId = item
				mtag = append(mtag, alarmGroupUser)
			}
			err = tx.Table("monitor_target_alarm_group").Create(&mtag).Error
			if err != nil {
				return err
			}
		}
		// registry consul service
		tags := make([]string, 0)
		tags = append(tags, monitorCluster.Code)
		tags = append(tags, monitorResource.Exporter)
		meta := make(map[string]string, 0)
		meta["resource"] = monitorResource.Code
		client := consul.GetClient()
		ip := strings.Split(m.Url,"/")[2]
		tempIp := strings.Split(ip,":")
		address := tempIp[0]
		port,err := strconv.Atoi(tempIp[1])
		if err != nil {
			return err
		}
		check := consulapi.AgentServiceCheck{
			HTTP:     m.Url,
			Interval: m.Interval,
		}
		serviceId := strconv.Itoa(m.Id)
		registration := consulapi.AgentServiceRegistration{
			ID: serviceId,
			Name: m.Name,
			Address: address,
			Port: port,
			Tags: tags,
			Meta:  meta,
			Check: &check,
		}
		err = client.Agent().ServiceRegister(&registration)
		return err
	})

	// reload prometheus
	err = utils.PrometheusReload(prometheusUrl)
	return m.Id, err
}

func UpdateMonitorTarget(m *models.MonitorTargetAdd) (int, error) {
	prometheusUrl := ""
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find prometheus
		prometheus := models.Prometheus{}
		err := tx.Table("prometheus").Where("id = ?", m.PrometheusId).Find(&prometheus).Error
		if err != nil {
			return err
		}
		prometheusUrl = prometheus.Url
		// find monitor cluster
		monitorCluster := models.MonitorCluster{}
		err = tx.Table("monitor_cluster").Where("id = ?", prometheus.MonitorClusterId).Find(&monitorCluster).Error
		if err != nil {
			return err
		}
		// find monitor resource
		monitorResource := models.MonitorResource{}
		err = tx.Table("monitor_resource").Where("id = ?", m.MonitorResourceId).Find(&monitorResource).Error
		if err != nil {
			return err
		}
		// update monitor target
		err = tx.Table("monitor_target").Where("id = ?", m.Id).Updates(m).Error
		if err != nil {
			return err
		}
		// delete monitor target alarm group
		err = tx.Table("monitor_target_alarm_group").Where("monitor_target_id = ?", m.Id).Delete(&models.MonitorTargetAlarmGroup{}).Error
		if err != nil {
			return err
		}
		// add monitor target alarm group
		if len(m.AlarmGroupIds) > 0 {
			var mtag []models.MonitorTargetAlarmGroup
			for _, item := range m.AlarmGroupIds{
				var alarmGroupUser models.MonitorTargetAlarmGroup
				alarmGroupUser.MonitorTargetId = m.Id
				alarmGroupUser.AlarmGroupId = item
				mtag = append(mtag, alarmGroupUser)
			}
			err = tx.Table("monitor_target_alarm_group").Create(&mtag).Error
			if err != nil {
				return err
			}
		}
		// registry consul service
		tags := make([]string, 0)
		tags = append(tags, monitorCluster.Code)
		tags = append(tags, monitorResource.Exporter)
		meta := make(map[string]string, 0)
		meta["resource"] = monitorResource.Code
		client := consul.GetClient()
		ip := strings.Split(m.Url,"/")[2]
		tempIp := strings.Split(ip,":")
		address := tempIp[0]
		port,err := strconv.Atoi(tempIp[1])
		if err != nil {
			return err
		}
		check := consulapi.AgentServiceCheck{
			HTTP:     m.Url,
			Interval: m.Interval,
		}
		serviceId := strconv.Itoa(m.Id)
		registration := consulapi.AgentServiceRegistration{
			ID:    serviceId,
			Name:  m.Name,
			Address: address,
			Port: port,
			Tags: tags,
			Meta:  meta,
			Check: &check,
		}
		err = client.Agent().ServiceRegister(&registration)
		return err
	})

	// reload prometheus
	err = utils.PrometheusReload(prometheusUrl)
	return m.Id, err
}

func DeleteMonitorTarget(id int) error {
	prometheusUrl := ""
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find monitor target
		monitorTarget := models.MonitorTarget{}
		txQuery := tx.Table("monitor_target").Where("id = ?", id)
		err := txQuery.Find(&monitorTarget).Error
		if err != nil {
			return err
		}
		// find prometheus
		prometheus := models.Prometheus{}
		err = tx.Table("prometheus").Where("id = ?", monitorTarget.PrometheusId).Find(&prometheus).Error
		if err != nil {
			return err
		}
		prometheusUrl = prometheus.Url
		// delete monitor target alarm group
		err = tx.Table("monitor_target_alarm_group").Where("monitor_target_id = ?", id).Delete(&models.MonitorTargetAlarmGroup{}).Error
		if err != nil {
			return err
		}
		// delete monitor target
		err = txQuery.Delete(&models.MonitorTarget{}).Error
		if err != nil {
			return err
		}
		// delete consul service
		client := consul.GetClient()
		serviceId := strconv.Itoa(monitorTarget.Id)
		err = client.Agent().ServiceDeregister(serviceId)
		return err
	})
	// reload prometheus
	err = utils.PrometheusReload(prometheusUrl)
	return err
}

type alarmGroupList struct {
	AlarmGroupId int `gorm:"column:alarm_group_id" json:"alarmGroupId"`
	AlarmGroupName string `gorm:"column:alarm_group_name" json:"alarmGroupName"`
	MonitorTargetId int `gorm:"column:monitor_target_id" json:"monitorTargetId"`
}

func GetMonitorTargetPage(pageIndex int, pageSize int, monitorClusterId int, prometheusId int, monitorResourceId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorTargetPage, 0)
	tx := myorm.GetOrmDB().Table("monitor_target")
	tx.Select("monitor_target.id","monitor_target.prometheus_id","monitor_target.monitor_resource_id","monitor_target.name","monitor_target.url","monitor_target.interval","monitor_target.remark","monitor_target.create_user","monitor_target.create_time","monitor_target.update_user","monitor_target.update_time","monitor_cluster.code as monitor_cluster_code","monitor_cluster.name as monitor_cluster_name","prometheus.name as prometheus_name","prometheus.monitor_cluster_id","monitor_resource.code as monitor_resource_code","monitor_resource.name as monitor_resource_name","monitor_resource.exporter")
	tx.Joins("left join prometheus on prometheus.id = monitor_target.prometheus_id")
	tx.Joins("left join monitor_resource on monitor_resource.id = monitor_target.monitor_resource_id")
	tx.Joins("left join monitor_cluster on monitor_cluster.id = prometheus.monitor_cluster_id")
	if monitorClusterId != 0 {
		tx.Where("prometheus.monitor_cluster_id = ?", monitorClusterId)
	}
	if prometheusId != 0 {
		tx.Where("monitor_target.prometheus_id = ?", prometheusId)
	}
	if monitorResourceId != 0 {
		tx.Where("monitor_target.monitor_resource_id = ?", monitorResourceId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("monitor_target.name like ? or monitor_target.url like ?", likeStr, likeStr)
	}

	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}
	alarmGroupList := make([]alarmGroupList, 0)
	myorm.GetOrmDB().Table("alarm_group").Select("monitor_target_alarm_group.alarm_group_id","alarm_group.name as alarm_group_name","monitor_target_alarm_group.monitor_target_id").Joins("left join monitor_target_alarm_group on monitor_target_alarm_group.alarm_group_id = alarm_group.id").Find(&alarmGroupList)
	for index, item := range dataList {
		for _, group := range alarmGroupList {
			if item.Id == group.MonitorTargetId {
				value := group
				dataList[index].GroupList = append(dataList[index].GroupList, &value)
			}
		}
	}

	pageData.Data = &dataList
	return pageData, nil
}