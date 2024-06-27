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

// ICefCustomCookieVisitor Parent: ICefCookieVisitor
type ICefCustomCookieVisitor interface {
	ICefCookieVisitor
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefCookieVisitor // procedure
}

// TCefCustomCookieVisitor Parent: TCefCookieVisitor
type TCefCustomCookieVisitor struct {
	TCefCookieVisitor
}

func NewCefCustomCookieVisitor(aEvents IChromiumEvents, aID int32) ICefCustomCookieVisitor {
	r1 := customCookieVisitorImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomCookieVisitor(r1)
}

func (m *TCefCustomCookieVisitor) AsInterface() ICefCookieVisitor {
	var resultCefCookieVisitor uintptr
	customCookieVisitorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefCookieVisitor)))
	return AsCefCookieVisitor(resultCefCookieVisitor)
}

var (
	customCookieVisitorImport       *imports.Imports = nil
	customCookieVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomCookieVisitor_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomCookieVisitor_Create", 0),
	}
)

func customCookieVisitorImportAPI() *imports.Imports {
	if customCookieVisitorImport == nil {
		customCookieVisitorImport = NewDefaultImports()
		customCookieVisitorImport.SetImportTable(customCookieVisitorImportTables)
		customCookieVisitorImportTables = nil
	}
	return customCookieVisitorImport
}
