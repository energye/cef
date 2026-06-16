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
