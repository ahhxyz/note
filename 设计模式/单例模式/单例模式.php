<?php
error_reporting(E_ALL);//生产环境将此值改成0
ini_set('magic_quotes_gpc','on');
ini_set('display_errors', TRUE);
ini_set('display_startup_errors', TRUE);
date_default_timezone_set('PRC');
define('APP_NAME','erp');
define('APP_DEBUG',false);
define('APP_PATH','./../Common/');

require('../Common/ThinkPHP/ThinkPHP.php');
?>
