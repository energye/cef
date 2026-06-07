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

// ICefComponent Parent: ICefBaseRefCounted
type ICefComponent interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the unique identifier for this component.
	GetId() string // function
	// GetName
	//  Returns the human-readable name of this component. Returns an NULL string
	//  if the component is not installed.
	GetName() string // function
	// GetVersion
	//  Returns the version of this component as a string (e.g., "1.2.3.4").
	//  Returns an NULL string if the component is not installed.
	GetVersion() string // function
	// GetState
	//  Returns the state of this component at the time this object was created. A
	//  component is considered installed when its state is one of:
	//  CEF_COMPONENT_STATE_UPDATED, CEF_COMPONENT_STATE_UP_TO_DATE, or
	//  CEF_COMPONENT_STATE_RUN.
	GetState() cefTypes.TCefComponentState // function
}

// ICefComponentRef Parent: ICefComponent ICefBaseRefCountedRef
type ICefComponentRef interface {
	ICefComponent
	ICefBaseRefCountedRef
	AsIntfComponent() uintptr
}

type TCefComponentRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefComponentRef) GetId() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefComponentRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefComponentRef) GetName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefComponentRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefComponentRef) GetVersion() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefComponentRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefComponentRef) GetState() cefTypes.TCefComponentState {
	if !m.IsValid() {
		return 0
	}
	r := cefComponentRefAPI().SysCallN(4, m.Instance())
	return cefTypes.TCefComponentState(r)
}

func (m *TCefComponentRef) AsIntfComponent() uintptr {
	return m.GetIntfPointer(0)
}

// ComponentRef  is static instance
var ComponentRef _ComponentRefClass

// _ComponentRefClass is class type defined by TCefComponentRef
type _ComponentRefClass uintptr

func (_ComponentRefClass) UnWrap(data uintptr) (result ICefComponent) {
	var resultPtr uintptr
	cefComponentRefAPI().SysCallN(5, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefComponentRef(resultPtr)
	return
}

// NewComponentRef class constructor
func NewComponentRef(data uintptr) ICefComponentRef {
	var componentPtr uintptr // ICefComponent
	r := cefComponentRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&componentPtr)))
	ret := AsCefComponentRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, componentPtr)
	}
	return ret
}

var (
	cefComponentRefOnce   base.Once
	cefComponentRefImport *imports.Imports = nil
)

func cefComponentRefAPI() *imports.Imports {
	cefComponentRefOnce.Do(func() {
		cefComponentRefImport = api.NewDefaultImports()
		cefComponentRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefComponentRef_Create", 0), // constructor NewComponentRef
			/* 1 */ imports.NewTable("TCefComponentRef_GetId", 0), // function GetId
			/* 2 */ imports.NewTable("TCefComponentRef_GetName", 0), // function GetName
			/* 3 */ imports.NewTable("TCefComponentRef_GetVersion", 0), // function GetVersion
			/* 4 */ imports.NewTable("TCefComponentRef_GetState", 0), // function GetState
			/* 5 */ imports.NewTable("TCefComponentRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefComponentRefImport
}
