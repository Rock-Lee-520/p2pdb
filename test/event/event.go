package main

import (
	"github.com/Rock-liyi/p2pdb/application/event" //注册事件监听
)

func main() {

	// 实例化一个角色
	d := new(event.Driver)

	// Register a callback named OnSkill
	event.RegisterEvent("OnSkill", d.OnEvent)

	// Register the global events on the OnSkill again
	event.RegisterEvent("OnSkill", event.GlobalEvent)

	// Call the event, and all of the registered functions with the same name are called
	event.CallEvent("OnSkill", 100)

}
