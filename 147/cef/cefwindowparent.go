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
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

// ICEFWindowParent Parent: ICEFWinControl
type ICEFWindowParent interface {
	ICEFWinControl
}

type TCEFWindowParent struct {
	TCEFWinControl
}

// NewWindowParent class constructor
func NewWindowParent(theOwner lcl.IComponent) ICEFWindowParent {
	r := cEFWindowParentAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCEFWindowParent(r)
}

// NewWindowParentParented class constructor
func NewWindowParentParented(parentWindow types.HWND) ICEFWindowParent {
	r := cEFWindowParentAPI().SysCallN(1, uintptr(parentWindow))
	return AsCEFWindowParent(r)
}

var (
	cEFWindowParentOnce   base.Once
	cEFWindowParentImport *imports.Imports = nil
)

func cEFWindowParentAPI() *imports.Imports {
	cEFWindowParentOnce.Do(func() {
		cEFWindowParentImport = api.NewDefaultImports()
		cEFWindowParentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFWindowParent_Create", 0), // constructor NewWindowParent
			/* 1 */ imports.NewTable("TCEFWindowParent_CreateParented", 0), // constructor NewWindowParentParented
		}
	})
	return cEFWindowParentImport
}
