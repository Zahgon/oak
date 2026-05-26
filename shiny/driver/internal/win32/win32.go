// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

// Package win32 implements a partial shiny screen driver using the Win32 API.
// It provides window, lifecycle, key, and mouse management, but no drawing.
// That is left to windriver (using GDI) or gldriver (using DirectX via ANGLE).
package win32

import (
	"sync"

	"github.com/oakmound/oak/v4/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// screenHWND is the handle to the "Screen window".
// The Screen window encapsulates all screen.Screen operations
// in an actual Windows window so they all run on the main thread.
// Since any messages sent to a window will be executed on the
// main thread, we can safely use the messages below.
var screenHWND HWND

const (
	msgCreateWindow = _WM_USER + iota
	msgShow
	msgQuit
	msgLast // WM_USER value https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-user
)

var msgCallbacks = func() *uint32 {
	u := new(uint32)
	*u = msgLast + 1
	return u
}()

// userWM is used to generate private (WM_USER and above) window message IDs
// for use by screenWindowWndProc and windowWndProc.
type userWM struct {
	sync.Mutex
	id uint32
}

// next id for the given userWM (which is a construct purely used to generate unique ids).
func (m *userWM) next() uint32 { _ = "STUB: not implemented"; return 0 }

// currentUserM gives a quick handle to globally mess with userWM.
var currentUserWM userWM

func newWindow(opts screen.WindowGenerator, class string) (HWND, error) {
	_ = "STUB: not implemented"
	return *new(HWND), nil
}

// This should be a feature, putting windows on the top layer

// This is interesting and we'll use it eventually
//SetWindowLongPtr(hwnd, GWL_STYLE, 0)
// TODO(andlabs): use proper nCmdShow
// TODO(andlabs): call UpdateWindow()

// WindowsStyle converts a screen.BorderStyle into a style and
// exStyle for a Windows window
func WindowsStyle(gen screen.WindowGenerator) (uint32, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// ResizeClientRect makes hwnd client rectangle opts.Width by opts.Height in size.
func ResizeClientRect(hwnd HWND, opts screen.WindowGenerator) error {
	_ = "STUB: not implemented"
	return nil
}

// Show shows a newly created window.
// It sends the appropriate lifecycle events, makes the window appear
// on the screen, and sends an initial size event.
//
// This is a separate step from NewWindow to give the driver a chance
// to setup its internal state for a window before events start being
// delivered.
func Show(hwnd HWND) { _ = "STUB: not implemented"; return }

// Release sends the close message to the specified window.
// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-close
func Release(hwnd HWND) { _ = "STUB: not implemented"; return }

// sendFocus change to the specified window.
// There is some value here but the panic is not safe for consumption.
// Consider: wrapper func or rewrite.
func sendFocus(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

func sendShow(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

func sendSizeEvent(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

func sendSize(hwnd HWND) { _ = "STUB: not implemented"; return }

// TODO(andlabs)

// TODO(andlabs): don't assume that PixelsPerPt == 1

func sendClose(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

func sendMouseEvent(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

// TODO: On a trackpad, a scroll can be a drawn-out affair with a
// distinct beginning and end. Should the intermediate events be
// DirNone?

// No-op.

// TODO: handle horizontal scrolling

// Precondition: this is called in immediate response to the message that triggered the event (so not after w.Send).
func keyModifiers() (m key.Modifiers) { _ = "STUB: not implemented"; return *new(key.Modifiers) }

// GetKeyState gets the key state at the time of the message, so this is what we want.

var (
	MouseEvent     func(hwnd HWND, e mouse.Event)
	PaintEvent     func(hwnd HWND, e paint.Event)
	SizeEvent      func(hwnd HWND, e size.Event)
	KeyEvent       func(hwnd HWND, e key.Event)
	LifecycleEvent func(hwnd HWND, e lifecycle.Stage)

	// TODO: use the golang.org/x/exp/shiny/driver/internal/lifecycler package
	// instead of or together with the LifecycleEvent callback?
)

func sendPaint(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

var screenMsgs = map[uint32]func(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr){}

func AddScreenMsg(fn func(hwnd HWND, uMsg uint32, wParam, lParam uintptr)) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func screenWindowWndProc(hwnd HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

//go:uintptrescapes

// SendScreenMessage is a perhaps poorly named wrapper for SendMessage where we know that lParam has a pointer in its call.
// Ends up calling https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew so thats cool.
func SendScreenMessage(screen HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

var windowMsgs = map[uint32]func(hwnd HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr){
	_WM_SETFOCUS:         sendFocus,
	_WM_KILLFOCUS:        sendFocus,
	_WM_PAINT:            sendPaint,
	msgShow:              sendShow,
	_WM_WINDOWPOSCHANGED: sendSizeEvent,
	_WM_CLOSE:            sendClose,

	_WM_LBUTTONDOWN: sendMouseEvent,
	_WM_LBUTTONUP:   sendMouseEvent,
	_WM_MBUTTONDOWN: sendMouseEvent,
	_WM_MBUTTONUP:   sendMouseEvent,
	_WM_RBUTTONDOWN: sendMouseEvent,
	_WM_RBUTTONUP:   sendMouseEvent,
	_WM_MOUSEMOVE:   sendMouseEvent,
	_WM_MOUSEWHEEL:  sendMouseEvent,

	_WM_KEYDOWN:         sendKeyEvent,
	_WM_KEYUP:           sendKeyEvent,
	_WM_INPUTLANGCHANGE: updateKeyboardLayout,
	// TODO case _WM_SYSKEYDOWN, _WM_SYSKEYUP:
}

// AddWindowMsg stores a given window manipulator so it can be accessed via syscalls.
// Stores a reference to the reference argument for the the given id.
func AddWindowMsg(fn func(hwnd HWND, uMsg uint32, wParam, lParam uintptr)) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// src: https://wiki.winehq.org/List_Of_Windows_Messages
// var unusedMessages = map[uint32]string{
// 	2:   "DESTROY",
// 	6:   "ACTIVATE",
// 	28:  "ACTIVATE_APP",
// 	32:  "SETCURSOR",
// 	70:  "WINDOWPOSCHANGING",
// 	130: "NCDESTROY",
// 	132: "NCHITTEST",
// 	134: "NCACTIVATE",
// 	144: "", // we get this, but its not documented in the source list
// 	160: "NCMOUSEMOVE",
// 	161: "NCLBUTTONDOWN",
// 	273: "COMMAND",
// 	274: "SYSCOMMAND",
// 	533: "CAPTURECHANGED",
// 	641: "IME_SETCONTEXT",
// 	642: "IME_NOTIFY",
// 	674: "NCMOUSELEAVE",
// }

func windowWndProc(hwnd HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	_ = "STUB: not implemented"
	return 0
}

//fmt.Printf("unused message %d, 0x%x, %v\n", uMsg, uMsg, unusedMessages[uMsg])

type newWindowParams struct {
	opts  screen.WindowGenerator
	w     HWND
	class string
	err   error
}

var nextWindow = new(int32)

// NewWindow attempts to register a screen on the given handle.
func NewWindow(screenHWND HWND, opts screen.WindowGenerator) (HWND, error) {
	_ = "STUB: not implemented"
	return *new(HWND), nil
}

func initWindowClass(class string) (err error) { _ = "STUB: not implemented"; return nil }

var nextScreenWindow = new(int32)

func initScreenWindow() (HWND, error) { _ = "STUB: not implemented"; return *new(HWND), nil }

var (
	windowStyle uint32 = WS_OVERLAPPEDWINDOW
)

var (
	hDefaultIcon   HICON
	hDefaultCursor HCURSOR
	hThisInstance  HINSTANCE
)

// initCommon attempts to set up some standard icons.
// TODO: Consider running this only once if successful.
func initCommon() (err error) { _ = "STUB: not implemented"; return nil }

// TODO(andlabs) hThisInstance

// Todo: this (and other globals) forces this package to only be able to run one window.
// Can we change this?
var (
	callbacksLock sync.RWMutex
	callbacks     = map[uint32]func(){}
)

// NewScreen sets up common infos and then attempts to create a new window.
func NewScreen() (HWND, error) { _ = "STUB: not implemented"; return *new(HWND), nil }

func Main(screenHWND HWND, f func()) error { _ = "STUB: not implemented"; return nil }

// TODO(andlabs): log an error if this fails?

// TODO(andlabs): unregister window class

// It does not matter which OS thread we are on.
// All that matters is that we confine all UI operations
// to the thread that created the respective window.

// Prime the pump.

// Main message pump.

// WM_QUIT
