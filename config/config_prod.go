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

func (m *Config) FrameworkPath() string {
	return m.Framework
}

func (m *Config) ChromiumPath() string {
	return m.Chromium.Dir
}
