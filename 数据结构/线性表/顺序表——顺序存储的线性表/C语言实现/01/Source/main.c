/*
 *测试用例
 *
 *
 */
#include <stdio.h>
#include "SeqList.c"
/*
 *遍历顺序表中的内容
 */
int   PrintAll(SeqList *sl){
	for (int i = 0; i < sl->Len; ++i){
		printf("节点名：%s；关键字：%s；节点值：%d\n",sl->Data[i].name,sl->Data[i].key,sl->Data[i].value );
	}
}                         

int main(){
	int i;
	SeqList sl;
	DATA data,*ptr;
	char key[15];
	Init(&sl);

	do{
		printf("输入添加节点的（名称、关键字、值）\n");
		fflush(stdin);		//	清空输入缓冲区
		scanf("%s%s%d",data.name,data.key,&data.value);
		if(!data.value||!Add(&sl,data)){
			printf("break\n");
			break;
		}
	}while(1);

	printf("顺序表中节点顺序是：\n");
	PrintAll(&sl);

	fflush(stdin);
	printf("请输入要取出的节点的索引：\n");
	scanf("%d",&i);
	ptr=FindByIndex(&sl,i);
	if(ptr){
		printf("索引为\"%d\"的节点为：[%s,%s,%d]\n",i,ptr->key,ptr->name,ptr->value);
	}

	fflush(stdin);
	printf("请输入要查找的节点的关键字：\n");

	scanf("%s",key);

	i=FindByKey(&sl,key);
	ptr=FindByIndex(&sl,i);
	if(ptr){
		printf("关键字为\"%s\"的节点为：[%s,%s,%d]\n",key,ptr->key,ptr->name,ptr->value);
	}
	getchar();
	return 0;

}
