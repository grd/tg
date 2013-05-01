// Wrapper functions for the Tcl C library.

#include <tcl8.5/tcl.h>
#include "tcl-wrap.h"
#include "_cgo_export.h"

//
// Go/C callback variables
//

ClientData clientData;
ClientData wrapClientData;
ClientData wrapClientDataDel;
Tcl_Interp *wrapInterp;
int wrapObjc;
struct Tcl_Obj * CONST * wrapObjv;
int wrapReturnVal;
Tcl_Command command;
Tcl_Interp *interp;
char *cmd_name;

//
// callback wrappers
//

int callbackHandler(ClientData clientData,
	Tcl_Interp *interp, int objc, struct Tcl_Obj * CONST * objv) {

	wrapClientData = clientData;
	wrapInterp = interp;
	wrapObjc = objc;
	wrapObjv = objv;
	wrapObjCmdProc();
	return wrapReturnVal;
};

void callbackDeleter(ClientData clientData) {
	wrapClientDataDel = clientData;
	wrapCmdDeleteProc();
};

void createObjCmd() {
     command = Tcl_CreateObjCommand(interp, cmd_name,
		callbackHandler, clientData, callbackDeleter);
};

void decrRefCount(Tcl_Obj *ptr) {
	Tcl_DecrRefCount(ptr);
};

