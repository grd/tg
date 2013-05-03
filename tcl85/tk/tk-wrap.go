//
// Package tk is a (Tcl/Tk) Tk wrapper for Go
//
package tk

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

func Init(interp *tcl.Interp) error {
	ci := (*C.Tcl_Interp)(unsafe.Pointer(interp))
	i := C.Tk_Init(ci)
	if i != tcl.OK {
		return interp.Error()
	}
	return nil
}

// The event loop
func MainLoop() {
	C.Tk_MainLoop()
}
