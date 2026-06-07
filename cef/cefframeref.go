//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefFrame Parent: ICefBaseRefCounted
type ICefFrame interface {
	ICefBaseRefCounted
	IsValid() bool                                                                       // function
	IsMain() bool                                                                        // function
	IsFocused() bool                                                                     // function
	GetName() string                                                                     // function
	GetIdentifier() int64                                                                // function
	GetParent() ICefFrame                                                                // function
	GetUrl() string                                                                      // function
	GetBrowser() ICefBrowser                                                             // function
	GetV8Context() ICefv8Context                                                         // function
	CreateUrlRequest(request ICefRequest, client IEngUrlrequestClient) ICefUrlRequest    // function
	Undo()                                                                               // procedure
	Redo()                                                                               // procedure
	Cut()                                                                                // procedure
	Copy()                                                                               // procedure
	Paste()                                                                              // procedure
	Del()                                                                                // procedure
	SelectAll()                                                                          // procedure
	ViewSource()                                                                         // procedure
	GetSource(visitor IEngStringVisitor)                                                 // procedure
	GetText(visitor IEngStringVisitor)                                                   // procedure
	LoadRequest(request ICefRequest)                                                     // procedure
	LoadUrl(url string)                                                                  // procedure
	ExecuteJavaScript(code string, scriptUrl string, startLine int32)                    // procedure
	VisitDom(visitor IEngDomVisitor)                                                     // procedure
	SendProcessMessage(targetProcess cefTypes.TCefProcessId, message ICefProcessMessage) // procedure
}

// ICefFrameRef Parent: ICefFrame ICefBaseRefCountedRef
type ICefFrameRef interface {
	ICefFrame
	ICefBaseRefCountedRef
	AsIntfFrame() uintptr
}

type TCefFrameRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefFrameRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefFrameRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefFrameRef) IsMain() bool {
	if !m.IsValid() {
		return false
	}
	r := cefFrameRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefFrameRef) IsFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := cefFrameRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefFrameRef) GetName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefFrameRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefFrameRef) GetIdentifier() (result int64) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefFrameRef) GetParent() (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefFrameRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefFrameRef) GetUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefFrameRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefFrameRef) GetBrowser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefFrameRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCefFrameRef) GetV8Context() (result ICefv8Context) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefFrameRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ContextRef(resultPtr)
	return
}

func (m *TCefFrameRef) CreateUrlRequest(request ICefRequest, client IEngUrlrequestClient) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefFrameRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (m *TCefFrameRef) Undo() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(12, m.Instance())
}

func (m *TCefFrameRef) Redo() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(13, m.Instance())
}

func (m *TCefFrameRef) Cut() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(14, m.Instance())
}

func (m *TCefFrameRef) Copy() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(15, m.Instance())
}

func (m *TCefFrameRef) Paste() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(16, m.Instance())
}

func (m *TCefFrameRef) Del() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(17, m.Instance())
}

func (m *TCefFrameRef) SelectAll() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(18, m.Instance())
}

func (m *TCefFrameRef) ViewSource() {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(19, m.Instance())
}

func (m *TCefFrameRef) GetSource(visitor IEngStringVisitor) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(visitor))
}

func (m *TCefFrameRef) GetText(visitor IEngStringVisitor) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(visitor))
}

func (m *TCefFrameRef) LoadRequest(request ICefRequest) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(request))
}

func (m *TCefFrameRef) LoadUrl(url string) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(23, m.Instance(), api.PasStr(url))
}

func (m *TCefFrameRef) ExecuteJavaScript(code string, scriptUrl string, startLine int32) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(24, m.Instance(), api.PasStr(code), api.PasStr(scriptUrl), uintptr(startLine))
}

func (m *TCefFrameRef) VisitDom(visitor IEngDomVisitor) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(visitor))
}

func (m *TCefFrameRef) SendProcessMessage(targetProcess cefTypes.TCefProcessId, message ICefProcessMessage) {
	if !m.IsValid() {
		return
	}
	cefFrameRefAPI().SysCallN(26, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(message))
}

