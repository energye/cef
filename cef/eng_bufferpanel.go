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

// IBufferPanel Parent: ICustomPanel
//
//	TBufferPanel is used by VCL and LCL applications with browsers in OSR mode
//	to draw the browser contents. See the SimpleOSRBrowser demo for more details.
type IBufferPanel interface {
	ICustomPanel
	// ChangeCompositionRange
	//  Calls ChangeCompositionRange in the IME handler.
	ChangeCompositionRange(selectionrange *TCefRange, characterbounds TCefRectDynArray)
	// ScanlineSize
	//  Returns the scanline size.
	ScanlineSize() int32 // property
	// BufferWidth
	//  Image width.
	BufferWidth() int32 // property
	// BufferHeight
	//  Image height.
	BufferHeight() int32 // property
	// BufferBits
	//  Returns a pointer to the buffer that stores the image.
	BufferBits() uintptr // property
	// ScreenScale
	//  Returns the screen scale.
	ScreenScale() (resultFloat32 float32) // property
	// ForcedDeviceScaleFactor
	//  Screen scale value used instead of the real one.
	ForcedDeviceScaleFactor() (resultFloat32 float32) // property
	// SetForcedDeviceScaleFactor Set ForcedDeviceScaleFactor
	SetForcedDeviceScaleFactor(AValue float32) // property
	// MustInitBuffer
	//  Clear the background image before copying the original buffer contents.
	MustInitBuffer() bool // property
	// SetMustInitBuffer Set MustInitBuffer
	SetMustInitBuffer(AValue bool) // property
	// Buffer
	//  Background bitmap.
	Buffer() IBitmap // property
	// OrigBuffer
	//  Copy of the raw main bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
	//  OrigBuffer will be transferred to the bitmap buffer before copying the bitmap buffer
	//  to the panel.
	OrigBuffer() ICEFBitmapBitBuffer // property
	// OrigBufferWidth
	//  Image width of the raw main bitmap buffer copy.
	OrigBufferWidth() int32 // property
	// OrigBufferHeight
	//  Image height of the raw main bitmap buffer copy.
	OrigBufferHeight() int32 // property
	// OrigPopupBuffer
	//  Copy of the raw popup bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
	OrigPopupBuffer() ICEFBitmapBitBuffer // property
	// OrigPopupBufferWidth
	//  Image width of the raw popup bitmap buffer copy.
	OrigPopupBufferWidth() int32 // property
	// OrigPopupBufferHeight
	//  Image height of the raw popup bitmap buffer copy.
	OrigPopupBufferHeight() int32 // property
	// OrigPopupBufferBits
	//  Returns a pointer to the raw popup bitmap buffer copye.
	OrigPopupBufferBits() uintptr // property
	// OrigPopupScanlineSize
	//  Returns the scanline size of the raw popup bitmap buffer copy.
	OrigPopupScanlineSize() int32 // property
	// ParentFormHandle
	//  Returns the handle of the parent form.
	ParentFormHandle() TCefWindowHandle // property
	// ParentForm
	//  Returns the parent form.
	ParentForm() ICustomForm // property
	// Transparent
	//  Set Transparent to True to use a WS_EX_TRANSPARENT window style in the panel
	//  and to call AlphaBlend in order to transfer the web contents from the bitmap
	//  buffer to the panel.
	//  If this property is False then BitBlt is used to transfer the web contents
	//  from the bitmap buffer to the panel.
	Transparent() bool // property
	// SetTransparent Set Transparent
	SetTransparent(AValue bool) // property
	// CopyOriginalBuffer
	//  When CopyOriginalBuffer is True then OrigBuffer will be used internally to copy of
	//  the raw main bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
	//  OrigBuffer will be transferred to the bitmap buffer before copying the buffer to the panel.
	//  This is necessary in GTK applications in order to avoid handling bitmaps in background
	//  threads.
	CopyOriginalBuffer() bool // property
	// SetCopyOriginalBuffer Set CopyOriginalBuffer
	SetCopyOriginalBuffer(AValue bool) // property
	DragCursor() TCursor               // property
	SetDragCursor(AValue TCursor)      // property
	DragKind() TDragKind               // property
	SetDragKind(AValue TDragKind)      // property
	DragMode() TDragMode               // property
	SetDragMode(AValue TDragMode)      // property
	ParentFont() bool                  // property
	SetParentFont(AValue bool)         // property
	ParentShowHint() bool              // property
	SetParentShowHint(AValue bool)     // property
	// SaveToFile
	//  Save the visible web contents as a bitmap file.
	SaveToFile(aFilename string) bool // function
	// InvalidatePanel
	//  Invalidate this panel.
	InvalidatePanel() bool // function
	// BeginBufferDraw
	//  Acquires the synchronization object before drawing into the background bitmap.
	BeginBufferDraw() bool // function
	// UpdateBufferDimensions
	//  Update the background bitmap size.
	UpdateBufferDimensions(aWidth, aHeight int32) bool // function
	// UpdateOrigBufferDimensions
	//  Update the image size of the original buffer copy.
	UpdateOrigBufferDimensions(aWidth, aHeight int32) bool // function
	// UpdateOrigPopupBufferDimensions
	//  Update the popup image size of the original buffer copy.
	UpdateOrigPopupBufferDimensions(aWidth, aHeight int32) bool // function
	// BufferIsResized
	//  Check if the background image buffers have the same dimensions as this panel. Returns true if they have the same size.
	BufferIsResized(aUseMutex bool) bool // function
	// EndBufferDraw
	//  Releases the synchronization object after drawing into the background bitmap.
	EndBufferDraw() // procedure
	// BufferDraw
	//  Draws aBitmap into the background bitmap buffer at the specified coordinates.
	//  <param name="x">x coordinate where the bitmap will be drawn.</param>
	//  <param name="y">y coordinate where the bitmap will be drawn.</param>
	//  <param name="aBitmap">Bitmap that will be drawn into the background bitmap.</param>
	BufferDraw(x, y int32, aBitmap IBitmap) // procedure
	// BufferDraw1
	//  Draws a part of aBitmap into the background bitmap buffer at the specified rectangle.
	//  <param name="aBitmap">Bitmap that will be drawn into the background bitmap.</param>
	//  <param name="aSrcRect">Rectangle that defines the area of aBitmap that will be drawn into the background bitmap.</param>
	//  <param name="aDstRect">Rectangle that defines the area of the background bitmap where aBitmap will be drawn.</param>
	BufferDraw1(aBitmap IBitmap, aSrcRect, aDstRect *TRect) // procedure
	// UpdateDeviceScaleFactor
	//  Update the FDeviceScaleFactor value with the current scale.
	UpdateDeviceScaleFactor() // procedure
	// CreateIMEHandler
	//  Creates the IME handler.
	CreateIMEHandler() // procedure
	// DrawOrigPopupBuffer
	//  Copy the contents from the original popup buffer copy to the main buffer copy.
	DrawOrigPopupBuffer(aSrcRect, aDstRect *TRect) // procedure
	// SetOnIMECancelComposition
	//  Event triggered when a WM_IME_ENDCOMPOSITION message is received because
	//  the IME ended composition.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/intl/wm-ime-endcomposition">See the WM_IME_ENDCOMPOSITION article.</a>
	SetOnIMECancelComposition(fn TNotify) // property event
	// SetOnIMECommitText
	//  Event triggered when a WM_IME_COMPOSITION message is received because
	//  the IME changed composition status as a result of a keystroke. This
	//  event is triggered after retrieving a composition result of the ongoing
	//  composition if it exists.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/intl/wm-ime-composition">See the WM_IME_COMPOSITION article.</a>
	SetOnIMECommitText(fn TOnIMECommitText) // property event
	// SetOnIMESetComposition
	//  Event triggered when a WM_IME_COMPOSITION message is received because
	//  the IME changed composition status as a result of a keystroke.
	//  This event is triggered after retrieving the current composition
	//  status of the ongoing composition.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/intl/wm-ime-composition">See the WM_IME_COMPOSITION article.</a>
	SetOnIMESetComposition(fn TOnIMESetComposition) // property event
	// SetOnCustomTouch
	//  Event triggered when a WM_TOUCH message is received. It notifies the
	//  window when one or more touch points, such as a finger or pen,
	//  touches a touch-sensitive digitizer surface.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/wintouch/wm-touchdown">See the WM_TOUCH article.</a>
	SetOnCustomTouch(fn TOnHandledMessage) // property event
	// SetOnPointerDown
	//  Event triggered when a WM_POINTERDOWN message is received.
	//  Posted when a pointer makes contact over the client area of a window.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/inputmsg/wm-pointerdown">See the WM_POINTERDOWN article.</a>
	SetOnPointerDown(fn TOnHandledMessage) // property event
	// SetOnPointerUp
	//  Event triggered when a WM_POINTERUP message is received.
	//  Posted when a pointer that made contact over the client area of a window breaks contact.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/inputmsg/wm-pointerup">See the WM_POINTERUP article.</a>
	SetOnPointerUp(fn TOnHandledMessage) // property event
	// SetOnPointerUpdate
	//  Event triggered when a WM_POINTERUPDATE message is received.
	//  Posted to provide an update on a pointer that made contact over the client area of a
	//  window or on a hovering uncaptured pointer over the client area of a window.
	//  <a href="https://learn.microsoft.com/en-us/windows/win32/inputmsg/wm-pointerupdate">See the WM_POINTERUPDATE article.</a>
	SetOnPointerUpdate(fn TOnHandledMessage) // property event
	// SetOnPaintParentBkg
	//  Event triggered before the AlphaBlend call that transfer the web contents from the
	//  bitmap buffer to the panel when the Transparent property is True.
	SetOnPaintParentBkg(fn TNotify)               // property event
	SetOnConstrainedResize(fn TConstrainedResize) // property event
	SetOnContextPopup(fn TContextPopup)           // property event
	SetOnDblClick(fn TNotify)                     // property event
	SetOnDragDrop(fn TDragDrop)                   // property event
	SetOnDragOver(fn TDragOver)                   // property event
	SetOnEndDock(fn TEndDrag)                     // property event
	SetOnEndDrag(fn TEndDrag)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfo)             // property event
	SetOnMouseDown(fn TMouse)                     // property event
	SetOnMouseMove(fn TMouseMove)                 // property event
	SetOnMouseUp(fn TMouse)                       // property event
	SetOnMouseWheel(fn TMouseWheel)               // property event
	SetOnStartDock(fn TStartDock)                 // property event
	SetOnStartDrag(fn TStartDrag)                 // property event
}

