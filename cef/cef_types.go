//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

type TCefAudioParameters struct {
	ChannelLayout   cefTypes.TCefChannelLayout // TCefChannelLayout
	SampleRate      int32                      // integer
	FramesPerBuffer int32                      // integer
}

type TCefBaseScoped struct {
	Del uintptr // procedure(self:PCefBaseScoped); stdcall
}

type TCefBoxLayoutSettings struct {
	Horizontal                    int32                      // Integer
	InsideBorderHorizontalSpacing int32                      // Integer
	InsideBorderVerticalSpacing   int32                      // Integer
	InsideBorderInsets            TCefInsets                 // TCefInsets
	BetweenChildSpacing           int32                      // Integer
	MainAxisAlignment             cefTypes.TCefAxisAlignment // TCefAxisAlignment
	CrossAxisAlignment            cefTypes.TCefAxisAlignment // TCefAxisAlignment
	MinimumCrossAxisSize          int32                      // Integer
	DefaultFlex                   int32                      // Integer
}

type TCefBrowserSettings struct {
	WindowlessFrameRate        int32              // Integer
	StandardFontFamily         string             // TCefString
	FixedFontFamily            string             // TCefString
	SerifFontFamily            string             // TCefString
	SansSerifFontFamily        string             // TCefString
	CursiveFontFamily          string             // TCefString
	FantasyFontFamily          string             // TCefString
	DefaultFontSize            int32              // Integer
	DefaultFixedFontSize       int32              // Integer
	MinimumFontSize            int32              // Integer
	MinimumLogicalFontSize     int32              // Integer
	DefaultEncoding            string             // TCefString
	RemoteFonts                cefTypes.TCefState // TCefState
	Javascript                 cefTypes.TCefState // TCefState
	JavascriptCloseWindows     cefTypes.TCefState // TCefState
	JavascriptAccessClipboard  cefTypes.TCefState // TCefState
	JavascriptDomPaste         cefTypes.TCefState // TCefState
	ImageLoading               cefTypes.TCefState // TCefState
	ImageShrinkStandaloneToFit cefTypes.TCefState // TCefState
	TextAreaResize             cefTypes.TCefState // TCefState
	TabToLinks                 cefTypes.TCefState // TCefState
	LocalStorage               cefTypes.TCefState // TCefState
	Databases                  cefTypes.TCefState // TCefState
	Webgl                      cefTypes.TCefState // TCefState
	BackgroundColor            cefTypes.TCefColor // TCefColor
	ChromeStatusBubble         cefTypes.TCefState // TCefState
	ChromeZoomBubble           cefTypes.TCefState // TCefState
}

type TCefCompositionUnderline struct {
	Range           TCefRange                              // TCefRange
	Color           cefTypes.TCefColor                     // TCefColor
	BackgroundColor cefTypes.TCefColor                     // TCefColor
	Thick           int32                                  // integer
	Style           cefTypes.TCefCompositionUnderlineStyle // TCefCompositionUnderlineStyle
}

type TCefCookie struct {
	Name       string                      // TCefString
	Value      string                      // TCefString
	Domain     string                      // TCefString
	Path       string                      // TCefString
	Secure     int32                       // Integer
	Httponly   int32                       // Integer
	Creation   int64                       // TCefBaseTime
	LastAccess int64                       // TCefBaseTime
	HasExpires int32                       // Integer
	Expires    int64                       // TCefBaseTime
	SameSite   cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority   cefTypes.TCefCookiePriority // TCefCookiePriority
}

type TCefCookieManagerRefSetCookieArgs struct {
	Url        string                      // ustring
	Name       string                      // ustring
	Value      string                      // ustring
	Domain     string                      // ustring
	Path       string                      // ustring
	Secure     types.LongBool              // Boolean
	Httponly   types.LongBool              // Boolean
	HasExpires types.LongBool              // Boolean
	Creation   types.TDateTime             // TDateTime
	LastAccess types.TDateTime             // TDateTime
	Expires    types.TDateTime             // TDateTime
	SameSite   cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority   cefTypes.TCefCookiePriority // TCefCookiePriority
	Callback   ICefSetCookieCallback       // ICefSetCookieCallback
}

