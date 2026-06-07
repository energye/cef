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

// ICefMediaRoute Parent: ICefBaseRefCounted
type ICefMediaRoute interface {
	ICefBaseRefCounted
	GetId() string                   // function
	GetSource() ICefMediaSource      // function
	GetSink() ICefMediaSink          // function
	SendRouteMessage(message string) // procedure
	Terminate()                      // procedure
}

// ICefMediaRouteRef Parent: ICefMediaRoute ICefBaseRefCountedRef
type ICefMediaRouteRef interface {
	ICefMediaRoute
	ICefBaseRefCountedRef
	AsIntfMediaRoute() uintptr
}

type TCefMediaRouteRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMediaRouteRef) GetId() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefMediaRouteRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefMediaRouteRef) GetSource() (result ICefMediaSource) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMediaRouteRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSourceRef(resultPtr)
	return
}

func (m *TCefMediaRouteRef) GetSink() (result ICefMediaSink) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMediaRouteRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSinkRef(resultPtr)
	return
}

func (m *TCefMediaRouteRef) SendRouteMessage(message string) {
	if !m.IsValid() {
		return
	}
	cefMediaRouteRefAPI().SysCallN(5, m.Instance(), api.PasStr(message))
}

func (m *TCefMediaRouteRef) Terminate() {
	if !m.IsValid() {
		return
	}
	cefMediaRouteRefAPI().SysCallN(6, m.Instance())
}

func (m *TCefMediaRouteRef) AsIntfMediaRoute() uintptr {
	return m.GetIntfPointer(0)
}

// MediaRouteRef  is static instance
var MediaRouteRef _MediaRouteRefClass

// _MediaRouteRefClass is class type defined by TCefMediaRouteRef
type _MediaRouteRefClass uintptr

func (_MediaRouteRefClass) UnWrap(data uintptr) (result ICefMediaRoute) {
	var resultPtr uintptr
	cefMediaRouteRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouteRef(resultPtr)
	return
}

// NewMediaRouteRef class constructor
func NewMediaRouteRef(data uintptr) ICefMediaRouteRef {
	var mediaRoutePtr uintptr // ICefMediaRoute
	r := cefMediaRouteRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&mediaRoutePtr)))
	ret := AsCefMediaRouteRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaRoutePtr)
	}
	return ret
}

var (
	cefMediaRouteRefOnce   base.Once
	cefMediaRouteRefImport *imports.Imports = nil
)

func cefMediaRouteRefAPI() *imports.Imports {
	cefMediaRouteRefOnce.Do(func() {
		cefMediaRouteRefImport = api.NewDefaultImports()
		cefMediaRouteRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMediaRouteRef_Create", 0), // constructor NewMediaRouteRef
			/* 1 */ imports.NewTable("TCefMediaRouteRef_GetId", 0), // function GetId
			/* 2 */ imports.NewTable("TCefMediaRouteRef_GetSource", 0), // function GetSource
			/* 3 */ imports.NewTable("TCefMediaRouteRef_GetSink", 0), // function GetSink
			/* 4 */ imports.NewTable("TCefMediaRouteRef_UnWrap", 0), // static function UnWrap
			/* 5 */ imports.NewTable("TCefMediaRouteRef_SendRouteMessage", 0), // procedure SendRouteMessage
			/* 6 */ imports.NewTable("TCefMediaRouteRef_Terminate", 0), // procedure Terminate
		}
	})
	return cefMediaRouteRefImport
}
