import socket

# 1 实例化一个phone
# 1.1 socket.AF_INET 表示基于网络通信, socket.SOCK_STREAM 表示基于tcp协议通信，详细含义后期解释STREAM
phone = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 2 绑定ip + port,以元祖方式绑定
phone.bind(('127.0.0.1', 8000))

# 3 表示可以同时可以接入几个连接
phone.listen(5)
print('---->')  # 确认listen没有卡

# 4 拿到元祖中第一个值(链接) 和第二个值(地址)
conn, addr = phone.accept()  # 等电话

# 5 可以接受多少字节的信息, msg为接受到的信息
msg = conn.recv(1024)  # 收消息
print('客户端发来的消息是: ', msg)

# 6 发送消息
conn.send(msg.upper())  # 发消息(发大写的收到的消息过去)

# 7 关掉发送消息链接
conn.close()
# 8 关掉链接
phone.close()