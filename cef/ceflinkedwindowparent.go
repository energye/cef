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
)

// ICEFLinkedWindowParent Parent: ICEFLinkedWinControlBase
type ICEFLinkedWindowParent interface {
	ICEFLinkedWinControlBase
	Chromium() IChromium         // property Chromium Getter
	SetChromium(value IChromium) // property Chromium Setter
}

type TCEFLinkedWindowParent struct {
	TCEFLinkedWinControlBase
}

func (m *TCEFLinkedWindowParent) Chromium() IChromium {
	if !m.IsValid() {
		return nil
	}
	r := cEFLinkedWindowParentAPI().SysCallN(1, 0, m.Instance())
	return AsChromium(r)
}

func (m *TCEFLinkedWindowParent) SetChromium(value IChromium) {
	if !m.IsValid() {
		return
	}
	cEFLinkedWindowParentAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLinkedWindowParent class constructor
func NewLinkedWindowParent(owner lcl.IComponent) ICEFLinkedWindowParent {
	r := cEFLinkedWindowParentAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCEFLinkedWindowParent(r)
}

var (
	cEFLinkedWindowParentOnce   base.Once
	cEFLinkedWindowParentImport *imports.Imports = nil
)

func cEFLinkedWindowParentAPI() *imports.Imports {
	cEFLinkedWindowParentOnce.Do(func() {
		cEFLinkedWindowParentImport = api.NewDefaultImports()
		cEFLinkedWindowParentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFLinkedWindowParent_Create", 0), // constructor NewLinkedWindowParent
			/* 1 */ imports.NewTable("TCEFLinkedWindowParent_Chromium", 0), // property Chromium
		}
	})
	return cEFLinkedWindowParentImport
}
