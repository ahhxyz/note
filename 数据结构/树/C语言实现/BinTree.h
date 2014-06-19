#include <stdio.h>

/*
typedef struct {
    char *name;
    int age;
}DATA;
*/

typedef struct BinTree{
    char *data;
    struct BinTree *left;
    struct BinTree *right;
}ChainBinTree;

ChainBinTree *init();//初始化二叉树

int add(ChainBinTree *p,ChainBinTree *node,int n);//添加节点到某个节点下面




