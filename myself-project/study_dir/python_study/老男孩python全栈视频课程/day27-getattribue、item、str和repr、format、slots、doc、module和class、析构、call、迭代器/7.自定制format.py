# class Format:
#     def __init__(self, year, month, day):
#         self.year = year
#         self.month = month
#         self.day = day

# f1 = Format('2019', '07', '18')
#
# data = '{0}'.format(f1)
# data2 = '{0.year} {0.month} {0.day}'.format(f1)
# print(data)
# print(data2)


format_dic = {
    'ymd': '{0.year} {0.month} {0.day}',
    'y:m:d': '{0.year}:{0.month}:{0.day}',
    'y-m-d': '{0.year}-{0.month}-{0.day}'
}


class Format:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    def __format__(self, format_spec):
        if not format_spec or format_spec not in format_dic:
            format_spec = "ymd"  # 赋予默认值
        fm = format_dic[format_spec]
        return fm.format(self)  # 必须要有返回值


f1 = Format('2019', '07', '18')
# print(f1.__format__('ymd'))  # 普通调用方式
# print(f1.__format__('y:m:d'))
# print(f1.__format__('y-m-d'))

print(format(f1, 'y:m:d'))  # 老师是这样写的
print(format(f1, 'aaaaa'))  # 给一个不存在的值
