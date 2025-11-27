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

// ICefBeforeDownloadCallback Parent: ICefBaseRefCountedRef
type ICefBeforeDownloadCallback interface {
	ICefBaseRefCountedRef
	// Cont
	//  Call to continue the download. Set |download_path| to the full file path
	//  for the download including the file name or leave blank to use the
	//  suggested name and the default temp directory. Set |show_dialog| to true
	//  (1) if you do wish to show the default "Save As" dialog.
	Cont(downloadPath string, showDialog bool) // procedure
}

// ICefBeforeDownloadCallbackRef Parent: ICefBeforeDownloadCallback
type ICefBeforeDownloadCallbackRef interface {
	ICefBeforeDownloadCallback
	AsIntfBeforeDownloadCallback() uintptr
}

type TCefBeforeDownloadCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefBeforeDownloadCallbackRef) Cont(downloadPath string, showDialog bool) {
	if !m.IsValid() {
		return
	}
	cefBeforeDownloadCallbackRefAPI().SysCallN(2, m.Instance(), api.PasStr(downloadPath), api.PasBool(showDialog))
}

func (m *TCefBeforeDownloadCallbackRef) AsIntfBeforeDownloadCallback() uintptr {
	return m.GetIntfPointer(0)
}

// BeforeDownloadCallbackRef  is static instance
var BeforeDownloadCallbackRef _BeforeDownloadCallbackRefClass

// _BeforeDownloadCallbackRefClass is class type defined by TCefBeforeDownloadCallbackRef
type _BeforeDownloadCallbackRefClass uintptr

func (_BeforeDownloadCallbackRefClass) UnWrap(data uintptr) (result ICefBeforeDownloadCallback) {
	var resultPtr uintptr
	cefBeforeDownloadCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBeforeDownloadCallbackRef(resultPtr)
	return
}

// NewBeforeDownloadCallbackRef class constructor
func NewBeforeDownloadCallbackRef(data uintptr) ICefBeforeDownloadCallbackRef {
	var beforeDownloadCallbackPtr uintptr // ICefBeforeDownloadCallback
	r := cefBeforeDownloadCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&beforeDownloadCallbackPtr)))
	ret := AsCefBeforeDownloadCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, beforeDownloadCallbackPtr)
	}
	return ret
}

var (
	cefBeforeDownloadCallbackRefOnce   base.Once
	cefBeforeDownloadCallbackRefImport *imports.Imports = nil
)

func cefBeforeDownloadCallbackRefAPI() *imports.Imports {
	cefBeforeDownloadCallbackRefOnce.Do(func() {
		cefBeforeDownloadCallbackRefImport = api.NewDefaultImports()
		cefBeforeDownloadCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBeforeDownloadCallbackRef_Create", 0), // constructor NewBeforeDownloadCallbackRef
			/* 1 */ imports.NewTable("TCefBeforeDownloadCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefBeforeDownloadCallbackRef_Cont", 0), // procedure Cont
		}
	})
	return cefBeforeDownloadCallbackRefImport
}
