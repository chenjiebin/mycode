//函数宏定义
//可以使用cpp进行查看
#define MAX(a,b) ((a)>(b)?(a):(b))
int a[] = { 9, 3, 5, 2, 1, 0, 8, 7, 6, 4 };
int max(int n)
{
	return n == 0 ? a[0] : MAX(a[n],max(n-1));
}

int main(void)
{
	max(9);
	return 0;
}