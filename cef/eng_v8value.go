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
//
//	Interface representing a V8 value handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8value_t))</a>
type ICefV8Value interface {
	ICefBaseRefCounted
	// ExecuteFunction
	//  Execute the function using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. |object| is the receiver ('this' object) of the function. If |object| is NULL the current context's global object will be used. |arguments| is the list of arguments that will be passed to the function. Returns the function return value on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	ExecuteFunction(obj ICefV8Value, arguments []ICefV8Value) ICefV8Value
	// ExecuteFunctionWithContext
	//  Execute the function using the specified V8 context. |object| is the receiver ('this' object) of the function. If |object| is NULL the specified context's global object will be used. |arguments| is the list of arguments that will be passed to the function. Returns the function return value on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	ExecuteFunctionWithContext(context ICefV8Context, obj ICefV8Value, arguments []ICefV8Value) ICefV8Value
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsUndefined
	//  True if the value type is undefined.
	IsUndefined() bool // function
	// IsNull
	//  True if the value type is null.
	IsNull() bool // function
	// IsBool
	//  True if the value type is bool.
	IsBool() bool // function
	// IsInt
	//  True if the value type is int.
	IsInt() bool // function
	// IsUInt
	//  True if the value type is unsigned int.
	IsUInt() bool // function
	// IsDouble
	//  True if the value type is double.
	IsDouble() bool // function
	// IsDate
	//  True if the value type is Date.
	IsDate() bool // function
	// IsString
	//  True if the value type is string.
	IsString() bool // function
	// IsObject
	//  True if the value type is object.
	IsObject() bool // function
	// IsArray
	//  True if the value type is array.
	IsArray() bool // function
	// IsArrayBuffer
	//  True if the value type is an ArrayBuffer.
	IsArrayBuffer() bool // function
	// IsFunction
	//  True if the value type is function.
	IsFunction() bool // function
	// IsPromise
	//  True if the value type is a Promise.
	IsPromise() bool // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefV8Value) bool // function
	// GetBoolValue
	//  Return a bool value.
	GetBoolValue() bool // function
	// GetIntValue
	//  Return an int value.
	GetIntValue() int32 // function
	// GetUIntValue
	//  Return an unsigned int value.
	GetUIntValue() uint32 // function
	// GetDoubleValue
	//  Return a double value.
	GetDoubleValue() (resultFloat64 float64) // function
	// GetDateValue
	//  Return a Date value.
	GetDateValue() (resultDateTime TDateTime) // function
	// GetStringValue
	//  Return a string value.
	GetStringValue() string // function
	// IsUserCreated
	//  Returns true (1) if this is a user created object.
	IsUserCreated() bool // function
	// HasException
	//  Returns true (1) if the last function call resulted in an exception. This attribute exists only in the scope of the current CEF value object.
	HasException() bool // function
	// GetException
	//  Returns the exception resulting from the last function call. This attribute exists only in the scope of the current CEF value object.
	GetException() ICefV8Exception // function
	// ClearException
	//  Clears the last exception and returns true (1) on success.
	ClearException() bool // function
	// WillRethrowExceptions
	//  Returns true (1) if this object will re-throw future exceptions. This attribute exists only in the scope of the current CEF value object.
	WillRethrowExceptions() bool // function
	// SetRethrowExceptions
	//  Set whether this object will re-throw future exceptions. By default exceptions are not re-thrown. If a exception is re-thrown the current context should not be accessed again until after the exception has been caught and not re-thrown. Returns true (1) on success. This attribute exists only in the scope of the current CEF value object.
	SetRethrowExceptions(rethrow bool) bool // function
	// HasValueByKey
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByKey(key string) bool // function
	// HasValueByIndex
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByIndex(index int32) bool // function
	// DeleteValueByKey
	//  Deletes the value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only and don't-delete values this function will return true (1) even though deletion failed.
	DeleteValueByKey(key string) bool // function
	// DeleteValueByIndex
	//  Deletes the value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly, deletion fails or an exception is thrown. For read-only and don't-delete values this function will return true (1) even though deletion failed.
	DeleteValueByIndex(index int32) bool // function
	// GetValueByKey
	//  Returns the value with the specified identifier on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	GetValueByKey(key string) ICefV8Value // function
	// GetValueByIndex
	//  Returns the value with the specified identifier on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	GetValueByIndex(index int32) ICefV8Value // function
	// SetValueByKey
	//  Associates a value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByKey(key string, value ICefV8Value, attribute TCefV8PropertyAttributes) bool // function
	// SetValueByIndex
	//  Associates a value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByIndex(index int32, value ICefV8Value) bool // function
	// SetValueByAccessor
	//  Registers an identifier and returns true (1) on success. Access to the identifier will be forwarded to the ICefV8Accessor instance passed to cef_v8value_create_object(). Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool // function
	// GetKeys
	//  Read the keys for the object's values into the specified vector. Integer- based keys will also be returned as strings.
	GetKeys(keys IStrings) int32 // function
	// SetUserData
	//  Sets the user data for this object and returns true (1) on success. Returns false (0) if this function is called incorrectly. This function can only be called on user created objects.
	SetUserData(data ICefV8Value) bool // function
	// GetUserData
	//  Returns the user data, if any, assigned to this object.
	GetUserData() ICefV8Value // function
	// GetExternallyAllocatedMemory
	//  Returns the amount of externally allocated memory registered for the object.
	GetExternallyAllocatedMemory() int32 // function
	// AdjustExternallyAllocatedMemory
	//  Adjusts the amount of registered external memory for the object. Used to give V8 an indication of the amount of externally allocated memory that is kept alive by JavaScript objects. V8 uses this information to decide when to perform global garbage collection. Each ICefv8Value tracks the amount of external memory associated with it and automatically decreases the global total by the appropriate amount on its destruction. |change_in_bytes| specifies the number of bytes to adjust by. This function returns the number of bytes associated with the object after the adjustment. This function can only be called on user created objects.
	AdjustExternallyAllocatedMemory(changeInBytes int32) int32 // function
	// GetArrayLength
	//  Returns the number of elements in the array.
	GetArrayLength() int32 // function
	// GetArrayBufferReleaseCallback
	//  Returns the ReleaseCallback object associated with the ArrayBuffer or NULL if the ArrayBuffer was not created with CreateArrayBuffer.
	GetArrayBufferReleaseCallback() ICefV8ArrayBufferReleaseCallback // function
	// NeuterArrayBuffer
	//  Prevent the ArrayBuffer from using it's memory block by setting the length to zero. This operation cannot be undone. If the ArrayBuffer was created with CreateArrayBuffer then ICefv8ArrayBufferReleaseCallback.ReleaseBuffer will be called to release the underlying buffer.
	NeuterArrayBuffer() bool // function
	// GetFunctionName
	//  Returns the function name.
	GetFunctionName() string // function
	// GetFunctionHandler
	//  Returns the function handler or NULL if not a CEF-created function.
	GetFunctionHandler() ICefV8Handler // function
	// ResolvePromise
	//  Resolve the Promise using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. |arg| is the argument passed to the resolved promise. Returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown.
	ResolvePromise(arg ICefV8Value) bool // function
	// RejectPromise
	//  Reject the Promise using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. Returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown.
	RejectPromise(errorMsg string) bool // function
}

