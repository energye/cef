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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/types"
)

// IBufferPanel Parent: ICustomPanel
type IBufferPanel interface {
	ICustomPanel
	SaveToFile(filename string) bool                                                         // function
	InvalidatePanel() bool                                                                   // function
	BeginBufferDraw() bool                                                                   // function
	UpdateBufferDimensions(width int32, height int32) bool                                   // function
	UpdateOrigBufferDimensions(width int32, height int32) bool                               // function
	UpdateOrigPopupBufferDimensions(width int32, height int32) bool                          // function
	BufferIsResized(useMutex bool) bool                                                      // function
	EndBufferDraw()                                                                          // procedure
	BufferDrawWithIntX2Bitmap(X int32, Y int32, bitmap lcl.IBitmap)                          // procedure
	BufferDrawWithBitmapRectX2(bitmap lcl.IBitmap, srcRect types.TRect, dstRect types.TRect) // procedure
	UpdateDeviceScaleFactor()                                                                // procedure
	CreateIMEHandler()                                                                       // procedure
	ChangeCompositionRange(selectionRange TCefRange, characterBounds ICefRectArray)          // procedure
	DrawOrigPopupBuffer(srcRect types.TRect, dstRect types.TRect)                            // procedure
	ScanlineSize() int32                                                                     // property ScanlineSize Getter
	BufferWidth() int32                                                                      // property BufferWidth Getter
	BufferHeight() int32                                                                     // property BufferHeight Getter
	BufferBits() uintptr                                                                     // property BufferBits Getter
	ScreenScale() float32                                                                    // property ScreenScale Getter
	ForcedDeviceScaleFactor() float32                                                        // property ForcedDeviceScaleFactor Getter
	SetForcedDeviceScaleFactor(value float32)                                                // property ForcedDeviceScaleFactor Setter
	MustInitBuffer() bool                                                                    // property MustInitBuffer Getter
	SetMustInitBuffer(value bool)                                                            // property MustInitBuffer Setter
	Buffer() lcl.IBitmap                                                                     // property Buffer Getter
	OrigBuffer() ICEFBitmapBitBuffer                                                         // property OrigBuffer Getter
	OrigBufferWidth() int32                                                                  // property OrigBufferWidth Getter
	OrigBufferHeight() int32                                                                 // property OrigBufferHeight Getter
	OrigPopupBuffer() ICEFBitmapBitBuffer                                                    // property OrigPopupBuffer Getter
	OrigPopupBufferWidth() int32                                                             // property OrigPopupBufferWidth Getter
	OrigPopupBufferHeight() int32                                                            // property OrigPopupBufferHeight Getter
	OrigPopupBufferBits() uintptr                                                            // property OrigPopupBufferBits Getter
	OrigPopupScanlineSize() int32                                                            // property OrigPopupScanlineSize Getter
	ParentFormHandle() cefTypes.TCefWindowHandle                                             // property ParentFormHandle Getter
	ParentForm() lcl.ICustomForm                                                             // property ParentForm Getter
	Transparent() bool                                                                       // property Transparent Getter
	SetTransparent(value bool)                                                               // property Transparent Setter
	CopyOriginalBuffer() bool                                                                // property CopyOriginalBuffer Getter
	SetCopyOriginalBuffer(value bool)                                                        // property CopyOriginalBuffer Setter
	DragCursor() types.TCursor                                                               // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                                       // property DragCursor Setter
	DragKind() types.TDragKind                                                               // property DragKind Getter
	SetDragKind(value types.TDragKind)                                                       // property DragKind Setter
	DragMode() types.TDragMode                                                               // property DragMode Getter
	SetDragMode(value types.TDragMode)                                                       // property DragMode Setter
	ParentFont() bool                                                                        // property ParentFont Getter
	SetParentFont(value bool)                                                                // property ParentFont Setter
	ParentShowHint() bool                                                                    // property ParentShowHint Getter
	SetParentShowHint(value bool)                                                            // property ParentShowHint Setter
	SetOnIMECancelComposition(fn TNotifyEvent)                                               // property event
	SetOnIMECommitText(fn TOnIMECommitTextEvent)                                             // property event
	SetOnIMESetComposition(fn TOnIMESetCompositionEvent)                                     // property event
	SetOnCustomTouch(fn TOnHandledMessageEvent)                                              // property event
	SetOnPointerDown(fn TOnHandledMessageEvent)                                              // property event
	SetOnPointerUp(fn TOnHandledMessageEvent)                                                // property event
	SetOnPointerUpdate(fn TOnHandledMessageEvent)                                            // property event
	SetOnPaintParentBkg(fn TNotifyEvent)                                                     // property event
	SetOnConstrainedResize(fn TConstrainedResizeEvent)                                       // property event
	SetOnContextPopup(fn TContextPopupEvent)                                                 // property event
	SetOnDblClick(fn TNotifyEvent)                                                           // property event
	SetOnDragDrop(fn TDragDropEvent)                                                         // property event
	SetOnDragOver(fn TDragOverEvent)                                                         // property event
	SetOnEndDock(fn TEndDragEvent)                                                           // property event
	SetOnEndDrag(fn TEndDragEvent)                                                           // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)                                                   // property event
	SetOnMouseDown(fn TMouseEvent)                                                           // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                                       // property event
	SetOnMouseUp(fn TMouseEvent)                                                             // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                                     // property event
	SetOnStartDock(fn TStartDockEvent)                                                       // property event
	SetOnStartDrag(fn TStartDragEvent)                                                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                                                         // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                         // property event
}

