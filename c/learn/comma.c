#include <stdio.h>

int main()
{
    int a = 5;
    a = (a=3*5, a*4), a+5;
    printf("%d\n", a);
}