type TCefCursorInfo struct {
	Hotspot          TCefPoint // TCefPoint
	ImageScaleFactor float32   // Single
	Buffer           uintptr   // Pointer
	Size             TCefSize  // TCefSize
}

type TCefDraggableRegion struct {
	Bounds    TCefRect // TCefRect
	Draggable int32    // Integer
}

type TCefInsets struct {
	Top    int32 // Integer
	Left   int32 // Integer
	Bottom int32 // Integer
	Right  int32 // Integer
}

type TCefKeyEvent struct {
	Kind                 cefTypes.TCefKeyEventType // TCefKeyEventType
	Modifiers            cefTypes.TCefEventFlags   // TCefEventFlags
	WindowsKeyCode       int32                     // Integer
	NativeKeyCode        int32                     // Integer
	IsSystemKey          int32                     // Integer
	Character            uint16                    // WideChar
	UnmodifiedCharacter  uint16                    // WideChar
	FocusOnEditableField int32                     // Integer
}

type TCefMouseEvent struct {
	X         int32                   // Integer
	Y         int32                   // Integer
	Modifiers cefTypes.TCefEventFlags // TCefEventFlags
}

type TCefPdfPrintSettings struct {
	Landscape               int32                           // Integer
	PrintBackground         int32                           // Integer
	Scale                   float64                         // double
	PaperWidth              float64                         // double
	PaperHeight             float64                         // double
	PreferCssPageSize       int32                           // Integer
	MarginType              cefTypes.TCefPdfPrintMarginType // TCefPdfPrintMarginType
	MarginTop               float64                         // double
	MarginRight             float64                         // double
	MarginBottom            float64                         // double
	MarginLeft              float64                         // double
	PageRanges              string                          // TCefString
	DisplayHeaderFooter     int32                           // Integer
	HeaderTemplate          string                          // TCefString
	FooterTemplate          string                          // TCefString
	GenerateTaggedPdf       int32                           // integer
	GenerateDocumentOutline int32                           // integer
}

type TCefPoint struct {
	X int32 // Integer
	Y int32 // Integer
}

type TCefPopupFeatures struct {
	X         int32 // Integer
	XSet      int32 // Integer
	Y         int32 // Integer
	YSet      int32 // Integer
	Width     int32 // Integer
	WidthSet  int32 // Integer
	Height    int32 // Integer
	HeightSet int32 // Integer
	IsPopup   int32 // Integer
}

type TCefPreferenceRegistrar struct {
	Base          TCefBaseScoped // TCefBaseScoped
	AddPreference uintptr        // function(self:PCefPreferenceRegistrar; const name:PCefString; default_value:PCefValue):Integer; stdcall
}

type TCefRange struct {
	From uint32 // cardinal
	To   uint32 // cardinal
}

type TCefRect struct {
	X      int32 // Integer
	Y      int32 // Integer
	Width  int32 // Integer
	Height int32 // Integer
}

type TCefRequestContextSettings struct {
	CachePath                        string // TCefString
	PersistSessionCookies            int32  // Integer
	PersistUserPreferences           int32  // Integer
	AcceptLanguageList               string // TCefString
	CookieableSchemesList            string // TCefString
	CookieableSchemesExcludeDefaults int32  // integer
}

type TCefScreenInfo struct {
	DeviceScaleFactor float32  // single
	Depth             int32    // integer
	DepthPerComponent int32    // integer
	IsMonochrome      int32    // integer
	Rect              TCefRect // TCefRect
	AvailableRect     TCefRect // TCefRect
}

type TCefSize struct {
	Width  int32 // Integer
	Height int32 // Integer
}

type TCefTime struct {
	Year        int32 // Integer
	Month       int32 // Integer
	DayOfWeek   int32 // Integer
	DayOfMonth  int32 // Integer
	Hour        int32 // Integer
	Minute      int32 // Integer
	Second      int32 // Integer
	Millisecond int32 // Integer
}

type TCefTouchEvent struct {
	Id            int32                        // integer
	X             float32                      // single
	Y             float32                      // single
	RadiusX       float32                      // single
	RadiusY       float32                      // single
	RotationAngle float32                      // single
	Pressure      float32                      // single
	Type          cefTypes.TCefTouchEeventType // TCefTouchEeventType
	Modifiers     cefTypes.TCefEventFlags      // TCefEventFlags
	PointerType   cefTypes.TCefPointerType     // TCefPointerType
}

