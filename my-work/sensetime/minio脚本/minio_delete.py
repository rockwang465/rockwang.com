from minio import Minio
from minio.error import ResponseError
import minio
import os
from datetime import datetime
import time


# import urllib3
class minio_delete():

    def __init__(self, url, access_key, secret_key):
        self.url = url
        self.access_key = access_key
        self.secret_key = secret_key

    def connect(self):
        minioClient = Minio(self.url, self.access_key, self.secret_key, secure=False)
        return minioClient

    def list_minio_bucket(self):
        bucket_list = []
        minioClient = self.connect()
        buckets = minioClient.list_buckets()
        for bucket in buckets:
            bucket_list.append(bucket.name)
        return bucket_list

    def remove_bucket(self, bucketname):
        try:
            minioClient.remove_bucket(bucketname)
        except ResponseError as err:
            print(err)

    def list_objects(self, bucketname):
        bucket_dict = {}
        obj_name = []
        # format_str = "%Y-%m-%d %H:%M:%S"
        minioClient = self.connect()
        objects = minioClient.list_objects_v2(bucketname, prefix=None, recursive=True)
        for obj in objects:
            # print(obj.bucket_name, obj.object_name.encode('utf-8').decode('utf-8'), obj.last_modified,obj.etag, obj.size, obj.content_type)
            object_name = obj.object_name.encode('utf-8').decode('utf-8')
            if obj.is_dir:
                continue
            # print(obj.last_modified)
            obj_name.append(
                {
                    # 'bucket_name': obj.bucket_name,
                    'object_name': object_name,
                    # 'is_dir': obj.is_dir,
                    'object_last_modified': obj.last_modified
                    # 'object_last_modified': datetime.strptime(obj.last_modified ,'%Y-%m-%d %H:%M:%S')
                    # 'object_last_modified':  time.strftime(format_str,time.localtime(obj.last_modified))
                }
            )
        bucket_dict[bucketname] = obj_name
        # print(obj_name)
        return bucket_dict

    def del_objects(self, bucketname, objects_to_delete):
        try:
            # force evaluation of the remove_objects() call by iterating over
            # the returned value.
            for del_err in minioClient.remove_objects(bucketname, objects_to_delete):
                print("Deletion Error: {}".format(del_err))
        except ResponseError as err:
            print(err)

    def minio_del(self):
        pass
