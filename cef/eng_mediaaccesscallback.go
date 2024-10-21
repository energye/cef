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

// ICefMediaAccessCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of media access permission requests.
//	This record is declared twice with almost identical parameters. "allowed_permissions" is defined as int and uint32.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_access_handler_capi.h">CEF source file: /include/capi/cef_media_access_handler_capi.h (cef_media_access_callback_t))</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_media_access_callback_t))</a>
type ICefMediaAccessCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Call to allow or deny media access. If this callback was initiated in response to a getUserMedia (indicated by CEF_MEDIA_PERMISSION_DEVICE_AUDIO_CAPTURE and/or CEF_MEDIA_PERMISSION_DEVICE_VIDEO_CAPTURE being set) then |allowed_permissions| must match |required_permissions| passed to OnRequestMediaAccessPermission.
	Cont(allowedpermissions TCefMediaAccessPermissionTypes) // procedure
	// Cancel
	//  Cancel the media access request.
	Cancel() // procedure
}

// TCefMediaAccessCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of media access permission requests.
//	This record is declared twice with almost identical parameters. "allowed_permissions" is defined as int and uint32.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_access_handler_capi.h">CEF source file: /include/capi/cef_media_access_handler_capi.h (cef_media_access_callback_t))</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_media_access_callback_t))</a>
type TCefMediaAccessCallback struct {
	TCefBaseRefCounted
}

// MediaAccessCallbackRef -> ICefMediaAccessCallback
var MediaAccessCallbackRef mediaAccessCallback

// mediaAccessCallback TCefMediaAccessCallback Ref
type mediaAccessCallback uintptr

func (m *mediaAccessCallback) UnWrap(data uintptr) ICefMediaAccessCallback {
	var resultCefMediaAccessCallback uintptr
	mediaAccessCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefMediaAccessCallback)))
	return AsCefMediaAccessCallback(resultCefMediaAccessCallback)
}

func (m *TCefMediaAccessCallback) Cont(allowedpermissions TCefMediaAccessPermissionTypes) {
	mediaAccessCallbackImportAPI().SysCallN(1, m.Instance(), uintptr(allowedpermissions))
}

func (m *TCefMediaAccessCallback) Cancel() {
	mediaAccessCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	mediaAccessCallbackImport       *imports.Imports = nil
	mediaAccessCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMediaAccessCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefMediaAccessCallback_Cont", 0),
		/*2*/ imports.NewTable("CefMediaAccessCallback_UnWrap", 0),
	}
)

func mediaAccessCallbackImportAPI() *imports.Imports {
	if mediaAccessCallbackImport == nil {
		mediaAccessCallbackImport = NewDefaultImports()
		mediaAccessCallbackImport.SetImportTable(mediaAccessCallbackImportTables)
		mediaAccessCallbackImportTables = nil
	}
	return mediaAccessCallbackImport
}
