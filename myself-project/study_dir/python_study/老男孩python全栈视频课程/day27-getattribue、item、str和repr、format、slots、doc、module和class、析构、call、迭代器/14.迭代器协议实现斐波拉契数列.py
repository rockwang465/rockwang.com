class Fib:
    def __init__(self):
        self.a = 1
        self.b = 1

    def __iter__(self):
        return self

    def __next__(self):
        self.a, self.b = self.b, self.a + self.b
        if self.b >= 30:
            raise StopIteration("大于30结束")
        return self.b


f1 = Fib()
print(next(f1))
print(next(f1))
print(next(f1))
print("+++++++++++++++++++++++++++++")
for i in f1:
    print(i)
