<?php
/**
 * 产品类，
 */ 
class Female{}

class Male{}
/*
 *
 */
class Controller{
    //工厂方法
    abstract public function getUser();
}

class 化妆品Controller extends Controller{

    public function getUser(){
        return new Female();
    }
}

class 数码Controller extends Controller{

    public function getUser(){
        return new Male();
    }
}

/*
 * 在实际实现中，需要配合服务定位器模式来确定要实例化的类
 */
?>
