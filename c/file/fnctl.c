#include <fcntl.h>
#include <stdio.h>

int main(void)
{
	//打开一个文件
	int f;
	f = open("test.log", O_RDWR | O_APPEND | O_CREAT, S_IRWXU | S_IRWXG | S_IRWXO);
	printf("%d\n", f);

	int flags;
	flags = fcntl(f, F_GETFL);
	printf("%d\n", flags);
	
	return 0;
}