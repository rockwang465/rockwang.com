#!/usr/bin/env bash
git_branch=`git branch| awk '{print $NF}'`  # master
commit_id=`git rev-parse --short HEAD`  # f6e8886
SERVER_NAME="senseguard-td-result-consume-sxyc2"
CHART_DIR="charts/senseguard-td-result-consume-sxyc2"
GUARD_VERSION="2.2.0"

DOCKER_REGISTRY="10.151.3.75/sensenebula-guard-std"
CHARTS_REGISTRY="http://10.151.3.75:8080/api/charts"

# 生成docker镜像并推送到仓库
# image镜像名称示例: 10.151.3.75/sensenebula-guard-std/senseguard-td-result-consume:2.1.0-v2.1.0-e6fff6
IMAGE_TAG="${SERVER_NAME}:${GUARD_VERSION}-${git_branch}-${commit_id}"
IMAGE_NAME="${DOCKER_REGISTRY}/${IMAGE_TAG}"
#echo ${IMAGE_NAME}
docker build -t ${IMAGE_NAME} .
docker push ${IMAGE_NAME}

# charts名称示例: senseguard-td-result-consume-2.2.0-v2.2.0-5a5816.tgz
CHART_TAG="${GUARD_VERSION}-${git_branch}-${commit_id}"
CHART_NAME="${SERVER_NAME}-${CHART_TAG}"
echo ${CHART_NAME}  # 效果: senseguard-td-result-consume-2.2.0-master-f6e8886

# 打包chart
cd ${CHART_DIR}
helm package --version=${CHART_TAG} . | echo true

# 推送chart到仓库
#echo "@${CHART_NAME}.tgz"
curl --data-binary "@${CHART_NAME}.tgz" ${CHARTS_REGISTRY}

# 删除本地chart文件
rm -f ${CHART_NAME}.tgz
cd ..

# 打印image和chart信息
echo -e "\n"
echo -e "##########################################################################################"
echo "docker image: ${IMAGE_NAME}"
echo "helm chart:   ${CHART_NAME}.tgz"
echo -e "##########################################################################################"