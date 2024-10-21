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

// ICefCommandLine Parent: ICefBaseRefCounted
//
//	Interface used to create and/or parse command line arguments. Arguments with "--", "-" and, on Windows, "/" prefixes are considered switches. Switches will always precede any arguments without switch prefixes. Switches can optionally have a value specified using the "=" delimiter (e.g. "-switch=value"). An argument of "--" will terminate switch parsing with all subsequent tokens, regardless of prefix, being interpreted as non-switch arguments. Switch names should be lowercase ASCII and will be converted to such if necessary. Switch values will retain the original case and UTF8 encoding. This interface can be used before cef_initialize() is called.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_line_capi.h">CEF source file: /include/capi/cef_command_line_capi.h (cef_command_line_t))</a>
type ICefCommandLine interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// Copy
	//  Returns a writable copy of this object.
	Copy() ICefCommandLine // function
	// GetCommandLineString
	//  Constructs and returns the represented command line string. Use this function cautiously because quoting behavior is unclear.
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
	//  Returns the value associated with the given switch. If the switch has no value or isn't present this function returns the NULL string.
	GetSwitchValue(name string) string // function
	// GetSwitches
	//  Returns the map of switch names and values. If a switch has no value an NULL string is returned.
	GetSwitches(switches *IStrings) bool // function
	// GetSwitches1
	//  Returns the map of switch names and values. If a switch has no value an NULL string is returned.
	GetSwitches1(switchKeys, switchValues *IStringList) bool // function
	// HasArguments
	//  True if there are remaining command line arguments.
	HasArguments() bool // function
	// InitFromArgv
	//  Initialize the command line with the specified |argc| and |argv| values. The first argument must be the name of the program. This function is only supported on non-Windows platforms.
	InitFromArgv(argc int32, argv PString) // procedure
	// InitFromString
	//  Initialize the command line with the string returned by calling GetCommandLineW(). This function is only supported on Windows.
	InitFromString(commandLine string) // procedure
	// Reset
	//  Reset the command-line switches and arguments but leave the program component unchanged.
	Reset() // procedure
	// GetArgv
	//  Retrieve the original command line string as a vector of strings. The argv array: `{ program, [(--|-|/)switch[=value]]*, [--], [argument]* }`
	GetArgv(args *IStrings) // procedure
	// SetProgram
	//  Set the program part of the command line string (the first item).
	SetProgram(prog string) // procedure
	// AppendSwitch
	//  Add a switch to the end of the command line.
	AppendSwitch(name string) // procedure
	// AppendSwitchWithValue
	//  Add a switch with the specified value to the end of the command line. If the switch has no value pass an NULL value string.
	AppendSwitchWithValue(name, value string) // procedure
	// GetArguments
	//  Get the remaining command line arguments.
	GetArguments(arguments *IStrings) // procedure
	// AppendArgument
	//  Add an argument to the end of the command line.
	AppendArgument(argument string) // procedure
	// PrependWrapper
	//  Insert a command before the current command. Common for debuggers, like "valgrind" or "gdb --args".
	PrependWrapper(wrapper string) // procedure
}

// TCefCommandLine Parent: TCefBaseRefCounted
//
//	Interface used to create and/or parse command line arguments. Arguments with "--", "-" and, on Windows, "/" prefixes are considered switches. Switches will always precede any arguments without switch prefixes. Switches can optionally have a value specified using the "=" delimiter (e.g. "-switch=value"). An argument of "--" will terminate switch parsing with all subsequent tokens, regardless of prefix, being interpreted as non-switch arguments. Switch names should be lowercase ASCII and will be converted to such if necessary. Switch values will retain the original case and UTF8 encoding. This interface can be used before cef_initialize() is called.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_line_capi.h">CEF source file: /include/capi/cef_command_line_capi.h (cef_command_line_t))</a>
type TCefCommandLine struct {
	TCefBaseRefCounted
}

// CommandLineRef -> ICefCommandLine
var CommandLineRef commandLine

// commandLine TCefCommandLine Ref
type commandLine uintptr

func (m *commandLine) UnWrap(data uintptr) ICefCommandLine {
	var resultCefCommandLine uintptr
	commandLineImportAPI().SysCallN(23, uintptr(data), uintptr(unsafePointer(&resultCefCommandLine)))
	return AsCefCommandLine(resultCefCommandLine)
}

func (m *commandLine) New() ICefCommandLine {
	var resultCefCommandLine uintptr
	commandLineImportAPI().SysCallN(19, uintptr(unsafePointer(&resultCefCommandLine)))
	return AsCefCommandLine(resultCefCommandLine)
}

func (m *commandLine) Global() ICefCommandLine {
	var resultCefCommandLine uintptr
	commandLineImportAPI().SysCallN(11, uintptr(unsafePointer(&resultCefCommandLine)))
	return AsCefCommandLine(resultCefCommandLine)
}

func (m *TCefCommandLine) IsValid() bool {
	r1 := commandLineImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefCommandLine) IsReadOnly() bool {
	r1 := commandLineImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefCommandLine) Copy() ICefCommandLine {
	var resultCefCommandLine uintptr
	commandLineImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefCommandLine)))
	return AsCefCommandLine(resultCefCommandLine)
}