type TCefTouchHandleState struct {
	TouchHandleId    int32                            // integer
	Flags            uint32                           // cardinal
	Enabled          int32                            // integer
	Orientation      cefTypes.TCefHorizontalAlignment // TCefHorizontalAlignment
	MirrorVertical   int32                            // integer
	MirrorHorizontal int32                            // integer
	Origin           TCefPoint                        // TCefPoint
	Alpha            float32                          // single
}

type TChromiumCoreSetCookieArgs struct {
	Url             string                      // ustring
	Name            string                      // ustring
	Value           string                      // ustring
	Domain          string                      // ustring
	Path            string                      // ustring
	Secure          types.LongBool              // Boolean
	Httponly        types.LongBool              // Boolean
	HasExpires      types.LongBool              // Boolean
	Creation        types.TDateTime             // TDateTime
	LastAccess      types.TDateTime             // TDateTime
	Expires         types.TDateTime             // TDateTime
	SameSite        cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority        cefTypes.TCefCookiePriority // TCefCookiePriority
	ASetImmediately types.LongBool              // boolean
	AID             int32                       // integer
}

type TChromiumCoreSimulateKeyEventArgs struct {
	Type                  cefTypes.TSimulatedCefKeyEventType // TSimulatedCefKeyEventType
	Modifiers             int32                              // integer
	Timestamp             float32                            // single
	Text                  string                             // ustring
	Unmodifiedtext        string                             // ustring
	KeyIdentifier         string                             // ustring
	Code                  string                             // ustring
	Key                   string                             // ustring
	WindowsVirtualKeyCode int32                              // integer
	NativeVirtualKeyCode  int32                              // integer
	AutoRepeat            types.LongBool                     // boolean
	IsKeypad              types.LongBool                     // boolean
	IsSystemKey           types.LongBool                     // boolean
	Location              cefTypes.TCefKeyLocation           // TCefKeyLocation
	Commands              cefTypes.TCefEditingCommand        // TCefEditingCommand
}

type TChromiumCoreSimulateMouseEventArgs struct {
	Type               cefTypes.TCefSimulatedMouseEventType // TCefSimulatedMouseEventType
	X                  float32                              // single
	Y                  float32                              // single
	Modifiers          int32                                // integer
	Timestamp          float32                              // single
	Button             cefTypes.TCefSimulatedMouseButton    // TCefSimulatedMouseButton
	Buttons            int32                                // integer
	ClickCount         int32                                // integer
	Force              float32                              // single
	TangentialPressure float32                              // single
	TiltX              float32                              // single
	TiltY              float32                              // single
	Twist              int32                                // integer
	DeltaX             float32                              // single
	DeltaY             float32                              // single
	PointerType        cefTypes.TCefSimulatedPointerType    // TCefSimulatedPointerType
}

func (m *TCefBrowserSettings) ToPas() *tCefBrowserSettings {
	if m == nil {
		return nil
	}
	return &tCefBrowserSettings{
		WindowlessFrameRate:        m.WindowlessFrameRate,
		StandardFontFamily:         api.PasStr(m.StandardFontFamily),
		FixedFontFamily:            api.PasStr(m.FixedFontFamily),
		SerifFontFamily:            api.PasStr(m.SerifFontFamily),
		SansSerifFontFamily:        api.PasStr(m.SansSerifFontFamily),
		CursiveFontFamily:          api.PasStr(m.CursiveFontFamily),
		FantasyFontFamily:          api.PasStr(m.FantasyFontFamily),
		DefaultFontSize:            m.DefaultFontSize,
		DefaultFixedFontSize:       m.DefaultFixedFontSize,
		MinimumFontSize:            m.MinimumFontSize,
		MinimumLogicalFontSize:     m.MinimumLogicalFontSize,
		DefaultEncoding:            api.PasStr(m.DefaultEncoding),
		RemoteFonts:                m.RemoteFonts,
		Javascript:                 m.Javascript,
		JavascriptCloseWindows:     m.JavascriptCloseWindows,
		JavascriptAccessClipboard:  m.JavascriptAccessClipboard,
		JavascriptDomPaste:         m.JavascriptDomPaste,
		ImageLoading:               m.ImageLoading,
		ImageShrinkStandaloneToFit: m.ImageShrinkStandaloneToFit,
		TextAreaResize:             m.TextAreaResize,
		TabToLinks:                 m.TabToLinks,
		LocalStorage:               m.LocalStorage,
		Databases:                  m.Databases,
		Webgl:                      m.Webgl,
		BackgroundColor:            m.BackgroundColor,
		ChromeStatusBubble:         m.ChromeStatusBubble,
		ChromeZoomBubble:           m.ChromeZoomBubble,
	}
}

