有2种方式来使用git
#一、创建本地仓库：
切换到要作为仓库的目录，创建一个本地的仓库：

    git init

首次运行git，需运行 
    
    git  git config --global user.email "you@example.com"
    
    git  git config --global user.name "Your Name"    
在仓库添加、编辑文件
将文件更改添加到本地仓库：

    git add
提交更改到了本地仓库：

    git commit -m “描述”  


如果要将本地库与远程仓库关联同步，那么可以采取下面2中方式：
	
* 1.https方式
* 2.ssh方式：
	
	在本地计算机使用RSA算法生成密钥：
    
    	ssh-keygen -t rsa -C "youremail@example.com"
将公钥上传到服务器或将公钥文件的内容填写在github等的相应地方；

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
首先安装git:
	
    $ yum install git -y
##小团队

* 1. 在服务器创建一个git用户:
    
    		git useradd demo@demo.com
    
	>如果是SSH方式，还要收集所有需要登录的用户的公钥，就是他们自己的id_rsa.pub文件，把所有公钥导      入到`/home/git/.ssh/authorized_keys`文件里，一行一个。


* 2. 在服务器初始化Git仓库:
	先选定一个目录作为Git仓库，这里以`/data/git`为例；
	在目录下输入命令：
    		$  git init --bare demo.git
		
	> 如果是从其他地方迁移GIT仓库则使用:`git clone --bare demo.git`

    > 更改目录权限：`chown -R git:git demo@demo.com`

* 3. 克隆远程仓库：
    $ git clone git@yourserver:/git/demo.git	
　　
> 如果本地已初始化了一个仓库，则直接将本地仓库与远程仓库关联：
    		git remote add origin demo@demo.com@yourserver:/git/demo.git

	
    
    
>注：　　yourserver为服务器地址(IP或域名)。    

>	设置长期保存密码： `git config --global credential.helper store`



---
　　

　　
* ## `大团队`:使用gitosis或gitolite来管理公钥

Gitosis和gitolite只是Git 服务管理工具，不需要它们一样可以搭建git服务器。

## GIT服务器迁移

* 迁移到自己的服务器：在新GIT服务器的git目录执行`git clone --bare 原git地址`即可
* 迁移到github等代码托管平台：除了需要现在平台创建仓库，还需要通过本地来执行以下操作:
	* 	`git clone --bare 原git地址`
	* 	`git push --mirror git平台上的仓库地址`

---


#常用GIT命令：
###分支操作：

* 查看分支：
	
		git branch

* 创建分支：  
*   
    	git branch name
    
* 切换分支：
    
    	git checkout name
* 创建并切换分支：
    
    	git checkout -b name

* 从远程分支提取

		git fetch origin
    
* 合并指定分支到当前分支：
    
    	git merge name
* 从远程获取最新代码并合并到本地分支

		git pull 远程分支origin 本地分支master
* 删除分支：
    
    	git branch -d name
    	
# 回滚到某一历史版本

* git revert 是生成一个新的提交来撤销某次提交，此次提交`之前`的commit都会被__保留__
* git reset 是`回到`某次提交，提交`及`之前的commit都会被保留，但是此次之后的修改都会被退回到暂存区

* 示例：

`git revert eda16b24f1c293708e7eba20692bd8a71a2ac929`

   恢复到eda16b24f1c293708e7eba20692bd8a71a2ac929这个commit之前的版本，
   即撤销该commit及之后的所有commit
   
   > 如果提示`is a merge but no -m option was given`的错误，这是因为你revert的那个commit是一个merge commit，它有两个parent, 
   Git不知道base是选哪个parent，就没法diff，所以就抱怨了，所以你要显示告诉Git用哪一个parent。
   一般来说，如果你在master上merge了dev分支,那么parent 1就是master，parent 2就是dev分支.
   根据情况，加上`-m 1`或`-m 2`参数即可



