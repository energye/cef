// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// Package Application environment initialization loads dynamic libraries

package initialize

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/locales/zh_CN"
	"github.com/energye/lcl/pkgs/i18n"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
)

// Initialize
// 初始化，运行时加载 LibLCL
func Initialize(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	// load lib dll
	loadLibLCL(libs, resources)
	// lib api
	api.APIInit()
	// rtl
	rtl.RtlInit()
	// version
	version.VersionInit()
	// zh_cn
	zh_CN.ZH_CNInit()
	// other api
	APIInit()
	// lcl
	lcl.LCLInit()
	// i18n
	i18n.I18NInit()
}
