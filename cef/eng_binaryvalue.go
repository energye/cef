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

// ICefBinaryValue Parent: ICefBaseRefCounted
type ICefBinaryValue interface {
	ICefBaseRefCounted
	IsValid() bool                                                        // function
	IsOwned() bool                                                        // function
	IsSame(that ICefBinaryValue) bool                                     // function
	IsEqual(that ICefBinaryValue) bool                                    // function
	Copy() ICefBinaryValue                                                // function
	GetSize() NativeUInt                                                  // function
	GetData(buffer uintptr, bufferSize, dataOffset NativeUInt) NativeUInt // function
}

// TCefBinaryValue Parent: TCefBaseRefCounted
type TCefBinaryValue struct {
	TCefBaseRefCounted
}

// BinaryValueRef -> ICefBinaryValue
var BinaryValueRef binaryValue

// binaryValue TCefBinaryValue Ref
type binaryValue uintptr

func (m *binaryValue) UnWrap(data uintptr) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	binaryValueImportAPI().SysCallN(8, uintptr(data), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *binaryValue) New(data uintptr, dataSize NativeUInt) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	binaryValueImportAPI().SysCallN(7, uintptr(data), uintptr(dataSize), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefBinaryValue) IsValid() bool {
	r1 := binaryValueImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsOwned() bool {
	r1 := binaryValueImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsSame(that ICefBinaryValue) bool {
	r1 := binaryValueImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBinaryValue) IsEqual(that ICefBinaryValue) bool {
	r1 := binaryValueImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBinaryValue) Copy() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	binaryValueImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefBinaryValue) GetSize() NativeUInt {
	r1 := binaryValueImportAPI().SysCallN(2, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefBinaryValue) GetData(buffer uintptr, bufferSize, dataOffset NativeUInt) NativeUInt {
	r1 := binaryValueImportAPI().SysCallN(1, m.Instance(), uintptr(buffer), uintptr(bufferSize), uintptr(dataOffset))
	return NativeUInt(r1)
}

var (
	binaryValueImport       *imports.Imports = nil
	binaryValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBinaryValue_Copy", 0),
		/*1*/ imports.NewTable("CefBinaryValue_GetData", 0),
		/*2*/ imports.NewTable("CefBinaryValue_GetSize", 0),
		/*3*/ imports.NewTable("CefBinaryValue_IsEqual", 0),
		/*4*/ imports.NewTable("CefBinaryValue_IsOwned", 0),
		/*5*/ imports.NewTable("CefBinaryValue_IsSame", 0),
		/*6*/ imports.NewTable("CefBinaryValue_IsValid", 0),
		/*7*/ imports.NewTable("CefBinaryValue_New", 0),
		/*8*/ imports.NewTable("CefBinaryValue_UnWrap", 0),
	}
)

func binaryValueImportAPI() *imports.Imports {
	if binaryValueImport == nil {
		binaryValueImport = NewDefaultImports()
		binaryValueImport.SetImportTable(binaryValueImportTables)
		binaryValueImportTables = nil
	}
	return binaryValueImport
}
