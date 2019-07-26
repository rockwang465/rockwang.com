class School:
    def __init__(self, name, addr):
        self.name = name
        self.addr = addr


class Course:
    def __init__(self, name, price, period, school):
        self.name = name
        self.price = price
        self.period = period
        self.school = school

school_choice = """1. 北京校区
2. 南京校区
3. 东京校区
"""

# s1 = School('oldboy', '北京校区')
# s2 = School('oldboy', '南京校区')
# s3 = School('oldboy', '东京校区')

course_choice = """1. Python课程
2. Go课程
3. Linux课程
"""

price_choice = """1. 15000 精英班
2. 11000 高级班
3. 6999 网络班
"""

period_choice = """1. 3个月 线下班
2. 6个月 周末班
3. 6个月 线上班
"""

while True:
    for i in school_choice.split("\n"):
        print(i)
    school_input = input("请输入校区>>>")
    print(course_choice)
    course_input = input("请输入课程>>>")
    print(price_choice)
    price_input = input("请输入价格>>>")
    print(period_choice)
    period_input = input("请输入周期>>>")
#
#     Course()

