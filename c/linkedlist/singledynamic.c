// 单向链表
// 动态生成的链表，通过用户输入，如果输入为0则中断输入
#include <stdio.h>
#include <stdlib.h>
#define LEN sizeof(struct node)

// 单向链表结构
struct node
{
	int num;
	struct node *next;
};

int n;

// 创建node节点
// 算法实现的思路是p1指向新开辟的结点，p2指向链表最后一个结点
struct node *creat(void)
{
	struct node *head;
	struct node *p1, *p2;

	n = 0;
	p1 = p2 = (struct node *)malloc(LEN);
	scanf("%d", &p1->num);

	head = NULL;
	while (p1->num != 0) {
		n = n + 1;
		if (n == 1) head = p1;
		else p2->next = p1;
		p2 = p1;

		p1 = (struct node *)malloc(LEN);
		scanf("%d", &p1->num);
	}
	p2->next = NULL;
	return(head);
}

int main(void)
{
	struct node *pt;
	pt = creat();
	// 打印出头
	printf("\n%d\n", pt->num);
	// 打印出整个链表
	do {
		printf("%d\n", pt->num);
		pt = pt->next;
	} while (pt != NULL);
	return 0;
}