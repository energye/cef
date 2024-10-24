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
//
//	Interface used to represent a frame in the browser window. When used in the browser process the functions of this interface may be called on any thread unless otherwise indicated in the comments. When used in the render process the functions of this interface may only be called on the main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_capi.h">CEF source file: /include/capi/cef_frame_capi.h (cef_frame_t))</a>
type ICefFrame interface {
	ICefBaseRefCounted
	// IsValid
	//  True if this object is currently attached to a valid frame.
	IsValid() bool // function
	// IsMain
	//  Returns true (1) if this is the main (top-level) frame.
	IsMain() bool // function
	// IsFocused
	//  Returns true (1) if this is the focused frame.
	IsFocused() bool // function
	// GetName
	//  Returns the name for this frame. If the frame has an assigned name (for example, set via the iframe "name" attribute) then that value will be returned. Otherwise a unique name will be constructed based on the frame parent hierarchy. The main (top-level) frame will always have an NULL name value.
	GetName() string // function
	// GetIdentifier
	//  Returns the globally unique identifier for this frame or < 0 if the underlying frame does not yet exist.
	GetIdentifier() (resultInt64 int64) // function
	// GetParent
	//  Returns the parent of this frame or NULL if this is the main (top-level) frame.
	GetParent() ICefFrame // function
	// GetUrl
	//  Returns the URL currently loaded in this frame.
	GetUrl() string // function
	// GetBrowser
	//  Returns the browser that this frame belongs to.
	GetBrowser() ICefBrowser // function
	// GetV8Context
	//  Get the V8 context associated with the frame. This function can only be called from the render process.
	GetV8Context() ICefV8Context // function
	// CreateUrlRequest
	//  Create a new URL request that will be treated as originating from this frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if you do not want the request to have this association, in which case it may be handled differently (see documentation on that function). A request created with this function may only originate from the browser process, and will behave as follows: - It may be intercepted by the client via CefResourceRequestHandler or CefSchemeHandlerFactory. - POST data may only contain a single element of type PDE_TYPE_FILE or PDE_TYPE_BYTES. The |request| object will be marked as read-only after calling this function.
	CreateUrlRequest(request ICefRequest, client ICefUrlRequestClient) ICefUrlRequest // function
	// Undo
	//  Execute undo in this frame.
	Undo() // procedure
	// Redo
	//  Execute redo in this frame.
	Redo() // procedure
	// Cut
	//  Execute cut in this frame.
	Cut() // procedure
	// Copy
	//  Execute copy in this frame.
	Copy() // procedure
	// Paste
	//  Execute paste in this frame.
	Paste() // procedure
	// Del
	//  Execute delete in this frame.
	Del() // procedure
	// SelectAll
	//  Execute select all in this frame.
	SelectAll() // procedure
	// ViewSource
	//  Save this frame's HTML source to a temporary file and open it in the default text viewing application. This function can only be called from the browser process.
	ViewSource() // procedure
	// GetSource
	//  Retrieve this frame's HTML source as a string sent to the specified visitor.
	GetSource(visitor ICefStringVisitor) // procedure
	// GetText
	//  Retrieve this frame's display text as a string sent to the specified visitor.
	GetText(visitor ICefStringVisitor) // procedure
	// LoadRequest
	//  Load the request represented by the |request| object. WARNING: This function will fail with "bad IPC message" reason INVALID_INITIATOR_ORIGIN (213) unless you first navigate to the request origin using some other mechanism (LoadURL, link click, etc).
	LoadRequest(request ICefRequest) // procedure
	// LoadUrl
	//  Load the specified |url|.
	LoadUrl(url string) // procedure
	// ExecuteJavaScript
	//  Execute a string of JavaScript code in this frame. The |script_url| parameter is the URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error. The |start_line| parameter is the base line number to use for error reporting.
	ExecuteJavaScript(code, scriptUrl string, startLine int32) // procedure
	// VisitDom
	//  Visit the DOM document. This function can only be called from the render process.
	VisitDom(visitor ICefDomVisitor) // procedure
	// SendProcessMessage
	//  Send a message to the specified |target_process|. Ownership of the message contents will be transferred and the |message| reference will be invalidated. Message delivery is not guaranteed in all cases (for example, if the browser is closing, navigating, or if the target process crashes). Send an ACK message back from the target process if confirmation is required.
	SendProcessMessage(targetProcess TCefProcessId, message ICefProcessMessage) // procedure
}

// TCefFrame Parent: TCefBaseRefCounted
//
//	Interface used to represent a frame in the browser window. When used in the browser process the functions of this interface may be called on any thread unless otherwise indicated in the comments. When used in the render process the functions of this interface may only be called on the main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_capi.h">CEF source file: /include/capi/cef_frame_capi.h (cef_frame_t))</a>
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
	value := NewTString()
	defer value.Free()
	frameImportAPI().SysCallN(7, m.Instance(), value.Instance())
	return value.Value()
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
	value := NewTString()
	defer value.Free()
	frameImportAPI().SysCallN(11, m.Instance(), value.Instance())
	return value.Value()
}

func (m *TCefFrame) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	frameImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefFrame) GetV8Context() ICefV8Context {
	var resultCefV8Context uintptr
	frameImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultCefV8Context)))
	return AsCefV8Context(resultCefV8Context)
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
