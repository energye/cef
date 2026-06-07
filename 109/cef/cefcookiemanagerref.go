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

// ICefCookieManager Parent: ICefBaseRefCounted
type ICefCookieManager interface {
	ICefBaseRefCounted
	VisitAllCookies(visitor IEngCookieVisitor) bool                                       // function
	VisitUrlCookies(url string, includeHttpOnly bool, visitor IEngCookieVisitor) bool     // function
	SetCookie(args TCefCookieManagerRefSetCookieArgs) bool                                // function
	DeleteCookies(url string, cookieName string, callback IEngDeleteCookiesCallback) bool // function
	FlushStore(callback IEngCompletionCallback) bool                                      // function
}

// ICefCookieManagerRef Parent: ICefCookieManager ICefBaseRefCountedRef
type ICefCookieManagerRef interface {
	ICefCookieManager
	ICefBaseRefCountedRef
	AsIntfCookieManager() uintptr
}

type TCefCookieManagerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefCookieManagerRef) VisitAllCookies(visitor IEngCookieVisitor) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCookieManagerRefAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(visitor))
	return api.GoBool(r)
}

func (m *TCefCookieManagerRef) VisitUrlCookies(url string, includeHttpOnly bool, visitor IEngCookieVisitor) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCookieManagerRefAPI().SysCallN(2, m.Instance(), api.PasStr(url), api.PasBool(includeHttpOnly), base.GetObjectUintptr(visitor))
	return api.GoBool(r)
}

func (m *TCefCookieManagerRef) SetCookie(args TCefCookieManagerRefSetCookieArgs) bool {
	if !m.IsValid() {
		return false
	}
	argsPtr := args.ToPas()
	r := cefCookieManagerRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(argsPtr)))
	return api.GoBool(r)
}

func (m *TCefCookieManagerRef) DeleteCookies(url string, cookieName string, callback IEngDeleteCookiesCallback) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCookieManagerRefAPI().SysCallN(4, m.Instance(), api.PasStr(url), api.PasStr(cookieName), base.GetObjectUintptr(callback))
	return api.GoBool(r)
}

func (m *TCefCookieManagerRef) FlushStore(callback IEngCompletionCallback) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCookieManagerRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(callback))
	return api.GoBool(r)
}

func (m *TCefCookieManagerRef) AsIntfCookieManager() uintptr {
	return m.GetIntfPointer(0)
}

// CookieManagerRef  is static instance
var CookieManagerRef _CookieManagerRefClass

// _CookieManagerRefClass is class type defined by TCefCookieManagerRef
type _CookieManagerRefClass uintptr

func (_CookieManagerRefClass) UnWrap(data uintptr) (result ICefCookieManager) {
	var resultPtr uintptr
	cefCookieManagerRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCookieManagerRef(resultPtr)
	return
}

func (_CookieManagerRefClass) Global(callback IEngCompletionCallback) (result ICefCookieManager) {
	var resultPtr uintptr
	cefCookieManagerRefAPI().SysCallN(7, base.GetObjectUintptr(callback), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCookieManagerRef(resultPtr)
	return
}

// NewCookieManagerRef class constructor
func NewCookieManagerRef(data uintptr) ICefCookieManagerRef {
	var cookieManagerPtr uintptr // ICefCookieManager
	r := cefCookieManagerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&cookieManagerPtr)))
	ret := AsCefCookieManagerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieManagerPtr)
	}
	return ret
}

var (
	cefCookieManagerRefOnce   base.Once
	cefCookieManagerRefImport *imports.Imports = nil
)

func cefCookieManagerRefAPI() *imports.Imports {
	cefCookieManagerRefOnce.Do(func() {
		cefCookieManagerRefImport = api.NewDefaultImports()
		cefCookieManagerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCookieManagerRef_Create", 0), // constructor NewCookieManagerRef
			/* 1 */ imports.NewTable("TCefCookieManagerRef_VisitAllCookies", 0), // function VisitAllCookies
			/* 2 */ imports.NewTable("TCefCookieManagerRef_VisitUrlCookies", 0), // function VisitUrlCookies
			/* 3 */ imports.NewTable("TCefCookieManagerRef_SetCookie", 0), // function SetCookie
			/* 4 */ imports.NewTable("TCefCookieManagerRef_DeleteCookies", 0), // function DeleteCookies
			/* 5 */ imports.NewTable("TCefCookieManagerRef_FlushStore", 0), // function FlushStore
			/* 6 */ imports.NewTable("TCefCookieManagerRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefCookieManagerRef_Global", 0), // static function Global
		}
	})
	return cefCookieManagerRefImport
}
