<?php
//分布式加减计数器
class Counter{
    public $count=0;

    public function  increment(){
        $this->count++;
    }

    public function  decrement(){
        $this->count--;
    }

    public function getCount(){
        return $this->count;
    }
}


//声明节点A
$nodeA=new Counter();
//声明节点B
$nodeB=new Counter();

$nodeA->increment();
$nodeA->increment();
$nodeA->increment();

$nodeB->decrement();

$count= $nodeA->getCount()+$nodeB->getCount();

var_dump($nodeA->getCount());
//int(3)
var_dump($nodeB->getCount());
//int(-1)
var_dump($count);
//int(2)