type tCefBrowserSettings struct {
	WindowlessFrameRate        int32              // Integer
	StandardFontFamily         uintptr            // TCefString
	FixedFontFamily            uintptr            // TCefString
	SerifFontFamily            uintptr            // TCefString
	SansSerifFontFamily        uintptr            // TCefString
	CursiveFontFamily          uintptr            // TCefString
	FantasyFontFamily          uintptr            // TCefString
	DefaultFontSize            int32              // Integer
	DefaultFixedFontSize       int32              // Integer
	MinimumFontSize            int32              // Integer
	MinimumLogicalFontSize     int32              // Integer
	DefaultEncoding            uintptr            // TCefString
	RemoteFonts                cefTypes.TCefState // TCefState
	Javascript                 cefTypes.TCefState // TCefState
	JavascriptCloseWindows     cefTypes.TCefState // TCefState
	JavascriptAccessClipboard  cefTypes.TCefState // TCefState
	JavascriptDomPaste         cefTypes.TCefState // TCefState
	ImageLoading               cefTypes.TCefState // TCefState
	ImageShrinkStandaloneToFit cefTypes.TCefState // TCefState
	TextAreaResize             cefTypes.TCefState // TCefState
	TabToLinks                 cefTypes.TCefState // TCefState
	LocalStorage               cefTypes.TCefState // TCefState
	Databases                  cefTypes.TCefState // TCefState
	Webgl                      cefTypes.TCefState // TCefState
	BackgroundColor            cefTypes.TCefColor // TCefColor
	ChromeStatusBubble         cefTypes.TCefState // TCefState
	ChromeZoomBubble           cefTypes.TCefState // TCefState
}

func (m *tCefBrowserSettings) ToGo() TCefBrowserSettings {
	if m == nil {
		return TCefBrowserSettings{}
	}
	return TCefBrowserSettings{
		WindowlessFrameRate:        m.WindowlessFrameRate,
		StandardFontFamily:         api.GoStr(m.StandardFontFamily),
		FixedFontFamily:            api.GoStr(m.FixedFontFamily),
		SerifFontFamily:            api.GoStr(m.SerifFontFamily),
		SansSerifFontFamily:        api.GoStr(m.SansSerifFontFamily),
		CursiveFontFamily:          api.GoStr(m.CursiveFontFamily),
		FantasyFontFamily:          api.GoStr(m.FantasyFontFamily),
		DefaultFontSize:            m.DefaultFontSize,
		DefaultFixedFontSize:       m.DefaultFixedFontSize,
		MinimumFontSize:            m.MinimumFontSize,
		MinimumLogicalFontSize:     m.MinimumLogicalFontSize,
		DefaultEncoding:            api.GoStr(m.DefaultEncoding),
		RemoteFonts:                m.RemoteFonts,
		Javascript:                 m.Javascript,
		JavascriptCloseWindows:     m.JavascriptCloseWindows,
		JavascriptAccessClipboard:  m.JavascriptAccessClipboard,
		JavascriptDomPaste:         m.JavascriptDomPaste,
		ImageLoading:               m.ImageLoading,
		ImageShrinkStandaloneToFit: m.ImageShrinkStandaloneToFit,
		TextAreaResize:             m.TextAreaResize,
		TabToLinks:                 m.TabToLinks,
		LocalStorage:               m.LocalStorage,
		Databases:                  m.Databases,
		Webgl:                      m.Webgl,
		BackgroundColor:            m.BackgroundColor,
		ChromeStatusBubble:         m.ChromeStatusBubble,
		ChromeZoomBubble:           m.ChromeZoomBubble,
	}
}
func (m *TCefCookie) ToPas() *tCefCookie {
	if m == nil {
		return nil
	}
	return &tCefCookie{
		Name:       api.PasStr(m.Name),
		Value:      api.PasStr(m.Value),
		Domain:     api.PasStr(m.Domain),
		Path:       api.PasStr(m.Path),
		Secure:     m.Secure,
		Httponly:   m.Httponly,
		Creation:   m.Creation,
		LastAccess: m.LastAccess,
		HasExpires: m.HasExpires,
		Expires:    m.Expires,
		SameSite:   m.SameSite,
		Priority:   m.Priority,
	}
}

