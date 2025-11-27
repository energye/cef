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

// ICEFTimerWorkScheduler Parent: lcl.IObject
type ICEFTimerWorkScheduler interface {
	lcl.IObject
	StopScheduler()                        // procedure
	ScheduleMessagePumpWork(delayMs int64) // procedure
	DepleteWorkCycles() uint32             // property DepleteWorkCycles Getter
	SetDepleteWorkCycles(value uint32)     // property DepleteWorkCycles Setter
	DepleteWorkDelay() uint32              // property DepleteWorkDelay Getter
	SetDepleteWorkDelay(value uint32)      // property DepleteWorkDelay Setter
	IsTimerPending() bool                  // property IsTimerPending Getter
	SetOnAllowDoWork(fn TOnAllowEvent)     // property event
}

type TCEFTimerWorkScheduler struct {
	lcl.TObject
}

func (m *TCEFTimerWorkScheduler) StopScheduler() {
	if !m.IsValid() {
		return
	}
	cEFTimerWorkSchedulerAPI().SysCallN(1, m.Instance())
}

func (m *TCEFTimerWorkScheduler) ScheduleMessagePumpWork(delayMs int64) {
	if !m.IsValid() {
		return
	}
	cEFTimerWorkSchedulerAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&delayMs)))
}

func (m *TCEFTimerWorkScheduler) DepleteWorkCycles() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFTimerWorkSchedulerAPI().SysCallN(3, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFTimerWorkScheduler) SetDepleteWorkCycles(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFTimerWorkSchedulerAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCEFTimerWorkScheduler) DepleteWorkDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFTimerWorkSchedulerAPI().SysCallN(4, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFTimerWorkScheduler) SetDepleteWorkDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFTimerWorkSchedulerAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCEFTimerWorkScheduler) IsTimerPending() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTimerWorkSchedulerAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTimerWorkScheduler) SetOnAllowDoWork(fn TOnAllowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAllowEvent(fn)
	base.SetEvent(m, 6, cEFTimerWorkSchedulerAPI(), api.MakeEventDataPtr(cb))
}

// NewTimerWorkScheduler class constructor
func NewTimerWorkScheduler() ICEFTimerWorkScheduler {
	r := cEFTimerWorkSchedulerAPI().SysCallN(0)
	return AsCEFTimerWorkScheduler(r)
}

var (
	cEFTimerWorkSchedulerOnce   base.Once
	cEFTimerWorkSchedulerImport *imports.Imports = nil
)

func cEFTimerWorkSchedulerAPI() *imports.Imports {
	cEFTimerWorkSchedulerOnce.Do(func() {
		cEFTimerWorkSchedulerImport = api.NewDefaultImports()
		cEFTimerWorkSchedulerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFTimerWorkScheduler_Create", 0), // constructor NewTimerWorkScheduler
			/* 1 */ imports.NewTable("TCEFTimerWorkScheduler_StopScheduler", 0), // procedure StopScheduler
			/* 2 */ imports.NewTable("TCEFTimerWorkScheduler_ScheduleMessagePumpWork", 0), // procedure ScheduleMessagePumpWork
			/* 3 */ imports.NewTable("TCEFTimerWorkScheduler_DepleteWorkCycles", 0), // property DepleteWorkCycles
			/* 4 */ imports.NewTable("TCEFTimerWorkScheduler_DepleteWorkDelay", 0), // property DepleteWorkDelay
			/* 5 */ imports.NewTable("TCEFTimerWorkScheduler_IsTimerPending", 0), // property IsTimerPending
			/* 6 */ imports.NewTable("TCEFTimerWorkScheduler_OnAllowDoWork", 0), // event OnAllowDoWork
		}
	})
	return cEFTimerWorkSchedulerImport
}
