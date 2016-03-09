<?php
/*
 * 观察者模式实际上一种事件系统
 * 场景设计
    1. 设计一个订单类
    2. 订单创建完成后，会做各种动作，比如发送EMAIL，或者改变订单状态等等。
    原始的方法，是将这些操作都写在create函数里面
    但是随着订单创建类的越来越庞大，这样的操作已经无法满足需求和快速变动
    这个时候，观察者模式出现了。
 *  
 */
//观察者设计模式能够更便利创建和查看目标对象状态的对象，并且提供和核心对象非耦合的置顶功能性。
//观察者设计模式非常常用，在一般复杂的WEB系统中，观察者模式可以帮你减轻代码设计的压力，降低代码耦合。
//以一个购物流程为例子

//抽象主体
interface Subject{
    public function addObserver(Observer $observer);
    public function deleteObserver(Observer $observer);
    public function notifyObservers();
}

class Order implements Subject{

    protected $_observers = []; // 存放观察容器

    //观察者新增
    public function addObserver(Observer $observer) {
        $this->_observers[] = $observer;
    }

    //运行观察者
    public function notifyObservers() {
        foreach($this->_observers as $observer){  //通知每一个观察者
            $observer->update();
        }
        return true;
    }

    public function deleteObserver(Observer $observer){
        
    }
    
    //下单购买流程
    public function create() {
        echo '购买成功';
        $this->addObserver(new Email);
        $this->notifyObservers();
    }
    
    /**
     * 支付
     */
    public function pay($id){
        echo '支付成功';
        $this->addObserver(new Email);
        $this->addObserver(new Finance);
        $this->addObserver(new Storage);    //添加一个仓库对象，以通知其发货
        $this->notifyObservers();
    }
}

//抽象观察者
interface Observer{
     public function update(Subject $subject);
}

class Email implements Observer{
    public  function update(Subject $subject) {
        $this->send($address, $title, $content);
    }
    
    
    /**
     * 发送邮件
     * @param string $address 邮件地址
     * @param string $title 邮件标题
     * @param string $content 邮件内容
     * @param resource $attachment 附件
     */
    public function send($address, $title, $content, $attachment = NULL){
        
    }
}


class Storage implements Observer{
    
    public function update(\Subject $subject) {
        $this->delivery($orderID);
    }
    
    /**
     * 发货
     * @param int $orderID 订单ID
     */
    public function delivery($orderID){
        
    }
    
}

/**
 * 财务类
 */
class Finance implements Observer {
    public  function update($order) {
        echo '改变订单状态';
        $this->record();    //通知财务入账
    }
    
    /**
     * 入账
     */
    public function record(){
        
    }
}
$ob = new order;
$ob->create();
$ob->pay($id);