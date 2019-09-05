from socket import *

ip_port = ('127.0.0.1', 8080)
# msg = 'hello'
buffer_size = 1024

tcp_client = socket(AF_INET, SOCK_STREAM)
tcp_client.connect(ip_port)

while True:
    msg = input('>>>:').strip()
    # 发信息:
    tcp_client.send(msg.encode('utf-8'))
    print('客户端已发送消息')
    # 收信息:
    data = tcp_client.recv(buffer_size)
    print('收到服务端的发来的消息: ', data.decode('utf-8'))  # 解码

tcp_client.close()
