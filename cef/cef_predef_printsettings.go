//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/tools/ptr"
	"unsafe"
)

func (m *TCefPrintSettings) SetPageRanges(ranges TRangeArray) {
	predefImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&ranges[0])), uintptr(int32(len(ranges))))
}

func (m *TCefPrintSettings) GetPageRanges() TRangeArray {
	var rangesPtr uintptr
	var size uintptr
	predefImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&rangesPtr)), uintptr(unsafePointer(&size)))
	if size > 0 {
		rangeSizeOf := unsafe.Sizeof(TCefRange{})
		ranges := make(TRangeArray, size, size)
		for i := 0; i < int(size); i++ {
			ranges[i] = *(*TCefRange)(ptr.GetParamPtr(rangesPtr, i*int(rangeSizeOf)))
		}
		return ranges
	}
	return nil
}
