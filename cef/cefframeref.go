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
	"github.com/energye/lcl/lcl"
)

// ICefFrame Parent: ICefBaseRefCounted
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
	//  Returns the name for this frame. If the frame has an assigned name (for
	//  example, set via the iframe "name" attribute) then that value will be
	//  returned. Otherwise a unique name will be constructed based on the frame
	//  parent hierarchy. The main (top-level) frame will always have an NULL name
	//  value.
	GetName() string // function
	// GetIdentifier
	//  Returns the globally unique identifier for this frame or empty if the
	//  underlying frame does not yet exist.
	GetIdentifier() string // function
	// GetParent
	//  Returns the parent of this frame or NULL if this is the main (top-level)
	//  frame.
	GetParent() ICefFrame // function
	// GetUrl
	//  Returns the URL currently loaded in this frame.
	GetUrl() string // function
	// GetBrowser
	//  Returns the browser that this frame belongs to.
	GetBrowser() ICefBrowser // function
	// GetV8Context
	//  Get the V8 context associated with the frame. This function can only be
	//  called from the render process.
	GetV8Context() ICefv8Context // function
	// CreateUrlRequest
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently (see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequest(request ICefRequest, client IEngUrlrequestClient) ICefUrlRequest // function
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
	//  Save this frame's HTML source to a temporary file and open it in the
	//  default text viewing application. This function can only be called from
	//  the browser process.
	ViewSource() // procedure
	// GetSource
	//  Retrieve this frame's HTML source as a string sent to the specified
	//  visitor.
	GetSource(visitor IEngStringVisitor) // procedure
	// GetText
	//  Retrieve this frame's display text as a string sent to the specified
	//  visitor.
	GetText(visitor IEngStringVisitor) // procedure
	// LoadRequest
	//  Load the request represented by the |request| object.
	//
	//  WARNING: This function will fail with "bad IPC message" reason
	//  INVALID_INITIATOR_ORIGIN (213) unless you first navigate to the request
	//  origin using some other mechanism (LoadURL, link click, etc).
	LoadRequest(request ICefRequest) // procedure
	// LoadUrl
	//  Load the specified |url|.
	LoadUrl(url string) // procedure
	// ExecuteJavaScript
	//  Execute a string of JavaScript code in this frame. The |script_url|
	//  parameter is the URL where the script in question can be found, if any.
	//  The renderer may request this URL to show the developer the source of the
	//  error. The |start_line| parameter is the base line number to use for
	//  error reporting.
	ExecuteJavaScript(code string, scriptUrl string, startLine int32) // procedure
	// VisitDom
	//  Visit the DOM document. This function can only be called from the render
	//  process.
	VisitDom(visitor IEngDomVisitor) // procedure
	// SendProcessMessage
	//  Send a message to the specified |target_process|. Ownership of the message
	//  contents will be transferred and the |message| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases (for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
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
	sw := lcl.NewStringWrap()
	cefFrameRefAPI().SysCallN(4, m.Instance(), sw.Instance())
	result = sw.Value()
	sw.Free()
	return
}

func (m *TCefFrameRef) GetIdentifier() (result string) {
	if !m.IsValid() {
		return ""
	}
	sw := lcl.NewStringWrap()
	cefFrameRefAPI().SysCallN(5, m.Instance(), sw.Instance())
	result = sw.Value()
	sw.Free()
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
	sw := lcl.NewStringWrap()
	cefFrameRefAPI().SysCallN(7, m.Instance(), sw.Instance())
	result = sw.Value()
	sw.Free()
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
