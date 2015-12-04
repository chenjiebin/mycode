#include <sys/socket.h>

/*
	int shutdown (int sockfd, int how);
	返回值：若成功，返回0；若出错，返回-1
	
	how
		SHUT_RD	关闭终端，无法从套接字读取数据
		SHUT_WR 关闭写端，无法从套接字发送数据
		SHUT_RDWR	无法读取数据，无法发送数据
*/