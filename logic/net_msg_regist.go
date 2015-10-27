package logic

import (
	"github.com/toophy/robot/help"
)

func RegMsgProc() {
	help.GetApp().RegMsgFunc(1, on_c2g_login)
}
