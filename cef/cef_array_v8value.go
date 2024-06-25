//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefV8ValueArray
//
//	[]ICefV8Value
type ICefV8ValueArray interface {
	Instance() uintptr
	Get(index int) ICefV8Value
	Size() int
	Free()
	Add(value ICefV8Value)
	Set(value []ICefV8Value)
}

// TCefV8ValueArray
//
//	[]ICefV8Value
type TCefV8ValueArray struct {
	instance unsafePointer
	count    int
	values   []ICefV8Value
}

// V8ValueArrayRef -> TCefV8ValueArray
var V8ValueArrayRef v8ValueArray

// v8ValueArray
type v8ValueArray uintptr

func (*v8ValueArray) New(count int, instance uintptr) ICefV8ValueArray {
	return &TCefV8ValueArray{
		instance: unsafePointer(instance),
		count:    count,
		values:   make([]ICefV8Value, count),
	}
}

func (m *TCefV8ValueArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefV8Value
func (m *TCefV8ValueArray) Get(index int) ICefV8Value {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefV8Value(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefV8Value 数组长度
func (m *TCefV8ValueArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefV8ValueArray) Free() {
	if m == nil {
		return
	}
	if m.values != nil {
		for i, v := range m.values {
			if v != nil {
				v.Free()
				m.values[i] = nil
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

func (m *TCefV8ValueArray) Add(value ICefV8Value) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefV8ValueArray) Set(value []ICefV8Value) {
	if m.values != nil {
		for i, v := range m.values {
			if v != nil && v.Instance() != 0 {
				v.Free()
				m.values[i] = nil
			}
		}
	}
	m.values = value
	m.count = len(value)
	m.instance = unsafePointer(m.values[0].Instance())
}
