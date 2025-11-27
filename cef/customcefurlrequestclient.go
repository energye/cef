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

// ICustomCefUrlrequestClient Parent: ICefUrlrequestClientOwn
type ICustomCefUrlrequestClient interface {
	ICefUrlrequestClientOwn
	AsIntfUrlrequestClient() uintptr
}

type TCustomCefUrlrequestClient struct {
	TCefUrlrequestClientOwn
}

func (m *TCustomCefUrlrequestClient) AsIntfUrlrequestClient() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCefUrlrequestClient class constructor
func NewCustomCefUrlrequestClient(events ICEFUrlRequestClientComponent) ICustomCefUrlrequestClient {
	var urlrequestClientPtr uintptr // ICefUrlrequestClient
	r := customCefUrlrequestClientAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&urlrequestClientPtr)))
	ret := AsCustomCefUrlrequestClient(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, urlrequestClientPtr)
	}
	return ret
}

var (
	customCefUrlrequestClientOnce   base.Once
	customCefUrlrequestClientImport *imports.Imports = nil
)

func customCefUrlrequestClientAPI() *imports.Imports {
	customCefUrlrequestClientOnce.Do(func() {
		customCefUrlrequestClientImport = api.NewDefaultImports()
		customCefUrlrequestClientImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCefUrlrequestClient_Create", 0), // constructor NewCustomCefUrlrequestClient
		}
	})
	return customCefUrlrequestClientImport
}
