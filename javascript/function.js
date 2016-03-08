


/**
 * 函数的返回值
 */

//返回值为基本数据类型
var v = 1;
function rtn(){
    return v;
}
var v2 = rtn();
v2 = 2;
console.log(v); //1

//返回值为引用类型
var v = {name : '张三'};
function rtn(){
    return v;
}
var v2 = rtn();
v2.name = '李四';
console.log(v.name); //李四