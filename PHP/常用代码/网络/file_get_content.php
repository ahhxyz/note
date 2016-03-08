<?php
/**
 * 使用file_get_contents
 */
$data = [
    'author' => '张三',
    'email' => '1@1.com'
];
$data = http_build_query($data);
$opts = [
    'http' => [
       'method' => 'POST',
       'header'=> "Content-type: application/x-www-form-urlencoded\r\n" . "Content-Length: " . strlen($data) . "\r\n",
       'content' => $data
    ]
];

$context = stream_context_create($opts);
$html = 
@file_get_contents('http://www.baidu.com', false, $context);
var_dump($html);