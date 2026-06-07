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
	"github.com/energye/lcl/types"
)

// ICEFOSRIMEHandler Parent: IObject
type ICEFOSRIMEHandler interface {
	IObject
	GetResult(param types.LParam, result *string) bool                                                                                   // function
	GetComposition(param types.LParam, compositionText *string, underlines *ICefCompositionUnderlineArray, compositionStart *int32) bool // function
	SetInputLanguage()                                                                                                                   // procedure
	CreateImeWindow()                                                                                                                    // procedure
	DestroyImeWindow()                                                                                                                   // procedure
	CleanupComposition()                                                                                                                 // procedure
	ResetComposition()                                                                                                                   // procedure
	EnableIME()                                                                                                                          // procedure
	DisableIME()                                                                                                                         // procedure
	CancelIME()                                                                                                                          // procedure
	UpdateCaretPosition(index int32)                                                                                                     // procedure
	ChangeCompositionRange(selectionRange TCefRange, characterBounds ICefRectArray)                                                      // procedure
	MoveImeWindow()                                                                                                                      // procedure
	IsComposing() bool                                                                                                                   // property IsComposing Getter
	InputLanguageID() types.LANGID                                                                                                       // property InputLanguageID Getter
	PrimaryLangID() uint16                                                                                                               // property PrimaryLangID Getter
	SubLangID() uint16                                                                                                                   // property SubLangID Getter
	Initialized() bool                                                                                                                   // property Initialized Getter
}

type TCEFOSRIMEHandler struct {
	TObject
}

func (m *TCEFOSRIMEHandler) GetResult(param types.LParam, result *string) bool {
	if !m.IsValid() {
		return false
	}
	resultPtr := api.PasStr(*result)
	r := cEFOSRIMEHandlerAPI().SysCallN(1, m.Instance(), uintptr(param), uintptr(base.UnsafePointer(&resultPtr)))
	*result = api.GoStr(resultPtr)
	return api.GoBool(r)
}

func (m *TCEFOSRIMEHandler) GetComposition(param types.LParam, compositionText *string, underlines *ICefCompositionUnderlineArray, compositionStart *int32) bool {
	if !m.IsValid() {
		return false
	}
	compositionTextPtr := api.PasStr(*compositionText)
	var underlinesPtr uintptr
	var underlinesCountPtr uintptr
	compositionStartPtr := uintptr(*compositionStart)
	r := cEFOSRIMEHandlerAPI().SysCallN(2, m.Instance(), uintptr(param), uintptr(base.UnsafePointer(&compositionTextPtr)), uintptr(base.UnsafePointer(&underlinesPtr)), uintptr(base.UnsafePointer(&underlinesCountPtr)), uintptr(base.UnsafePointer(&compositionStartPtr)))
	*compositionText = api.GoStr(compositionTextPtr)
	*underlines = NewCefCompositionUnderlineDynArray(int(underlinesCountPtr), underlinesPtr)
	*compositionStart = int32(compositionStartPtr)
	return api.GoBool(r)
}

func (m *TCEFOSRIMEHandler) SetInputLanguage() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(3, m.Instance())
}

func (m *TCEFOSRIMEHandler) CreateImeWindow() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(4, m.Instance())
}

func (m *TCEFOSRIMEHandler) DestroyImeWindow() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(5, m.Instance())
}

func (m *TCEFOSRIMEHandler) CleanupComposition() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(6, m.Instance())
}

func (m *TCEFOSRIMEHandler) ResetComposition() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(7, m.Instance())
}

func (m *TCEFOSRIMEHandler) EnableIME() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(8, m.Instance())
}

func (m *TCEFOSRIMEHandler) DisableIME() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(9, m.Instance())
}

func (m *TCEFOSRIMEHandler) CancelIME() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(10, m.Instance())
}

func (m *TCEFOSRIMEHandler) UpdateCaretPosition(index int32) {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(11, m.Instance(), uintptr(index))
}

func (m *TCEFOSRIMEHandler) ChangeCompositionRange(selectionRange TCefRange, characterBounds ICefRectArray) {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&selectionRange)), characterBounds.Instance(), uintptr(int32(characterBounds.Count())))
}

