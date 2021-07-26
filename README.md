# portal

## meta标签
1、作用
2. 便于grafana过滤区分展示
3. 资源利用率统计

2、meta
* cluster: openstack_1 openstack_2 k8s_1 k8s_2 ceph_1 ceph_2 other

* resource: centos ubuntu windows centos_vm ubuntu_vm windows_vm

* team: bigdata payment dba openstack k8sl

* project: project_1 project_2 nova-control nova-compute k8s-master k8s-node

* app: app_1 app_2

## prometheus
1、reload  
curl -X POST http://192.168.219.128:9090/-/reload

## consul
### consul service
1、service list  
curl http://127.0.0.1:8500/v1/agent/services

2、registry service  
curl -X PUT -d '{"id": "node-exporter-192.168.219.128","name": "node-exporter-192.168.219.128","address": "192.168.219.128","port": 9100,"tags": ["portal","node-exporter"],"meta": {"env": "dev"},"checks": [{"http": "http://192.168.219.128:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register
curl -X PUT -d '{"id": "redis-exporter-192.168.219.128","name": "redis-exporter-192.168.219.128","address": "192.168.219.128","port": 9121,"tags": ["portal","redis-exporter"],"meta": {"env": "dev"},"checks": [{"http": "http://192.168.219.128:9121/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register
curl -X PUT -d '{"id": "node-exporter-192.168.219.159","name": "node-exporter-192.168.219.159","address": "192.168.219.159","port": 9100,"tags": ["openstack","node-exporter"],"meta": {"env": "dev"},"checks": [{"http": "http://192.168.219.159:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register

3、delete service  
curl --request PUT http://127.0.0.1:8500/v1/agent/service/deregister/node-exporter

### consul kv
1、get key
curl --header "X-Consul-Token: 111111111111" http://127.0.0.1:8500/v1/kv/prometheus/portal/alertings?recurse=true