type tCefCookie struct {
	Name       uintptr                     // TCefString
	Value      uintptr                     // TCefString
	Domain     uintptr                     // TCefString
	Path       uintptr                     // TCefString
	Secure     int32                       // Integer
	Httponly   int32                       // Integer
	Creation   int64                       // TCefBaseTime
	LastAccess int64                       // TCefBaseTime
	HasExpires int32                       // Integer
	Expires    int64                       // TCefBaseTime
	SameSite   cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority   cefTypes.TCefCookiePriority // TCefCookiePriority
}

func (m *tCefCookie) ToGo() TCefCookie {
	if m == nil {
		return TCefCookie{}
	}
	return TCefCookie{
		Name:       api.GoStr(m.Name),
		Value:      api.GoStr(m.Value),
		Domain:     api.GoStr(m.Domain),
		Path:       api.GoStr(m.Path),
		Secure:     m.Secure,
		Httponly:   m.Httponly,
		Creation:   m.Creation,
		LastAccess: m.LastAccess,
		HasExpires: m.HasExpires,
		Expires:    m.Expires,
		SameSite:   m.SameSite,
		Priority:   m.Priority,
	}
}
func (m *TCefCookieManagerRefSetCookieArgs) ToPas() *tCefCookieManagerRefSetCookieArgs {
	if m == nil {
		return nil
	}
	return &tCefCookieManagerRefSetCookieArgs{
		Url:        api.PasStr(m.Url),
		Name:       api.PasStr(m.Name),
		Value:      api.PasStr(m.Value),
		Domain:     api.PasStr(m.Domain),
		Path:       api.PasStr(m.Path),
		Secure:     m.Secure,
		Httponly:   m.Httponly,
		HasExpires: m.HasExpires,
		Creation:   m.Creation,
		LastAccess: m.LastAccess,
		Expires:    m.Expires,
		SameSite:   m.SameSite,
		Priority:   m.Priority,
		Callback:   m.Callback,
	}
}

type tCefCookieManagerRefSetCookieArgs struct {
	Url        uintptr                     // ustring
	Name       uintptr                     // ustring
	Value      uintptr                     // ustring
	Domain     uintptr                     // ustring
	Path       uintptr                     // ustring
	Secure     types.LongBool              // Boolean
	Httponly   types.LongBool              // Boolean
	HasExpires types.LongBool              // Boolean
	Creation   types.TDateTime             // TDateTime
	LastAccess types.TDateTime             // TDateTime
	Expires    types.TDateTime             // TDateTime
	SameSite   cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority   cefTypes.TCefCookiePriority // TCefCookiePriority
	Callback   ICefSetCookieCallback       // ICefSetCookieCallback
}

