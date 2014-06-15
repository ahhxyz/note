
#include <stdio.h>
#include "InsertSort.c"
#define LEN 10

int main(){    
    int a[LEN]={90,10,36,45,22,12,56,78,32,99}; 
    printf("未排序：\n");
     for(x=0;x<10;x++){
        printf("%d\n",a[x]);                  
     } 
     printf("已排序：\n");   
    InsertSort(a,LEN);  
    for(x=0;x<10;x++){
        printf("%d\n",a[x]);                  
     }       
}

