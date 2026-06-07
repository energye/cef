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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/types"
)

// ICefBrowserHost Parent: ICefBaseRefCounted
type ICefBrowserHost interface {
	ICefBaseRefCounted
	GetBrowser() ICefBrowser                                                                                                                              // function
	TryCloseBrowser() bool                                                                                                                                // function
	GetWindowHandle() cefTypes.TCefWindowHandle                                                                                                           // function
	GetOpenerWindowHandle() cefTypes.TCefWindowHandle                                                                                                     // function
	HasView() bool                                                                                                                                        // function
	GetRequestContext() ICefRequestContext                                                                                                                // function
	GetZoomLevel() float64                                                                                                                                // function
	HasDevTools() bool                                                                                                                                    // function
	SendDevToolsMessage(message string) bool                                                                                                              // function
	ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32                                                               // function
	AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) ICefRegistration                                                                     // function
	IsWindowRenderingDisabled() bool                                                                                                                      // function
	GetWindowlessFrameRate() int32                                                                                                                        // function
	GetVisibleNavigationEntry() ICefNavigationEntry                                                                                                       // function
	GetExtension() ICefExtension                                                                                                                          // function
	IsBackgroundHost() bool                                                                                                                               // function
	IsAudioMuted() bool                                                                                                                                   // function
	CloseBrowser(forceClose bool)                                                                                                                         // procedure
	SetFocus(focus bool)                                                                                                                                  // procedure
	SetZoomLevel(zoomLevel float64)                                                                                                                       // procedure
	RunFileDialog(mode cefTypes.TCefFileDialogMode, title string, defaultFilePath string, acceptFilters lcl.IStrings, callback IEngRunFileDialogCallback) // procedure
	StartDownload(url string)                                                                                                                             // procedure
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback IEngDownloadImageCallback)                             // procedure
	Print()                                                                                                                                               // procedure
	PrintToPdf(path string, settings TCefPdfPrintSettings, callback IEngPdfPrintCallback)                                                                 // procedure
	Find(searchText string, forward bool, matchCase bool, findNext bool)                                                                                  // procedure
	StopFinding(clearSelection bool)                                                                                                                      // procedure
	ShowDevTools(windowInfo TCefWindowInfo, client IEngClient, settings TCefBrowserSettings, inspectElementAt TCefPoint)                                  // procedure
	CloseDevTools()                                                                                                                                       // procedure
	GetNavigationEntries(visitor IEngNavigationEntryVisitor, currentOnly bool)                                                                            // procedure
	ReplaceMisspelling(word string)                                                                                                                       // procedure
	AddWordToDictionary(word string)                                                                                                                      // procedure
	WasResized()                                                                                                                                          // procedure
	NotifyScreenInfoChanged()                                                                                                                             // procedure
	WasHidden(hidden bool)                                                                                                                                // procedure
	Invalidate(type_ cefTypes.TCefPaintElementType)                                                                                                       // procedure
	SendExternalBeginFrame()                                                                                                                              // procedure
	SendKeyEvent(event TCefKeyEvent)                                                                                                                      // procedure
	SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32)                                         // procedure
	SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool)                                                                                             // procedure
	SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32)                                                                                 // procedure
	SendTouchEvent(event TCefTouchEvent)                                                                                                                  // procedure
	SendCaptureLostEvent()                                                                                                                                // procedure
	NotifyMoveOrResizeStarted()                                                                                                                           // procedure
	SetWindowlessFrameRate(frameRate int32)                                                                                                               // procedure
	IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange)                        // procedure
	IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32)                                                                       // procedure
	IMEFinishComposingText(keepSelection bool)                                                                                                            // procedure
	IMECancelComposition()                                                                                                                                // procedure
	DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations)                                              // procedure
	DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations)                                                                      // procedure
	DragTargetDragLeave()                                                                                                                                 // procedure
	DragTargetDrop(event TCefMouseEvent)                                                                                                                  // procedure
	DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation)                                                                                    // procedure
	DragSourceSystemDragEnded()                                                                                                                           // procedure
	SetAccessibilityState(accessibilityState cefTypes.TCefState)                                                                                          // procedure
	SetAutoResizeEnabled(enabled bool, minSize TCefSize, maxSize TCefSize)                                                                                // procedure
	SetAudioMuted(mute bool)                                                                                                                              // procedure
}

