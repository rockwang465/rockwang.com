# 公司服务使用python做客户端进行安装
## 1. API文档
+ API文档地址 `http://10.5.6.14:30001/#`
+ API请求地址 `http://10.5.6.14:30001/#`

## 2.部署服务
+ 名称空间
  - logging、component、nebula、default

## 3.脚本使用思路
+ 执行方式
  - help帮助: `python server_deploy_client.py --help`
  - 安装一个名称空间的服务: `python server_deploy_client.py logging`
  - 单独安装名称空间下的一个服务安装: `python server_deploy_client.py component minio`
  
+ 脚本调用外部包问题
  - 固定路径: `/opt/SenseNebula-G-v1.1.0+20190621134305`
  
## 4.要安装的服务
+ devops
  - logging的elasticsearch
  - monitoring的Prometheus
+ component组件
  - cassandra、kafka、minio、osg、seaweedfs、zookeeper、mysql、redis
  - 初始化bucker、mysql
+ license-ca加密狗
  - 安装
  - 激活
+ nebula组件
  - access-control-process、engine-image-ingress-service、engine-image-process-service、engine-static-feature-db、engine-timespace-feature-db、engine-video-ingress-service、engine-video-process-service、tailing-detection-comparison-service
+ gurade业务层
  - senseguard-ac-result-consume-default senseguard-bulk-tool-default senseguard-device-management-default senseguard-lbs-default senseguard-lib-auth-default senseguard-log-default senseguard-map-management-default senseguard-message-center-default senseguard-oauth2-default senseguard-records-export-default senseguard-records-management-default senseguard-rule-management-default senseguard-scheduled-default senseguard-search-face-default senseguard-target-export-default senseguard-td-result-consume-default senseguard-timezone-management-default senseguard-tools-default senseguard-uums-default senseguard-watchlist-management-default


  