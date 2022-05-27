
<?php

class Message{
    public $id;     //事件唯一id， 发送一次递增一次
    public $timestamp;   //系统毫秒时间戳
    public $msg;    //发生消息内容
    public $nodeId; //节点id
    public $hostId; //主机id
    public $lastId; //因果关系事件id,为0代表没有因果关系
    public $type; //事件type sending or receiving
    public $receivingId; //事件接收时,唯一id
}



/**
 * 比较逻辑时钟,返回后发生的事件id
 */
function compareClock(array $a, array $b){
    /*
     *同一个节点发出的,事件具有因果关系
     *先发出的在前，比较事件id
    */
    if($a['nodeId']==$b['nodeId']){
         $maxId=getMaxId($a['id'],$b['id']);
         return $maxId;
    }

    /*
     *不同节点发出的消息，timestamp相等表示出现并发,按照指定规则收敛
     *消息发出的比接收的大
     */
    if($a['timestamp']==$b['timestamp']){

        if($a['type']=="sending" && $b['timestamp']=="receiving"){
            $a['timestamp'] =  $a['timestamp']+1;  
               
        }

        if($a['type']=="receiving" && $b['timestamp']=="sending"){
            $b['timestamp'] =  $b['timestamp']+1;   
           
        }

    }

    if($a['timestamp']==$b['timestamp']){
         /**
         * 不同节点，都是发送 或者都是接收的请求
         * 按照last-writer-win 收敛
         */
        $maxId=getMaxIdByRid($a['receivingId'],$b['receivingId'],$a['id'],$b['id']);
        return $maxId;
    }else{
        $maxId=getMaxIdByTimestamp($a['timestamp'],$b['timestamp'],$a['id'],$b['id']);
        return $maxId; 
    }
}



