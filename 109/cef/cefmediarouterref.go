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

// ICefMediaRouter Parent: ICefBaseRefCounted
type ICefMediaRouter interface {
	ICefBaseRefCounted
	AddObserver(observer IEngMediaObserver) ICefRegistration                                       // function
	GetSource(urn string) ICefMediaSource                                                          // function
	NotifyCurrentSinks()                                                                           // procedure
	CreateRoute(source ICefMediaSource, sink ICefMediaSink, callback IEngMediaRouteCreateCallback) // procedure
	NotifyCurrentRoutes()                                                                          // procedure
}

// ICefMediaRouterRef Parent: ICefMediaRouter ICefBaseRefCountedRef
type ICefMediaRouterRef interface {
	ICefMediaRouter
	ICefBaseRefCountedRef
	AsIntfMediaRouter() uintptr
}

type TCefMediaRouterRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMediaRouterRef) AddObserver(observer IEngMediaObserver) (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMediaRouterRefAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(observer), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TCefMediaRouterRef) GetSource(urn string) (result ICefMediaSource) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMediaRouterRefAPI().SysCallN(2, m.Instance(), api.PasStr(urn), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSourceRef(resultPtr)
	return
}

func (m *TCefMediaRouterRef) NotifyCurrentSinks() {
	if !m.IsValid() {
		return
	}
	cefMediaRouterRefAPI().SysCallN(5, m.Instance())
}

func (m *TCefMediaRouterRef) CreateRoute(source ICefMediaSource, sink ICefMediaSink, callback IEngMediaRouteCreateCallback) {
	if !m.IsValid() {
		return
	}
	cefMediaRouterRefAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(source), base.GetObjectUintptr(sink), base.GetObjectUintptr(callback))
}

func (m *TCefMediaRouterRef) NotifyCurrentRoutes() {
	if !m.IsValid() {
		return
	}
	cefMediaRouterRefAPI().SysCallN(7, m.Instance())
}

func (m *TCefMediaRouterRef) AsIntfMediaRouter() uintptr {
	return m.GetIntfPointer(0)
}

// MediaRouterRef  is static instance
var MediaRouterRef _MediaRouterRefClass

// _MediaRouterRefClass is class type defined by TCefMediaRouterRef
type _MediaRouterRefClass uintptr

func (_MediaRouterRefClass) UnWrap(data uintptr) (result ICefMediaRouter) {
	var resultPtr uintptr
	cefMediaRouterRefAPI().SysCallN(3, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouterRef(resultPtr)
	return
}

func (_MediaRouterRefClass) Global() (result ICefMediaRouter) {
	var resultPtr uintptr
	cefMediaRouterRefAPI().SysCallN(4, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouterRef(resultPtr)
	return
}

// NewMediaRouterRef class constructor
func NewMediaRouterRef(data uintptr) ICefMediaRouterRef {
	var mediaRouterPtr uintptr // ICefMediaRouter
	r := cefMediaRouterRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&mediaRouterPtr)))
	ret := AsCefMediaRouterRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaRouterPtr)
	}
	return ret
}

var (
	cefMediaRouterRefOnce   base.Once
	cefMediaRouterRefImport *imports.Imports = nil
)

func cefMediaRouterRefAPI() *imports.Imports {
	cefMediaRouterRefOnce.Do(func() {
		cefMediaRouterRefImport = api.NewDefaultImports()
		cefMediaRouterRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMediaRouterRef_Create", 0), // constructor NewMediaRouterRef
			/* 1 */ imports.NewTable("TCefMediaRouterRef_AddObserver", 0), // function AddObserver
			/* 2 */ imports.NewTable("TCefMediaRouterRef_GetSource", 0), // function GetSource
			/* 3 */ imports.NewTable("TCefMediaRouterRef_UnWrap", 0), // static function UnWrap
			/* 4 */ imports.NewTable("TCefMediaRouterRef_Global", 0), // static function Global
			/* 5 */ imports.NewTable("TCefMediaRouterRef_NotifyCurrentSinks", 0), // procedure NotifyCurrentSinks
			/* 6 */ imports.NewTable("TCefMediaRouterRef_CreateRoute", 0), // procedure CreateRoute
			/* 7 */ imports.NewTable("TCefMediaRouterRef_NotifyCurrentRoutes", 0), // procedure NotifyCurrentRoutes
		}
	})
	return cefMediaRouterRefImport
}
