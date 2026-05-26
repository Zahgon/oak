package x11

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
)

func MoveWindow(xc *xgb.Conn, xw xproto.Window, x, y, width, height int) (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func ToggleFullScreen(xutil *xgbutil.XUtil, win xproto.Window) error {
	_ = "STUB: not implemented"
	return nil
}

func SetFullScreen(xutil *xgbutil.XUtil, win xproto.Window, fullscreen bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ToggleTopMost(xutil *xgbutil.XUtil, win xproto.Window) error {
	_ = "STUB: not implemented"
	return nil
}

func SetTopMost(xutil *xgbutil.XUtil, win xproto.Window, topMost bool) error {
	_ = "STUB: not implemented"
	return nil
}

func SetBorderless(xutil *xgbutil.XUtil, win xproto.Window, borderless bool) error {
	_ = "STUB: not implemented"
	return nil
}
