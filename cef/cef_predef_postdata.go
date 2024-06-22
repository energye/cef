//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

func (m *TCefPostData) GetElements(elementsCount *NativeUInt, elements *ICefPostDataElementArray) {
	var result uintptr
	predefImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(elementsCount)), uintptr(unsafePointer(&result)))
	*elements = PostDataElementArrayRef.New(int(*elementsCount), result)
}
