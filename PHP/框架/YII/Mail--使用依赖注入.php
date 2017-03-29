<?php
/**
 * Created by PhpStorm.
 * User: wuxin
 * Date: 2017/3/29
 * Time: 14:27
 */

namespace frontend\components;

use Yii;
use yii\base\Component;


class Mail extends Component{

    //主机地址
    private $_host;

    //端口
    private $_port;

    //用户名
    private $_username;

    //密码
    private $_password;

    //发送者对象
    private $_mailer;

    //邮件标题
    private $_title;

    //邮件正文
    private $_message;

    //邮件内容模板
    private $_template;
    /**
     * 构造函数.
     * @param \PHPMailer $mailer 依赖注入
     * @param array $config
     * Author keven.wu
     * Date 2017-03-28
     * Time 14:33
     * Last modify date
     */
    public function __construct(\PHPMailer $mailer, array $config = [])
    {
        parent::__construct($config);
        $this->_mailer = $mailer;
    }

    /**
     * 设置host属性
     * @param $value string 属性值
     * @return void
     * Author keven.wu
     * Date 2017-03-28
     */
    public function setHost($value){
        $this->_host = $value;
    }

    /**
     * 设置port属性
     * @param $value string 属性值
     * @return void
     * Author keven.wu
     * Date 2017-03-28
     */
    public function setPort($value){
        $this->_port = $value;
    }

    /**
     * 设置username属性
     * @param $value string 属性值
     * @return void
     * Author keven.wu
     * Date 2017-03-28
     */
    public function setUsername($value){
        $this->_username = $value;
    }

    /**
     * 设置password属性
     * @param $value string 属性值
     * @return void
     * Author keven.wu
     * Date 2017-03-28
     */
    public function setPassword($value){
        $this->_password = $value;
    }




    public function send(){

    }

}
$mailConfig = [
    'host' => '',
    'username' => '',
    'password' => '',
];
//构造函数注入
Yii::createObject(Mail::class, [\PHPMailer::class, $mailConfig]);
//属性注入，需要修改Mail类
Yii::createObject([
    'class' => Mail::class,
    'mailer' => \PHPMailer::class,
    'host' => $mailConfig['host']
]);