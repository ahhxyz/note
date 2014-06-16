<?php 
/**
 * 里氏替换原则
 * 组合模式
 */
abstract class MenuComponent{
    public $name;

    public abstract function getName();

    public abstract function add(MenuComponent $menu);

    public abstract function remove(MenuComponent $menu);

    public abstract function getChild($i);

    public abstract function show();
}

class MenuItem extends MenuComponent{
    public function __construct($name){
        $this->name = $name;
    }
    public function getName(){
        return $this->name;
    }
    public function add(MenuComponent $menu){
        return false;
    }
    public function remove(MenuComponent $menu){
        return false;
    }

    public function getChild($i){
        return null;
    }
    public function show(){
        echo " |-".$this->getName()."\n";
    }
}

class Menu extends MenuComponent{
    public $menuComponents = array();
    public function __construct($name){
        $this->name = $name;
    }

    public function getName(){
        return $this->name;
    }

    public function add(MenuComponent $menu){
        $this->menuComponents[] = $menu;
    }

    public function remove(MenuComponent $menu){
        $key = array_search($menu, $this->menuComponents);
        if($key !== false){
            unset($this->menuComponents[$key]);
        }    
    }

    public function getChild($i){
        if(isset($this->menuComponents[$i])){
            return $this->menuComponents[$i];
        }
        return null;
    }

    public function show(){
        echo "|-" . $this->getName() . "\n";
        foreach($this->menuComponents as $v){
            $v->show();
        }
    }
}

class client{
 public function run(){
  $menu1 = new Menu('经部');
  $menuItem1 = new MenuItem('论语');
  $menuItem2 = new MenuItem('大学');
  $menuItem3 = new MenuItem('尚书');
  $menuItem4 = new MenuItem('礼记');
  
  $menu1->add($menuItem1);
  $menu1->add($menuItem2);
  $menu1->add($menuItem3);
  $menu1->add($menuItem4);
  $menu1->show();
  $menu1->remove($menuItem4);
  echo PHP_EOL."删除经部中的《礼记》：".PHP_EOL;
  $menu1->show();
  
  echo PHP_EOL."给《经部》创建一个上级分类：四库全书".PHP_EOL;
  $menu2 = new Menu("四库全书");
  $menu2->add($menu1);
  $menu2->show();
  }
}


$test = new client();
$test->run();
?>