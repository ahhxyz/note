
var promise = $.Deferred().resolve('随便什么数据');   //$.Deferred().resolve('随便什么数据').promise();


promise.then(function(data){
    var _def = $.Deferred();
    setTimeout(function(){
        console.log(1);

        _def.resolve('随便什么数据');

    }, 2000);
    return _def.promise();  //这是最关键的地方, 必须返回一个promise

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