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

// ICefMenuButtonPressedLock Parent: ICefBaseRefCounted
type ICefMenuButtonPressedLock interface {
	ICefBaseRefCounted
}

// ICefMenuButtonPressedLockRef Parent: ICefMenuButtonPressedLock ICefBaseRefCountedRef
type ICefMenuButtonPressedLockRef interface {
	ICefMenuButtonPressedLock
	ICefBaseRefCountedRef
	AsIntfMenuButtonPressedLock() uintptr
}

type TCefMenuButtonPressedLockRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMenuButtonPressedLockRef) AsIntfMenuButtonPressedLock() uintptr {
	return m.GetIntfPointer(0)
}

// MenuButtonPressedLockRef  is static instance
var MenuButtonPressedLockRef _MenuButtonPressedLockRefClass

// _MenuButtonPressedLockRefClass is class type defined by TCefMenuButtonPressedLockRef
type _MenuButtonPressedLockRefClass uintptr

// UnWrap
//
//	Returns a ICefMenuButtonPressedLock instance using a PCefMenuButtonPressedLock data pointer.
func (_MenuButtonPressedLockRefClass) UnWrap(data uintptr) (result ICefMenuButtonPressedLock) {
	var resultPtr uintptr
	cefMenuButtonPressedLockRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuButtonPressedLockRef(resultPtr)
	return
}

// NewMenuButtonPressedLockRef class constructor
func NewMenuButtonPressedLockRef(data uintptr) ICefMenuButtonPressedLockRef {
	var menuButtonPressedLockPtr uintptr // ICefMenuButtonPressedLock
	r := cefMenuButtonPressedLockRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&menuButtonPressedLockPtr)))
	ret := AsCefMenuButtonPressedLockRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, menuButtonPressedLockPtr)
	}
	return ret
}

var (
	cefMenuButtonPressedLockRefOnce   base.Once
	cefMenuButtonPressedLockRefImport *imports.Imports = nil
)

func cefMenuButtonPressedLockRefAPI() *imports.Imports {
	cefMenuButtonPressedLockRefOnce.Do(func() {
		cefMenuButtonPressedLockRefImport = api.NewDefaultImports()
		cefMenuButtonPressedLockRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMenuButtonPressedLockRef_Create", 0), // constructor NewMenuButtonPressedLockRef
			/* 1 */ imports.NewTable("TCefMenuButtonPressedLockRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefMenuButtonPressedLockRefImport
}
