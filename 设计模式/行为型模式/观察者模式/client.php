<?php
header("Content-type:text/html;charset=utf-8");
include 'observe.php';
//客户端
class Client{
	public static function test(){
            $subject = new ConcreteSubject();  //创建一个主体对象

            //新增第一个观察者
            $observer1 = new ConcreteObserver('奥巴马');
            //将这个观察者注册到主体当中
            $subject->add($observer1);
            //通知
            $subject->notifyObservers();

            //新增第二个观察者
            $observer2 = new ConcreteObserver('山姆');
            $subject->add($observer2);
            //通知
            $subject->notifyObservers();

            //删除观察者1
            $subject->delete($observer1);
            //通知
            $subject->notifyObservers();

	}
}

Client::test()


?>