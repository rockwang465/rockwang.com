class A:
    def test(self):
        print("A")

class B(A):
    pass
    # def test(self):
    #     print("B")

class C(A):
    # def test(self):
    #     print("C")
    pass

class D(B):
    pass
    # def test(self):
    #     print("D")

class E(C):
    pass
    # def test(self):
    #     print("E")

class F(D, E):
    pass
    # def test(self):
    #     print("F")

f1 = F()
# f1.test()  #F类中有test函数时，结果为F
# f1.test()  #F类中没有test函数时，结果为D
# f1.test()  #F类、D类都没有test函数时，结果为B
# f1.test()  #F类、D类、B类都没有test函数时，查找顺序改变，切换到右侧查找，结果为E
# f1.test()  #左侧F类、D类、B类，以及右侧E类都没有test函数时，结果为C
# f1.test()  #左侧F类、D类、B类，以及右侧E类 、C类都没有test函数时，结果为A

# 总结上面结果: 左侧:F->D->B-->右侧:E->C->A

print(F.__mro__)

# 继承顺序:
# 当类为经典类时，多继承按深度优先
# 当类为新式类时，多继承按广度(从左侧开始找)优先（python3为新式类）
