class Traffic:  # 父类
    def __init__(self, name, speed, load, energy):
        self.name = name
        self.speed = speed
        self.load = load
        self.energy = energy

    def run(self):
        print("%s 号线地铁 开动啦" % self.line)


class Subway(Traffic):  # 子类调用父类
    def __init__(self, name, speed, load, energy, line):  # 这里要传入所有参数
        # Traffic.__init__(self, name, speed, load, energy)  # 这里引用父类的变量参数，下面就不用一个一个self.xxx = xxx重复定义了，直接用父类的定义了
        # super().__init__(name, speed, load, energy)  # 用super()替代父类的继承
        super(Subway, self).__init__(name, speed, load, energy)  # 这样写和上面super()其实是一模一样的
        self.line = line

    def show_info(self):
        print("%s %s号线地铁，当前时速: %s ,当前承载量: %s ,使用能用: %s" % (self.name, self.line, self.speed, self.load, self.energy))

    def run(self):  # 引用父组件的run方法
        # Traffic.run(self)
        super()


line13 = Subway('上海地铁', '300km/h', '2050人', 'electricity', '17')
line13.show_info()
line13.run()
