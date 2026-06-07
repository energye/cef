//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

type TCefWindowInfo struct {
	ExStyle                    types.DWORD               // DWORD
	WindowName                 string                    // TCefString
	Style                      types.DWORD               // DWORD
	Bounds                     TCefRect                  // TCefRect
	ParentWindow               cefTypes.TCefWindowHandle // TCefWindowHandle
	Menu                       types.HMENU               // HMENU
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
		ExStyle:                    m.ExStyle,
		WindowName:                 api.PasStr(m.WindowName),
		Style:                      m.Style,
		Bounds:                     m.Bounds,
		ParentWindow:               m.ParentWindow,
		Menu:                       m.Menu,
		WindowlessRenderingEnabled: m.WindowlessRenderingEnabled,
		SharedTextureEnabled:       m.SharedTextureEnabled,
		ExternalBeginFrameEnabled:  m.ExternalBeginFrameEnabled,
		Window:                     m.Window,
	}
}

type tCefWindowInfo struct {
	ExStyle                    types.DWORD               // DWORD
	WindowName                 uintptr                   // TCefString
	Style                      types.DWORD               // DWORD
	Bounds                     TCefRect                  // TCefRect
	ParentWindow               cefTypes.TCefWindowHandle // TCefWindowHandle
	Menu                       types.HMENU               // HMENU
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
		ExStyle:                    m.ExStyle,
		WindowName:                 api.GoStr(m.WindowName),
		Style:                      m.Style,
		Bounds:                     m.Bounds,
		ParentWindow:               m.ParentWindow,
		Menu:                       m.Menu,
		WindowlessRenderingEnabled: m.WindowlessRenderingEnabled,
		SharedTextureEnabled:       m.SharedTextureEnabled,
		ExternalBeginFrameEnabled:  m.ExternalBeginFrameEnabled,
		Window:                     m.Window,
	}
}
