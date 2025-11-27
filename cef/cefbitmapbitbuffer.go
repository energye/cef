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
	"github.com/energye/lcl/types"
)

// ICEFBitmapBitBuffer Parent: lcl.IObject
type ICEFBitmapBitBuffer interface {
	lcl.IObject
	// UpdateSize
	//  Updates the image size.
	UpdateSize(width int32, height int32) // procedure
	// Width
	//  Image width.
	Width() int32 // property Width Getter
	// Height
	//  Image height.
	Height() int32 // property Height Getter
	// BufferLength
	//  Buffer length.
	BufferLength() int32 // property BufferLength Getter
	// Empty
	//  Returns true if the buffer is empty.
	Empty() bool // property Empty Getter
	// Scanline
	//  Returns a pointer to the first byte in of the Y scnaline.
	Scanline(Y int32) types.PByte // property Scanline Getter
	// ScanlineSize
	//  Returns the scanline size.
	ScanlineSize() int32 // property ScanlineSize Getter
	// BufferScanlineSize
	//  Returns the real buffer scanline size.
	BufferScanlineSize() int32 // property BufferScanlineSize Getter
	// BufferBits
	//  Returns a pointer to the buffer that stores the image.
	BufferBits() uintptr // property BufferBits Getter
}

type TCEFBitmapBitBuffer struct {
	lcl.TObject
}

func (m *TCEFBitmapBitBuffer) UpdateSize(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	cEFBitmapBitBufferAPI().SysCallN(1, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TCEFBitmapBitBuffer) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCEFBitmapBitBuffer) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TCEFBitmapBitBuffer) BufferLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCEFBitmapBitBuffer) Empty() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFBitmapBitBufferAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFBitmapBitBuffer) Scanline(Y int32) types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(6, m.Instance(), uintptr(Y))
	return types.PByte(r)
}

func (m *TCEFBitmapBitBuffer) ScanlineSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TCEFBitmapBitBuffer) BufferScanlineSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TCEFBitmapBitBuffer) BufferBits() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cEFBitmapBitBufferAPI().SysCallN(9, m.Instance())
	return uintptr(r)
}

// NewBitmapBitBuffer class constructor
func NewBitmapBitBuffer(width int32, height int32) ICEFBitmapBitBuffer {
	r := cEFBitmapBitBufferAPI().SysCallN(0, uintptr(width), uintptr(height))
	return AsCEFBitmapBitBuffer(r)
}

var (
	cEFBitmapBitBufferOnce   base.Once
	cEFBitmapBitBufferImport *imports.Imports = nil
)

func cEFBitmapBitBufferAPI() *imports.Imports {
	cEFBitmapBitBufferOnce.Do(func() {
		cEFBitmapBitBufferImport = api.NewDefaultImports()
		cEFBitmapBitBufferImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFBitmapBitBuffer_Create", 0), // constructor NewBitmapBitBuffer
			/* 1 */ imports.NewTable("TCEFBitmapBitBuffer_UpdateSize", 0), // procedure UpdateSize
			/* 2 */ imports.NewTable("TCEFBitmapBitBuffer_Width", 0), // property Width
			/* 3 */ imports.NewTable("TCEFBitmapBitBuffer_Height", 0), // property Height
			/* 4 */ imports.NewTable("TCEFBitmapBitBuffer_BufferLength", 0), // property BufferLength
			/* 5 */ imports.NewTable("TCEFBitmapBitBuffer_Empty", 0), // property Empty
			/* 6 */ imports.NewTable("TCEFBitmapBitBuffer_Scanline", 0), // property Scanline
			/* 7 */ imports.NewTable("TCEFBitmapBitBuffer_ScanlineSize", 0), // property ScanlineSize
			/* 8 */ imports.NewTable("TCEFBitmapBitBuffer_BufferScanlineSize", 0), // property BufferScanlineSize
			/* 9 */ imports.NewTable("TCEFBitmapBitBuffer_BufferBits", 0), // property BufferBits
		}
	})
	return cEFBitmapBitBufferImport
}
