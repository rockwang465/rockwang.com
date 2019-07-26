class Open:
    def __init__(self, name):
        self.name = name

    def __enter__(self):
        print("执行了enter")
        return self  # 返回值将赋值给f

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("执行了exit")
        print(exc_type)
        print(exc_val)
        print(exc_tb)
        return True


# with Open as f 有点类似于实例化: f = Open()，但实际上仅实例化相同，而触发过程是不同的，是__enter__返回值给f的。
with Open('a.txt') as f:
    print(f)
    print(f.name)
    print(bucunzaidezhi)  # 操作一个不存在的值，报错NameError
    print("------->")

print('end+++++++++++++++')

# 总结:
# 1. with Open ---> 会触发 Open.__enter__(),拿到返回值。
# 2. as f ---> 等同于 f = 返回值
# 3. with Open as f 等同于 f = Open.__enter__()
# 4. 执行代码块过程:
#    (__exit__中的三个值: exc_type, exc_val, exc_tb)
#    A. with Open中代码无异常时，先执行__enter__ ; 再执行with Open下面的代码 ; 然后执行__exit__,它的三个参数都为None.
#    B. with Open中代码有异常时:
#       则程序触发 __exit__ 模块 , 且with Open后面的代码就不会再执行了(代表了with Open内的代码整个运行结束了)。
#       exc_type提示: <class 'NameError'> , 表示报错类型class
#       exc_val提示: name 'bucunzaidezhi' is not defined , 表示报错的值
#       exc_tb提示: <traceback object at 0x00C14A80> , 表示报错的追踪信息
#       I.  如果 __exit__ 中有 return True， 则把所有异常信息吞了，只打印 exc_type, exc_val, exc_tb 的信息，并执行最后的 print('end+++')。
#      II.  如果 __exit__ 中没有 return True， 则把所有异常信息都吐出来了，会显示报错信息。后面的任何代码都不会再执行了，包括print('end+++')。
#     III.  __exit__ 运行完毕后，代表这整个with语句的所有代码执行完毕(上面也说了)。


