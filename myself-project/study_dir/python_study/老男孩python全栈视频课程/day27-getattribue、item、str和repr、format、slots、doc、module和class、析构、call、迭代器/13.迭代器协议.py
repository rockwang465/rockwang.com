class Foo:
    def __init__(self, num):
        self.num = num

    def __iter__(self):
        return self

    def __next__(self):
        if self.num == 15:
            raise StopIteration("到%s结束了" % self.num) # 捕捉到这个错误则结束
        self.num += 1
        return self.num


f1 = Foo(10)
# print(f1.__next__())
# print(f1.__next__())
# print(f1.__next__())
# print(f1.__next__())

for i in f1:
    print(i)


# 输出结果:
# 11
# 12
# 13
# 14
# 15