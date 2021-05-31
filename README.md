# portal

# consul
curl http://127.0.0.1:8500/v1/agent/services
curl -X PUT -d '{"id": "node-exporter","name": "node-exporter-172.30.12.167",meta{"env":"dev",},"check": {"http": "http://192.168.219.128:9100/metrics", "interval": "1m"}}'  http://127.0.0.1:8500/v1/agent/service/register
curl --request PUT http://127.0.0.1:8500/v1/agent/service/deregister/cadvisor-exporter-192.168.219.128