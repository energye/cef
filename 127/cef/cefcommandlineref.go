//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

// ICefCommandLine Parent: ICefBaseRefCounted
type ICefCommandLine interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions
	//  if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may
	//  expose read-only objects.
	IsReadOnly() bool // function
	// Copy
	//  Returns a writable copy of this object.
	Copy() ICefCommandLine // function
	// GetCommandLineString
	//  Constructs and returns the represented command line string. Use this
	//  function cautiously because quoting behavior is unclear.
	GetCommandLineString() string // function
	// GetProgram
	//  Get the program part of the command line string (the first item).
	GetProgram() string // function
	// HasSwitches
	//  Returns true (1) if the command line has switches.
	HasSwitches() bool // function
	// HasSwitch
	//  Returns true (1) if the command line contains the given switch.
	HasSwitch(name string) bool // function
	// GetSwitchValue
	//  Returns the value associated with the given switch. If the switch has no
	//  value or isn't present this function returns the NULL string.
	GetSwitchValue(name string) string // function
	// GetSwitches
	//  Returns the map of switch names and values. If a switch has no value an
	//  NULL string is returned.
	GetSwitches(switches *lcl.IStrings) bool // function
	// HasArguments
	//  True if there are remaining command line arguments.
	HasArguments() bool // function
	// InitFromArgv
	//  Initialize the command line with the specified |argc| and |argv| values.
	//  The first argument must be the name of the program. This function is only
	//  supported on non-Windows platforms.
	InitFromArgv(argc int32, argv types.PPAnsiChar) // procedure
	// InitFromString
	//  Initialize the command line with the string returned by calling
	//  GetCommandLineW(). This function is only supported on Windows.
	InitFromString(commandLine string) // procedure
	// Reset
	//  Reset the command-line switches and arguments but leave the program
	//  component unchanged.
	Reset() // procedure
	// GetArgv
	//  Retrieve the original command line string as a vector of strings. The argv
	//  array: `{ program, [(--|-|/)switch[=value]]*, [--], [argument]* }`
	GetArgv(args *lcl.IStrings) // procedure
	// SetProgram
	//  Set the program part of the command line string (the first item).
	SetProgram(prog string) // procedure
	// AppendSwitch
	//  Add a switch to the end of the command line.
	AppendSwitch(name string) // procedure
	// AppendSwitchWithValue
	//  Add a switch with the specified value to the end of the command line. If
	//  the switch has no value pass an NULL value string.
	AppendSwitchWithValue(name string, value string) // procedure
	// GetArguments
	//  Get the remaining command line arguments.
	GetArguments(arguments *lcl.IStrings) // procedure
	// AppendArgument
	//  Add an argument to the end of the command line.
	AppendArgument(argument string) // procedure
	// PrependWrapper
	//  Insert a command before the current command. Common for debuggers, like
	//  "valgrind" or "gdb --args".
	PrependWrapper(wrapper string) // procedure
}

// ICefCommandLineRef Parent: ICefCommandLine ICefBaseRefCountedRef
type ICefCommandLineRef interface {
	ICefCommandLine
	ICefBaseRefCountedRef
	AsIntfCommandLine() uintptr
}

type TCefCommandLineRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefCommandLineRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefCommandLineRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefCommandLineRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) Copy() (result ICefCommandLine) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefCommandLineRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCommandLineRef(resultPtr)
	return
}

func (m *TCefCommandLineRef) GetCommandLineString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCommandLineRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCommandLineRef) GetProgram() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCommandLineRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCommandLineRef) HasSwitches() bool {
	if !m.IsValid() {
		return false
	}
	r := cefCommandLineRefAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) HasSwitch(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCommandLineRefAPI().SysCallN(7, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) GetSwitchValue(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCommandLineRefAPI().SysCallN(8, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCommandLineRef) GetSwitches(switches *lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	switchesPtr := base.GetObjectUintptr(*switches)
	r := cefCommandLineRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&switchesPtr)))
	*switches = lcl.AsStrings(switchesPtr)
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) HasArguments() bool {
	if !m.IsValid() {
		return false
	}
	r := cefCommandLineRefAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCommandLineRef) InitFromArgv(argc int32, argv types.PPAnsiChar) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(14, m.Instance(), uintptr(argc), uintptr(argv))
}

func (m *TCefCommandLineRef) InitFromString(commandLine string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(15, m.Instance(), api.PasStr(commandLine))
}

func (m *TCefCommandLineRef) Reset() {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(16, m.Instance())
}

func (m *TCefCommandLineRef) GetArgv(args *lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	argsPtr := base.GetObjectUintptr(*args)
	cefCommandLineRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&argsPtr)))
	*args = lcl.AsStrings(argsPtr)
}

func (m *TCefCommandLineRef) SetProgram(prog string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(18, m.Instance(), api.PasStr(prog))
}

func (m *TCefCommandLineRef) AppendSwitch(name string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(19, m.Instance(), api.PasStr(name))
}

