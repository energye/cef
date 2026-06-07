//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
)

type TCefWindowInfo struct {
	WindowName                 string                    // TCefString
	Bounds                     TCefRect                  // TCefRect
	ParentWindow               cefTypes.TCefWindowHandle // TCefWindowHandle
	WindowlessRenderingEnabled int32                     // Integer
	SharedTextureEnabled       int32                     // Integer
	ExternalBeginFrameEnabled  int32                     // Integer
	Window                     cefTypes.TCefWindowHandle // TCefWindowHandle
}

func (m *TCefWindowInfo) ToPas() *tCefWindowInfo {
	if m == nil {
		return nil
	}
	return &tCefWindowInfo{
		WindowName:                 api.PasStr(m.WindowName),
		Bounds:                     m.Bounds,
		ParentWindow:               m.ParentWindow,
		WindowlessRenderingEnabled: m.WindowlessRenderingEnabled,
		SharedTextureEnabled:       m.SharedTextureEnabled,
		ExternalBeginFrameEnabled:  m.ExternalBeginFrameEnabled,
		Window:                     m.Window,
	}
}

type tCefWindowInfo struct {
	WindowName                 uintptr                   // TCefString
	Bounds                     TCefRect                  // TCefRect
	ParentWindow               cefTypes.TCefWindowHandle // TCefWindowHandle
	WindowlessRenderingEnabled int32                     // Integer
	SharedTextureEnabled       int32                     // Integer
	ExternalBeginFrameEnabled  int32                     // Integer
	Window                     cefTypes.TCefWindowHandle // TCefWindowHandle
}

func (m *tCefWindowInfo) ToGo() TCefWindowInfo {
	if m == nil {
		return TCefWindowInfo{}
	}
	return TCefWindowInfo{
		WindowName:                 api.GoStr(m.WindowName),
		Bounds:                     m.Bounds,
		ParentWindow:               m.ParentWindow,
		WindowlessRenderingEnabled: m.WindowlessRenderingEnabled,
		SharedTextureEnabled:       m.SharedTextureEnabled,
		ExternalBeginFrameEnabled:  m.ExternalBeginFrameEnabled,
		Window:                     m.Window,
	}
}
