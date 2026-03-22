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

// ICefRegistration Parent: ICefBaseRefCounted
type ICefRegistration interface {
	ICefBaseRefCounted
}

// ICefRegistrationRef Parent: ICefRegistration ICefBaseRefCountedRef
type ICefRegistrationRef interface {
	ICefRegistration
	ICefBaseRefCountedRef
	AsIntfRegistration() uintptr
}

type TCefRegistrationRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefRegistrationRef) AsIntfRegistration() uintptr {
	return m.GetIntfPointer(0)
}

// RegistrationRef  is static instance
var RegistrationRef _RegistrationRefClass

// _RegistrationRefClass is class type defined by TCefRegistrationRef
type _RegistrationRefClass uintptr

func (_RegistrationRefClass) UnWrap(data uintptr) (result ICefRegistration) {
	var resultPtr uintptr
	cefRegistrationRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

// NewRegistrationRef class constructor
func NewRegistrationRef(data uintptr) ICefRegistrationRef {
	var registrationPtr uintptr // ICefRegistration
	r := cefRegistrationRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&registrationPtr)))
	ret := AsCefRegistrationRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, registrationPtr)
	}
	return ret
}

var (
	cefRegistrationRefOnce   base.Once
	cefRegistrationRefImport *imports.Imports = nil
)

func cefRegistrationRefAPI() *imports.Imports {
	cefRegistrationRefOnce.Do(func() {
		cefRegistrationRefImport = api.NewDefaultImports()
		cefRegistrationRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefRegistrationRef_Create", 0), // constructor NewRegistrationRef
			/* 1 */ imports.NewTable("TCefRegistrationRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefRegistrationRefImport
}
