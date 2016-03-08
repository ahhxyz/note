<?php

interface gun {
    public function shoot();
}

class Handgun implements gun{
    
    public function shoot() {
        echo '手枪';
    }
}


function test(gun $g){
    $g->shoot();
}