// ICefBrowserHostRef Parent: ICefBrowserHost ICefBaseRefCountedRef
type ICefBrowserHostRef interface {
	ICefBrowserHost
	ICefBaseRefCountedRef
	AsIntfBrowserHost() uintptr
}

type TCefBrowserHostRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefBrowserHostRef) GetBrowser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCefBrowserHostRef) GetOpenerWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(4, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCefBrowserHostRef) HasView() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetRequestContext() (result ICefRequestContext) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) GetZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefBrowserHostRef) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) SendDevToolsMessage(message string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(9, m.Instance(), api.PasStr(message))
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(10, m.Instance(), uintptr(messageId), api.PasStr(method), base.GetObjectUintptr(params))
	return int32(r)
}

func (m *TCefBrowserHostRef) AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(observer), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) IsWindowRenderingDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetWindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TCefBrowserHostRef) GetVisibleNavigationEntry() (result ICefNavigationEntry) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefNavigationEntryRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) GetExtension() (result ICefExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefExtensionRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) IsBackgroundHost() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) IsAudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(19, m.Instance(), api.PasBool(forceClose))
}

func (m *TCefBrowserHostRef) SetFocus(focus bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(20, m.Instance(), api.PasBool(focus))
}

func (m *TCefBrowserHostRef) SetZoomLevel(zoomLevel float64) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&zoomLevel)))
}

func (m *TCefBrowserHostRef) RunFileDialog(mode cefTypes.TCefFileDialogMode, title string, defaultFilePath string, acceptFilters lcl.IStrings, callback IEngRunFileDialogCallback) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(22, m.Instance(), uintptr(mode), api.PasStr(title), api.PasStr(defaultFilePath), base.GetObjectUintptr(acceptFilters), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) StartDownload(url string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(23, m.Instance(), api.PasStr(url))
}

func (m *TCefBrowserHostRef) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback IEngDownloadImageCallback) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(24, m.Instance(), api.PasStr(imageUrl), api.PasBool(isFavicon), uintptr(maxImageSize), api.PasBool(bypassCache), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) Print() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(25, m.Instance())
}

func (m *TCefBrowserHostRef) PrintToPdf(path string, settings TCefPdfPrintSettings, callback IEngPdfPrintCallback) {
	if !m.IsValid() {
		return
	}
	settingsPtr := settings.ToPas()
	cefBrowserHostRefAPI().SysCallN(26, m.Instance(), api.PasStr(path), uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) Find(searchText string, forward bool, matchCase bool, findNext bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(27, m.Instance(), api.PasStr(searchText), api.PasBool(forward), api.PasBool(matchCase), api.PasBool(findNext))
}

func (m *TCefBrowserHostRef) StopFinding(clearSelection bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(28, m.Instance(), api.PasBool(clearSelection))
}

func (m *TCefBrowserHostRef) ShowDevTools(windowInfo TCefWindowInfo, client IEngClient, settings TCefBrowserSettings, inspectElementAt TCefPoint) {
	if !m.IsValid() {
		return
	}
	windowInfoPtr := windowInfo.ToPas()
	settingsPtr := settings.ToPas()
	cefBrowserHostRefAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(windowInfoPtr)), base.GetObjectUintptr(client), uintptr(base.UnsafePointer(settingsPtr)), uintptr(base.UnsafePointer(&inspectElementAt)))
}

func (m *TCefBrowserHostRef) CloseDevTools() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(30, m.Instance())
}

func (m *TCefBrowserHostRef) GetNavigationEntries(visitor IEngNavigationEntryVisitor, currentOnly bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(visitor), api.PasBool(currentOnly))
}

func (m *TCefBrowserHostRef) ReplaceMisspelling(word string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(32, m.Instance(), api.PasStr(word))
}

