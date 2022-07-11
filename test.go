package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
发布者
事件驱动架构是计算机科学家中一种高度可扩展年的范例,它允许我们可以多方系统异步处理事件
事件总线是 发布/ 订阅模式的实现,其中发布者发布数据,并且感兴趣的订阅者可以监听这些数据并基于
这些数据作出处理,是发布者与订阅者松耦合,发布者将数据事件发布到事件总线,总线负责将他们发送给订阅者
传统的实现事件总线的方法会涉及到使用回调,订阅者通常实现接口,然后事件总线通过接口传播数据
使用Go的并发模型,大多数地方可以使用channel来代替回调,在本文中,我们将重点介绍如何使用channle 来实现事件总线
基于主题 topic 的事件 发布者发布到主题,订阅者可以收听他们
*/
// 定义数据结构
// 实现事件总线,我们需要定义要传递数据结构,
// 我们可以使用struct 简单地创建一个新的数据类型,定义一个DataEvent的结构体如下.
// 这里 我们已经将基础数据定义为接口,这意味着他可以是任何值,我们还将主题定义为结构的成员,
// 订阅者可能会收听多个主题,因此,我们将通过主题让订阅者可以区分不同的事件做法 is good

type DataEvent struct {
	Data  interface{}
	Topic string
}

// 为事件总线定义了我们主要的数据结构,我们还需要一种方法传递它,
// 为此,我们可以定义一个可以传播DataEvent 和DataChannel 类型
//DataChannel 是一个能接受DataEvent的channel

type DataChannel chan DataEvent

// DataChannelSlice是一个包涵DataChannels 数据切片
//DataChannelSlice的创建是为了保留DataChannel 的切片并轻松引用他们
type DatachannelSlice []DataChannel

// 事件总线
// EventBus 存储有关订阅者感兴趣的特定主题的信息
/**
EventBus有 subscribes, 这是一个包涵DataChannelSlice 的map 我们使用互斥锁来保护并发访问的读写
通过使用map 和定义topics,它允许我们轻松组织事件.主题被视为map的键,当有人发布它时,
可以通过键轻松找到主题,然后将事件传播到channnel 中以进行异步处理
*/
type EventBus struct {
	subscribers map[string]DatachannelSlice
	rm          sync.RWMutex
}

// 订阅主题
/**
对于订阅主题,使用channle 它像传统方法中的回调一样,当发布者向主题发布数据时channel 将接收数据
简单地说,我们将订阅者添加到channle 切片中然后给该结构加锁,最后在爱操作后将其解锁
*/
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

// 发布主题
/**
 发布事件,发布者需要提供广播给订阅者所需的主题和数据
在次方法中,首先问你检查主题是否存在任何订阅者
然后我们只是简单遍历与主题有关的channel 切片 把事件发布给他们
注: 在发布方法中使用了Goroutine 来避免阻塞发布者
*/
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		/**
			这样做是因为切片引用相同的数组,即使他们是按值传递的
		因此我们在使用我们的元素创建一个新切片,从而能正确地保持锁定
		*/
		channels := append(DatachannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlice DatachannelSlice) {
			for _, ch := range dataChannelSlice {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent %v\n", ch, data.Topic, data.Data)
}

func publishersTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

var eb = &EventBus{
	subscribers: map[string]DatachannelSlice{},
}

// 测试
/**
我们需要创建一个事件总线的实例,在实际场景中,
可以从包中到处单个EventBus,使其像单例一样运行
*/
func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic3", ch3)

	go publishersTo("topic1", "Hitopic 1")
	go publishersTo("topic2", "Welcome to topic2")
	go publishersTo("topic3", "Welcome to topic3")
	d := <-ch1
	printDataEvent("ch1", d)

	d2 := <-ch2
	printDataEvent("ch2", d2)

	d3 := <-ch3
	printDataEvent("ch3", d3)

	// for {
	// 	select {
	// 	case d := <-ch1:
	// 		go printDataEvent("ch1", d)
	// 	case d := <-ch2:
	// 		go printDataEvent("ch2", d)
	// 	case d := <-ch3:
	// 		go printDataEvent("ch3", d)
	// 	}
	// }
}
