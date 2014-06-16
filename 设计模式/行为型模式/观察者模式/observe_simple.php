<?php
header("Content-type:text/html;charset=utf-8");
/*
 *观察者模式：定义对象间一种一对多的依赖关系（即，多个类(观察者)依赖于某一个类），当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并自动更新。
 *
 *主体就是被观察者、被依赖项，而依赖项就是观察者
 *适用情形：类之间是组合(不是继承)关系的情况。凡是存在多个类依赖于一个类这样的一种依赖关系时，都可以用观察者模式来管理这几个类彼此之间的依赖关系，实现程序的松耦合
 */



//被依赖者
class ConcreteSubject{
     private $_observers;
     public function __construct(){
          $this->_observers = array();  //在主体内部创建一个观察者的对象列表，这个数组的元素都是对象类型
     }
 
     public function add($observer){
          return array_push($this->_observers,$observer);  //新增一个观察者
     }
     
     
 
     public function notifyObservers(){
          if(!is_array($this->_observers)){
			return false;
          }
          foreach($this->_observers as $observer){  //通知每一个观察者
			$observer->update();
          }
          return true;
     }
}
 

//依赖者
class ConcreteObserver{
     private $_name;
     public function __construct($name){
          $this->_name = $name;
     }
 
     public function update(){
          echo '观察者',$this->_name.'已收到更新通知'.PHP_EOL;
     }
}
 



?>