func (m *TCefBrowserHostRef) AddWordToDictionary(word string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(33, m.Instance(), api.PasStr(word))
}

func (m *TCefBrowserHostRef) WasResized() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(34, m.Instance())
}

func (m *TCefBrowserHostRef) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(35, m.Instance())
}

func (m *TCefBrowserHostRef) WasHidden(hidden bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(36, m.Instance(), api.PasBool(hidden))
}

func (m *TCefBrowserHostRef) Invalidate(type_ cefTypes.TCefPaintElementType) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(37, m.Instance(), uintptr(type_))
}

func (m *TCefBrowserHostRef) SendExternalBeginFrame() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(38, m.Instance())
}

func (m *TCefBrowserHostRef) SendKeyEvent(event TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(39, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(40, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(type_), api.PasBool(mouseUp), uintptr(clickCount))
}

func (m *TCefBrowserHostRef) SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(41, m.Instance(), uintptr(base.UnsafePointer(&event)), api.PasBool(mouseLeave))
}

func (m *TCefBrowserHostRef) SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(42, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCefBrowserHostRef) SendTouchEvent(event TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(43, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(44, m.Instance())
}

func (m *TCefBrowserHostRef) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(45, m.Instance())
}

func (m *TCefBrowserHostRef) SetWindowlessFrameRate(frameRate int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(46, m.Instance(), uintptr(frameRate))
}

func (m *TCefBrowserHostRef) IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(47, m.Instance(), api.PasStr(text), underlines.Instance(), uintptr(int32(underlines.Count())), uintptr(base.UnsafePointer(&replacementRange)), uintptr(base.UnsafePointer(&selectionRange)))
}

func (m *TCefBrowserHostRef) IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(48, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&replacementRange)), uintptr(relativeCursorPos))
}

func (m *TCefBrowserHostRef) IMEFinishComposingText(keepSelection bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(49, m.Instance(), api.PasBool(keepSelection))
}

func (m *TCefBrowserHostRef) IMECancelComposition() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(50, m.Instance())
}

func (m *TCefBrowserHostRef) DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(51, m.Instance(), base.GetObjectUintptr(dragData), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TCefBrowserHostRef) DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(52, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TCefBrowserHostRef) DragTargetDragLeave() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(53, m.Instance())
}

func (m *TCefBrowserHostRef) DragTargetDrop(event TCefMouseEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(54, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(55, m.Instance(), uintptr(X), uintptr(Y), uintptr(op))
}

func (m *TCefBrowserHostRef) DragSourceSystemDragEnded() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(56, m.Instance())
}

func (m *TCefBrowserHostRef) SetAccessibilityState(accessibilityState cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(57, m.Instance(), uintptr(accessibilityState))
}

func (m *TCefBrowserHostRef) SetAutoResizeEnabled(enabled bool, minSize TCefSize, maxSize TCefSize) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(58, m.Instance(), api.PasBool(enabled), uintptr(base.UnsafePointer(&minSize)), uintptr(base.UnsafePointer(&maxSize)))
}

func (m *TCefBrowserHostRef) SetAudioMuted(mute bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(59, m.Instance(), api.PasBool(mute))
}

func (m *TCefBrowserHostRef) AsIntfBrowserHost() uintptr {
	return m.GetIntfPointer(0)
}

// BrowserHostRef  is static instance
var BrowserHostRef _BrowserHostRefClass

// _BrowserHostRefClass is class type defined by TCefBrowserHostRef
type _BrowserHostRefClass uintptr

func (_BrowserHostRefClass) UnWrap(data uintptr) (result ICefBrowserHost) {
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(18, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserHostRef(resultPtr)
	return
}

// NewBrowserHostRef class constructor
func NewBrowserHostRef(data uintptr) ICefBrowserHostRef {
	var browserHostPtr uintptr // ICefBrowserHost
	r := cefBrowserHostRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&browserHostPtr)))
	ret := AsCefBrowserHostRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserHostPtr)
	}
	return ret
}

var (
	cefBrowserHostRefOnce   base.Once
	cefBrowserHostRefImport *imports.Imports = nil
)

