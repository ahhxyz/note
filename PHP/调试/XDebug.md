XDebug是作为PHP扩展来安装的。

关于配置：
xdebug启动后，在phpinfo()的xdebug的配置列表那里就可以查看到所有的配置项，根据需要在php.ini中配置即可。
![](https://github.com/sunbay/note/blob/master/PHP/%E8%B0%83%E8%AF%95/xdebug.png)

结合netbeans调试时的常见问题：
>* 1.一直是：等待连接netbeans-xdebug。
>>解决办法：netbeans里面的ID、端口要和phpinfo中xdebug的IDE Key及php.ini配置里面xdebug.remote_port一致。
