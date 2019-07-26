class Foo:
    def __call__(self, *args, **kwargs):
        print("实例执行啦 obj()")


f1 = Foo()
f1()  # 执行是调用Foo下的__call__方法
# Foo()  # 执行是xxxxx下的__call__方法
