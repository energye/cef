//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package base

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
)

// SetCommandLine MacOS set command line args
func SetCommandLine(args lcl.IStrings) {
	cefApplicationDefAPI().SysCallN(0, args.Instance())
}

func AddCrDelegate() {
	cefApplicationDefAPI().SysCallN(1)
}

func SetGlobalCEFApplication(application uintptr) {
	cefApplicationDefAPI().SysCallN(2, application)
}

func DestroyGlobalCEFApplication() {
	cefApplicationDefAPI().SysCallN(3)
}

func SetGlobalCEFWorkSchedule(workSchedule uintptr) {
	cefApplicationDefAPI().SysCallN(4, workSchedule)
}

func DestroyGlobalCEFWorkSchedule() {
	cefApplicationDefAPI().SysCallN(5)
}

func CEFLibVersion() (major, minor, release, build int32) {
	/*
		procedure CEFLibVersion(out AMajor, AMinor, ARelease, ABuild: Integer); extdecl;
		begin
		  AMajor := Integer(CEF_SUPPORTED_VERSION_MAJOR);
		  AMinor := Integer(CEF_SUPPORTED_VERSION_MINOR);
		  ARelease := Integer(CEF_SUPPORTED_VERSION_RELEASE);
		  ABuild := Integer(CEF_SUPPORTED_VERSION_BUILD);
		end;
	*/
	cefApplicationDefAPI().SysCallN(6, uintptr(base.UnsafePointer(&major)), uintptr(base.UnsafePointer(&minor)),
		uintptr(base.UnsafePointer(&release)), uintptr(base.UnsafePointer(&build)))
	return
}

var (
	cefApplicationDefOnce   base.Once
	cefApplicationDefImport *imports.Imports = nil
)

func cefApplicationDefAPI() *imports.Imports {
	cefApplicationDefOnce.Do(func() {
		cefApplicationDefImport = api.NewDefaultImports()
		cefApplicationDefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("DarwinSetCommandLine", 0),
			/* 1 */ imports.NewTable("DarwinAddCrDelegate", 0),
			/* 2 */ imports.NewTable("SetGlobalCEFApplication", 0),
			/* 3 */ imports.NewTable("DestroyGlobalCEFApplication", 0),
			/* 4 */ imports.NewTable("SetGlobalCEFWorkSchedule", 0),
			/* 5 */ imports.NewTable("DestroyGlobalCEFWorkSchedule", 0),
			/* 6 */ imports.NewTable("CEFLibVersion", 0),
		}
	})
	return cefApplicationDefImport
}
