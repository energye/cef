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

	cefTypes "github.com/energye/cef/types"
)

// ICEFSentinel Parent: IComponent
type ICEFSentinel interface {
	IComponent
	// Start
	//  Start checking all the CEF subprocesses.
	Start() // procedure
	// Status
	//  Status of this component.
	Status() cefTypes.TSentinelStatus // property Status Getter
	// ChildProcCount
	//  Number of CEF subprocesses.
	ChildProcCount() int32 // property ChildProcCount Getter
	// DelayPerProcMs
	//  Delay per subprocess in milliseconds. This delay is used to calculate how much time to wait until this component checks the CEF subprocesses again.
	DelayPerProcMs() uint32         // property DelayPerProcMs Getter
	SetDelayPerProcMs(value uint32) // property DelayPerProcMs Setter
	// MinInitDelayMs
	//  Minimum initial delay in milliseconds. This is the minimum time to wait until this component checks the CEF subprocesses again.
	MinInitDelayMs() uint32         // property MinInitDelayMs Getter
	SetMinInitDelayMs(value uint32) // property MinInitDelayMs Setter
	// FinalDelayMs
	//  Final delay in milliseconds. This is an extra delay to wait after enough CEF subprocesses are closed.
	FinalDelayMs() uint32         // property FinalDelayMs Getter
	SetFinalDelayMs(value uint32) // property FinalDelayMs Setter
	// MinChildProcs
	//  Minimum number of CEF subprocesses. When ChildProcCount reaches this value it's considered safe to trigger OnClose.
	MinChildProcs() int32         // property MinChildProcs Getter
	SetMinChildProcs(value int32) // property MinChildProcs Setter
	// MaxCheckCount
	//  Maximum number of times this component will check the CEF subprocesses.
	MaxCheckCount() int32         // property MaxCheckCount Getter
	SetMaxCheckCount(value int32) // property MaxCheckCount Setter
	SetOnClose(fn TNotifyEvent)   // property event
}

type TCEFSentinel struct {
	TComponent
}

func (m *TCEFSentinel) Start() {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(1, m.Instance())
}

func (m *TCEFSentinel) Status() cefTypes.TSentinelStatus {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(2, m.Instance())
	return cefTypes.TSentinelStatus(r)
}

func (m *TCEFSentinel) ChildProcCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TCEFSentinel) DelayPerProcMs() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(4, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFSentinel) SetDelayPerProcMs(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCEFSentinel) MinInitDelayMs() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(5, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFSentinel) SetMinInitDelayMs(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCEFSentinel) FinalDelayMs() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(6, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFSentinel) SetFinalDelayMs(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCEFSentinel) MinChildProcs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCEFSentinel) SetMinChildProcs(value int32) {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCEFSentinel) MaxCheckCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFSentinelAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCEFSentinel) SetMaxCheckCount(value int32) {
	if !m.IsValid() {
		return
	}
	cEFSentinelAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCEFSentinel) SetOnClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, cEFSentinelAPI(), api.MakeEventDataPtr(cb))
}

// NewSentinel class constructor
func NewSentinel(owner lcl.IComponent) ICEFSentinel {
	r := cEFSentinelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCEFSentinel(r)
}

var (
	cEFSentinelOnce   base.Once
	cEFSentinelImport *imports.Imports = nil
)

func cEFSentinelAPI() *imports.Imports {
	cEFSentinelOnce.Do(func() {
		cEFSentinelImport = api.NewDefaultImports()
		cEFSentinelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFSentinel_Create", 0), // constructor NewSentinel
			/* 1 */ imports.NewTable("TCEFSentinel_Start", 0), // procedure Start
			/* 2 */ imports.NewTable("TCEFSentinel_Status", 0), // property Status
			/* 3 */ imports.NewTable("TCEFSentinel_ChildProcCount", 0), // property ChildProcCount
			/* 4 */ imports.NewTable("TCEFSentinel_DelayPerProcMs", 0), // property DelayPerProcMs
			/* 5 */ imports.NewTable("TCEFSentinel_MinInitDelayMs", 0), // property MinInitDelayMs
			/* 6 */ imports.NewTable("TCEFSentinel_FinalDelayMs", 0), // property FinalDelayMs
			/* 7 */ imports.NewTable("TCEFSentinel_MinChildProcs", 0), // property MinChildProcs
			/* 8 */ imports.NewTable("TCEFSentinel_MaxCheckCount", 0), // property MaxCheckCount
			/* 9 */ imports.NewTable("TCEFSentinel_OnClose", 0), // event OnClose
		}
	})
	return cEFSentinelImport
}