func (m *TCefCommandLine) GetCommandLineString() string {
	r1 := commandLineImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCefCommandLine) GetProgram() string {
	r1 := commandLineImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefCommandLine) HasSwitches() bool {
	r1 := commandLineImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCefCommandLine) HasSwitch(name string) bool {
	r1 := commandLineImportAPI().SysCallN(13, m.Instance(), PascalStr(name))
	return GoBool(r1)
}

func (m *TCefCommandLine) GetSwitchValue(name string) string {
	r1 := commandLineImportAPI().SysCallN(8, m.Instance(), PascalStr(name))
	return GoStr(r1)
}

func (m *TCefCommandLine) GetSwitches(switches *IStrings) bool {
	var result0 uintptr
	r1 := commandLineImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&result0)))
	*switches = AsStrings(result0)
	return GoBool(r1)
}

func (m *TCefCommandLine) GetSwitches1(switchKeys, switchValues *IStringList) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := commandLineImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*switchKeys = AsStringList(result0)
	*switchValues = AsStringList(result1)
	return GoBool(r1)
}

func (m *TCefCommandLine) HasArguments() bool {
	r1 := commandLineImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TCefCommandLine) InitFromArgv(argc int32, argv PString) {
	commandLineImportAPI().SysCallN(15, m.Instance(), uintptr(argc), uintptr(argv))
}

func (m *TCefCommandLine) InitFromString(commandLine string) {
	commandLineImportAPI().SysCallN(16, m.Instance(), PascalStr(commandLine))
}

func (m *TCefCommandLine) Reset() {
	commandLineImportAPI().SysCallN(21, m.Instance())
}

func (m *TCefCommandLine) GetArgv(args *IStrings) {
	var result0 uintptr
	commandLineImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&result0)))
	*args = AsStrings(result0)
}

func (m *TCefCommandLine) SetProgram(prog string) {
	commandLineImportAPI().SysCallN(22, m.Instance(), PascalStr(prog))
}

func (m *TCefCommandLine) AppendSwitch(name string) {
	commandLineImportAPI().SysCallN(1, m.Instance(), PascalStr(name))
}

func (m *TCefCommandLine) AppendSwitchWithValue(name, value string) {
	commandLineImportAPI().SysCallN(2, m.Instance(), PascalStr(name), PascalStr(value))
}

func (m *TCefCommandLine) GetArguments(arguments *IStrings) {
	var result0 uintptr
	commandLineImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&result0)))
	*arguments = AsStrings(result0)
}

func (m *TCefCommandLine) AppendArgument(argument string) {
	commandLineImportAPI().SysCallN(0, m.Instance(), PascalStr(argument))
}

func (m *TCefCommandLine) PrependWrapper(wrapper string) {
	commandLineImportAPI().SysCallN(20, m.Instance(), PascalStr(wrapper))
}

var (
	commandLineImport       *imports.Imports = nil
	commandLineImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCommandLine_AppendArgument", 0),
		/*1*/ imports.NewTable("CefCommandLine_AppendSwitch", 0),
		/*2*/ imports.NewTable("CefCommandLine_AppendSwitchWithValue", 0),
		/*3*/ imports.NewTable("CefCommandLine_Copy", 0),
		/*4*/ imports.NewTable("CefCommandLine_GetArguments", 0),
		/*5*/ imports.NewTable("CefCommandLine_GetArgv", 0),
		/*6*/ imports.NewTable("CefCommandLine_GetCommandLineString", 0),
		/*7*/ imports.NewTable("CefCommandLine_GetProgram", 0),
		/*8*/ imports.NewTable("CefCommandLine_GetSwitchValue", 0),
		/*9*/ imports.NewTable("CefCommandLine_GetSwitches", 0),
		/*10*/ imports.NewTable("CefCommandLine_GetSwitches1", 0),
		/*11*/ imports.NewTable("CefCommandLine_Global", 0),
		/*12*/ imports.NewTable("CefCommandLine_HasArguments", 0),
		/*13*/ imports.NewTable("CefCommandLine_HasSwitch", 0),
		/*14*/ imports.NewTable("CefCommandLine_HasSwitches", 0),
		/*15*/ imports.NewTable("CefCommandLine_InitFromArgv", 0),
		/*16*/ imports.NewTable("CefCommandLine_InitFromString", 0),
		/*17*/ imports.NewTable("CefCommandLine_IsReadOnly", 0),
		/*18*/ imports.NewTable("CefCommandLine_IsValid", 0),
		/*19*/ imports.NewTable("CefCommandLine_New", 0),
		/*20*/ imports.NewTable("CefCommandLine_PrependWrapper", 0),
		/*21*/ imports.NewTable("CefCommandLine_Reset", 0),
		/*22*/ imports.NewTable("CefCommandLine_SetProgram", 0),
		/*23*/ imports.NewTable("CefCommandLine_UnWrap", 0),
	}
)

func commandLineImportAPI() *imports.Imports {
	if commandLineImport == nil {
		commandLineImport = NewDefaultImports()
		commandLineImport.SetImportTable(commandLineImportTables)
		commandLineImportTables = nil
	}
	return commandLineImport
}
