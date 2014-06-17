<?php
/**
 * 组合模式
 * 场景：在游戏中敌我双方对战时，我们需要评估战斗结果：胜利还是失败。
 * 那么我们就要首先获取双方的对应建制的战斗值。
 * 而每一方的战斗值是不同级别战斗单位战斗值的和。
 * 
 * 
 * 该抽象类的抽象方法是分支节点和叶节点的功能的并集
 * 各个子节点根据自身来决定如何实现这些抽象方法，也就是如何组合这些抽象方法。
 */
abstract class FightUnit{
    protected $value=NULL;
    protected $name=NULL;//战斗单位名称
    protected $values=NULL;
    
    
    public function __construct($name,$value=NULL) {
        $this->name=$name;
        $this->value=$value?$value:NULL;
    }
    abstract function valueCount();//返回对应战斗单位的总战斗值
    abstract function add(FightUnit $unit);
    //abstract function attact();
    //abstract function fightResult();
}


/**
 * 作战单位：步兵
 */

class Soldier extends FightUnit{
    
    public function add(FightUnit $unit) {
        return $unit;
    }
    
    public function valueCount() {
        return $this->value;
    }
}

/**
 * 作战单位：特种兵
 *
 */


class SpecialSoldier extends FightUnit{
    public function add(FightUnit $unit){
        return false;
    }
    
    public function valueCount() {
        return $this->value;
    }
}


/**
 * 作战单位：班
 */

class  Squads extends FightUnit{
    public $units=array();


    public function add(FightUnit $unit) {
        array_push($this->units, $unit);
    }
    
    public function valueCount(){
        $value=0;
        foreach ($this->units as $unit){
            $value+=$unit->value;
        }
        return $value;
    }
    
    public function removeUnit(){
        
    }
    
}

$步兵甲=new Soldier("步兵甲",100);
$步兵乙=new Soldier("步兵乙",100);
$特种兵甲=new SpecialSoldier("特种兵甲",1000);

$战斗小组=new Squads("战斗小组甲");
$战斗小组->add($步兵甲);
$战斗小组->add($步兵乙);
echo PHP_EOL."战斗小组甲的战斗力是：".$战斗小组->valueCount();
$战斗小组->add($特种兵甲);
echo PHP_EOL."增加一个特种兵后，战斗小组甲的战斗力是：".$战斗小组->valueCount();

