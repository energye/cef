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

// IEngSettingObserver Parent: ICefSettingObserverOwn
type IEngSettingObserver interface {
	ICefSettingObserverOwn
	SetOnSettingObserverSettingChanged(fn TOnSettingObserverSettingChangedEvent) // property event
	AsIntfSettingObserver() uintptr
}

type TEngSettingObserver struct {
	TCefSettingObserverOwn
}

func (m *TEngSettingObserver) SetOnSettingObserverSettingChanged(fn TOnSettingObserverSettingChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSettingObserverSettingChangedEvent(fn)
	base.SetEvent(m, 1, engSettingObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngSettingObserver) AsIntfSettingObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngSettingObserver class constructor
func NewEngSettingObserver() IEngSettingObserver {
	var settingObserverPtr uintptr // ICefSettingObserver
	r := engSettingObserverAPI().SysCallN(0, uintptr(base.UnsafePointer(&settingObserverPtr)))
	ret := AsEngSettingObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, settingObserverPtr)
	}
	return ret
}

var (
	engSettingObserverOnce   base.Once
	engSettingObserverImport *imports.Imports = nil
)

func engSettingObserverAPI() *imports.Imports {
	engSettingObserverOnce.Do(func() {
		engSettingObserverImport = api.NewDefaultImports()
		engSettingObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngSettingObserver_Create", 0), // constructor NewEngSettingObserver
			/* 1 */ imports.NewTable("TEngSettingObserver_OnSettingObserverSettingChanged", 0), // event OnSettingObserverSettingChanged
		}
	})
	return engSettingObserverImport
}