func (m *TCefFrameRef) AsIntfFrame() uintptr {
	return m.GetIntfPointer(0)
}

// FrameRef  is static instance
var FrameRef _FrameRefClass

// _FrameRefClass is class type defined by TCefFrameRef
type _FrameRefClass uintptr

func (_FrameRefClass) UnWrap(data uintptr) (result ICefFrame) {
	var resultPtr uintptr
	cefFrameRefAPI().SysCallN(11, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

// NewFrameRef class constructor
func NewFrameRef(data uintptr) ICefFrameRef {
	var framePtr uintptr // ICefFrame
	r := cefFrameRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&framePtr)))
	ret := AsCefFrameRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, framePtr)
	}
	return ret
}

var (
	cefFrameRefOnce   base.Once
	cefFrameRefImport *imports.Imports = nil
)

func cefFrameRefAPI() *imports.Imports {
	cefFrameRefOnce.Do(func() {
		cefFrameRefImport = api.NewDefaultImports()
		cefFrameRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefFrameRef_Create", 0), // constructor NewFrameRef
			/* 1 */ imports.NewTable("TCefFrameRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefFrameRef_IsMain", 0), // function IsMain
			/* 3 */ imports.NewTable("TCefFrameRef_IsFocused", 0), // function IsFocused
			/* 4 */ imports.NewTable("TCefFrameRef_GetName", 0), // function GetName
			/* 5 */ imports.NewTable("TCefFrameRef_GetIdentifier", 0), // function GetIdentifier
			/* 6 */ imports.NewTable("TCefFrameRef_GetParent", 0), // function GetParent
			/* 7 */ imports.NewTable("TCefFrameRef_GetUrl", 0), // function GetUrl
			/* 8 */ imports.NewTable("TCefFrameRef_GetBrowser", 0), // function GetBrowser
			/* 9 */ imports.NewTable("TCefFrameRef_GetV8Context", 0), // function GetV8Context
			/* 10 */ imports.NewTable("TCefFrameRef_CreateUrlRequest", 0), // function CreateUrlRequest
			/* 11 */ imports.NewTable("TCefFrameRef_UnWrap", 0), // static function UnWrap
			/* 12 */ imports.NewTable("TCefFrameRef_Undo", 0), // procedure Undo
			/* 13 */ imports.NewTable("TCefFrameRef_Redo", 0), // procedure Redo
			/* 14 */ imports.NewTable("TCefFrameRef_Cut", 0), // procedure Cut
			/* 15 */ imports.NewTable("TCefFrameRef_Copy", 0), // procedure Copy
			/* 16 */ imports.NewTable("TCefFrameRef_Paste", 0), // procedure Paste
			/* 17 */ imports.NewTable("TCefFrameRef_Del", 0), // procedure Del
			/* 18 */ imports.NewTable("TCefFrameRef_SelectAll", 0), // procedure SelectAll
			/* 19 */ imports.NewTable("TCefFrameRef_ViewSource", 0), // procedure ViewSource
			/* 20 */ imports.NewTable("TCefFrameRef_GetSource", 0), // procedure GetSource
			/* 21 */ imports.NewTable("TCefFrameRef_GetText", 0), // procedure GetText
			/* 22 */ imports.NewTable("TCefFrameRef_LoadRequest", 0), // procedure LoadRequest
			/* 23 */ imports.NewTable("TCefFrameRef_LoadUrl", 0), // procedure LoadUrl
			/* 24 */ imports.NewTable("TCefFrameRef_ExecuteJavaScript", 0), // procedure ExecuteJavaScript
			/* 25 */ imports.NewTable("TCefFrameRef_VisitDom", 0), // procedure VisitDom
			/* 26 */ imports.NewTable("TCefFrameRef_SendProcessMessage", 0), // procedure SendProcessMessage
		}
	})
	return cefFrameRefImport
}
