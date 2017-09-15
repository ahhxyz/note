
var promise = $.Deferred().resolve('随便什么数据');   //$.Deferred().resolve('随便什么数据').promise();


promise.then(function(data){
    var _def = $.Deferred();
    setTimeout(function(){
        console.log(1);

        _def.resolve('随便什么数据');

    }, 2000);
    return _def.promise();  //这是最关键的地方, then的回调函数必须继续返回一个promise，这样才能保证该then执行完后再执行后续的then

})
.then(function(data){
    var _def = $.Deferred();
    setTimeout(function(){
        console.log(2);

        _def.resolve('随便什么数据');

    }, 1000);
    return _def.promise();
})
.then(function(data){
    var _def = $.Deferred();
    setTimeout(function(){
        console.log(3);

        _def.resolve('随便什么数据');

    }, 500);
    return _def.promise();
});



/**
 * 第二种方法
 */

function runAsync(){
    var def = $.Deferred();
    //做一些异步操作
    setTimeout(function(){
        console.log('执行完成');
        def.resolve('随便什么数据');
    }, 1000);
    return def.promise(); //就在这里调用
}

var d = runAsync();

d.then(function(data){
    console.log(1, data);
    return runAsync();
})
.then(function(data){
    console.log(2, data);
    return runAsync();
})
.then(function(data){
    console.log(3, data);
});