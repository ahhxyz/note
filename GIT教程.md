有2种方式来使用git
#一、创建本地仓库：
切换到要作为仓库的目录，创建一个本地的仓库：

    git init
    
在仓库添加、编辑文件
将文件更改添加到本地仓库：

    git add
提交更改到了本地仓库：

    git commit -m “描述”  


如果要将本地库与远程仓库关联同步，那么可以采取下面2中方式：
>>* 1.https方式
>>* 2.ssh方式：
>>>在本地计算机使用RSA算法生成密钥：
    
    ssh-keygen -t rsa -C "youremail@example.com"
>>将公钥上传到服务器或将公钥文件的内容填写在github等的相应地方；

#####关联远程仓库
将本地库与远程库关联：

    git remote add origin https://github.com/sunbay
将本地仓库提交到远程仓库：

    git push [-u] origin master    
首次提交需要加上-u参数


`远程仓库和本地仓库没有不同，只是本地库是属于有开发者个人的，而远程库则是属于一个团队的。`





#二、克隆远程仓库：
1.从远程仓库拉取代码：
    
    git clone 
2.提交到远程库：

    git push https://github.com/sunbay



#GIT服务器搭建：
    $ yum install git -y
>* ##小团队
* 1.创建一个git用户:

    git useradd git
>>>>如果是SSH方式，还要收集所有需要登录的用户的公钥，就是他们自己的id_rsa.pub文件，把所有公钥导      入到/home/git/.ssh/authorized_keys文件里，一行一个。
　　
* 2.初始化Git仓库:
先选定一个目录作为Git仓库，这里以/data/git为例；
在目录下输入命令：

    $  git init --bare demo.git
    chown -R git:git demo.git
* 3.初始化本地仓库:

    git init
    

　　设置长期保存密码：  
    
    git config --global credential.helper store
　　
* 4.将本地仓库与远程仓库关联：
>>>>    
    git remote add origin git@yourserver:/git/demo.git

* 5.克隆远程仓库：


    $ git clone git@yourserver:/git/demo.git
    
    
    

　　yourserver为服务器地址(IP或域名)。
　　
　　
>* ##`大团队`:使用gitosis或gitolite来管理公钥
Gitosis和gitolite只是Git 服务管理工具，不需要它们一样可以搭建git服务器。


#常用GIT命令：
分支操作：
1.查看分支：

    git branch
2.创建分支：

    git branch name
    
3.切换分支：

    git checkout name
4.创建并切换分支：

    git checkout -b name
5.合并指定分支到当前分支：

    git merge name
6.删除分支

    git branch -d name