// TBufferPanel Parent: TCustomPanel
//
//	TBufferPanel is used by VCL and LCL applications with browsers in OSR mode
//	to draw the browser contents. See the SimpleOSRBrowser demo for more details.
type TBufferPanel struct {
	TCustomPanel
	iMECancelCompositionPtr uintptr
	iMECommitTextPtr        uintptr
	iMESetCompositionPtr    uintptr
	customTouchPtr          uintptr
	pointerDownPtr          uintptr
	pointerUpPtr            uintptr
	pointerUpdatePtr        uintptr
	paintParentBkgPtr       uintptr
	constrainedResizePtr    uintptr
	contextPopupPtr         uintptr
	dblClickPtr             uintptr
	dragDropPtr             uintptr
	dragOverPtr             uintptr
	endDockPtr              uintptr
	endDragPtr              uintptr
	getSiteInfoPtr          uintptr
	mouseDownPtr            uintptr
	mouseMovePtr            uintptr
	mouseUpPtr              uintptr
	mouseWheelPtr           uintptr
	startDockPtr            uintptr
	startDragPtr            uintptr
}

func NewBufferPanel(aOwner IComponent) IBufferPanel {
	r1 := bufferPanelImportAPI().SysCallN(11, GetObjectUintptr(aOwner))
	return AsBufferPanel(r1)
}

