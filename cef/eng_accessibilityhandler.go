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

// IAccessibilityHandler Parent: ICefAccessibilityHandler
//
//	Implement this interface to receive accessibility notification when
//	accessibility events have been registered. The functions of this interface
//	will be called on the UI thread.
//	<a cref="uCEFTypes|TCefAccessibilityHandler">Implements TCefAccessibilityHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t)</a>
type IAccessibilityHandler interface {
	ICefAccessibilityHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefAccessibilityHandler // procedure
	// SetOnTreeChange
	//  Called after renderer process sends accessibility tree changes to the
	//  browser process.
	SetOnTreeChange(fn TOnAccessibility) // property event
	// SetOnLocationChange
	//  Called after renderer process sends accessibility location changes to the
	//  browser process.
	SetOnLocationChange(fn TOnAccessibility) // property event
}

// TAccessibilityHandler Parent: TCefAccessibilityHandler
//
//	Implement this interface to receive accessibility notification when
//	accessibility events have been registered. The functions of this interface
//	will be called on the UI thread.
//	<a cref="uCEFTypes|TCefAccessibilityHandler">Implements TCefAccessibilityHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t)</a>
type TAccessibilityHandler struct {
	TCefAccessibilityHandler
	treeChangePtr     uintptr
	locationChangePtr uintptr
}

func NewAccessibilityHandler() IAccessibilityHandler {
	r1 := accessibilityHandlerImportAPI().SysCallN(1)
	return AsAccessibilityHandler(r1)
}

func (m *TAccessibilityHandler) AsInterface() ICefAccessibilityHandler {
	var resultCefAccessibilityHandler uintptr
	accessibilityHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefAccessibilityHandler)))
	return AsCefAccessibilityHandler(resultCefAccessibilityHandler)
}

func (m *TAccessibilityHandler) SetOnTreeChange(fn TOnAccessibility) {
	if m.treeChangePtr != 0 {
		RemoveEventElement(m.treeChangePtr)
	}
	m.treeChangePtr = MakeEventDataPtr(fn)
	accessibilityHandlerImportAPI().SysCallN(3, m.Instance(), m.treeChangePtr)
}

func (m *TAccessibilityHandler) SetOnLocationChange(fn TOnAccessibility) {
	if m.locationChangePtr != 0 {
		RemoveEventElement(m.locationChangePtr)
	}
	m.locationChangePtr = MakeEventDataPtr(fn)
	accessibilityHandlerImportAPI().SysCallN(2, m.Instance(), m.locationChangePtr)
}

var (
	accessibilityHandlerImport       *imports.Imports = nil
	accessibilityHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("AccessibilityHandler_AsInterface", 0),
		/*1*/ imports.NewTable("AccessibilityHandler_Create", 0),
		/*2*/ imports.NewTable("AccessibilityHandler_SetOnLocationChange", 0),
		/*3*/ imports.NewTable("AccessibilityHandler_SetOnTreeChange", 0),
	}
)

func accessibilityHandlerImportAPI() *imports.Imports {
	if accessibilityHandlerImport == nil {
		accessibilityHandlerImport = NewDefaultImports()
		accessibilityHandlerImport.SetImportTable(accessibilityHandlerImportTables)
		accessibilityHandlerImportTables = nil
	}
	return accessibilityHandlerImport
}
