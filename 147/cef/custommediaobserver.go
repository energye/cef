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

// ICustomMediaObserver Parent: ICefMediaObserverOwn
type ICustomMediaObserver interface {
	ICefMediaObserverOwn
	AsIntfMediaObserver() uintptr
}

type TCustomMediaObserver struct {
	TCefMediaObserverOwn
}

func (m *TCustomMediaObserver) AsIntfMediaObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomMediaObserver class constructor
func NewCustomMediaObserver(events IChromiumCore) ICustomMediaObserver {
	var mediaObserverPtr uintptr // ICefMediaObserver
	r := customMediaObserverAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&mediaObserverPtr)))
	ret := AsCustomMediaObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaObserverPtr)
	}
	return ret
}

var (
	customMediaObserverOnce   base.Once
	customMediaObserverImport *imports.Imports = nil
)

func customMediaObserverAPI() *imports.Imports {
	customMediaObserverOnce.Do(func() {
		customMediaObserverImport = api.NewDefaultImports()
		customMediaObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomMediaObserver_Create", 0), // constructor NewCustomMediaObserver
		}
	})
	return customMediaObserverImport
}
