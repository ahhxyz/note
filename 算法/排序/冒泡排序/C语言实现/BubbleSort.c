include "BubbleSort.h"
/*
************************************************************************
*1.对整个序列从头到尾进行遍历
*2.在每次遍历过程中，对待排序序列从后向前进行不断扫描，
*3.当发现相邻的2个记录中后者大于前者时， 将这2个记录交换。
*4.完成了本次遍历的所有扫描后，再重复尚书步骤直到记录结尾。
*这样，小的记录就逐渐从后向前移动，最终完成排序。
*冒泡排序在排序过程中不构建已排序序列，在未完成排序之前，所有元素都是未排序的，
*也就是说，排序过程中的元素序列式没有已排序和未排序之分的。
***********************************************************************/


void BubbleSort(int a[],int len){
    for (int i = 0; i < len; ++i){//对序列从头到尾进行遍历
        for (int j = len-1; j > i && a[j-1] > a[j]; j--){ //每次遍历时，对序列从尾到头进行扫描，
            int tmp=a[j-1];
            a[j-1]=a[j];
            a[j]=tmp;
        }	
    }
}
