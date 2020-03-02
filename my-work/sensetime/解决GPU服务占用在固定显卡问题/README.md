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
  