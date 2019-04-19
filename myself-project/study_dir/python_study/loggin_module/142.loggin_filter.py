# _*_ coding: utf-8 _*_

# logging 高级玩法
# python使用logging模块记录日志涉及4个主要类:
# logger 提供了应用程序可以直接使用的接口。 (主要调用的程序)
# handler 讲logger创建的日志记录发送到合适的目录输出。(配置屏幕输出、或日志文件输出目的地)
# filter 提供了对输出的日志进行过滤。(对应什么级别的日志，能否输出)
# formatter 决定日志记录的最终输出格式。 (定义输出日志的格式)

import logging

class IgnoreBackupLogFilter(logging.Filter):
    """忽略带db backup 内容的日志打印"""
    def fitler(self, record):  #固定写法
        return "db backup" not in record.getMessage()
        # return返回的结果是True 或False，判断 "db backup" 是否在打印的日志中


# 1.1 生成 logger 对象
logger = logging.getLogger("web")
logger.setLevel(logging.DEBUG)  # 设置全局日志级别，如果不设置默认为Warning级别。

# 1.2 把filter 对象添加到logger中
logger.addFilter(IgnoreBackupLogFilter())

# 2.1 生成 handler 对象(设置日志信息输出的地方: 屏幕、日志文件)
ch = logging.StreamHandler()  # 这是往屏幕输出的配置
fh = logging.FileHandler("nginx_web.log")  # 这是往日志文件输出的配置

# 2.2 设置当前日志级别
# ch.setLevel(logging.DEBUG)  # 设置输出到屏幕的是INFO级别
# fh.setLevel(logging.WARNING)  # 设置输出到日志文件的是WARNING级别

# 2.3 把handler对象 绑定到logger 上
logger.addHandler(ch)  # 绑定屏幕
logger.addHandler(fh)  # 绑定日志文件

# 3.1 生成 formatter 对象(打印日志的格式配置)
file_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s ')
console_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(lineno)d - %(message)s ')

# 3.2  把 formatter 对象绑定到 handler 对象上
ch.setFormatter(console_formatter)
fh.setFormatter(file_formatter)

# 这里是测试打印日志的地方（此处受全局日志级别、当前日志级别的影响）
logger.debug("test debug log")  # 打印debug级别日志
logger.debug("test debug db backup log")  # 打印debug级别日志
logger.info("test info log")  # 打印info级别日志
logger.warning("test warning log")  # 打印warning级别日志
logger.error("test error log and db backup")  # 打印error级别日志
logger.critical("test critical log\n")  # 打印critical级别日志
# 全局定义INFO，屏幕是DEBUG,日志文件是WARNING,所以最后输出的结果为：
#    屏幕输出INFO以上级别的日志(因为全局高于定义的屏幕级别)
#    日志文件输出WARNING以上级别的日志(因为全局低于定义的日志文件级别)

# 注意:
# A. 定义全局 日志输出级别 的代码为:  logger.setLevel(logging.INFO)
# B. 定义当前屏幕/日志文件 日志输出级别 的代码为：
#     ch.setLevel(logging.DEBUG)    #设置输出到屏幕的是INFO级别
#     fh.setLevel(logging.WARNING)  #设置输出到日志文件的是WARNING级别
# C. 最后打印日志信息 的代码为：
#     logger.debug("test debug log\n")  #打印debug级别日志
#     logger.error("test error log\n")  #打印debug级别日志

# 总结：
#   A. 全局日志级别没有设置时，默认为warning级别。所以低于WARNING级别的日志都不会被打印出来的。
#   B. 当全局日志级别高于当前定义的日志级别时，按照全局的来。
#      例如: 全局为Warning，当前屏幕输出为INFO,则INFO级别的日志不会打印到屏幕。
#   C. 当全局日志级别低于当前定义的日志级别时，按当前定义的来。
#      例如：全局为DEBUG，当前屏幕输出为WARNING，则低于WARNING级别的日志是不会打印到屏幕的。
