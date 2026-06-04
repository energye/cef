//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !prod

package config

import (
	"encoding/json"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"os"
	"path/filepath"
)

func init() {
	homeDIR := exec.HomeDir
	if homeDIR == "" {
		println("[Warning] failed to obtain the current user directory.")
	} else {
		configPath := filepath.Join(homeDIR, ".energy", "config.json")
		if tool.IsExist(configPath) {
			data, err := os.ReadFile(configPath)
			if err != nil {
				println("[ERROR] ReadFile .energy Error:", err.Error())
				return
			}
			err = json.Unmarshal(data, GConfig)
			if err != nil {
				println("[ERROR] Unmarshal .energy file Error:", err.Error())
				return
			}
		}
	}
}

func (m *Config) FrameworkPath() string {
	return m.Framework
}

func (m *Config) ChromiumPath() string {
	return resolveChromiumPath(m.Chromium.Dir)
}
