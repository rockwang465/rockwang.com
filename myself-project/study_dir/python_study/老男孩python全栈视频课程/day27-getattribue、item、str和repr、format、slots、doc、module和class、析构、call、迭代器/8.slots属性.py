class Foo:
    # __slots__ = ['name', 'age']  # 相当于定义了一个字典，替换了__dict__. {'name':None,'age': None}
    __slots__ = 'name'  # 这里代表只能用 name 这一个属性，实例化后定义其他属性就报错。


f1 = Foo()
f1.name = 'rock'
print(f1.name)
print(f1.__slots__)
# print(f1.__dict__)  # dict现在永不了了，因为有了slots了。
f1.job = 17  # 没有job时，则不允许定义此用法的。
