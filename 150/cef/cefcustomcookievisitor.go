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

// ICefCustomCookieVisitor Parent: ICefCookieVisitorOwn
type ICefCustomCookieVisitor interface {
	ICefCookieVisitorOwn
	AsIntfCookieVisitor() uintptr
}

type TCefCustomCookieVisitor struct {
	TCefCookieVisitorOwn
}

func (m *TCefCustomCookieVisitor) AsIntfCookieVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCookieVisitor class constructor
func NewCustomCookieVisitor(events IChromiumCore, iD int32) ICefCustomCookieVisitor {
	var cookieVisitorPtr uintptr // ICefCookieVisitor
	r := cefCustomCookieVisitorAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(iD), uintptr(base.UnsafePointer(&cookieVisitorPtr)))
	ret := AsCefCustomCookieVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieVisitorPtr)
	}
	return ret
}

var (
	cefCustomCookieVisitorOnce   base.Once
	cefCustomCookieVisitorImport *imports.Imports = nil
)

func cefCustomCookieVisitorAPI() *imports.Imports {
	cefCustomCookieVisitorOnce.Do(func() {
		cefCustomCookieVisitorImport = api.NewDefaultImports()
		cefCustomCookieVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomCookieVisitor_Create", 0), // constructor NewCustomCookieVisitor
		}
	})
	return cefCustomCookieVisitorImport
}