func cefBrowserHostRefAPI() *imports.Imports {
	cefBrowserHostRefOnce.Do(func() {
		cefBrowserHostRefImport = api.NewDefaultImports()
		cefBrowserHostRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserHostRef_Create", 0), // constructor NewBrowserHostRef
			/* 1 */ imports.NewTable("TCefBrowserHostRef_GetBrowser", 0), // function GetBrowser
			/* 2 */ imports.NewTable("TCefBrowserHostRef_TryCloseBrowser", 0), // function TryCloseBrowser
			/* 3 */ imports.NewTable("TCefBrowserHostRef_GetWindowHandle", 0), // function GetWindowHandle
			/* 4 */ imports.NewTable("TCefBrowserHostRef_GetOpenerWindowHandle", 0), // function GetOpenerWindowHandle
			/* 5 */ imports.NewTable("TCefBrowserHostRef_HasView", 0), // function HasView
			/* 6 */ imports.NewTable("TCefBrowserHostRef_GetRequestContext", 0), // function GetRequestContext
			/* 7 */ imports.NewTable("TCefBrowserHostRef_GetZoomLevel", 0), // function GetZoomLevel
			/* 8 */ imports.NewTable("TCefBrowserHostRef_HasDevTools", 0), // function HasDevTools
			/* 9 */ imports.NewTable("TCefBrowserHostRef_SendDevToolsMessage", 0), // function SendDevToolsMessage
			/* 10 */ imports.NewTable("TCefBrowserHostRef_ExecuteDevToolsMethod", 0), // function ExecuteDevToolsMethod
			/* 11 */ imports.NewTable("TCefBrowserHostRef_AddDevToolsMessageObserver", 0), // function AddDevToolsMessageObserver
			/* 12 */ imports.NewTable("TCefBrowserHostRef_IsWindowRenderingDisabled", 0), // function IsWindowRenderingDisabled
			/* 13 */ imports.NewTable("TCefBrowserHostRef_GetWindowlessFrameRate", 0), // function GetWindowlessFrameRate
			/* 14 */ imports.NewTable("TCefBrowserHostRef_GetVisibleNavigationEntry", 0), // function GetVisibleNavigationEntry
			/* 15 */ imports.NewTable("TCefBrowserHostRef_GetExtension", 0), // function GetExtension
			/* 16 */ imports.NewTable("TCefBrowserHostRef_IsBackgroundHost", 0), // function IsBackgroundHost
			/* 17 */ imports.NewTable("TCefBrowserHostRef_IsAudioMuted", 0), // function IsAudioMuted
			/* 18 */ imports.NewTable("TCefBrowserHostRef_UnWrap", 0), // static function UnWrap
			/* 19 */ imports.NewTable("TCefBrowserHostRef_CloseBrowser", 0), // procedure CloseBrowser
			/* 20 */ imports.NewTable("TCefBrowserHostRef_SetFocus", 0), // procedure SetFocus
			/* 21 */ imports.NewTable("TCefBrowserHostRef_SetZoomLevel", 0), // procedure SetZoomLevel
			/* 22 */ imports.NewTable("TCefBrowserHostRef_RunFileDialog", 0), // procedure RunFileDialog
			/* 23 */ imports.NewTable("TCefBrowserHostRef_StartDownload", 0), // procedure StartDownload
			/* 24 */ imports.NewTable("TCefBrowserHostRef_DownloadImage", 0), // procedure DownloadImage
			/* 25 */ imports.NewTable("TCefBrowserHostRef_Print", 0), // procedure Print
			/* 26 */ imports.NewTable("TCefBrowserHostRef_PrintToPdf", 0), // procedure PrintToPdf
			/* 27 */ imports.NewTable("TCefBrowserHostRef_Find", 0), // procedure Find
			/* 28 */ imports.NewTable("TCefBrowserHostRef_StopFinding", 0), // procedure StopFinding
			/* 29 */ imports.NewTable("TCefBrowserHostRef_ShowDevTools", 0), // procedure ShowDevTools
			/* 30 */ imports.NewTable("TCefBrowserHostRef_CloseDevTools", 0), // procedure CloseDevTools
			/* 31 */ imports.NewTable("TCefBrowserHostRef_GetNavigationEntries", 0), // procedure GetNavigationEntries
			/* 32 */ imports.NewTable("TCefBrowserHostRef_ReplaceMisspelling", 0), // procedure ReplaceMisspelling
			/* 33 */ imports.NewTable("TCefBrowserHostRef_AddWordToDictionary", 0), // procedure AddWordToDictionary
			/* 34 */ imports.NewTable("TCefBrowserHostRef_WasResized", 0), // procedure WasResized
			/* 35 */ imports.NewTable("TCefBrowserHostRef_NotifyScreenInfoChanged", 0), // procedure NotifyScreenInfoChanged
			/* 36 */ imports.NewTable("TCefBrowserHostRef_WasHidden", 0), // procedure WasHidden
			/* 37 */ imports.NewTable("TCefBrowserHostRef_Invalidate", 0), // procedure Invalidate
			/* 38 */ imports.NewTable("TCefBrowserHostRef_SendExternalBeginFrame", 0), // procedure SendExternalBeginFrame
			/* 39 */ imports.NewTable("TCefBrowserHostRef_SendKeyEvent", 0), // procedure SendKeyEvent
			/* 40 */ imports.NewTable("TCefBrowserHostRef_SendMouseClickEvent", 0), // procedure SendMouseClickEvent
			/* 41 */ imports.NewTable("TCefBrowserHostRef_SendMouseMoveEvent", 0), // procedure SendMouseMoveEvent
			/* 42 */ imports.NewTable("TCefBrowserHostRef_SendMouseWheelEvent", 0), // procedure SendMouseWheelEvent
			/* 43 */ imports.NewTable("TCefBrowserHostRef_SendTouchEvent", 0), // procedure SendTouchEvent
			/* 44 */ imports.NewTable("TCefBrowserHostRef_SendCaptureLostEvent", 0), // procedure SendCaptureLostEvent
			/* 45 */ imports.NewTable("TCefBrowserHostRef_NotifyMoveOrResizeStarted", 0), // procedure NotifyMoveOrResizeStarted
			/* 46 */ imports.NewTable("TCefBrowserHostRef_SetWindowlessFrameRate", 0), // procedure SetWindowlessFrameRate
			/* 47 */ imports.NewTable("TCefBrowserHostRef_IMESetComposition", 0), // procedure IMESetComposition
			/* 48 */ imports.NewTable("TCefBrowserHostRef_IMECommitText", 0), // procedure IMECommitText
			/* 49 */ imports.NewTable("TCefBrowserHostRef_IMEFinishComposingText", 0), // procedure IMEFinishComposingText
			/* 50 */ imports.NewTable("TCefBrowserHostRef_IMECancelComposition", 0), // procedure IMECancelComposition
			/* 51 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragEnter", 0), // procedure DragTargetDragEnter
			/* 52 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragOver", 0), // procedure DragTargetDragOver
			/* 53 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragLeave", 0), // procedure DragTargetDragLeave
			/* 54 */ imports.NewTable("TCefBrowserHostRef_DragTargetDrop", 0), // procedure DragTargetDrop
			/* 55 */ imports.NewTable("TCefBrowserHostRef_DragSourceEndedAt", 0), // procedure DragSourceEndedAt
			/* 56 */ imports.NewTable("TCefBrowserHostRef_DragSourceSystemDragEnded", 0), // procedure DragSourceSystemDragEnded
			/* 57 */ imports.NewTable("TCefBrowserHostRef_SetAccessibilityState", 0), // procedure SetAccessibilityState
			/* 58 */ imports.NewTable("TCefBrowserHostRef_SetAutoResizeEnabled", 0), // procedure SetAutoResizeEnabled
			/* 59 */ imports.NewTable("TCefBrowserHostRef_SetAudioMuted", 0), // procedure SetAudioMuted
		}
	})
	return cefBrowserHostRefImport
}
