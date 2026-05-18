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

// ICEFFileDialogInfo Parent: IObject
type ICEFFileDialogInfo interface {
	IObject
	Clear()                                   // procedure
	Mode() uint32                             // property Mode Getter
	SetMode(value uint32)                     // property Mode Setter
	Title() string                            // property Title Getter
	SetTitle(value string)                    // property Title Setter
	DefaultFilePath() string                  // property DefaultFilePath Getter
	SetDefaultFilePath(value string)          // property DefaultFilePath Setter
	SetAcceptFilters(value lcl.IStrings)      // property AcceptFilters Setter
	SetAcceptExtensions(value lcl.IStrings)   // property AcceptExtensions Setter
	SetAcceptDescriptions(value lcl.IStrings) // property AcceptDescriptions Setter
	Callback() ICefFileDialogCallback         // property Callback Getter
	SetCallback(value ICefFileDialogCallback) // property Callback Setter
	DialogFilter() string                     // property DialogFilter Getter
	DialogType() cefTypes.TCEFDialogType      // property DialogType Getter
	DefaultAudioFileDesc() string             // property DefaultAudioFileDesc Getter
	SetDefaultAudioFileDesc(value string)     // property DefaultAudioFileDesc Setter
	DefaultVideoFileDesc() string             // property DefaultVideoFileDesc Getter
	SetDefaultVideoFileDesc(value string)     // property DefaultVideoFileDesc Setter
	DefaultTextFileDesc() string              // property DefaultTextFileDesc Getter
	SetDefaultTextFileDesc(value string)      // property DefaultTextFileDesc Setter
	DefaultImageFileDesc() string             // property DefaultImageFileDesc Getter
	SetDefaultImageFileDesc(value string)     // property DefaultImageFileDesc Setter
	DefaultAllFileDesc() string               // property DefaultAllFileDesc Getter
	SetDefaultAllFileDesc(value string)       // property DefaultAllFileDesc Setter
	DefaultUnknownFileDesc() string           // property DefaultUnknownFileDesc Getter
	SetDefaultUnknownFileDesc(value string)   // property DefaultUnknownFileDesc Setter
}

type TCEFFileDialogInfo struct {
	TObject
}

func (m *TCEFFileDialogInfo) Clear() {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(1, m.Instance())
}

func (m *TCEFFileDialogInfo) Mode() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFFileDialogInfoAPI().SysCallN(2, 0, m.Instance())
	return uint32(r)
}

func (m *TCEFFileDialogInfo) SetMode(value uint32) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCEFFileDialogInfo) Title() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultFilePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultFilePath(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) SetAcceptFilters(value lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFFileDialogInfo) SetAcceptExtensions(value lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFFileDialogInfo) SetAcceptDescriptions(value lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFFileDialogInfo) Callback() (result ICefFileDialogCallback) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFFileDialogInfoAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFileDialogCallbackRef(resultPtr)
	return
}

func (m *TCEFFileDialogInfo) SetCallback(value ICefFileDialogCallback) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFFileDialogInfo) DialogFilter() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) DialogType() cefTypes.TCEFDialogType {
	if !m.IsValid() {
		return 0
	}
	r := cEFFileDialogInfoAPI().SysCallN(10, m.Instance())
	return cefTypes.TCEFDialogType(r)
}

func (m *TCEFFileDialogInfo) DefaultAudioFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultAudioFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(11, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultVideoFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultVideoFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultTextFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultTextFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultImageFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultImageFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(14, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultAllFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(15, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultAllFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFFileDialogInfo) DefaultUnknownFileDesc() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFFileDialogInfoAPI().SysCallN(16, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFFileDialogInfo) SetDefaultUnknownFileDesc(value string) {
	if !m.IsValid() {
		return
	}
	cEFFileDialogInfoAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

// NewFileDialogInfo class constructor
func NewFileDialogInfo() ICEFFileDialogInfo {
	r := cEFFileDialogInfoAPI().SysCallN(0)
	return AsCEFFileDialogInfo(r)
}

var (
	cEFFileDialogInfoOnce   base.Once
	cEFFileDialogInfoImport *imports.Imports = nil
)

func cEFFileDialogInfoAPI() *imports.Imports {
	cEFFileDialogInfoOnce.Do(func() {
		cEFFileDialogInfoImport = api.NewDefaultImports()
		cEFFileDialogInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFFileDialogInfo_Create", 0), // constructor NewFileDialogInfo
			/* 1 */ imports.NewTable("TCEFFileDialogInfo_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TCEFFileDialogInfo_Mode", 0), // property Mode
			/* 3 */ imports.NewTable("TCEFFileDialogInfo_Title", 0), // property Title
			/* 4 */ imports.NewTable("TCEFFileDialogInfo_DefaultFilePath", 0), // property DefaultFilePath
			/* 5 */ imports.NewTable("TCEFFileDialogInfo_AcceptFilters", 0), // property AcceptFilters
			/* 6 */ imports.NewTable("TCEFFileDialogInfo_AcceptExtensions", 0), // property AcceptExtensions
			/* 7 */ imports.NewTable("TCEFFileDialogInfo_AcceptDescriptions", 0), // property AcceptDescriptions
			/* 8 */ imports.NewTable("TCEFFileDialogInfo_Callback", 0), // property Callback
			/* 9 */ imports.NewTable("TCEFFileDialogInfo_DialogFilter", 0), // property DialogFilter
			/* 10 */ imports.NewTable("TCEFFileDialogInfo_DialogType", 0), // property DialogType
			/* 11 */ imports.NewTable("TCEFFileDialogInfo_DefaultAudioFileDesc", 0), // property DefaultAudioFileDesc
			/* 12 */ imports.NewTable("TCEFFileDialogInfo_DefaultVideoFileDesc", 0), // property DefaultVideoFileDesc
			/* 13 */ imports.NewTable("TCEFFileDialogInfo_DefaultTextFileDesc", 0), // property DefaultTextFileDesc
			/* 14 */ imports.NewTable("TCEFFileDialogInfo_DefaultImageFileDesc", 0), // property DefaultImageFileDesc
			/* 15 */ imports.NewTable("TCEFFileDialogInfo_DefaultAllFileDesc", 0), // property DefaultAllFileDesc
			/* 16 */ imports.NewTable("TCEFFileDialogInfo_DefaultUnknownFileDesc", 0), // property DefaultUnknownFileDesc
		}
	})
	return cEFFileDialogInfoImport
}
