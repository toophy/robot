package logic

import (
	"github.com/toophy/robot/help"
)

func Main_go() {
	RegMsgProc()

	go help.GetApp().Connect("to_gate", "tcp", "127.0.0.1:8001", OnListenRet)
}

func OnListenRet(typ string, name string, id int, info string) bool {
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

		c := help.GetApp().GetConnById(id)

		msg_len := 0
		var msg help.Ty_net_msg
		var stream help.Ty_msg_stream
		msg.InitNetMsg()
		stream.InitMsgStream(&msg)
		stream.WriteU2(msg_len)

		var c2g_login int = 1
		stream.WriteU2(c2g_login)
		acc := "古老茅"
		stream.WriteString(&acc)
		msg.Send(c.Conn)

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
