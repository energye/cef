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
}

// TCustomCefUrlrequestClient Parent: TCefUrlRequestClient
type TCustomCefUrlrequestClient struct {
	TCefUrlRequestClient
}

func NewCustomCefUrlrequestClient(events ICEFUrlRequestClientEvents) ICustomCefUrlrequestClient {
	r1 := customCefUrlrequestClientImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomCefUrlrequestClient(r1)
}

var (
	customCefUrlrequestClientImport       *imports.Imports = nil
	customCefUrlrequestClientImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCefUrlrequestClient_Create", 0),
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
