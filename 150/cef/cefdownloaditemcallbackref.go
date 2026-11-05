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
)

// ICefDownloadItemCallback Parent: ICefBaseRefCounted
type ICefDownloadItemCallback interface {
	ICefBaseRefCounted
	// Cancel
	//  Call to cancel the download.
	Cancel() // procedure
	// Pause
	//  Call to pause the download.
	Pause() // procedure
	// Resume
	//  Call to resume the download.
	Resume() // procedure
}

// ICefDownloadItemCallbackRef Parent: ICefDownloadItemCallback ICefBaseRefCountedRef
type ICefDownloadItemCallbackRef interface {
	ICefDownloadItemCallback
	ICefBaseRefCountedRef
	AsIntfDownloadItemCallback() uintptr
}

type TCefDownloadItemCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDownloadItemCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefDownloadItemCallbackRefAPI().SysCallN(2, m.Instance())
}

func (m *TCefDownloadItemCallbackRef) Pause() {
	if !m.IsValid() {
		return
	}
	cefDownloadItemCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefDownloadItemCallbackRef) Resume() {
	if !m.IsValid() {
		return
	}
	cefDownloadItemCallbackRefAPI().SysCallN(4, m.Instance())
}

func (m *TCefDownloadItemCallbackRef) AsIntfDownloadItemCallback() uintptr {
	return m.GetIntfPointer(0)
}

// DownloadItemCallbackRef  is static instance
var DownloadItemCallbackRef _DownloadItemCallbackRefClass

// _DownloadItemCallbackRefClass is class type defined by TCefDownloadItemCallbackRef
type _DownloadItemCallbackRefClass uintptr

func (_DownloadItemCallbackRefClass) UnWrap(data uintptr) (result ICefDownloadItemCallback) {
	var resultPtr uintptr
	cefDownloadItemCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDownloadItemCallbackRef(resultPtr)
	return
}

// NewDownloadItemCallbackRef class constructor
func NewDownloadItemCallbackRef(data uintptr) ICefDownloadItemCallbackRef {
	var downloadItemCallbackPtr uintptr // ICefDownloadItemCallback
	r := cefDownloadItemCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&downloadItemCallbackPtr)))
	ret := AsCefDownloadItemCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadItemCallbackPtr)
	}
	return ret
}

var (
	cefDownloadItemCallbackRefOnce   base.Once
	cefDownloadItemCallbackRefImport *imports.Imports = nil
)

func cefDownloadItemCallbackRefAPI() *imports.Imports {
	cefDownloadItemCallbackRefOnce.Do(func() {
		cefDownloadItemCallbackRefImport = api.NewDefaultImports()
		cefDownloadItemCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDownloadItemCallbackRef_Create", 0), // constructor NewDownloadItemCallbackRef
			/* 1 */ imports.NewTable("TCefDownloadItemCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefDownloadItemCallbackRef_Cancel", 0), // procedure Cancel
			/* 3 */ imports.NewTable("TCefDownloadItemCallbackRef_Pause", 0), // procedure Pause
			/* 4 */ imports.NewTable("TCefDownloadItemCallbackRef_Resume", 0), // procedure Resume
		}
	})
	return cefDownloadItemCallbackRefImport
}
