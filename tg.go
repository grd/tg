//
// Package tg is a simple Tcl/Tk GUI for Go, inspired by the
// Inferno (plan9) library
//
package tg

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

import (
	"github.com/grd/tg/tcl85/tcl"
	"github.com/grd/tg/tcl85/tk"
	"log"
)

// Initializing Tcl environment
func Init(arg string) (interp *tcl.Interp, err error) {
	tcl.FindExecutable(arg)
	return getInterp()
}

// Get the Tcl/Tk interpreter
func getInterp() (interp *tcl.Interp, err error) {
	interp = tcl.CreateInterp()

	err = tcl.Init(interp)
	if err != nil {
		return
	}

	err = tk.Init(interp)
	if err != nil {
		return
	}

	err = tcl.Eval(interp, "namespace eval gt {}")

	return
}

func DoEval(str string) error {
	log.Println(str)
	interp := GetInterp()
	err := tcl.Eval(interp, str)
	if err != nil {
		return err
	}
	return nil
}

// The event loop
func Run() {
	tk.MainLoop()
}

func Cmd(interp *tcl.Interp, script string) error {
	return tcl.Eval(interp, script)
}

// Widgets

func Button(options string) {
	// ".b" -text "Hello World" -command hello
}

func Pack(options string) {
	// ".b" -padx 25 -pady 10
}
