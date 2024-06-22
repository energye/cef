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

// ICEFBitmapBitBuffer Parent: IObject
//
//	Class that stores a copy of the raw bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
type ICEFBitmapBitBuffer interface {
	IObject
	// Width
	//  Image width.
	Width() int32 // property
	// Height
	//  Image height.
	Height() int32 // property
	// BufferLength
	//  Buffer length.
	BufferLength() int32 // property
	// Empty
	//  Returns true if the buffer is empty.
	Empty() bool // property
	// Scanline
	//  Returns a pointer to the first byte in of the Y scnaline.
	Scanline(y int32) PByte // property
	// ScanlineSize
	//  Returns the scanline size.
	ScanlineSize() int32 // property
	// BufferScanlineSize
	//  Returns the real buffer scanline size.
	BufferScanlineSize() int32 // property
	// BufferBits
	//  Returns a pointer to the buffer that stores the image.
	BufferBits() uintptr // property
	// UpdateSize
	//  Updates the image size.
	UpdateSize(aWidth, aHeight int32) // procedure
}

// TCEFBitmapBitBuffer Parent: TObject
//
//	Class that stores a copy of the raw bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
type TCEFBitmapBitBuffer struct {
	TObject
}

func NewCEFBitmapBitBuffer(aWidth, aHeight int32) ICEFBitmapBitBuffer {
	r1 := bitmapBitBufferImportAPI().SysCallN(3, uintptr(aWidth), uintptr(aHeight))
	return AsCEFBitmapBitBuffer(r1)
}

func (m *TCEFBitmapBitBuffer) Width() int32 {
	r1 := bitmapBitBufferImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) Height() int32 {
	r1 := bitmapBitBufferImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferLength() int32 {
	r1 := bitmapBitBufferImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) Empty() bool {
	r1 := bitmapBitBufferImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCEFBitmapBitBuffer) Scanline(y int32) PByte {
	r1 := bitmapBitBufferImportAPI().SysCallN(6, m.Instance(), uintptr(y))
	return PByte(r1)
}

func (m *TCEFBitmapBitBuffer) ScanlineSize() int32 {
	r1 := bitmapBitBufferImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferScanlineSize() int32 {
	r1 := bitmapBitBufferImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferBits() uintptr {
	r1 := bitmapBitBufferImportAPI().SysCallN(0, m.Instance())
	return uintptr(r1)
}

func (m *TCEFBitmapBitBuffer) UpdateSize(aWidth, aHeight int32) {
	bitmapBitBufferImportAPI().SysCallN(8, m.Instance(), uintptr(aWidth), uintptr(aHeight))
}

var (
	bitmapBitBufferImport       *imports.Imports = nil
	bitmapBitBufferImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFBitmapBitBuffer_BufferBits", 0),
		/*1*/ imports.NewTable("CEFBitmapBitBuffer_BufferLength", 0),
		/*2*/ imports.NewTable("CEFBitmapBitBuffer_BufferScanlineSize", 0),
		/*3*/ imports.NewTable("CEFBitmapBitBuffer_Create", 0),
		/*4*/ imports.NewTable("CEFBitmapBitBuffer_Empty", 0),
		/*5*/ imports.NewTable("CEFBitmapBitBuffer_Height", 0),
		/*6*/ imports.NewTable("CEFBitmapBitBuffer_Scanline", 0),
		/*7*/ imports.NewTable("CEFBitmapBitBuffer_ScanlineSize", 0),
		/*8*/ imports.NewTable("CEFBitmapBitBuffer_UpdateSize", 0),
		/*9*/ imports.NewTable("CEFBitmapBitBuffer_Width", 0),
	}
)

func bitmapBitBufferImportAPI() *imports.Imports {
	if bitmapBitBufferImport == nil {
		bitmapBitBufferImport = NewDefaultImports()
		bitmapBitBufferImport.SetImportTable(bitmapBitBufferImportTables)
		bitmapBitBufferImportTables = nil
	}
	return bitmapBitBufferImport
}
