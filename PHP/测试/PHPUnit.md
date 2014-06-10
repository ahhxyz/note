>#安装Skeleton框架生成器:
    PHPUunit只是个测试框架，并不会生成测试代码，要生成测试代码，就要再安装skeleton.
    安装前要确保执行了前面的添加频道的操作。
>>* 1.列出服务器上phpunit分类下的可用包：
>	
	pear remote-list -c phpunit

>>* 2.安装：
>
	pear install phpunit/PHPUnit_SkeletonGenerator
>>>安装时的常见错误：
>>>>* 如果提示依赖包满足不了要求就安装或更新对应包，如：
>>>>
	pear install components.ez.no/ConsoleTools
>>>>* `windows`下如果提示无法写入cache，就删除$/pear/cache下面的所有文件再进行安装

安装完成后，使用：
	
	$ phpunit-skeleton 
来生成测试用例。

断言的格式是：`@assert (0,0) == 0`，少了中间的空格就无法正确生成测试用例

