//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

func (m *TCefV8Value) ExecuteFunction(obj ICefV8Value, arguments []ICefV8Value) ICefV8Value {
	var resultCefV8Value uintptr
	args := V8ValueArrayRef.New(len(arguments), 0)
	args.Set(arguments)
	predefV8ValueImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(obj), args.Instance(), uintptr(int32(args.Size())), uintptr(unsafePointer(&resultCefV8Value)))
	args.Free()
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) ExecuteFunctionWithContext(context ICefV8Context, obj ICefV8Value, arguments []ICefV8Value) ICefV8Value {
	var resultCefV8Value uintptr
	args := V8ValueArrayRef.New(len(arguments), 0)
	args.Set(arguments)
	predefV8ValueImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(context), GetObjectUintptr(obj), args.Instance(), uintptr(int32(args.Size())), uintptr(unsafePointer(&resultCefV8Value)))
	args.Free()
	return AsCefV8Value(resultCefV8Value)
}

var (
	predefV8ValueImport       *imports.Imports = nil
	predefV8ValueImportTables                  = []*imports.Table{
		imports.NewTable("CefV8Value_ExecuteFunction", 0),
		imports.NewTable("CefV8Value_ExecuteFunctionWithContext", 0),
	}
)

func predefV8ValueImportAPI() *imports.Imports {
	if predefV8ValueImport == nil {
		predefV8ValueImport = NewDefaultImports()
		predefV8ValueImport.SetImportTable(predefV8ValueImportTables)
		predefV8ValueImportTables = nil
	}
	return predefV8ValueImport
}
