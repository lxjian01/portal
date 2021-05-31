package monitor

import (
	consulapi "github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"portal/global/consul"
	"portal/global/myorm"
	"portal/httpd/models"
	"portal/httpd/utils"
)

func AddMonitorTarget(m *models.MonitorTarget) (int, error) {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		// find monitor cluster
		monitorCluster := models.MonitorCluster{}
		err := tx.Table("monitor_cluster").Where("monitor_cluster_id = ?", m.MonitorClusterId).Find(&monitorCluster).Error
		if err != nil {
			return err
		}
		// find monitor component
		monitorComponent := models.MonitorComponent{}
		err = tx.Table("monitor_component").Where("monitor_component_id = ?", m.MonitorComponentId).Find(&monitorComponent).Error
		if err != nil {
			return err
		}
		// add monitor target
		err = tx.Table("monitor_target").Create(m).Error
		if err != nil {
			return err
		}
		// registry consul service
		meta := make(map[string]string, 0)
		meta["cluster"] = monitorCluster.Code
		meta["component"] = monitorComponent.Code
		meta["exporter"] = monitorComponent.Exporter
		client := consul.GetClient()
		check := consulapi.AgentServiceCheck{
			HTTP:     m.Url,
			Interval: m.Interval,
		}
		registration := consulapi.AgentServiceRegistration{
			ID:    m.Url,
			Name:  m.Name,
			Meta:  meta,
			Check: &check,
		}
		err = client.Agent().ServiceRegister(&registration)
		return err
	})
	return m.Id, err
}

func DeleteMonitorTarget(id int) error {
	err := myorm.GetOrmDB().Transaction(func(tx *gorm.DB) error {
		monitorTarget := models.MonitorTarget{}
		txQuery := tx.Table("monitor_cluster").Where("monitor_target_id = ?", id)
		err := txQuery.Find(&monitorTarget).Error
		if err != nil {
			return err
		}
		err = txQuery.Delete(&models.MonitorTarget{}).Error
		if err != nil {
			return err
		}
		client := consul.GetClient()
		err = client.Agent().ServiceDeregister(monitorTarget.Url)
		return err
	})
	return err
}

func GetMonitorTargetPage(pageIndex int, pageSize int, monitorClusterId int, monitorComponentId int, alarmGroupId int, keywords string) (*utils.PageData, error) {
	dataList := make([]models.MonitorTargetPage, 0)
	tx := myorm.GetOrmDB().Table("monitor_target")
	tx.Select("monitor_target.id","monitor_target.monitor_cluster_id","monitor_target.monitor_component_id","monitor_target.name","monitor_target.url","monitor_target.interval","monitor_target.remark","monitor_target.create_user","monitor_target.create_time","monitor_target.update_user","monitor_target.update_time","monitor_cluster.code as monitor_cluster_code","monitor_cluster.name as monitor_cluster_name","monitor_component.code as monitor_component_code","monitor_component.name as monitor_component_name","monitor_component.exporter")
	tx.Joins("left join monitor_cluster on monitor_cluster.id = monitor_target.monitor_cluster_id")
	tx.Joins("left join monitor_component on monitor_component.id = monitor_target.monitor_component_id")
	if monitorClusterId != 0 {
		tx.Where("monitor_cluster_id = ?", monitorClusterId)
	}
	if monitorComponentId != 0 {
		tx.Where("monitor_component_id = ?", monitorComponentId)
	}
	if alarmGroupId != 0 {
		tx.Where("alarm_group_id = ?", alarmGroupId)
	}
	if keywords != "" {
		likeStr := "%" + keywords + "%"
		tx.Where("name like ? or url like ?", likeStr)
	}

	pageData, err := utils.GetPageData(tx, pageIndex, pageSize, &dataList)
	if err != nil {
		return nil, err
	}

	pageData.Data = &dataList
	return pageData, nil
}