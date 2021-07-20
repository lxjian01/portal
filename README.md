# portal

## meta标签
1、作用
2. 便于grafana过滤区分展示
3. 资源利用率统计

2、meta
* cluster: openstack_1 openstack_2 k8s_1 k8s_2 ceph_1 ceph_2 other

* componentl: centos ubuntu windows centos_vm ubuntu_vm windows_vm

* team: bigdata payment dba openstack k8sl

* project: project_1 project_2 nova-control nova-compute k8s-master k8s-node

* app: app_1 app_2

## prometheus
1、reload  
curl -X POST http://192.168.219.128:9090/-/reload

## consul
1、service list  
curl http://127.0.0.1:8500/v1/agent/services

2、registry service  
curl -X PUT -d '{"id": "node-exporter","name": "node-exporter-172.30.12.167","address": "192.168.219.128","port": 9100,"tags": ["dev"],"meta": {"env": "dev"},"checks": [{"http": "http://192.168.219.128:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register

3、delete service  
curl --request PUT http://127.0.0.1:8500/v1/agent/service/deregister/node-exporter

