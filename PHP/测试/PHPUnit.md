
>#使用XAMPP
>>XAMPP中已自带了pear和PHPUnit，无需进行安装。
>#使用PHP官方包
>>* ##1.首先安装pear:
>>>如果安装目录没有go-pear.bat，则去http://pear.php.net/go-pear.phar下载。
然后再安装目录执行`php go-pear.phar`，
就会安装pear，并生成一个pear目录和pear.bat。
如果按照目录存在go-bear.bat则直接在命令行运行也能完成pear的安装。

>>* ###2.完成pear的安装后，就要安装PHPUnit了
>>>* 1.首先添加频道：`pear channel-discover pear.phpunit.de`
>>>* 2.安装PHPUnit:  `pear install phpunit/PHPUnit` 。完成后会在PHP安装目录生成phpunit和phpunit.bat文件。



>#安装Skeleton框架生成器:
    PHPUunit只是个测试框架，并不会生成测试代码，要生成测试代码，就要再安装skeleton.
    安装前要确保执行了前面的添加频道的操作。
>>* 1.列出服务器上面可用的：`pear remote-list -c phpunit`
>>* 2.安装：`pear install phpunit/PHPUnit_SkeletonGenerator`
>>>安装时的常见错误：
>>>>* 如果提示依赖包满足不了要求就安装或更新对应包，如：pear install components.ez.no/ConsoleTools
>>>>* 如果提示无法写入cache，就删除$/pear/cache下面的所有文件再进行安装

断言的格式是：`@assert (0,0) == 0`，少了中间的空格就无法正确生成测试用例
常用pear命令：

已安装的：
pear list

列出可用更新：
pear list-upgrades

更新所有：
pear upgrades-all
