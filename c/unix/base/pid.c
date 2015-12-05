#include "apue.h"

int main(void) {
	printf("hello world from proess Id %ld\n", (long)getpid());
	exit(0);
}