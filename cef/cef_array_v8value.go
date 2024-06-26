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
	values   []uintptr
}

// V8ValueArrayRef -> TCefV8ValueArray
var V8ValueArrayRef v8ValueArray

// v8ValueArray
type v8ValueArray uintptr

// New 创建一个 ICefV8Value 数组引用维护
//
//	count: 初始默认大小
//	instance: 默认 nil(0)
func (*v8ValueArray) New(count int, instance uintptr) ICefV8ValueArray {
	return &TCefV8ValueArray{
		instance: unsafePointer(instance),
		count:    count,
		values:   make([]uintptr, count),
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
		if result == 0 {
			result = getParamOf(index, m.Instance())
			m.values[index] = result
		}
		return AsCefV8Value(result)
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
			if v != 0 {
				m.values[i] = 0
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

// Add 添加一个 ICefV8Value 后指针实例被重新改变
func (m *TCefV8ValueArray) Add(value ICefV8Value) {
	m.values = append(m.values, value.Instance())
	m.instance = unsafePointer(&m.values[0])
	m.count++
}

// Set 设置 []ICefV8Value 集合, 会覆盖已有的
// 如果value为空或大小为0则清空现有
func (m *TCefV8ValueArray) Set(value []ICefV8Value) {
	m.count = len(value)
	if m.count > 0 {
		m.values = make([]uintptr, m.count, m.count)
		for i, _ := range m.values {
			m.values[i] = value[i].Instance()
		}
		m.instance = unsafePointer(&m.values[0])
	} else {
		m.values = make([]uintptr, 0)
		m.instance = nil
	}
}
