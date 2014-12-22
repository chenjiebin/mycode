#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>

int main(void)
{
	//打开一个文件
	int f;
	f = open("test.log", O_RDWR | O_APPEND | O_CREAT, S_IRWXU | S_IRWXG | S_IRWXO);
	printf("%d\n", f);

	//写
	char wbuf[13] = "Hello world!\n";
	write(f, wbuf, sizeof(wbuf));

	//将文件偏移量移到文件开始
	lseek(f, 0, SEEK_SET);

	//读
	char rbuf[12];
	int size;
	size = read(f, rbuf, sizeof(rbuf) - 1);
	printf("%d\n", size);
	rbuf[size] = '\0';
	printf("%s\n", rbuf);

	//关闭文件
	close(f);
	
	return 0;
}