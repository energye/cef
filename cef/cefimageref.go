//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefImage Parent: ICefBaseRefCounted
type ICefImage interface {
	ICefBaseRefCounted
	// IsEmpty
	//  Returns true (1) if this Image is NULL.
	IsEmpty() bool // function
	// IsSame
	//  Returns true (1) if this Image and |that| Image share the same underlying
	//  storage. Will also return true (1) if both images are NULL.
	IsSame(that ICefImage) bool // function
	// AddBitmap
	//  Add a bitmap image representation for |scale_factor|. Only 32-bit
	//  RGBA/BGRA formats are supported. |pixel_width| and |pixel_height| are the
	//  bitmap representation size in pixel coordinates. |pixel_data| is the array
	//  of pixel data and should be |pixel_width| x |pixel_height| x 4 bytes in
	//  size. |color_type| and |alpha_type| values specify the pixel format.
	AddBitmap(scaleFactor float32, pixelWidth int32, pixelHeight int32, colorType cefTypes.TCefColorType, alphaType cefTypes.TCefAlphaType, pixelData uintptr, pixelDataSize cefTypes.NativeUInt) bool // function
	// AddPng
	//  Add a PNG image representation for |scale_factor|. |png_data| is the image
	//  data of size |png_data_size|. Any alpha transparency in the PNG data will
	//  be maintained.
	AddPng(scaleFactor float32, pngData uintptr, pngDataSize cefTypes.NativeUInt) bool // function
	// AddJpeg
	//  Create a JPEG image representation for |scale_factor|. |jpeg_data| is the
	//  image data of size |jpeg_data_size|. The JPEG format does not support
	//  transparency so the alpha byte will be set to 0xFF for all pixels.
	AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize cefTypes.NativeUInt) bool // function
	// GetWidth
	//  Returns the image width in density independent pixel (DIP) units.
	GetWidth() cefTypes.NativeUInt // function
	// GetHeight
	//  Returns the image height in density independent pixel (DIP) units.
	GetHeight() cefTypes.NativeUInt // function
	// HasRepresentation
	//  Returns true (1) if this image contains a representation for
	//  |scale_factor|.
	HasRepresentation(scaleFactor float32) bool // function
	// RemoveRepresentation
	//  Removes the representation for |scale_factor|. Returns true (1) on
	//  success.
	RemoveRepresentation(scaleFactor float32) bool // function
	// GetRepresentationInfo
	//  Returns information for the representation that most closely matches
	//  |scale_factor|. |actual_scale_factor| is the actual scale factor for the
	//  representation. |pixel_width| and |pixel_height| are the representation
	//  size in pixel coordinates. Returns true (1) on success.
	GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth *int32, pixelHeight *int32) bool // function
	// GetAsBitmap
	//  Returns the bitmap representation that most closely matches
	//  |scale_factor|. Only 32-bit RGBA/BGRA formats are supported. |color_type|
	//  and |alpha_type| values specify the desired output pixel format.
	//  |pixel_width| and |pixel_height| are the output representation size in
	//  pixel coordinates. Returns a ICefBinaryValue containing the pixel data
	//  on success or NULL on failure.
	GetAsBitmap(scaleFactor float32, colorType cefTypes.TCefColorType, alphaType cefTypes.TCefAlphaType, pixelWidth *int32, pixelHeight *int32) ICefBinaryValue // function
	// GetAsPng
	//  Returns the PNG representation that most closely matches |scale_factor|.
	//  If |with_transparency| is true (1) any alpha transparency in the image
	//  will be represented in the resulting PNG data. |pixel_width| and
	//  |pixel_height| are the output representation size in pixel coordinates.
	//  Returns a ICefBinaryValue containing the PNG image data on success or
	//  NULL on failure.
	GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth *int32, pixelHeight *int32) ICefBinaryValue // function
	// GetAsJpeg
	//  Returns the JPEG representation that most closely matches |scale_factor|.
	//  |quality| determines the compression level with 0 == lowest and 100 ==
	//  highest. The JPEG format does not support alpha transparency and the alpha
	//  channel, if any, will be discarded. |pixel_width| and |pixel_height| are
	//  the output representation size in pixel coordinates. Returns a
	//  ICefBinaryValue containing the JPEG image data on success or NULL on
	//  failure.
	GetAsJpeg(scaleFactor float32, quality int32, pixelWidth *int32, pixelHeight *int32) ICefBinaryValue // function
}

// ICefImageRef Parent: ICefImage ICefBaseRefCountedRef
type ICefImageRef interface {
	ICefImage
	ICefBaseRefCountedRef
	AsIntfImage() uintptr
}

type TCefImageRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefImageRef) IsEmpty() bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefImageRef) IsSame(that ICefImage) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefImageRef) AddBitmap(scaleFactor float32, pixelWidth int32, pixelHeight int32, colorType cefTypes.TCefColorType, alphaType cefTypes.TCefAlphaType, pixelData uintptr, pixelDataSize cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(pixelWidth), uintptr(pixelHeight), uintptr(colorType), uintptr(alphaType), uintptr(pixelData), uintptr(pixelDataSize))
	return api.GoBool(r)
}

func (m *TCefImageRef) AddPng(scaleFactor float32, pngData uintptr, pngDataSize cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(pngData), uintptr(pngDataSize))
	return api.GoBool(r)
}

func (m *TCefImageRef) AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(jpegData), uintptr(jpegDataSize))
	return api.GoBool(r)
}

func (m *TCefImageRef) GetWidth() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefImageRefAPI().SysCallN(6, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefImageRef) GetHeight() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefImageRefAPI().SysCallN(7, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefImageRef) HasRepresentation(scaleFactor float32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)))
	return api.GoBool(r)
}

