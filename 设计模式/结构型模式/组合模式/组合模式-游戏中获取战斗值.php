<?php
/**
 * 策略模式
 * 游戏中获取不同级别战斗单位的战斗值
 * 
 */
abstract class FightUnit{
    protected $combatValue=100;//单兵的战斗值，就是叶节点的值
    public function combatValue();
}

/**
 * 作战单位：单兵
 *
 */


class Soldier extends FightUnit{
    public function combatValue(){
        return $this->combatValue;
    }
}


/**
 * 作战单位：班
 */

class Company extends FightUnit{
    public $soldiers=array();


    public function combatValue() {
        ;
    }
    
    public function addUnit(){
        
    }
    
    public function removeUnit(){
        
    }
    
}