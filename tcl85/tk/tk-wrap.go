package tk

/*
#cgo CFLAGS: -I/usr/include/tcl8.5
#cgo LDFLAGS: -ltk8.5
#include <tk.h>
*/
import "C"

import (
	"github.com/grd/tg/tcl85/tcl"
	"unsafe"
)

func getInterp() (interp *tcl.Interp, err error) {
	interp = tcl.CreateInterp()

	err = tcl.Init(interp)
	if err != nil {
		return
	}

	err = InitTk(interp)
	if err != nil {
		return
	}

	err = tcl.Eval(interp, "namespace eval gt {}")

	return
}

// for initializing Tcl environment
func Init(arg string) (interp *tcl.Interp, err error) {
	tcl.FindExecutable(arg)
	return getInterp()
}

func InitTk(interp *tcl.Interp) error {
	ci := (*C.Tcl_Interp)(unsafe.Pointer(interp))
	i := C.Tk_Init(ci)
	if i != tcl.OK {
		return interp.Error()
	}
	return nil
}

// The event loop
func Run() {
	C.Tk_MainLoop()
}

func Cmd(interp *tcl.Interp, script string) error {
	return tcl.Eval(interp, script)
}

// Widgets

func Button(options string) {
	// ".b" -text "Say Hello" -command hello
}

func Pack(options string) {
	// ".b" -padx 20 -pady 6
}
