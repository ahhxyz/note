<?php

class Stack {
    public $top;
    public $data;
    
    public function __construct() {
        $this->top = -1;
        $this->data = [];
    }
    
    public function push($item){
        $this->data[++$this->top] = $item;
    }
}
