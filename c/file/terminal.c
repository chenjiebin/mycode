#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <fcntl.h>
#include <errno.h>
#include <string.h>

#define MSG_TRY "try again\n"
#define MSG_TIMEOUT "timeout\n"

void blockTerminal()
{
	char buf[10];
	int n;
	n = read(STDIN_FILENO, buf, 10);
	if (n < 0) {
		perror("read STDIN_FILENO");
		exit(1);
	}
	write(STDOUT_FILENO, buf, n);
}

void nonblockTerminal()
{
	char buf[10];
	int fd, n;
	fd = open("/dev/tty", O_RDONLY|O_NONBLOCK);
	if (fd < 0) {
		perror("open /dev/tty");
		exit(1);
	}

	tryagain:
	n = read(fd, buf, 10);
	if (n < 0) {
		if (errno == EAGAIN) {
			sleep(1);
			write(STDOUT_FILENO, MSG_TRY, strlen(MSG_TRY));
			goto tryagain;
		}
		perror("read /dev/tty");
		exit(1);
	}

	write(STDOUT_FILENO, buf, n);
	close(fd);
}

void nonblockTerminalAndTimeout()
{
	char buf[10];
	int fd, n, i;
	fd = open("/dev/tty", O_RDONLY|O_NONBLOCK);
	if (fd < 0) {
		perror("open /dev/tty");
		exit(1);
	}

	for (i = 0; i < 5; i++) {
		n = read(fd, buf, 10);
		if (n >= 0) {
			break;
		}
		if (errno != EAGAIN) {
			perror("read /dev/tty");
			exit(1);
		}
		sleep(1);
		write(STDOUT_FILENO, MSG_TRY, strlen(MSG_TRY));
	}
	if (i == 5) {
		write(STDOUT_FILENO, MSG_TIMEOUT, strlen(MSG_TIMEOUT));
	} else {
		write(STDOUT_FILENO, buf, n);
	}

	close(fd);
}

int main(void)
{
	nonblockTerminalAndTimeout();
	return 0;
}

