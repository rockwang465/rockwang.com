class School:
    def __init__(self, name, addr, ):
        self.name = name
        self.addr = addr


class Course:
    def __init__(self, name, price, period, school):
        self.name = name
        self.price = price
        self.period = period


s1 = School('oldboy', '北京')
s2 = School('oldboy', '南京')
s3 = School('oldboy', '东京')

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
    choice_school = input("请输入对应的学校校区>>>:")
    choice_course = input("请输入课程>>>:")
    choice_price = input("请输入价格>>>:")
    choice_period = input("请输入周期>>>:")
    # c = Course('linux', 8000, '3month', menu[choice_school])
    c = Course(choice_course, choice_price, choice_period, menu[choice_school])
    print("欢迎选择 %s %s校区 的 %s 课程，周期为 %s, 价格 %s" % (menu[choice_school].name, menu[choice_school].addr, c.name, c.period, c.price ))
