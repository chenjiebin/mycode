/* 
 * File:   main.c
 * Author: chenjiebin
 *
 * Created on October 16, 2014, 12:42 PM
 */

#include <stdio.h>
#include <stdlib.h>

/*
 * 
 */
int main(int argc, char** argv) {
    int count[4];
    
    count[0] = 7;
    count[1] = count[0] *2;
    ++count[2];
    
    int i;
    
    for (i = 0;i < 4;i++) {
        printf("count[%d]=%d\n", i, count[i]);
    }
    
    printf("Hello, world.\n");

    return (EXIT_SUCCESS);
}

