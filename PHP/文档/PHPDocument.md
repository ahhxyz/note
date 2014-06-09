##安装：
>* 1.直接下载安装：
    
    https://github.com/pear/PhpDocumentor
然后将该文件夹放到PHP的安装目录，也就是php.exe的同目录
>* 2.使用pear安装：
    pear install PhpDocumentor

##使用：
>>* 1.最基本：
    phpdoc phpdoc -o HTML:frames:earthli -f 要生成文档的PHP文件或目录的路径 -t 文档的生成目录

>>>>>>* 1.首次运行phpdoc要选择配置文件，根据提示选择即可。
>>>>>>* 2.生成后，在文档目录用浏览器打开index.html接口查看文档。