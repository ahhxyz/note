/*
 *二叉树
 */


#include <stdio.h>
#include <stdlib.h>
#include "BinTree.h"

/**
 * 以一个ChainBinTree类型的指针作为根节点来初始化二叉树
 * @param node 动态申请的指向ChainBinTree类型的指针
 * @return  (ChainBinTree *)
 */
ChainBinTree *init(){
    ChainBinTree *node;
    if(node=(ChainBinTree *)malloc(sizeof(ChainBinTree))){
        printf("请输入根节点的值：");
        scanf("%d",&node->data);
        node->left=NULL;
        node->right=NULL;
        
        return node;
    }
    return NULL;
}


int add(ChainBinTree *p,ChainBinTree *node,int n){
    if(!p){
        printf("父节点不存在！");
        return 0;
    }
    
    switch(n){
        case 1:
            if(p->left){
                printf("左子树节点已存在！");
                return 0;
            }
            p->left=node;
            break;
        case 2:
            if(p->right){
                printf("右子树节点已存在！");
            }
            p->right=node;
            break;            
        default:
            printf("参数错误");
            return 0;
    }
    
    return 1;
    
}

/**
 * 根据节点值来返回节点指针。可用来将数据插入到指定节点下
 * @param d
 * @param data
 * @return 
 */
ChainBinTree *find(ChainBinTree *d,char *data){
    if(!d){
        return NULL;
    }
    
    if(strcmp(d->data,data)==0){
        return d;
    }
    
    if(d->left->data==data){
        return d;
    }
    
    if(d->left->data==data){
        return d;
    }
}