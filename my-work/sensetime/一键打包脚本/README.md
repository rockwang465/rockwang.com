## 执行脚本
`python main.py`

## 依赖关系
+ A.get_version.py
  - 用于获取标准环境的charts和images的版本信息，合并后，导出文件 versions.json 。
+ B.connect_10.py
  - 连接到10服务器
  - 将versions.json 和 pack_charts_images.py 推到 10.5.6.10 的固定目录中。
  - 在10.5.6.10上执行 pack_charts_images.py 脚本，开始打包。
+ C.pack_charts_images.py
  - 用于在10.5.6.10上打包charts和images。
  