class BlackMedium:
    feture = 'Ugly'

    def __init__(self, name, addr):
        self.name = name
        self.addr = addr

    def sell_house(self):
        print('【%s】 正在卖房子，傻逼才买呢' % self.name)

    def rent_house(self):
        print('【%s】 正在租房子，傻逼才租呢' % self.name)


# print(hasattr(BlackMedium, 'feture'))
# getattr()

b1 = BlackMedium('万成置地', '天露园')
# b1.name--->b1.__dic__['name']
# print(b1.__dict__)

# b1.name
# b1.sell_house
# print(hasattr(b1, 'name'))  # 查看是否有这个属性，有则True，没有则False
# print(hasattr(b1, 'sell_house'))
# print(hasattr(b1, 'sell_houseasdfsa'))

# print(getattr(b1, 'name'))  # 获取属性的值
# print(getattr(b1, 'rent_house'))  # 获取函数的地址
# func = getattr(b1, 'rent_house')  # 这样也可以用来执行
# func()  # 执行这个函数

# print(getattr(b1, 'rent_houseasdfsa'))  # 没有则报错
# print(getattr(b1, 'rent_houseasdfsa', '没有这个属性'))  # 没有的话，则返回后面的默认属性 '没有这个属性'

# b1.sb=True  # 普通的设置属性的方法
# setattr(b1, 'sb', True)  # setattr设置增加新属性的方法
# setattr(b1,'name','SB')  # 修改属性值
# setattr(b1, 'func', lambda x: x + 1)  # 设置函数属性
# setattr(b1, 'func1', lambda self:self.name+'sb')  # 函数属性中调用自身的self属性
# print(b1.__dict__)  # 字典中查看是否有设置的属性

# print(b1.func(10))  # 打印这个函数
# print(b1.func1(b1))
# del b1.sb  # 普通删除属性值的方法
# delattr(b1,'sb')  #  delattr删除对应的属性值
print(b1.__dict__)
