# 脚本介绍
## 脚本作用
+ 由于公司的打包之前，需要更新versions.yaml 和 images.yaml。这里的作用就是于此。
+ update_charts.py 是用于更新versions.yaml文件的。
+ update_images.py 是用于更新images.yaml文件的。

## 脚本依赖与生成
+ 两个脚本，分别依赖于旧版本的versions.yaml 和 images.yaml 文件。
+ 脚本需要在最新版本的服务器上执行，这样获取的才是最新的charts和images版本。
+ 且生成的文件是以 `new_` 开头的，如 `new_images.yaml`和`new_versions.yaml`。
+ 建议使用 venv 环境后，再执行脚本。

## 执行打印屏幕信息
+ update_charts.py 出现的提示信息
```
(venv) [root@nebula-test-68 python-pack]# python update_charts.py

[Info] : [local-volume-provisioner] chart version is right
[Error] : The current machine was not found this chart name : infra-console-service
[Error] : The current machine was not found this chart name : infra-frontend-service
[Warning] : [aurora] chart version not right, the old version is : 1.1.0-v1.1.0-222039 , the new version is : 1.2.0-v1.2.0-001-1c3436
... ...
```

+ update_images.py 出现的提示信息
``` 
(venv) [root@nebula-test-68 python-pack]# python update_images.py
[Info] : [coreos/prometheus-operator] image version is right
[Warning] :  [component/mc] image name not found in [kubectl get pods --all-namespaces -o yaml], Now use the image from [docker images] , the version is : RELEASE.2019-02-13T19-48-27Z
[Warning] :  [mysql/mysql-router] image version not right, the old version is : 8.0.16-20190618 , the new version is : 8.0.16-8269810
... ...
```