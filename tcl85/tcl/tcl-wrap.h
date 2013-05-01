// Wrapper functions for the Tcl C library.

#include <tcl8.5/tcl.h>


// Go/C callback variables

extern ClientData clientData;
extern ClientData wrapClientData;
extern ClientData wrapClientDataDel;
extern Tcl_Interp *wrapInterp;
extern int wrapObjc;
extern struct Tcl_Obj * CONST * wrapObjv;
extern int wrapReturnVal;
extern Tcl_Command command;
extern Tcl_Interp *interp;
extern char *cmd_name;


// Go callback wrappers

extern void wrapObjCmdProc();
extern void wrapCmdDeleteProc();


// C callback wrappers

extern int callbackHandler(ClientData clientData,
	Tcl_Interp *interp, int objc, struct Tcl_Obj * CONST * objv);
extern void callbackDeleter(ClientData clientData);
extern void createObjCmd();


extern void decrRefCount(Tcl_Obj *ptr);

