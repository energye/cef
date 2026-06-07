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
	"github.com/energye/lcl/types"
)

// ICEFWorkScheduler Parent: IComponent
type ICEFWorkScheduler interface {
	IComponent
	ScheduleMessagePumpWork(delayMs int64)   // procedure
	StopScheduler()                          // procedure
	CreateThread()                           // procedure
	Priority() types.TThreadPriority         // property Priority Getter
	SetPriority(value types.TThreadPriority) // property Priority Setter
	DefaultInterval() int32                  // property DefaultInterval Getter
	SetDefaultInterval(value int32)          // property DefaultInterval Setter
	DepleteWorkCycles() uint32               // property DepleteWorkCycles Getter
	SetDepleteWorkCycles(value uint32)       // property DepleteWorkCycles Setter
	DepleteWorkDelay() uint32                // property DepleteWorkDelay Getter
	SetDepleteWorkDelay(value uint32)        // property DepleteWorkDelay Setter
	UseQueueThread() bool                    // property UseQueueThread Getter
	SetUseQueueThread(value bool)            // property UseQueueThread Setter
}

type TCEFWorkScheduler struct {
	TComponent
}

func (m *TCEFWorkScheduler) ScheduleMessagePumpWork(delayMs int64) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&delayMs)))
}

func (m *TCEFWorkScheduler) StopScheduler() {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(3, m.Instance())
}

func (m *TCEFWorkScheduler) CreateThread() {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(4, m.Instance())
}

func (m *TCEFWorkScheduler) Priority() types.TThreadPriority {
	if !m.IsValid() {
		return 0
	}
	r := cEFWorkSchedulerAPI().SysCallN(5, 0, m.Instance())
	return types.TThreadPriority(r)
}

func (m *TCEFWorkScheduler) SetPriority(value types.TThreadPriority) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWorkScheduler) DefaultInterval() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFWorkSchedulerAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCEFWorkScheduler) SetDefaultInterval(value int32) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWorkScheduler) DepleteWorkCycles() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFWorkSchedulerAPI().SysCallN(7, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFWorkScheduler) SetDepleteWorkCycles(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWorkScheduler) DepleteWorkDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFWorkSchedulerAPI().SysCallN(8, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFWorkScheduler) SetDepleteWorkDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWorkScheduler) UseQueueThread() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWorkSchedulerAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWorkScheduler) SetUseQueueThread(value bool) {
	if !m.IsValid() {
		return
	}
	cEFWorkSchedulerAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

// NewWorkScheduler class constructor
func NewWorkScheduler(owner lcl.IComponent) ICEFWorkScheduler {
	r := cEFWorkSchedulerAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCEFWorkScheduler(r)
}

// NewWorkSchedulerDelayed class constructor
func NewWorkSchedulerDelayed() ICEFWorkScheduler {
	r := cEFWorkSchedulerAPI().SysCallN(1)
	return AsCEFWorkScheduler(r)
}

var (
	cEFWorkSchedulerOnce   base.Once
	cEFWorkSchedulerImport *imports.Imports = nil
)

func cEFWorkSchedulerAPI() *imports.Imports {
	cEFWorkSchedulerOnce.Do(func() {
		cEFWorkSchedulerImport = api.NewDefaultImports()
		cEFWorkSchedulerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFWorkScheduler_Create", 0), // constructor NewWorkScheduler
			/* 1 */ imports.NewTable("TCEFWorkScheduler_CreateDelayed", 0), // constructor NewWorkSchedulerDelayed
			/* 2 */ imports.NewTable("TCEFWorkScheduler_ScheduleMessagePumpWork", 0), // procedure ScheduleMessagePumpWork
			/* 3 */ imports.NewTable("TCEFWorkScheduler_StopScheduler", 0), // procedure StopScheduler
			/* 4 */ imports.NewTable("TCEFWorkScheduler_CreateThread", 0), // procedure CreateThread
			/* 5 */ imports.NewTable("TCEFWorkScheduler_Priority", 0), // property Priority
			/* 6 */ imports.NewTable("TCEFWorkScheduler_DefaultInterval", 0), // property DefaultInterval
			/* 7 */ imports.NewTable("TCEFWorkScheduler_DepleteWorkCycles", 0), // property DepleteWorkCycles
			/* 8 */ imports.NewTable("TCEFWorkScheduler_DepleteWorkDelay", 0), // property DepleteWorkDelay
			/* 9 */ imports.NewTable("TCEFWorkScheduler_UseQueueThread", 0), // property UseQueueThread
		}
	})
	return cEFWorkSchedulerImport
}