func (m *TBufferPanel) ScanlineSize() int32 {
	r1 := bufferPanelImportAPI().SysCallN(39, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) BufferWidth() int32 {
	r1 := bufferPanelImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) BufferHeight() int32 {
	r1 := bufferPanelImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) BufferBits() uintptr {
	r1 := bufferPanelImportAPI().SysCallN(4, m.Instance())
	return uintptr(r1)
}

func (m *TBufferPanel) ScreenScale() (resultFloat32 float32) {
	bufferPanelImportAPI().SysCallN(40, m.Instance(), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TBufferPanel) ForcedDeviceScaleFactor() (resultFloat32 float32) {
	bufferPanelImportAPI().SysCallN(19, 0, m.Instance(), uintptr(unsafePointer(&resultFloat32)), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TBufferPanel) SetForcedDeviceScaleFactor(AValue float32) {
	bufferPanelImportAPI().SysCallN(19, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TBufferPanel) MustInitBuffer() bool {
	r1 := bufferPanelImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBufferPanel) SetMustInitBuffer(AValue bool) {
	bufferPanelImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBufferPanel) Buffer() IBitmap {
	r1 := bufferPanelImportAPI().SysCallN(3, m.Instance())
	return AsBitmap(r1)
}

func (m *TBufferPanel) OrigBuffer() ICEFBitmapBitBuffer {
	var resultCEFBitmapBitBuffer uintptr
	bufferPanelImportAPI().SysCallN(25, m.Instance(), uintptr(unsafePointer(&resultCEFBitmapBitBuffer)))
	return AsCEFBitmapBitBuffer(resultCEFBitmapBitBuffer)
}

func (m *TBufferPanel) OrigBufferWidth() int32 {
	r1 := bufferPanelImportAPI().SysCallN(27, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) OrigBufferHeight() int32 {
	r1 := bufferPanelImportAPI().SysCallN(26, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) OrigPopupBuffer() ICEFBitmapBitBuffer {
	var resultCEFBitmapBitBuffer uintptr
	bufferPanelImportAPI().SysCallN(28, m.Instance(), uintptr(unsafePointer(&resultCEFBitmapBitBuffer)))
	return AsCEFBitmapBitBuffer(resultCEFBitmapBitBuffer)
}

func (m *TBufferPanel) OrigPopupBufferWidth() int32 {
	r1 := bufferPanelImportAPI().SysCallN(31, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) OrigPopupBufferHeight() int32 {
	r1 := bufferPanelImportAPI().SysCallN(30, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) OrigPopupBufferBits() uintptr {
	r1 := bufferPanelImportAPI().SysCallN(29, m.Instance())
	return uintptr(r1)
}

func (m *TBufferPanel) OrigPopupScanlineSize() int32 {
	r1 := bufferPanelImportAPI().SysCallN(32, m.Instance())
	return int32(r1)
}

func (m *TBufferPanel) ParentFormHandle() TCefWindowHandle {
	r1 := bufferPanelImportAPI().SysCallN(36, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TBufferPanel) ParentForm() ICustomForm {
	var resultCustomForm uintptr
	bufferPanelImportAPI().SysCallN(35, m.Instance(), uintptr(unsafePointer(&resultCustomForm)))
	return AsCustomForm(resultCustomForm)
}

func (m *TBufferPanel) Transparent() bool {
	r1 := bufferPanelImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBufferPanel) SetTransparent(AValue bool) {
	bufferPanelImportAPI().SysCallN(63, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBufferPanel) CopyOriginalBuffer() bool {
	r1 := bufferPanelImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBufferPanel) SetCopyOriginalBuffer(AValue bool) {
	bufferPanelImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBufferPanel) DragCursor() TCursor {
	r1 := bufferPanelImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TBufferPanel) SetDragCursor(AValue TCursor) {
	bufferPanelImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TBufferPanel) DragKind() TDragKind {
	r1 := bufferPanelImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TBufferPanel) SetDragKind(AValue TDragKind) {
	bufferPanelImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TBufferPanel) DragMode() TDragMode {
	r1 := bufferPanelImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TBufferPanel) SetDragMode(AValue TDragMode) {
	bufferPanelImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TBufferPanel) ParentFont() bool {
	r1 := bufferPanelImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBufferPanel) SetParentFont(AValue bool) {
	bufferPanelImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBufferPanel) ParentShowHint() bool {
	r1 := bufferPanelImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBufferPanel) SetParentShowHint(AValue bool) {
	bufferPanelImportAPI().SysCallN(37, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBufferPanel) SaveToFile(aFilename string) bool {
	r1 := bufferPanelImportAPI().SysCallN(38, m.Instance(), PascalStr(aFilename))
	return GoBool(r1)
}

func (m *TBufferPanel) InvalidatePanel() bool {
	r1 := bufferPanelImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TBufferPanel) BeginBufferDraw() bool {
	r1 := bufferPanelImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TBufferPanel) UpdateBufferDimensions(aWidth, aHeight int32) bool {
	r1 := bufferPanelImportAPI().SysCallN(64, m.Instance(), uintptr(aWidth), uintptr(aHeight))
	return GoBool(r1)
}

func (m *TBufferPanel) UpdateOrigBufferDimensions(aWidth, aHeight int32) bool {
	r1 := bufferPanelImportAPI().SysCallN(66, m.Instance(), uintptr(aWidth), uintptr(aHeight))
	return GoBool(r1)
}

func (m *TBufferPanel) UpdateOrigPopupBufferDimensions(aWidth, aHeight int32) bool {
	r1 := bufferPanelImportAPI().SysCallN(67, m.Instance(), uintptr(aWidth), uintptr(aHeight))
	return GoBool(r1)
}

func (m *TBufferPanel) BufferIsResized(aUseMutex bool) bool {
	r1 := bufferPanelImportAPI().SysCallN(8, m.Instance(), PascalBool(aUseMutex))
	return GoBool(r1)
}

func (m *TBufferPanel) EndBufferDraw() {
	bufferPanelImportAPI().SysCallN(18, m.Instance())
}

func (m *TBufferPanel) BufferDraw(x, y int32, aBitmap IBitmap) {
	bufferPanelImportAPI().SysCallN(5, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(aBitmap))
}

func (m *TBufferPanel) BufferDraw1(aBitmap IBitmap, aSrcRect, aDstRect *TRect) {
	bufferPanelImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(aBitmap), uintptr(unsafePointer(aSrcRect)), uintptr(unsafePointer(aDstRect)))
}

func (m *TBufferPanel) UpdateDeviceScaleFactor() {
	bufferPanelImportAPI().SysCallN(65, m.Instance())
}

func (m *TBufferPanel) CreateIMEHandler() {
	bufferPanelImportAPI().SysCallN(12, m.Instance())
}

func (m *TBufferPanel) DrawOrigPopupBuffer(aSrcRect, aDstRect *TRect) {
	bufferPanelImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(aSrcRect)), uintptr(unsafePointer(aDstRect)))
}

func (m *TBufferPanel) SetOnIMECancelComposition(fn TNotify) {
	if m.iMECancelCompositionPtr != 0 {
		RemoveEventElement(m.iMECancelCompositionPtr)
	}
	m.iMECancelCompositionPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(50, m.Instance(), m.iMECancelCompositionPtr)
}

func (m *TBufferPanel) SetOnIMECommitText(fn TOnIMECommitText) {
	if m.iMECommitTextPtr != 0 {
		RemoveEventElement(m.iMECommitTextPtr)
	}
	m.iMECommitTextPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(51, m.Instance(), m.iMECommitTextPtr)
}

func (m *TBufferPanel) SetOnIMESetComposition(fn TOnIMESetComposition) {
	if m.iMESetCompositionPtr != 0 {
		RemoveEventElement(m.iMESetCompositionPtr)
	}
	m.iMESetCompositionPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(52, m.Instance(), m.iMESetCompositionPtr)
}

func (m *TBufferPanel) SetOnCustomTouch(fn TOnHandledMessage) {
	if m.customTouchPtr != 0 {
		RemoveEventElement(m.customTouchPtr)
	}
	m.customTouchPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(43, m.Instance(), m.customTouchPtr)
}

func (m *TBufferPanel) SetOnPointerDown(fn TOnHandledMessage) {
	if m.pointerDownPtr != 0 {
		RemoveEventElement(m.pointerDownPtr)
	}
	m.pointerDownPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(58, m.Instance(), m.pointerDownPtr)
}

func (m *TBufferPanel) SetOnPointerUp(fn TOnHandledMessage) {
	if m.pointerUpPtr != 0 {
		RemoveEventElement(m.pointerUpPtr)
	}
	m.pointerUpPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(59, m.Instance(), m.pointerUpPtr)
}

func (m *TBufferPanel) SetOnPointerUpdate(fn TOnHandledMessage) {
	if m.pointerUpdatePtr != 0 {
		RemoveEventElement(m.pointerUpdatePtr)
	}
	m.pointerUpdatePtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(60, m.Instance(), m.pointerUpdatePtr)
}

func (m *TBufferPanel) SetOnPaintParentBkg(fn TNotify) {
	if m.paintParentBkgPtr != 0 {
		RemoveEventElement(m.paintParentBkgPtr)
	}
	m.paintParentBkgPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(57, m.Instance(), m.paintParentBkgPtr)
}

func (m *TBufferPanel) SetOnConstrainedResize(fn TConstrainedResize) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(41, m.Instance(), m.constrainedResizePtr)
}

func (m *TBufferPanel) SetOnContextPopup(fn TContextPopup) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(42, m.Instance(), m.contextPopupPtr)
}

func (m *TBufferPanel) SetOnDblClick(fn TNotify) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(44, m.Instance(), m.dblClickPtr)
}

func (m *TBufferPanel) SetOnDragDrop(fn TDragDrop) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(45, m.Instance(), m.dragDropPtr)
}

func (m *TBufferPanel) SetOnDragOver(fn TDragOver) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(46, m.Instance(), m.dragOverPtr)
}

func (m *TBufferPanel) SetOnEndDock(fn TEndDrag) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(47, m.Instance(), m.endDockPtr)
}

func (m *TBufferPanel) SetOnEndDrag(fn TEndDrag) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(48, m.Instance(), m.endDragPtr)
}

func (m *TBufferPanel) SetOnGetSiteInfo(fn TGetSiteInfo) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(49, m.Instance(), m.getSiteInfoPtr)
}

