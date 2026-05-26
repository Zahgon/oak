package win32

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	_ = "STUB: not implemented"

	// "The operation completed successfully"
	return nil
}

// TODO: add more here, after collecting data on the common
// error values see on Windows. (perhaps when running
// all.bat?)

var (
	moduser32  = windows.NewLazySystemDLL("user32.dll")
	modshell32 = syscall.NewLazyDLL("shell32.dll")
)

var (
	procShell_NotifyIconW  = modshell32.NewProc("Shell_NotifyIconW")
	procRegisterClass      = moduser32.NewProc("RegisterClassW")
	procIsZoomed           = moduser32.NewProc("IsZoomed")
	procLoadIcon           = moduser32.NewProc("LoadIconW")
	procLoadImageW         = moduser32.NewProc("LoadImageW")
	procLoadCursor         = moduser32.NewProc("LoadCursorW")
	procShowWindow         = moduser32.NewProc("ShowWindow")
	procCreateWindowEx     = moduser32.NewProc("CreateWindowExW")
	procDestroyWindow      = moduser32.NewProc("DestroyWindow")
	procDefWindowProc      = moduser32.NewProc("DefWindowProcW")
	procPostQuitMessage    = moduser32.NewProc("PostQuitMessage")
	procGetMessage         = moduser32.NewProc("GetMessageW")
	procTranslateMessage   = moduser32.NewProc("TranslateMessage")
	procDispatchMessage    = moduser32.NewProc("DispatchMessageW")
	procSendMessage        = moduser32.NewProc("SendMessageW")
	procPostMessage        = moduser32.NewProc("PostMessageW")
	procSetWindowText      = moduser32.NewProc("SetWindowTextW")
	procGetWindowRect      = moduser32.NewProc("GetWindowRect")
	procGetWindow          = moduser32.NewProc("GetWindow")
	procMoveWindow         = moduser32.NewProc("MoveWindow")
	procScreenToClient     = moduser32.NewProc("ScreenToClient")
	procSetWindowLong      = moduser32.NewProc("SetWindowLongW")
	procGetClientRect      = moduser32.NewProc("GetClientRect")
	procGetDC              = moduser32.NewProc("GetDC")
	procReleaseDC          = moduser32.NewProc("ReleaseDC")
	procSetWindowPos       = moduser32.NewProc("SetWindowPos")
	procGetKeyboardLayout  = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState   = moduser32.NewProc("GetKeyboardState")
	procMonitorFromWindow  = moduser32.NewProc("MonitorFromWindow")
	procGetMonitorInfo     = moduser32.NewProc("GetMonitorInfoW")
	procGetKeyState        = moduser32.NewProc("GetKeyState")
	procToUnicodeEx        = moduser32.NewProc("ToUnicodeEx")
	procLoadCursorFromFile = moduser32.NewProc("LoadCursorFromFileW")
	procCreateCursor       = moduser32.NewProc("CreateCursor")
	procSetClassLongPtr    = moduser32.NewProc("SetClassLongPtrW")
	procGetCursorPos       = moduser32.NewProc("GetCursorPos")
)

func _GetKeyboardLayout(threadID uint32) (locale syscall.Handle) {
	_ = "STUB: not implemented"
	return *new(syscall.Handle)
}

func _GetKeyboardState(lpKeyState *byte) (err error) { _ = "STUB: not implemented"; return nil }

func _GetKeyState(virtkey int32) (keystatus int16) { _ = "STUB: not implemented"; return 0 }

func _PostQuitMessage(exitCode int32) { _ = "STUB: not implemented"; return }

func _ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32, dwhkl syscall.Handle) (ret int32) {
	_ = "STUB: not implemented"
	return 0
}

func IsZoomed(hwnd HWND) bool { _ = "STUB: not implemented"; return false }

