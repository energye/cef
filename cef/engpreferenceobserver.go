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

// IEngPreferenceObserver Parent: ICefPreferenceObserverOwn
type IEngPreferenceObserver interface {
	ICefPreferenceObserverOwn
	SetOnPreferenceObserverPreferenceChanged(fn TOnPreferenceObserverPreferenceChangedEvent) // property event
	AsIntfPreferenceObserver() uintptr
}

type TEngPreferenceObserver struct {
	TCefPreferenceObserverOwn
}

func (m *TEngPreferenceObserver) SetOnPreferenceObserverPreferenceChanged(fn TOnPreferenceObserverPreferenceChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPreferenceObserverPreferenceChangedEvent(fn)
	base.SetEvent(m, 1, engPreferenceObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPreferenceObserver) AsIntfPreferenceObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngPreferenceObserver class constructor
func NewEngPreferenceObserver() IEngPreferenceObserver {
	var preferenceObserverPtr uintptr // ICefPreferenceObserver
	r := engPreferenceObserverAPI().SysCallN(0, uintptr(base.UnsafePointer(&preferenceObserverPtr)))
	ret := AsEngPreferenceObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, preferenceObserverPtr)
	}
	return ret
}

var (
	engPreferenceObserverOnce   base.Once
	engPreferenceObserverImport *imports.Imports = nil
)

func engPreferenceObserverAPI() *imports.Imports {
	engPreferenceObserverOnce.Do(func() {
		engPreferenceObserverImport = api.NewDefaultImports()
		engPreferenceObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngPreferenceObserver_Create", 0), // constructor NewEngPreferenceObserver
			/* 1 */ imports.NewTable("TEngPreferenceObserver_OnPreferenceObserverPreferenceChanged", 0), // event OnPreferenceObserverPreferenceChanged
		}
	})
	return engPreferenceObserverImport
}
