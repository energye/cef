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
)

// ICefMediaSink Parent: ICefBaseRefCounted
type ICefMediaSink interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the ID for this sink.
	GetId() string // function
	// GetName
	//  Returns the name of this sink.
	GetName() string // function
	// GetIconType
	//  Returns the icon type for this sink.
	GetIconType() cefTypes.TCefMediaSinkIconType // function
	// IsCastSink
	//  Returns true (1) if this sink accepts content via Cast.
	IsCastSink() bool // function
	// IsDialSink
	//  Returns true (1) if this sink accepts content via DIAL.
	IsDialSink() bool // function
	// IsCompatibleWith
	//  Returns true (1) if this sink is compatible with |source|.
	IsCompatibleWith(source ICefMediaSource) bool // function
	// GetDeviceInfo
	//  Asynchronously retrieves device info.
	GetDeviceInfo(callback IEngMediaSinkDeviceInfoCallback) // procedure
}

// ICefMediaSinkRef Parent: ICefMediaSink ICefBaseRefCountedRef
type ICefMediaSinkRef interface {
	ICefMediaSink
	ICefBaseRefCountedRef
	AsIntfMediaSink() uintptr
}

type TCefMediaSinkRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMediaSinkRef) GetId() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefMediaSinkRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefMediaSinkRef) GetName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefMediaSinkRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefMediaSinkRef) GetIconType() cefTypes.TCefMediaSinkIconType {
	if !m.IsValid() {
		return 0
	}
	r := cefMediaSinkRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefMediaSinkIconType(r)
}

func (m *TCefMediaSinkRef) IsCastSink() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMediaSinkRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMediaSinkRef) IsDialSink() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMediaSinkRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMediaSinkRef) IsCompatibleWith(source ICefMediaSource) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMediaSinkRefAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(source))
	return api.GoBool(r)
}

func (m *TCefMediaSinkRef) GetDeviceInfo(callback IEngMediaSinkDeviceInfoCallback) {
	if !m.IsValid() {
		return
	}
	cefMediaSinkRefAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefMediaSinkRef) AsIntfMediaSink() uintptr {
	return m.GetIntfPointer(0)
}

// MediaSinkRef  is static instance
var MediaSinkRef _MediaSinkRefClass

// _MediaSinkRefClass is class type defined by TCefMediaSinkRef
type _MediaSinkRefClass uintptr

func (_MediaSinkRefClass) UnWrap(data uintptr) (result ICefMediaSink) {
	var resultPtr uintptr
	cefMediaSinkRefAPI().SysCallN(7, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSinkRef(resultPtr)
	return
}

// NewMediaSinkRef class constructor
func NewMediaSinkRef(data uintptr) ICefMediaSinkRef {
	var mediaSinkPtr uintptr // ICefMediaSink
	r := cefMediaSinkRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&mediaSinkPtr)))
	ret := AsCefMediaSinkRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaSinkPtr)
	}
	return ret
}

var (
	cefMediaSinkRefOnce   base.Once
	cefMediaSinkRefImport *imports.Imports = nil
)

func cefMediaSinkRefAPI() *imports.Imports {
	cefMediaSinkRefOnce.Do(func() {
		cefMediaSinkRefImport = api.NewDefaultImports()
		cefMediaSinkRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMediaSinkRef_Create", 0), // constructor NewMediaSinkRef
			/* 1 */ imports.NewTable("TCefMediaSinkRef_GetId", 0), // function GetId
			/* 2 */ imports.NewTable("TCefMediaSinkRef_GetName", 0), // function GetName
			/* 3 */ imports.NewTable("TCefMediaSinkRef_GetIconType", 0), // function GetIconType
			/* 4 */ imports.NewTable("TCefMediaSinkRef_IsCastSink", 0), // function IsCastSink
			/* 5 */ imports.NewTable("TCefMediaSinkRef_IsDialSink", 0), // function IsDialSink
			/* 6 */ imports.NewTable("TCefMediaSinkRef_IsCompatibleWith", 0), // function IsCompatibleWith
			/* 7 */ imports.NewTable("TCefMediaSinkRef_UnWrap", 0), // static function UnWrap
			/* 8 */ imports.NewTable("TCefMediaSinkRef_GetDeviceInfo", 0), // procedure GetDeviceInfo
		}
	})
	return cefMediaSinkRefImport
}
