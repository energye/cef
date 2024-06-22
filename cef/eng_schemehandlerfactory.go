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

// ISchemeHandlerFactory Parent: ICefSchemeHandlerFactory
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type ISchemeHandlerFactory interface {
	ICefSchemeHandlerFactory
	// SetOnNew
	//  Return a new resource handler instance to handle the request or an NULL
	//  reference to allow default handling of the request. |browser| and |frame|
	//  will be the browser window and frame respectively that originated the
	//  request or NULL if the request did not originate from a browser window
	//  (for example, if the request came from ICefUrlRequest). The |request|
	//  object passed to this function cannot be modified.
	SetOnNew(fn TOnSchemeHandlerFactoryNew) // property event
}

// TSchemeHandlerFactory Parent: TCefSchemeHandlerFactory
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type TSchemeHandlerFactory struct {
	TCefSchemeHandlerFactory
	newPtr uintptr
}

func NewSchemeHandlerFactory(aClass TCefResourceHandlerClass) ISchemeHandlerFactory {
	r1 := schemeHandlerFactoryImportAPI().SysCallN(0, uintptr(aClass))
	return AsSchemeHandlerFactory(r1)
}

func (m *TSchemeHandlerFactory) SetOnNew(fn TOnSchemeHandlerFactoryNew) {
	if m.newPtr != 0 {
		RemoveEventElement(m.newPtr)
	}
	m.newPtr = MakeEventDataPtr(fn)
	schemeHandlerFactoryImportAPI().SysCallN(1, m.Instance(), m.newPtr)
}

var (
	schemeHandlerFactoryImport       *imports.Imports = nil
	schemeHandlerFactoryImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SchemeHandlerFactory_Create", 0),
		/*1*/ imports.NewTable("SchemeHandlerFactory_SetOnNew", 0),
	}
)

func schemeHandlerFactoryImportAPI() *imports.Imports {
	if schemeHandlerFactoryImport == nil {
		schemeHandlerFactoryImport = NewDefaultImports()
		schemeHandlerFactoryImport.SetImportTable(schemeHandlerFactoryImportTables)
		schemeHandlerFactoryImportTables = nil
	}
	return schemeHandlerFactoryImport
}
