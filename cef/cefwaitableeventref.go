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

// ICefWaitableEvent Parent: ICefBaseRefCounted
type ICefWaitableEvent interface {
	ICefBaseRefCounted
	// IsSignaled
	//  Returns true (1) if the event is in the signaled state, else false (0). If
	//  the event was created with |automatic_reset| set to true (1) then calling
	//  this function will also cause a reset.
	IsSignaled() bool // function
	// TimedWait
	//  Wait up to |max_ms| milliseconds for the event to be signaled. Returns
	//  true (1) if the event was signaled. A return value of false (0) does not
	//  necessarily mean that |max_ms| was exceeded. This function will not return
	//  until after the call to signal() has completed. This function cannot be
	//  called on the browser process UI or IO threads.
	TimedWait(maxMs int64) bool // function
	// Reset
	//  Put the event in the un-signaled state.
	Reset() // procedure
	// Signal
	//  Put the event in the signaled state. This causes any thread blocked on
	//  Wait to be woken up.
	Signal() // procedure
	// Wait
	//  Wait indefinitely for the event to be signaled. This function will not
	//  return until after the call to signal() has completed. This function
	//  cannot be called on the browser process UI or IO threads.
	Wait() // procedure
}

// ICefWaitableEventRef Parent: ICefWaitableEvent ICefBaseRefCountedRef
type ICefWaitableEventRef interface {
	ICefWaitableEvent
	ICefBaseRefCountedRef
	AsIntfWaitableEvent() uintptr
}

type TCefWaitableEventRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefWaitableEventRef) IsSignaled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWaitableEventRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWaitableEventRef) TimedWait(maxMs int64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefWaitableEventRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&maxMs)))
	return api.GoBool(r)
}

func (m *TCefWaitableEventRef) Reset() {
	if !m.IsValid() {
		return
	}
	cefWaitableEventRefAPI().SysCallN(5, m.Instance())
}

func (m *TCefWaitableEventRef) Signal() {
	if !m.IsValid() {
		return
	}
	cefWaitableEventRefAPI().SysCallN(6, m.Instance())
}

func (m *TCefWaitableEventRef) Wait() {
	if !m.IsValid() {
		return
	}
	cefWaitableEventRefAPI().SysCallN(7, m.Instance())
}

func (m *TCefWaitableEventRef) AsIntfWaitableEvent() uintptr {
	return m.GetIntfPointer(0)
}

// WaitableEventRef  is static instance
var WaitableEventRef _WaitableEventRefClass

// _WaitableEventRefClass is class type defined by TCefWaitableEventRef
type _WaitableEventRefClass uintptr

func (_WaitableEventRefClass) UnWrap(data uintptr) (result ICefWaitableEvent) {
	var resultPtr uintptr
	cefWaitableEventRefAPI().SysCallN(3, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWaitableEventRef(resultPtr)
	return
}

func (_WaitableEventRefClass) New(automaticReset bool, initiallySignaled bool) (result ICefWaitableEvent) {
	var resultPtr uintptr
	cefWaitableEventRefAPI().SysCallN(4, api.PasBool(automaticReset), api.PasBool(initiallySignaled), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWaitableEventRef(resultPtr)
	return
}

// NewWaitableEventRef class constructor
func NewWaitableEventRef(data uintptr) ICefWaitableEventRef {
	var waitableEventPtr uintptr // ICefWaitableEvent
	r := cefWaitableEventRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&waitableEventPtr)))
	ret := AsCefWaitableEventRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, waitableEventPtr)
	}
	return ret
}

var (
	cefWaitableEventRefOnce   base.Once
	cefWaitableEventRefImport *imports.Imports = nil
)

func cefWaitableEventRefAPI() *imports.Imports {
	cefWaitableEventRefOnce.Do(func() {
		cefWaitableEventRefImport = api.NewDefaultImports()
		cefWaitableEventRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefWaitableEventRef_Create", 0), // constructor NewWaitableEventRef
			/* 1 */ imports.NewTable("TCefWaitableEventRef_IsSignaled", 0), // function IsSignaled
			/* 2 */ imports.NewTable("TCefWaitableEventRef_TimedWait", 0), // function TimedWait
			/* 3 */ imports.NewTable("TCefWaitableEventRef_UnWrap", 0), // static function UnWrap
			/* 4 */ imports.NewTable("TCefWaitableEventRef_New", 0), // static function New
			/* 5 */ imports.NewTable("TCefWaitableEventRef_Reset", 0), // procedure Reset
			/* 6 */ imports.NewTable("TCefWaitableEventRef_Signal", 0), // procedure Signal
			/* 7 */ imports.NewTable("TCefWaitableEventRef_Wait", 0), // procedure Wait
		}
	})
	return cefWaitableEventRefImport
}
