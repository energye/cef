//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICefBrowserHost Parent: ICefBaseRefCounted
type ICefBrowserHost interface {
	ICefBaseRefCounted
	IMESetComposition(text string, underlines TCefCompositionUnderlineDynArray, replacementrange, selectionrange *TCefRange)
	GetBrowser() ICefBrowser                                                                                                          // function
	TryCloseBrowser() bool                                                                                                            // function
	GetWindowHandle() TCefWindowHandle                                                                                                // function
	GetOpenerWindowHandle() TCefWindowHandle                                                                                          // function
	HasView() bool                                                                                                                    // function
	GetClient() ICefClient                                                                                                            // function
	GetRequestContext() ICefRequestContext                                                                                            // function
	CanZoom(command TCefZoomCommand) bool                                                                                             // function
	GetDefaultZoomLevel() (resultFloat64 float64)                                                                                     // function
	GetZoomLevel() (resultFloat64 float64)                                                                                            // function
	HasDevTools() bool                                                                                                                // function
	SendDevToolsMessage(message string) bool                                                                                          // function
	ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32                                           // function
	AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration                                                 // function
	IsWindowRenderingDisabled() bool                                                                                                  // function
	GetWindowlessFrameRate() int32                                                                                                    // function
	GetVisibleNavigationEntry() ICefNavigationEntry                                                                                   // function
	GetExtension() ICefExtension                                                                                                      // function
	IsBackgroundHost() bool                                                                                                           // function
	IsAudioMuted() bool                                                                                                               // function
	IsFullscreen() bool                                                                                                               // function
	CloseBrowser(forceClose bool)                                                                                                     // procedure
	SetFocus(focus bool)                                                                                                              // procedure
	Zoom(command TCefZoomCommand)                                                                                                     // procedure
	SetZoomLevel(zoomLevel float64)                                                                                                   // procedure
	RunFileDialog(mode TCefFileDialogMode, title, defaultFilePath string, acceptFilters IStrings, callback ICefRunFileDialogCallback) // procedure
	StartDownload(url string)                                                                                                         // procedure
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback ICefDownloadImageCallback)         // procedure
	Print()                                                                                                                           // procedure
	PrintToPdf(path string, settings *TCefPdfPrintSettings, callback ICefPdfPrintCallback)                                            // procedure
	Find(searchText string, forward, matchCase, findNext bool)                                                                        // procedure
	StopFinding(clearSelection bool)                                                                                                  // procedure
	ShowDevTools(windowInfo *TCefWindowInfo, client ICefClient, settings *TCefBrowserSettings, inspectElementAt *TCefPoint)           // procedure
	CloseDevTools()                                                                                                                   // procedure
	GetNavigationEntries(visitor ICefNavigationEntryVisitor, currentOnly bool)                                                        // procedure
	ReplaceMisspelling(word string)                                                                                                   // procedure
	AddWordToDictionary(word string)                                                                                                  // procedure
	WasResized()                                                                                                                      // procedure
	NotifyScreenInfoChanged()                                                                                                         // procedure
	WasHidden(hidden bool)                                                                                                            // procedure
	Invalidate(type_ TCefPaintElementType)                                                                                            // procedure
	SendExternalBeginFrame()                                                                                                          // procedure
	SendKeyEvent(event *TCefKeyEvent)                                                                                                 // procedure
	SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32)                             // procedure
	SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool)                                                                        // procedure
	SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32)                                                                  // procedure
	SendTouchEvent(event *TCefTouchEvent)                                                                                             // procedure
	SendCaptureLostEvent()                                                                                                            // procedure
	NotifyMoveOrResizeStarted()                                                                                                       // procedure
	SetWindowlessFrameRate(frameRate int32)                                                                                           // procedure
	IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32)                                                  // procedure
	IMEFinishComposingText(keepselection bool)                                                                                        // procedure
	IMECancelComposition()                                                                                                            // procedure
	DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations)                                  // procedure
	DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations)                                                          // procedure
	DragTargetDragLeave()                                                                                                             // procedure
	DragTargetDrop(event *TCefMouseEvent)                                                                                             // procedure
	DragSourceEndedAt(x, y int32, op TCefDragOperation)                                                                               // procedure
	DragSourceSystemDragEnded()                                                                                                       // procedure
	SetAccessibilityState(accessibilityState TCefState)                                                                               // procedure
	SetAutoResizeEnabled(enabled bool, minsize, maxsize *TCefSize)                                                                    // procedure
	SetAudioMuted(mute bool)                                                                                                          // procedure
	ExitFullscreen(willcauseresize bool)                                                                                              // procedure
}

