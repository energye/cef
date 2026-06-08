//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package base

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/tool/ptr"
)

type Callback struct {
	Name string
	Cb   func(getVal func(i int) uintptr)
}

// 移除事件，释放相关的引用
func removeEventCallbackProc(f uintptr) uintptr {
	api.RemoveEventElement(f)
	return 0
}

func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := api.PtrToElementValue(f)
	if fn != nil {
		if cb, ok := fn.(*Callback); ok {
			defer func() {
				if err := recover(); err != nil {
					// TODO 增加回调到用户
					println("[ERROR] CallbackEvent:", cb.Name, err.(error).Error())
				}
			}()
			getVal := func(i int) uintptr {
				return ptr.GetParamOf(i, args)
			}
			cb.Cb(getVal)
		}
	}
	return 0
}

func NewCallback(name string, cb func(getVal func(i int) uintptr)) *Callback {
	return &Callback{Name: name, Cb: cb}
}

func GetPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}
