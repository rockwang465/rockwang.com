class Typed:
    def __init__(self, key, expected_type):
        self.key = key
        self.expected_type = expected_type

    def __get__(self, instance, owner):
        print('get方法')
        # print('instance参数【%s】' %instance)
        # print('owner参数【%s】' %owner)
        return instance.__dict__[self.key]

    def __set__(self, instance, value):
        print('set方法')
        # print('instance参数【%s】' % instance)
        # print('value参数【%s】' % value)
        # print('====>',self)
        if not isinstance(value, self.expected_type):  # 5. 判断传入的实参是否符合对应的type类型，如名字必须为str类型的值。
            # print('你传入的类型不是字符串，错误')
            # return
            raise TypeError('%s 传入的类型不是%s' % (self.key, self.expected_type))
        instance.__dict__[self.key] = value

    def __delete__(self, instance):
        print('delete方法')
        # print('instance参数【%s】' % instance)
        instance.__dict__.pop(self.key)


# 3. 定义deco进行传参操作
def deco(**kwargs):  # kwargs={'name':str,'age':int}
    def wrapper(obj):  # obj=People
        for key, val in kwargs.items():  # (('name',str),('age',int))
            # 4. [难点] setattr 设置属性值到People这个class中，key为输入的name、age等，value为对应的描述符；其实就是在做  People.name=Typed('name',str)这种设置属性值
            # 而  People.name=Typed('name',str) 就是为了解决 People中定义的 name=Typed('name',str) 的操作
            setattr(obj, key, Typed(key, val))  # Typed(key, val)的值为<__main__.Typed object at 0x02D11DB0> 这样的描述符， 而它的字典__dict__为{'key': 'name', 'expected_type': <class 'str'>}
            # 解释为: setattr(People,'name',Typed('name',str)) # People.name=Typed('name',str)
        return obj

    return wrapper


# 2. @deco中设置传入的对应值必须为对应的类型，如name为str，age为int
@deco(name=str, age=int, salary=float, gender=str, height=str)  # 先执行deco() ==> 后执行 @wrapper ，相当于执行 People=wrapper(People)
class People:
    # 1.原来是定义这里，一个一个传入值。 后面优化用@deco传入值，简化了代码(29-43行代码就是为了优化下面几行代码的定义)
    # name=Typed('name',str)
    # age=Typed('age',int)
    # salary=Typed('salary',int)

    def __init__(self, name, age, salary, gender, height):
        self.name = name
        self.age = age
        self.salary = salary
        self.gender = gender
        self.height = height


# p1=People('213',13.3,13.3,'x','y')
p2 = People('rock', 28, 16.0, 'man', '173cm')  # 6. 实例化
print(People.__dict__)  # 打印People中的字典是否有对应的key和描述符(<__main__.Typed object at 0x00F0BF90>)
