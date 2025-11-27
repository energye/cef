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

	cefTypes "github.com/energye/cef/types"
)

// ICEFJson Parent: lcl.IObject
type ICEFJson interface {
	lcl.IObject
}

type TCEFJson struct {
	lcl.TObject
}

// Json  is static instance
var Json _JsonClass

// _JsonClass is class type defined by TCEFJson
type _JsonClass uintptr

// ReadValue
//
//	Returns the ICefValue value at the specified key.
func (_JsonClass) ReadValue(dictionary ICefDictionaryValue, key string, value *ICefValue) bool {
	valuePtr := base.GetObjectUintptr(*value)
	r := cEFJsonAPI().SysCallN(0, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCefValueRef(valuePtr)
	return api.GoBool(r)
}

// ReadBoolean
//
//	Returns the boolean value at the specified key.
func (_JsonClass) ReadBoolean(dictionary ICefDictionaryValue, key string, value *bool) bool {
	r := cEFJsonAPI().SysCallN(1, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(value)))
	return api.GoBool(r)
}

// ReadInteger
//
//	Returns the integer value at the specified key.
func (_JsonClass) ReadInteger(dictionary ICefDictionaryValue, key string, value *int32) bool {
	valuePtr := uintptr(*value)
	r := cEFJsonAPI().SysCallN(2, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = int32(valuePtr)
	return api.GoBool(r)
}

// ReadDouble
//
//	Returns the double value at the specified key.
func (_JsonClass) ReadDouble(dictionary ICefDictionaryValue, key string, value *float64) bool {
	valuePtr := uintptr(*value)
	r := cEFJsonAPI().SysCallN(3, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&value)))
	*value = float64(valuePtr)
	return api.GoBool(r)
}

