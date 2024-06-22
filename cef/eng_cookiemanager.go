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

// ICefCookieManager Parent: ICefBaseRefCounted
type ICefCookieManager interface {
	ICefBaseRefCounted
	VisitAllCookies(visitor ICefCookieVisitor) bool                                                                                                                                                                             // function
	VisitUrlCookies(url string, includeHttpOnly bool, visitor ICefCookieVisitor) bool                                                                                                                                           // function
	SetCookie(url, name, value, domain, path string, secure, httponly, hasExpires bool, creation, lastAccess, expires TDateTime, samesite TCefCookieSameSite, priority TCefCookiePriority, callback ICefSetCookieCallback) bool // function
	DeleteCookies(url, cookieName string, callback ICefDeleteCookiesCallback) bool                                                                                                                                              // function
	FlushStore(callback ICefCompletionCallback) bool                                                                                                                                                                            // function
}

// TCefCookieManager Parent: TCefBaseRefCounted
type TCefCookieManager struct {
	TCefBaseRefCounted
}

// CookieManagerRef -> ICefCookieManager
var CookieManagerRef cookieManager

// cookieManager TCefCookieManager Ref
type cookieManager uintptr

func (m *cookieManager) UnWrap(data uintptr) ICefCookieManager {
	var resultCefCookieManager uintptr
	cookieManagerImportAPI().SysCallN(4, uintptr(data), uintptr(unsafePointer(&resultCefCookieManager)))
	return AsCefCookieManager(resultCefCookieManager)
}

func (m *cookieManager) Global(callback ICefCompletionCallback) ICefCookieManager {
	var resultCefCookieManager uintptr
	cookieManagerImportAPI().SysCallN(2, GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefCookieManager)))
	return AsCefCookieManager(resultCefCookieManager)
}

func (m *TCefCookieManager) VisitAllCookies(visitor ICefCookieVisitor) bool {
	r1 := cookieManagerImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(visitor))
	return GoBool(r1)
}

func (m *TCefCookieManager) VisitUrlCookies(url string, includeHttpOnly bool, visitor ICefCookieVisitor) bool {
	r1 := cookieManagerImportAPI().SysCallN(6, m.Instance(), PascalStr(url), PascalBool(includeHttpOnly), GetObjectUintptr(visitor))
	return GoBool(r1)
}

func (m *TCefCookieManager) SetCookie(url, name, value, domain, path string, secure, httponly, hasExpires bool, creation, lastAccess, expires TDateTime, samesite TCefCookieSameSite, priority TCefCookiePriority, callback ICefSetCookieCallback) bool {
	r1 := cookieManagerImportAPI().SysCallN(3, m.Instance(), PascalStr(url), PascalStr(name), PascalStr(value), PascalStr(domain), PascalStr(path), PascalBool(secure), PascalBool(httponly), PascalBool(hasExpires), uintptr(unsafePointer(&creation)), uintptr(unsafePointer(&lastAccess)), uintptr(unsafePointer(&expires)), uintptr(samesite), uintptr(priority), GetObjectUintptr(callback))
	return GoBool(r1)
}

func (m *TCefCookieManager) DeleteCookies(url, cookieName string, callback ICefDeleteCookiesCallback) bool {
	r1 := cookieManagerImportAPI().SysCallN(0, m.Instance(), PascalStr(url), PascalStr(cookieName), GetObjectUintptr(callback))
	return GoBool(r1)
}

func (m *TCefCookieManager) FlushStore(callback ICefCompletionCallback) bool {
	r1 := cookieManagerImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(callback))
	return GoBool(r1)
}

var (
	cookieManagerImport       *imports.Imports = nil
	cookieManagerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCookieManager_DeleteCookies", 0),
		/*1*/ imports.NewTable("CefCookieManager_FlushStore", 0),
		/*2*/ imports.NewTable("CefCookieManager_Global", 0),
		/*3*/ imports.NewTable("CefCookieManager_SetCookie", 0),
		/*4*/ imports.NewTable("CefCookieManager_UnWrap", 0),
		/*5*/ imports.NewTable("CefCookieManager_VisitAllCookies", 0),
		/*6*/ imports.NewTable("CefCookieManager_VisitUrlCookies", 0),
	}
)

func cookieManagerImportAPI() *imports.Imports {
	if cookieManagerImport == nil {
		cookieManagerImport = NewDefaultImports()
		cookieManagerImport.SetImportTable(cookieManagerImportTables)
		cookieManagerImportTables = nil
	}
	return cookieManagerImport
}
