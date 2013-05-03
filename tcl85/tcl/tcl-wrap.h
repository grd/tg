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

