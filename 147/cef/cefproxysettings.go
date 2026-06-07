//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICEFProxySettings Parent: IObject
type ICEFProxySettings interface {
	IObject
	NoProxyServer() bool         // property NoProxyServer Getter
	SetNoProxyServer(value bool) // property NoProxyServer Setter
	AutoDetect() bool            // property AutoDetect Getter
	SetAutoDetect(value bool)    // property AutoDetect Setter
	ByPassList() string          // property ByPassList Getter
	SetByPassList(value string)  // property ByPassList Setter
	PacUrl() string              // property PacUrl Getter
	SetPacUrl(value string)      // property PacUrl Setter
	Server() string              // property Server Getter
	SetServer(value string)      // property Server Setter
}

type TCEFProxySettings struct {
	TObject
}

func (m *TCEFProxySettings) NoProxyServer() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFProxySettingsAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFProxySettings) SetNoProxyServer(value bool) {
	if !m.IsValid() {
		return
	}
	cEFProxySettingsAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFProxySettings) AutoDetect() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFProxySettingsAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFProxySettings) SetAutoDetect(value bool) {
	if !m.IsValid() {
		return
	}
	cEFProxySettingsAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFProxySettings) ByPassList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFProxySettingsAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFProxySettings) SetByPassList(value string) {
	if !m.IsValid() {
		return
	}
	cEFProxySettingsAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFProxySettings) PacUrl() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFProxySettingsAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFProxySettings) SetPacUrl(value string) {
	if !m.IsValid() {
		return
	}
	cEFProxySettingsAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFProxySettings) Server() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFProxySettingsAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFProxySettings) SetServer(value string) {
	if !m.IsValid() {
		return
	}
	cEFProxySettingsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

// NewProxySettings class constructor
func NewProxySettings() ICEFProxySettings {
	r := cEFProxySettingsAPI().SysCallN(0)
	return AsCEFProxySettings(r)
}

var (
	cEFProxySettingsOnce   base.Once
	cEFProxySettingsImport *imports.Imports = nil
)

func cEFProxySettingsAPI() *imports.Imports {
	cEFProxySettingsOnce.Do(func() {
		cEFProxySettingsImport = api.NewDefaultImports()
		cEFProxySettingsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFProxySettings_Create", 0), // constructor NewProxySettings
			/* 1 */ imports.NewTable("TCEFProxySettings_NoProxyServer", 0), // property NoProxyServer
			/* 2 */ imports.NewTable("TCEFProxySettings_AutoDetect", 0), // property AutoDetect
			/* 3 */ imports.NewTable("TCEFProxySettings_ByPassList", 0), // property ByPassList
			/* 4 */ imports.NewTable("TCEFProxySettings_PacUrl", 0), // property PacUrl
			/* 5 */ imports.NewTable("TCEFProxySettings_Server", 0), // property Server
		}
	})
	return cEFProxySettingsImport
}
