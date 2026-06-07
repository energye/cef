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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/109/types"
)

// ICEFUrlRequestClientEvents Parent: IObject
type ICEFUrlRequestClientEvents interface {
	IObject
}

// ICEFUrlRequestClientComponent Parent: ICEFUrlRequestClientEvents IComponent
type ICEFUrlRequestClientComponent interface {
	ICEFUrlRequestClientEvents
	IComponent
	AddURLRequest()                                   // procedure
	Client() IEngUrlrequestClient                     // property Client Getter
	ThreadID() cefTypes.TCefThreadId                  // property ThreadID Getter
	SetThreadID(value cefTypes.TCefThreadId)          // property ThreadID Setter
	SetOnRequestComplete(fn TOnRequestComplete)       // property event
	SetOnUploadProgress(fn TOnUploadProgress)         // property event
	SetOnDownloadProgress(fn TOnDownloadProgress)     // property event
	SetOnDownloadData(fn TOnDownloadData)             // property event
	SetOnGetAuthCredentials(fn TOnGetAuthCredentials) // property event
	SetOnCreateURLRequest(fn TNotifyEvent)            // property event
	AsIntfUrlRequestClientEvents() uintptr
}

type TCEFUrlRequestClientComponent struct {
	TComponent
}

func (m *TCEFUrlRequestClientComponent) AddURLRequest() {
	if !m.IsValid() {
		return
	}
	cEFUrlRequestClientComponentAPI().SysCallN(1, m.Instance())
}

func (m *TCEFUrlRequestClientComponent) Client() (result IEngUrlrequestClient) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFUrlRequestClientComponentAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngUrlrequestClient(resultPtr)
	return
}

func (m *TCEFUrlRequestClientComponent) ThreadID() cefTypes.TCefThreadId {
	if !m.IsValid() {
		return 0
	}
	r := cEFUrlRequestClientComponentAPI().SysCallN(3, 0, m.Instance())
	return cefTypes.TCefThreadId(r)
}

func (m *TCEFUrlRequestClientComponent) SetThreadID(value cefTypes.TCefThreadId) {
	if !m.IsValid() {
		return
	}
	cEFUrlRequestClientComponentAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCEFUrlRequestClientComponent) SetOnRequestComplete(fn TOnRequestComplete) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestComplete(fn)
	base.SetEvent(m, 4, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) SetOnUploadProgress(fn TOnUploadProgress) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUploadProgress(fn)
	base.SetEvent(m, 5, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) SetOnDownloadProgress(fn TOnDownloadProgress) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadProgress(fn)
	base.SetEvent(m, 6, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) SetOnDownloadData(fn TOnDownloadData) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadData(fn)
	base.SetEvent(m, 7, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) SetOnGetAuthCredentials(fn TOnGetAuthCredentials) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAuthCredentials(fn)
	base.SetEvent(m, 8, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) SetOnCreateURLRequest(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, cEFUrlRequestClientComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFUrlRequestClientComponent) AsIntfUrlRequestClientEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewUrlRequestClientComponent class constructor
func NewUrlRequestClientComponent(owner lcl.IComponent) ICEFUrlRequestClientComponent {
	var urlRequestClientEventsPtr uintptr // ICEFUrlRequestClientEvents
	r := cEFUrlRequestClientComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&urlRequestClientEventsPtr)))
	ret := AsCEFUrlRequestClientComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, urlRequestClientEventsPtr)
	}
	return ret
}

var (
	cEFUrlRequestClientComponentOnce   base.Once
	cEFUrlRequestClientComponentImport *imports.Imports = nil
)

func cEFUrlRequestClientComponentAPI() *imports.Imports {
	cEFUrlRequestClientComponentOnce.Do(func() {
		cEFUrlRequestClientComponentImport = api.NewDefaultImports()
		cEFUrlRequestClientComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFUrlRequestClientComponent_Create", 0), // constructor NewUrlRequestClientComponent
			/* 1 */ imports.NewTable("TCEFUrlRequestClientComponent_AddURLRequest", 0), // procedure AddURLRequest
			/* 2 */ imports.NewTable("TCEFUrlRequestClientComponent_Client", 0), // property Client
			/* 3 */ imports.NewTable("TCEFUrlRequestClientComponent_ThreadID", 0), // property ThreadID
			/* 4 */ imports.NewTable("TCEFUrlRequestClientComponent_OnRequestComplete", 0), // event OnRequestComplete
			/* 5 */ imports.NewTable("TCEFUrlRequestClientComponent_OnUploadProgress", 0), // event OnUploadProgress
			/* 6 */ imports.NewTable("TCEFUrlRequestClientComponent_OnDownloadProgress", 0), // event OnDownloadProgress
			/* 7 */ imports.NewTable("TCEFUrlRequestClientComponent_OnDownloadData", 0), // event OnDownloadData
			/* 8 */ imports.NewTable("TCEFUrlRequestClientComponent_OnGetAuthCredentials", 0), // event OnGetAuthCredentials
			/* 9 */ imports.NewTable("TCEFUrlRequestClientComponent_OnCreateURLRequest", 0), // event OnCreateURLRequest
		}
	})
	return cEFUrlRequestClientComponentImport
}
