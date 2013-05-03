// Wrapper functions for the Tcl C library.

/* 
 * Copyright (C) 2013 G.vd.Schoot
 * 
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or (at
 * your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.
 */

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