func (m *TCefImageRef) RemoveRepresentation(scaleFactor float32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefImageRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)))
	return api.GoBool(r)
}

func (m *TCefImageRef) GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth *int32, pixelHeight *int32) bool {
	if !m.IsValid() {
		return false
	}
	actualScaleFactorPtr := uintptr(*actualScaleFactor)
	pixelWidthPtr := uintptr(*pixelWidth)
	pixelHeightPtr := uintptr(*pixelHeight)
	r := cefImageRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(base.UnsafePointer(&actualScaleFactor)), uintptr(base.UnsafePointer(&pixelWidthPtr)), uintptr(base.UnsafePointer(&pixelHeightPtr)))
	*actualScaleFactor = float32(actualScaleFactorPtr)
	*pixelWidth = int32(pixelWidthPtr)
	*pixelHeight = int32(pixelHeightPtr)
	return api.GoBool(r)
}

func (m *TCefImageRef) GetAsBitmap(scaleFactor float32, colorType cefTypes.TCefColorType, alphaType cefTypes.TCefAlphaType, pixelWidth *int32, pixelHeight *int32) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	pixelWidthPtr := uintptr(*pixelWidth)
	pixelHeightPtr := uintptr(*pixelHeight)
	var resultPtr uintptr
	cefImageRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(colorType), uintptr(alphaType), uintptr(base.UnsafePointer(&pixelWidthPtr)), uintptr(base.UnsafePointer(&pixelHeightPtr)), uintptr(base.UnsafePointer(&resultPtr)))
	*pixelWidth = int32(pixelWidthPtr)
	*pixelHeight = int32(pixelHeightPtr)
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefImageRef) GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth *int32, pixelHeight *int32) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	pixelWidthPtr := uintptr(*pixelWidth)
	pixelHeightPtr := uintptr(*pixelHeight)
	var resultPtr uintptr
	cefImageRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), api.PasBool(withTransparency), uintptr(base.UnsafePointer(&pixelWidthPtr)), uintptr(base.UnsafePointer(&pixelHeightPtr)), uintptr(base.UnsafePointer(&resultPtr)))
	*pixelWidth = int32(pixelWidthPtr)
	*pixelHeight = int32(pixelHeightPtr)
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefImageRef) GetAsJpeg(scaleFactor float32, quality int32, pixelWidth *int32, pixelHeight *int32) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	pixelWidthPtr := uintptr(*pixelWidth)
	pixelHeightPtr := uintptr(*pixelHeight)
	var resultPtr uintptr
	cefImageRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)), uintptr(quality), uintptr(base.UnsafePointer(&pixelWidthPtr)), uintptr(base.UnsafePointer(&pixelHeightPtr)), uintptr(base.UnsafePointer(&resultPtr)))
	*pixelWidth = int32(pixelWidthPtr)
	*pixelHeight = int32(pixelHeightPtr)
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefImageRef) AsIntfImage() uintptr {
	return m.GetIntfPointer(0)
}

// ImageRef  is static instance
var ImageRef _ImageRefClass

// _ImageRefClass is class type defined by TCefImageRef
type _ImageRefClass uintptr

func (_ImageRefClass) UnWrap(data uintptr) (result ICefImage) {
	var resultPtr uintptr
	cefImageRefAPI().SysCallN(14, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (_ImageRefClass) New() (result ICefImage) {
	var resultPtr uintptr
	cefImageRefAPI().SysCallN(15, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

// NewImageRef class constructor
func NewImageRef(data uintptr) ICefImageRef {
	var imagePtr uintptr // ICefImage
	r := cefImageRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&imagePtr)))
	ret := AsCefImageRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, imagePtr)
	}
	return ret
}

var (
	cefImageRefOnce   base.Once
	cefImageRefImport *imports.Imports = nil
)

func cefImageRefAPI() *imports.Imports {
	cefImageRefOnce.Do(func() {
		cefImageRefImport = api.NewDefaultImports()
		cefImageRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefImageRef_Create", 0), // constructor NewImageRef
			/* 1 */ imports.NewTable("TCefImageRef_IsEmpty", 0), // function IsEmpty
			/* 2 */ imports.NewTable("TCefImageRef_IsSame", 0), // function IsSame
			/* 3 */ imports.NewTable("TCefImageRef_AddBitmap", 0), // function AddBitmap
			/* 4 */ imports.NewTable("TCefImageRef_AddPng", 0), // function AddPng
			/* 5 */ imports.NewTable("TCefImageRef_AddJpeg", 0), // function AddJpeg
			/* 6 */ imports.NewTable("TCefImageRef_GetWidth", 0), // function GetWidth
			/* 7 */ imports.NewTable("TCefImageRef_GetHeight", 0), // function GetHeight
			/* 8 */ imports.NewTable("TCefImageRef_HasRepresentation", 0), // function HasRepresentation
			/* 9 */ imports.NewTable("TCefImageRef_RemoveRepresentation", 0), // function RemoveRepresentation
			/* 10 */ imports.NewTable("TCefImageRef_GetRepresentationInfo", 0), // function GetRepresentationInfo
			/* 11 */ imports.NewTable("TCefImageRef_GetAsBitmap", 0), // function GetAsBitmap
			/* 12 */ imports.NewTable("TCefImageRef_GetAsPng", 0), // function GetAsPng
			/* 13 */ imports.NewTable("TCefImageRef_GetAsJpeg", 0), // function GetAsJpeg
			/* 14 */ imports.NewTable("TCefImageRef_UnWrap", 0), // static function UnWrap
			/* 15 */ imports.NewTable("TCefImageRef_New", 0), // static function New
		}
	})
	return cefImageRefImport
}
