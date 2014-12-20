#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>

int main(void)
{
	//打开一个文件
	int f;
	f = open("test.log", O_RDWR|O_APPEND|O_CREAT, "0666");
	printf("%d\n", f);

	//写
	char buf[12] = "Hello world!";
	write(f, buf, sizeof(buf));

	//读
	char rbuf[12];
	lseek(f, 0, SEEK_SET);
	read(f, rbuf, 12);
	printf("%s\n", rbuf);

	close(f);
	
	return 0;
}