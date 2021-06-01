# portal

# prometheus
curl -X POST http://192.168.219.128:9090/-/reload

# consul
curl http://127.0.0.1:8500/v1/agent/services
curl -X PUT -d '{"id": "node-exporter","name": "node-exporter-172.30.12.167","address": "192.168.219.128","port": 9100,"meta": {"env": "dev"},"checks": [{"http": "http://192.168.219.128:9100/metrics", "interval": "5s"}]}'  http://127.0.0.1:8500/v1/agent/service/register
curl --request PUT http://127.0.0.1:8500/v1/agent/service/deregister/node-exporter

