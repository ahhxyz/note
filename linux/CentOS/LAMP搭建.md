##安装apache：##
  	#  yum install httpd
  	#  chkconfig --levels 235 httpd on
  	#  /etc/init.d/httpd start
  此时输入localhost就可以看到apache的欢迎界面。

##安装pgsql93：
	# yum install http://yum.postgresql.org/9.3/redhat/rhel-6-x86_64/pgdg-redhat93-9.3-1.noarch.rpm
	# yum install postgresql93-server postgresql93-contrib
	# service postgresql-9.3 initdb
	# chkconfig postgresql-9.3 on
	# service postgresql-9.3 start



##安装mysql;
  	#  yum install mysql mysql-server
  	#  chkconfig --levels 235 mysqld on
  	#  /etc/init.d/mysqld start
  	#  mysql_secure_installation

##安装php：##
  	#  yum install php
  	

>###php5.5的安装方法：###
>>#####CentOS/RHEL 6.x:
	rpm -Uvh http://mirror.webtatic.com/yum/el6/latest.rpm
	yum install php55w php55w-opcache -y
>###安装PHP扩展
>>
	yum  install php55w-pgsql php55w-gd -y

安装完成后：
>	
	#  /etc/init.d/httpd restart

