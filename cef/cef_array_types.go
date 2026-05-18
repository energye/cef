//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package cef

import "github.com/energye/lcl/base"

// ICefRectArray = array[0..(High(integer) div SizeOf(TCefRect)) - 1] of TCefRect;
type ICefRectArray interface {
	base.IArraySlice[TCefRect]
}

// ICefRangeArray = array[0..(High(integer) div SizeOf(TCefRange)) - 1] of TCefRange;
// 内部和外部数组维护
type ICefRangeArray interface {
	base.IArraySlice[TCefRange]
}

// ICefFrameIdentifierArray = array of int64
//type ICefFrameIdentifierArray interface {
//	IArraySlice[int64]
//}

// ICefDraggableRegionArray = array[0..(High(integer) div SizeOf(TCefDraggableRegion)) - 1] of TCefDraggableRegion;
// 内部和外部数组维护
type ICefDraggableRegionArray interface {
	base.IArraySlice[TCefDraggableRegion]
}

// ICefCompositionUnderlineArray = array[0..(High(integer) div SizeOf(TCefCompositionUnderline)) - 1] of TCefCompositionUnderline;
// 内部和外部数组维护
type ICefCompositionUnderlineArray interface {
	base.IArraySlice[TCefCompositionUnderline]
}

// NewCefRectArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefRectArray(count int, instance uintptr) ICefRectArray {
	return base.NewStructArraySlice[TCefRect](count, instance)
}
func NewCefRectDynArray(count int, instance uintptr) ICefRectArray {
	return NewCefRectArray(count, instance)
}

// NewCefRangeArray 初始化 TCefRange 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefRangeArray(count int, instance uintptr) ICefRangeArray {
	return base.NewStructArraySlice[TCefRange](count, instance)
}

// NewCefFrameIdentifierArray 初始化 int64 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
//func NewCefFrameIdentifierArray(count int, instance uintptr) ICefFrameIdentifierArray {
//	return NewStructArraySlice[int64](count, instance)
//}

// NewCefDraggableRegionArray 初始化 TCefDraggableRegion 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefDraggableRegionArray(count int, instance uintptr) ICefDraggableRegionArray {
	return base.NewStructArraySlice[TCefDraggableRegion](count, instance)
}

// NewCefCompositionUnderlineArray 初始化 TCefCompositionUnderline 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefCompositionUnderlineArray(count int, instance uintptr) ICefCompositionUnderlineArray {
	return base.NewStructArraySlice[TCefCompositionUnderline](count, instance)
}

func NewCefCompositionUnderlineDynArray(count int, instance uintptr) ICefCompositionUnderlineArray {
	return NewCefCompositionUnderlineArray(count, instance)
}

// ICefBinaryValueArray = array of ICefBinaryValue
type ICefBinaryValueArray interface {
	base.IArraySlice[ICefBinaryValue]
}

// ICefv8ValueArray = array of ICefv8Value
type ICefv8ValueArray interface {
	base.IArraySlice[ICefv8Value]
}

// ICefX509CertificateArray = array of ICefX509Certificate
type ICefX509CertificateArray interface {
	base.IArraySlice[ICefX509Certificate]
}

// ICefPostDataElementArray = array of ICefPostDataElement
type ICefPostDataElementArray interface {
	base.IArraySlice[ICefPostDataElement]
}

// ICefMediaSinkArray = array of ICefMediaSink
type ICefMediaSinkArray interface {
	base.IArraySlice[ICefMediaSink]
}

// ICefMediaRouteArray = array of ICefMediaRoute
type ICefMediaRouteArray interface {
	base.IArraySlice[ICefMediaRoute]
}

// ICefDisplayArray = array of ICefDisplay
type ICefDisplayArray interface {
	base.IArraySlice[ICefDisplay]
}

// NewCefBinaryValueArray 初始化 ICefBinaryValue 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefBinaryValueArray(count int, instance uintptr) ICefBinaryValueArray {
	return base.NewInterfaceArraySlice[ICefBinaryValue](count, instance, func(obj any) ICefBinaryValue {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefBinaryValueRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefv8ValueArray 初始化 ICefv8Value 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefv8ValueArray(count int, instance uintptr) ICefv8ValueArray {
	return base.NewInterfaceArraySlice[ICefv8Value](count, instance, func(obj any) ICefv8Value {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefv8ValueRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefX509CertificateArray 初始化 ICEFX509Certificate 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefX509CertificateArray(count int, instance uintptr) ICefX509CertificateArray {
	return base.NewInterfaceArraySlice[ICefX509Certificate](count, instance, func(obj any) ICefX509Certificate {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCEFX509CertificateRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefPostDataElementArray 初始化 ICefPostDataElement 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefPostDataElementArray(count int, instance uintptr) ICefPostDataElementArray {
	return base.NewInterfaceArraySlice[ICefPostDataElement](count, instance, func(obj any) ICefPostDataElement {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefPostDataElementRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefMediaSinkArray 初始化 ICefMediaSinkRef 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefMediaSinkArray(count int, instance uintptr) ICefMediaSinkArray {
	return base.NewInterfaceArraySlice[ICefMediaSink](count, instance, func(obj any) ICefMediaSink {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefMediaSinkRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefMediaRouteArray 初始化 ICefMediaRoute 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefMediaRouteArray(count int, instance uintptr) ICefMediaRouteArray {
	return base.NewInterfaceArraySlice[ICefMediaRoute](count, instance, func(obj any) ICefMediaRoute {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefMediaRouteRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}

// NewCefDisplayArray 初始化 ICefDisplay 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewCefDisplayArray(count int, instance uintptr) ICefDisplayArray {
	return base.NewInterfaceArraySlice[ICefDisplay](count, instance, func(obj any) ICefDisplay {
		instance := base.GetInstance(obj)
		if instance == nil {
			return nil
		}
		result := new(TCefDisplayRef)
		base.SetObjectInstance(result, instance)
		return result
	})
}
