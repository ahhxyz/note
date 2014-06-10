##selinux
>查看：
>	
	$ /usr/sbin/sestatus -v
SELinux status: enabled就表示开启状态。
>###永久关闭：
>
	# vi /etc/selinux/config 
将`SELINUX=enforcing`改为`SELINUX=disabled`
>###临时
>>关闭：
>>
	# /usr/sbin/setenforce 0    
开启：
>>	
	# /usr/sbin/setenforce 1

#虚拟主机配置
网站路径如果是在/home下面，就会提示找不到目录；所以放到/var下面

##vhosts.conf:
>###PHP5.5-
>
	<VirtualHost  *:80>
	ServerName www.demo.com
	DocumentRoot /var/phpApp/demo
    #  <Directory />
    #   Options FollowSymLinks
    #   AllowOverride all
    #   Order allow,deny
    #   Allow from all 
    #  </Directory>
	</VirtualHost>
>###PHP5.5+
>
	<VirtualHost *:80>
	ServerName www.demo.com
	DocumentRoot "/var/phpApp/demo"
      <Directory />
       Options FollowSymLinks
       AllowOverride all
       Require all granted 
      </Directory>
	</VirtualHost>

如果是配置的本地域名，则还需要在host加上`ServerName`的值：

	# echo 127.0.0.1 demo >> /etc/host
之后，在浏览器输入：`http:demo`就可以访问站点了。

>#####另外，如果需要还要加上www的配置，不然原先的localhost就打不开
>
	<VirtualHost *:80>    
    	DocumentRoot "/var/www/home"
    	ServerName localhost    
	</VirtualHost>