func (m *tCefCookieManagerRefSetCookieArgs) ToGo() TCefCookieManagerRefSetCookieArgs {
	if m == nil {
		return TCefCookieManagerRefSetCookieArgs{}
	}
	return TCefCookieManagerRefSetCookieArgs{
		Url:        api.GoStr(m.Url),
		Name:       api.GoStr(m.Name),
		Value:      api.GoStr(m.Value),
		Domain:     api.GoStr(m.Domain),
		Path:       api.GoStr(m.Path),
		Secure:     m.Secure,
		Httponly:   m.Httponly,
		HasExpires: m.HasExpires,
		Creation:   m.Creation,
		LastAccess: m.LastAccess,
		Expires:    m.Expires,
		SameSite:   m.SameSite,
		Priority:   m.Priority,
		Callback:   m.Callback,
	}
}
func (m *TCefPdfPrintSettings) ToPas() *tCefPdfPrintSettings {
	if m == nil {
		return nil
	}
	return &tCefPdfPrintSettings{
		Landscape:               m.Landscape,
		PrintBackground:         m.PrintBackground,
		Scale:                   m.Scale,
		PaperWidth:              m.PaperWidth,
		PaperHeight:             m.PaperHeight,
		PreferCssPageSize:       m.PreferCssPageSize,
		MarginType:              m.MarginType,
		MarginTop:               m.MarginTop,
		MarginRight:             m.MarginRight,
		MarginBottom:            m.MarginBottom,
		MarginLeft:              m.MarginLeft,
		PageRanges:              api.PasStr(m.PageRanges),
		DisplayHeaderFooter:     m.DisplayHeaderFooter,
		HeaderTemplate:          api.PasStr(m.HeaderTemplate),
		FooterTemplate:          api.PasStr(m.FooterTemplate),
		GenerateTaggedPdf:       m.GenerateTaggedPdf,
		GenerateDocumentOutline: m.GenerateDocumentOutline,
	}
}

type tCefPdfPrintSettings struct {
	Landscape               int32                           // Integer
	PrintBackground         int32                           // Integer
	Scale                   float64                         // double
	PaperWidth              float64                         // double
	PaperHeight             float64                         // double
	PreferCssPageSize       int32                           // Integer
	MarginType              cefTypes.TCefPdfPrintMarginType // TCefPdfPrintMarginType
	MarginTop               float64                         // double
	MarginRight             float64                         // double
	MarginBottom            float64                         // double
	MarginLeft              float64                         // double
	PageRanges              uintptr                         // TCefString
	DisplayHeaderFooter     int32                           // Integer
	HeaderTemplate          uintptr                         // TCefString
	FooterTemplate          uintptr                         // TCefString
	GenerateTaggedPdf       int32                           // integer
	GenerateDocumentOutline int32                           // integer
}

func (m *tCefPdfPrintSettings) ToGo() TCefPdfPrintSettings {
	if m == nil {
		return TCefPdfPrintSettings{}
	}
	return TCefPdfPrintSettings{
		Landscape:               m.Landscape,
		PrintBackground:         m.PrintBackground,
		Scale:                   m.Scale,
		PaperWidth:              m.PaperWidth,
		PaperHeight:             m.PaperHeight,
		PreferCssPageSize:       m.PreferCssPageSize,
		MarginType:              m.MarginType,
		MarginTop:               m.MarginTop,
		MarginRight:             m.MarginRight,
		MarginBottom:            m.MarginBottom,
		MarginLeft:              m.MarginLeft,
		PageRanges:              api.GoStr(m.PageRanges),
		DisplayHeaderFooter:     m.DisplayHeaderFooter,
		HeaderTemplate:          api.GoStr(m.HeaderTemplate),
		FooterTemplate:          api.GoStr(m.FooterTemplate),
		GenerateTaggedPdf:       m.GenerateTaggedPdf,
		GenerateDocumentOutline: m.GenerateDocumentOutline,
	}
}
func (m *TCefRequestContextSettings) ToPas() *tCefRequestContextSettings {
	if m == nil {
		return nil
	}
	return &tCefRequestContextSettings{
		CachePath:                        api.PasStr(m.CachePath),
		PersistSessionCookies:            m.PersistSessionCookies,
		PersistUserPreferences:           m.PersistUserPreferences,
		AcceptLanguageList:               api.PasStr(m.AcceptLanguageList),
		CookieableSchemesList:            api.PasStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: m.CookieableSchemesExcludeDefaults,
	}
}

type tCefRequestContextSettings struct {
	CachePath                        uintptr // TCefString
	PersistSessionCookies            int32   // Integer
	PersistUserPreferences           int32   // Integer
	AcceptLanguageList               uintptr // TCefString
	CookieableSchemesList            uintptr // TCefString
	CookieableSchemesExcludeDefaults int32   // integer
}

