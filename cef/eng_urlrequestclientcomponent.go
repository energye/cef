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

// ICEFUrlRequestClientComponent Parent: IComponent
//
//	The TCEFUrlRequestClientComponent class puts together all CEF URL request procedures, functions, properties and events in one place.
type ICEFUrlRequestClientComponent interface {
	IComponent
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICEFUrlRequestClientEvents // procedure
	// Client
	//  Returns the client.
	Client() ICefUrlRequestClient // property
	// ThreadID
	//  CEF thread used to create the URLRequest. Most of the client events will be executed on the same thread.
	ThreadID() TCefThreadId // property
	// SetThreadID Set ThreadID
	SetThreadID(AValue TCefThreadId) // property
	// AddURLRequest
	//  Create the URLRequest in the context of TCEFUrlRequestClientComponent.ThreadId, which is the CEF UI thread by default.
	AddURLRequest() // procedure
	// SetOnRequestComplete
	//  Notifies the client that the request has completed. Use the
	//  ICefUrlRequest.GetRequestStatus function to determine if the request
	//  was successful or not.
	//  This event will be called on the TCEFUrlRequestClientComponent.ThreadId thread, which is the CEF UI thread by default.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h(cef_urlrequest_client_t)</a>
	SetOnRequestComplete(fn TOnRequestCompleteRcc) // property event
	// SetOnUploadProgress
	//  Notifies the client of upload progress. |current| denotes the number of
	//  bytes sent so far and |total| is the total size of uploading data(or -1
	//  if chunked upload is enabled). This function will only be called if the
	//  UR_FLAG_REPORT_UPLOAD_PROGRESS flag is set on the request.
	//  This event will be called on the TCEFUrlRequestClientComponent.ThreadId thread, which is the CEF UI thread by default.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h(cef_urlrequest_client_t)</a>
	SetOnUploadProgress(fn TOnUploadProgressRcc) // property event
	// SetOnDownloadProgress
	//  Notifies the client of download progress. |current| denotes the number of
	//  bytes received up to the call and |total| is the expected total size of
	//  the response(or -1 if not determined).
	//  This event will be called on the TCEFUrlRequestClientComponent.ThreadId thread, which is the CEF UI thread by default.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h(cef_urlrequest_client_t)</a>
	SetOnDownloadProgress(fn TOnDownloadProgressRcc) // property event
	// SetOnDownloadData
	//  Called when some part of the response is read. |data| contains the current
	//  bytes received since the last call. This function will not be called if
	//  the UR_FLAG_NO_DOWNLOAD_DATA flag is set on the request.
	//  This event will be called on the TCEFUrlRequestClientComponent.ThreadId thread, which is the CEF UI thread by default.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h(cef_urlrequest_client_t)</a>
	SetOnDownloadData(fn TOnDownloadDataRcc) // property event
	// SetOnGetAuthCredentials
	//  Called on the IO thread when the browser needs credentials from the user.
	//  |isProxy| indicates whether the host is a proxy server. |host| contains
	//  the hostname and |port| contains the port number. Return true(1) to
	//  continue the request and call ICefAuthCallback.cont() when the
	//  authentication information is available. If the request has an associated
	//  browser/frame then returning false(0) will result in a call to
	//  GetAuthCredentials on the ICefRequestHandler associated with that
	//  browser, if any. Otherwise, returning false(0) will cancel the request
	//  immediately. This function will only be called for requests initiated from
	//  the browser process.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h(cef_urlrequest_client_t)</a>
	SetOnGetAuthCredentials(fn TOnGetAuthCredentialsRcc) // property event
	// SetOnCreateURLRequest
	//  Event triggered when the URLRequest has been created.
	SetOnCreateURLRequest(fn TNotifyRcc) // property event
}

// TCEFUrlRequestClientComponent Parent: TComponent
//
//	The TCEFUrlRequestClientComponent class puts together all CEF URL request procedures, functions, properties and events in one place.
type TCEFUrlRequestClientComponent struct {
	TComponent
	requestCompletePtr    uintptr
	uploadProgressPtr     uintptr
	downloadProgressPtr   uintptr
	downloadDataPtr       uintptr
	getAuthCredentialsPtr uintptr
	createURLRequestPtr   uintptr
}

func NewCEFUrlRequestClientComponent(aOwner IComponent) ICEFUrlRequestClientComponent {
	r1 := urlRequestClientComponentImportAPI().SysCallN(3, GetObjectUintptr(aOwner))
	return AsCEFUrlRequestClientComponent(r1)
}

