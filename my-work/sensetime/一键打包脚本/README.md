## 1. 执行脚本
`python main.py`

## 2. 脚本作用与依赖关系
### 2.1 标准环境机器上
+ A.get_version.py
  - 用于获取标准环境的charts和images的版本信息，合并后，导出文件 versions.json 。
+ B.connect_10.py
  - 连接到10服务器。
  - 将versions.json 和 pack_charts_images.py 推到 10.5.6.10 的固定目录中。
  - 在10.5.6.10上执行 pack_main.py 脚本，开始打包。

### 2.2 10.5.6.10机器上
+ A.pack_main.py
  - 用于在10.5.6.10上打包charts、images、base 。
  - 依赖pack_charts_images.py 和 pack_base.py 。
+ B.pack_charts_images.py 
  - 用于打包 charts和images。
+ C.pack_base.py
  - 用于打包 base目录 、 versions.json 、 和整个release封装。
  - base包含:  yum-data、 tools、 license-ca、 infra-ansible.tgz 。
  - versions.json为images和charts的版本记录文件。