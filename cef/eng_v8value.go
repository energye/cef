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

// ICefV8Value Parent: ICefBaseRefCounted
type ICefV8Value interface {
	ICefBaseRefCounted
	IsValid() bool                                                                                             // function
	IsUndefined() bool                                                                                         // function
	IsNull() bool                                                                                              // function
	IsBool() bool                                                                                              // function
	IsInt() bool                                                                                               // function
	IsUInt() bool                                                                                              // function
	IsDouble() bool                                                                                            // function
	IsDate() bool                                                                                              // function
	IsString() bool                                                                                            // function
	IsObject() bool                                                                                            // function
	IsArray() bool                                                                                             // function
	IsArrayBuffer() bool                                                                                       // function
	IsFunction() bool                                                                                          // function
	IsPromise() bool                                                                                           // function
	IsSame(that ICefV8Value) bool                                                                              // function
	GetBoolValue() bool                                                                                        // function
	GetIntValue() int32                                                                                        // function
	GetUIntValue() uint32                                                                                      // function
	GetDoubleValue() (resultFloat64 float64)                                                                   // function
	GetDateValue() (resultDateTime TDateTime)                                                                  // function
	GetStringValue() string                                                                                    // function
	IsUserCreated() bool                                                                                       // function
	HasException() bool                                                                                        // function
	GetException() ICefV8Exception                                                                             // function
	ClearException() bool                                                                                      // function
	WillRethrowExceptions() bool                                                                               // function
	SetRethrowExceptions(rethrow bool) bool                                                                    // function
	HasValueByKey(key string) bool                                                                             // function
	HasValueByIndex(index int32) bool                                                                          // function
	DeleteValueByKey(key string) bool                                                                          // function
	DeleteValueByIndex(index int32) bool                                                                       // function
	GetValueByKey(key string) ICefV8Value                                                                      // function
	GetValueByIndex(index int32) ICefV8Value                                                                   // function
	SetValueByKey(key string, value ICefV8Value, attribute TCefV8PropertyAttributes) bool                      // function
	SetValueByIndex(index int32, value ICefV8Value) bool                                                       // function
	SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool     // function
	GetKeys(keys IStrings) int32                                                                               // function
	SetUserData(data ICefV8Value) bool                                                                         // function
	GetUserData() ICefV8Value                                                                                  // function
	GetExternallyAllocatedMemory() int32                                                                       // function
	AdjustExternallyAllocatedMemory(changeInBytes int32) int32                                                 // function
	GetArrayLength() int32                                                                                     // function
	GetArrayBufferReleaseCallback() ICefV8ArrayBufferReleaseCallback                                           // function
	NeuterArrayBuffer() bool                                                                                   // function
	GetFunctionName() string                                                                                   // function
	GetFunctionHandler() ICefV8Handler                                                                         // function
	ExecuteFunction(obj ICefV8Value, arguments ICefV8ValueArray) ICefV8Value                                   // function
	ExecuteFunctionWithContext(context ICefV8Context, obj ICefV8Value, arguments ICefV8ValueArray) ICefV8Value // function
	ResolvePromise(arg ICefV8Value) bool                                                                       // function
	RejectPromise(errorMsg string) bool                                                                        // function
}

// TCefV8Value Parent: TCefBaseRefCounted
type TCefV8Value struct {
	TCefBaseRefCounted
}

// V8ValueRef -> ICefV8Value
var V8ValueRef v8Value

// v8Value TCefV8Value Ref
type v8Value uintptr

