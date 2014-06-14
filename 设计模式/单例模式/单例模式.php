<?php
class Config{
/*-----------------必须部分-----------------------------------------------------*/
    static private $instance=NULL;


    private function __construct() {
        
    }
    
    private function __clone() {}
    
    static function getInstance(){
        if(self::$instance=NULL){
            self::$instance=new Config;
        }
        
        return self::$instance;
        
    }
/*-------------------------------------------------------------------------------/
 * ---------------------END 必须部分----------------------------------------------/
 *----------------------------------------------------------------------------- */
    
    
    /**
     * 该方法是实例方法，仅作示范
     */
    public function demo(){
        echo "demo";
    }
}
//单例的Config类一律通过Config::getInstance()来调用Config类的实例方法。
Config::getInstance()->demo();
?>
