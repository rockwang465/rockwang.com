# _*_ coding:utf-8 _*_

import logging

logging.basicConfig(filename="log-test.log",
                    level=logging.INFO,
                    format='%(asctime)s - %(levelname)s : %(filename)s:%(funcName)s:%(lineno)d : %(message)s',
                    datefmt='%Y-%m-%d %I:%M:%S %p')


def say_err():
    logging.error("There are has error")


say_err()

logging.debug("This message is debug")
logging.info("This message is info")
logging.warning("This message is warning")
logging.critical("This message is critical\n")
