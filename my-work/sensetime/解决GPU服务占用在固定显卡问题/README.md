# 脚本说明
## 脚本作用
+ 要求不同服务占用在不同显卡上

## 显卡占用说明
+ 3个服务服务名及占用显卡名
  - ips(engine-face-extract-service pod服务)，显卡服务中名称: ./engine-image-process-service
  - tfd(engine-timespace-feature-db pod服务)，显卡服务中名称: /gpu-searcher/search-worker
  - struct-tfd(engine-struct-timespace-feature-db pod服务)，显卡服务中名称: /gpu-searcher/search-worker
  - vps(engine-video-process-service pod服务)，显卡服务中名称: ./video-process-service-worker

+ 占用显卡要求(默认都为4张显卡)
  - ips和tfd和struct-tfd三个服务(3个gpu服务)同时放在一张卡上。
  - vps的3个worker(3个gpu服务)分别放在另外3张卡上(即一个worker占一张卡)。

## 脚本使用方法
+ 放于crontab中进行定时触发检测即可。

## 注意
+ 1.这里默认是使用override.yaml安装这几个服务的。
+ 2.因为脚本是调用sophon(8000)平台，而sophon平台是获取helm client中的override.yaml文件的。
+ 3.也就是说，如果你没有用override.yaml 或者在sophon中配置单独的配置，那么通过8000的api是拿不到config的任何配置的。
+ 4.则脚本里调的8000接口中，config中无数据，则后面的脚本操作完全无效。
+ 5.所以，特别注意: 一定要使用override.yaml进行安装服务，否则此脚本完全无效。