// TCefV8Value Parent: TCefBaseRefCounted
//
//	Interface representing a V8 value handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8value_t))</a>
type TCefV8Value struct {
	TCefBaseRefCounted
}

// V8ValueRef -> ICefV8Value
var V8ValueRef v8Value

// v8Value TCefV8Value Ref
type v8Value uintptr

func (m *v8Value) UnWrap(data uintptr) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(60, uintptr(data), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewUndefined() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(52, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewNull() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(47, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewBool(value bool) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(42, PascalBool(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewInt(value int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(46, uintptr(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewUInt(value uint32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(51, uintptr(value), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewDouble(value float64) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(44, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewDate(value TDateTime) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(43, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewString(str string) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(50, PascalStr(str), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewObject(accessor ICefV8Accessor, interceptor ICefV8Interceptor) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(48, GetObjectUintptr(accessor), GetObjectUintptr(interceptor), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewArray(len int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(40, uintptr(len), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewArrayBuffer(buffer uintptr, length NativeUInt, callback ICefV8ArrayBufferReleaseCallback) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(41, uintptr(buffer), uintptr(length), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewFunction(name string, handler ICefV8Handler) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(45, PascalStr(name), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *v8Value) NewPromise() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(49, uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) IsValid() bool {
	r1 := v8ValueImportAPI().SysCallN(38, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsUndefined() bool {
	r1 := v8ValueImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsNull() bool {
	r1 := v8ValueImportAPI().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsBool() bool {
	r1 := v8ValueImportAPI().SysCallN(25, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsInt() bool {
	r1 := v8ValueImportAPI().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsUInt() bool {
	r1 := v8ValueImportAPI().SysCallN(35, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsDouble() bool {
	r1 := v8ValueImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsDate() bool {
	r1 := v8ValueImportAPI().SysCallN(26, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsString() bool {
	r1 := v8ValueImportAPI().SysCallN(34, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsObject() bool {
	r1 := v8ValueImportAPI().SysCallN(31, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsArray() bool {
	r1 := v8ValueImportAPI().SysCallN(23, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(24, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsFunction() bool {
	r1 := v8ValueImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsPromise() bool {
	r1 := v8ValueImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) IsSame(that ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(33, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefV8Value) GetBoolValue() bool {
	r1 := v8ValueImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetIntValue() int32 {
	r1 := v8ValueImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) GetUIntValue() uint32 {
	r1 := v8ValueImportAPI().SysCallN(16, m.Instance())
	return uint32(r1)
}

func (m *TCefV8Value) GetDoubleValue() (resultFloat64 float64) {
	v8ValueImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefV8Value) GetDateValue() (resultDateTime TDateTime) {
	v8ValueImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefV8Value) GetStringValue() string {
	r1 := v8ValueImportAPI().SysCallN(15, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Value) IsUserCreated() bool {
	r1 := v8ValueImportAPI().SysCallN(37, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) HasException() bool {
	r1 := v8ValueImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetException() ICefV8Exception {
	var resultCefV8Exception uintptr
	v8ValueImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefV8Value) ClearException() bool {
	r1 := v8ValueImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) WillRethrowExceptions() bool {
	r1 := v8ValueImportAPI().SysCallN(61, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) SetRethrowExceptions(rethrow bool) bool {
	r1 := v8ValueImportAPI().SysCallN(55, m.Instance(), PascalBool(rethrow))
	return GoBool(r1)
}

func (m *TCefV8Value) HasValueByKey(key string) bool {
	r1 := v8ValueImportAPI().SysCallN(22, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefV8Value) HasValueByIndex(index int32) bool {
	r1 := v8ValueImportAPI().SysCallN(21, m.Instance(), uintptr(index))
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
	v8ValueImportAPI().SysCallN(19, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) GetValueByIndex(index int32) ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(18, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) SetValueByKey(key string, value ICefV8Value, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(59, m.Instance(), PascalStr(key), GetObjectUintptr(value), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefV8Value) SetValueByIndex(index int32, value ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(58, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefV8Value) SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool {
	r1 := v8ValueImportAPI().SysCallN(57, m.Instance(), PascalStr(key), uintptr(settings), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefV8Value) GetKeys(keys IStrings) int32 {
	r1 := v8ValueImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(keys))
	return int32(r1)
}

func (m *TCefV8Value) SetUserData(data ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(56, m.Instance(), GetObjectUintptr(data))
	return GoBool(r1)
}

func (m *TCefV8Value) GetUserData() ICefV8Value {
	var resultCefV8Value uintptr
	v8ValueImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Value) GetExternallyAllocatedMemory() int32 {
	r1 := v8ValueImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	r1 := v8ValueImportAPI().SysCallN(0, m.Instance(), uintptr(changeInBytes))
	return int32(r1)
}

func (m *TCefV8Value) GetArrayLength() int32 {
	r1 := v8ValueImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TCefV8Value) GetArrayBufferReleaseCallback() ICefV8ArrayBufferReleaseCallback {
	var resultCefV8ArrayBufferReleaseCallback uintptr
	v8ValueImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefV8ArrayBufferReleaseCallback)))
	return AsCefV8ArrayBufferReleaseCallback(resultCefV8ArrayBufferReleaseCallback)
}

func (m *TCefV8Value) NeuterArrayBuffer() bool {
	r1 := v8ValueImportAPI().SysCallN(39, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Value) GetFunctionName() string {
	r1 := v8ValueImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Value) GetFunctionHandler() ICefV8Handler {
	var resultCefV8Handler uintptr
	v8ValueImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultCefV8Handler)))
	return AsCefV8Handler(resultCefV8Handler)
}

func (m *TCefV8Value) ResolvePromise(arg ICefV8Value) bool {
	r1 := v8ValueImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(arg))
	return GoBool(r1)
}

func (m *TCefV8Value) RejectPromise(errorMsg string) bool {
	r1 := v8ValueImportAPI().SysCallN(53, m.Instance(), PascalStr(errorMsg))
	return GoBool(r1)
}

var (
	v8ValueImport       *imports.Imports = nil
	v8ValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8Value_AdjustExternallyAllocatedMemory", 0),
		/*1*/ imports.NewTable("CefV8Value_ClearException", 0),
		/*2*/ imports.NewTable("CefV8Value_DeleteValueByIndex", 0),
		/*3*/ imports.NewTable("CefV8Value_DeleteValueByKey", 0),
		/*4*/ imports.NewTable("CefV8Value_GetArrayBufferReleaseCallback", 0),
		/*5*/ imports.NewTable("CefV8Value_GetArrayLength", 0),
		/*6*/ imports.NewTable("CefV8Value_GetBoolValue", 0),
		/*7*/ imports.NewTable("CefV8Value_GetDateValue", 0),
		/*8*/ imports.NewTable("CefV8Value_GetDoubleValue", 0),
		/*9*/ imports.NewTable("CefV8Value_GetException", 0),
		/*10*/ imports.NewTable("CefV8Value_GetExternallyAllocatedMemory", 0),
		/*11*/ imports.NewTable("CefV8Value_GetFunctionHandler", 0),
		/*12*/ imports.NewTable("CefV8Value_GetFunctionName", 0),
		/*13*/ imports.NewTable("CefV8Value_GetIntValue", 0),
		/*14*/ imports.NewTable("CefV8Value_GetKeys", 0),
		/*15*/ imports.NewTable("CefV8Value_GetStringValue", 0),
		/*16*/ imports.NewTable("CefV8Value_GetUIntValue", 0),
		/*17*/ imports.NewTable("CefV8Value_GetUserData", 0),
		/*18*/ imports.NewTable("CefV8Value_GetValueByIndex", 0),
		/*19*/ imports.NewTable("CefV8Value_GetValueByKey", 0),
		/*20*/ imports.NewTable("CefV8Value_HasException", 0),
		/*21*/ imports.NewTable("CefV8Value_HasValueByIndex", 0),
		/*22*/ imports.NewTable("CefV8Value_HasValueByKey", 0),
		/*23*/ imports.NewTable("CefV8Value_IsArray", 0),
		/*24*/ imports.NewTable("CefV8Value_IsArrayBuffer", 0),
		/*25*/ imports.NewTable("CefV8Value_IsBool", 0),
		/*26*/ imports.NewTable("CefV8Value_IsDate", 0),
		/*27*/ imports.NewTable("CefV8Value_IsDouble", 0),
		/*28*/ imports.NewTable("CefV8Value_IsFunction", 0),
		/*29*/ imports.NewTable("CefV8Value_IsInt", 0),
		/*30*/ imports.NewTable("CefV8Value_IsNull", 0),
		/*31*/ imports.NewTable("CefV8Value_IsObject", 0),
		/*32*/ imports.NewTable("CefV8Value_IsPromise", 0),
		/*33*/ imports.NewTable("CefV8Value_IsSame", 0),
		/*34*/ imports.NewTable("CefV8Value_IsString", 0),
		/*35*/ imports.NewTable("CefV8Value_IsUInt", 0),
		/*36*/ imports.NewTable("CefV8Value_IsUndefined", 0),
		/*37*/ imports.NewTable("CefV8Value_IsUserCreated", 0),
		/*38*/ imports.NewTable("CefV8Value_IsValid", 0),
		/*39*/ imports.NewTable("CefV8Value_NeuterArrayBuffer", 0),
		/*40*/ imports.NewTable("CefV8Value_NewArray", 0),
		/*41*/ imports.NewTable("CefV8Value_NewArrayBuffer", 0),
		/*42*/ imports.NewTable("CefV8Value_NewBool", 0),
		/*43*/ imports.NewTable("CefV8Value_NewDate", 0),
		/*44*/ imports.NewTable("CefV8Value_NewDouble", 0),
		/*45*/ imports.NewTable("CefV8Value_NewFunction", 0),
		/*46*/ imports.NewTable("CefV8Value_NewInt", 0),
		/*47*/ imports.NewTable("CefV8Value_NewNull", 0),
		/*48*/ imports.NewTable("CefV8Value_NewObject", 0),
		/*49*/ imports.NewTable("CefV8Value_NewPromise", 0),
		/*50*/ imports.NewTable("CefV8Value_NewString", 0),
		/*51*/ imports.NewTable("CefV8Value_NewUInt", 0),
		/*52*/ imports.NewTable("CefV8Value_NewUndefined", 0),
		/*53*/ imports.NewTable("CefV8Value_RejectPromise", 0),
		/*54*/ imports.NewTable("CefV8Value_ResolvePromise", 0),
		/*55*/ imports.NewTable("CefV8Value_SetRethrowExceptions", 0),
		/*56*/ imports.NewTable("CefV8Value_SetUserData", 0),
		/*57*/ imports.NewTable("CefV8Value_SetValueByAccessor", 0),
		/*58*/ imports.NewTable("CefV8Value_SetValueByIndex", 0),
		/*59*/ imports.NewTable("CefV8Value_SetValueByKey", 0),
		/*60*/ imports.NewTable("CefV8Value_UnWrap", 0),
		/*61*/ imports.NewTable("CefV8Value_WillRethrowExceptions", 0),
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
