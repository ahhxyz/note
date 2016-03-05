#前言
PHP中数组是基于哈希表实现的
通常哈希表的这些操作时间复杂度为O(1)。 这也是它被钟爱的原因。
正是因为哈希表在使用上的便利性及效率上的表现，目前大部分动态语言的实现中都使用了哈希表。

所有的用户端定义的变量保存在一个符号表里，而这个符号表其实就是一个HashTable，它的每一个元素都是一个zval*类型的变量。不仅如此，保存用户定义的函数、类、资源等的容器都是以HashTable的形式在内核中实现的。

#基本概念

哈希表是一种通过哈希函数，将特定的键映射到特定值的一种数据结构，它维护键和值之间一一对应关系。

* 键(key)：用于操作数据的标示，例如PHP数组中的索引，或者字符串键等等。
* 槽(slot/bucket)：哈希表中用于保存数据的一个单元，也就是数据真正存放的容器。
* 哈希函数(hash function)：将key映射(map)到数据应该存放的slot所在位置的函数。
* 哈希冲突(hash collision)：哈希函数将两个不同的key映射到同一个索引的情况。

哈希表可以理解为数组的扩展或者关联数组，数组使用数字下标来寻址，如果关键字(key)的范围较小且是数字的话， 我们可以直接使用数组来完成哈希表，而如果关键字范围太大，如果直接使用数组我们需要为所有可能的key申请空间。 很多情况下这是不现实的。即使空间足够，空间利用率也会很低，这并不理想。同时键也可能并不是数字， 在PHP中尤为如此，所以人们使用一种映射函数(哈希函数)来将key映射到特定的域中：

	h(key) -> index
通过合理设计的哈希函数，我们就能将key映射到合适的范围，因为我们的key空间可以很大(例如字符串key)， 在映射到一个较小的空间中时可能会出现两个不同的key映射被到同一个index上的情况， 这就是我们所说的出现了冲突。 目前解决hash冲突的方法主要有两种：链接法和开放寻址法。


每个HashTable中的元素都有两部分组成：索引与值， 每个元素的值都是一个独立的zval（确切的说应该是指向某个zval的指针）。


HashTable结构：


	typedef struct _hashtable {
    	uint nTableSize; //表长度，并非元素个数
 		uint nTableMask;//表的掩码，始终等于nTableSize-1
 		uint nNumOfElements;//存储的元素个数
 		ulong nNextFreeElement;//指向下一个空的元素位置
 		Bucket *pInternalPointer;//foreach循环时，用来记录当前遍历到的元素位置
 		Bucket *pListHead;
 		Bucket *pListTail;
 		Bucket **arBuckets;//存储的元素数组
 		dtor_func_t pDestructor;//析构函数
 		zend_bool persistent;//是否持久保存。从这可以发现，PHP数组是可以实现持久保存在内存中的，而无需每次请求都重新加载。
 		unsigned char nApplyCount;
 		zend_bool bApplyProtection;
	} HashTable;
    
bucket结构：

	typedef struct bucket {
     	ulong h; //数组索引
     	uint nKeyLength; //字符串索引的长度
     	void *pData; //实际数据的存储地址
     	void *pDataPtr; //引入的数据存储地址
     	struct bucket *pListNext;
     	struct bucket *pListLast;
     	struct bucket *pNext; //双向链表的下一个元素的地址
     	struct bucket *pLast;//双向链表的下一个元素地址
     	char arKey[1]; /* Must be last element */
	} Bucket;