// ReadString
//
//	Returns the ustring value at the specified key.
func (_JsonClass) ReadString(dictionary ICefDictionaryValue, key string, value *string) bool {
	valuePtr := api.PasStr(*value)
	r := cEFJsonAPI().SysCallN(4, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = api.GoStr(valuePtr)
	return api.GoBool(r)
}

// ReadBinary
//
//	Returns the ICefBinaryValue value at the specified key.
func (_JsonClass) ReadBinary(dictionary ICefDictionaryValue, key string, value *ICefBinaryValue) bool {
	valuePtr := base.GetObjectUintptr(*value)
	r := cEFJsonAPI().SysCallN(5, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCefBinaryValueRef(valuePtr)
	return api.GoBool(r)
}

// ReadDictionary
//
//	Returns the ICefDictionaryValue value at the specified key.
func (_JsonClass) ReadDictionary(dictionary ICefDictionaryValue, key string, value *ICefDictionaryValue) bool {
	valuePtr := base.GetObjectUintptr(*value)
	r := cEFJsonAPI().SysCallN(6, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCefDictionaryValueRef(valuePtr)
	return api.GoBool(r)
}

// ReadList
//
//	Returns the ICefListValue value at the specified key.
func (_JsonClass) ReadList(dictionary ICefDictionaryValue, key string, value *ICefListValue) bool {
	valuePtr := base.GetObjectUintptr(*value)
	r := cEFJsonAPI().SysCallN(7, base.GetObjectUintptr(dictionary), api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCefListValueRef(valuePtr)
	return api.GoBool(r)
}

// ParseWithStringJsonParserOptions
//
//	Parses the specified |json_string| and returns a dictionary or list
//	representation. If JSON parsing fails this function returns NULL.
func (_JsonClass) ParseWithStringJsonParserOptions(jsonString string, options cefTypes.TCefJsonParserOptions) (result ICefValue) {
	var resultPtr uintptr
	cEFJsonAPI().SysCallN(8, api.PasStr(jsonString), uintptr(options), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

// ParseWithPointerNativeUIntJsonParserOptions
//
//	Parses the specified UTF8-encoded |json| buffer of size |json_size| and
//	returns a dictionary or list representation. If JSON parsing fails this
//	function returns NULL.
func (_JsonClass) ParseWithPointerNativeUIntJsonParserOptions(json uintptr, jsonSize cefTypes.NativeUInt, options cefTypes.TCefJsonParserOptions) (result ICefValue) {
	var resultPtr uintptr
	cEFJsonAPI().SysCallN(9, uintptr(json), uintptr(jsonSize), uintptr(options), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

// ParseAndReturnError
//
//	Parses the specified |json_string| and returns a dictionary or list
//	representation. If JSON parsing fails this function returns NULL and
//	populates |error_msg_out| with a formatted error message.
func (_JsonClass) ParseAndReturnError(jsonString string, options cefTypes.TCefJsonParserOptions, outErrorMsgOut *string) (result ICefValue) {
	var errorMsgOutPtr uintptr
	var resultPtr uintptr
	cEFJsonAPI().SysCallN(10, api.PasStr(jsonString), uintptr(options), uintptr(base.UnsafePointer(&errorMsgOutPtr)), uintptr(base.UnsafePointer(&resultPtr)))
	*outErrorMsgOut = api.GoStr(errorMsgOutPtr)
	result = AsCefValueRef(resultPtr)
	return
}

// WriteWithValueJsonWriterOptions
//
//	Generates a JSON string from the specified root |node|.
//	Returns an NULL string on failure. This function
//	requires exclusive access to |node| including any underlying data.
func (_JsonClass) WriteWithValueJsonWriterOptions(node ICefValue, options cefTypes.TCefJsonWriterOptions) string {
	r := cEFJsonAPI().SysCallN(11, base.GetObjectUintptr(node), uintptr(options))
	return api.GoStr(r)
}

// WriteWithDictionaryValueJsonWriterOptions
//
//	Generates a JSON string from the specified root |node|.
//	Returns an NULL string on failure. This function
//	requires exclusive access to |node| including any underlying data.
func (_JsonClass) WriteWithDictionaryValueJsonWriterOptions(node ICefDictionaryValue, options cefTypes.TCefJsonWriterOptions) string {
	r := cEFJsonAPI().SysCallN(12, base.GetObjectUintptr(node), uintptr(options))
	return api.GoStr(r)
}

// WriteWithValueStringList
//
//	Generates a JSON string from the specified root |node|.
//	Returns an NULL string on failure. This function
//	requires exclusive access to |node| including any underlying data.
func (_JsonClass) WriteWithValueStringList(node ICefValue, rsltStrings *lcl.IStringList) bool {
	rsltStringsPtr := base.GetObjectUintptr(*rsltStrings)
	r := cEFJsonAPI().SysCallN(13, base.GetObjectUintptr(node), uintptr(base.UnsafePointer(&rsltStringsPtr)))
	*rsltStrings = lcl.AsStringList(rsltStringsPtr)
	return api.GoBool(r)
}

// WriteWithDictionaryValueStringList
//
//	Generates a JSON string from the specified root |node|.
//	Returns an NULL string on failure. This function
//	requires exclusive access to |node| including any underlying data.
func (_JsonClass) WriteWithDictionaryValueStringList(node ICefDictionaryValue, rsltStrings *lcl.IStringList) bool {
	rsltStringsPtr := base.GetObjectUintptr(*rsltStrings)
	r := cEFJsonAPI().SysCallN(14, base.GetObjectUintptr(node), uintptr(base.UnsafePointer(&rsltStringsPtr)))
	*rsltStrings = lcl.AsStringList(rsltStringsPtr)
	return api.GoBool(r)
}

// SaveToFileWithValueString
//
//	Saves the JSON data in |node| to a file in aFileName.
func (_JsonClass) SaveToFileWithValueString(node ICefValue, fileName string) bool {
	r := cEFJsonAPI().SysCallN(15, base.GetObjectUintptr(node), api.PasStr(fileName))
	return api.GoBool(r)
}

// SaveToFileWithDictionaryValueString
//
//	Saves the JSON data in |node| to a file in aFileName.
func (_JsonClass) SaveToFileWithDictionaryValueString(node ICefDictionaryValue, fileName string) bool {
	r := cEFJsonAPI().SysCallN(16, base.GetObjectUintptr(node), api.PasStr(fileName))
	return api.GoBool(r)
}

// LoadFromFile
//
//	Loads the JSON data in |aFileName| using the |encoding| and returns an ICefValue node in |aRsltNode|.
func (_JsonClass) LoadFromFile(fileName string, rsltNode *ICefValue, options cefTypes.TCefJsonParserOptions) bool {
	rsltNodePtr := base.GetObjectUintptr(*rsltNode)
	r := cEFJsonAPI().SysCallN(17, api.PasStr(fileName), uintptr(base.UnsafePointer(&rsltNodePtr)), uintptr(options))
	*rsltNode = AsCefValueRef(rsltNodePtr)
	return api.GoBool(r)
}

var (
	cEFJsonOnce   base.Once
	cEFJsonImport *imports.Imports = nil
)

func cEFJsonAPI() *imports.Imports {
	cEFJsonOnce.Do(func() {
		cEFJsonImport = api.NewDefaultImports()
		cEFJsonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFJson_ReadValue", 0), // static function ReadValue
			/* 1 */ imports.NewTable("TCEFJson_ReadBoolean", 0), // static function ReadBoolean
			/* 2 */ imports.NewTable("TCEFJson_ReadInteger", 0), // static function ReadInteger
			/* 3 */ imports.NewTable("TCEFJson_ReadDouble", 0), // static function ReadDouble
			/* 4 */ imports.NewTable("TCEFJson_ReadString", 0), // static function ReadString
			/* 5 */ imports.NewTable("TCEFJson_ReadBinary", 0), // static function ReadBinary
			/* 6 */ imports.NewTable("TCEFJson_ReadDictionary", 0), // static function ReadDictionary
			/* 7 */ imports.NewTable("TCEFJson_ReadList", 0), // static function ReadList
			/* 8 */ imports.NewTable("TCEFJson_ParseWithStringJsonParserOptions", 0), // static function ParseWithStringJsonParserOptions
			/* 9 */ imports.NewTable("TCEFJson_ParseWithPointerNativeUIntJsonParserOptions", 0), // static function ParseWithPointerNativeUIntJsonParserOptions
			/* 10 */ imports.NewTable("TCEFJson_ParseAndReturnError", 0), // static function ParseAndReturnError
			/* 11 */ imports.NewTable("TCEFJson_WriteWithValueJsonWriterOptions", 0), // static function WriteWithValueJsonWriterOptions
			/* 12 */ imports.NewTable("TCEFJson_WriteWithDictionaryValueJsonWriterOptions", 0), // static function WriteWithDictionaryValueJsonWriterOptions
			/* 13 */ imports.NewTable("TCEFJson_WriteWithValueStringList", 0), // static function WriteWithValueStringList
			/* 14 */ imports.NewTable("TCEFJson_WriteWithDictionaryValueStringList", 0), // static function WriteWithDictionaryValueStringList
			/* 15 */ imports.NewTable("TCEFJson_SaveToFileWithValueString", 0), // static function SaveToFileWithValueString
			/* 16 */ imports.NewTable("TCEFJson_SaveToFileWithDictionaryValueString", 0), // static function SaveToFileWithDictionaryValueString
			/* 17 */ imports.NewTable("TCEFJson_LoadFromFile", 0), // static function LoadFromFile
		}
	})
	return cEFJsonImport
}