func (m *TCEFUrlRequestClientComponent) AsInterface() ICEFUrlRequestClientEvents {
	var resultCEFUrlRequestClientEvents uintptr
	urlRequestClientComponentImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCEFUrlRequestClientEvents)))
	return AsCEFUrlRequestClientEvents(resultCEFUrlRequestClientEvents)
}

func (m *TCEFUrlRequestClientComponent) Client() ICefUrlRequestClient {
	var resultCefUrlRequestClient uintptr
	urlRequestClientComponentImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefUrlRequestClient)))
	return AsCefUrlRequestClient(resultCefUrlRequestClient)
}

func (m *TCEFUrlRequestClientComponent) ThreadID() TCefThreadId {
	r1 := urlRequestClientComponentImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TCefThreadId(r1)
}

func (m *TCEFUrlRequestClientComponent) SetThreadID(AValue TCefThreadId) {
	urlRequestClientComponentImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFUrlRequestClientComponent) AddURLRequest() {
	urlRequestClientComponentImportAPI().SysCallN(0, m.Instance())
}

func (m *TCEFUrlRequestClientComponent) SetOnRequestComplete(fn TOnRequestCompleteRcc) {
	if m.requestCompletePtr != 0 {
		RemoveEventElement(m.requestCompletePtr)
	}
	m.requestCompletePtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(8, m.Instance(), m.requestCompletePtr)
}

func (m *TCEFUrlRequestClientComponent) SetOnUploadProgress(fn TOnUploadProgressRcc) {
	if m.uploadProgressPtr != 0 {
		RemoveEventElement(m.uploadProgressPtr)
	}
	m.uploadProgressPtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(9, m.Instance(), m.uploadProgressPtr)
}

func (m *TCEFUrlRequestClientComponent) SetOnDownloadProgress(fn TOnDownloadProgressRcc) {
	if m.downloadProgressPtr != 0 {
		RemoveEventElement(m.downloadProgressPtr)
	}
	m.downloadProgressPtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(6, m.Instance(), m.downloadProgressPtr)
}

func (m *TCEFUrlRequestClientComponent) SetOnDownloadData(fn TOnDownloadDataRcc) {
	if m.downloadDataPtr != 0 {
		RemoveEventElement(m.downloadDataPtr)
	}
	m.downloadDataPtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(5, m.Instance(), m.downloadDataPtr)
}

func (m *TCEFUrlRequestClientComponent) SetOnGetAuthCredentials(fn TOnGetAuthCredentialsRcc) {
	if m.getAuthCredentialsPtr != 0 {
		RemoveEventElement(m.getAuthCredentialsPtr)
	}
	m.getAuthCredentialsPtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(7, m.Instance(), m.getAuthCredentialsPtr)
}

func (m *TCEFUrlRequestClientComponent) SetOnCreateURLRequest(fn TNotifyRcc) {
	if m.createURLRequestPtr != 0 {
		RemoveEventElement(m.createURLRequestPtr)
	}
	m.createURLRequestPtr = MakeEventDataPtr(fn)
	urlRequestClientComponentImportAPI().SysCallN(4, m.Instance(), m.createURLRequestPtr)
}

var (
	urlRequestClientComponentImport       *imports.Imports = nil
	urlRequestClientComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFUrlRequestClientComponent_AddURLRequest", 0),
		/*1*/ imports.NewTable("CEFUrlRequestClientComponent_AsInterface", 0),
		/*2*/ imports.NewTable("CEFUrlRequestClientComponent_Client", 0),
		/*3*/ imports.NewTable("CEFUrlRequestClientComponent_Create", 0),
		/*4*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnCreateURLRequest", 0),
		/*5*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnDownloadData", 0),
		/*6*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnDownloadProgress", 0),
		/*7*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnGetAuthCredentials", 0),
		/*8*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnRequestComplete", 0),
		/*9*/ imports.NewTable("CEFUrlRequestClientComponent_SetOnUploadProgress", 0),
		/*10*/ imports.NewTable("CEFUrlRequestClientComponent_ThreadID", 0),
	}
)

func urlRequestClientComponentImportAPI() *imports.Imports {
	if urlRequestClientComponentImport == nil {
		urlRequestClientComponentImport = NewDefaultImports()
		urlRequestClientComponentImport.SetImportTable(urlRequestClientComponentImportTables)
		urlRequestClientComponentImportTables = nil
	}
	return urlRequestClientComponentImport
}
