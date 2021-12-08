// +build windows

package win32api

import (
	"golang.org/x/sys/windows"
)

var (
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
)

var (
	procFlashWindowEx       = moduser32.NewProc("FlashWindowEx")
	procGetForegroundWindow = moduser32.NewProc("GetForegroundWindow")
	procGetActiveWindow     = moduser32.NewProc("GetActiveWindow")
	procShowWindow          = moduser32.NewProc("ShowWindow")

	procGetConsoleWindow = modkernel32.NewProc("GetConsoleWindow")
)

// type FLASHWINFO struct {
// 	cbSize    uint
// 	hwnd      HWND
// 	dwFlags   uint32
// 	uCount    uint
// 	dwTimeout uint32
// }

// var (
// 	FLASHW_TIMERNOFG = 0x0000000C
// 	FLASHW_TRAY      = 0x00000002
// )

// func FlashWindowEx() {
// 	var fwi FLASHWINFO
// 	fwi.cbSize = uint(unsafe.Sizeof(fwi))

// 	fwi.hwnd = GetConsoleWindow()
// 	fwi.dwFlags = uint32(FLASHW_TIMERNOFG | FLASHW_TRAY)
// 	fwi.uCount = 10
// 	fwi.dwTimeout = 1000

// 	r0, _, e0 := procFlashWindowEx.Call(uintptr(unsafe.Pointer(&fwi)))
// 	fmt.Print(r0, e0)
// }

func GetConsoleWindow() HWND {
	ret, _, _ := procGetConsoleWindow.Call()

	return HWND(ret)
}

var isHidden bool = false

func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0

}

const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

// HideConsoleWindow hides the console window
func HideConsoleWindow() {
	ShowWindow(GetConsoleWindow(), SW_HIDE)
	isHidden = true
}

// ShowConsoleWindow shows the console window
func ShowConsoleWindow() {
	ShowWindow(GetConsoleWindow(), SW_SHOW)
	isHidden = false
}

// ToggleShowConsoleWindow toggles the visibility of the console window
func ToggleShowConsoleWindow() {
	if isHidden {
		ShowConsoleWindow()
	} else {
		HideConsoleWindow()
	}
}
