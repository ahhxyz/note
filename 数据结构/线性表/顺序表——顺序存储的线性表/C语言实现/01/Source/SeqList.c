/*
 *定义操作顺序表的函数；
 *
 */


#include <stdio.h>
#include <string.h>
#include "../Header/SeqList.h"

/*
 *初始化顺序表；
 */
void Init(SeqList *sl){    						

	sl->Len=0;
}


/*
 *返回顺序表的元素数量；
 */

int   Length(SeqList *sl){  						

	return (sl->Len);
}


/*
 *向顺序表中添加元素
 *
 */
int   Add(SeqList *sl,DATA data){           		

	if (sl->Len>=MAXSIZE){
		printf("顺序表已满，不能再添加结点\n");
	    return 0;
	}
	sl->Data[sl->Len++]=data;
	return 1;

}

/*
 *向顺序表中插入元素
 *@param:
 *@n:相对应当前顺序表的起始位置的偏移,其实就是数组Data的索引；
 */
int   Insert(SeqList *sl,int n, DATA data){		
	if(sl->Len > MAXSIZE){
		printf("顺序表已满，不能再添加结点\n");
		return 0;
	}
	if (n < 0||n > sl->Len-2){
		printf("节点的序号错误！\n");
		return 0;
	}
	//将插入位置n及其后的每个元素向后移动一位
	for (int i = sl->Len-2;i>=n; --i){
		sl->Data[i+1]=sl->Data[i];
	}
	sl->Data[n]=data;
	++sl->Len;
	return 1;

}
/*
 *删除顺序表中的元素
 */
int   Delete(SeqList *sl,int n){
	if (n < 0||n > sl->Len-2){
		printf("节点的序号错误！\n");
		return 0;
	}
	//将删除位置n以后的每个元素向前移动一位：
	for (int i =n; i<sl->Len-1; ++i){
		sl->Data[i]=sl->Data[i+1];
	}
	sl->Len--;
	return 1;

}
/*
 *根据序号返回元素
 */
DATA *FindByIndex(SeqList *sl,int n){
	if (n < 0||n > sl->Len-1){
		printf("节点的序号错误！\n");
		return 0;
	}
	
	return 	&(sl->Data[n]);
}       
/*
 *按关键字查找
 */         
int   FindByKey(SeqList *sl,char *key){
	for (int i = 0; i <sl->Len; ++i)	{
		if(strcmp(sl->Data[i].key,key)==0){			//如果找到所需节点
			return i;								//返回索引
		}
	}
	return -1;

}
