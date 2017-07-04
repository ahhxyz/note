<?php
$arr2 = ['a','s','d'];
foreach ($arr2 as $k => &$v){
    echo $k." ".$v."<br>";
}
/*
 * 输出：
    0 a
    1 s
    2 d
 */
//unset($v);
/**
 * 由于变量$v已经在前面初始化并且为引用类型，所以：
 * 1. 第一次循环开始后，$arr已变成了['a','s','a']
 * 2. 第二次循环开始后，$arr已变成了['a','s','s']
 */
foreach ($arr2 as $k => $v){
    echo $k." ";
    echo $v." ".current($arr2)."<br>";
}
/*
 * 输出：
    0 a a
    1 s a
    2 s a
 */

