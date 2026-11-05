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

// ICustomDevToolsMessageObserver Parent: ICEFDevToolsMessageObserverOwn
type ICustomDevToolsMessageObserver interface {
	ICEFDevToolsMessageObserverOwn
	AsIntfDevToolsMessageObserver() uintptr
}

type TCustomDevToolsMessageObserver struct {
	TCEFDevToolsMessageObserverOwn
}

func (m *TCustomDevToolsMessageObserver) AsIntfDevToolsMessageObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDevToolsMessageObserver class constructor
func NewCustomDevToolsMessageObserver(events IChromiumCore) ICustomDevToolsMessageObserver {
	var devToolsMessageObserverPtr uintptr // ICefDevToolsMessageObserver
	r := customDevToolsMessageObserverAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&devToolsMessageObserverPtr)))
	ret := AsCustomDevToolsMessageObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, devToolsMessageObserverPtr)
	}
	return ret
}

var (
	customDevToolsMessageObserverOnce   base.Once
	customDevToolsMessageObserverImport *imports.Imports = nil
)

func customDevToolsMessageObserverAPI() *imports.Imports {
	customDevToolsMessageObserverOnce.Do(func() {
		customDevToolsMessageObserverImport = api.NewDefaultImports()
		customDevToolsMessageObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDevToolsMessageObserver_Create", 0), // constructor NewCustomDevToolsMessageObserver
		}
	})
	return customDevToolsMessageObserverImport
}
