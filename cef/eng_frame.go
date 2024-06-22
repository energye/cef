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

// ICefFrame Parent: ICefBaseRefCounted
type ICefFrame interface {
	ICefBaseRefCounted
	IsValid() bool                                                                    // function
	IsMain() bool                                                                     // function
	IsFocused() bool                                                                  // function
	GetName() string                                                                  // function
	GetIdentifier() (resultInt64 int64)                                               // function
	GetParent() ICefFrame                                                             // function
	GetUrl() string                                                                   // function
	GetBrowser() ICefBrowser                                                          // function
	GetV8Context() ICefv8Context                                                      // function
	CreateUrlRequest(request ICefRequest, client ICefUrlRequestClient) ICefUrlRequest // function
	Undo()                                                                            // procedure
	Redo()                                                                            // procedure
	Cut()                                                                             // procedure
	Copy()                                                                            // procedure
	Paste()                                                                           // procedure
	Del()                                                                             // procedure
	SelectAll()                                                                       // procedure
	ViewSource()                                                                      // procedure
	GetSource(visitor ICefStringVisitor)                                              // procedure
	GetText(visitor ICefStringVisitor)                                                // procedure
	LoadRequest(request ICefRequest)                                                  // procedure
	LoadUrl(url string)                                                               // procedure
	ExecuteJavaScript(code, scriptUrl string, startLine int32)                        // procedure
	VisitDom(visitor ICefDomVisitor)                                                  // procedure
	SendProcessMessage(targetProcess TCefProcessId, message ICefProcessMessage)       // procedure
}

// TCefFrame Parent: TCefBaseRefCounted
type TCefFrame struct {
	TCefBaseRefCounted
}

// FrameRef -> ICefFrame
var FrameRef frame

// frame TCefFrame Ref
type frame uintptr

func (m *frame) UnWrap(data uintptr) ICefFrame {
	var resultCefFrame uintptr
	frameImportAPI().SysCallN(22, uintptr(data), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefFrame) IsValid() bool {
	r1 := frameImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefFrame) IsMain() bool {
	r1 := frameImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCefFrame) IsFocused() bool {
	r1 := frameImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefFrame) GetName() string {
	r1 := frameImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefFrame) GetIdentifier() (resultInt64 int64) {
	frameImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefFrame) GetParent() ICefFrame {
	var resultCefFrame uintptr
	frameImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefFrame) GetUrl() string {
	r1 := frameImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefFrame) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	frameImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefFrame) GetV8Context() ICefv8Context {
	var resultCefv8Context uintptr
	frameImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultCefv8Context)))
	return AsCefv8Context(resultCefv8Context)
}

func (m *TCefFrame) CreateUrlRequest(request ICefRequest, client ICefUrlRequestClient) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	frameImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(request), GetObjectUintptr(client), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *TCefFrame) Undo() {
	frameImportAPI().SysCallN(23, m.Instance())
}

func (m *TCefFrame) Redo() {
	frameImportAPI().SysCallN(19, m.Instance())
}

func (m *TCefFrame) Cut() {
	frameImportAPI().SysCallN(2, m.Instance())
}

func (m *TCefFrame) Copy() {
	frameImportAPI().SysCallN(0, m.Instance())
}

func (m *TCefFrame) Paste() {
	frameImportAPI().SysCallN(18, m.Instance())
}

func (m *TCefFrame) Del() {
	frameImportAPI().SysCallN(3, m.Instance())
}

func (m *TCefFrame) SelectAll() {
	frameImportAPI().SysCallN(20, m.Instance())
}

func (m *TCefFrame) ViewSource() {
	frameImportAPI().SysCallN(24, m.Instance())
}

func (m *TCefFrame) GetSource(visitor ICefStringVisitor) {
	frameImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(visitor))
}

func (m *TCefFrame) GetText(visitor ICefStringVisitor) {
	frameImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(visitor))
}

func (m *TCefFrame) LoadRequest(request ICefRequest) {
	frameImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(request))
}

func (m *TCefFrame) LoadUrl(url string) {
	frameImportAPI().SysCallN(17, m.Instance(), PascalStr(url))
}

func (m *TCefFrame) ExecuteJavaScript(code, scriptUrl string, startLine int32) {
	frameImportAPI().SysCallN(4, m.Instance(), PascalStr(code), PascalStr(scriptUrl), uintptr(startLine))
}

func (m *TCefFrame) VisitDom(visitor ICefDomVisitor) {
	frameImportAPI().SysCallN(25, m.Instance(), GetObjectUintptr(visitor))
}

func (m *TCefFrame) SendProcessMessage(targetProcess TCefProcessId, message ICefProcessMessage) {
	frameImportAPI().SysCallN(21, m.Instance(), uintptr(targetProcess), GetObjectUintptr(message))
}

var (
	frameImport       *imports.Imports = nil
	frameImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefFrame_Copy", 0),
		/*1*/ imports.NewTable("CefFrame_CreateUrlRequest", 0),
		/*2*/ imports.NewTable("CefFrame_Cut", 0),
		/*3*/ imports.NewTable("CefFrame_Del", 0),
		/*4*/ imports.NewTable("CefFrame_ExecuteJavaScript", 0),
		/*5*/ imports.NewTable("CefFrame_GetBrowser", 0),
		/*6*/ imports.NewTable("CefFrame_GetIdentifier", 0),
		/*7*/ imports.NewTable("CefFrame_GetName", 0),
		/*8*/ imports.NewTable("CefFrame_GetParent", 0),
		/*9*/ imports.NewTable("CefFrame_GetSource", 0),
		/*10*/ imports.NewTable("CefFrame_GetText", 0),
		/*11*/ imports.NewTable("CefFrame_GetUrl", 0),
		/*12*/ imports.NewTable("CefFrame_GetV8Context", 0),
		/*13*/ imports.NewTable("CefFrame_IsFocused", 0),
		/*14*/ imports.NewTable("CefFrame_IsMain", 0),
		/*15*/ imports.NewTable("CefFrame_IsValid", 0),
		/*16*/ imports.NewTable("CefFrame_LoadRequest", 0),
		/*17*/ imports.NewTable("CefFrame_LoadUrl", 0),
		/*18*/ imports.NewTable("CefFrame_Paste", 0),
		/*19*/ imports.NewTable("CefFrame_Redo", 0),
		/*20*/ imports.NewTable("CefFrame_SelectAll", 0),
		/*21*/ imports.NewTable("CefFrame_SendProcessMessage", 0),
		/*22*/ imports.NewTable("CefFrame_UnWrap", 0),
		/*23*/ imports.NewTable("CefFrame_Undo", 0),
		/*24*/ imports.NewTable("CefFrame_ViewSource", 0),
		/*25*/ imports.NewTable("CefFrame_VisitDom", 0),
	}
)

func frameImportAPI() *imports.Imports {
	if frameImport == nil {
		frameImport = NewDefaultImports()
		frameImport.SetImportTable(frameImportTables)
		frameImportTables = nil
	}
	return frameImport
}
