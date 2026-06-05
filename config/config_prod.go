//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod

package config

import (
	"github.com/energye/lcl/tool/exec"
	"path/filepath"
)

func init() {
	GConfig.Chromium.Dir = "@" // default current path
}

// ChromiumPath returns the CEF framework directory
func (m *Config) ChromiumPath() string {
	// -tags prod: handle CEF directory loading in production mode
	// Dir: @ indicates relative path to the binary executable;
	// otherwise use custom configuration or compile-time path injected via -X
	dir := m.Chromium.Dir
	if dir != "" && dir[0] == '@' {
		execDir := exec.AppDir()
		dir = filepath.Join(execDir, dir[1:])
	}
	return dor
}
