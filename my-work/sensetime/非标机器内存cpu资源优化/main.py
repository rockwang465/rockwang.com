# encoding: utf-8

import os
import sys
import yaml
import time

ns = ['component', 'logging', 'monitoring', 'nebula']
optimization_server_name = {
    'component': ['cassandra', 'kafka'],
    'logging': ['elasticsearch'],
    'monitoring': ['prometheus-operator'],
    'nebula': ['engine-timespace-feature-db']
}
charts_version = {
    'component': [],
    'logging': [],
    'monitoring': [],
    'nebula': []
}
packages_path = '/opt/optimization'

# tmp_override_file 中定义/tmp/目录下override.yaml文件名
tmp_override_file = {
    'component': {},
    'logging': {},
    'monitoring': {},
    'nebula': {}
}


# 获取charts包
# A. 通过helm list | grep 'xxx' 获取版本号，写入到optimization_server.txt中，
#    再从optimization_server.txt中读取版本号，然后通过10.151.3.75获取charts包。
# B. 通过helm list | grep 'xxx' 获取版本号，并从本地/data/charts/获取charts包。 -- 不合适
# C. 通过helm list | grep 'xxx' 获取版本号，并从10.151.3.75获取charts包。 -- 不合适

# 1. helm list 获取版本号，并写入optimization_server.txt中
class get_charts_packages:
    def __init__(self):
        self.charts_version = charts_version
        self.tmp_override_file = tmp_override_file

    # 获取helm list 中的版本信息
    def get_helm_version(self):
        for key_ns in optimization_server_name:
            for server_name in optimization_server_name.get(key_ns):
                res = os.popen("helm list | grep %s | grep -v elasticsearch-curator | awk '{print $9}'" % server_name)
                res_data = res.read().strip()
                if res_data:
                    self.charts_version[key_ns].append(res_data)
                else:
                    print('Error : Not found server name : [%s]' % server_name)
        # print("Info : Here is the chart information :")
        # print(self.charts_version)
        print("\n")

    # 通过上面获取的版本信息，这里进行fetch下载
    def fetch_helm_packages(self):
        for key_ns in self.charts_version:
            for charts_info in charts_version.get(key_ns):
                if not os.path.exists(packages_path):
                    print("Info : create directory %s" % packages_path)
                    os.mkdir(packages_path)
                print(
                    "Info : [cd %s && helm fetch http://10.151.3.75:8080/charts/%s.tgz]" % (packages_path, charts_info))
                res = os.system(
                    "cd %s && helm fetch http://10.151.3.75:8080/charts/%s.tgz" % (packages_path, charts_info))
                if res != 0:
                    print("Error : failure to [helm fetch http://10.151.3.75:8080/charts/%s.tgz]")
                    sys.exit(1)

    # 获取 /tmp/下各服务的override.values.yaml文件名
    def get_override_name(self):
        for key_name in optimization_server_name:
            # print(optimization_server_name.get(key_name))
            for server_name in optimization_server_name.get(key_name):
                # print(server_name)
                res = os.popen('ls /tmp/%s-%s*' % (server_name, key_name))
                server_file = res.read().strip()
                if server_file:
                    # 添加对应服务在/tmp/目录下的overrid文件名
                    self.tmp_override_file[key_name][
                        server_name] = server_file  # tmp_override_file['component']['kafka']='/tmp/kafka-component.1573202714'
                else:
                    print("Error : Not found [/tmp/%s-%s....] file" % (server_name, key_name))
                    sys.exit(1)
        # print(self.tmp_override_file)


