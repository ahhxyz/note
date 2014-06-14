这里以命令行CLI执行PHP为例，
	$ php -f test.php

CLI的主源代码文件在{PHPSRC}/sapi/cli/php_cli.c,整个过程就从这个文件中的 main()函数执行，内部的的执行主要可以分为以下几个步骤：

* 1.解析命令行参数
* 2.初始化环境
* 3.解析PHP代码，并执行生成的opcode
* 4.清理并推出

下面重点说明`步骤3`的主要逻辑：

* 1.`{PHPSRC}/main/main.c`中的[`php_execute_script(&file_handle TSRMLS_CC)`](https://github.com/php/php-src/blob/master/main/main.c#L2466)
