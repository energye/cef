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

// ICEFLinkedWindowParent Parent: ICEFLinkedWinControlBase
//
//	This component can be used by VCL and LCL applications. It has the
//	same purpose as TCEFWindowParent but it has a Chromium property to
//	link it directly to a TChromium component.
//	TCEFLinkedWindowParent resizes the child controls created by CEF
//	for browsers in normal mode and sets the browser focus using the
//	linked TChromium component. TCEFWindowParent and TCEFLinkedWindowParent
//	work fine in Windows and you can used any of them but you can't use
//	TCEFWindowParent in Linux or MacOS.
type ICEFLinkedWindowParent interface {
	ICEFLinkedWinControlBase
	// Chromium
	//  TChromium instance used by this component.
	Chromium() IChromium // property
	// SetChromium Set Chromium
	SetChromium(AValue IChromium) // property
}

// TCEFLinkedWindowParent Parent: TCEFLinkedWinControlBase
//
//	This component can be used by VCL and LCL applications. It has the
//	same purpose as TCEFWindowParent but it has a Chromium property to
//	link it directly to a TChromium component.
//	TCEFLinkedWindowParent resizes the child controls created by CEF
//	for browsers in normal mode and sets the browser focus using the
//	linked TChromium component. TCEFWindowParent and TCEFLinkedWindowParent
//	work fine in Windows and you can used any of them but you can't use
//	TCEFWindowParent in Linux or MacOS.
type TCEFLinkedWindowParent struct {
	TCEFLinkedWinControlBase
}

func NewCEFLinkedWindowParent(aOwner IComponent) ICEFLinkedWindowParent {
	r1 := linkedWindowParentImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsCEFLinkedWindowParent(r1)
}

func (m *TCEFLinkedWindowParent) Chromium() IChromium {
	r1 := linkedWindowParentImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsChromium(r1)
}

func (m *TCEFLinkedWindowParent) SetChromium(AValue IChromium) {
	linkedWindowParentImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

var (
	linkedWindowParentImport       *imports.Imports = nil
	linkedWindowParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFLinkedWindowParent_Chromium", 0),
		/*1*/ imports.NewTable("CEFLinkedWindowParent_Create", 0),
	}
)

func linkedWindowParentImportAPI() *imports.Imports {
	if linkedWindowParentImport == nil {
		linkedWindowParentImport = NewDefaultImports()
		linkedWindowParentImport.SetImportTable(linkedWindowParentImportTables)
		linkedWindowParentImportTables = nil
	}
	return linkedWindowParentImport
}
