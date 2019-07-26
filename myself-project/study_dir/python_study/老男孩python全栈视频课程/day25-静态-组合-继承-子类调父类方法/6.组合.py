class School:
    def __init__(self, name, addr):
        self.name = name
        self.addr = addr


class Course:
    def __init__(self, name, price, period, school):
        self.name = name
        self.price = price
        self.period = period
        self.school = school.name  # 拿到学校的名字
        self.addr = school.addr  #拿到校区地址
    def print_info(self):
        print("欢迎选择 %s %s校区, 周期 %s, 价格 %s" % (self.school, self.addr, self.period, self.price))

s1 = School('oldboy', '北京')
s2 = School('oldboy', '南京')
s3 = School('oldboy', '东京')

# 情况1： s1为上面的传入，放到Course类中传参进去，而且是把整个s1字典都传进去了
# c1 = Course('linux', 8000, '3month', s1)
# print(c1.__dict__)
# print(c1.school)

# 情况2： 当学校这里有3个，北京 南京 东京，为了方便用户选择，所以进行input交互更合适：
msg = """
1 老男孩 北京校区
2 老男孩 南京校区
3 老男孩 东京校区 
"""
while True:
    print(msg)
    menu = {
        '1': s1,
        '2': s2,
        '3': s3
    }
    choice = input("请输入对应的学校校区>>>:")
    c = Course('linux', 8000, '3month', menu[choice])
    c.print_info()
