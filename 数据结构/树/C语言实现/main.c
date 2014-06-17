/*
 *二叉树
 */


#include <stdio.h>
#include <stdlib.h>
#include "BinTree.c"

int main(){
    ChainBinTree *d,*node;
    printf("申请内存来存放一个节点，并以此节点来初始化二叉树\n");
    d=init();
    //node=find(d,"A");
    printf("%s\n",d->data);
    //添加节点
    /*
    do{
        printf("请输入父节点的值和节点的值：");
        if(node=(ChainBinTree *)malloc(sizeof()))
    }
     */
    
}