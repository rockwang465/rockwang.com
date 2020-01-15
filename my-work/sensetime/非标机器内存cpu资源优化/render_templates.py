#!/usr/bin/env python
# encoding: utf-8

import os
from jinja2 import Template
import time
import yaml

docker_registry = "registry.sensenebula.io:5000"
chartmuseum = "http://127.0.0.1:38080"

# override_file 中定义如: /opt/optimization/templates/logging/目录下elasticsearch.values.yaml（override）文件名
override_file = {
    'component': {},
    'logging': {},
    'monitoring': {},
    'nebula': {}
}


# 3. 使用jinja2将templates下的模板文件渲染后放入/opt/optimization/templates下
class template_render:
    def __init__(self, all_nodes, packages_path):
        self.render_vars = {}
        self.all_nodes = all_nodes
        self.packages_path = packages_path
        self.override_file = override_file

    # 完善传入进来的渲染模板的变量
    def init_args(self):
        base_dir = os.getcwd()
        self.templates_dir = base_dir + "/templates/"  # 获取templates的绝对路径
        # print(self.templates_dir)

        if len(self.all_nodes) > 1:
            self.deploy_mode = "distributed"
            self.license_ca = {"master": self.all_nodes[0]['ip'], "slave": self.all_nodes[1]['ip']}
        elif len(self.all_nodes) == 1:
            self.deploy_mode = "standalone"
            self.license_ca = {"master": self.all_nodes[0]['ip'], "slave": "127.0.0.1"}
        else:
            print("Error: get nodes number error, the number is : [%s], please check. " % len(self.all_nodes))

        self.render_vars = {
            "template_dir": self.templates_dir,
            "deploy_mode": self.deploy_mode,
            "docker_registry": docker_registry,
            "chartmuseum": chartmuseum,
            "license_ca": self.license_ca,
            "nodes": self.all_nodes,
            "hosts": [node['hostname'] for node in self.all_nodes]
        }

    # 渲染传入进来的文件
    def RenderConfig(self, tfile, values, ns, values_file_name, server_name):
        override_path = self.packages_path + "/override_yaml/"
        if not os.path.exists(override_path):
            os.mkdir(override_path)

        save_path = override_path + ns  # 保存的路径: /opt/optimization/override_yaml/logging
        if not os.path.exists(save_path):
            os.mkdir(save_path)

        save_file_name = save_path + "/" + values_file_name
        # print(save_path, save_file_name)
        if not os.path.exists(save_path):
            # print("Info: not exist [%s] directory, create now ." % save_path)
            os.mkdir(save_path)

        content = ""
        with open(tfile) as fp:
            content = fp.read()
        t = Template(content)
        render_res = t.render(**values)

        with open(save_file_name, 'w') as f:
            f.write(render_res)
        print("Info: save render file to [%s]" % save_file_name)

        # self.override_file字典示例: {"logging":{"elasticsearch": "/opt/optimization/logging/elsaticsearch.values.yml"}}
        self.override_file[ns][server_name] = save_file_name

    # 获取模板文件名，传参给RenderConfig进行渲染
    def get_template_file(self, charts_version):
        print("\n")
        time.sleep(1)
        for ns in charts_version:
            for server_name in charts_version.get(ns):
                values_file_name = server_name + ".values.yaml"
                # print(values_file_name)
                self.template_file_name = self.templates_dir + ns + "/" + values_file_name
                self.RenderConfig(self.template_file_name, self.render_vars, ns, values_file_name, server_name)


