/*
 *1.定义顺序表的数据结构；
 *2.声明操作数据表的函数；
 */


#include <stdio.h>
#include <string.h>
#define MAXSIZE 100                    			//定义线性表的最大长度
typedef struct {
	char name[20];
	char key[15];    							//节点关键字
	int  value;
}DATA;

typedef struct{                  				 //定义顺序表结构
	DATA Data[MAXSIZE];          				 //保存顺序表的数组。由于C语言中的数组也是采用的顺序存储的方式，因此可以使用数组来模拟顺序表的保存形式。
	int Len;                       				 //顺序表已存结构的数量
}SeqList;

/*
 *顺序表操作函数的声明
 *
 *
 */

void  Init(SeqList *sl);    					//初始化顺序表；
int   Length(SeqList *sl);  					//返回顺序表的元素数量
int   Add(SeqList *sl,DATA data);          		//向顺序表中添加元素
int   Insert(SeqList *sl,int n, DATA data);		//向顺序表中插入元素
int   Delete(SeqList *sl,int n);                //删除顺序表中的元素
DATA* FindByIndex(SeqList *sl,int n);           //根据序号返回元素
int   FindByKey(SeqList *sl,char *key);         //按关键字查找
