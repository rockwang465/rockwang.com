# -*- coding: utf-8 -*-
import logging
import paramiko
import xlrd

logger = logging.getLogger("image_check")
logger.setLevel(logging.INFO)
formatter = logging.Formatter("[%(asctime)-15s][%(levelname)-5s] %(message)s", datefmt="%Y-%m-%d %H:%M:%S")
file_handler = logging.FileHandler("image_check.log", "w")
file_handler.setFormatter(formatter)
logger.addHandler(file_handler)

# 引擎层组件
NebulaEngineNamespace = "nebula"
# 业务层组件
NebulaSenseguardNamespace = "default"
# 基础组件
NebulaBasicComponentNamespace = "component"
# 日志组件
NebulaLogging = "logging"
# 监控组件
NebulaMonitoring = "monitoring"
# k8s组件
NebulaK8S = "kube-system"

# 远程服务器地址, 请根据实际情况修改!!!
RemoteHostname = "10.151.3.82"
# 远程服务器连接用户, 请根据实际情况修改!!!
RemoteUsername = "root"
# 远程服务器连接密码, 请根据实际情况修改!!!
RemotePassword = "Nebula123$%^"

# 为了方便操作, 请从confluence台账页面下载excel, 并保存到本地, 请根据实际情况修改!!!
OfflineExcel = "【SN-G V2.2.0】版本记录.xlsx"
# 模块名所在列号 (表格中为字母表示, 比如列号C代表第3列), 请根据实际情况修改!!!
ModuleNameCol = 3
# 版本号所在列号 (表格中为字母表示, 比如列号BH代表第60列), 请根据实际情况修改!!!
ModuleVersionCol = 60

# 非标准镜像服务, 请检查列表完整性!!!
InvalidImage = [
    "nginx-ingress",
    "kubernetes",
    "design",
    "infra-init-scripts",
    "infra-sophon-service",
    "infra-frontend-service",
    "tools",
    "mysql",
    "etcd",
    "elasticsearch"
]


def get_nebula_pod_image_list_from_k8s(namespaces):
    ssh_c = paramiko.SSHClient()
    # 允许连接不在~/.ssh/know_hosts文件中的服务器
    ssh_c.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    # 连接远程服务器
    ssh_c.connect(hostname=RemoteHostname, port=22, username=RemoteUsername, password=RemotePassword)

    image_list = {}
    for ns in namespaces:
        _, stdout, _ = ssh_c.exec_command(
            'kubectl get pods -n {} --no-headers -o custom-columns=":metadata.name"'.format(ns))
        for line in stdout.readlines():
            pod = line.strip()
            _, stdout_inner, _ = ssh_c.exec_command(
                "kubectl get pods -n {} {} -o yaml | grep image: | head -n 1".format(ns, pod))
            image, tag = "", ""
            for l in stdout_inner.readlines():
                all_parts = l.strip().split("/")
                image_part = all_parts[len(all_parts) - 1]
                image, tag = image_part.split(":")[0], image_part.split(":")[1]
                image_list[image.strip()] = tag.strip()
                break
            logger.info("[k8s - {}] Pod {}, 镜像 {}:{}".format(ns, pod, image, tag))

    # 关闭连接
    ssh_c.close()
    return image_list


# 请将表格中横线划掉的版本号删除, 再重新保存!!!
def get_nebula_pod_image_list_from_offline_excel(offline_excel):
    image_list = {}

    workbook = xlrd.open_workbook(offline_excel)
    sheets = workbook.sheet_names()

    print("共获取{}个sheet, 请选择从哪个sheet中读取信息\n".format(len(sheets)))
    for idx, sheet in enumerate(sheets):
        print("[{}] {}".format(idx, sheet))
    index = input("\n请输入序号:")
    sheet = workbook.sheet_by_index(int(index))

    n_rows, n_cols = sheet.nrows, sheet.ncols
    for col in range(n_cols)[1:]:
        image = sheet.cell_value(col, ModuleNameCol - 1)
        tag = sheet.cell_value(col, ModuleVersionCol - 1)
        if (image == "APK") or (tag == "NA") or (image in InvalidImage):
            continue
        # 比如ips或vps都存在多个版本号(t4/p4)
        if len(tag.strip().split("\n")) == 1:
            image_list[image.strip()] = [tag.strip()]
        else:
            image_list[image.strip()] = []
            for t in tag.strip().split("\n"):
                image_list[image.strip()].append(t.strip())

    return image_list


if __name__ == "__main__":
    namespaces = [
        NebulaEngineNamespace,
        NebulaSenseguardNamespace,
        NebulaBasicComponentNamespace,
        NebulaLogging,
        NebulaMonitoring,
        NebulaK8S
    ]
    image_list_from_k8s = get_nebula_pod_image_list_from_k8s(namespaces)
    image_list_from_excel = get_nebula_pod_image_list_from_offline_excel(OfflineExcel)

    # k8s上的镜像版本号得去满足confluence台账页面上的版本号
    for image, tags in image_list_from_excel.items():
        found = False
        found_tag = ""
        for tag in tags:
            if image_list_from_k8s[image] == tag:
                found = True
                found_tag = tag
                break
        if found:
            logger.info("镜像{}匹配通过, 版本号为{}".format(image, found_tag))
        else:
            logger.error("镜像{}匹配失败, 请检查机器上的版本号".format(image))
