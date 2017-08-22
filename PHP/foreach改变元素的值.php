<?php
$arr = ['a','s','d'];
foreach ($arr as $k => $v){
    $arr[$k] = $v . '_changed';
    var_dump($arr);
}


