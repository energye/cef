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
	"runtime"
	"sort"
	"strconv"
	"strings"
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

func (m *Config) ChromiumPath() string {
	if m.Chromium.Version != "" {
		return filepath.Join(m.Chromium.Dir, m.Chromium.Version)
	}
	return resolveChromiumPath(m.Chromium.Dir)
}

// resolveChromiumPath 从 chromium 根目录中查找匹配当前系统和架构的最新版本目录
// 目录格式: {os}_{arch}_{version}，例如 windows_amd64_127.3.5
func resolveChromiumPath(chromiumDir string) string {
	if chromiumDir == "" {
		println("[Warning] chromium directory is empty, please check the config: chromium.dir")
		return chromiumDir
	}
	if !tool.IsExist(chromiumDir) {
		println("[Warning] chromium directory does not exist:", chromiumDir)
		return chromiumDir
	}
	prefix := runtime.GOOS + "_" + runtime.GOARCH + "_"
	entries, err := os.ReadDir(chromiumDir)
	if err != nil {
		println("[Warning] failed to read chromium directory:", chromiumDir, "error:", err.Error())
		return chromiumDir
	}
	// 收集所有匹配当前系统架构的版本目录
	type versionDir struct {
		name    string
		version []int
	}
	var matched []versionDir
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		versionStr := strings.TrimPrefix(name, prefix)
		parts := strings.Split(versionStr, ".")
		if len(parts) < 2 {
			continue
		}
		var nums []int
		valid := true
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				valid = false
				break
			}
			nums = append(nums, n)
		}
		if !valid {
			continue
		}
		matched = append(matched, versionDir{name: name, version: nums})
	}
	if len(matched) == 0 {
		println("[Warning] no matching CEF version directory found, prefix:", prefix, "in:", chromiumDir)
		return chromiumDir
	}
	// 按版本号降序排序，取最新版本
	sort.Slice(matched, func(i, j int) bool {
		a, b := matched[i].version, matched[j].version
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] > b[k]
			}
		}
		return len(a) > len(b)
	})
	return filepath.Join(chromiumDir, matched[0].name)
}
