# 例1 : deco的用法
# def deco(obj):
#     obj.x = 1  # 定义了固定的值 x y z
#     obj.y = 2
#     obj.z = 3
#     return obj  # 返回Foo、Bar
#
# @deco  # 相当于 Foo = deco(Foo)
# class Foo:
#     pass
#
# @deco
# class Bar:
#     pass
#
# print(Foo.x)


# 例2 : 解决传输固定值的问题
def out(**kwargs):
    def deco(obj):
        # obj.x = 1  # 定义了固定的值 x y z
        print('inner-->', kwargs)
        # print('obj-->', obj, obj.__dict__)
        for k, v in kwargs.items():  # 循环kwargs这个字典
            # print(k, v)
            # obj.__dict__[k] = v  # 设置值，报错: TypeError: 'mappingproxy' object does not support item assignment
            # obj.k = v  # 这样不行，这是加了一个key属性 'k': 3 ， 导致k不是变量，而是字符串了
            setattr(obj, k, v)  # 这样才行
        print('结果-->', obj.__dict__)
        return obj  # 返回Foo、Bar

    print('==>', kwargs)
    return deco


@out(x=1, y=2, z=3)  # 首先out()是先执行这个函数，然后会返回deco函数名，则现在out()的值等于deco。
# 再执行@deco语法糖(Foo=deco(Foo) )
class Foo:
    pass


@out(name='egon', age=18)
class Bar:
    pass

print(Bar.name)