func (m *tCefRequestContextSettings) ToGo() TCefRequestContextSettings {
	if m == nil {
		return TCefRequestContextSettings{}
	}
	return TCefRequestContextSettings{
		CachePath:                        api.GoStr(m.CachePath),
		PersistSessionCookies:            m.PersistSessionCookies,
		PersistUserPreferences:           m.PersistUserPreferences,
		AcceptLanguageList:               api.GoStr(m.AcceptLanguageList),
		CookieableSchemesList:            api.GoStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: m.CookieableSchemesExcludeDefaults,
	}
}
func (m *TChromiumCoreSetCookieArgs) ToPas() *tChromiumCoreSetCookieArgs {
	if m == nil {
		return nil
	}
	return &tChromiumCoreSetCookieArgs{
		Url:             api.PasStr(m.Url),
		Name:            api.PasStr(m.Name),
		Value:           api.PasStr(m.Value),
		Domain:          api.PasStr(m.Domain),
		Path:            api.PasStr(m.Path),
		Secure:          m.Secure,
		Httponly:        m.Httponly,
		HasExpires:      m.HasExpires,
		Creation:        m.Creation,
		LastAccess:      m.LastAccess,
		Expires:         m.Expires,
		SameSite:        m.SameSite,
		Priority:        m.Priority,
		ASetImmediately: m.ASetImmediately,
		AID:             m.AID,
	}
}

type tChromiumCoreSetCookieArgs struct {
	Url             uintptr                     // ustring
	Name            uintptr                     // ustring
	Value           uintptr                     // ustring
	Domain          uintptr                     // ustring
	Path            uintptr                     // ustring
	Secure          types.LongBool              // Boolean
	Httponly        types.LongBool              // Boolean
	HasExpires      types.LongBool              // Boolean
	Creation        types.TDateTime             // TDateTime
	LastAccess      types.TDateTime             // TDateTime
	Expires         types.TDateTime             // TDateTime
	SameSite        cefTypes.TCefCookieSameSite // TCefCookieSameSite
	Priority        cefTypes.TCefCookiePriority // TCefCookiePriority
	ASetImmediately types.LongBool              // boolean
	AID             int32                       // integer
}

func (m *tChromiumCoreSetCookieArgs) ToGo() TChromiumCoreSetCookieArgs {
	if m == nil {
		return TChromiumCoreSetCookieArgs{}
	}
	return TChromiumCoreSetCookieArgs{
		Url:             api.GoStr(m.Url),
		Name:            api.GoStr(m.Name),
		Value:           api.GoStr(m.Value),
		Domain:          api.GoStr(m.Domain),
		Path:            api.GoStr(m.Path),
		Secure:          m.Secure,
		Httponly:        m.Httponly,
		HasExpires:      m.HasExpires,
		Creation:        m.Creation,
		LastAccess:      m.LastAccess,
		Expires:         m.Expires,
		SameSite:        m.SameSite,
		Priority:        m.Priority,
		ASetImmediately: m.ASetImmediately,
		AID:             m.AID,
	}
}
func (m *TChromiumCoreSimulateKeyEventArgs) ToPas() *tChromiumCoreSimulateKeyEventArgs {
	if m == nil {
		return nil
	}
	return &tChromiumCoreSimulateKeyEventArgs{
		Type:                  m.Type,
		Modifiers:             m.Modifiers,
		Timestamp:             m.Timestamp,
		Text:                  api.PasStr(m.Text),
		Unmodifiedtext:        api.PasStr(m.Unmodifiedtext),
		KeyIdentifier:         api.PasStr(m.KeyIdentifier),
		Code:                  api.PasStr(m.Code),
		Key:                   api.PasStr(m.Key),
		WindowsVirtualKeyCode: m.WindowsVirtualKeyCode,
		NativeVirtualKeyCode:  m.NativeVirtualKeyCode,
		AutoRepeat:            m.AutoRepeat,
		IsKeypad:              m.IsKeypad,
		IsSystemKey:           m.IsSystemKey,
		Location:              m.Location,
		Commands:              m.Commands,
	}
}

