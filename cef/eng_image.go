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

// ICefImage Parent: ICefBaseRefCounted
type ICefImage interface {
	ICefBaseRefCounted
	IsEmpty() bool                                                                                                                                                    // function
	IsSame(that ICefImage) bool                                                                                                                                       // function
	AddBitmap(scaleFactor float32, pixelWidth, pixelHeight int32, colorType TCefColorType, alphaType TCefAlphaType, pixelData uintptr, pixelDataSize NativeUInt) bool // function
	AddPng(scaleFactor float32, pngData uintptr, pngDataSize NativeUInt) bool                                                                                         // function
	AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize NativeUInt) bool                                                                                      // function
	GetWidth() NativeUInt                                                                                                                                             // function
	GetHeight() NativeUInt                                                                                                                                            // function
	HasRepresentation(scaleFactor float32) bool                                                                                                                       // function
	RemoveRepresentation(scaleFactor float32) bool                                                                                                                    // function
	GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth, pixelHeight *int32) bool                                                       // function
	GetAsBitmap(scaleFactor float32, colorType TCefColorType, alphaType TCefAlphaType, pixelWidth, pixelHeight *int32) ICefBinaryValue                                // function
	GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth, pixelHeight *int32) ICefBinaryValue                                                              // function
	GetAsJpeg(scaleFactor float32, quality int32, pixelWidth, pixelHeight *int32) ICefBinaryValue                                                                     // function
}

// TCefImage Parent: TCefBaseRefCounted
type TCefImage struct {
	TCefBaseRefCounted
}

// ImageRef -> ICefImage
var ImageRef image

// image TCefImage Ref
type image uintptr

func (m *image) UnWrap(data uintptr) ICefImage {
	var resultCefImage uintptr
	mageImportAPI().SysCallN(14, uintptr(data), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *image) New() ICefImage {
	var resultCefImage uintptr
	mageImportAPI().SysCallN(12, uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefImage) IsEmpty() bool {
	r1 := mageImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TCefImage) IsSame(that ICefImage) bool {
	r1 := mageImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefImage) AddBitmap(scaleFactor float32, pixelWidth, pixelHeight int32, colorType TCefColorType, alphaType TCefAlphaType, pixelData uintptr, pixelDataSize NativeUInt) bool {
	r1 := mageImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(pixelWidth), uintptr(pixelHeight), uintptr(colorType), uintptr(alphaType), uintptr(pixelData), uintptr(pixelDataSize))
	return GoBool(r1)
}

func (m *TCefImage) AddPng(scaleFactor float32, pngData uintptr, pngDataSize NativeUInt) bool {
	r1 := mageImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(pngData), uintptr(pngDataSize))
	return GoBool(r1)
}

func (m *TCefImage) AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize NativeUInt) bool {
	r1 := mageImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(jpegData), uintptr(jpegDataSize))
	return GoBool(r1)
}

func (m *TCefImage) GetWidth() NativeUInt {
	r1 := mageImportAPI().SysCallN(8, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefImage) GetHeight() NativeUInt {
	r1 := mageImportAPI().SysCallN(6, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefImage) HasRepresentation(scaleFactor float32) bool {
	r1 := mageImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&scaleFactor)))
	return GoBool(r1)
}

func (m *TCefImage) RemoveRepresentation(scaleFactor float32) bool {
	r1 := mageImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&scaleFactor)))
	return GoBool(r1)
}

func (m *TCefImage) GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth, pixelHeight *int32) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	r1 := mageImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*actualScaleFactor = float32(result1)
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return GoBool(r1)
}

func (m *TCefImage) GetAsBitmap(scaleFactor float32, colorType TCefColorType, alphaType TCefAlphaType, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result3 uintptr
	var result4 uintptr
	var resultCefBinaryValue uintptr
	mageImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(colorType), uintptr(alphaType), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result3)
	*pixelHeight = int32(result4)
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefImage) GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result2 uintptr
	var result3 uintptr
	var resultCefBinaryValue uintptr
	mageImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&scaleFactor)), PascalBool(withTransparency), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefImage) GetAsJpeg(scaleFactor float32, quality int32, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result2 uintptr
	var result3 uintptr
	var resultCefBinaryValue uintptr
	mageImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(quality), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return AsCefBinaryValue(resultCefBinaryValue)
}

var (
	mageImport       *imports.Imports = nil
	mageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefImage_AddBitmap", 0),
		/*1*/ imports.NewTable("CefImage_AddJpeg", 0),
		/*2*/ imports.NewTable("CefImage_AddPng", 0),
		/*3*/ imports.NewTable("CefImage_GetAsBitmap", 0),
		/*4*/ imports.NewTable("CefImage_GetAsJpeg", 0),
		/*5*/ imports.NewTable("CefImage_GetAsPng", 0),
		/*6*/ imports.NewTable("CefImage_GetHeight", 0),
		/*7*/ imports.NewTable("CefImage_GetRepresentationInfo", 0),
		/*8*/ imports.NewTable("CefImage_GetWidth", 0),
		/*9*/ imports.NewTable("CefImage_HasRepresentation", 0),
		/*10*/ imports.NewTable("CefImage_IsEmpty", 0),
		/*11*/ imports.NewTable("CefImage_IsSame", 0),
		/*12*/ imports.NewTable("CefImage_New", 0),
		/*13*/ imports.NewTable("CefImage_RemoveRepresentation", 0),
		/*14*/ imports.NewTable("CefImage_UnWrap", 0),
	}
)

func mageImportAPI() *imports.Imports {
	if mageImport == nil {
		mageImport = NewDefaultImports()
		mageImport.SetImportTable(mageImportTables)
		mageImportTables = nil
	}
	return mageImport
}
