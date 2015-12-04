#include <arpa/inet.h>

/*
	? 16.3.1 字节序是处理器架构特性(大端、小端)

	TCP/IP协议栈使用大端字节序


	uint32_t htonl(uint32_t hostint32)
	返回值：以主机字节序表示的32位整数

	uint16_t htons(uint16_t hostint16)
	返回值：以主机字节序表示的16位整数

	uint32_t ntohl(uint32_t netint32)
	返回值：以网络字节序表示的32位整数

	uint16_t ntohs(uint16_t netint16)
	返回值：以网络字节序表示的16位整数

*/