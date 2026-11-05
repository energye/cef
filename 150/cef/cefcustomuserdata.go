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

// ICefCustomUserData Parent: ICefBaseRefCountedOwn
type ICefCustomUserData interface {
	ICefBaseRefCountedOwn
	GetUserDataType() uintptr // function
	GetUserData() uintptr     // function
	AsIntfCustomUserData() uintptr
}

type TCefCustomUserData struct {
	TCefBaseRefCountedOwn
}

func (m *TCefCustomUserData) GetUserDataType() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomUserDataAPI().SysCallN(1, m.Instance())
	return uintptr(r)
}

func (m *TCefCustomUserData) GetUserData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomUserDataAPI().SysCallN(2, m.Instance())
	return uintptr(r)
}

func (m *TCefCustomUserData) AsIntfCustomUserData() uintptr {
	return m.GetIntfPointer(0)
}

// CustomUserData  is static instance
var CustomUserData _CustomUserDataClass

// _CustomUserDataClass is class type defined by TCefCustomUserData
type _CustomUserDataClass uintptr

func (_CustomUserDataClass) UnWrap(data uintptr) (result ICefCustomUserData) {
	var resultPtr uintptr
	cefCustomUserDataAPI().SysCallN(3, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCustomUserData(resultPtr)
	return
}

// NewCustomUserData class constructor
func NewCustomUserData(userDataType uintptr, userData uintptr) ICefCustomUserData {
	var customUserDataPtr uintptr // ICefCustomUserData
	r := cefCustomUserDataAPI().SysCallN(0, uintptr(userDataType), uintptr(userData), uintptr(base.UnsafePointer(&customUserDataPtr)))
	ret := AsCefCustomUserData(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customUserDataPtr)
	}
	return ret
}

var (
	cefCustomUserDataOnce   base.Once
	cefCustomUserDataImport *imports.Imports = nil
)

func cefCustomUserDataAPI() *imports.Imports {
	cefCustomUserDataOnce.Do(func() {
		cefCustomUserDataImport = api.NewDefaultImports()
		cefCustomUserDataImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomUserData_Create", 0), // constructor NewCustomUserData
			/* 1 */ imports.NewTable("TCefCustomUserData_GetUserDataType", 0), // function GetUserDataType
			/* 2 */ imports.NewTable("TCefCustomUserData_GetUserData", 0), // function GetUserData
			/* 3 */ imports.NewTable("TCefCustomUserData_UnWrap", 0), // static function UnWrap
		}
	})
	return cefCustomUserDataImport
}
