/*
 * 1.使用构造函数来定义类的属性及执行初始化操作
 *   构造函数也是一种数据结构，类似于C/golang的结构体struct
 *   构造函数也是一种JavaScript对象，不要用其他语言中函数的思路去理解JavaScript中的函数
 */
function Class(){
    this.property01 = '';
    this.property02 = '';
    //this.init() //实例化时要执行的方法
}

/*
 * 2.使用原型来定义实例的共享成员，即方法和共享的成员变量
 */

Class.prototype.init = function(){
    this.property01 = 1;
}

Class.prototype.say = function(){
    alert(this.property01);
}

//静态方法
Class.run = function(){
    console.log(this);	//此处的this就是上面定义的构造函数，this.name就是'Class'
}