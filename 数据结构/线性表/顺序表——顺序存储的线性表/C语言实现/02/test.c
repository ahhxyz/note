/*
 *测试用例
 *
 *
 */
#include <stdio.h>
#include "lib/SeqList.c"
/*
 *遍历顺序表中的内容
 */
int   PrintAll(SeqList *sl){
	for (int i = 0; i < sl->Len-1; ++i){
		printf("节点名：%s\n",sl->Data );
	}
}                         

int main(){
	int i;
	SeqList sl;
	char Data,*ptr;
	char key[15];
	Init(&sl);

	do{
		printf("输入添加节点的名称\n");
		fflush(stdin);		//	清空输入缓冲区
		scanf("%s",&Data);
		if(Data=="exit"){
			break;
		}
		if(!Add(&sl,Data)){
			continue;
		}
	}while(1);

	printf("顺序表中节点顺序是：\n");
	PrintAll(&sl);

	fflush(stdin);
	printf("请输入要取出的节点的索引：\n");
	scanf("%d",&i);
	ptr=FindByIndex(&sl,i);
	if(ptr){
		printf("索引为%d的节点为：[%s%s%d]\n",i,ptr.Data);
	}

	fflush(stdin);
	printf("请输入要查找的节点的关键字：\n");

	scanf("%s",key);
	i=FindByKey(&sl,key);
	ptr=FindByIndex(&sl,i);
	if(ptr){
		printf("关键字为%s的节点为：[%s%s%d]\n",key,ptr.Data);
	}
	getch();
	return 0;

}
