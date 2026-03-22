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

// ICefMediaSource Parent: ICefBaseRefCounted
type ICefMediaSource interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the ID (media source URN or URL) for this source.
	GetId() string // function
	// IsCastSource
	//  Returns true (1) if this source outputs its content via Cast.
	IsCastSource() bool // function
	// IsDialSource
	//  Returns true (1) if this source outputs its content via DIAL.
	IsDialSource() bool // function
}

// ICefMediaSourceRef Parent: ICefMediaSource ICefBaseRefCountedRef
type ICefMediaSourceRef interface {
	ICefMediaSource
	ICefBaseRefCountedRef
	AsIntfMediaSource() uintptr
}

type TCefMediaSourceRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMediaSourceRef) GetId() string {
	if !m.IsValid() {
		return ""
	}
	r := cefMediaSourceRefAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCefMediaSourceRef) IsCastSource() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMediaSourceRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMediaSourceRef) IsDialSource() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMediaSourceRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMediaSourceRef) AsIntfMediaSource() uintptr {
	return m.GetIntfPointer(0)
}

// MediaSourceRef  is static instance
var MediaSourceRef _MediaSourceRefClass

// _MediaSourceRefClass is class type defined by TCefMediaSourceRef
type _MediaSourceRefClass uintptr

func (_MediaSourceRefClass) UnWrap(data uintptr) (result ICefMediaSource) {
	var resultPtr uintptr
	cefMediaSourceRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSourceRef(resultPtr)
	return
}

// NewMediaSourceRef class constructor
func NewMediaSourceRef(data uintptr) ICefMediaSourceRef {
	var mediaSourcePtr uintptr // ICefMediaSource
	r := cefMediaSourceRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&mediaSourcePtr)))
	ret := AsCefMediaSourceRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaSourcePtr)
	}
	return ret
}

var (
	cefMediaSourceRefOnce   base.Once
	cefMediaSourceRefImport *imports.Imports = nil
)

func cefMediaSourceRefAPI() *imports.Imports {
	cefMediaSourceRefOnce.Do(func() {
		cefMediaSourceRefImport = api.NewDefaultImports()
		cefMediaSourceRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMediaSourceRef_Create", 0), // constructor NewMediaSourceRef
			/* 1 */ imports.NewTable("TCefMediaSourceRef_GetId", 0), // function GetId
			/* 2 */ imports.NewTable("TCefMediaSourceRef_IsCastSource", 0), // function IsCastSource
			/* 3 */ imports.NewTable("TCefMediaSourceRef_IsDialSource", 0), // function IsDialSource
			/* 4 */ imports.NewTable("TCefMediaSourceRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefMediaSourceRefImport
}
