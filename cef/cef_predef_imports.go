//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

var (
	predefImport       *imports.Imports = nil
	predefImportTables                  = []*imports.Table{
		// TCefPrintSettings
		/*0*/ imports.NewTable("CefPrintSettings_SetPageRanges", 0),
		/*1*/ imports.NewTable("CefPrintSettings_GetPageRanges", 0),
		// TCefBrowserHost
		/*2*/ imports.NewTable("ICefBrowserHost_IMESetComposition", 0),
		// TBufferPanel
		/*3*/ imports.NewTable("BufferPanel_ChangeCompositionRange", 0),
		// TChromiumCore
		/*4*/ imports.NewTable("ChromiumCore_IMESetComposition", 0),
		// TCefX509Certificate
		/*5*/ imports.NewTable("ICefX509Certificate_GetDEREncodedIssuerChain", 0),
		/*6*/ imports.NewTable("ICefX509Certificate_GetPEMEncodedIssuerChain", 0),
		// TCefDisplay
		/*7*/ imports.NewTable("CefDisplay_GetAlls", 0),
		// TCefPostData
		/*8*/ imports.NewTable("CefPostData_GetElements", 0),
		// TChromiumCore
		/*9*/ imports.NewTable("ChromiumCore_SetCookie", 0),
		// Energy Other
		/*10*/ imports.NewTable("Eng_SetCommandLine", 0),
		/*11*/ imports.NewTable("Eng_AddCrDelegate", 0),
		// GlobalApp
		/*12*/ imports.NewTable("GlobalApp_SetGlobalCEFApp", 0),
		/*13*/ imports.NewTable("GlobalApp_DestroyGlobalCEFApp", 0),
		/*14*/ imports.NewTable("GlobalApp_SpecificVersion", 0),
		// Global TCEFWorkScheduler
		/*15*/ imports.NewTable("GlobalCEFWorkScheduler_SetGlobalWorkScheduler", 0),
		/*16*/ imports.NewTable("GlobalCEFWorkScheduler_DestroyGlobalCEFWorkScheduler", 0),
		// TCefRequest
		/*17*/ imports.NewTable("CefRequest_GetHeaderMap", 0),
		// TCefResponse
		/*18*/ imports.NewTable("CefResponse_GetHeaderMap", 0),
	}
)

func predefImportAPI() *imports.Imports {
	if predefImport == nil {
		predefImport = api.NewDefaultImports()
		predefImport.SetImportTable(predefImportTables)
		predefImportTables = nil
	}
	return predefImport
}
