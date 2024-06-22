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

// ICEFJson Parent: IObject
type ICEFJson interface {
	IObject
}

// TCEFJson Parent: TObject
type TCEFJson struct {
	TObject
}

// JsonRef -> ICEFJson
var JsonRef json

// json TCEFJson Ref
type json uintptr

func (m *json) ReadValue(aDictionary ICefDictionaryValue, aKey string, aValue *ICefValue) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(11, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefValue(result2)
	return GoBool(r1)
}

func (m *json) ReadBoolean(aDictionary ICefDictionaryValue, aKey string, aValue *bool) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(5, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = GoBool(result2)
	return GoBool(r1)
}

func (m *json) ReadInteger(aDictionary ICefDictionaryValue, aKey string, aValue *int32) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(8, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = int32(result2)
	return GoBool(r1)
}

func (m *json) ReadDouble(aDictionary ICefDictionaryValue, aKey string, aValue *float64) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(7, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = float64(result2)
	return GoBool(r1)
}

func (m *json) ReadString(aDictionary ICefDictionaryValue, aKey string, aValue *string) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(10, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = GoStr(result2)
	return GoBool(r1)
}

func (m *json) ReadBinary(aDictionary ICefDictionaryValue, aKey string, aValue *ICefBinaryValue) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(4, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefBinaryValue(result2)
	return GoBool(r1)
}

func (m *json) ReadDictionary(aDictionary ICefDictionaryValue, aKey string, aValue *ICefDictionaryValue) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(6, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefDictionaryValue(result2)
	return GoBool(r1)
}

func (m *json) ReadList(aDictionary ICefDictionaryValue, aKey string, aValue *ICefListValue) bool {
	var result2 uintptr
	r1 := jsonImportAPI().SysCallN(9, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefListValue(result2)
	return GoBool(r1)
}

func (m *json) Parse(jsonString string, options TCefJsonParserOptions) ICefValue {
	var resultCefValue uintptr
	jsonImportAPI().SysCallN(1, PascalStr(jsonString), uintptr(options), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *json) Parse1(json uintptr, jsonsize NativeUInt, options TCefJsonParserOptions) ICefValue {
	var resultCefValue uintptr
	jsonImportAPI().SysCallN(2, uintptr(json), uintptr(jsonsize), uintptr(options), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *json) ParseAndReturnError(jsonString string, options TCefJsonParserOptions, outErrorMsgOut *string) ICefValue {
	var result2 uintptr
	var resultCefValue uintptr
	jsonImportAPI().SysCallN(3, PascalStr(jsonString), uintptr(options), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&resultCefValue)))
	*outErrorMsgOut = GoStr(result2)
	return AsCefValue(resultCefValue)
}

func (m *json) Write(node ICefValue, options TCefJsonWriterOptions) string {
	r1 := jsonImportAPI().SysCallN(14, GetObjectUintptr(node), uintptr(options))
	return GoStr(r1)
}

func (m *json) Write1(node ICefDictionaryValue, options TCefJsonWriterOptions) string {
	r1 := jsonImportAPI().SysCallN(15, GetObjectUintptr(node), uintptr(options))
	return GoStr(r1)
}

func (m *json) Write2(node ICefValue, aRsltStrings *IStringList) bool {
	var result1 uintptr
	r1 := jsonImportAPI().SysCallN(16, GetObjectUintptr(node), uintptr(unsafePointer(&result1)))
	*aRsltStrings = AsStringList(result1)
	return GoBool(r1)
}

func (m *json) Write3(node ICefDictionaryValue, aRsltStrings *IStringList) bool {
	var result1 uintptr
	r1 := jsonImportAPI().SysCallN(17, GetObjectUintptr(node), uintptr(unsafePointer(&result1)))
	*aRsltStrings = AsStringList(result1)
	return GoBool(r1)
}

func (m *json) SaveToFile(node ICefValue, aFileName string) bool {
	r1 := jsonImportAPI().SysCallN(12, GetObjectUintptr(node), PascalStr(aFileName))
	return GoBool(r1)
}

func (m *json) SaveToFile1(node ICefDictionaryValue, aFileName string) bool {
	r1 := jsonImportAPI().SysCallN(13, GetObjectUintptr(node), PascalStr(aFileName))
	return GoBool(r1)
}

func (m *json) LoadFromFile(aFileName string, aRsltNode *ICefValue, options TCefJsonParserOptions) bool {
	var result1 uintptr
	r1 := jsonImportAPI().SysCallN(0, PascalStr(aFileName), uintptr(unsafePointer(&result1)), uintptr(options))
	*aRsltNode = AsCefValue(result1)
	return GoBool(r1)
}

var (
	jsonImport       *imports.Imports = nil
	jsonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFJson_LoadFromFile", 0),
		/*1*/ imports.NewTable("CEFJson_Parse", 0),
		/*2*/ imports.NewTable("CEFJson_Parse1", 0),
		/*3*/ imports.NewTable("CEFJson_ParseAndReturnError", 0),
		/*4*/ imports.NewTable("CEFJson_ReadBinary", 0),
		/*5*/ imports.NewTable("CEFJson_ReadBoolean", 0),
		/*6*/ imports.NewTable("CEFJson_ReadDictionary", 0),
		/*7*/ imports.NewTable("CEFJson_ReadDouble", 0),
		/*8*/ imports.NewTable("CEFJson_ReadInteger", 0),
		/*9*/ imports.NewTable("CEFJson_ReadList", 0),
		/*10*/ imports.NewTable("CEFJson_ReadString", 0),
		/*11*/ imports.NewTable("CEFJson_ReadValue", 0),
		/*12*/ imports.NewTable("CEFJson_SaveToFile", 0),
		/*13*/ imports.NewTable("CEFJson_SaveToFile1", 0),
		/*14*/ imports.NewTable("CEFJson_Write", 0),
		/*15*/ imports.NewTable("CEFJson_Write1", 0),
		/*16*/ imports.NewTable("CEFJson_Write2", 0),
		/*17*/ imports.NewTable("CEFJson_Write3", 0),
	}
)

func jsonImportAPI() *imports.Imports {
	if jsonImport == nil {
		jsonImport = NewDefaultImports()
		jsonImport.SetImportTable(jsonImportTables)
		jsonImportTables = nil
	}
	return jsonImport
}
