// 单向链表
// 静态生成的链表
#include <stdio.h>

// 单向链表结构
struct node
{
	int num;
	struct node *next;
};

int main(void)
{

	struct node a, b, c, *head, *p;
	a.num = 7;
	b.num = 10;
	c.num = 5;
	head = &a;
	a.next = &b;
	b.next = &c;
	c.next = NULL;
	p=head;
	do {
		printf("%d\n", p->num);
		p = p->next;
	} while (p!=NULL);
	return 0;
}