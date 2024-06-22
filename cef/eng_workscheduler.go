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

// ICEFWorkScheduler Parent: IComponent
//
//	Implementation of an external message pump for VCL and LCL.
//	Read the GlobalCEFApp.OnScheduleMessagePumpWork documentation for all the details.
type ICEFWorkScheduler interface {
	IComponent
	// Priority
	//  Priority of TCEFWorkSchedulerThread in Windows.
	Priority() TThreadPriority // property
	// SetPriority Set Priority
	SetPriority(AValue TThreadPriority) // property
	// DefaultInterval
	//  Default interval in milliseconds to do the next GlobalCEFApp.DoMessageLoopWork call.
	DefaultInterval() int32 // property
	// SetDefaultInterval Set DefaultInterval
	SetDefaultInterval(AValue int32) // property
	// DepleteWorkCycles
	//  Number of cycles used to deplete the remaining messages in the work loop.
	DepleteWorkCycles() uint32 // property
	// SetDepleteWorkCycles Set DepleteWorkCycles
	SetDepleteWorkCycles(AValue uint32) // property
	// DepleteWorkDelay
	//  Delay in milliseconds between the cycles used to deplete the remaining messages in the work loop.
	DepleteWorkDelay() uint32 // property
	// SetDepleteWorkDelay Set DepleteWorkDelay
	SetDepleteWorkDelay(AValue uint32) // property
	// UseQueueThread
	//  Use a custom queue thread instead of Windows messages or any other way to schedule the next pump work.
	UseQueueThread() bool // property
	// SetUseQueueThread Set UseQueueThread
	SetUseQueueThread(AValue bool) // property
	// ScheduleMessagePumpWork
	//  TCEFWorkScheduler destructor.
	//  Called from GlobalCEFApp.OnScheduleMessagePumpWork to schedule
	//  a GlobalCEFApp.DoMessageLoopWork call asynchronously to perform a single
	//  iteration of CEF message loop processing.
	//  <param name="delay_ms">Requested delay in milliseconds.</param>
	ScheduleMessagePumpWork(delayms int64) // procedure
	// StopScheduler
	//  Stop the scheduler. This function must be called after the destruction of all the forms in the application.
	StopScheduler() // procedure
	// CreateThread
	//  Creates all the internal threads used by TCEFWorkScheduler.
	CreateThread() // procedure
}

// TCEFWorkScheduler Parent: TComponent
//
//	Implementation of an external message pump for VCL and LCL.
//	Read the GlobalCEFApp.OnScheduleMessagePumpWork documentation for all the details.
type TCEFWorkScheduler struct {
	TComponent
}

func NewCEFWorkScheduler(aOwner IComponent) ICEFWorkScheduler {
	r1 := workSchedulerImportAPI().SysCallN(0, GetObjectUintptr(aOwner))
	return AsCEFWorkScheduler(r1)
}

func NewCEFWorkSchedulerDelayed() ICEFWorkScheduler {
	r1 := workSchedulerImportAPI().SysCallN(1)
	return AsCEFWorkScheduler(r1)
}

func (m *TCEFWorkScheduler) Priority() TThreadPriority {
	r1 := workSchedulerImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TThreadPriority(r1)
}

func (m *TCEFWorkScheduler) SetPriority(AValue TThreadPriority) {
	workSchedulerImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DefaultInterval() int32 {
	r1 := workSchedulerImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFWorkScheduler) SetDefaultInterval(AValue int32) {
	workSchedulerImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DepleteWorkCycles() uint32 {
	r1 := workSchedulerImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFWorkScheduler) SetDepleteWorkCycles(AValue uint32) {
	workSchedulerImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DepleteWorkDelay() uint32 {
	r1 := workSchedulerImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFWorkScheduler) SetDepleteWorkDelay(AValue uint32) {
	workSchedulerImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) UseQueueThread() bool {
	r1 := workSchedulerImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFWorkScheduler) SetUseQueueThread(AValue bool) {
	workSchedulerImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFWorkScheduler) ScheduleMessagePumpWork(delayms int64) {
	workSchedulerImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&delayms)))
}

func (m *TCEFWorkScheduler) StopScheduler() {
	workSchedulerImportAPI().SysCallN(8, m.Instance())
}

func (m *TCEFWorkScheduler) CreateThread() {
	workSchedulerImportAPI().SysCallN(2, m.Instance())
}

var (
	workSchedulerImport       *imports.Imports = nil
	workSchedulerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFWorkScheduler_Create", 0),
		/*1*/ imports.NewTable("CEFWorkScheduler_CreateDelayed", 0),
		/*2*/ imports.NewTable("CEFWorkScheduler_CreateThread", 0),
		/*3*/ imports.NewTable("CEFWorkScheduler_DefaultInterval", 0),
		/*4*/ imports.NewTable("CEFWorkScheduler_DepleteWorkCycles", 0),
		/*5*/ imports.NewTable("CEFWorkScheduler_DepleteWorkDelay", 0),
		/*6*/ imports.NewTable("CEFWorkScheduler_Priority", 0),
		/*7*/ imports.NewTable("CEFWorkScheduler_ScheduleMessagePumpWork", 0),
		/*8*/ imports.NewTable("CEFWorkScheduler_StopScheduler", 0),
		/*9*/ imports.NewTable("CEFWorkScheduler_UseQueueThread", 0),
	}
)

func workSchedulerImportAPI() *imports.Imports {
	if workSchedulerImport == nil {
		workSchedulerImport = NewDefaultImports()
		workSchedulerImport.SetImportTable(workSchedulerImportTables)
		workSchedulerImportTables = nil
	}
	return workSchedulerImport
}
