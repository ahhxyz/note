通过使用 Rx.Observable.create 或者是创建操作符，创建一个Observable； 
Observable 被 Observer（观察者） 订阅； 
在执行时 向观察者发送next / error / complete 通知；同时执行过程可以被 终止。 
Observable 类型的实例具备了以上四个方面的特性，与其他类型如：Observer 和 Subscription 紧密相关。

我们重点关注以下四个方面：

创建
订阅
执行
终止









1. Rx.Observable.create()创建的可观察对象就是被观察者
2. 可观察对象的subscribe()用来添加一个观察者并返回
 