func RegisterClass(wc *_WNDCLASS) (atom uint16, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func LoadIcon(instance HINSTANCE, iconName uintptr) (HICON, error) {
	_ = "STUB: not implemented"
	return *new(HICON), nil
}

func LoadCursor(instance HINSTANCE, cursorName uintptr) (HCURSOR, error) {
	_ = "STUB: not implemented"
	return *new(HCURSOR), nil
}

func ShowWindow(hwnd HWND, cmdshow int) bool { _ = "STUB: not implemented"; return false }

func CreateWindowEx(exStyle uint32, className, windowName *uint16,
	style uint32, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param uintptr) (HWND, error) {
	_ = "STUB: not implemented"
	return *new(HWND), nil
}

func DestroyWindow(hwnd HWND) bool { _ = "STUB: not implemented"; return false }

const (
	ICON_BIG   = 1
	ICON_SMALL = 0
)

func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) (uintptr, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func TranslateMessage(msg *MSG) bool { _ = "STUB: not implemented"; return false }

func DispatchMessage(msg *MSG) uintptr { _ = "STUB: not implemented"; return 0 }

// SendMessage to the specified window.
// Wrapper around the proc: https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	_ = "STUB: not implemented"
	return 0
}

func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	_ = "STUB: not implemented"
	return false
}

func SetWindowText(hwnd HWND, text string) { _ = "STUB: not implemented"; return }

func GetWindowRect(hwnd HWND) (*RECT, error) { _ = "STUB: not implemented"; return nil, nil }

func MoveWindow(hwnd HWND, x, y, width, height int32, repaint bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ScreenToClient(hwnd HWND, x, y int) (X, Y int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func SetWindowLong(hwnd HWND, index int, value int32) int32 { _ = "STUB: not implemented"; return 0 }

func GetClientRect(hwnd HWND) (*RECT, error) { _ = "STUB: not implemented"; return nil, nil }

func GetDC(hwnd HWND) (HDC, error) { _ = "STUB: not implemented"; return *new(HDC), nil }

func ReleaseDC(hwnd HWND, hDC HDC) bool { _ = "STUB: not implemented"; return false }

func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int32, uFlags uint) bool {
	_ = "STUB: not implemented"
	return false
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	_ = "STUB: not implemented"
	return *new(HMONITOR)
}

func LoadImage(hInst uintptr, name *uint16, type_ uint32, cx, cy int32, fuLoad uint32) (uintptr, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Shell_NotifyIcon(action Shell_NotifyAction, data *NOTIFYICONDATA) bool {
	_ = "STUB: not implemented"
	return false
}

func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	_ = "STUB: not implemented"
	return false
}

func CreateCursor(hinst HINSTANCE, x, y, w, h int32, andMask, xorMask []byte) HCURSOR {
	_ = "STUB: not implemented"
	return *new(HCURSOR)
}

func SetClassLongPtr(hwnd HWND, param ClassLongParam, val uintptr) bool {
	_ = "STUB: not implemented"
	return false
}

type ClassLongParam int32

const (
	// Sets the size, in bytes, of the extra memory associated with the class. Setting this value does not change the number of extra bytes already allocated.
	GCL_CBCLSEXTRA ClassLongParam = -20
	// Sets the size, in bytes, of the extra window memory associated with each window in the class. Setting this value does not change the number of extra bytes already allocated. For information on how to access this memory, see SetWindowLongPtr.
	GCL_CBWNDEXTRA ClassLongParam = -18
	// Replaces a handle to the background brush associated with the class.
	GCLP_HBRBACKGROUND ClassLongParam = -10
	// Replaces a handle to the cursor associated with the class.
	GCLP_HCURSOR ClassLongParam = -12
	// Replaces a handle to the icon associated with the class.
	GCLP_HICON ClassLongParam = -14
	// Retrieves a handle to the small icon associated with the class.
	GCLP_HICONSM ClassLongParam = -34
	// Replaces a handle to the module that registered the class.
	GCLP_HMODULE ClassLongParam = -16
	// Replaces the pointer to the menu name string. The string identifies the menu resource associated with the class.
	GCLP_MENUNAME ClassLongParam = -8
	// Replaces the window-class style bits.
	GCL_STYLE ClassLongParam = -26
	// Replaces the pointer to the window procedure associated with the class.
	GCLP_WNDPROC ClassLongParam = -24
)

func GetCursorPos() (x, y int, ok bool) { _ = "STUB: not implemented"; return 0, 0, false }
