class People:
    _star = 'earth'
    __country = 'china'
    def __init__(self,id ,name,age,salary):
        self.id = id
        self.name = name
        self.age = age
        self.salary = salary

    def get_id(self):
        print('我是私有方法啊，我找到的id是 [%s]' % self.id)