func (m *TCefCommandLineRef) AppendSwitchWithValue(name string, value string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(20, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TCefCommandLineRef) GetArguments(arguments *lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	argumentsPtr := base.GetObjectUintptr(*arguments)
	cefCommandLineRefAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&argumentsPtr)))
	*arguments = lcl.AsStrings(argumentsPtr)
}

func (m *TCefCommandLineRef) AppendArgument(argument string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(22, m.Instance(), api.PasStr(argument))
}

func (m *TCefCommandLineRef) PrependWrapper(wrapper string) {
	if !m.IsValid() {
		return
	}
	cefCommandLineRefAPI().SysCallN(23, m.Instance(), api.PasStr(wrapper))
}

func (m *TCefCommandLineRef) AsIntfCommandLine() uintptr {
	return m.GetIntfPointer(0)
}

// CommandLineRef  is static instance
var CommandLineRef _CommandLineRefClass

// _CommandLineRefClass is class type defined by TCefCommandLineRef
type _CommandLineRefClass uintptr

// UnWrap
//
//	Returns a ICefCommandLine instance using a PCefCommandLine data pointer.
func (_CommandLineRefClass) UnWrap(data uintptr) (result ICefCommandLine) {
	var resultPtr uintptr
	cefCommandLineRefAPI().SysCallN(11, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCommandLineRef(resultPtr)
	return
}

// New
//
//	Create a new ICefCommandLine instance.
func (_CommandLineRefClass) New() (result ICefCommandLine) {
	var resultPtr uintptr
	cefCommandLineRefAPI().SysCallN(12, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCommandLineRef(resultPtr)
	return
}

// Global
//
//	Returns the singleton global ICefCommandLine object. The returned object
//	will be read-only.
func (_CommandLineRefClass) Global() (result ICefCommandLine) {
	var resultPtr uintptr
	cefCommandLineRefAPI().SysCallN(13, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCommandLineRef(resultPtr)
	return
}

// NewCommandLineRef class constructor
func NewCommandLineRef(data uintptr) ICefCommandLineRef {
	var commandLinePtr uintptr // ICefCommandLine
	r := cefCommandLineRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&commandLinePtr)))
	ret := AsCefCommandLineRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, commandLinePtr)
	}
	return ret
}

var (
	cefCommandLineRefOnce   base.Once
	cefCommandLineRefImport *imports.Imports = nil
)

func cefCommandLineRefAPI() *imports.Imports {
	cefCommandLineRefOnce.Do(func() {
		cefCommandLineRefImport = api.NewDefaultImports()
		cefCommandLineRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCommandLineRef_Create", 0), // constructor NewCommandLineRef
			/* 1 */ imports.NewTable("TCefCommandLineRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefCommandLineRef_IsReadOnly", 0), // function IsReadOnly
			/* 3 */ imports.NewTable("TCefCommandLineRef_Copy", 0), // function Copy
			/* 4 */ imports.NewTable("TCefCommandLineRef_GetCommandLineString", 0), // function GetCommandLineString
			/* 5 */ imports.NewTable("TCefCommandLineRef_GetProgram", 0), // function GetProgram
			/* 6 */ imports.NewTable("TCefCommandLineRef_HasSwitches", 0), // function HasSwitches
			/* 7 */ imports.NewTable("TCefCommandLineRef_HasSwitch", 0), // function HasSwitch
			/* 8 */ imports.NewTable("TCefCommandLineRef_GetSwitchValue", 0), // function GetSwitchValue
			/* 9 */ imports.NewTable("TCefCommandLineRef_GetSwitches", 0), // function GetSwitches
			/* 10 */ imports.NewTable("TCefCommandLineRef_HasArguments", 0), // function HasArguments
			/* 11 */ imports.NewTable("TCefCommandLineRef_UnWrap", 0), // static function UnWrap
			/* 12 */ imports.NewTable("TCefCommandLineRef_New", 0), // static function New
			/* 13 */ imports.NewTable("TCefCommandLineRef_Global", 0), // static function Global
			/* 14 */ imports.NewTable("TCefCommandLineRef_InitFromArgv", 0), // procedure InitFromArgv
			/* 15 */ imports.NewTable("TCefCommandLineRef_InitFromString", 0), // procedure InitFromString
			/* 16 */ imports.NewTable("TCefCommandLineRef_Reset", 0), // procedure Reset
			/* 17 */ imports.NewTable("TCefCommandLineRef_GetArgv", 0), // procedure GetArgv
			/* 18 */ imports.NewTable("TCefCommandLineRef_SetProgram", 0), // procedure SetProgram
			/* 19 */ imports.NewTable("TCefCommandLineRef_AppendSwitch", 0), // procedure AppendSwitch
			/* 20 */ imports.NewTable("TCefCommandLineRef_AppendSwitchWithValue", 0), // procedure AppendSwitchWithValue
			/* 21 */ imports.NewTable("TCefCommandLineRef_GetArguments", 0), // procedure GetArguments
			/* 22 */ imports.NewTable("TCefCommandLineRef_AppendArgument", 0), // procedure AppendArgument
			/* 23 */ imports.NewTable("TCefCommandLineRef_PrependWrapper", 0), // procedure PrependWrapper
		}
	})
	return cefCommandLineRefImport
}
