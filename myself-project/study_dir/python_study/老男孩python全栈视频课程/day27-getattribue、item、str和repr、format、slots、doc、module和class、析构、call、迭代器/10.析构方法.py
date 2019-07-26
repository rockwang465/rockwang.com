class Foo:
    def __init__(self, name):
        self.name = name

    def __del__(self):
        print("析构方法执行啦")


f1 = Foo('rock')
del f1.name
print("----------")
