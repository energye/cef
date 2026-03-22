//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package cef

import (
	"os"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
)

// Init CEF
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	if tool.IsDarwin() {
	}
	// LCL init
	lcl.Init(libs, resources)
	defer func() {
		if err := recover(); err != nil {
			println(err)
			os.Exit(1)
		}
	}()
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()

	if tool.IsDarwin() {
		// MacOS SetCommandLine
		argsList := lcl.NewStringList()
		for _, v := range os.Args {
			argsList.Add(v)
		}
		SetCommandLine(argsList)
		argsList.Free()
	}

	// 注册 CEF 对象事件回调
	api.SetEventCallback(eventCallback, api.EctCEF)
	// 注册 CEF 对象事件回调移除
	api.SetEventCallback(removeEventCallback, api.EctCEFRemove)

}
