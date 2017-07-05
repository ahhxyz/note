# 可观察对象 (Observable)

| MagicQ        | 单值 | 多值   |
| ------------- |:-------:|:-----:|
| 拉取(Pull)      | 函数 | 遍历器|
| 推送(Push) |  Promise  |    Observable |

Observable（可观察对象）是基于推送（Push）运行时执行（lazy）的__多值__集合，为解决基于推送的__多值__问题。

首先，我们需要创建一个可观察对象：

	var observable = Rx.Observable.create(function (observer) {
	  observer.next(1);
	  observer.next(2);
	  observer.next(3);
	  setTimeout(() => {
	    observer.next(4);
	    observer.complete();
	  }, 1000);
	});

为得到observable推送的值，我们需要订阅（subscribe）这个Observable。当observable被订阅后，会立即（同步地）推送1， 2， 3 三个值；1秒之后，继续推送4这个值，最后结束（推送结束通知）：

	
	console.log('just before subscribe');
	observable.subscribe({
	  next: x => console.log('got value ' + x),
	  error: err => console.error('something wrong occurred: ' + err),
	  complete: () => console.log('done'),
	});
	console.log('just after subscribe');
	
程序执行后，将在控制台输出如下结果：
	
	just before subscribe
	got value 1
	got value 2
	got value 3
	just after subscribe
	got value 4
	done
# angular中的可观察对象 (Observable)

Http服务中的每个方法都返回一个 HTTP Response对象的Observable实例。

我们的HeroService中把那个Observable对象转换成了Promise（承诺），并把这个承诺返回给了调用者。 这一节，我们将学会直接返回Observable，并且讨论何时以及为何那样做会更好。


转换成承诺通常是更好地选择，我们通常会要求http.get()获取`单块`数据。`只要接收到数据，就算完成`。 使用承诺这种形式的结果是让调用方更容易写，并且承诺已经在 JavaScript 程序员中被广泛接受了。

__但是请求并非总是“一次性”的。我们可以开始一个请求， 并且取消它，在服务器对第一个请求作出回应之前，再开始另一个不同的请求 。 像这样一个请求-取消-新请求的序列对于Promise来说是很难实现的，但是对Observable来说则很容易。__