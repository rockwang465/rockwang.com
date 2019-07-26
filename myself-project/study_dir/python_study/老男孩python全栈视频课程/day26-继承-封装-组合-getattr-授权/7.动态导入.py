# from import_dynamic import *
# test1()
# _test2()

import importlib
m = importlib.import_module('import_dynamic')
m._test2()  # 这样是可以正常调用的