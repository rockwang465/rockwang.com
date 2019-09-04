class Lazyproperty:
    def __init__(self, func):
        print('Lazyproperty=====>', func)
        self.func = func

    def __get__(self, instance, owner):
        print("-->__get__方法")
        # 5.2 这里 self.func就是area，instance为Room实例化的r1，owner为Room自身
        res = self.func(instance)  # 5.3 则执行了area(self)这个函数
        return res  # 5.4 返回执行的结果


class Room:
    def __init__(self, name, width, length):
        self.name = name
        self.width = width
        self.length = length

    # @property  # 静态方法，就是用 .area 就执行了area函数，并拿到了返回值
    @Lazyproperty  # 相当于执行: area = property(area)
    def area(self):
        return self.width * self.length


r1 = Room('厨房', 3.2, 2.5)
# print(r1.area)  # 1.使用@property ，结果为8.0
# print(r1.__dict__)  # 2.1 结果为: {'name': '厨房', 'width': 3.2, 'length': 2.5}, 里面是没有area的
# print(Room.__dict__)  # 2.2 结果中有 :   'area': <__main__.Lazyproperty object at 0x00530EF0>
# print(r1.area)  # 3. 使用@Lazyproperty ，结果为描述符 <__main__.Lazyproperty object at 0x00530EF0>，所以根本没有触发area这个函数的运行
                  # 因为area里的函数没有执行，而是执行的@Lazyproperty，即area=Lazyproperty(area)
                  # 因为area代理了下面函数area，而area现在为非数据描述符，r1.area中只能找到这个非数据描述符
# print(r1.area.func(r1))  # 4. 现在只能这样触发，这样肯定不行。

print(r1.area)  # 5.1 由于r1.area只能找到代理的area非数据描述符方法，则触发了__get__方法，
