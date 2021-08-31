# AlertManager
versio: 0.21.0

## 参数说明：
* group_b: 告警组名
* group_wai: 当收到告警的时候，等待10秒确认时间内是否有新告警，如有则一并发送
* group_interva: 发送前等待时间，配置中设置接收到告警后再等待20秒发送
* repeat_interva: 重复告警周期时间，由于是测试所以设置比较短，在生产环境中一般设置为1h
* receive: 指示信息推送给谁，此处设置的值必须在receivers中能够找到
* webhook_config: 调用api的url地址，实验环境使用192.168.111.1创建的一个api

## API：
> https://www.kancloud.cn/pshizhsysu/prometheus/1872669  

**Alert**
```
GET    /api/v2/alerts
POST   /api/v2/alerts
```

**AlertGroup**
```
GET    /api/v2/alerts/groups
```

**General**
```
GET    /api/v2/status
```

**Receiver**
```
GET    /api/v2/receivers
```

**Silence**
```
GET    /api/v2/silences
POST   /api/v2/silences
GET    /api/v2/silence/{silenceID}
DELETE /api/v2/silence/{silenceID}
```

## MANAGEMENT API
Alertmanager provides a set of management API to ease automation and integrations.

**Health check**
```
GET /-/healthy
```
This endpoint always returns 200 and should be used to check Alertmanager health.

**Readiness check**
```
GET /-/ready
```
This endpoint returns 200 when Alertmanager is ready to serve traffic (i.e. respond to queries).

**Reload**
```
POST /-/reload
```
This endpoint triggers a reload of the Alertmanager configuration file.  
An alternative way to trigger a configuration reload is by sending a SIGHUP to the Alertmanager process.