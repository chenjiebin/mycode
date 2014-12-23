#include <stdio.h>

int main(void)
{
	int count[4] = {3, 2}, i;
	for (i = 0; i < 4; i++)
		printf("count[%d]=%d\n", i, count[i]);
	return 0;
}