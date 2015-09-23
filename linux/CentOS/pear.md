全称为PHP扩展与应用库(PHP Extension and Application Repository)，这个库是用`PHP代码`编写的。
#安装pear
其实就是安装pear的包管理器

切换到php的安装目录
	
	$ wget http://pear.php.net/go-pear.phar
	$ php go-pear.phar
> 如果第二步没任何反应，则执行：`php -d phar.require_hash=0 go-pear.phar`
根据提示按`Enter`键就完成了pear的安装，此时，pear命令已生效。

>关于`windows`下安装的说明：
>>* 1.PHP的安装目录没有pear.bat才需要安装pear。
>>* 2.安装后的pear及其里面包的bat文件都是生成在php的安装目录下面的。而linux则被安装在了`/usr/bin`下面，如：`/usr/bin/pear`.
>>* 3.XAMPP包已经内置pear和phpunit等

#安装pear应用程序包：

>* 一、安装PHPUnit:
>>	
>
	pear channel-discover pear.phpunit.de	
	pear install phpunit/PHPUnit
安装完成，`phpunit`命令生效。





#常用pear命令：

更新PEAR
	
	$ pear upgrade PEAR

已安装的包：

	$ pear list

列出包的可用更新：

	$ pear list-upgrades

更新所有：

	$ pear upgrades-all