type tChromiumCoreSimulateKeyEventArgs struct {
	Type                  cefTypes.TSimulatedCefKeyEventType // TSimulatedCefKeyEventType
	Modifiers             int32                              // integer
	Timestamp             float32                            // single
	Text                  uintptr                            // ustring
	Unmodifiedtext        uintptr                            // ustring
	KeyIdentifier         uintptr                            // ustring
	Code                  uintptr                            // ustring
	Key                   uintptr                            // ustring
	WindowsVirtualKeyCode int32                              // integer
	NativeVirtualKeyCode  int32                              // integer
	AutoRepeat            types.LongBool                     // boolean
	IsKeypad              types.LongBool                     // boolean
	IsSystemKey           types.LongBool                     // boolean
	Location              cefTypes.TCefKeyLocation           // TCefKeyLocation
	Commands              cefTypes.TCefEditingCommand        // TCefEditingCommand
}

func (m *tChromiumCoreSimulateKeyEventArgs) ToGo() TChromiumCoreSimulateKeyEventArgs {
	if m == nil {
		return TChromiumCoreSimulateKeyEventArgs{}
	}
	return TChromiumCoreSimulateKeyEventArgs{
		Type:                  m.Type,
		Modifiers:             m.Modifiers,
		Timestamp:             m.Timestamp,
		Text:                  api.GoStr(m.Text),
		Unmodifiedtext:        api.GoStr(m.Unmodifiedtext),
		KeyIdentifier:         api.GoStr(m.KeyIdentifier),
		Code:                  api.GoStr(m.Code),
		Key:                   api.GoStr(m.Key),
		WindowsVirtualKeyCode: m.WindowsVirtualKeyCode,
		NativeVirtualKeyCode:  m.NativeVirtualKeyCode,
		AutoRepeat:            m.AutoRepeat,
		IsKeypad:              m.IsKeypad,
		IsSystemKey:           m.IsSystemKey,
		Location:              m.Location,
		Commands:              m.Commands,
	}
}
func (m *TChromiumCoreSimulateMouseEventArgs) ToPas() *tChromiumCoreSimulateMouseEventArgs {
	if m == nil {
		return nil
	}
	return &tChromiumCoreSimulateMouseEventArgs{
		Type:               m.Type,
		X:                  m.X,
		Y:                  m.Y,
		Modifiers:          m.Modifiers,
		Timestamp:          m.Timestamp,
		Button:             m.Button,
		Buttons:            m.Buttons,
		ClickCount:         m.ClickCount,
		Force:              m.Force,
		TangentialPressure: m.TangentialPressure,
		TiltX:              m.TiltX,
		TiltY:              m.TiltY,
		Twist:              m.Twist,
		DeltaX:             m.DeltaX,
		DeltaY:             m.DeltaY,
		PointerType:        m.PointerType,
	}
}

type tChromiumCoreSimulateMouseEventArgs struct {
	Type               cefTypes.TCefSimulatedMouseEventType // TCefSimulatedMouseEventType
	X                  float32                              // single
	Y                  float32                              // single
	Modifiers          int32                                // integer
	Timestamp          float32                              // single
	Button             cefTypes.TCefSimulatedMouseButton    // TCefSimulatedMouseButton
	Buttons            int32                                // integer
	ClickCount         int32                                // integer
	Force              float32                              // single
	TangentialPressure float32                              // single
	TiltX              float32                              // single
	TiltY              float32                              // single
	Twist              int32                                // integer
	DeltaX             float32                              // single
	DeltaY             float32                              // single
	PointerType        cefTypes.TCefSimulatedPointerType    // TCefSimulatedPointerType
}

func (m *tChromiumCoreSimulateMouseEventArgs) ToGo() TChromiumCoreSimulateMouseEventArgs {
	if m == nil {
		return TChromiumCoreSimulateMouseEventArgs{}
	}
	return TChromiumCoreSimulateMouseEventArgs{
		Type:               m.Type,
		X:                  m.X,
		Y:                  m.Y,
		Modifiers:          m.Modifiers,
		Timestamp:          m.Timestamp,
		Button:             m.Button,
		Buttons:            m.Buttons,
		ClickCount:         m.ClickCount,
		Force:              m.Force,
		TangentialPressure: m.TangentialPressure,
		TiltX:              m.TiltX,
		TiltY:              m.TiltY,
		Twist:              m.Twist,
		DeltaX:             m.DeltaX,
		DeltaY:             m.DeltaY,
		PointerType:        m.PointerType,
	}
}
