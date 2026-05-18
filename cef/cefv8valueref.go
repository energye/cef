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

	cefTypes "github.com/energye/cef/types"
)

// ICefv8Value Parent: ICefBaseRefCounted
type ICefv8Value interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed
	//  on the current thread. Do not call any other functions if this function
	//  returns false (0).
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
	//  Returns true (1) if this object is pointing to the same handle as |that|
	//  object.
	IsSame(that ICefv8Value) bool // function
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
	GetDoubleValue() float64 // function
	// GetDateValue
	//  Return a Date value.
	GetDateValue() types.TDateTime // function
	// GetStringValue
	//  Return a string value.
	GetStringValue() string // function
	// IsUserCreated
	//  Returns true (1) if this is a user created object.
	IsUserCreated() bool // function
	// HasException
	//  Returns true (1) if the last function call resulted in an exception. This
	//  attribute exists only in the scope of the current CEF value object.
	HasException() bool // function
	// GetException
	//  Returns the exception resulting from the last function call. This
	//  attribute exists only in the scope of the current CEF value object.
	GetException() ICefV8Exception // function
	// ClearException
	//  Clears the last exception and returns true (1) on success.
	ClearException() bool // function
	// WillRethrowExceptions
	//  Returns true (1) if this object will re-throw future exceptions. This
	//  attribute exists only in the scope of the current CEF value object.
	WillRethrowExceptions() bool // function
	// SetRethrowExceptions
	//  Set whether this object will re-throw future exceptions. By default
	//  exceptions are not re-thrown. If a exception is re-thrown the current
	//  context should not be accessed again until after the exception has been
	//  caught and not re-thrown. Returns true (1) on success. This attribute
	//  exists only in the scope of the current CEF value object.
	SetRethrowExceptions(rethrow bool) bool // function
	// HasValueByKey
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByKey(key string) bool // function
	// HasValueByIndex
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByIndex(index int32) bool // function
	// DeleteValueByKey
	//  Deletes the value with the specified identifier and returns true (1) on
	//  success. Returns false (0) if this function is called incorrectly or an
	//  exception is thrown. For read-only and don't-delete values this function
	//  will return true (1) even though deletion failed.
	DeleteValueByKey(key string) bool // function
	// DeleteValueByIndex
	//  Deletes the value with the specified identifier and returns true (1) on
	//  success. Returns false (0) if this function is called incorrectly,
	//  deletion fails or an exception is thrown. For read-only and don't-delete
	//  values this function will return true (1) even though deletion failed.
	DeleteValueByIndex(index int32) bool // function
	// GetValueByKey
	//  Returns the value with the specified identifier on success. Returns NULL
	//  if this function is called incorrectly or an exception is thrown.
	GetValueByKey(key string) ICefv8Value // function
	// GetValueByIndex
	//  Returns the value with the specified identifier on success. Returns NULL
	//  if this function is called incorrectly or an exception is thrown.
	GetValueByIndex(index int32) ICefv8Value // function
	// SetValueByKey
	//  Associates a value with the specified identifier and returns true (1) on
	//  success. Returns false (0) if this function is called incorrectly or an
	//  exception is thrown. For read-only values this function will return true
	//  (1) even though assignment failed.
	SetValueByKey(key string, value ICefv8Value, attribute cefTypes.TCefV8PropertyAttributes) bool // function
	// SetValueByIndex
	//  Associates a value with the specified identifier and returns true (1) on
	//  success. Returns false (0) if this function is called incorrectly or an
	//  exception is thrown. For read-only values this function will return true
	//  (1) even though assignment failed.
	SetValueByIndex(index int32, value ICefv8Value) bool // function
	// SetValueByAccessor
	//  Registers an identifier and returns true (1) on success. Access to the
	//  identifier will be forwarded to the ICefV8Accessor instance passed to
	//  cef_v8value_create_object(). Returns false (0) if this
	//  function is called incorrectly or an exception is thrown. For read-only
	//  values this function will return true (1) even though assignment failed.
	SetValueByAccessor(key string, attribute cefTypes.TCefV8PropertyAttributes) bool // function
	// GetKeys
	//  Read the keys for the object's values into the specified vector. Integer-
	//  based keys will also be returned as strings.
	GetKeys(keys lcl.IStrings) int32 // function
	// SetUserData
	//  Sets the user data for this object and returns true (1) on success.
	//  Returns false (0) if this function is called incorrectly. This function
	//  can only be called on user created objects.
	SetUserData(data ICefv8Value) bool // function
	// GetUserData
	//  Returns the user data, if any, assigned to this object.
	GetUserData() ICefv8Value // function
	// GetExternallyAllocatedMemory
	//  Returns the amount of externally allocated memory registered for the
	//  object.
	GetExternallyAllocatedMemory() int32 // function
	// AdjustExternallyAllocatedMemory
	//  Adjusts the amount of registered external memory for the object. Used to
	//  give V8 an indication of the amount of externally allocated memory that is
	//  kept alive by JavaScript objects. V8 uses this information to decide when
	//  to perform global garbage collection. Each ICefv8Value tracks the amount
	//  of external memory associated with it and automatically decreases the
	//  global total by the appropriate amount on its destruction.
	//  |change_in_bytes| specifies the number of bytes to adjust by. This
	//  function returns the number of bytes associated with the object after the
	//  adjustment. This function can only be called on user created objects.
	AdjustExternallyAllocatedMemory(changeInBytes int32) int32 // function
	// GetArrayLength
	//  Returns the number of elements in the array.
	GetArrayLength() int32 // function
	// GetArrayBufferReleaseCallback
	//  Returns the ReleaseCallback object associated with the ArrayBuffer or NULL
	//  if the ArrayBuffer was not created with CreateArrayBuffer.
	GetArrayBufferReleaseCallback() IEngV8ArrayBufferReleaseCallback // function
	// NeuterArrayBuffer
	//  Prevent the ArrayBuffer from using it's memory block by setting the length
	//  to zero. This operation cannot be undone. If the ArrayBuffer was created
	//  with CreateArrayBuffer then
	//  ICefv8ArrayBufferReleaseCallback.ReleaseBuffer will be called to
	//  release the underlying buffer.
	NeuterArrayBuffer() bool // function
	// GetArrayBufferByteLength
	//  Returns the length (in bytes) of the ArrayBuffer.
	GetArrayBufferByteLength() cefTypes.NativeUInt // function
	// GetArrayBufferData
	//  Returns a pointer to the beginning of the memory block for this
	//  ArrayBuffer backing store. The returned pointer is valid as long as the
	//  ICefv8value is alive.
	GetArrayBufferData() uintptr // function
	// GetFunctionName
	//  Returns the function name.
	GetFunctionName() string // function
	// GetFunctionHandler
	//  Returns the function handler or NULL if not a CEF-created function.
	GetFunctionHandler() IEngV8Handler // function
	// ExecuteFunction
	//  Execute the function using the current V8 context. This function should
	//  only be called from within the scope of a ICefv8Handler or
	//  ICefV8Accessor callback, or in combination with calling enter() and
	//  exit() on a stored ICefv8Context reference. |object| is the receiver
	//  ('this' object) of the function. If |object| is NULL the current context's
	//  global object will be used. |arguments| is the list of arguments that will
	//  be passed to the function. Returns the function return value on success.
	//  Returns NULL if this function is called incorrectly or an exception is
	//  thrown.
	ExecuteFunction(obj ICefv8Value, arguments ICefv8ValueArray) ICefv8Value // function
	// ExecuteFunctionWithContext
	//  Execute the function using the specified V8 context. |object| is the
	//  receiver ('this' object) of the function. If |object| is NULL the
	//  specified context's global object will be used. |arguments| is the list of
	//  arguments that will be passed to the function. Returns the function return
	//  value on success. Returns NULL if this function is called incorrectly or
	//  an exception is thrown.
	ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefv8ValueArray) ICefv8Value // function
	// ResolvePromise
	//  Resolve the Promise using the current V8 context. This function should
	//  only be called from within the scope of a ICefv8Handler or
	//  ICefV8Accessor callback, or in combination with calling enter() and
	//  exit() on a stored ICefv8Context reference. |arg| is the argument passed
	//  to the resolved promise. Returns true (1) on success. Returns false (0) if
	//  this function is called incorrectly or an exception is thrown.
	ResolvePromise(arg ICefv8Value) bool // function
	// RejectPromise
	//  Reject the Promise using the current V8 context. This function should only
	//  be called from within the scope of a ICefv8Handler or ICefV8Accessor
	//  callback, or in combination with calling enter() and exit() on a stored
	//  ICefv8Context reference. Returns true (1) on success. Returns false (0)
	//  if this function is called incorrectly or an exception is thrown.
	RejectPromise(errorMsg string) bool // function
}

