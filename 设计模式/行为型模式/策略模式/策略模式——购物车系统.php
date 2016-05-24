<?php
/**
 * 假设现在要设计一个购物车系统，一个最简单的情况就是把所有货品的单价乘上数量，但是实际情况肯定比这个复杂。
 * 例如：对普通会员无折扣，对银牌会员提供9折优惠，对金牌会员提供8折优惠。
 */
/**
 * 将经常变化的功能独立出来，也可以使用抽象类
 */
interface Strategy{
    public function price($price);//不同级别会员的折扣
}


/**
 * 普通会员的价格策略类
 */
class OrdinaryMember implements Strategy{
    
    
    /**
     * 返回普通会员的折扣值
     * @param void
     * @return int 折扣值
     */
    public function price($price) {
        echo "普通会员无折扣";
        return $price;
    }
}


/**
 * 银牌会员的价格策略类
 */
class SilverMember implements Strategy{
    public function price($price) {
        return 0.9 * $price;
    }
}


/**
 * 金牌会员的价格策略类
 */
class GoldMember implements Strategy{
    public function price($price) {
        return 0.8 * $price;
    }
}

class Pirce{
    //具体的策略类实例
    private $strategyInstace = NULL;
    
    /**
     * 传入具体的策略类对象
     * @param Strategh $instance 具体的策略类对象
     */
    public function __construct(Strategy $instance) {
        $this->strategyInstace = $instance;
    }
    
    public function quote($price){
        return $this->strategyInstace->price($price);
    }
}

/**
 * 客户端操作
 */
$goldMember = new GoldMember(); //可使用工厂来获取策略类
$price = new Pirce($goldMember);
$finalPrice = $price->quote(888);
echo "最终价格：".$finalPrice;


