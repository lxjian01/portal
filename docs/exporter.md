## exporter
* node exporter  
linux主机监控，提供cpu、内存、磁盘io、网络io、分区等监控数据  
https://github.com/prometheus/node_exporter

* wmi exporter  
windows主机监控，与node_exporter类似  
https://github.com/martinlindhe/wmi_exporter

* blackbox exporter  
提供HTTP, HTTPS, DNS, TCP and ICMP等地址探测监控功能，主要用于存活性监控  
https://github.com/prometheus/blackbox_exporter

* bind exporter  
监控 named/dns v9+服务  
https://github.com/digitalocean/bind_exporter

* cadvisor  
google开源的docker监控，k8s集成该exporter  
https://github.com/google/cadvisor

* docker exporter  
自研的docker监控exporter，cadvisor在收集openstack controller时因为docker使用主机网络，导致网络监控数据达到10W并且是无效数据，因此自研docker exporter用于openstack环境组件docker监控
 
* lbaas exporter  
自研exporter，用于查询openstack所有的lbaas并对其ip进行ping监控和相应的haproxy监控

* libvirt exporter  
自研exporter，在云主机未安装node_exporter时可通过libvirt获取云主机cpu、内存、磁盘io、网络io监控数据，无法获取文件系统的监控数据

* qga exporter  
自研exporter，通过openstack的节点上安装qga_exporter，通过libvirt获取云主机的node_exporter监控信息，加载到prometheus中，解决公有云主机和虚拟机网络不通的问题  
github.com/libvirt/libvirt-go  
github.com/prometheus/client_golang

* openstack exporter  
自研exporter，监控openstack服务状态、endpoint存活性、资源分配等数据

* openstack exporter vpc  
自研exporter，监控openstack网络相关监控数据

* openstack instance autodiscover  
自研组件，主要用于自动发现openstack云主机变化，新增或删除相应的prometheus中的监控target；目前生产环境云主机镜像已内置node_exporter或wmi_exporter

* vsphere exporter  
自研exporter，用于监控vcenter，包括物理机、虚拟机状态，ping状态，以及cpu\内存\磁盘\文件系统\网络等监控数据

* 其它  
其它如ceph、haproxy、memcached、mongodb、mysql、postgres、rabbitmq、redis、snmp exporter均是监控对应软件或服务的exporter