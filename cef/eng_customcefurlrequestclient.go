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

// ICustomCefUrlrequestClient Parent: ICefUrlRequestClient
type ICustomCefUrlrequestClient interface {
	ICefUrlRequestClient
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefUrlRequestClient // procedure
}

// TCustomCefUrlrequestClient Parent: TCefUrlRequestClient
type TCustomCefUrlrequestClient struct {
	TCefUrlRequestClient
}

func NewCustomCefUrlrequestClient(events ICEFUrlRequestClientEvents) ICustomCefUrlrequestClient {
	r1 := customCefUrlrequestClientImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomCefUrlrequestClient(r1)
}

func (m *TCustomCefUrlrequestClient) AsInterface() ICefUrlRequestClient {
	var resultCefUrlRequestClient uintptr
	customCefUrlrequestClientImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefUrlRequestClient)))
	return AsCefUrlRequestClient(resultCefUrlRequestClient)
}

var (
	customCefUrlrequestClientImport       *imports.Imports = nil
	customCefUrlrequestClientImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCefUrlrequestClient_AsInterface", 0),
		/*1*/ imports.NewTable("CustomCefUrlrequestClient_Create", 0),
	}
)

func customCefUrlrequestClientImportAPI() *imports.Imports {
	if customCefUrlrequestClientImport == nil {
		customCefUrlrequestClientImport = NewDefaultImports()
		customCefUrlrequestClientImport.SetImportTable(customCefUrlrequestClientImportTables)
		customCefUrlrequestClientImportTables = nil
	}
	return customCefUrlrequestClientImport
}