# 4. 定义修改需要优化服务的包中values.yaml及override文件中的request.memory request.cpu的大小，及configmap的jvm大小的函数方法
class modify_request_values:
    def __init__(self, override_file):
        self.override_file = override_file
        # print(self.override_file)

    # 修改/tmp 目录下的override文件中的cpu memory
    def modify_override_args(self, server_name, override_yaml_path):  # modify_file_path为需要修改cpu memory的文件
        # print(modify_file_path)
        # A. yaml转为json
        with open(override_yaml_path, 'r') as override_yaml_data:
            data = yaml.load(override_yaml_data)

        file = '/tmp/' + server_name + '.values.yaml'  # /tmp/kafka.values.yaml
        # B. 确认是否有request， 有则修改json中的:resources.requests.cpu 和 resources.requests.memory
        if 'data' in data:
            # print("data in yaml")
            data["data"]["resources"]["requests"]["cpu"] = 0.1
            data["data"]["resources"]["requests"]["memory"] = 0.1
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中
        elif "resources" in data:
            # print("resources in yaml")
            data["resources"]["requests"]["cpu"] = 0.1
            data["resources"]["requests"]["memory"] = 0.1
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中
        else:
            with open(file, "w") as f:
                yaml.dump(data, f)  # 保存json字符串到文件中
        print("Info: save optimized file to [%s]" % file)

    # 修改服务包中的values.yaml文件
    def modify_values_args(self, server_name, values_yaml_path):  # values_yaml_path 为服务包中的values.yaml
        bak_values_yaml = values_yaml_path + "-bak"
        res = os.system("\cp %s %s" % (values_yaml_path, bak_values_yaml))
        with open(bak_values_yaml, "r") as f:
            data = yaml.load(f)
        n = 0
        # print("\n")
        # print(bak_values_yaml)
        # print(data)
        # if data.get("prometheus").get("prometheusSpec").get("resources").get("requests"):
        if data.get("prometheus"):
            data["prometheus"]["prometheusSpec"]["resources"]["requests"]["cpu"] = 0.1
            data["prometheus"]["prometheusSpec"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("client").get("resources").get("requests"):
        if data.get("client"):
            data["client"]["resources"]["requests"]["cpu"] = 0.1
            data["client"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("data").get("resources").get("requests"):
        if data.get("data"):
            data["data"]["resources"]["requests"]["memory"] = 0.1
            data["data"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("worker").get("resources").get("requests"):
        if data.get("worker"):
            data["worker"]["resources"]["requests"]["cpu"] = 0.1
            data["worker"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("coordinator").get("resources").get("requests"):
        if data.get("coordinator"):
            data["coordinator"]["resources"]["requests"]["cpu"] = 0.1
            data["coordinator"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("reducer").get("resources").get("requests"):
        if data.get("reducer"):
            data["reducer"]["resources"]["requests"]["cpu"] = 0.1
            data["reducer"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("resources").get("requests"):
        if data.get("resources"):
            data["resources"]["requests"]["cpu"] = 0.1
            data["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("master").get("resources").get("requests"):
        if data.get("master"):
            data["master"]["resources"]["requests"]["cpu"] = 0.1
            data["master"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("volume").get("resources").get("requests"):
        if data.get("volume"):
            data["volume"]["resources"]["requests"]["cpu"] = 0.1
            data["volume"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        # if data.get("proxy").get("resources").get("requests"):
        if data.get("proxy"):
            data["proxy"]["resources"]["requests"]["cpu"] = 0.1
            data["proxy"]["resources"]["requests"]["memory"] = 0.1
            n += 1
        if server_name == "cassandra":  # cassandra 有个configmap的jvm大小配置，这里单独拿出来
            path1 = os.path.dirname(values_yaml_path)  # 取路径
            configmap_path = path1 + "/templates/config.yml"  # 这个是jinja2模板文件，所以无法使用yaml转为json格式
            os.system("sed -i 's/-Xms31G/-Xms5G/g' %s" % configmap_path)
            os.system("sed -i 's/-Xmx31G/-Xmx5G/g' %s" % configmap_path)
        if n == 0:
            print("Error : [%s] service not found resource define ." % server_name)
            sys.exit(1)
        with open(values_yaml_path, "w") as f2:  # 保存修改好的data到的服务的values.yaml中
            yaml.dump(data, f2)
