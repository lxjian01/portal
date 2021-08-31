# thanos
version: 0.17.2

The test uses local storage. In the production environment, you are advised to use S3 storage or refer to other storage configurations on the official website

## thanos 组件
* sidecar: sidecar 是伴随 prometheus 的主要组件，部署时 sidecar 容器和 prometheus 容器在同一个 pod 里。sidecar 主要功能有：1.读取和归档对象存储中的数据；2.管理 prometheus 的配置和生命周期；3.将外部标签注入 prometheus 配置中，并区分 Prometheus 的副本；4.调用 prometheus 的 promQL 接口，提供给 thanos query 查询。
* ruler: 除了 Prometheus 的规则外，Thanos Ruler 基本上执行与查询器相同的操作。唯一的区别是它可以与 Thanos 组件进行通信。ruler 是可选组件，可根据需求评估是否使用。我们大部分的告警使用了 prometheus 自身的 rule 功能，因为告警需要最新的指标。prometheus 副本数量的告警，可以使用 ruler 实现。
* query: query 是 thanos 的指标查询入口，使用 grafana 或 prometheus client 查询时，均可用 query 地址取代 prometheus 地址。可以部署多个副本，实现 query 的高可用。query 的主要功能有：1.监听 HTTP 并将查询转换为 Thanos gRPC 格式；2.汇总来自不同来源的查询结果，并且可以从 Sidecar 和 Store 中读取数据；3.在高可用设置中，Thanos Query 可以对查询结果进行去重。
* store gateway: store gateway 将对象存储的数据暴露给 thanos query 去查询。store gateway 在对象存储桶中的历史数据之上实现 store API，主要充当 API 网关，因此不需要大量的本地磁盘空间。在启动时加入 Thanos 集群，并暴露其可以访问的数据。store gateway 保存了少量本地磁盘远程块与存储桶的同步信息，通常可以安全地在重新启动期间删除数据，但这样做会增加启动时间。
* compact: compact 使用 Prometheus 2.0 存储引擎，对对象存储中的数据进行压缩分块。通常不与安全语义并发，确保每个存储桶必须部署一个。compact 还会进行数据下采样，40 小时后 5 分钟下采样，10 天后 1 小时下采样。下采样能加速大时间区间监控数据查询的速度。