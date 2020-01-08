# 脚本说明
## 脚本作用
+ 此脚本用于cpu、内存的资源优化
+ 将requests的cpu和memory缩减到0.5

## 脚本执行命令
+ `source /root/venv/bin/active`
+ `python main.py`
+ 注意: 请将脚本放在集群的第一台master上执行

## 脚本相关文件位置
+ 执行后生成/opt/optimization/目录
+ 将templates模板文件渲染后放入到/opt/optimization/override_yaml/下的名称空间中
+ 然后将上面的模板文件中cpu、memory大小的资源缩减后，把override文件放在了/tmp/目录下，已 `service_name.values.yaml` 方式命名

## 脚本功能使用
+ 如果后期需要使用override文件，建议使用/tmp目录下的。如果/tmp下文件被清理了，可以修改main.py脚本，只将更新部分的函数注释掉即可