import time


class FileHandle:
    def __init__(self, filename, mode="r", encoding="utf-8"):
        self.file = open(filename, mode, encoding=encoding)  # 这里定义为系统的open方法打开文件
        self.mode = mode
        self.encoding = encoding

    def write(self, line):  # 要求: 在写入文件时，前面加上时间。 -->替换系统的write()方法
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        self.file.write("%s : %s" % (now, line))

    def __getattr__(self, item):
        # self.file.read()  #正常读取一个文件，用系统方法，是后面加read()
        # 但是用了 __getattr__ ，item此时是字符串，我们可以用通过getattr来调用字符串
        # print(item, type(item))  #type(item)为str，例如 "read" "write"等
        return getattr(self.file, item)  # 通过getattr拿到self.file的item方法，并返回


f1 = FileHandle('text.txt', 'w+')
f1.write('---------->\n')  # 当此处调用write、read等方法时，由于FileHandle中没有write、read，
f1.write('CPU负载过高\n')
f1.write('内存剩余不足\n')
f1.write('磁盘空间不足\n')
# 则会执行__getattr__中的代码。
f1.seek(0)  # 移到0位置
data = f1.read()
print(data)

# 此组合授权优点：
#   1.f1.file.read() 调用时，显示检查当前class中是否有read函数方法，没有才走__getattr__
#   2.f1.file.read() 调用时，read是以字符串方法上传到__getattr__中的item上的，而 getattr()中的item正好来检测是否有这个字符串对应的调用方法。
#   3.f1.file.write() 【牛逼的地方】--->调用时，是使用定义的函数wirte，虽然里面还是系统的write方法，但增加了功能。 增加了时间戳在写入的数据中。
#   总结: 其实完全是靠 __getattr__ 作为中转。
