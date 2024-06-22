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

// ICEFSentinel Parent: IComponent
//
//	TCEFSentinel is used as a timer that checks the number of running
//	CEF processes when you close all browsers before shutdown.
//	This component is only used as a last resort when there's an unresolved
//	shutdown issue in CEF or CEF4Delphi that generates exceptions when the
//	application is closed.
type ICEFSentinel interface {
	IComponent
	// Status
	//  Status of this component.
	Status() TSentinelStatus // property
	// ChildProcCount
	//  Number of CEF subprocesses.
	ChildProcCount() int32 // property
	// DelayPerProcMs
	//  Delay per subprocess in milliseconds. This delay is used to calculate how much time to wait until this component checks the CEF subprocesses again.
	DelayPerProcMs() uint32 // property
	// SetDelayPerProcMs Set DelayPerProcMs
	SetDelayPerProcMs(AValue uint32) // property
	// MinInitDelayMs
	//  Minimum initial delay in milliseconds. This is the minimum time to wait until this component checks the CEF subprocesses again.
	MinInitDelayMs() uint32 // property
	// SetMinInitDelayMs Set MinInitDelayMs
	SetMinInitDelayMs(AValue uint32) // property
	// FinalDelayMs
	//  Final delay in milliseconds. This is an extra delay to wait after enough CEF subprocesses are closed.
	FinalDelayMs() uint32 // property
	// SetFinalDelayMs Set FinalDelayMs
	SetFinalDelayMs(AValue uint32) // property
	// MinChildProcs
	//  Minimum number of CEF subprocesses. When ChildProcCount reaches this value it's considered safe to trigger OnClose.
	MinChildProcs() int32 // property
	// SetMinChildProcs Set MinChildProcs
	SetMinChildProcs(AValue int32) // property
	// MaxCheckCount
	//  Maximum number of times this component will check the CEF subprocesses.
	MaxCheckCount() int32 // property
	// SetMaxCheckCount Set MaxCheckCount
	SetMaxCheckCount(AValue int32) // property
	// Start
	//  Start checking all the CEF subprocesses.
	Start() // procedure
	// SetOnClose
	//  Event triggered when enought CEF subprocesses are closed.
	SetOnClose(fn TNotify) // property event
}

// TCEFSentinel Parent: TComponent
//
//	TCEFSentinel is used as a timer that checks the number of running
//	CEF processes when you close all browsers before shutdown.
//	This component is only used as a last resort when there's an unresolved
//	shutdown issue in CEF or CEF4Delphi that generates exceptions when the
//	application is closed.
type TCEFSentinel struct {
	TComponent
	closePtr uintptr
}

func NewCEFSentinel(aOwner IComponent) ICEFSentinel {
	r1 := sentinelImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsCEFSentinel(r1)
}

func (m *TCEFSentinel) Status() TSentinelStatus {
	r1 := sentinelImportAPI().SysCallN(9, m.Instance())
	return TSentinelStatus(r1)
}

func (m *TCEFSentinel) ChildProcCount() int32 {
	r1 := sentinelImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

func (m *TCEFSentinel) DelayPerProcMs() uint32 {
	r1 := sentinelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFSentinel) SetDelayPerProcMs(AValue uint32) {
	sentinelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFSentinel) MinInitDelayMs() uint32 {
	r1 := sentinelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFSentinel) SetMinInitDelayMs(AValue uint32) {
	sentinelImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFSentinel) FinalDelayMs() uint32 {
	r1 := sentinelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFSentinel) SetFinalDelayMs(AValue uint32) {
	sentinelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFSentinel) MinChildProcs() int32 {
	r1 := sentinelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFSentinel) SetMinChildProcs(AValue int32) {
	sentinelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFSentinel) MaxCheckCount() int32 {
	r1 := sentinelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFSentinel) SetMaxCheckCount(AValue int32) {
	sentinelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFSentinel) Start() {
	sentinelImportAPI().SysCallN(8, m.Instance())
}

func (m *TCEFSentinel) SetOnClose(fn TNotify) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	sentinelImportAPI().SysCallN(7, m.Instance(), m.closePtr)
}

var (
	sentinelImport       *imports.Imports = nil
	sentinelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFSentinel_ChildProcCount", 0),
		/*1*/ imports.NewTable("CEFSentinel_Create", 0),
		/*2*/ imports.NewTable("CEFSentinel_DelayPerProcMs", 0),
		/*3*/ imports.NewTable("CEFSentinel_FinalDelayMs", 0),
		/*4*/ imports.NewTable("CEFSentinel_MaxCheckCount", 0),
		/*5*/ imports.NewTable("CEFSentinel_MinChildProcs", 0),
		/*6*/ imports.NewTable("CEFSentinel_MinInitDelayMs", 0),
		/*7*/ imports.NewTable("CEFSentinel_SetOnClose", 0),
		/*8*/ imports.NewTable("CEFSentinel_Start", 0),
		/*9*/ imports.NewTable("CEFSentinel_Status", 0),
	}
)

func sentinelImportAPI() *imports.Imports {
	if sentinelImport == nil {
		sentinelImport = NewDefaultImports()
		sentinelImport.SetImportTable(sentinelImportTables)
		sentinelImportTables = nil
	}
	return sentinelImport
}
