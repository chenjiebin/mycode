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

	char buf[12] = "Hello world!";
	write(f, buf, 12);
	
	return 0;
}