type TBufferPanel struct {
	TCustomPanel
}

func (m *TBufferPanel) SaveToFile(filename string) bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(1, m.Instance(), api.PasStr(filename))
	return api.GoBool(r)
}

func (m *TBufferPanel) InvalidatePanel() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) BeginBufferDraw() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) UpdateBufferDimensions(width int32, height int32) bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(4, m.Instance(), uintptr(width), uintptr(height))
	return api.GoBool(r)
}

func (m *TBufferPanel) UpdateOrigBufferDimensions(width int32, height int32) bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(5, m.Instance(), uintptr(width), uintptr(height))
	return api.GoBool(r)
}

func (m *TBufferPanel) UpdateOrigPopupBufferDimensions(width int32, height int32) bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(6, m.Instance(), uintptr(width), uintptr(height))
	return api.GoBool(r)
}

func (m *TBufferPanel) BufferIsResized(useMutex bool) bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(7, m.Instance(), api.PasBool(useMutex))
	return api.GoBool(r)
}

func (m *TBufferPanel) EndBufferDraw() {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(8, m.Instance())
}

func (m *TBufferPanel) BufferDrawWithIntX2Bitmap(X int32, Y int32, bitmap lcl.IBitmap) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(9, m.Instance(), uintptr(X), uintptr(Y), base.GetObjectUintptr(bitmap))
}

func (m *TBufferPanel) BufferDrawWithBitmapRectX2(bitmap lcl.IBitmap, srcRect types.TRect, dstRect types.TRect) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(bitmap), uintptr(base.UnsafePointer(&srcRect)), uintptr(base.UnsafePointer(&dstRect)))
}

func (m *TBufferPanel) UpdateDeviceScaleFactor() {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(11, m.Instance())
}

func (m *TBufferPanel) CreateIMEHandler() {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(12, m.Instance())
}

func (m *TBufferPanel) ChangeCompositionRange(selectionRange TCefRange, characterBounds ICefRectArray) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&selectionRange)), characterBounds.Instance(), uintptr(int32(characterBounds.Count())))
}

func (m *TBufferPanel) DrawOrigPopupBuffer(srcRect types.TRect, dstRect types.TRect) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&srcRect)), uintptr(base.UnsafePointer(&dstRect)))
}

