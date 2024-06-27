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

// IEnergyStringMap Parent: ICefStringMap
//
//	CEF string maps are a set of key/value string pairs.
type IEnergyStringMap interface {
	ICefStringMap
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringMap // procedure
}

// TEnergyStringMap Parent: TCefStringMap
//
//	CEF string maps are a set of key/value string pairs.
type TEnergyStringMap struct {
	TCefStringMap
}

func NewEnergyStringMap() IEnergyStringMap {
	r1 := energyStringMapImportAPI().SysCallN(1)
	return AsEnergyStringMap(r1)
}

func (m *TEnergyStringMap) AsInterface() ICefStringMap {
	var resultCefStringMap uintptr
	energyStringMapImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringMap)))
	return AsCefStringMap(resultCefStringMap)
}

var (
	energyStringMapImport       *imports.Imports = nil
	energyStringMapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("EnergyStringMap_AsInterface", 0),
		/*1*/ imports.NewTable("EnergyStringMap_Create", 0),
	}
)

func energyStringMapImportAPI() *imports.Imports {
	if energyStringMapImport == nil {
		energyStringMapImport = NewDefaultImports()
		energyStringMapImport.SetImportTable(energyStringMapImportTables)
		energyStringMapImportTables = nil
	}
	return energyStringMapImport
}
