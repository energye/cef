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

// IEngViewDelegate Parent: ICefViewDelegateOwn
type IEngViewDelegate interface {
	ICefViewDelegateOwn
	SetOnViewGetPreferredSize(fn TOnViewGetPreferredSizeEvent)   // property event
	SetOnViewGetMinimumSize(fn TOnViewGetMinimumSizeEvent)       // property event
	SetOnViewGetMaximumSize(fn TOnViewGetMaximumSizeEvent)       // property event
	SetOnViewGetHeightForWidth(fn TOnViewGetHeightForWidthEvent) // property event
	SetOnViewParentViewChanged(fn TOnViewParentViewChangedEvent) // property event
	SetOnViewChildViewChanged(fn TOnViewChildViewChangedEvent)   // property event
	SetOnViewWindowChanged(fn TOnViewWindowChangedEvent)         // property event
	SetOnViewLayoutChanged(fn TOnViewLayoutChangedEvent)         // property event
	SetOnViewFocus(fn TOnViewFocusEvent)                         // property event
	SetOnViewBlur(fn TOnViewBlurEvent)                           // property event
	SetOnViewThemeChanged(fn TOnViewThemeChangedEvent)           // property event
	AsIntfViewDelegate() uintptr
}

type TEngViewDelegate struct {
	TCefViewDelegateOwn
}

func (m *TEngViewDelegate) SetOnViewGetPreferredSize(fn TOnViewGetPreferredSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewGetPreferredSizeEvent(fn)
	base.SetEvent(m, 1, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewGetMinimumSize(fn TOnViewGetMinimumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewGetMinimumSizeEvent(fn)
	base.SetEvent(m, 2, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewGetMaximumSize(fn TOnViewGetMaximumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewGetMaximumSizeEvent(fn)
	base.SetEvent(m, 3, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewGetHeightForWidth(fn TOnViewGetHeightForWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewGetHeightForWidthEvent(fn)
	base.SetEvent(m, 4, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewParentViewChanged(fn TOnViewParentViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewParentViewChangedEvent(fn)
	base.SetEvent(m, 5, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewChildViewChanged(fn TOnViewChildViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewChildViewChangedEvent(fn)
	base.SetEvent(m, 6, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewWindowChanged(fn TOnViewWindowChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewWindowChangedEvent(fn)
	base.SetEvent(m, 7, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewLayoutChanged(fn TOnViewLayoutChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewLayoutChangedEvent(fn)
	base.SetEvent(m, 8, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewFocus(fn TOnViewFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewFocusEvent(fn)
	base.SetEvent(m, 9, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewBlur(fn TOnViewBlurEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewBlurEvent(fn)
	base.SetEvent(m, 10, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) SetOnViewThemeChanged(fn TOnViewThemeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnViewThemeChangedEvent(fn)
	base.SetEvent(m, 11, engViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngViewDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngViewDelegate class constructor
func NewEngViewDelegate() IEngViewDelegate {
	var viewDelegatePtr uintptr // ICefViewDelegate
	r := engViewDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngViewDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewDelegatePtr)
	}
	return ret
}

var (
	engViewDelegateOnce   base.Once
	engViewDelegateImport *imports.Imports = nil
)

func engViewDelegateAPI() *imports.Imports {
	engViewDelegateOnce.Do(func() {
		engViewDelegateImport = api.NewDefaultImports()
		engViewDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngViewDelegate_Create", 0), // constructor NewEngViewDelegate
			/* 1 */ imports.NewTable("TEngViewDelegate_OnViewGetPreferredSize", 0), // event OnViewGetPreferredSize
			/* 2 */ imports.NewTable("TEngViewDelegate_OnViewGetMinimumSize", 0), // event OnViewGetMinimumSize
			/* 3 */ imports.NewTable("TEngViewDelegate_OnViewGetMaximumSize", 0), // event OnViewGetMaximumSize
			/* 4 */ imports.NewTable("TEngViewDelegate_OnViewGetHeightForWidth", 0), // event OnViewGetHeightForWidth
			/* 5 */ imports.NewTable("TEngViewDelegate_OnViewParentViewChanged", 0), // event OnViewParentViewChanged
			/* 6 */ imports.NewTable("TEngViewDelegate_OnViewChildViewChanged", 0), // event OnViewChildViewChanged
			/* 7 */ imports.NewTable("TEngViewDelegate_OnViewWindowChanged", 0), // event OnViewWindowChanged
			/* 8 */ imports.NewTable("TEngViewDelegate_OnViewLayoutChanged", 0), // event OnViewLayoutChanged
			/* 9 */ imports.NewTable("TEngViewDelegate_OnViewFocus", 0), // event OnViewFocus
			/* 10 */ imports.NewTable("TEngViewDelegate_OnViewBlur", 0), // event OnViewBlur
			/* 11 */ imports.NewTable("TEngViewDelegate_OnViewThemeChanged", 0), // event OnViewThemeChanged
		}
	})
	return engViewDelegateImport
}
