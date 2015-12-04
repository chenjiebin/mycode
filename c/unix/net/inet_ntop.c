#include <arpa/inet.h>

/*
	const char *inet_ntop(int domain, const void *restrict addr,
						char *restrict str, socklen_t size);
	返回值：若成功，返回地址字符串指针；若出错，返回NULL

	int inet_pton(int domain, const char * restrict str, void *restrict addr);
	返回值：若成功，返回1；若格式无效，返回0；若出错，返回-1
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>


int main(void)
{
	char IPdotdec[20];
	struct in_addr s;

	printf("Please input IP address: ");
	scanf("%s", IPdotdec);

	inet_pton(AF_INET, IPdotdec, (void *)&s);
	printf("inet_pton:0x%x\n", s.s_addr);
	printf("inet_pton:%d\n", s.s_addr);
	
	inet_ntop(AF_INET, (void *)&s, IPdotdec, 16);
	printf("inet_ntop:%s\n", IPdotdec);

	return 0;
}