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

	cefTypes "github.com/energye/cef/147/types"
)

// ICefBytesWriteHandler Parent: ICefWriteHandlerOwn
type ICefBytesWriteHandler interface {
	ICefWriteHandlerOwn
	Write(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt // function
	Seek(offset int64, whence int32) int32                                                  // function
	Tell() int64                                                                            // function
	Flush() int32                                                                           // function
	MayBlock() bool                                                                         // function
	GetData() uintptr                                                                       // function
	GetDataSize() int64                                                                     // function
	AsIntfWriteHandler() uintptr
}

type TCefBytesWriteHandler struct {
	TCefWriteHandlerOwn
}

func (m *TCefBytesWriteHandler) Write(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBytesWriteHandlerAPI().SysCallN(1, m.Instance(), uintptr(ptr), uintptr(size), uintptr(N))
	return cefTypes.NativeUInt(r)
}

func (m *TCefBytesWriteHandler) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBytesWriteHandlerAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(whence))
	return int32(r)
}

func (m *TCefBytesWriteHandler) Tell() (result int64) {
	if !m.IsValid() {
		return
	}
	cefBytesWriteHandlerAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefBytesWriteHandler) Flush() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBytesWriteHandlerAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCefBytesWriteHandler) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBytesWriteHandlerAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBytesWriteHandler) GetData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefBytesWriteHandlerAPI().SysCallN(6, m.Instance())
	return uintptr(r)
}

func (m *TCefBytesWriteHandler) GetDataSize() (result int64) {
	if !m.IsValid() {
		return
	}
	cefBytesWriteHandlerAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefBytesWriteHandler) AsIntfWriteHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewBytesWriteHandler class constructor
func NewBytesWriteHandler(grow cefTypes.NativeUInt) ICefBytesWriteHandler {
	var writeHandlerPtr uintptr // ICefWriteHandler
	r := cefBytesWriteHandlerAPI().SysCallN(0, uintptr(grow), uintptr(base.UnsafePointer(&writeHandlerPtr)))
	ret := AsCefBytesWriteHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, writeHandlerPtr)
	}
	return ret
}

var (
	cefBytesWriteHandlerOnce   base.Once
	cefBytesWriteHandlerImport *imports.Imports = nil
)

func cefBytesWriteHandlerAPI() *imports.Imports {
	cefBytesWriteHandlerOnce.Do(func() {
		cefBytesWriteHandlerImport = api.NewDefaultImports()
		cefBytesWriteHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBytesWriteHandler_Create", 0), // constructor NewBytesWriteHandler
			/* 1 */ imports.NewTable("TCefBytesWriteHandler_Write", 0), // function Write
			/* 2 */ imports.NewTable("TCefBytesWriteHandler_Seek", 0), // function Seek
			/* 3 */ imports.NewTable("TCefBytesWriteHandler_Tell", 0), // function Tell
			/* 4 */ imports.NewTable("TCefBytesWriteHandler_Flush", 0), // function Flush
			/* 5 */ imports.NewTable("TCefBytesWriteHandler_MayBlock", 0), // function MayBlock
			/* 6 */ imports.NewTable("TCefBytesWriteHandler_GetData", 0), // function GetData
			/* 7 */ imports.NewTable("TCefBytesWriteHandler_GetDataSize", 0), // function GetDataSize
		}
	})
	return cefBytesWriteHandlerImport
}