func (m *TBufferPanel) SetOnMouseDown(fn TMouse) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(53, m.Instance(), m.mouseDownPtr)
}

func (m *TBufferPanel) SetOnMouseMove(fn TMouseMove) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(54, m.Instance(), m.mouseMovePtr)
}

func (m *TBufferPanel) SetOnMouseUp(fn TMouse) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(55, m.Instance(), m.mouseUpPtr)
}

func (m *TBufferPanel) SetOnMouseWheel(fn TMouseWheel) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(56, m.Instance(), m.mouseWheelPtr)
}

func (m *TBufferPanel) SetOnStartDock(fn TStartDock) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(61, m.Instance(), m.startDockPtr)
}

func (m *TBufferPanel) SetOnStartDrag(fn TStartDrag) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	bufferPanelImportAPI().SysCallN(62, m.Instance(), m.startDragPtr)
}

var (
	bufferPanelImport       *imports.Imports = nil
	bufferPanelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BufferPanel_BeginBufferDraw", 0),
		/*1*/ imports.NewTable("BufferPanel_BevelEdges", 0),
		/*2*/ imports.NewTable("BufferPanel_BevelKind", 0),
		/*3*/ imports.NewTable("BufferPanel_Buffer", 0),
		/*4*/ imports.NewTable("BufferPanel_BufferBits", 0),
		/*5*/ imports.NewTable("BufferPanel_BufferDraw", 0),
		/*6*/ imports.NewTable("BufferPanel_BufferDraw1", 0),
		/*7*/ imports.NewTable("BufferPanel_BufferHeight", 0),
		/*8*/ imports.NewTable("BufferPanel_BufferIsResized", 0),
		/*9*/ imports.NewTable("BufferPanel_BufferWidth", 0),
		/*10*/ imports.NewTable("BufferPanel_CopyOriginalBuffer", 0),
		/*11*/ imports.NewTable("BufferPanel_Create", 0),
		/*12*/ imports.NewTable("BufferPanel_CreateIMEHandler", 0),
		/*13*/ imports.NewTable("BufferPanel_Ctl3D", 0),
		/*14*/ imports.NewTable("BufferPanel_DragCursor", 0),
		/*15*/ imports.NewTable("BufferPanel_DragKind", 0),
		/*16*/ imports.NewTable("BufferPanel_DragMode", 0),
		/*17*/ imports.NewTable("BufferPanel_DrawOrigPopupBuffer", 0),
		/*18*/ imports.NewTable("BufferPanel_EndBufferDraw", 0),
		/*19*/ imports.NewTable("BufferPanel_ForcedDeviceScaleFactor", 0),
		/*20*/ imports.NewTable("BufferPanel_FullRepaint", 0),
		/*21*/ imports.NewTable("BufferPanel_InvalidatePanel", 0),
		/*22*/ imports.NewTable("BufferPanel_Locked", 0),
		/*23*/ imports.NewTable("BufferPanel_MustInitBuffer", 0),
		/*24*/ imports.NewTable("BufferPanel_OnCanResize", 0),
		/*25*/ imports.NewTable("BufferPanel_OrigBuffer", 0),
		/*26*/ imports.NewTable("BufferPanel_OrigBufferHeight", 0),
		/*27*/ imports.NewTable("BufferPanel_OrigBufferWidth", 0),
		/*28*/ imports.NewTable("BufferPanel_OrigPopupBuffer", 0),
		/*29*/ imports.NewTable("BufferPanel_OrigPopupBufferBits", 0),
		/*30*/ imports.NewTable("BufferPanel_OrigPopupBufferHeight", 0),
		/*31*/ imports.NewTable("BufferPanel_OrigPopupBufferWidth", 0),
		/*32*/ imports.NewTable("BufferPanel_OrigPopupScanlineSize", 0),
		/*33*/ imports.NewTable("BufferPanel_ParentCtl3D", 0),
		/*34*/ imports.NewTable("BufferPanel_ParentFont", 0),
		/*35*/ imports.NewTable("BufferPanel_ParentForm", 0),
		/*36*/ imports.NewTable("BufferPanel_ParentFormHandle", 0),
		/*37*/ imports.NewTable("BufferPanel_ParentShowHint", 0),
		/*38*/ imports.NewTable("BufferPanel_SaveToFile", 0),
		/*39*/ imports.NewTable("BufferPanel_ScanlineSize", 0),
		/*40*/ imports.NewTable("BufferPanel_ScreenScale", 0),
		/*41*/ imports.NewTable("BufferPanel_SetOnConstrainedResize", 0),
		/*42*/ imports.NewTable("BufferPanel_SetOnContextPopup", 0),
		/*43*/ imports.NewTable("BufferPanel_SetOnCustomTouch", 0),
		/*44*/ imports.NewTable("BufferPanel_SetOnDblClick", 0),
		/*45*/ imports.NewTable("BufferPanel_SetOnDragDrop", 0),
		/*46*/ imports.NewTable("BufferPanel_SetOnDragOver", 0),
		/*47*/ imports.NewTable("BufferPanel_SetOnEndDock", 0),
		/*48*/ imports.NewTable("BufferPanel_SetOnEndDrag", 0),
		/*49*/ imports.NewTable("BufferPanel_SetOnGetSiteInfo", 0),
		/*50*/ imports.NewTable("BufferPanel_SetOnIMECancelComposition", 0),
		/*51*/ imports.NewTable("BufferPanel_SetOnIMECommitText", 0),
		/*52*/ imports.NewTable("BufferPanel_SetOnIMESetComposition", 0),
		/*53*/ imports.NewTable("BufferPanel_SetOnMouseDown", 0),
		/*54*/ imports.NewTable("BufferPanel_SetOnMouseMove", 0),
		/*55*/ imports.NewTable("BufferPanel_SetOnMouseUp", 0),
		/*56*/ imports.NewTable("BufferPanel_SetOnMouseWheel", 0),
		/*57*/ imports.NewTable("BufferPanel_SetOnPaintParentBkg", 0),
		/*58*/ imports.NewTable("BufferPanel_SetOnPointerDown", 0),
		/*59*/ imports.NewTable("BufferPanel_SetOnPointerUp", 0),
		/*60*/ imports.NewTable("BufferPanel_SetOnPointerUpdate", 0),
		/*61*/ imports.NewTable("BufferPanel_SetOnStartDock", 0),
		/*62*/ imports.NewTable("BufferPanel_SetOnStartDrag", 0),
		/*63*/ imports.NewTable("BufferPanel_Transparent", 0),
		/*64*/ imports.NewTable("BufferPanel_UpdateBufferDimensions", 0),
		/*65*/ imports.NewTable("BufferPanel_UpdateDeviceScaleFactor", 0),
		/*66*/ imports.NewTable("BufferPanel_UpdateOrigBufferDimensions", 0),
		/*67*/ imports.NewTable("BufferPanel_UpdateOrigPopupBufferDimensions", 0),
	}
)

func bufferPanelImportAPI() *imports.Imports {
	if bufferPanelImport == nil {
		bufferPanelImport = NewDefaultImports()
		bufferPanelImport.SetImportTable(bufferPanelImportTables)
		bufferPanelImportTables = nil
	}
	return bufferPanelImport
}
