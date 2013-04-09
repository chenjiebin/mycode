#include <stdio.h>

struct complex_struct { double x, y; } z;

struct complex_struct add_complex(struct complex_struct z1, struct complex_struct z2);

int main(void)
{
    struct complex_struct z;

    double x = 3.0;
    z.x = x;
    z.y = 4.0;
    if (z.y < 0)
        printf("z=%f%fi\n", z.x, z.y);
    else
        printf("z=%f+%fi\n", z.x, z.y);

    //两个函数相加
    z = add_complex(z, z);
    printf("z=%f+%fi\n", z.x, z.y);
        
    return 0;
    
}

struct complex_struct add_complex(struct complex_struct z1, struct complex_struct z2)
{
    z1.x = z1.x + z2.x;
    z1.y = z1.y + z2.y;
    return z1;
}




