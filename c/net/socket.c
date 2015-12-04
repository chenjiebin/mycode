#include <sys/socket.h>


/*
	int socket (int domain, int type, int protocol)
	返回值：若成功，返回文件（套接字）描述符；若出错，返回-1

	domain
		AF_INET		IPv4因特网域
		AF_INET6	IPv6因特网域
		AF_UNIX		UNIX域
		AF_UPSPEC	未指定
	type
		SOCK_DGRAM		固定长度的，无连接的、不可靠的报文传输
		SOCK_RAW		IP协议的数据报接口(在POSIX.1中为可选)
		SOCK_SEQPACKET	固定长度的、有序的、可靠的、面向连接的报文传递
		SOCK_STREAM		有序的、可靠的、双向的、面向连接的字节流
	protocol
		IPPROTO_IP 		IPv4网际协议
		IPPROTO_IPV6	IPv6网际协议（在POSIX.1中为可选）
		IPPROTO_ICMP	因特网控制报文协议（Internet Control Message Protocol）
		IPPROTO_RAW		原始IP数据包协议（在POSIX.1中为可选）
		IPPROTO_TCP		传输控制协议
		IPPROTO_UPD		用户数据报协议（User Datagram Protocol）
	
	domain为AF_INET，type为SOCK_DGRAM默认协议为UDP
	domain为AF_INET，type为SOCK_STREAM默认协议为TCP

	套接字可以传递的函数
	close		释放套接字
	dup和dup2	和一般文件描述符一样复制
	fchdir		失败，并且将errno设置为ENOTDIR
	fchomod		未指定
	fchown		由实现定义
	fcntl		支持一些命令，包括F_DUPFD、F_DUPFD_CLOEXEC、F_GETED,F_GETFL,F_GETOWN,F_SETFD,F_SETFL,F_SETOWN
	fdatasync和fsync		由实现定义
	fstat		支持一些stat结构成员，但如何支持由实现定义
	ftruncate	未指定
	ioctl		支持部分命令、依赖于底层设备驱动
	lseek		由现实定义（通常失败时会将errno设为ESPIPE）
	mmap		未指定
	poll		正常工作
	pread和pwrite	失败时会将errno设为ESPIPE
	read和readv	与没有任何标志位的recv等价
	select		正常工作
	write		与没有任何标志位的send等价


	? 报文和字节流的区别，SOCK_SEQPACKET具体示例
	? 
*/



