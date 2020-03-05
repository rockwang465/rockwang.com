from datetime import timedelta
from datetime import datetime
from minio_delete import *
from psutil import disk_usage
import socket

format_str = "%Y-%m-%d %H:%M:%S"


# 获取本机ip
def get_ip():
    myname = socket.getfqdn(socket.gethostname())
    myaddr = socket.gethostbyname(myname)
    return myaddr


# 获取15天阈值
def time_threshold():
    delay_time = datetime.now() - timedelta(days=15)
    time = delay_time.strftime(format_str)
    # datetime.datetime.strptime(delay_time)
    # print(time)
    # print(delay_time)
    return time


# 检查minio 挂载目录
def disk():
    path = 'D:/'
    total, used, free, percent = disk_usage(path)
    return percent


def analysis_object(bucket_name, obj_list):
    # del_obj_list = {}
    # obj = sorted(list, key=lambda item: item['object_last_modified'])
    # print(obj)
    # global format_str
    del_obj = []
    threshold_time = time_threshold()
    # 判断天数
    for item in obj_list.get(bucket_name):
        obj_time = item['object_last_modified'].strftime(format_str)
        # print(obj_time)
        if obj_time < threshold_time:
            del_obj.append(item.get('object_name'))
    # del_obj_list[bucket_name] = del_obj
    return bucket_name, del_obj


def analysis_old_object(bucket_name, obj_list):
    del_old_obj = []
    all_obj = obj_list.get(bucket_name)
    obj = sorted(all_obj, key=lambda item: item['object_last_modified'], reverse=True)
    print(obj[:1])
    for item in obj[:1]:
        del_old_obj.append(item['object_name'])
    return bucket_name, del_old_obj


def start(object, bucket):
    disk_percent = disk()
    obj_list = object.list_objects(bucket)
    # print(analysis_old_object('senseguard-map-management', obj_list))
    # 未满95%，删除超过15天的数据
    if disk_percent < 95:
        del_bucket, obj = analysis_object(bucket, obj_list)
        print(del_bucket, obj)
    else:
        # 满95%，删除最老的，直到90%
        while true:
            del_old_bucket, old_obj = analysis_old_object(bucket, obj_list)
            # del_objects(del_old_bucket,old_obj)
            if disk_percent < 90:
                break


if __name__ == '__main__':
    ip = get_ip()
    minio = minio_delete('10.151.3.94:31900', 'minio', 'minio123')
    buckets = minio.list_minio_bucket()
    # print(buckets)
    for item in buckets:
        start(minio, item)

    # minio.del_objects()
