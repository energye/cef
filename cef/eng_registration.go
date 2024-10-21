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
//
//	Generic callback interface used for managing the lifespan of a registration.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_registration_capi.h">CEF source file: /include/capi/cef_registration_capi.h (cef_registration_t))</a>
type ICefRegistration interface {
	ICefBaseRefCounted
}

// TCefRegistration Parent: TCefBaseRefCounted
//
//	Generic callback interface used for managing the lifespan of a registration.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_registration_capi.h">CEF source file: /include/capi/cef_registration_capi.h (cef_registration_t))</a>
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
