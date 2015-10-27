// main.go
package main

import (
	"github.com/toophy/robot/config"
	"github.com/toophy/robot/help"
)

// Gogame framework version.
const (
	VERSION = "0.0.2"
)

func main() {
	help.GetApp().Start(config.LogDir, config.ProfFile)

	// 主协程
	go main_go()

	// 等待结束
	help.GetApp().WaitExit()
}

func main_go() {
	RegMsgProc()

	go help.GetApp().Listen("main_listen", "tcp", ":8001", OnListen)
}

func OnListen(typ string, name string, id int, info string) bool {
	switch typ {
	case "listen failed":
		println(name + " : Listen failed[" + info)

	case "listen ok":
		println(name + " : Listen ok.")

	case "accept failed":
		println(info)
		return false

	case "accept ok":
		println(info)

	case "connect failed":
		if len(name) > 0 {
			println(name + " : Connect failed[" + info)
		} else {
			println("Conn[", id, "] : Connect failed["+info)
		}

	case "connect ok":
		if len(name) > 0 {
			println(name + " : Connect ok")
		} else {
			println("Conn[", id, "] : Connect ok")
		}

	case "read failed":
		if len(name) > 0 {
			println(name + " : Connect read[" + info)
		} else {
			println("Conn[", id, "] : Connect read["+info)
		}

	case "pre close":
		if len(name) > 0 {
			println(name + " : Connect pre close")
		} else {
			println("Conn[", id, "] : Connect pre close")
		}

	case "close failed":
		if len(name) > 0 {
			println(name + " : Connect close failed[" + info)
		} else {
			println("Conn[", id, "] : Connect close failed["+info)
		}

	case "close ok":
		if len(name) > 0 {
			println(name + " : Connect close ok.")
		} else {
			println("Conn[", id, "] : Connect close ok.")
		}
	}

	return true
}

func RegMsgProc() {
	help.GetApp().RegMsgFunc(1, on_c2g_login)
}

func on_c2g_login(c *help.ClientConn) {
	if c.Id > 0 {
		name := c.Stream.ReadStr()
		println(name)
	}
}
