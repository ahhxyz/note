#include <stdio.h>


typedef struct node{
    char data;
    struct node *left;
	struct node *right;
}Node



Node *init();//初始化二叉树

int add(Node *p, Node *node, int n);//添加节点到某个节点下面




