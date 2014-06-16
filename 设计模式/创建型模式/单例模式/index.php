<?php
class db{
	private $dbName;
    private static $instance=false;
    public $name;
    private function __construct($dbName){
    	$this->dbName=$dbName;
    }
	
    public static function getDb($dbName){

		if(self::$instance==false){

			self::$instance=new db($dbName);
			echo "创建新实例！";
		}
		return self::$instance;
    }
}
?>