// TCefBrowserHost Parent: TCefBaseRefCounted
type TCefBrowserHost struct {
	TCefBaseRefCounted
}

// BrowserHostRef -> ICefBrowserHost
var BrowserHostRef browserHost

// browserHost TCefBrowserHost Ref
type browserHost uintptr

func (m *browserHost) UnWrap(data uintptr) ICefBrowserHost {
	var resultCefBrowserHost uintptr
	browserHostImportAPI().SysCallN(60, uintptr(data), uintptr(unsafePointer(&resultCefBrowserHost)))
	return AsCefBrowserHost(resultCefBrowserHost)
}

func (m *TCefBrowserHost) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	browserHostImportAPI().SysCallN(15, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefBrowserHost) TryCloseBrowser() bool {
	r1 := browserHostImportAPI().SysCallN(59, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetWindowHandle() TCefWindowHandle {
	r1 := browserHostImportAPI().SysCallN(23, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefBrowserHost) GetOpenerWindowHandle() TCefWindowHandle {
	r1 := browserHostImportAPI().SysCallN(20, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefBrowserHost) HasView() bool {
	r1 := browserHostImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetClient() ICefClient {
	var resultCefClient uintptr
	browserHostImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(&resultCefClient)))
	return AsCefClient(resultCefClient)
}

func (m *TCefBrowserHost) GetRequestContext() ICefRequestContext {
	var resultCefRequestContext uintptr
	browserHostImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TCefBrowserHost) CanZoom(command TCefZoomCommand) bool {
	r1 := browserHostImportAPI().SysCallN(2, m.Instance(), uintptr(command))
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetDefaultZoomLevel() (resultFloat64 float64) {
	browserHostImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefBrowserHost) GetZoomLevel() (resultFloat64 float64) {
	browserHostImportAPI().SysCallN(25, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefBrowserHost) HasDevTools() bool {
	r1 := browserHostImportAPI().SysCallN(26, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) SendDevToolsMessage(message string) bool {
	r1 := browserHostImportAPI().SysCallN(43, m.Instance(), PascalStr(message))
	return GoBool(r1)
}

func (m *TCefBrowserHost) ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32 {
	r1 := browserHostImportAPI().SysCallN(12, m.Instance(), uintptr(messageid), PascalStr(method), GetObjectUintptr(params))
	return int32(r1)
}

func (m *TCefBrowserHost) AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration {
	var resultCefRegistration uintptr
	browserHostImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(observer), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TCefBrowserHost) IsWindowRenderingDisabled() bool {
	r1 := browserHostImportAPI().SysCallN(35, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetWindowlessFrameRate() int32 {
	r1 := browserHostImportAPI().SysCallN(24, m.Instance())
	return int32(r1)
}

func (m *TCefBrowserHost) GetVisibleNavigationEntry() ICefNavigationEntry {
	var resultCefNavigationEntry uintptr
	browserHostImportAPI().SysCallN(22, m.Instance(), uintptr(unsafePointer(&resultCefNavigationEntry)))
	return AsCefNavigationEntry(resultCefNavigationEntry)
}

func (m *TCefBrowserHost) GetExtension() ICefExtension {
	var resultCefExtension uintptr
	browserHostImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TCefBrowserHost) IsBackgroundHost() bool {
	r1 := browserHostImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) IsAudioMuted() bool {
	r1 := browserHostImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) IsFullscreen() bool {
	r1 := browserHostImportAPI().SysCallN(34, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) CloseBrowser(forceClose bool) {
	browserHostImportAPI().SysCallN(3, m.Instance(), PascalBool(forceClose))
}

func (m *TCefBrowserHost) SetFocus(focus bool) {
	browserHostImportAPI().SysCallN(53, m.Instance(), PascalBool(focus))
}

func (m *TCefBrowserHost) Zoom(command TCefZoomCommand) {
	browserHostImportAPI().SysCallN(63, m.Instance(), uintptr(command))
}

func (m *TCefBrowserHost) SetZoomLevel(zoomLevel float64) {
	browserHostImportAPI().SysCallN(55, m.Instance(), uintptr(unsafePointer(&zoomLevel)))
}

func (m *TCefBrowserHost) RunFileDialog(mode TCefFileDialogMode, title, defaultFilePath string, acceptFilters IStrings, callback ICefRunFileDialogCallback) {
	browserHostImportAPI().SysCallN(41, m.Instance(), uintptr(mode), PascalStr(title), PascalStr(defaultFilePath), GetObjectUintptr(acceptFilters), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) StartDownload(url string) {
	browserHostImportAPI().SysCallN(57, m.Instance(), PascalStr(url))
}

func (m *TCefBrowserHost) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback ICefDownloadImageCallback) {
	browserHostImportAPI().SysCallN(5, m.Instance(), PascalStr(imageUrl), PascalBool(isFavicon), uintptr(maxImageSize), PascalBool(bypassCache), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) Print() {
	browserHostImportAPI().SysCallN(38, m.Instance())
}

func (m *TCefBrowserHost) PrintToPdf(path string, settings *TCefPdfPrintSettings, callback ICefPdfPrintCallback) {
	browserHostImportAPI().SysCallN(39, m.Instance(), PascalStr(path), uintptr(unsafePointer(settings)), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) Find(searchText string, forward, matchCase, findNext bool) {
	browserHostImportAPI().SysCallN(14, m.Instance(), PascalStr(searchText), PascalBool(forward), PascalBool(matchCase), PascalBool(findNext))
}

func (m *TCefBrowserHost) StopFinding(clearSelection bool) {
	browserHostImportAPI().SysCallN(58, m.Instance(), PascalBool(clearSelection))
}

func (m *TCefBrowserHost) ShowDevTools(windowInfo *TCefWindowInfo, client ICefClient, settings *TCefBrowserSettings, inspectElementAt *TCefPoint) {
	inArgs0 := windowInfo.Pointer()
	browserHostImportAPI().SysCallN(56, m.Instance(), uintptr(unsafePointer(inArgs0)), GetObjectUintptr(client), uintptr(unsafePointer(settings)), uintptr(unsafePointer(inspectElementAt)))
}

func (m *TCefBrowserHost) CloseDevTools() {
	browserHostImportAPI().SysCallN(4, m.Instance())
}

func (m *TCefBrowserHost) GetNavigationEntries(visitor ICefNavigationEntryVisitor, currentOnly bool) {
	browserHostImportAPI().SysCallN(19, m.Instance(), GetObjectUintptr(visitor), PascalBool(currentOnly))
}

func (m *TCefBrowserHost) ReplaceMisspelling(word string) {
	browserHostImportAPI().SysCallN(40, m.Instance(), PascalStr(word))
}

func (m *TCefBrowserHost) AddWordToDictionary(word string) {
	browserHostImportAPI().SysCallN(1, m.Instance(), PascalStr(word))
}

func (m *TCefBrowserHost) WasResized() {
	browserHostImportAPI().SysCallN(62, m.Instance())
}

func (m *TCefBrowserHost) NotifyScreenInfoChanged() {
	browserHostImportAPI().SysCallN(37, m.Instance())
}

func (m *TCefBrowserHost) WasHidden(hidden bool) {
	browserHostImportAPI().SysCallN(61, m.Instance(), PascalBool(hidden))
}

func (m *TCefBrowserHost) Invalidate(type_ TCefPaintElementType) {
	browserHostImportAPI().SysCallN(31, m.Instance(), uintptr(type_))
}

func (m *TCefBrowserHost) SendExternalBeginFrame() {
	browserHostImportAPI().SysCallN(44, m.Instance())
}

func (m *TCefBrowserHost) SendKeyEvent(event *TCefKeyEvent) {
	browserHostImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	browserHostImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(event)), uintptr(type_), PascalBool(mouseUp), uintptr(clickCount))
}

func (m *TCefBrowserHost) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	browserHostImportAPI().SysCallN(47, m.Instance(), uintptr(unsafePointer(event)), PascalBool(mouseLeave))
}

func (m *TCefBrowserHost) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	browserHostImportAPI().SysCallN(48, m.Instance(), uintptr(unsafePointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCefBrowserHost) SendTouchEvent(event *TCefTouchEvent) {
	browserHostImportAPI().SysCallN(49, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) SendCaptureLostEvent() {
	browserHostImportAPI().SysCallN(42, m.Instance())
}

func (m *TCefBrowserHost) NotifyMoveOrResizeStarted() {
	browserHostImportAPI().SysCallN(36, m.Instance())
}

func (m *TCefBrowserHost) SetWindowlessFrameRate(frameRate int32) {
	browserHostImportAPI().SysCallN(54, m.Instance(), uintptr(frameRate))
}

func (m *TCefBrowserHost) IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32) {
	browserHostImportAPI().SysCallN(29, m.Instance(), PascalStr(text), uintptr(unsafePointer(replacementrange)), uintptr(relativecursorpos))
}

func (m *TCefBrowserHost) IMEFinishComposingText(keepselection bool) {
	browserHostImportAPI().SysCallN(30, m.Instance(), PascalBool(keepselection))
}

func (m *TCefBrowserHost) IMECancelComposition() {
	browserHostImportAPI().SysCallN(28, m.Instance())
}

func (m *TCefBrowserHost) DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations) {
	browserHostImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(dragData), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TCefBrowserHost) DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations) {
	browserHostImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TCefBrowserHost) DragTargetDragLeave() {
	browserHostImportAPI().SysCallN(9, m.Instance())
}

func (m *TCefBrowserHost) DragTargetDrop(event *TCefMouseEvent) {
	browserHostImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) DragSourceEndedAt(x, y int32, op TCefDragOperation) {
	browserHostImportAPI().SysCallN(6, m.Instance(), uintptr(x), uintptr(y), uintptr(op))
}

func (m *TCefBrowserHost) DragSourceSystemDragEnded() {
	browserHostImportAPI().SysCallN(7, m.Instance())
}

func (m *TCefBrowserHost) SetAccessibilityState(accessibilityState TCefState) {
	browserHostImportAPI().SysCallN(50, m.Instance(), uintptr(accessibilityState))
}

func (m *TCefBrowserHost) SetAutoResizeEnabled(enabled bool, minsize, maxsize *TCefSize) {
	browserHostImportAPI().SysCallN(52, m.Instance(), PascalBool(enabled), uintptr(unsafePointer(minsize)), uintptr(unsafePointer(maxsize)))
}

func (m *TCefBrowserHost) SetAudioMuted(mute bool) {
	browserHostImportAPI().SysCallN(51, m.Instance(), PascalBool(mute))
}

func (m *TCefBrowserHost) ExitFullscreen(willcauseresize bool) {
	browserHostImportAPI().SysCallN(13, m.Instance(), PascalBool(willcauseresize))
}

var (
	browserHostImport       *imports.Imports = nil
	browserHostImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBrowserHost_AddDevToolsMessageObserver", 0),
		/*1*/ imports.NewTable("CefBrowserHost_AddWordToDictionary", 0),
		/*2*/ imports.NewTable("CefBrowserHost_CanZoom", 0),
		/*3*/ imports.NewTable("CefBrowserHost_CloseBrowser", 0),
		/*4*/ imports.NewTable("CefBrowserHost_CloseDevTools", 0),
		/*5*/ imports.NewTable("CefBrowserHost_DownloadImage", 0),
		/*6*/ imports.NewTable("CefBrowserHost_DragSourceEndedAt", 0),
		/*7*/ imports.NewTable("CefBrowserHost_DragSourceSystemDragEnded", 0),
		/*8*/ imports.NewTable("CefBrowserHost_DragTargetDragEnter", 0),
		/*9*/ imports.NewTable("CefBrowserHost_DragTargetDragLeave", 0),
		/*10*/ imports.NewTable("CefBrowserHost_DragTargetDragOver", 0),
		/*11*/ imports.NewTable("CefBrowserHost_DragTargetDrop", 0),
		/*12*/ imports.NewTable("CefBrowserHost_ExecuteDevToolsMethod", 0),
		/*13*/ imports.NewTable("CefBrowserHost_ExitFullscreen", 0),
		/*14*/ imports.NewTable("CefBrowserHost_Find", 0),
		/*15*/ imports.NewTable("CefBrowserHost_GetBrowser", 0),
		/*16*/ imports.NewTable("CefBrowserHost_GetClient", 0),
		/*17*/ imports.NewTable("CefBrowserHost_GetDefaultZoomLevel", 0),
		/*18*/ imports.NewTable("CefBrowserHost_GetExtension", 0),
		/*19*/ imports.NewTable("CefBrowserHost_GetNavigationEntries", 0),
		/*20*/ imports.NewTable("CefBrowserHost_GetOpenerWindowHandle", 0),
		/*21*/ imports.NewTable("CefBrowserHost_GetRequestContext", 0),
		/*22*/ imports.NewTable("CefBrowserHost_GetVisibleNavigationEntry", 0),
		/*23*/ imports.NewTable("CefBrowserHost_GetWindowHandle", 0),
		/*24*/ imports.NewTable("CefBrowserHost_GetWindowlessFrameRate", 0),
		/*25*/ imports.NewTable("CefBrowserHost_GetZoomLevel", 0),
		/*26*/ imports.NewTable("CefBrowserHost_HasDevTools", 0),
		/*27*/ imports.NewTable("CefBrowserHost_HasView", 0),
		/*28*/ imports.NewTable("CefBrowserHost_IMECancelComposition", 0),
		/*29*/ imports.NewTable("CefBrowserHost_IMECommitText", 0),
		/*30*/ imports.NewTable("CefBrowserHost_IMEFinishComposingText", 0),
		/*31*/ imports.NewTable("CefBrowserHost_Invalidate", 0),
		/*32*/ imports.NewTable("CefBrowserHost_IsAudioMuted", 0),
		/*33*/ imports.NewTable("CefBrowserHost_IsBackgroundHost", 0),
		/*34*/ imports.NewTable("CefBrowserHost_IsFullscreen", 0),
		/*35*/ imports.NewTable("CefBrowserHost_IsWindowRenderingDisabled", 0),
		/*36*/ imports.NewTable("CefBrowserHost_NotifyMoveOrResizeStarted", 0),
		/*37*/ imports.NewTable("CefBrowserHost_NotifyScreenInfoChanged", 0),
		/*38*/ imports.NewTable("CefBrowserHost_Print", 0),
		/*39*/ imports.NewTable("CefBrowserHost_PrintToPdf", 0),
		/*40*/ imports.NewTable("CefBrowserHost_ReplaceMisspelling", 0),
		/*41*/ imports.NewTable("CefBrowserHost_RunFileDialog", 0),
		/*42*/ imports.NewTable("CefBrowserHost_SendCaptureLostEvent", 0),
		/*43*/ imports.NewTable("CefBrowserHost_SendDevToolsMessage", 0),
		/*44*/ imports.NewTable("CefBrowserHost_SendExternalBeginFrame", 0),
		/*45*/ imports.NewTable("CefBrowserHost_SendKeyEvent", 0),
		/*46*/ imports.NewTable("CefBrowserHost_SendMouseClickEvent", 0),
		/*47*/ imports.NewTable("CefBrowserHost_SendMouseMoveEvent", 0),
		/*48*/ imports.NewTable("CefBrowserHost_SendMouseWheelEvent", 0),
		/*49*/ imports.NewTable("CefBrowserHost_SendTouchEvent", 0),
		/*50*/ imports.NewTable("CefBrowserHost_SetAccessibilityState", 0),
		/*51*/ imports.NewTable("CefBrowserHost_SetAudioMuted", 0),
		/*52*/ imports.NewTable("CefBrowserHost_SetAutoResizeEnabled", 0),
		/*53*/ imports.NewTable("CefBrowserHost_SetFocus", 0),
		/*54*/ imports.NewTable("CefBrowserHost_SetWindowlessFrameRate", 0),
		/*55*/ imports.NewTable("CefBrowserHost_SetZoomLevel", 0),
		/*56*/ imports.NewTable("CefBrowserHost_ShowDevTools", 0),
		/*57*/ imports.NewTable("CefBrowserHost_StartDownload", 0),
		/*58*/ imports.NewTable("CefBrowserHost_StopFinding", 0),
		/*59*/ imports.NewTable("CefBrowserHost_TryCloseBrowser", 0),
		/*60*/ imports.NewTable("CefBrowserHost_UnWrap", 0),
		/*61*/ imports.NewTable("CefBrowserHost_WasHidden", 0),
		/*62*/ imports.NewTable("CefBrowserHost_WasResized", 0),
		/*63*/ imports.NewTable("CefBrowserHost_Zoom", 0),
	}
)

func browserHostImportAPI() *imports.Imports {
	if browserHostImport == nil {
		browserHostImport = NewDefaultImports()
		browserHostImport.SetImportTable(browserHostImportTables)
		browserHostImportTables = nil
	}
	return browserHostImport
}
