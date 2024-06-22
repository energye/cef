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

// ICEFFileDialogInfo Parent: IObject
type ICEFFileDialogInfo interface {
	IObject
	Mode() uint32                              // property
	SetMode(AValue uint32)                     // property
	Title() string                             // property
	SetTitle(AValue string)                    // property
	DefaultFilePath() string                   // property
	SetDefaultFilePath(AValue string)          // property
	SetAcceptFilters(AValue IStrings)          // property
	Callback() ICefFileDialogCallback          // property
	SetCallback(AValue ICefFileDialogCallback) // property
	DialogFilter() string                      // property
	DialogType() TCEFDialogType                // property
	DefaultAudioFileDesc() string              // property
	SetDefaultAudioFileDesc(AValue string)     // property
	DefaultVideoFileDesc() string              // property
	SetDefaultVideoFileDesc(AValue string)     // property
	DefaultTextFileDesc() string               // property
	SetDefaultTextFileDesc(AValue string)      // property
	DefaultImageFileDesc() string              // property
	SetDefaultImageFileDesc(AValue string)     // property
	DefaultAllFileDesc() string                // property
	SetDefaultAllFileDesc(AValue string)       // property
	DefaultUnknownFileDesc() string            // property
	SetDefaultUnknownFileDesc(AValue string)   // property
	Clear()                                    // procedure
}

// TCEFFileDialogInfo Parent: TObject
type TCEFFileDialogInfo struct {
	TObject
}

func NewCEFFileDialogInfo() ICEFFileDialogInfo {
	r1 := fileDialogInfoImportAPI().SysCallN(3)
	return AsCEFFileDialogInfo(r1)
}

func (m *TCEFFileDialogInfo) Mode() uint32 {
	r1 := fileDialogInfoImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFFileDialogInfo) SetMode(AValue uint32) {
	fileDialogInfoImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFFileDialogInfo) Title() string {
	r1 := fileDialogInfoImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetTitle(AValue string) {
	fileDialogInfoImportAPI().SysCallN(14, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultFilePath() string {
	r1 := fileDialogInfoImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultFilePath(AValue string) {
	fileDialogInfoImportAPI().SysCallN(6, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) SetAcceptFilters(AValue IStrings) {
	fileDialogInfoImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCEFFileDialogInfo) Callback() ICefFileDialogCallback {
	var resultCefFileDialogCallback uintptr
	fileDialogInfoImportAPI().SysCallN(1, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefFileDialogCallback)))
	return AsCefFileDialogCallback(resultCefFileDialogCallback)
}

func (m *TCEFFileDialogInfo) SetCallback(AValue ICefFileDialogCallback) {
	fileDialogInfoImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFFileDialogInfo) DialogFilter() string {
	r1 := fileDialogInfoImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) DialogType() TCEFDialogType {
	r1 := fileDialogInfoImportAPI().SysCallN(12, m.Instance())
	return TCEFDialogType(r1)
}

func (m *TCEFFileDialogInfo) DefaultAudioFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultAudioFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultVideoFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultVideoFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultTextFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultTextFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(8, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultImageFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultImageFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultAllFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultAllFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultUnknownFileDesc() string {
	r1 := fileDialogInfoImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultUnknownFileDesc(AValue string) {
	fileDialogInfoImportAPI().SysCallN(9, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) Clear() {
	fileDialogInfoImportAPI().SysCallN(2, m.Instance())
}

var (
	fileDialogInfoImport       *imports.Imports = nil
	fileDialogInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFFileDialogInfo_AcceptFilters", 0),
		/*1*/ imports.NewTable("CEFFileDialogInfo_Callback", 0),
		/*2*/ imports.NewTable("CEFFileDialogInfo_Clear", 0),
		/*3*/ imports.NewTable("CEFFileDialogInfo_Create", 0),
		/*4*/ imports.NewTable("CEFFileDialogInfo_DefaultAllFileDesc", 0),
		/*5*/ imports.NewTable("CEFFileDialogInfo_DefaultAudioFileDesc", 0),
		/*6*/ imports.NewTable("CEFFileDialogInfo_DefaultFilePath", 0),
		/*7*/ imports.NewTable("CEFFileDialogInfo_DefaultImageFileDesc", 0),
		/*8*/ imports.NewTable("CEFFileDialogInfo_DefaultTextFileDesc", 0),
		/*9*/ imports.NewTable("CEFFileDialogInfo_DefaultUnknownFileDesc", 0),
		/*10*/ imports.NewTable("CEFFileDialogInfo_DefaultVideoFileDesc", 0),
		/*11*/ imports.NewTable("CEFFileDialogInfo_DialogFilter", 0),
		/*12*/ imports.NewTable("CEFFileDialogInfo_DialogType", 0),
		/*13*/ imports.NewTable("CEFFileDialogInfo_Mode", 0),
		/*14*/ imports.NewTable("CEFFileDialogInfo_Title", 0),
	}
)

func fileDialogInfoImportAPI() *imports.Imports {
	if fileDialogInfoImport == nil {
		fileDialogInfoImport = NewDefaultImports()
		fileDialogInfoImport.SetImportTable(fileDialogInfoImportTables)
		fileDialogInfoImportTables = nil
	}
	return fileDialogInfoImport
}