// ICefv8ValueRef Parent: ICefv8Value ICefBaseRefCountedRef
type ICefv8ValueRef interface {
	ICefv8Value
	ICefBaseRefCountedRef
	AsIntfV8Value() uintptr
}

type TCefv8ValueRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefv8ValueRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsUndefined() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsNull() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsBool() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsInt() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsUInt() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsDouble() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsDate() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsString() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsObject() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsArray() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsArrayBuffer() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsFunction() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsPromise() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) IsSame(that ICefv8Value) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetBoolValue() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetIntValue() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(17, m.Instance())
	return int32(r)
}

func (m *TCefv8ValueRef) GetUIntValue() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(18, m.Instance())
	return uint32(r)
}

func (m *TCefv8ValueRef) GetDoubleValue() (result float64) {
	if !m.IsValid() {
		return
	}
	cefv8ValueRefAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefv8ValueRef) GetDateValue() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefv8ValueRefAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefv8ValueRef) GetStringValue() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefv8ValueRefAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefv8ValueRef) IsUserCreated() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(22, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) HasException() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetException() (result ICefV8Exception) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8ExceptionRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) ClearException() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(25, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) WillRethrowExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(26, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) SetRethrowExceptions(rethrow bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(27, m.Instance(), api.PasBool(rethrow))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) HasValueByKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(28, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) HasValueByIndex(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(29, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) DeleteValueByKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(30, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) DeleteValueByIndex(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(31, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetValueByKey(key string) (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(32, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) GetValueByIndex(index int32) (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(33, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) SetValueByKey(key string, value ICefv8Value, attribute cefTypes.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(34, m.Instance(), api.PasStr(key), base.GetObjectUintptr(value), uintptr(attribute))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) SetValueByIndex(index int32, value ICefv8Value) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(35, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) SetValueByAccessor(key string, attribute cefTypes.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(36, m.Instance(), api.PasStr(key), uintptr(attribute))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetKeys(keys lcl.IStrings) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(37, m.Instance(), base.GetObjectUintptr(keys))
	return int32(r)
}

func (m *TCefv8ValueRef) SetUserData(data ICefv8Value) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(38, m.Instance(), base.GetObjectUintptr(data))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetUserData() (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(39, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) GetExternallyAllocatedMemory() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(40, m.Instance())
	return int32(r)
}

func (m *TCefv8ValueRef) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(41, m.Instance(), uintptr(changeInBytes))
	return int32(r)
}

func (m *TCefv8ValueRef) GetArrayLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(42, m.Instance())
	return int32(r)
}

func (m *TCefv8ValueRef) GetArrayBufferReleaseCallback() (result IEngV8ArrayBufferReleaseCallback) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(43, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngV8ArrayBufferReleaseCallback(resultPtr)
	return
}

func (m *TCefv8ValueRef) NeuterArrayBuffer() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(44, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) GetArrayBufferByteLength() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(45, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefv8ValueRef) GetArrayBufferData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefv8ValueRefAPI().SysCallN(46, m.Instance())
	return uintptr(r)
}

func (m *TCefv8ValueRef) GetFunctionName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefv8ValueRefAPI().SysCallN(47, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefv8ValueRef) GetFunctionHandler() (result IEngV8Handler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(48, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngV8Handler(resultPtr)
	return
}

func (m *TCefv8ValueRef) ExecuteFunction(obj ICefv8Value, arguments ICefv8ValueArray) (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(49, m.Instance(), base.GetObjectUintptr(obj), arguments.Instance(), uintptr(int32(arguments.Count())), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefv8ValueArray) (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(50, m.Instance(), base.GetObjectUintptr(context), base.GetObjectUintptr(obj), arguments.Instance(), uintptr(int32(arguments.Count())), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ValueRef) ResolvePromise(arg ICefv8Value) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(51, m.Instance(), base.GetObjectUintptr(arg))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) RejectPromise(errorMsg string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ValueRefAPI().SysCallN(52, m.Instance(), api.PasStr(errorMsg))
	return api.GoBool(r)
}

func (m *TCefv8ValueRef) AsIntfV8Value() uintptr {
	return m.GetIntfPointer(0)
}

// V8ValueRef  is static instance
var V8ValueRef _V8ValueRefClass

// _V8ValueRefClass is class type defined by TCefv8ValueRef
type _V8ValueRefClass uintptr

// UnWrap
//
//	Returns a ICefv8Value instance using a PCefv8Value data pointer.
func (_V8ValueRefClass) UnWrap(data uintptr) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(53, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewUndefined
//
//	Create a new ICefv8Value object of type undefined.
func (_V8ValueRefClass) NewUndefined() (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(54, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewNull
//
//	Create a new ICefv8Value object of type null.
func (_V8ValueRefClass) NewNull() (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(55, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewBool
//
//	Create a new ICefv8Value object of type bool.
func (_V8ValueRefClass) NewBool(value bool) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(56, api.PasBool(value), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewInt
//
//	Create a new ICefv8Value object of type int.
func (_V8ValueRefClass) NewInt(value int32) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(57, uintptr(value), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewUInt
//
//	Create a new ICefv8Value object of type unsigned int.
func (_V8ValueRefClass) NewUInt(value uint32) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(58, uintptr(value), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewDouble
//
//	Create a new ICefv8Value object of type double.
func (_V8ValueRefClass) NewDouble(value float64) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(59, uintptr(base.UnsafePointer(&value)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewDate
//
//	Create a new ICefv8Value object of type Date. This function should only be
//	called from within the scope of a ICefRenderProcessHandler,
//	ICefv8Handler or ICefv8Accessor callback, or in combination with calling
//	enter() and exit() on a stored ICefv8Context reference.
func (_V8ValueRefClass) NewDate(value types.TDateTime) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(60, uintptr(base.UnsafePointer(&value)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewString
//
//	Create a new ICefv8Value object of type string.
func (_V8ValueRefClass) NewString(str string) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(61, api.PasStr(str), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewObject
//
//	Create a new ICefv8Value object of type object with optional accessor
//	and/or interceptor. This function should only be called from within the
//	scope of a ICefRenderProcessHandler, ICefv8Handler or ICefv8Accessor
//	callback, or in combination with calling enter() and exit() on a stored
//	ICefv8Context reference.
func (_V8ValueRefClass) NewObject(accessor IEngV8Accessor, interceptor IEngV8Interceptor) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(62, base.GetObjectUintptr(accessor), base.GetObjectUintptr(interceptor), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewArray
//
//	Create a new ICefv8Value object of type array with the specified |length|.
//	If |length| is negative the returned array will have length 0. This function
//	should only be called from within the scope of a
//	ICefRenderProcessHandler, ICefv8Handler or ICefv8Accessor callback,
//	or in combination with calling enter() and exit() on a stored
//	ICefv8Context reference.
func (_V8ValueRefClass) NewArray(len int32) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(63, uintptr(len), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewArrayBuffer
//
//	Create a new ICefv8Value object of type ArrayBuffer which wraps the
//	provided |buffer| of size |length| bytes. The ArrayBuffer is externalized,
//	meaning that it does not own |buffer|. The caller is responsible for freeing
//	|buffer| when requested via a call to
//	ICefv8ArrayBufferReleaseCallback.ReleaseBuffer. This function should
//	only be called from within the scope of a ICefRenderProcessHandler,
//	ICefv8Handler or ICefv8Accessor callback, or in combination with calling
//	enter() and exit() on a stored ICefv8Context reference.
func (_V8ValueRefClass) NewArrayBuffer(buffer uintptr, length cefTypes.NativeUInt, callback IEngV8ArrayBufferReleaseCallback) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(64, uintptr(buffer), uintptr(length), base.GetObjectUintptr(callback), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewFunction
//
//	Create a new ICefv8Value object of type function. This function should
//	only be called from within the scope of a ICefRenderProcessHandler,
//	ICefv8Handler or ICefv8Accessor callback, or in combination with calling
//	enter() and exit() on a stored ICefv8Context reference.
func (_V8ValueRefClass) NewFunction(name string, handler IEngV8Handler) (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(65, api.PasStr(name), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewPromise
//
//	Create a new ICefv8Value object of type Promise. This function should only
//	be called from within the scope of a ICefRenderProcessHandler,
//	ICefv8Handler or ICefv8Accessor callback, or in combination with calling
//	enter() and exit() on a stored ICefv8Context reference.
func (_V8ValueRefClass) NewPromise() (result ICefv8Value) {
	var resultPtr uintptr
	cefv8ValueRefAPI().SysCallN(66, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

// NewV8ValueRef class constructor
func NewV8ValueRef(data uintptr) ICefv8ValueRef {
	var v8ValuePtr uintptr // ICefv8Value
	r := cefv8ValueRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8ValuePtr)))
	ret := AsCefv8ValueRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8ValuePtr)
	}
	return ret
}

var (
	cefv8ValueRefOnce   base.Once
	cefv8ValueRefImport *imports.Imports = nil
)

func cefv8ValueRefAPI() *imports.Imports {
	cefv8ValueRefOnce.Do(func() {
		cefv8ValueRefImport = api.NewDefaultImports()
		cefv8ValueRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefv8ValueRef_Create", 0), // constructor NewV8ValueRef
			/* 1 */ imports.NewTable("TCefv8ValueRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefv8ValueRef_IsUndefined", 0), // function IsUndefined
			/* 3 */ imports.NewTable("TCefv8ValueRef_IsNull", 0), // function IsNull
			/* 4 */ imports.NewTable("TCefv8ValueRef_IsBool", 0), // function IsBool
			/* 5 */ imports.NewTable("TCefv8ValueRef_IsInt", 0), // function IsInt
			/* 6 */ imports.NewTable("TCefv8ValueRef_IsUInt", 0), // function IsUInt
			/* 7 */ imports.NewTable("TCefv8ValueRef_IsDouble", 0), // function IsDouble
			/* 8 */ imports.NewTable("TCefv8ValueRef_IsDate", 0), // function IsDate
			/* 9 */ imports.NewTable("TCefv8ValueRef_IsString", 0), // function IsString
			/* 10 */ imports.NewTable("TCefv8ValueRef_IsObject", 0), // function IsObject
			/* 11 */ imports.NewTable("TCefv8ValueRef_IsArray", 0), // function IsArray
			/* 12 */ imports.NewTable("TCefv8ValueRef_IsArrayBuffer", 0), // function IsArrayBuffer
			/* 13 */ imports.NewTable("TCefv8ValueRef_IsFunction", 0), // function IsFunction
			/* 14 */ imports.NewTable("TCefv8ValueRef_IsPromise", 0), // function IsPromise
			/* 15 */ imports.NewTable("TCefv8ValueRef_IsSame", 0), // function IsSame
			/* 16 */ imports.NewTable("TCefv8ValueRef_GetBoolValue", 0), // function GetBoolValue
			/* 17 */ imports.NewTable("TCefv8ValueRef_GetIntValue", 0), // function GetIntValue
			/* 18 */ imports.NewTable("TCefv8ValueRef_GetUIntValue", 0), // function GetUIntValue
			/* 19 */ imports.NewTable("TCefv8ValueRef_GetDoubleValue", 0), // function GetDoubleValue
			/* 20 */ imports.NewTable("TCefv8ValueRef_GetDateValue", 0), // function GetDateValue
			/* 21 */ imports.NewTable("TCefv8ValueRef_GetStringValue", 0), // function GetStringValue
			/* 22 */ imports.NewTable("TCefv8ValueRef_IsUserCreated", 0), // function IsUserCreated
			/* 23 */ imports.NewTable("TCefv8ValueRef_HasException", 0), // function HasException
			/* 24 */ imports.NewTable("TCefv8ValueRef_GetException", 0), // function GetException
			/* 25 */ imports.NewTable("TCefv8ValueRef_ClearException", 0), // function ClearException
			/* 26 */ imports.NewTable("TCefv8ValueRef_WillRethrowExceptions", 0), // function WillRethrowExceptions
			/* 27 */ imports.NewTable("TCefv8ValueRef_SetRethrowExceptions", 0), // function SetRethrowExceptions
			/* 28 */ imports.NewTable("TCefv8ValueRef_HasValueByKey", 0), // function HasValueByKey
			/* 29 */ imports.NewTable("TCefv8ValueRef_HasValueByIndex", 0), // function HasValueByIndex
			/* 30 */ imports.NewTable("TCefv8ValueRef_DeleteValueByKey", 0), // function DeleteValueByKey
			/* 31 */ imports.NewTable("TCefv8ValueRef_DeleteValueByIndex", 0), // function DeleteValueByIndex
			/* 32 */ imports.NewTable("TCefv8ValueRef_GetValueByKey", 0), // function GetValueByKey
			/* 33 */ imports.NewTable("TCefv8ValueRef_GetValueByIndex", 0), // function GetValueByIndex
			/* 34 */ imports.NewTable("TCefv8ValueRef_SetValueByKey", 0), // function SetValueByKey
			/* 35 */ imports.NewTable("TCefv8ValueRef_SetValueByIndex", 0), // function SetValueByIndex
			/* 36 */ imports.NewTable("TCefv8ValueRef_SetValueByAccessor", 0), // function SetValueByAccessor
			/* 37 */ imports.NewTable("TCefv8ValueRef_GetKeys", 0), // function GetKeys
			/* 38 */ imports.NewTable("TCefv8ValueRef_SetUserData", 0), // function SetUserData
			/* 39 */ imports.NewTable("TCefv8ValueRef_GetUserData", 0), // function GetUserData
			/* 40 */ imports.NewTable("TCefv8ValueRef_GetExternallyAllocatedMemory", 0), // function GetExternallyAllocatedMemory
			/* 41 */ imports.NewTable("TCefv8ValueRef_AdjustExternallyAllocatedMemory", 0), // function AdjustExternallyAllocatedMemory
			/* 42 */ imports.NewTable("TCefv8ValueRef_GetArrayLength", 0), // function GetArrayLength
			/* 43 */ imports.NewTable("TCefv8ValueRef_GetArrayBufferReleaseCallback", 0), // function GetArrayBufferReleaseCallback
			/* 44 */ imports.NewTable("TCefv8ValueRef_NeuterArrayBuffer", 0), // function NeuterArrayBuffer
			/* 45 */ imports.NewTable("TCefv8ValueRef_GetArrayBufferByteLength", 0), // function GetArrayBufferByteLength
			/* 46 */ imports.NewTable("TCefv8ValueRef_GetArrayBufferData", 0), // function GetArrayBufferData
			/* 47 */ imports.NewTable("TCefv8ValueRef_GetFunctionName", 0), // function GetFunctionName
			/* 48 */ imports.NewTable("TCefv8ValueRef_GetFunctionHandler", 0), // function GetFunctionHandler
			/* 49 */ imports.NewTable("TCefv8ValueRef_ExecuteFunction", 0), // function ExecuteFunction
			/* 50 */ imports.NewTable("TCefv8ValueRef_ExecuteFunctionWithContext", 0), // function ExecuteFunctionWithContext
			/* 51 */ imports.NewTable("TCefv8ValueRef_ResolvePromise", 0), // function ResolvePromise
			/* 52 */ imports.NewTable("TCefv8ValueRef_RejectPromise", 0), // function RejectPromise
			/* 53 */ imports.NewTable("TCefv8ValueRef_UnWrap", 0), // static function UnWrap
			/* 54 */ imports.NewTable("TCefv8ValueRef_NewUndefined", 0), // static function NewUndefined
			/* 55 */ imports.NewTable("TCefv8ValueRef_NewNull", 0), // static function NewNull
			/* 56 */ imports.NewTable("TCefv8ValueRef_NewBool", 0), // static function NewBool
			/* 57 */ imports.NewTable("TCefv8ValueRef_NewInt", 0), // static function NewInt
			/* 58 */ imports.NewTable("TCefv8ValueRef_NewUInt", 0), // static function NewUInt
			/* 59 */ imports.NewTable("TCefv8ValueRef_NewDouble", 0), // static function NewDouble
			/* 60 */ imports.NewTable("TCefv8ValueRef_NewDate", 0), // static function NewDate
			/* 61 */ imports.NewTable("TCefv8ValueRef_NewString", 0), // static function NewString
			/* 62 */ imports.NewTable("TCefv8ValueRef_NewObject", 0), // static function NewObject
			/* 63 */ imports.NewTable("TCefv8ValueRef_NewArray", 0), // static function NewArray
			/* 64 */ imports.NewTable("TCefv8ValueRef_NewArrayBuffer", 0), // static function NewArrayBuffer
			/* 65 */ imports.NewTable("TCefv8ValueRef_NewFunction", 0), // static function NewFunction
			/* 66 */ imports.NewTable("TCefv8ValueRef_NewPromise", 0), // static function NewPromise
		}
	})
	return cefv8ValueRefImport
}