class modify_request_values:
    # 修改/tmp 目录下的override文件中的cpu memory
    def modify_tmp_args(self, server_name, tmp_yaml_path):  # modify_file_path为需要修改cpu memory的文件
        # print(modify_file_path)
        # A. yaml转为json
        with open(tmp_yaml_path, 'r') as tmp_yaml_data:
            data = yaml.load(tmp_yaml_data)

        file = '/tmp/' + server_name + '.values.yaml'  # /tmp/kafka.values.yaml
        # B. 确认是否有request， 有则修改json中的:resources.requests.cpu 和 resources.requests.memory
        if 'data' in data:
            print("data in yaml")
            data["data"]["resources"]["requests"]["cpu"] = 0.5
            data["data"]["resources"]["requests"]["memory"] = 0.5
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中
        elif "resources" in data:
            # print("resources in yaml")
            data["resources"]["requests"]["cpu"] = 0.5
            data["resources"]["requests"]["memory"] = 0.5
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中
        else:
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中

    # 修改服务包中的values.yaml文件
    def modify_values_args(self, server_name, values_yaml_path):  # values_yaml_path 为服务包中的values.yaml
        bak_values_yaml = values_yaml_path + "-bak"
        res = os.system("\cp %s %s" % (values_yaml_path, bak_values_yaml))
        with open(bak_values_yaml, "r") as f:
            data = yaml.load(f)
        if server_name == "prometheus-operator":  # prometheus-operator的values.yaml中requests字段定义不同
            data["prometheus"]["prometheusSpec"]["resources"]["requests"]["cpu"] = 0.5
            data["prometheus"]["prometheusSpec"]["resources"]["requests"]["memory"] = 0.5
        elif server_name == "elasticsearch":
            data["client"]["resources"]["requests"]["cpu"] = 0.5
            data["client"]["resources"]["requests"]["memory"] = 0.5
            data["data"]["resources"]["requests"]["memory"] = 0.5
            data["data"]["resources"]["requests"]["memory"] = 0.5
        elif server_name == "engine-timespace-feature-db":
            data["worker"]["resources"]["requests"]["cpu"] = 0.5
            data["worker"]["resources"]["requests"]["memory"] = 0.5
            data["coordinator"]["resources"]["requests"]["cpu"] = 0.5
            data["coordinator"]["resources"]["requests"]["memory"] = 0.5
            data["reducer"]["resources"]["requests"]["cpu"] = 0.5
            data["reducer"]["resources"]["requests"]["memory"] = 0.5
        elif server_name == "cassandra":  # cassandra 有个configmap的jvm大小配置，这里单独拿出来
            data["resources"]["requests"]["cpu"] = 0.5
            data["resources"]["requests"]["memory"] = 0.5
            path1 = os.path.dirname(values_yaml_path)  # 取路径
            configmap_path = path1 + "/templates/config.yml"  # 这个是jinja2模板文件，所以无法使用yaml转为json格式
            os.system("sed -i 's/-Xms31G/-Xms5G/g' %s" % configmap_path)
            os.system("sed -i 's/-Xmx31G/-Xmx5G/g' %s" % configmap_path)
        elif server_name == "kafka":  # resources.requests，待考虑其他文件从这里走
            data["resources"]["requests"]["cpu"] = 0.5
            data["resources"]["requests"]["memory"] = 0.5
        else:
            print("Error : [%s] This service is not defined, please define first." % server_name)
            sys.exit(1)
            # 新的服务，需要修改 optimization_server_name字典，和 这里确定values.yaml的定义格式。最好上面/tmp/server_name 中也确认下。

        with open(values_yaml_path, "w") as f2:  # 保存修改好的data到的服务的values.yaml中
            yaml.dump(data, f2)


# 获取之前保存字典中的override文件路径，及各服务的values.yaml的路径
class get_modify_file:
    # 修改/tmp/目录下override文件中request的memory及cpu的大小
    def modify_tmp_override(self, tmp_override_file):
        # print(get_packages.tmp_override_file)
        for key_name in tmp_override_file:  # key_name: component
            # print(tmp_override_file)
            # print(key_name)
            for name in tmp_override_file.get(
                    key_name):  # tmp_override_file.get(key_name): {'elasticsearch': '/tmp/elasticsearch-logging.1573202709'}
                server_name = name
                tmp_yaml_path = tmp_override_file.get(key_name).get(name)
                modify_values.modify_tmp_args(server_name, tmp_yaml_path)  # 传入服务名，文件路径

    # 修改values.yaml中的cpu memory
    def modify_values_yaml(self):
        # 解压所有fetch下来的charts包
        res = os.system("cd %s && for i in `ls *.tgz` ; do tar xf $i ; done" % packages_path)
        if res != 0:
            print("Error : failure to [cd %s && ls *.tgz | xargs tar xf]" % packages_path)
            sys.exit(1)
        for key_name in optimization_server_name:
            for name in optimization_server_name.get(key_name):
                values_yaml_path = packages_path + "/" + name + "/values.yaml"
                server_name = name
                modify_values.modify_values_args(server_name, values_yaml_path)


# 资源优化后，开始更新服务
class update_optimization_service:
    def upgrade_service(self):
        print(1)
        for key_ns in optimization_server_name:
            for server_name in optimization_server_name.get(key_ns):
                tmp_override_new_file = "/tmp/" + server_name + ".values.yaml"
                # print(tmp_override_new_file)
                os.chdir("%s/%s" % (packages_path, server_name))
                # print("%s/%s" % (packages_path, server_name))
                os.system("helm upgrade -i %s-%s --namespace=%s -f /tmp/%s.values.yaml ." % (
                    server_name, key_ns, key_ns, server_name))
                # print("helm upgrade -i %s-%s --namespace=%s -f /tmp/%s.values.yaml ." % (server_name, key_ns, key_ns, server_name))
                time.sleep(2)


if __name__ == '__main__':
    # 1. 获取需要优化服务的版本信息、拉取对应版本的包、找到/tmp/ 目录下的override文件
    get_packages = get_charts_packages()
    get_packages.get_helm_version()  # A.获取需要优化服务的版本信息
    # get_packages.fetch_helm_packages()  # B.拉取对应版本的包
    get_packages.get_override_name()  # C.找到/tmp/ 目录下的override文件

    # 2. 定义修改需要优化服务的包中values.yaml及override文件中的request.memory request.cpu的大小，及configmap的jvm大小的函数方法
    modify_values = modify_request_values()

    # 3. 获取之前保存字典中的override文件路径，及各服务的values.yaml的路径
    get_file = get_modify_file()
    get_file.modify_tmp_override(get_packages.tmp_override_file)  # 修改/tmp/目录下的override文件
    get_file.modify_values_yaml()  # 修改服务中values.yaml文件

    # 4. 资源优化后，开始更新服务
    update_service = update_optimization_service()
    update_service.upgrade_service()
