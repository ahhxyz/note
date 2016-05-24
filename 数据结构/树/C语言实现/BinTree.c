/*
 *二叉树
 * 代码未测试
 */


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "BinTree.h"

/**
 * 以一个Node类型的指针作为根节点来初始化二叉树
 * @param node 动态申请的指向Node类型的指针
 * @return  (Node *)
 */
Node *init(){
    Node *node;
    if(node = (Node *) malloc( sizeof(Node) ) ){
        printf("请输入根节点的值：");
        fgets(node->data, 10, stdin);
        node->left = NULL;
        node->right = NULL;
        printf("%s\n", node->data);
        return node;
    }
    return NULL;
}


int add(Node *p, Node *node, int n){
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
            p->left = node;
            break;
        case 2:
            if(p->right){
                printf("右子树节点已存在！");
            }
            p->right = node;
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
Node *find(Node *d,char *data){
    if(!d){
        return NULL;
    }
    
    if(strcmp(d->data,data) == 0){
        return d;
    }
    
    if(strcmp(d->left->data,data) == 0){
        return d;
    }
    
    if(strcmp(d->right->data,data) == 0){
        return d;
    }
}