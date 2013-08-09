#include <stdio.h>

int main(void)
{
    extern char **environ;
    int i, j;
    for (i=0; environ[i]!=NULL; i++)
        printf("%s\n", environ[i]);
    return 0;
}