func (m *TCEFOSRIMEHandler) MoveImeWindow() {
	if !m.IsValid() {
		return
	}
	cEFOSRIMEHandlerAPI().SysCallN(13, m.Instance())
}

func (m *TCEFOSRIMEHandler) IsComposing() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOSRIMEHandlerAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFOSRIMEHandler) InputLanguageID() types.LANGID {
	if !m.IsValid() {
		return 0
	}
	r := cEFOSRIMEHandlerAPI().SysCallN(15, m.Instance())
	return types.LANGID(r)
}

func (m *TCEFOSRIMEHandler) PrimaryLangID() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cEFOSRIMEHandlerAPI().SysCallN(16, m.Instance())
	return uint16(r)
}

func (m *TCEFOSRIMEHandler) SubLangID() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cEFOSRIMEHandlerAPI().SysCallN(17, m.Instance())
	return uint16(r)
}

func (m *TCEFOSRIMEHandler) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOSRIMEHandlerAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

// NewOSRIMEHandler class constructor
func NewOSRIMEHandler(hWND types.HWND) ICEFOSRIMEHandler {
	r := cEFOSRIMEHandlerAPI().SysCallN(0, uintptr(hWND))
	return AsCEFOSRIMEHandler(r)
}

var (
	cEFOSRIMEHandlerOnce   base.Once
	cEFOSRIMEHandlerImport *imports.Imports = nil
)

func cEFOSRIMEHandlerAPI() *imports.Imports {
	cEFOSRIMEHandlerOnce.Do(func() {
		cEFOSRIMEHandlerImport = api.NewDefaultImports()
		cEFOSRIMEHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFOSRIMEHandler_Create", 0), // constructor NewOSRIMEHandler
			/* 1 */ imports.NewTable("TCEFOSRIMEHandler_GetResult", 0), // function GetResult
			/* 2 */ imports.NewTable("TCEFOSRIMEHandler_GetComposition", 0), // function GetComposition
			/* 3 */ imports.NewTable("TCEFOSRIMEHandler_SetInputLanguage", 0), // procedure SetInputLanguage
			/* 4 */ imports.NewTable("TCEFOSRIMEHandler_CreateImeWindow", 0), // procedure CreateImeWindow
			/* 5 */ imports.NewTable("TCEFOSRIMEHandler_DestroyImeWindow", 0), // procedure DestroyImeWindow
			/* 6 */ imports.NewTable("TCEFOSRIMEHandler_CleanupComposition", 0), // procedure CleanupComposition
			/* 7 */ imports.NewTable("TCEFOSRIMEHandler_ResetComposition", 0), // procedure ResetComposition
			/* 8 */ imports.NewTable("TCEFOSRIMEHandler_EnableIME", 0), // procedure EnableIME
			/* 9 */ imports.NewTable("TCEFOSRIMEHandler_DisableIME", 0), // procedure DisableIME
			/* 10 */ imports.NewTable("TCEFOSRIMEHandler_CancelIME", 0), // procedure CancelIME
			/* 11 */ imports.NewTable("TCEFOSRIMEHandler_UpdateCaretPosition", 0), // procedure UpdateCaretPosition
			/* 12 */ imports.NewTable("TCEFOSRIMEHandler_ChangeCompositionRange", 0), // procedure ChangeCompositionRange
			/* 13 */ imports.NewTable("TCEFOSRIMEHandler_MoveImeWindow", 0), // procedure MoveImeWindow
			/* 14 */ imports.NewTable("TCEFOSRIMEHandler_IsComposing", 0), // property IsComposing
			/* 15 */ imports.NewTable("TCEFOSRIMEHandler_InputLanguageID", 0), // property InputLanguageID
			/* 16 */ imports.NewTable("TCEFOSRIMEHandler_PrimaryLangID", 0), // property PrimaryLangID
			/* 17 */ imports.NewTable("TCEFOSRIMEHandler_SubLangID", 0), // property SubLangID
			/* 18 */ imports.NewTable("TCEFOSRIMEHandler_Initialized", 0), // property Initialized
		}
	})
	return cEFOSRIMEHandlerImport
}
