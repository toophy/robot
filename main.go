// main.go
package main

import (
	"github.com/toophy/robot/config"
	"github.com/toophy/robot/help"
	"github.com/toophy/robot/logic"
)

// Gogame framework version.
const (
	VERSION = "0.0.2"
)

func main() {
	help.GetApp().Start(config.LogDir, config.ProfFile)

	// 主协程
	go logic.Main_go()

	// 等待结束
	help.GetApp().WaitExit()
}
