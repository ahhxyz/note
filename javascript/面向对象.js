/*
 * 使用构造函数来定义类的属性及属性的初始化
 */
function Class(){
    this.property01 = '';
    this.property02 = '';
    //this.init() //实例化时要执行的方法
}

/*
 * 使用原型来定义实例的共享成员，即方法和共享的成员变量
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