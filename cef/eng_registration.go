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

// ICefRegistration Parent: ICefBaseRefCounted
type ICefRegistration interface {
	ICefBaseRefCounted
}

// TCefRegistration Parent: TCefBaseRefCounted
type TCefRegistration struct {
	TCefBaseRefCounted
}

// RegistrationRef -> ICefRegistration
var RegistrationRef registration

// registration TCefRegistration Ref
type registration uintptr

func (m *registration) UnWrap(data uintptr) ICefRegistration {
	var resultCefRegistration uintptr
	registrationImportAPI().SysCallN(0, uintptr(data), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

var (
	registrationImport       *imports.Imports = nil
	registrationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefRegistration_UnWrap", 0),
	}
)

func registrationImportAPI() *imports.Imports {
	if registrationImport == nil {
		registrationImport = NewDefaultImports()
		registrationImport.SetImportTable(registrationImportTables)
		registrationImportTables = nil
	}
	return registrationImport
}
