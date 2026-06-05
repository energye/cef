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
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"os"
	"sync"
)

var initOnce sync.Once

// Init CEF
func Init() {
	initOnce.Do(func() {
		if tool.IsDarwin() {
			// MacOS SetCommandLine
			argsList := lcl.NewStringList()
			for _, v := range os.Args {
				argsList.Add(v)
			}
			SetCommandLine(argsList)
			argsList.Free()
		}

		// Register CEF object event callback
		api.SetEventCallback(eventCallback, api.EctCEF)
		// Unregister CEF object event callback
		api.SetEventCallback(removeEventCallback, api.EctCEFRemove)
	})
}
