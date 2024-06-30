//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

// 函数工具 - windows

package cef

import (
	"github.com/energye/lcl/api/winapi"
	"github.com/energye/lcl/types"
	. "github.com/energye/lcl/types/keys"
	"github.com/energye/lcl/types/messages"
)

func CefIsKeyDown(wparam types.WPARAM) bool {
	return winapi.GetKeyState(types.Integer(wparam)) < 0
}

func CefIsKeyToggled(wparam types.WPARAM) bool {
	return (int16(winapi.GetKeyState(types.Integer(wparam))) & 0x1) != 0
}

func GetCefMouseModifiersByWPARAM(wparam types.WPARAM) (result TCefEventFlags) {
	result = EVENTFLAG_NONE
	if (wparam & messages.MK_CONTROL) != 0 {
		result = result | EVENTFLAG_CONTROL_DOWN
	}
	if (wparam & messages.MK_SHIFT) != 0 {
		result = result | EVENTFLAG_SHIFT_DOWN
	}
	if (wparam & messages.MK_LBUTTON) != 0 {
		result = result | EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if (wparam & messages.MK_MBUTTON) != 0 {
		result = result | EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if (wparam & messages.MK_RBUTTON) != 0 {
		result = result | EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(VkMenu) {
		result = result | EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(VkNumLock) {
		result = result | EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(VkCapital) {
		result = result | EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefMouseModifiers() (result TCefEventFlags) {
	result = EVENTFLAG_NONE
	if CefIsKeyDown(messages.MK_CONTROL) {
		result = result | EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(messages.MK_SHIFT) {
		result = result | EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(messages.MK_LBUTTON) {
		result = result | EVENTFLAG_LEFT_MOUSE_BUTTON
	}
	if CefIsKeyDown(messages.MK_MBUTTON) {
		result = result | EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	if CefIsKeyDown(messages.MK_RBUTTON) {
		result = result | EVENTFLAG_RIGHT_MOUSE_BUTTON
	}
	if CefIsKeyDown(VkMenu) {
		result = result | EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(VkNumLock) {
		result = result | EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(VkCapital) {
		result = result | EVENTFLAG_CAPS_LOCK_ON
	}
	return
}

func GetCefKeyboardModifiers(aWparam types.WPARAM, aLparam types.LPARAM) (result TCefEventFlags) {
	result = EVENTFLAG_NONE
	if CefIsKeyDown(VkShift) {
		result = result | EVENTFLAG_SHIFT_DOWN
	}
	if CefIsKeyDown(VkControl) {
		result = result | EVENTFLAG_CONTROL_DOWN
	}
	if CefIsKeyDown(VkMenu) {
		result = result | EVENTFLAG_ALT_DOWN
	}
	if CefIsKeyToggled(VkNumLock) {
		result = result | EVENTFLAG_NUM_LOCK_ON
	}
	if CefIsKeyToggled(VkCapital) {
		result = result | EVENTFLAG_CAPS_LOCK_ON
	}
	switch aWparam {
	case VkReturn:
		if ((aLparam >> 16) & KF_EXTENDED) != 0 {
			result = result | EVENTFLAG_IS_KEY_PAD
		}
	case VkInsert, VkDelete, VkHome, VkEnd, VkPrior, VkNext, VkUp, VkDown, VkLeft, VkRight:
		if ((aLparam >> 16) & winapi.KF_EXTENDED) == 0 {
			result = result | EVENTFLAG_IS_KEY_PAD
		}
	case VkNumLock, VkNumpad0, VkNumpad1, VkNumpad2, VkNumpad3, VkNumpad4, VkNumpad5, VkNumpad6, VkNumpad7,
		VkNumpad8, VkNumpad9, VkDivide, VkMultiply, VkSubtract, VkAdd, VkDecimal, VkClear:
		result = result | EVENTFLAG_IS_KEY_PAD
	case VkShift:
		if CefIsKeyDown(VkLShift) {
			result = result | EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(VkRShift) {
			result = result | EVENTFLAG_IS_RIGHT
		}
	case VkControl:
		if CefIsKeyDown(VkLControl) {
			result = result | EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(VkRControl) {
			result = result | EVENTFLAG_IS_RIGHT
		}
	case VkMenu:
		if CefIsKeyDown(VkLMenu) {
			result = result | EVENTFLAG_IS_LEFT
		} else if CefIsKeyDown(VkRMenu) {
			result = result | EVENTFLAG_IS_RIGHT
		}
	case VkLWin:
		result = result | EVENTFLAG_IS_LEFT
	case VkRWin:
		result = result | EVENTFLAG_IS_RIGHT
	}
	return
}
