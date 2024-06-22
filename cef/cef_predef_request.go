//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// GetHeaderMap
func (m *TCefRequest) GetHeaderMap() ICefStringMultimap {
	var result uintptr
	predefImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(&result)))
	if result > 0 {
		return AsCefStringMultimap(result)
	}
	return nil
}
