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

// IEngUrlrequestClient Parent: ICefUrlrequestClientOwn
type IEngUrlrequestClient interface {
	ICefUrlrequestClientOwn
	SetOnUrlrequestClientGetAuthCredentials(fn TOnUrlrequestClientGetAuthCredentialsEvent) // property event
	SetOnUrlrequestClientRequestComplete(fn TOnUrlrequestClientRequestCompleteEvent)       // property event
	SetOnUrlrequestClientUploadProgress(fn TOnUrlrequestClientUploadProgressEvent)         // property event
	SetOnUrlrequestClientDownloadProgress(fn TOnUrlrequestClientDownloadProgressEvent)     // property event
	SetOnUrlrequestClientDownloadData(fn TOnUrlrequestClientDownloadDataEvent)             // property event
	AsIntfUrlrequestClient() uintptr
}

type TEngUrlrequestClient struct {
	TCefUrlrequestClientOwn
}

func (m *TEngUrlrequestClient) SetOnUrlrequestClientGetAuthCredentials(fn TOnUrlrequestClientGetAuthCredentialsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUrlrequestClientGetAuthCredentialsEvent(fn)
	base.SetEvent(m, 1, engUrlrequestClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngUrlrequestClient) SetOnUrlrequestClientRequestComplete(fn TOnUrlrequestClientRequestCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUrlrequestClientRequestCompleteEvent(fn)
	base.SetEvent(m, 2, engUrlrequestClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngUrlrequestClient) SetOnUrlrequestClientUploadProgress(fn TOnUrlrequestClientUploadProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUrlrequestClientUploadProgressEvent(fn)
	base.SetEvent(m, 3, engUrlrequestClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngUrlrequestClient) SetOnUrlrequestClientDownloadProgress(fn TOnUrlrequestClientDownloadProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUrlrequestClientDownloadProgressEvent(fn)
	base.SetEvent(m, 4, engUrlrequestClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngUrlrequestClient) SetOnUrlrequestClientDownloadData(fn TOnUrlrequestClientDownloadDataEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUrlrequestClientDownloadDataEvent(fn)
	base.SetEvent(m, 5, engUrlrequestClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngUrlrequestClient) AsIntfUrlrequestClient() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngUrlrequestClient class constructor
func NewEngUrlrequestClient() IEngUrlrequestClient {
	var urlrequestClientPtr uintptr // ICefUrlrequestClient
	r := engUrlrequestClientAPI().SysCallN(0, uintptr(base.UnsafePointer(&urlrequestClientPtr)))
	ret := AsEngUrlrequestClient(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, urlrequestClientPtr)
	}
	return ret
}

var (
	engUrlrequestClientOnce   base.Once
	engUrlrequestClientImport *imports.Imports = nil
)

func engUrlrequestClientAPI() *imports.Imports {
	engUrlrequestClientOnce.Do(func() {
		engUrlrequestClientImport = api.NewDefaultImports()
		engUrlrequestClientImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngUrlrequestClient_Create", 0), // constructor NewEngUrlrequestClient
			/* 1 */ imports.NewTable("TEngUrlrequestClient_OnUrlrequestClientGetAuthCredentials", 0), // event OnUrlrequestClientGetAuthCredentials
			/* 2 */ imports.NewTable("TEngUrlrequestClient_OnUrlrequestClientRequestComplete", 0), // event OnUrlrequestClientRequestComplete
			/* 3 */ imports.NewTable("TEngUrlrequestClient_OnUrlrequestClientUploadProgress", 0), // event OnUrlrequestClientUploadProgress
			/* 4 */ imports.NewTable("TEngUrlrequestClient_OnUrlrequestClientDownloadProgress", 0), // event OnUrlrequestClientDownloadProgress
			/* 5 */ imports.NewTable("TEngUrlrequestClient_OnUrlrequestClientDownloadData", 0), // event OnUrlrequestClientDownloadData
		}
	})
	return engUrlrequestClientImport
}
