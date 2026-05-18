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

type TCefAcceleratedPaintInfo struct {
	Planes     uintptr                // array [0..pred(CEF_KACCELERATEDPAINTMAXPLANES)] of TCefAcceleratedPaintNativePixmapPlaneInfo
	PlanesSize int32                  // SizeOf: array [0..pred(CEF_KACCELERATEDPAINTMAXPLANES)] of TCefAcceleratedPaintNativePixmapPlaneInfo
	PlaneCount int32                  // integer
	Modifier   uint64                 // uint64
	Format     cefTypes.TCefColorType // TCefColorType
}

type TCefAcceleratedPaintNativePixmapPlaneInfo struct {
	Stride uint32 // Cardinal
	Offset uint64 // uint64
	Size   uint64 // uint64
	Fd     int32  // integer
}

type TCefWindowInfo struct {
	WindowName                 string                    // TCefString
	Bounds                     TCefRect                  // TCefRect
	ParentWindow               cefTypes.TCefWindowHandle // TCefWindowHandle
	WindowlessRenderingEnabled int32                     // Integer
	SharedTextureEnabled       int32                     // Integer
	ExternalBeginFrameEnabled  int32                     // Integer
	Window                     cefTypes.TCefWindowHandle // TCefWindowHandle
	RuntimeStyle               cefTypes.TCefRuntimeStyle // TCefRuntimeStyle
}

func (m *TCefAcceleratedPaintInfo) ToPas() *tCefAcceleratedPaintInfo {
	if m == nil {
		return nil
	}
	return &tCefAcceleratedPaintInfo{
		Planes:     m.Planes,
		PlanesSize: m.PlanesSize,
		PlaneCount: m.PlaneCount,
		Modifier:   m.Modifier,
		Format:     m.Format,
	}
}

type tCefAcceleratedPaintInfo struct {
	Planes     uintptr                // array [0..pred(CEF_KACCELERATEDPAINTMAXPLANES)] of TCefAcceleratedPaintNativePixmapPlaneInfo
	PlanesSize int32                  // SizeOf: array [0..pred(CEF_KACCELERATEDPAINTMAXPLANES)] of TCefAcceleratedPaintNativePixmapPlaneInfo
	PlaneCount int32                  // integer
	Modifier   uint64                 // uint64
	Format     cefTypes.TCefColorType // TCefColorType
}

func (m *tCefAcceleratedPaintInfo) ToGo() TCefAcceleratedPaintInfo {
	if m == nil {
		return TCefAcceleratedPaintInfo{}
	}
	return TCefAcceleratedPaintInfo{
		Planes:     m.Planes,
		PlanesSize: m.PlanesSize,
		PlaneCount: m.PlaneCount,
		Modifier:   m.Modifier,
		Format:     m.Format,
	}
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
		RuntimeStyle:               m.RuntimeStyle,
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
	RuntimeStyle               cefTypes.TCefRuntimeStyle // TCefRuntimeStyle
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
		RuntimeStyle:               m.RuntimeStyle,
	}
}
