class area:
    def __init__(self, name, addr, width, length):
        self.name = name
        self.addr = addr
        self.__width = width
        self.__length = length  # 这里不想给别人调用宽度和长度，用双下划线

    def calc(self):  # calc作为接口给别人使用
        # 下面自己引用长度、宽度
        print("%s 家的%s 的面积为: %s 平方米" % (self.name, self.addr, self.__length * self.__width))

    def width2(self):  # 但是以后可能需要用宽度、长度的时候，只能做成接口给别人使用了
        return self.__width

a1 = area('alex', '厕所', 5, 8)
a1.calc()
