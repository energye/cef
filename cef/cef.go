//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/base"
	cefTypes "github.com/energye/cef/cef/types"
	"sync"
)

var (
	gLibVersionOnce                   = sync.Once{}
	gMajor, gmMinor, gRelease, gBuild int32
)

func LibVersion() (major, minor, release, build int) {
	gLibVersionOnce.Do(func() {
		gMajor, gmMinor, gRelease, gBuild = base.CEFLibVersion()
	})
	return int(gMajor), int(gmMinor), int(gRelease), int(gBuild)
}

func RunOnMainThread(fn func()) {
	if fn == nil {
		return
	}
	if MiscFunc.CefCurrentlyOn(cefTypes.TID_UI) {
		fn()
		return
	}
	task := NewEngTask()
	task.SetOnTaskExecute(func() {
		fn()
	})
	task = AsEngTask(task.AsIntfTask())
	MiscFunc.CefPostTask(cefTypes.TID_UI, task)
	task.Release()
}
