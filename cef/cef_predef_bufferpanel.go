//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ChangeCompositionRange windows
func (m *TBufferPanel) ChangeCompositionRange(selectionrange *TCefRange, characterbounds TCefRectDynArray) {
	predefImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(selectionrange)), uintptr(unsafePointer(&characterbounds[0])), uintptr(int32(len(characterbounds))))
}
