#!/usr/bin/env python
# -*- coding: utf-8 -*-

import subprocess,time,requests,json

ip="10.5.6.66"
port="30080"
##ns=["addons","component","default","helm","ingress","kube-public","kube-system","logging","monitoring","mysql-operator"]
topic=["stream.features.automobile","stream.features.automobile.garbage","stream.features.cyclist","stream.features.cyclist.garbage","stream.features.face_24602","stream.features.face_24602.garbage","stream.features.pedestrian","stream.features.pedestrian.garbage","stream.rws.td.comparison","stream.sensekeeper.biz","stream.sensekeeper.rwstsdb","sync.stream.features.face_24602"]
bucket=["keeper_face","video_automobile_cropped","video_automobile_panoramic","video_cyclist_cropped","video_cyclist_panoramic","video_face","video_panoramic","video_pedestrian_cropped","video_pedestrian_panoramic"]


def check_server_usage():
  print("您好: 您当前测试的机器为:",ip)

def check_node_state():
  state=subprocess.run('kubectl get nodes',stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
  print("\n1. 检查k8s node状态".ljust(80,'*'))
  check_node_value=(str(state.stdout, encoding='utf-8'))
#  print(check_node_value)
  if 'NotReady' in check_node_value.split():
    print(check_node_value)
    print("Error : node is [ NotReady ]")
    time.sleep(4)
  else:
    print("Node正常")
    time.sleep(1)
  print("\n".rjust(84,'*'))


def check_pods_state():
  state=subprocess.run('kubectl get pods --all-namespaces | grep -v Running | grep -v NAME',stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
  print("\n2. 检查k8s所有pod服务状态".ljust(76,'*'))
  if state.stdout:
    check_pods_value=(str(state.stdout, encoding='utf-8'))
    print(check_pods_value)
    print("Error : pods有问题")
    time.sleep(4)
  else:
    print("Pod正常")
    time.sleep(1)
  print("\n".ljust(84,'*'))


def check_license_state():
  state=subprocess.run("/usr/local/bin/license_client status | grep 'status is' | awk -F: '{print $2}'",stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
  check_license_value=(str(state.stdout, encoding='utf-8'))
  print("\n3. 检查加密狗状态".ljust(77,'*'))
#  print(check_license_value)
  if 'alive' not in check_license_value.split():
    print(check_license_value)
    print("Error : 加密狗 has error")
    time.sleep(4)
  else: 
    print("加密狗正常")
    time.sleep(1)
  print("\n".rjust(84,'*'))

def check_topic_state():
  state=subprocess.run("kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --list --zookeeper zookeeper-default:2181/kafka | grep -v 'consumer_offsets'",stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
  ##state=subprocess.run("kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --list --zookeeper zookeeper-default:2181/kafka | egrep -v 'consumer_offsets|auto'",stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
  check_topic_value=(str(state.stdout, encoding='utf-8'))
  print("\n4. 检查topic状态".ljust(80,'*'))
#  print(check_topic_value)
  topic_num=0
  for i in topic:
    if i not in check_topic_value.split():
      print("Error : 缺少topic,请创建 : ",i)
      topic_num=topic_num+1
  if topic_num == 0:
    print("topic正常")
    time.sleep(1)
  else:
    print("\n")
    print(check_topic_value)
    time.sleep(4)
  print("\n".rjust(84,'*'))


def check_bucket_state():
  url='http://'+ip+':'+port+'/components/osg-default/v1'
  req = requests.get(url)
  print("\n5. 检查bucket状态".ljust(80,'*'))
  req_value=eval(req.text)   #转成字典
#  print(req_value)
#  print(type(req_value))
  if 'buckets' in req_value:
    url_bucket=[]
    for k1 in req_value['buckets']:
      if 'name' in k1:
        ##print(k1['name'])
        url_bucket.append(k1['name'])
        ##print(url_bucket)
      else:
        print("Error : 无任何buckets，请全部创建")
  else:
    print("Error : 无任何buckets，请全部创建")
  
  bucket_num=0
  for i in bucket:
    if i not in url_bucket:
      print("Error : 缺少bucket,请创建 : ",i)
      bucket_num=bucket_num+1
      time.sleep(1)
  if bucket_num == 0:
    print("bucket正常")
    time.sleep(1)
  print("\n".rjust(84,'*'))


check_server_usage()
check_node_state()
check_pods_state()
check_license_state()
check_topic_state()
check_bucket_state()
