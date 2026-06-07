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

// ICefv8BackingStore Parent: ICefBaseRefCounted
type ICefv8BackingStore interface {
	ICefBaseRefCounted
	// Data
	//  Returns a pointer to the allocated memory, or nullptr if the backing store
	//  has been consumed or is otherwise invalid. The pointer is safe to
	//  read/write from any thread. The caller must ensure all writes are complete
	//  before passing this object to
	//  TCefv8ValueRef.NewArrayBufferFromBackingStore(). Pointers obtained
	//  from this function should not be retained after calling
	//  TCefv8ValueRef.NewArrayBufferFromBackingStore(), as the memory will
	//  then be owned by the ArrayBuffer and subject to V8 garbage collection.
	Data() uintptr // function
	// ByteLength
	//  Returns the size of the allocated memory in bytes, or 0 if the backing
	//  store has been consumed.
	ByteLength() cefTypes.NativeUInt // function
	// IsValid
	//  Returns true (1) if this backing store has not yet been consumed by
	//  TCefv8ValueRef.NewArrayBufferFromBackingStore().
	IsValid() bool // function
}

// ICefv8BackingStoreRef Parent: ICefv8BackingStore ICefBaseRefCountedRef
type ICefv8BackingStoreRef interface {
	ICefv8BackingStore
	ICefBaseRefCountedRef
	AsIntfV8BackingStore() uintptr
}

type TCefv8BackingStoreRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefv8BackingStoreRef) Data() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefv8BackingStoreRefAPI().SysCallN(1, m.Instance())
	return uintptr(r)
}

func (m *TCefv8BackingStoreRef) ByteLength() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefv8BackingStoreRefAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefv8BackingStoreRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefv8BackingStoreRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8BackingStoreRef) AsIntfV8BackingStore() uintptr {
	return m.GetIntfPointer(0)
}

// V8BackingStoreRef  is static instance
var V8BackingStoreRef _V8BackingStoreRefClass

// _V8BackingStoreRefClass is class type defined by TCefv8BackingStoreRef
type _V8BackingStoreRefClass uintptr

func (_V8BackingStoreRefClass) UnWrap(data uintptr) (result ICefv8BackingStore) {
	var resultPtr uintptr
	cefv8BackingStoreRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8BackingStoreRef(resultPtr)
	return
}

func (_V8BackingStoreRefClass) New(byteLength cefTypes.NativeUInt) (result ICefv8BackingStore) {
	var resultPtr uintptr
	cefv8BackingStoreRefAPI().SysCallN(5, uintptr(byteLength), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8BackingStoreRef(resultPtr)
	return
}

// NewV8BackingStoreRef class constructor
func NewV8BackingStoreRef(data uintptr) ICefv8BackingStoreRef {
	var v8BackingStorePtr uintptr // ICefv8BackingStore
	r := cefv8BackingStoreRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8BackingStorePtr)))
	ret := AsCefv8BackingStoreRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8BackingStorePtr)
	}
	return ret
}

var (
	cefv8BackingStoreRefOnce   base.Once
	cefv8BackingStoreRefImport *imports.Imports = nil
)

func cefv8BackingStoreRefAPI() *imports.Imports {
	cefv8BackingStoreRefOnce.Do(func() {
		cefv8BackingStoreRefImport = api.NewDefaultImports()
		cefv8BackingStoreRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefv8BackingStoreRef_Create", 0), // constructor NewV8BackingStoreRef
			/* 1 */ imports.NewTable("TCefv8BackingStoreRef_data", 0), // function Data
			/* 2 */ imports.NewTable("TCefv8BackingStoreRef_ByteLength", 0), // function ByteLength
			/* 3 */ imports.NewTable("TCefv8BackingStoreRef_IsValid", 0), // function IsValid
			/* 4 */ imports.NewTable("TCefv8BackingStoreRef_UnWrap", 0), // static function UnWrap
			/* 5 */ imports.NewTable("TCefv8BackingStoreRef_New", 0), // static function New
		}
	})
	return cefv8BackingStoreRefImport
}
