#!/usr/bin/env python
# -*- coding: utf-8 -*-

from jinja2 import Environment, PackageLoader
import os

docker_registry = "10.5.6.14:5000"
# standalone or distributed
deploy_mode = "standalone"
k8s_node = "10.5.6."

namespaces = ['default', 'nebula', 'component', 'logging', 'monitoring']
# component_temp = ['cassandra', 'kafka', 'minio', 'mysql', 'osg', 'prometheus_mysql_exporter', 'prometheus_redis_exporter', 'redisoperator', 'seaweedfs', 'zookeeper']

env = Environment(loader=PackageLoader('jinja2_files'))

groups = [{
    "kube_master": ["nebula-test-14", "nebula-test-18", "nebula-test-19"],
    "kube_node": []
}]


# use namespace name to create directory
def create_dir():
    for dir in namespaces:
        # path = 'jinja2_files/templates/' + dir
        path = 'result_files/' + dir
        isExists = os.path.exists(path)
        if not isExists:
            # if not exists directory, then create directory
            os.makedirs(path)
            print('[Info] :' + path + ' create successfully')
        else:
            # if directory already exists,then notice is exists
            print('[Info] :' + path + ' directory is exists')


# component template files
def define_temp_services():
    # render1 = ['cassandra', 'kafka', 'minio']  # just defined "standalone"
    # render2 = ['mysql']  # have defined {{ groups['kube-master']|first }}
    for ns in namespaces:
        recode = os.popen('ls ./jinja2_files/templates/%s/' % ns)
        files_list = recode.readlines()
        print(files_list)
        for file in files_list:
            print("[Info] : Now start transfer the %s ---------->start" % file)
            # os.system('touch %s' % file)
            file = file.replace("\n", "")  # let "\n" replace empty
            # get jinja2 template file
            # [example] : template = env.get_template('test.j2')
            file_path = ns + '/' + file
            template = env.get_template(file_path)

            # render template
            # [example] : content = template.render(name='liuhao', age='18', country='China')
            # content = template.render(docker_registry=docker_registry, deploy_mode="standalone")
            content = template.render(docker_registry=docker_registry, deploy_mode="standalone", groups=groups)
            # content = template.render(docker_registry=docker_registry)

            print("[Info] : Now was finished transfer the %s ---------->end" % file)

            # save result to values.yaml
            with open('./result_files/' + ns + '/' + file, 'w') as fp:
                fp.write(content)


create_dir()
define_temp_services()
