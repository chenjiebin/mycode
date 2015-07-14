#include <stdio.h>
#define X 3
#undef X
#define X 2

int main(void) {
    printf("hello");
    printf("%d", X);
}