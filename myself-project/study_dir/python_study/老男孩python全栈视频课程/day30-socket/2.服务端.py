from socket import *

ip_port = ('127.0.0.1', 8080)
back_log = 5
buffer_size = 1024

tcp_server = socket(AF_INET, SOCK_STREAM)
tcp_server.bind(ip_port)
tcp_server.listen(back_log)
print('服务端开始运行了')

conn, addr = tcp_server.accept()  # 只有等客户端连进来，否则就会阻塞在这里。
print('双向链接是: ', conn)
print('客户端地址是: ', addr)
# 双向链接是:  <socket.socket fd=556, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 8080), raddr=('127.0.0.1', 58942)>
# 客户端地址是:  ('127.0.0.1', 58942)

while True:
    # 收信息:
    data = conn.recv(buffer_size)

    # 发信息:
    print('收到客户端发来的信息: ', data.decode('utf-8'))
    conn.send(data.upper())

conn.close()
tcp_server.close()
