<?php

$post_ =array (
	'author' => 'Gonn',
	'mail'=>'gonn@nowamagic.net',
	'url'=>'http://www.nowamagic.net/',
	'text'=>'欢迎访问简明现代魔法');



/**
 * 使用curl
 */



$data=http_build_query($post_);

/**
 * 使用fsockopen
 */

$fp = fsockopen("nowamagic.net", 80, $errno, $errstr, 5);

$out="POST http://nowamagic.net/news/1/comment HTTP/1.1\r\n";
$out.="Host: typecho.org\r\n";
$out.="User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; zh-CN; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13"."\r\n";
$out.="Content-type: application/x-www-form-urlencoded\r\n";
$out.="PHPSESSID=082b0cc33cc7e6df1f87502c456c3eb0\r\n";
$out.="Content-Length: " . strlen($data) . "\r\n";
$out.="Connection: close\r\n\r\n";
$out.=$data."\r\n\r\n";

fwrite($fp, $out);//发送POST请求

while (!feof($fp))
{
    echo fgets($fp, 1280);
}

fclose($fp);


/**
 * 使用file_get_contents
 */
$opts = [
	'http' => [
	   'method' => 'POST',
	   'header'=> "Content-type: application/x-www-form-urlencoded\r\n" . "Content-Length: " . strlen($data) . "\r\n",
	   'content' => $data
        ]
];

$context = stream_context_create($opts);
file_get_contents('http://nowamagic.net/news/1/comment', false, $context);