func (m *TBufferPanel) ScanlineSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(15, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) BufferWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(16, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) BufferHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(17, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) BufferBits() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(18, m.Instance())
	return uintptr(r)
}

func (m *TBufferPanel) ScreenScale() (result float32) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBufferPanel) ForcedDeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(20, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBufferPanel) SetForcedDeviceScaleFactor(value float32) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(20, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TBufferPanel) MustInitBuffer() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) SetMustInitBuffer(value bool) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TBufferPanel) Buffer() lcl.IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := bufferPanelAPI().SysCallN(22, m.Instance())
	return lcl.AsBitmap(r)
}

func (m *TBufferPanel) OrigBuffer() ICEFBitmapBitBuffer {
	if !m.IsValid() {
		return nil
	}
	r := bufferPanelAPI().SysCallN(23, m.Instance())
	return AsCEFBitmapBitBuffer(r)
}

func (m *TBufferPanel) OrigBufferWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(24, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) OrigBufferHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) OrigPopupBuffer() ICEFBitmapBitBuffer {
	if !m.IsValid() {
		return nil
	}
	r := bufferPanelAPI().SysCallN(26, m.Instance())
	return AsCEFBitmapBitBuffer(r)
}

func (m *TBufferPanel) OrigPopupBufferWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(27, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) OrigPopupBufferHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(28, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) OrigPopupBufferBits() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(29, m.Instance())
	return uintptr(r)
}

func (m *TBufferPanel) OrigPopupScanlineSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(30, m.Instance())
	return int32(r)
}

func (m *TBufferPanel) ParentFormHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(31, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TBufferPanel) ParentForm() lcl.ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := bufferPanelAPI().SysCallN(32, m.Instance())
	return lcl.AsCustomForm(r)
}

func (m *TBufferPanel) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TBufferPanel) CopyOriginalBuffer() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(34, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) SetCopyOriginalBuffer(value bool) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(34, 1, m.Instance(), api.PasBool(value))
}

