# consul
version: 1.9.4

## consul service
**service list**  
curl http://127.0.0.1:8500/v1/agent/services

**registry service**  
* curl -X PUT -d '{"id": "node-exporter-192.168.219.128","name": "node-exporter-192.168.219.128","address": "192.168.219.128","port": 9100,"tags": ["portal","node-exporter"],"meta": {"resource": "centos_vm"},"checks": [{"http": "http://192.168.219.128:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register
* curl -X PUT -d '{"id": "redis-exporter-192.168.219.128","name": "redis-exporter-192.168.219.128","address": "192.168.219.128","port": 9121,"tags": ["portal","redis-exporter"],"meta": {"resource": "centos_vm"},"checks": [{"http": "http://192.168.219.128:9121/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register
* curl -X PUT -d '{"id": "node-exporter-192.168.219.159","name": "node-exporter-192.168.219.159","address": "192.168.219.159","port": 9100,"tags": ["openstack","node-exporter"],"meta": {"resource": "centos_vm"},"checks": [{"http": "http://192.168.219.159:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register

**delete service**  
curl --request PUT http://127.0.0.1:8500/v1/agent/service/deregister/node-exporter

## consul kv
**get key**  
curl --header "X-Consul-Token: 111111111111" http://127.0.0.1:8500/v1/kv/prometheus/portal/alertings?recurse=true

