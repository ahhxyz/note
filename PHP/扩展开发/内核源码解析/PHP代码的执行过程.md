这里以命令行CLI执行PHP为例，
	$ php -f test.php

CLI的主源代码文件在{PHPSRC}/sapi/cli/php_cli.c,整个过程就从这个文件中的 main()函数执行，内部的的执行主要可以分为以下几个步骤：

* 1.解析命令行参数
* 2.初始化环境
* 3.解析PHP代码，并执行生成的opcode
* 4.清理并推出

下面重点说明`步骤3`的主要逻辑：

* 1.执行`{PHPSRC}/main/main.c`中的[`php_execute_script(zend_file_handle *primary_file TSRMLS_DC)`](https://github.com/php/php-src/blob/master/main/main.c#L2466)函数：
	
	`zend_file_handle`类型是zend对`文件句柄`的一个封装，这个句柄指向了正在执行的PHP文件，这里就是demo.php。
* 2.`php_execute_script()`调用`zend_execute_scripts()`：
		
	在`zend_execute_scripts()`中，最重要的就是：
>* a.调用`ZEND_API zend_op_array *(*zend_compile_file)(zend_file_handle *file_handle, int type TSRMLS_DC)`函数指针来解析PHP文件，最后返回一个zend_op_array()类型的指针，执行了生成的opcode；
>* b.调用`ZEND_API extern void (*zend_execute)(zend_op_array *op_array TSRMLS_DC)`来执行`opcode，在这个过程中，就实现了PHP的各个功能`。