func (m *TBufferPanel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(35, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TBufferPanel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(36, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TBufferPanel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := bufferPanelAPI().SysCallN(37, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TBufferPanel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(38, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(38, 1, m.Instance(), api.PasBool(value))
}

func (m *TBufferPanel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := bufferPanelAPI().SysCallN(39, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBufferPanel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	bufferPanelAPI().SysCallN(39, 1, m.Instance(), api.PasBool(value))
}

func (m *TBufferPanel) SetOnIMECancelComposition(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 40, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnIMECommitText(fn TOnIMECommitTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIMECommitTextEvent(fn)
	base.SetEvent(m, 41, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnIMESetComposition(fn TOnIMESetCompositionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIMESetCompositionEvent(fn)
	base.SetEvent(m, 42, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnCustomTouch(fn TOnHandledMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnHandledMessageEvent(fn)
	base.SetEvent(m, 43, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnPointerDown(fn TOnHandledMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnHandledMessageEvent(fn)
	base.SetEvent(m, 44, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnPointerUp(fn TOnHandledMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnHandledMessageEvent(fn)
	base.SetEvent(m, 45, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnPointerUpdate(fn TOnHandledMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnHandledMessageEvent(fn)
	base.SetEvent(m, 46, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnPaintParentBkg(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 47, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 48, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 49, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 50, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 51, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 52, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 53, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 54, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 55, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 56, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 57, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 58, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 59, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 60, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 61, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 62, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBufferPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 63, bufferPanelAPI(), api.MakeEventDataPtr(cb))
}

// NewBufferPanel class constructor
func NewBufferPanel(owner lcl.IComponent) IBufferPanel {
	r := bufferPanelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsBufferPanel(r)
}

var (
	bufferPanelOnce   base.Once
	bufferPanelImport *imports.Imports = nil
)

func bufferPanelAPI() *imports.Imports {
	bufferPanelOnce.Do(func() {
		bufferPanelImport = api.NewDefaultImports()
		bufferPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBufferPanel_Create", 0), // constructor NewBufferPanel
			/* 1 */ imports.NewTable("TBufferPanel_SaveToFile", 0), // function SaveToFile
			/* 2 */ imports.NewTable("TBufferPanel_InvalidatePanel", 0), // function InvalidatePanel
			/* 3 */ imports.NewTable("TBufferPanel_BeginBufferDraw", 0), // function BeginBufferDraw
			/* 4 */ imports.NewTable("TBufferPanel_UpdateBufferDimensions", 0), // function UpdateBufferDimensions
			/* 5 */ imports.NewTable("TBufferPanel_UpdateOrigBufferDimensions", 0), // function UpdateOrigBufferDimensions
			/* 6 */ imports.NewTable("TBufferPanel_UpdateOrigPopupBufferDimensions", 0), // function UpdateOrigPopupBufferDimensions
			/* 7 */ imports.NewTable("TBufferPanel_BufferIsResized", 0), // function BufferIsResized
			/* 8 */ imports.NewTable("TBufferPanel_EndBufferDraw", 0), // procedure EndBufferDraw
			/* 9 */ imports.NewTable("TBufferPanel_BufferDrawWithIntX2Bitmap", 0), // procedure BufferDrawWithIntX2Bitmap
			/* 10 */ imports.NewTable("TBufferPanel_BufferDrawWithBitmapRectX2", 0), // procedure BufferDrawWithBitmapRectX2
			/* 11 */ imports.NewTable("TBufferPanel_UpdateDeviceScaleFactor", 0), // procedure UpdateDeviceScaleFactor
			/* 12 */ imports.NewTable("TBufferPanel_CreateIMEHandler", 0), // procedure CreateIMEHandler
			/* 13 */ imports.NewTable("TBufferPanel_ChangeCompositionRange", 0), // procedure ChangeCompositionRange
			/* 14 */ imports.NewTable("TBufferPanel_DrawOrigPopupBuffer", 0), // procedure DrawOrigPopupBuffer
			/* 15 */ imports.NewTable("TBufferPanel_ScanlineSize", 0), // property ScanlineSize
			/* 16 */ imports.NewTable("TBufferPanel_BufferWidth", 0), // property BufferWidth
			/* 17 */ imports.NewTable("TBufferPanel_BufferHeight", 0), // property BufferHeight
			/* 18 */ imports.NewTable("TBufferPanel_BufferBits", 0), // property BufferBits
			/* 19 */ imports.NewTable("TBufferPanel_ScreenScale", 0), // property ScreenScale
			/* 20 */ imports.NewTable("TBufferPanel_ForcedDeviceScaleFactor", 0), // property ForcedDeviceScaleFactor
			/* 21 */ imports.NewTable("TBufferPanel_MustInitBuffer", 0), // property MustInitBuffer
			/* 22 */ imports.NewTable("TBufferPanel_Buffer", 0), // property Buffer
			/* 23 */ imports.NewTable("TBufferPanel_OrigBuffer", 0), // property OrigBuffer
			/* 24 */ imports.NewTable("TBufferPanel_OrigBufferWidth", 0), // property OrigBufferWidth
			/* 25 */ imports.NewTable("TBufferPanel_OrigBufferHeight", 0), // property OrigBufferHeight
			/* 26 */ imports.NewTable("TBufferPanel_OrigPopupBuffer", 0), // property OrigPopupBuffer
			/* 27 */ imports.NewTable("TBufferPanel_OrigPopupBufferWidth", 0), // property OrigPopupBufferWidth
			/* 28 */ imports.NewTable("TBufferPanel_OrigPopupBufferHeight", 0), // property OrigPopupBufferHeight
			/* 29 */ imports.NewTable("TBufferPanel_OrigPopupBufferBits", 0), // property OrigPopupBufferBits
			/* 30 */ imports.NewTable("TBufferPanel_OrigPopupScanlineSize", 0), // property OrigPopupScanlineSize
			/* 31 */ imports.NewTable("TBufferPanel_ParentFormHandle", 0), // property ParentFormHandle
			/* 32 */ imports.NewTable("TBufferPanel_ParentForm", 0), // property ParentForm
			/* 33 */ imports.NewTable("TBufferPanel_Transparent", 0), // property Transparent
			/* 34 */ imports.NewTable("TBufferPanel_CopyOriginalBuffer", 0), // property CopyOriginalBuffer
			/* 35 */ imports.NewTable("TBufferPanel_DragCursor", 0), // property DragCursor
			/* 36 */ imports.NewTable("TBufferPanel_DragKind", 0), // property DragKind
			/* 37 */ imports.NewTable("TBufferPanel_DragMode", 0), // property DragMode
			/* 38 */ imports.NewTable("TBufferPanel_ParentFont", 0), // property ParentFont
			/* 39 */ imports.NewTable("TBufferPanel_ParentShowHint", 0), // property ParentShowHint
			/* 40 */ imports.NewTable("TBufferPanel_OnIMECancelComposition", 0), // event OnIMECancelComposition
			/* 41 */ imports.NewTable("TBufferPanel_OnIMECommitText", 0), // event OnIMECommitText
			/* 42 */ imports.NewTable("TBufferPanel_OnIMESetComposition", 0), // event OnIMESetComposition
			/* 43 */ imports.NewTable("TBufferPanel_OnCustomTouch", 0), // event OnCustomTouch
			/* 44 */ imports.NewTable("TBufferPanel_OnPointerDown", 0), // event OnPointerDown
			/* 45 */ imports.NewTable("TBufferPanel_OnPointerUp", 0), // event OnPointerUp
			/* 46 */ imports.NewTable("TBufferPanel_OnPointerUpdate", 0), // event OnPointerUpdate
			/* 47 */ imports.NewTable("TBufferPanel_OnPaintParentBkg", 0), // event OnPaintParentBkg
			/* 48 */ imports.NewTable("TBufferPanel_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 49 */ imports.NewTable("TBufferPanel_OnContextPopup", 0), // event OnContextPopup
			/* 50 */ imports.NewTable("TBufferPanel_OnDblClick", 0), // event OnDblClick
			/* 51 */ imports.NewTable("TBufferPanel_OnDragDrop", 0), // event OnDragDrop
			/* 52 */ imports.NewTable("TBufferPanel_OnDragOver", 0), // event OnDragOver
			/* 53 */ imports.NewTable("TBufferPanel_OnEndDock", 0), // event OnEndDock
			/* 54 */ imports.NewTable("TBufferPanel_OnEndDrag", 0), // event OnEndDrag
			/* 55 */ imports.NewTable("TBufferPanel_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 56 */ imports.NewTable("TBufferPanel_OnMouseDown", 0), // event OnMouseDown
			/* 57 */ imports.NewTable("TBufferPanel_OnMouseMove", 0), // event OnMouseMove
			/* 58 */ imports.NewTable("TBufferPanel_OnMouseUp", 0), // event OnMouseUp
			/* 59 */ imports.NewTable("TBufferPanel_OnMouseWheel", 0), // event OnMouseWheel
			/* 60 */ imports.NewTable("TBufferPanel_OnStartDock", 0), // event OnStartDock
			/* 61 */ imports.NewTable("TBufferPanel_OnStartDrag", 0), // event OnStartDrag
			/* 62 */ imports.NewTable("TBufferPanel_OnMouseEnter", 0), // event OnMouseEnter
			/* 63 */ imports.NewTable("TBufferPanel_OnMouseLeave", 0), // event OnMouseLeave
		}
	})
	return bufferPanelImport
}
