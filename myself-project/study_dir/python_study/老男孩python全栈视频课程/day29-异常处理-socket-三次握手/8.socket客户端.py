import socket

phone = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

phone.connect(('127.0.0.1', 8000))  # 拨通电话

phone.send('hello'.encode('utf-8'))  # 必须用二进制bytes才可以
data = phone.recv(1024)
print('收到服务端的发来的消息: ', data)

phone.close()