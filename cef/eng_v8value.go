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

// ICefv8Value Parent: ICefBaseRefCounted
type ICefv8Value interface {
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
	IsSame(that ICefv8Value) bool                                                                              // function
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
	GetValueByKey(key string) ICefv8Value                                                                      // function
	GetValueByIndex(index int32) ICefv8Value                                                                   // function
	SetValueByKey(key string, value ICefv8Value, attribute TCefV8PropertyAttributes) bool                      // function
	SetValueByIndex(index int32, value ICefv8Value) bool                                                       // function
	SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool     // function
	GetKeys(keys IStrings) int32                                                                               // function
	SetUserData(data ICefv8Value) bool                                                                         // function
	GetUserData() ICefv8Value                                                                                  // function
	GetExternallyAllocatedMemory() int32                                                                       // function
	AdjustExternallyAllocatedMemory(changeInBytes int32) int32                                                 // function
	GetArrayLength() int32                                                                                     // function
	GetArrayBufferReleaseCallback() ICefv8ArrayBufferReleaseCallback                                           // function
	NeuterArrayBuffer() bool                                                                                   // function
	GetFunctionName() string                                                                                   // function
	GetFunctionHandler() ICefv8Handler                                                                         // function
	ExecuteFunction(obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value                                   // function
	ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value // function
	ResolvePromise(arg ICefv8Value) bool                                                                       // function
	RejectPromise(errorMsg string) bool                                                                        // function
}

// TCefv8Value Parent: TCefBaseRefCounted
type TCefv8Value struct {
	TCefBaseRefCounted
}

// V8ValueRef -> ICefv8Value
var V8ValueRef v8Value

// v8Value TCefv8Value Ref
type v8Value uintptr

func (m *v8Value) UnWrap(data uintptr) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(62, uintptr(data), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewUndefined() ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(54, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewNull() ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(49, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewBool(value bool) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(44, PascalBool(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewInt(value int32) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(48, uintptr(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewUInt(value uint32) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(53, uintptr(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewDouble(value float64) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(46, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewDate(value TDateTime) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(45, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewString(str string) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(52, PascalStr(str), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewObject(accessor ICefV8Accessor, interceptor ICefV8Interceptor) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(50, GetObjectUintptr(accessor), GetObjectUintptr(interceptor), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewArray(len int32) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(42, uintptr(len), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewArrayBuffer(buffer uintptr, length NativeUInt, callback ICefv8ArrayBufferReleaseCallback) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(43, uintptr(buffer), uintptr(length), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewFunction(name string, handler ICefv8Handler) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(47, PascalStr(name), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewPromise() ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(51, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) IsValid() bool {
	r1 := v8ValueImportAPI().SysCallN(40, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsUndefined() bool {
	r1 := v8ValueImportAPI().SysCallN(38, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsNull() bool {
	r1 := v8ValueImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsBool() bool {
	r1 := v8ValueImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsInt() bool {
	r1 := v8ValueImportAPI().SysCallN(31, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsUInt() bool {
	r1 := v8ValueImportAPI().SysCallN(37, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsDouble() bool {
	r1 := v8ValueImportAPI().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsDate() bool {
	r1 := v8ValueImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsString() bool {
	r1 := v8ValueImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsObject() bool {
	r1 := v8ValueImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsArray() bool {
	r1 := v8ValueImportAPI().SysCallN(25, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(26, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsFunction() bool {
	r1 := v8ValueImportAPI().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsPromise() bool {
	r1 := v8ValueImportAPI().SysCallN(34, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsSame(that ICefv8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(35, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefv8Value) GetBoolValue() bool {
	r1 := v8ValueImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetIntValue() int32 {
	r1 := v8ValueImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) GetUIntValue() uint32 {
	r1 := v8ValueImportAPI().SysCallN(18, m.Instance())
	return uint32(r1)
}

func (m *TCefv8Value) GetDoubleValue() (resultFloat64 float64) {
	v8ValueImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefv8Value) GetDateValue() (resultDateTime TDateTime) {
	v8ValueImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefv8Value) GetStringValue() string {
	r1 := v8ValueImportAPI().SysCallN(17, m.Instance())
	return GoStr(r1)
}

func (m *TCefv8Value) IsUserCreated() bool {
	r1 := v8ValueImportAPI().SysCallN(39, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) HasException() bool {
	r1 := v8ValueImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetException() ICefV8Exception {
	var resultCefV8Exception uintptr
	v8ValueImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefv8Value) ClearException() bool {
	r1 := v8ValueImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) WillRethrowExceptions() bool {
	r1 := v8ValueImportAPI().SysCallN(63, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) SetRethrowExceptions(rethrow bool) bool {
	r1 := v8ValueImportAPI().SysCallN(57, m.Instance(), PascalBool(rethrow))
	return GoBool(r1)
}

func (m *TCefv8Value) HasValueByKey(key string) bool {
	r1 := v8ValueImportAPI().SysCallN(24, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefv8Value) HasValueByIndex(index int32) bool {
	r1 := v8ValueImportAPI().SysCallN(23, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefv8Value) DeleteValueByKey(key string) bool {
	r1 := v8ValueImportAPI().SysCallN(3, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefv8Value) DeleteValueByIndex(index int32) bool {
	r1 := v8ValueImportAPI().SysCallN(2, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefv8Value) GetValueByKey(key string) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(21, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) GetValueByIndex(index int32) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(20, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) SetValueByKey(key string, value ICefv8Value, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(61, m.Instance(), PascalStr(key), GetObjectUintptr(value), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefv8Value) SetValueByIndex(index int32, value ICefv8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(60, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefv8Value) SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(59, m.Instance(), PascalStr(key), uintptr(settings), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefv8Value) GetKeys(keys IStrings) int32 {
	r1 := v8ValueImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(keys))
	return int32(r1)
}

func (m *TCefv8Value) SetUserData(data ICefv8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(58, m.Instance(), GetObjectUintptr(data))
	return GoBool(r1)
}

func (m *TCefv8Value) GetUserData() ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(19, m.Instance(), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) GetExternallyAllocatedMemory() int32 {
	r1 := v8ValueImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	r1 := v8ValueImportAPI().SysCallN(0, m.Instance(), uintptr(changeInBytes))
	return int32(r1)
}

func (m *TCefv8Value) GetArrayLength() int32 {
	r1 := v8ValueImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) GetArrayBufferReleaseCallback() ICefv8ArrayBufferReleaseCallback {
	var resultCefv8ArrayBufferReleaseCallback uintptr
	v8ValueImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefv8ArrayBufferReleaseCallback)))
	return AsCefv8ArrayBufferReleaseCallback(resultCefv8ArrayBufferReleaseCallback)
}

func (m *TCefv8Value) NeuterArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(41, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetFunctionName() string {
	r1 := v8ValueImportAPI().SysCallN(14, m.Instance())
	return GoStr(r1)
}

func (m *TCefv8Value) GetFunctionHandler() ICefv8Handler {
	var resultCefv8Handler uintptr
	v8ValueImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefv8Handler)))
	return AsCefv8Handler(resultCefv8Handler)
}

func (m *TCefv8Value) ExecuteFunction(obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value {
	var resultCefv8Value uintptr
	v8ValueImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(context), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) ResolvePromise(arg ICefv8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(56, m.Instance(), GetObjectUintptr(arg))
	return GoBool(r1)
}

func (m *TCefv8Value) RejectPromise(errorMsg string) bool {
	r1 := v8ValueImportAPI().SysCallN(55, m.Instance(), PascalStr(errorMsg))
	return GoBool(r1)
}

var (
	v8ValueImport       *imports.Imports = nil
	v8ValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Cefv8Value_AdjustExternallyAllocatedMemory", 0),
		/*1*/ imports.NewTable("Cefv8Value_ClearException", 0),
		/*2*/ imports.NewTable("Cefv8Value_DeleteValueByIndex", 0),
		/*3*/ imports.NewTable("Cefv8Value_DeleteValueByKey", 0),
		/*4*/ imports.NewTable("Cefv8Value_ExecuteFunction", 0),
		/*5*/ imports.NewTable("Cefv8Value_ExecuteFunctionWithContext", 0),
		/*6*/ imports.NewTable("Cefv8Value_GetArrayBufferReleaseCallback", 0),
		/*7*/ imports.NewTable("Cefv8Value_GetArrayLength", 0),
		/*8*/ imports.NewTable("Cefv8Value_GetBoolValue", 0),
		/*9*/ imports.NewTable("Cefv8Value_GetDateValue", 0),
		/*10*/ imports.NewTable("Cefv8Value_GetDoubleValue", 0),
		/*11*/ imports.NewTable("Cefv8Value_GetException", 0),
		/*12*/ imports.NewTable("Cefv8Value_GetExternallyAllocatedMemory", 0),
		/*13*/ imports.NewTable("Cefv8Value_GetFunctionHandler", 0),
		/*14*/ imports.NewTable("Cefv8Value_GetFunctionName", 0),
		/*15*/ imports.NewTable("Cefv8Value_GetIntValue", 0),
		/*16*/ imports.NewTable("Cefv8Value_GetKeys", 0),
		/*17*/ imports.NewTable("Cefv8Value_GetStringValue", 0),
		/*18*/ imports.NewTable("Cefv8Value_GetUIntValue", 0),
		/*19*/ imports.NewTable("Cefv8Value_GetUserData", 0),
		/*20*/ imports.NewTable("Cefv8Value_GetValueByIndex", 0),
		/*21*/ imports.NewTable("Cefv8Value_GetValueByKey", 0),
		/*22*/ imports.NewTable("Cefv8Value_HasException", 0),
		/*23*/ imports.NewTable("Cefv8Value_HasValueByIndex", 0),
		/*24*/ imports.NewTable("Cefv8Value_HasValueByKey", 0),
		/*25*/ imports.NewTable("Cefv8Value_IsArray", 0),
		/*26*/ imports.NewTable("Cefv8Value_IsArrayBuffer", 0),
		/*27*/ imports.NewTable("Cefv8Value_IsBool", 0),
		/*28*/ imports.NewTable("Cefv8Value_IsDate", 0),
		/*29*/ imports.NewTable("Cefv8Value_IsDouble", 0),
		/*30*/ imports.NewTable("Cefv8Value_IsFunction", 0),
		/*31*/ imports.NewTable("Cefv8Value_IsInt", 0),
		/*32*/ imports.NewTable("Cefv8Value_IsNull", 0),
		/*33*/ imports.NewTable("Cefv8Value_IsObject", 0),
		/*34*/ imports.NewTable("Cefv8Value_IsPromise", 0),
		/*35*/ imports.NewTable("Cefv8Value_IsSame", 0),
		/*36*/ imports.NewTable("Cefv8Value_IsString", 0),
		/*37*/ imports.NewTable("Cefv8Value_IsUInt", 0),
		/*38*/ imports.NewTable("Cefv8Value_IsUndefined", 0),
		/*39*/ imports.NewTable("Cefv8Value_IsUserCreated", 0),
		/*40*/ imports.NewTable("Cefv8Value_IsValid", 0),
		/*41*/ imports.NewTable("Cefv8Value_NeuterArrayBuffer", 0),
		/*42*/ imports.NewTable("Cefv8Value_NewArray", 0),
		/*43*/ imports.NewTable("Cefv8Value_NewArrayBuffer", 0),
		/*44*/ imports.NewTable("Cefv8Value_NewBool", 0),
		/*45*/ imports.NewTable("Cefv8Value_NewDate", 0),
		/*46*/ imports.NewTable("Cefv8Value_NewDouble", 0),
		/*47*/ imports.NewTable("Cefv8Value_NewFunction", 0),
		/*48*/ imports.NewTable("Cefv8Value_NewInt", 0),
		/*49*/ imports.NewTable("Cefv8Value_NewNull", 0),
		/*50*/ imports.NewTable("Cefv8Value_NewObject", 0),
		/*51*/ imports.NewTable("Cefv8Value_NewPromise", 0),
		/*52*/ imports.NewTable("Cefv8Value_NewString", 0),
		/*53*/ imports.NewTable("Cefv8Value_NewUInt", 0),
		/*54*/ imports.NewTable("Cefv8Value_NewUndefined", 0),
		/*55*/ imports.NewTable("Cefv8Value_RejectPromise", 0),
		/*56*/ imports.NewTable("Cefv8Value_ResolvePromise", 0),
		/*57*/ imports.NewTable("Cefv8Value_SetRethrowExceptions", 0),
		/*58*/ imports.NewTable("Cefv8Value_SetUserData", 0),
		/*59*/ imports.NewTable("Cefv8Value_SetValueByAccessor", 0),
		/*60*/ imports.NewTable("Cefv8Value_SetValueByIndex", 0),
		/*61*/ imports.NewTable("Cefv8Value_SetValueByKey", 0),
		/*62*/ imports.NewTable("Cefv8Value_UnWrap", 0),
		/*63*/ imports.NewTable("Cefv8Value_WillRethrowExceptions", 0),
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
