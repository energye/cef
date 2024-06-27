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

// IDomVisitor Parent: ICefDomVisitor
//
//	Interface to implement for visiting the DOM. The functions of this interface
//	will be called on the render process main thread.
//	<a cref="uCEFTypes|TCefDomVisitor">Implements TCefDomVisitor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t)</a>
type IDomVisitor interface {
	ICefDomVisitor
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDomVisitor // procedure
	// SetOnDomVisitor
	//  Method executed for visiting the DOM. The document object passed to this
	//  function represents a snapshot of the DOM at the time this function is
	//  executed. DOM objects are only valid for the scope of this function. Do
	//  not keep references to or attempt to access any DOM objects outside the
	//  scope of this function.
	SetOnDomVisitor(fn TOnDomVisitor) // property event
}

// TDomVisitor Parent: TCefDomVisitor
//
//	Interface to implement for visiting the DOM. The functions of this interface
//	will be called on the render process main thread.
//	<a cref="uCEFTypes|TCefDomVisitor">Implements TCefDomVisitor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t)</a>
type TDomVisitor struct {
	TCefDomVisitor
	domVisitorPtr uintptr
}

func NewDomVisitor() IDomVisitor {
	r1 := domVisitorImportAPI().SysCallN(1)
	return AsDomVisitor(r1)
}

func (m *TDomVisitor) AsInterface() ICefDomVisitor {
	var resultCefDomVisitor uintptr
	domVisitorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDomVisitor)))
	return AsCefDomVisitor(resultCefDomVisitor)
}

func (m *TDomVisitor) SetOnDomVisitor(fn TOnDomVisitor) {
	if m.domVisitorPtr != 0 {
		RemoveEventElement(m.domVisitorPtr)
	}
	m.domVisitorPtr = MakeEventDataPtr(fn)
	domVisitorImportAPI().SysCallN(2, m.Instance(), m.domVisitorPtr)
}

var (
	domVisitorImport       *imports.Imports = nil
	domVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DomVisitor_AsInterface", 0),
		/*1*/ imports.NewTable("DomVisitor_Create", 0),
		/*2*/ imports.NewTable("DomVisitor_SetOnDomVisitor", 0),
	}
)

func domVisitorImportAPI() *imports.Imports {
	if domVisitorImport == nil {
		domVisitorImport = NewDefaultImports()
		domVisitorImport.SetImportTable(domVisitorImportTables)
		domVisitorImportTables = nil
	}
	return domVisitorImport
}
