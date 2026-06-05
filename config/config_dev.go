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

	var latestName string
	var latestVersion []int
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() || !strings.HasPrefix(name, prefix) {
			continue
		}

		version, ok := parseVersion(name[len(prefix):])
		if !ok {
			continue
		}
		if latestName == "" || newerVersion(version, latestVersion) {
			latestName = name
			latestVersion = version
		}
	}

	if latestName == "" {
		println("[Warning] no matching CEF version directory found, prefix:", prefix, "in:", chromiumDir)
		return chromiumDir
	}
	return filepath.Join(chromiumDir, latestName)
}

func parseVersion(version string) ([]int, bool) {
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return nil, false
	}

	nums := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, false
		}
		nums = append(nums, num)
	}
	return nums, true
}

func newerVersion(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return len(a) > len(b)
}
