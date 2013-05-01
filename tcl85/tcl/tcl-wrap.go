// Wrapper functions for the Tcl C library.
package tcl

/*
#cgo CFLAGS: -I/usr/include/tcl8.5
#cgo LDFLAGS: -ltcl8.5
#include <stdlib.h>
#include <tcl.h>
#include "tcl-wrap.h"
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

const (
	OK          = C.TCL_OK
	ERROR       = C.TCL_ERROR
	RETURN      = C.TCL_RETURN
	BREAK       = C.TCL_BREAK
	CONTINUE    = C.TCL_CONTINUE
	RESULT_SIZE = C.TCL_RESULT_SIZE
)

//Types for linked variables
const (
	LINK_INT       = C.TCL_LINK_INT
	LINK_DOUBLE    = C.TCL_LINK_DOUBLE
	LINK_BOOLEAN   = C.TCL_LINK_BOOLEAN
	LINK_STRING    = C.TCL_LINK_STRING
	LINK_WIDE_INT  = C.TCL_LINK_WIDE_INT
	LINK_CHAR      = C.TCL_LINK_CHAR
	LINK_UCHAR     = C.TCL_LINK_UCHAR
	LINK_SHORT     = C.TCL_LINK_SHORT
	LINK_USHORT    = C.TCL_LINK_USHORT
	LINK_UINT      = C.TCL_LINK_UINT
	LINK_LONG      = C.TCL_LINK_LONG
	LINK_ULONG     = C.TCL_LINK_ULONG
	LINK_FLOAT     = C.TCL_LINK_FLOAT
	LINK_WIDE_UINT = C.TCL_LINK_WIDE_UINT
	LINK_READ_ONLY = C.TCL_LINK_READ_ONLY
)

const (
	VOLATILE = 1
	STATIC   = 0
	DYNAMIC  = 3
)

type Command C.Tcl_Command
type FreeProc C.Tcl_FreeProc
type Interp C.Tcl_Interp
type Obj C.Tcl_Obj
type ClientData unsafe.Pointer // Arbitrary client data

type Callbacker interface {
	Func(int, *Interp, []Obj) int
	DeleteFunc(int)
}

// Only used for the internal C callback funcs
var callback Callbacker

func (p *Interp) Result() string {
	return C.GoString(p.result)
}

func (p *Interp) Error() error {
	return errors.New(C.GoString(p.result))
}

func CreateInterp() *Interp {
	return (*Interp)(C.Tcl_CreateInterp())
}

func CreateObjCommand(interp *Interp, cmdName string,
	clientData int, cb Callbacker) Command {

	C.interp = (*C.Tcl_Interp)(interp)
	C.cmd_name = C.CString(cmdName)
	defer C.free(unsafe.Pointer(C.cmd_name))
	C.clientData = (C.ClientData)(unsafe.Pointer(&clientData))
	callback = cb

	C.createObjCmd()
	return Command(C.command)
}

func DeleteCommand(interp *Interp, cmdName string) error {
	cc := C.CString(cmdName)
	defer C.free(unsafe.Pointer(cc))

	res := C.Tcl_DeleteCommand((*C.Tcl_Interp)(interp), cc)
	if res != OK {
		return interp.Error()
	}
	return nil
}

func DecrRefCount(objPtr *Obj) {
	C.decrRefCount((*C.Tcl_Obj)(objPtr))
}

func Eval(interp *Interp, script string) error {
	cs := C.CString(script)
	defer C.free(unsafe.Pointer(cs))
	cinterp := (*C.Tcl_Interp)(interp)
	res := C.Tcl_Eval(cinterp, cs)
	if res != OK {
		return interp.Error()
	}
	return nil
}

func FindExecutable(arg string) {
	argv0 := C.CString(arg)
	C.Tcl_FindExecutable(argv0)
	C.free(unsafe.Pointer(argv0))
}

func GetDoubleFromObj(interp *Interp, objPtr *Obj) (float64, error) {
	var double C.double
	ct := C.Tcl_GetDoubleFromObj((*C.Tcl_Interp)(interp), (*C.Tcl_Obj)(objPtr), &double)
	if ct != OK {
		return 0, interp.Error()
	}
	return float64(double), nil
}

func GetIntFromObj(interp *Interp, objPtr *Obj) (int, error) {
	var cint C.int
	ct := C.Tcl_GetIntFromObj((*C.Tcl_Interp)(interp), (*C.Tcl_Obj)(objPtr), &cint)
	if ct != OK {
		return 0, interp.Error()
	}
	return int(cint), nil
}

func GetObjResult(interp *Interp) *Obj {
	ci := (*C.Tcl_Interp)(interp)
	return (*Obj)(C.Tcl_GetObjResult(ci))
}

func GetString(objPtr *Obj) string {
	return C.GoString(C.Tcl_GetString((*C.Tcl_Obj)(objPtr)))
}

func Init(interp *Interp) error {
	ct := C.Tcl_Init((*C.Tcl_Interp)(interp))
	if ct != OK {
		return interp.Error()
	}
	return nil
}

func LinkVar(interp *Interp, varName, addr string, Type int) error {
	cvn := C.CString(varName)
	defer C.free(unsafe.Pointer(cvn))
	caddr := C.CString(addr)
	defer C.free(unsafe.Pointer(caddr))

	ct := C.Tcl_LinkVar((*C.Tcl_Interp)(interp), cvn, caddr, C.int(Type))
	if ct != OK {
		return interp.Error()
	}
	return nil
}

func ListObjIndex(interp *Interp, listPtr *Obj, index int) (*Obj, error) {
	var objPtrPtr *C.Tcl_Obj
	ct := C.Tcl_ListObjIndex((*C.Tcl_Interp)(interp), (*C.Tcl_Obj)(listPtr), C.int(index), &objPtrPtr)
	if ct != OK {
		return nil, interp.Error()
	}
	return (*Obj)(objPtrPtr), nil
}

func ListObjLength(interp *Interp, listPtr *Obj) (int, error) {
	var cint C.int
	ct := C.Tcl_ListObjLength((*C.Tcl_Interp)(interp), (*C.Tcl_Obj)(listPtr), &cint)
	if ct != OK {
		return 0, interp.Error()
	}
	return int(cint), nil
}

func Merge(args []string) string {
	argc := len(args)
	argv := make([](*C.char), argc)

	for i := range args {
		argv[i] = C.CString(args[i])
	}

	cres := C.Tcl_Merge(C.int(argc), (**C.char)(&argv[0]))

	res := C.GoString(cres)

	C.Tcl_Free(cres)

	for i := range args {
		C.free(unsafe.Pointer(argv[i]))
	}

	return res
}

func NewBooleanObj(value bool) *Obj {
	var cval C.int // "false"
	if value {
		cval = 1
	} // "true"

	return (*Obj)(C.Tcl_NewBooleanObj(cval))
}

func NewDoubleObj(value float64) *Obj {
	return (*Obj)(C.Tcl_NewDoubleObj(C.double(value)))
}

func NewLongObj(value int64) *Obj {
	return (*Obj)(C.Tcl_NewLongObj(C.long(value)))
}

func NewStringObj(data string) *Obj {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	cl := C.int(len(data))
	return (*Obj)(C.Tcl_NewStringObj(cs, cl))
}

func SetObjResult(interp *Interp, objPtr *Obj) {
	C.Tcl_SetObjResult((*C.Tcl_Interp)(interp), (*C.Tcl_Obj)(objPtr))
}

func SetResult(interp *Interp, result string, freeProc *FreeProc) {
	cs := C.CString(result)
	defer C.free(unsafe.Pointer(cs))
	C.Tcl_SetResult((*C.Tcl_Interp)(interp), cs, (*C.Tcl_FreeProc)(unsafe.Pointer(freeProc)))
}

func SplitList(list string) []string {
	var argc C.int
	var argv **C.char
	var v []string
	cl := C.CString(list)
	defer C.free(unsafe.Pointer(cl))

	if len(list) == 0 {
		return nil
	}

	if C.Tcl_SplitList(nil, cl, &argc, &argv) != OK {
		// Not a list. Could be a quoted string containing funnies, e.g. {"}.
		return nil
	}

	ref := &reflect.SliceHeader{uintptr(unsafe.Pointer(argv)), int(argc), int(argc)}
	args := *(*[]*C.char)(unsafe.Pointer(ref))

	if argc == 0 {
		v = []string{""}
	} else if argc == 1 {
		v = []string{C.GoString(args[0])}
	} else {
		v = make([]string, int(argc))
		for i := 0; i < int(argc); i++ {
			v[i] = C.GoString(args[i])
		}
	}
	C.Tcl_Free((*C.char)(unsafe.Pointer(argv)))
	return v
}

func UnlinkVar(interp *Interp, varName string) {
	cs := C.CString(varName)
	defer C.free(unsafe.Pointer(cs))
	C.Tcl_UnlinkVar((*C.Tcl_Interp)(interp), cs)
}

func UpdateLinkedVar(interp *Interp, varName string) {
	cs := C.CString(varName)
	defer C.free(unsafe.Pointer(cs))
	C.Tcl_UpdateLinkedVar((*C.Tcl_Interp)(interp), cs)
}

//export wrapObjCmdProc
func wrapObjCmdProc() {
	clientData := (*int)(unsafe.Pointer(C.wrapClientData))
	interp := (*Interp)(C.wrapInterp)

	h := &reflect.SliceHeader{uintptr(unsafe.Pointer(C.wrapObjv)), int(C.wrapObjc), int(C.wrapObjc)}
	obj := *(*[]Obj)(unsafe.Pointer(h))

	ret := callback.Func(*clientData, interp, obj)
	C.wrapReturnVal = C.int(ret)
}

//export wrapCmdDeleteProc
func wrapCmdDeleteProc() {
	clientData := (*int)(unsafe.Pointer(C.wrapClientData))
	callback.DeleteFunc(*clientData)
}