func (m *v8Value) UnWrap(data uintptr) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(62, uintptr(data), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewUndefined() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(54, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewNull() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(49, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewBool(value bool) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(44, PascalBool(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewInt(value int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(48, uintptr(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewUInt(value uint32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(53, uintptr(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewDouble(value float64) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(46, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewDate(value TDateTime) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(45, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewString(str string) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(52, PascalStr(str), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewObject(accessor ICefV8Accessor, interceptor ICefV8Interceptor) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(50, GetObjectUintptr(accessor), GetObjectUintptr(interceptor), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewArray(len int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(42, uintptr(len), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewArrayBuffer(buffer uintptr, length NativeUInt, callback ICefV8ArrayBufferReleaseCallback) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(43, uintptr(buffer), uintptr(length), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewFunction(name string, handler ICefV8Handler) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(47, PascalStr(name), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewPromise() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(51, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) IsValid() bool {
	r1 := v8ValueImportAPI().SysCallN(40, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsUndefined() bool {
	r1 := v8ValueImportAPI().SysCallN(38, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsNull() bool {
	r1 := v8ValueImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsBool() bool {
	r1 := v8ValueImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsInt() bool {
	r1 := v8ValueImportAPI().SysCallN(31, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsUInt() bool {
	r1 := v8ValueImportAPI().SysCallN(37, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsDouble() bool {
	r1 := v8ValueImportAPI().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsDate() bool {
	r1 := v8ValueImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsString() bool {
	r1 := v8ValueImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsObject() bool {
	r1 := v8ValueImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsArray() bool {
	r1 := v8ValueImportAPI().SysCallN(25, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(26, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsFunction() bool {
	r1 := v8ValueImportAPI().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsPromise() bool {
	r1 := v8ValueImportAPI().SysCallN(34, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsSame(that ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(35, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefV8Value) GetBoolValue() bool {
	r1 := v8ValueImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetIntValue() int32 {
	r1 := v8ValueImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) GetUIntValue() uint32 {
	r1 := v8ValueImportAPI().SysCallN(18, m.Instance())
	return uint32(r1)
}

func (m *TCefV8Value) GetDoubleValue() (resultFloat64 float64) {
	v8ValueImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefV8Value) GetDateValue() (resultDateTime TDateTime) {
	v8ValueImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefV8Value) GetStringValue() string {
	r1 := v8ValueImportAPI().SysCallN(17, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Value) IsUserCreated() bool {
	r1 := v8ValueImportAPI().SysCallN(39, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) HasException() bool {
	r1 := v8ValueImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetException() ICefV8Exception {
	var resultCefV8Exception uintptr
	v8ValueImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefV8Value) ClearException() bool {
	r1 := v8ValueImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) WillRethrowExceptions() bool {
	r1 := v8ValueImportAPI().SysCallN(63, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) SetRethrowExceptions(rethrow bool) bool {
	r1 := v8ValueImportAPI().SysCallN(57, m.Instance(), PascalBool(rethrow))
	return GoBool(r1)
}

func (m *TCefV8Value) HasValueByKey(key string) bool {
	r1 := v8ValueImportAPI().SysCallN(24, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefV8Value) HasValueByIndex(index int32) bool {
	r1 := v8ValueImportAPI().SysCallN(23, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefV8Value) DeleteValueByKey(key string) bool {
	r1 := v8ValueImportAPI().SysCallN(3, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefV8Value) DeleteValueByIndex(index int32) bool {
	r1 := v8ValueImportAPI().SysCallN(2, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefV8Value) GetValueByKey(key string) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(21, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) GetValueByIndex(index int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(20, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) SetValueByKey(key string, value ICefV8Value, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(61, m.Instance(), PascalStr(key), GetObjectUintptr(value), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefV8Value) SetValueByIndex(index int32, value ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(60, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefV8Value) SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(59, m.Instance(), PascalStr(key), uintptr(settings), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefV8Value) GetKeys(keys IStrings) int32 {
	r1 := v8ValueImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(keys))
	return int32(r1)
}

func (m *TCefV8Value) SetUserData(data ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(58, m.Instance(), GetObjectUintptr(data))
	return GoBool(r1)
}

func (m *TCefV8Value) GetUserData() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(19, m.Instance(), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) GetExternallyAllocatedMemory() int32 {
	r1 := v8ValueImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	r1 := v8ValueImportAPI().SysCallN(0, m.Instance(), uintptr(changeInBytes))
	return int32(r1)
}

func (m *TCefV8Value) GetArrayLength() int32 {
	r1 := v8ValueImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) GetArrayBufferReleaseCallback() ICefV8ArrayBufferReleaseCallback {
	var resultCefV8ArrayBufferReleaseCallback uintptr
	v8ValueImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefV8ArrayBufferReleaseCallback)))
	return AsCefV8ArrayBufferReleaseCallback(resultCefV8ArrayBufferReleaseCallback)
}

func (m *TCefV8Value) NeuterArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(41, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetFunctionName() string {
	r1 := v8ValueImportAPI().SysCallN(14, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Value) GetFunctionHandler() ICefV8Handler {
	var resultCefV8Handler uintptr
	v8ValueImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefV8Handler)))
	return AsCefV8Handler(resultCefV8Handler)
}

func (m *TCefV8Value) ExecuteFunction(obj ICefV8Value, arguments ICefV8ValueArray) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) ExecuteFunctionWithContext(context ICefV8Context, obj ICefV8Value, arguments ICefV8ValueArray) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(context), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) ResolvePromise(arg ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(56, m.Instance(), GetObjectUintptr(arg))
	return GoBool(r1)
}

func (m *TCefV8Value) RejectPromise(errorMsg string) bool {
	r1 := v8ValueImportAPI().SysCallN(55, m.Instance(), PascalStr(errorMsg))
	return GoBool(r1)
}

var (
	v8ValueImport       *imports.Imports = nil
	v8ValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8Value_AdjustExternallyAllocatedMemory", 0),
		/*1*/ imports.NewTable("CefV8Value_ClearException", 0),
		/*2*/ imports.NewTable("CefV8Value_DeleteValueByIndex", 0),
		/*3*/ imports.NewTable("CefV8Value_DeleteValueByKey", 0),
		/*4*/ imports.NewTable("CefV8Value_ExecuteFunction", 0),
		/*5*/ imports.NewTable("CefV8Value_ExecuteFunctionWithContext", 0),
		/*6*/ imports.NewTable("CefV8Value_GetArrayBufferReleaseCallback", 0),
		/*7*/ imports.NewTable("CefV8Value_GetArrayLength", 0),
		/*8*/ imports.NewTable("CefV8Value_GetBoolValue", 0),
		/*9*/ imports.NewTable("CefV8Value_GetDateValue", 0),
		/*10*/ imports.NewTable("CefV8Value_GetDoubleValue", 0),
		/*11*/ imports.NewTable("CefV8Value_GetException", 0),
		/*12*/ imports.NewTable("CefV8Value_GetExternallyAllocatedMemory", 0),
		/*13*/ imports.NewTable("CefV8Value_GetFunctionHandler", 0),
		/*14*/ imports.NewTable("CefV8Value_GetFunctionName", 0),
		/*15*/ imports.NewTable("CefV8Value_GetIntValue", 0),
		/*16*/ imports.NewTable("CefV8Value_GetKeys", 0),
		/*17*/ imports.NewTable("CefV8Value_GetStringValue", 0),
		/*18*/ imports.NewTable("CefV8Value_GetUIntValue", 0),
		/*19*/ imports.NewTable("CefV8Value_GetUserData", 0),
		/*20*/ imports.NewTable("CefV8Value_GetValueByIndex", 0),
		/*21*/ imports.NewTable("CefV8Value_GetValueByKey", 0),
		/*22*/ imports.NewTable("CefV8Value_HasException", 0),
		/*23*/ imports.NewTable("CefV8Value_HasValueByIndex", 0),
		/*24*/ imports.NewTable("CefV8Value_HasValueByKey", 0),
		/*25*/ imports.NewTable("CefV8Value_IsArray", 0),
		/*26*/ imports.NewTable("CefV8Value_IsArrayBuffer", 0),
		/*27*/ imports.NewTable("CefV8Value_IsBool", 0),
		/*28*/ imports.NewTable("CefV8Value_IsDate", 0),
		/*29*/ imports.NewTable("CefV8Value_IsDouble", 0),
		/*30*/ imports.NewTable("CefV8Value_IsFunction", 0),
		/*31*/ imports.NewTable("CefV8Value_IsInt", 0),
		/*32*/ imports.NewTable("CefV8Value_IsNull", 0),
		/*33*/ imports.NewTable("CefV8Value_IsObject", 0),
		/*34*/ imports.NewTable("CefV8Value_IsPromise", 0),
		/*35*/ imports.NewTable("CefV8Value_IsSame", 0),
		/*36*/ imports.NewTable("CefV8Value_IsString", 0),
		/*37*/ imports.NewTable("CefV8Value_IsUInt", 0),
		/*38*/ imports.NewTable("CefV8Value_IsUndefined", 0),
		/*39*/ imports.NewTable("CefV8Value_IsUserCreated", 0),
		/*40*/ imports.NewTable("CefV8Value_IsValid", 0),
		/*41*/ imports.NewTable("CefV8Value_NeuterArrayBuffer", 0),
		/*42*/ imports.NewTable("CefV8Value_NewArray", 0),
		/*43*/ imports.NewTable("CefV8Value_NewArrayBuffer", 0),
		/*44*/ imports.NewTable("CefV8Value_NewBool", 0),
		/*45*/ imports.NewTable("CefV8Value_NewDate", 0),
		/*46*/ imports.NewTable("CefV8Value_NewDouble", 0),
		/*47*/ imports.NewTable("CefV8Value_NewFunction", 0),
		/*48*/ imports.NewTable("CefV8Value_NewInt", 0),
		/*49*/ imports.NewTable("CefV8Value_NewNull", 0),
		/*50*/ imports.NewTable("CefV8Value_NewObject", 0),
		/*51*/ imports.NewTable("CefV8Value_NewPromise", 0),
		/*52*/ imports.NewTable("CefV8Value_NewString", 0),
		/*53*/ imports.NewTable("CefV8Value_NewUInt", 0),
		/*54*/ imports.NewTable("CefV8Value_NewUndefined", 0),
		/*55*/ imports.NewTable("CefV8Value_RejectPromise", 0),
		/*56*/ imports.NewTable("CefV8Value_ResolvePromise", 0),
		/*57*/ imports.NewTable("CefV8Value_SetRethrowExceptions", 0),
		/*58*/ imports.NewTable("CefV8Value_SetUserData", 0),
		/*59*/ imports.NewTable("CefV8Value_SetValueByAccessor", 0),
		/*60*/ imports.NewTable("CefV8Value_SetValueByIndex", 0),
		/*61*/ imports.NewTable("CefV8Value_SetValueByKey", 0),
		/*62*/ imports.NewTable("CefV8Value_UnWrap", 0),
		/*63*/ imports.NewTable("CefV8Value_WillRethrowExceptions", 0),
	}
)

func v8ValueImportAPI() *imports.Imports {
	if v8ValueImport == nil {
		v8ValueImport = NewDefaultImports()
		v8ValueImport.SetImportTable(v8ValueImportTables)
		v8ValueImportTables = nil
	}
	return v8ValueImport
}
