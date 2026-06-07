//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// TCefAlphaType ENUM
//
//	include/internal/cef_types.h (cef_alpha_type_t)
type TCefAlphaType = int32

const (
	CEF_ALPHA_TYPE_OPAQUE TCefAlphaType = iota
	CEF_ALPHA_TYPE_PREMULTIPLIED
	CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// TCefAplicationStatus ENUM
type TCefAplicationStatus = int32

const (
	AsLoading TCefAplicationStatus = iota
	AsLoaded
	AsInitialized
	AsShuttingDown
	AsUnloaded
	AsErrorMissingFiles
	AsErrorDLLVersion
	AsErrorLoadingLibrary
	AsErrorInitializingLibrary
	AsErrorExecutingProcess
)

// TCefAutoplayPolicy ENUM
type TCefAutoplayPolicy = int32

const (
	AppDefault TCefAutoplayPolicy = iota
	AppDocumentUserActivationRequired
	AppNoUserGestureRequired
	AppUserGestureRequired
)

// TCefBatterySaverModeState ENUM
//
//	Values used by the battery saver mode state preference
//	https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h
type TCefBatterySaverModeState = int32

const (
	BsmsDisabled              TCefBatterySaverModeState = iota
	BsmsEnabledBelowThreshold                           = 1
	BsmsEnabledOnBattery                                = 2
	BsmsEnabled                                         = 3
	BsmsDefault                                         = 4
)

// TCefBrowserNavigation ENUM
//
//	Used by TCefBrowserNavigationTask to navigate in the right CEF thread
type TCefBrowserNavigation = int32

const (
	BnBack TCefBrowserNavigation = iota
	BnForward
	BnReload
	BnReloadIgnoreCache
	BnStopLoad
)

// TCefButtonState ENUM
//
//	include/internal/cef_types.h (cef_button_state_t)
type TCefButtonState = int32

const (
	CEF_BUTTON_STATE_NORMAL TCefButtonState = iota
	CEF_BUTTON_STATE_HOVERED
	CEF_BUTTON_STATE_PRESSED
	CEF_BUTTON_STATE_DISABLED
)

// TCefChannelLayout ENUM
//
//	include/internal/cef_types.h (cef_channel_layout_t)
type TCefChannelLayout = int32

const (
	CEF_CHANNEL_LAYOUT_NONE TCefChannelLayout = iota
	CEF_CHANNEL_LAYOUT_UNSUPPORTED
	CEF_CHANNEL_LAYOUT_MONO
	CEF_CHANNEL_LAYOUT_STEREO
	CEF_CHANNEL_LAYOUT_2_1
	CEF_CHANNEL_LAYOUT_SURROUND
	CEF_CHANNEL_LAYOUT_4_0
	CEF_CHANNEL_LAYOUT_2_2
	CEF_CHANNEL_LAYOUT_QUAD
	CEF_CHANNEL_LAYOUT_5_0
	CEF_CHANNEL_LAYOUT_5_1
	CEF_CHANNEL_LAYOUT_5_0_BACK
	CEF_CHANNEL_LAYOUT_5_1_BACK
	CEF_CHANNEL_LAYOUT_7_0
	CEF_CHANNEL_LAYOUT_7_1
	CEF_CHANNEL_LAYOUT_7_1_WIDE
	CEF_CHANNEL_LAYOUT_STEREO_DOWNMIX
	CEF_CHANNEL_LAYOUT_2POINT1
	CEF_CHANNEL_LAYOUT_3_1
	CEF_CHANNEL_LAYOUT_4_1
	CEF_CHANNEL_LAYOUT_6_0
	CEF_CHANNEL_LAYOUT_6_0_FRONT
	CEF_CHANNEL_LAYOUT_HEXAGONAL
	CEF_CHANNEL_LAYOUT_6_1
	CEF_CHANNEL_LAYOUT_6_1_BACK
	CEF_CHANNEL_LAYOUT_6_1_FRONT
	CEF_CHANNEL_LAYOUT_7_0_FRONT
	CEF_CHANNEL_LAYOUT_7_1_WIDE_BACK
	CEF_CHANNEL_LAYOUT_OCTAGONAL
	CEF_CHANNEL_LAYOUT_DISCRETE
	CEF_CHANNEL_LAYOUT_STEREO_AND_KEYBOARD_MIC
	CEF_CHANNEL_LAYOUT_4_1_QUAD_SIDE
	CEF_CHANNEL_LAYOUT_BITSTREAM
	CEF_CHANNEL_LAYOUT_5_1_4_DOWNMIX
)

// TCefClearDataStorageTypes ENUM
type TCefClearDataStorageTypes = int32

const (
	CdstAppCache TCefClearDataStorageTypes = iota
	CdstCookies
	CdstFileSystems
	CdstIndexeddb
	CdstLocalStorage
	CdstShaderCache
	CdstWebsql
	CdstServiceWorkers
	CdstCacheStorage
	CdstAll
)

// TCefCloseBrowserAction ENUM
//
//	Used in TChromium.Onclose
//	-------------------------
//	cbaCancel : stop closing the browser
//	cbaClose : continue closing the browser
//	cbaDelay : stop closing the browser momentarily. Used when the application
//	needs to execute some custom processes before closing the
//	browser. This is usually needed to destroy a TCEFWindowParent
//	in the main thread before closing the browser.
type TCefCloseBrowserAction = int32

const (
	CbaClose TCefCloseBrowserAction = iota
	CbaDelay
	CbaCancel
)

// TCefColorModel ENUM
//
//	include/internal/cef_types.h (cef_color_model_t)
type TCefColorModel = int32

const (
	COLOR_MODEL_UNKNOWN TCefColorModel = iota
	COLOR_MODEL_GRAY
	COLOR_MODEL_COLOR
	COLOR_MODEL_CMYK
	COLOR_MODEL_CMY
	COLOR_MODEL_KCMY
	COLOR_MODEL_CMY_K
	COLOR_MODEL_BLACK
	COLOR_MODEL_GRAYSCALE
	COLOR_MODEL_RGB
	COLOR_MODEL_RGB16
	COLOR_MODEL_RGBA
	COLOR_MODEL_COLORMODE_COLOR
	COLOR_MODEL_COLORMODE_MONOCHROME
	COLOR_MODEL_HP_COLOR_COLOR
	COLOR_MODEL_HP_COLOR_BLACK
	COLOR_MODEL_PRINTOUTMODE_NORMAL
	COLOR_MODEL_PRINTOUTMODE_NORMAL_GRAY
	COLOR_MODEL_PROCESSCOLORMODEL_CMYK
	COLOR_MODEL_PROCESSCOLORMODEL_GREYSCALE
	COLOR_MODEL_PROCESSCOLORMODEL_RGB
)

// TCefColorType ENUM
//
//	include/internal/cef_types.h (cef_color_type_t)
type TCefColorType = int32

const (
	CEF_COLOR_TYPE_RGBA_8888 TCefColorType = iota
	CEF_COLOR_TYPE_BGRA_8888
)

// TCefCOMInitMode ENUM
//
//	include/internal/cef_types.h (cef_com_init_mode_t)
type TCefCOMInitMode = int32

const (
	COM_INIT_MODE_NONE TCefCOMInitMode = iota
	COM_INIT_MODE_STA
	COM_INIT_MODE_MTA
)

// TCefCompositionUnderlineStyle ENUM
//
//	include/internal/cef_types.h (cef_composition_underline_style_t)
type TCefCompositionUnderlineStyle = int32

const (
	CEF_CUS_SOLID TCefCompositionUnderlineStyle = iota
	CEF_CUS_DOT
	CEF_CUS_DASH
	CEF_CUS_NONE
)

// TCefContextMenuMediaType ENUM
//
//	include/internal/cef_types.h (cef_context_menu_media_type_t)
type TCefContextMenuMediaType = int32

const (
	CM_MEDIATYPE_NONE TCefContextMenuMediaType = iota
	CM_MEDIATYPE_IMAGE
	CM_MEDIATYPE_VIDEO
	CM_MEDIATYPE_AUDIO
	CM_MEDIATYPE_CANVAS
	CM_MEDIATYPE_FILE
	CM_MEDIATYPE_PLUGIN
)

// TCefCookiePref ENUM
//
//	Used in TChromium preferences to allow or block cookies.
type TCefCookiePref = int32

const (
	CpDefault TCefCookiePref = iota
	CpAllow
	CpBlock
)

// TCefCookieSameSite ENUM
//
//	include/internal/cef_types.h (cef_cookie_same_site_t)
type TCefCookieSameSite = int32

const (
	CEF_COOKIE_SAME_SITE_UNSPECIFIED TCefCookieSameSite = iota
	CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	CEF_COOKIE_SAME_SITE_LAX_MODE
	CEF_COOKIE_SAME_SITE_STRICT_MODE
)

// TCefCrossAxisAlignment ENUM
//
//	include/internal/cef_types.h (cef_cross_axis_alignment_t)
type TCefCrossAxisAlignment = int32

const (
	CEF_CROSS_AXIS_ALIGNMENT_STRETCH TCefCrossAxisAlignment = iota
	CEF_CROSS_AXIS_ALIGNMENT_START
	CEF_CROSS_AXIS_ALIGNMENT_CENTER
	CEF_CROSS_AXIS_ALIGNMENT_END
)

// TCefCursorType ENUM
//
//	include/internal/cef_types.h (cef_cursor_type_t)
type TCefCursorType = int32

const (
	CT_POINTER TCefCursorType = iota
	CT_CROSS
	CT_HAND
	CT_IBEAM
	CT_WAIT
	CT_HELP
	CT_EASTRESIZE
	CT_NORTHRESIZE
	CT_NORTHEASTRESIZE
	CT_NORTHWESTRESIZE
	CT_SOUTHRESIZE
	CT_SOUTHEASTRESIZE
	CT_SOUTHWESTRESIZE
	CT_WESTRESIZE
	CT_NORTHSOUTHRESIZE
	CT_EASTWESTRESIZE
	CT_NORTHEASTSOUTHWESTRESIZE
	CT_NORTHWESTSOUTHEASTRESIZE
	CT_COLUMNRESIZE
	CT_ROWRESIZE
	CT_MIDDLEPANNING
	CT_EASTPANNING
	CT_NORTHPANNING
	CT_NORTHEASTPANNING
	CT_NORTHWESTPANNING
	CT_SOUTHPANNING
	CT_SOUTHEASTPANNING
	CT_SOUTHWESTPANNING
	CT_WESTPANNING
	CT_MOVE
	CT_VERTICALTEXT
	CT_CELL
	CT_CONTEXTMENU
	CT_ALIAS
	CT_PROGRESS
	CT_NODROP
	CT_COPY
	CT_NONE
	CT_NOTALLOWED
	CT_ZOOMIN
	CT_ZOOMOUT
	CT_GRAB
	CT_GRABBING
	CT_MIDDLE_PANNING_VERTICAL
	CT_MIDDLE_PANNING_HORIZONTAL
	CT_CUSTOM
	CT_DND_NONE
	CT_DND_MOVE
	CT_DND_COPY
	CT_DND_LIN
)

// TCEFDialogType ENUM
//
//	Used by TCEFFileDialogInfo
type TCEFDialogType = int32

const (
	DtOpen TCEFDialogType = iota
	DtOpenMultiple
	DtOpenFolder
	DtSave
)

// TCefDomDocumentType ENUM
//
//	include/internal/cef_types.h (cef_dom_document_type_t)
type TCefDomDocumentType = int32

const (
	DOM_DOCUMENT_TYPE_UNKNOWN TCefDomDocumentType = iota
	DOM_DOCUMENT_TYPE_HTML
	DOM_DOCUMENT_TYPE_XHTML
	DOM_DOCUMENT_TYPE_PLUGIN
)

// TCefDomEventPhase ENUM
//
//	include/internal/cef_types.h (cef_dom_event_phase_t)
type TCefDomEventPhase = int32

const (
	DOM_EVENT_PHASE_UNKNOWN TCefDomEventPhase = iota
	DOM_EVENT_PHASE_CAPTURING
	DOM_EVENT_PHASE_AT_TARGET
	DOM_EVENT_PHASE_BUBBLING
)

// TCefDomNodeType ENUM
//
//	include/internal/cef_types.h (cef_dom_node_type_t)
type TCefDomNodeType = int32

const (
	DOM_NODE_TYPE_UNSUPPORTED TCefDomNodeType = iota
	DOM_NODE_TYPE_ELEMENT
	DOM_NODE_TYPE_ATTRIBUTE
	DOM_NODE_TYPE_TEXT
	DOM_NODE_TYPE_CDATA_SECTION
	DOM_NODE_TYPE_PROCESSING_INSTRUCTIONS
	DOM_NODE_TYPE_COMMENT
	DOM_NODE_TYPE_DOCUMENT
	DOM_NODE_TYPE_DOCUMENT_TYPE
	DOM_NODE_TYPE_DOCUMENT_FRAGMENT
)

// TCefFocusSource ENUM
//
//	include/internal/cef_types.h (cef_focus_source_t)
type TCefFocusSource = int32

const (
	FOCUS_SOURCE_NAVIGATION TCefFocusSource = iota
	FOCUS_SOURCE_SYSTEM
)

// TCefHorizontalAlignment ENUM
//
//	include/internal/cef_types.h (cef_horizontal_alignment_t)
type TCefHorizontalAlignment = int32

const (
	CEF_HORIZONTAL_ALIGNMENT_LEFT TCefHorizontalAlignment = iota
	CEF_HORIZONTAL_ALIGNMENT_CENTER
	CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// TCefJsDialogType ENUM
//
//	include/internal/cef_types.h (cef_jsdialog_type_t)
type TCefJsDialogType = int32

const (
	JSDIALOGTYPE_ALERT TCefJsDialogType = iota
	JSDIALOGTYPE_CONFIRM
	JSDIALOGTYPE_PROMPT
)

// TCefJsonParserOptions ENUM
//
//	include/internal/cef_types.h (cef_json_parser_options_t)
type TCefJsonParserOptions = int32

const (
	JSON_PARSER_RFC                   TCefJsonParserOptions = iota
	JSON_PARSER_ALLOW_TRAILING_COMMAS                       = 1 << 0
)

// TCefKeyEventType ENUM
//
//	include/internal/cef_types.h (cef_key_event_type_t)
type TCefKeyEventType = int32

const (
	KEYEVENT_RAWKEYDOWN TCefKeyEventType = iota
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

// TCefMainAxisAlignment ENUM
//
//	include/internal/cef_types.h (cef_main_axis_alignment_t)
type TCefMainAxisAlignment = int32

const (
	CEF_MAIN_AXIS_ALIGNMENT_START TCefMainAxisAlignment = iota
	CEF_MAIN_AXIS_ALIGNMENT_CENTER
	CEF_MAIN_AXIS_ALIGNMENT_END
)

// TCefMediaRouteConnectionState ENUM
//
//	include/internal/cef_types.h (cef_media_route_connection_state_t)
type TCefMediaRouteConnectionState = int32

const (
	CEF_MRCS_UNKNOWN TCefMediaRouteConnectionState = iota
	CEF_MRCS_CONNECTING
	CEF_MRCS_CONNECTED
	CEF_MRCS_CLOSED
	CEF_MRCS_TERMINATED
)

// TCefMediaSinkIconType ENUM
//
//	include/internal/cef_types.h (cef_media_sink_icon_type_t)
type TCefMediaSinkIconType = int32

const (
	CEF_MSIT_CAST TCefMediaSinkIconType = iota
	CEF_MSIT_CAST_AUDIO_GROUP
	CEF_MSIT_CAST_AUDIO
	CEF_MSIT_MEETING
	CEF_MSIT_HANGOUT
	CEF_MSIT_EDUCATION
	CEF_MSIT_WIRED_DISPLAY
	CEF_MSIT_GENERIC
	CEF_MSIT_TOTAL_COUNT
)

// TCefMediaType ENUM
//
//	Used by TCefMediaSinkInfo and TCefMediaSourceInfo
type TCefMediaType = int32

const (
	MtCast TCefMediaType = iota
	MtDial
	MtUnknown
)

// TCefMenuAnchorPosition ENUM
//
//	include/internal/cef_types.h (cef_menu_anchor_position_t)
type TCefMenuAnchorPosition = int32

const (
	CEF_MENU_ANCHOR_TOPLEFT TCefMenuAnchorPosition = iota
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

// TCefMenuColorType ENUM
//
//	include/internal/cef_types.h (cef_menu_color_type_t)
type TCefMenuColorType = int32

const (
	CEF_MENU_COLOR_TEXT TCefMenuColorType = iota
	CEF_MENU_COLOR_TEXT_HOVERED
	CEF_MENU_COLOR_TEXT_ACCELERATOR
	CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED
	CEF_MENU_COLOR_BACKGROUND
	CEF_MENU_COLOR_BACKGROUND_HOVERED
	CEF_MENU_COLOR_COUNT
)

// TCefMenuItemType ENUM
//
//	include/internal/cef_types.h (cef_menu_item_type_t)
type TCefMenuItemType = int32

const (
	MENUITEMTYPE_NONE TCefMenuItemType = iota
	MENUITEMTYPE_COMMAND
	MENUITEMTYPE_CHECK
	MENUITEMTYPE_RADIO
	MENUITEMTYPE_SEPARATOR
	MENUITEMTYPE_SUBMENU
)

// TCefMessageLoopType ENUM
//
//	include/internal/cef_types.h (cef_message_loop_type_t)
type TCefMessageLoopType = int32

const (
	ML_TYPE_DEFAULT TCefMessageLoopType = iota
	ML_TYPE_UI
	ML_TYPE_IO
)

// TCefMouseButtonType ENUM
//
//	include/internal/cef_types.h (cef_mouse_button_type_t)
type TCefMouseButtonType = int32

const (
	MBT_LEFT TCefMouseButtonType = iota
	MBT_MIDDLE
	MBT_RIGHT
)

// TCefNavigationType ENUM
//
//	include/internal/cef_types.h (cef_navigation_type_t)
type TCefNavigationType = int32

const (
	NAVIGATION_LINK_CLICKED TCefNavigationType = iota
	NAVIGATION_FORM_SUBMITTED
	NAVIGATION_BACK_FORWARD
	NAVIGATION_RELOAD
	NAVIGATION_FORM_RESUBMITTED
	NAVIGATION_OTHER
)

// TCefNetLogCaptureMode ENUM
//
//	Values used by the --net-log-capture-mode command line switch.
//	Sets the granularity of events to capture in the network log.
//	https://source.chromium.org/chromium/chromium/src/+/main:content/browser/network_service_instance_impl.cc
//	https://source.chromium.org/chromium/chromium/src/+/main:net/log/net_log_capture_mode.h
type TCefNetLogCaptureMode = int32

const (
	NlcmDefault TCefNetLogCaptureMode = iota
	NlcmIncludeSensitive
	NlcmEverything
)

// TCefPaintElementType ENUM
//
//	include/internal/cef_types.h (cef_paint_element_type_t)
type TCefPaintElementType = int32

const (
	PET_VIEW TCefPaintElementType = iota
	PET_POPUP
)

// TCefPathKey ENUM
//
//	include/internal/cef_types.h (cef_path_key_t)
type TCefPathKey = int32

const (
	PK_DIR_CURRENT TCefPathKey = iota
	PK_DIR_EXE
	PK_DIR_MODULE
	PK_DIR_TEMP
	PK_FILE_EXE
	PK_FILE_MODULE
	PK_LOCAL_APP_DATA
	PK_USER_DATA
	PK_DIR_RESOURCES
)

// TCefPdfPrintMarginType ENUM
//
//	include/internal/cef_types.h (cef_pdf_print_margin_type_t)
type TCefPdfPrintMarginType = int32

const (
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	PDF_PRINT_MARGIN_NONE
	PDF_PRINT_MARGIN_CUSTOM
)

// TCefPermissionRequestResult ENUM
//
//	include/internal/cef_types.h (cef_permission_request_result_t)
type TCefPermissionRequestResult = int32

const (
	CEF_PERMISSION_RESULT_ACCEPT TCefPermissionRequestResult = iota
	CEF_PERMISSION_RESULT_DENY
	CEF_PERMISSION_RESULT_DISMISS
	CEF_PERMISSION_RESULT_IGNORE
)

// TCefPointerType ENUM
//
//	include/internal/cef_types.h (cef_pointer_type_t)
type TCefPointerType = int32

const (
	CEF_POINTER_TYPE_TOUCH TCefPointerType = iota
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

// TCefPostDataElementType ENUM
//
//	include/internal/cef_types.h (cef_postdataelement_type_t)
type TCefPostDataElementType = int32

const (
	PDE_TYPE_EMPTY TCefPostDataElementType = iota
	PDE_TYPE_BYTES
	PDE_TYPE_FILE
)

// TCefPreferencesType ENUM
//
//	include/internal/cef_types.h (cef_preferences_type_t)
type TCefPreferencesType = int32

const (
	CEF_PREFERENCES_TYPE_GLOBAL TCefPreferencesType = iota
	CEF_PREFERENCES_TYPE_REQUEST_CONTEXT
)

// TCefProcessId ENUM
//
//	include/internal/cef_types.h (cef_process_id_t)
type TCefProcessId = int32

const (
	PID_BROWSER TCefProcessId = iota
	PID_RENDERER
)

// TCefProcessType ENUM
type TCefProcessType = int32

const (
	PtBrowser TCefProcessType = iota
	PtRenderer
	PtZygote
	PtGPU
	PtUtility
	PtBroker
	PtCrashpad
	PtOther
)

// TCefProxyScheme ENUM
type TCefProxyScheme = int32

const (
	PsHTTP TCefProxyScheme = iota
	PsSOCKS4
	PsSOCKS5
)

// TCefReferrerPolicy ENUM
//
//	include/internal/cef_types.h (cef_referrer_policy_t)
type TCefReferrerPolicy = int32

const (
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_FROM_SECURE_TO_INSECURE TCefReferrerPolicy = iota
	REFERRER_POLICY_REDUCE_REFERRER_GRANULARITY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_ONLY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_NEVER_CLEAR_REFERRER
	REFERRER_POLICY_ORIGIN
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_CLEAR_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	REFERRER_POLICY_NO_REFERRER
)

// TCefResourceType ENUM
//
//	include/internal/cef_types.h (cef_resource_type_t)
type TCefResourceType = int32

const (
	RT_MAIN_FRAME TCefResourceType = iota
	RT_SUB_FRAME
	RT_STYLESHEET
	RT_SCRIPT
	RT_IMAGE
	RT_FONT_RESOURCE
	RT_SUB_RESOURCE
	RT_OBJECT
	RT_MEDIA
	RT_WORKER
	RT_SHARED_WORKER
	RT_PREFETCH
	RT_FAVICON
	RT_XHR
	RT_PING
	RT_SERVICE_WORKER
	RT_CSP_REPORT
	RT_PLUGIN_RESOURCE
	RT_EMPTY_FILLER_TYPE_DO_NOT_USE
	RT_NAVIGATION_PRELOAD_MAIN_FRAME
	RT_NAVIGATION_PRELOAD_SUB_FRAME
)

// TCefResponseFilterStatus ENUM
//
//	include/internal/cef_types.h (cef_response_filter_status_t)
type TCefResponseFilterStatus = int32

const (
	RESPONSE_FILTER_NEED_MORE_DATA TCefResponseFilterStatus = iota
	RESPONSE_FILTER_DONE
	RESPONSE_FILTER_ERROR
)

// TCefReturnValue ENUM
//
//	include/internal/cef_types.h (cef_return_value_t)
type TCefReturnValue = int32

const (
	RV_CANCEL TCefReturnValue = iota
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

// TCefScaleFactor ENUM
//
//	include/internal/cef_types.h (cef_scale_factor_t)
type TCefScaleFactor = int32

const (
	SCALE_FACTOR_NONE TCefScaleFactor = iota
	SCALE_FACTOR_100P
	SCALE_FACTOR_125P
	SCALE_FACTOR_133P
	SCALE_FACTOR_140P
	SCALE_FACTOR_150P
	SCALE_FACTOR_180P
	SCALE_FACTOR_200P
	SCALE_FACTOR_250P
	SCALE_FACTOR_300P
)

// TCefState ENUM
//
//	include/internal/cef_types.h (cef_state_t)
type TCefState = int32

const (
	STATE_DEFAULT TCefState = iota
	STATE_ENABLED
	STATE_DISABLED
)

// TCefStorageType ENUM
//
//	include/internal/cef_types.h (cef_storage_type_t)
type TCefStorageType = int32

const (
	ST_LOCALSTORAGE TCefStorageType = iota
	ST_SESSIONSTORAGE
)

// TCefTerminationStatus ENUM
//
//	include/internal/cef_types.h (cef_termination_status_t)
type TCefTerminationStatus = int32

const (
	TS_ABNORMAL_TERMINATION TCefTerminationStatus = iota
	TS_PROCESS_WAS_KILLED
	TS_PROCESS_CRASHED
	TS_PROCESS_OOM
)

// TCefTestCertType ENUM
//
//	include/internal/cef_types.h (cef_test_cert_type_t)
type TCefTestCertType = int32

const (
	CEF_TEST_CERT_OK_IP TCefTestCertType = iota
	CEF_TEST_CERT_OK_DOMAIN
	CEF_TEST_CERT_EXPIRED
)

// TCefTextInpuMode ENUM
//
//	include/internal/cef_types.h (cef_text_input_mode_t)
type TCefTextInpuMode = int32

const (
	CEF_TEXT_INPUT_MODE_DEFAULT TCefTextInpuMode = iota
	CEF_TEXT_INPUT_MODE_NONE
	CEF_TEXT_INPUT_MODE_TEXT
	CEF_TEXT_INPUT_MODE_TEL
	CEF_TEXT_INPUT_MODE_URL
	CEF_TEXT_INPUT_MODE_EMAIL
	CEF_TEXT_INPUT_MODE_NUMERIC
	CEF_TEXT_INPUT_MODE_DECIMAL
	CEF_TEXT_INPUT_MODE_SEARCH
)

// TCefTextStyle ENUM
//
//	include/internal/cef_types.h (cef_text_style_t)
type TCefTextStyle = int32

const (
	CEF_TEXT_STYLE_BOLD TCefTextStyle = iota
	CEF_TEXT_STYLE_ITALIC
	CEF_TEXT_STYLE_STRIKE
	CEF_TEXT_STYLE_DIAGONAL_STRIKE
	CEF_TEXT_STYLE_UNDERLINE
)

// TCefThreadId ENUM
//
//	include/internal/cef_types.h (cef_thread_id_t)
type TCefThreadId = int32

const (
	TID_UI TCefThreadId = iota
	TID_FILE_BACKGROUND
	TID_FILE_USER_VISIBLE
	TID_FILE_USER_BLOCKING
	TID_PROCESS_LAUNCHER
	TID_IO
	TID_RENDERER
)

// TCefThreadPriority ENUM
//
//	include/internal/cef_types.h (cef_thread_priority_t)
type TCefThreadPriority = int32

const (
	TP_BACKGROUND TCefThreadPriority = iota
	TP_NORMAL
	TP_DISPLAY
	TP_REALTIME_AUDIO
)

// TCefTouchEeventType ENUM
//
//	include/internal/cef_types.h (cef_touch_event_type_t)
type TCefTouchEeventType = int32

const (
	CEF_TET_RELEASED TCefTouchEeventType = iota
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

// TCefUrlRequestStatus ENUM
//
//	include/internal/cef_types.h (cef_urlrequest_status_t)
type TCefUrlRequestStatus = int32

const (
	UR_UNKNOWN TCefUrlRequestStatus = iota
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

// TCefValueType ENUM
//
//	include/internal/cef_types.h (cef_value_type_t)
type TCefValueType = int32

const (
	VTYPE_INVALID TCefValueType = iota
	VTYPE_NULL
	VTYPE_BOOL
	VTYPE_INT
	VTYPE_DOUBLE
	VTYPE_STRING
	VTYPE_BINARY
	VTYPE_DICTIONARY
	VTYPE_LIST
)

// TCefWebRTCHandlingPolicy ENUM
type TCefWebRTCHandlingPolicy = int32

const (
	HpDefault TCefWebRTCHandlingPolicy = iota
	HpDefaultPublicAndPrivateInterfaces
	HpDefaultPublicInterfaceOnly
	HpDisableNonProxiedUDP
)

// TCefWindowOpenDisposition ENUM
//
//	include/internal/cef_types.h (cef_window_open_disposition_t)
type TCefWindowOpenDisposition = int32

const (
	WOD_UNKNOWN TCefWindowOpenDisposition = iota
	WOD_CURRENT_TAB
	WOD_SINGLETON_TAB
	WOD_NEW_FOREGROUND_TAB
	WOD_NEW_BACKGROUND_TAB
	WOD_NEW_POPUP
	WOD_NEW_WINDOW
	WOD_SAVE_TO_DISK
	WOD_OFF_THE_RECORD
	WOD_IGNORE_ACTION
	WOD_SWITCH_TO_TAB
	WOD_NEW_PICTURE_IN_PICTURE
)

// TCefXmlEncodingType ENUM
//
//	include/internal/cef_types.h (cef_xml_encoding_type_t)
type TCefXmlEncodingType = int32

const (
	XML_ENCODING_NONE TCefXmlEncodingType = iota
	XML_ENCODING_UTF8
	XML_ENCODING_UTF16LE
	XML_ENCODING_UTF16BE
	XML_ENCODING_ASCII
)

// TCefXmlNodeType ENUM
//
//	include/internal/cef_types.h (cef_xml_node_type_t)
type TCefXmlNodeType = int32

const (
	XML_NODE_UNSUPPORTED TCefXmlNodeType = iota
	XML_NODE_PROCESSING_INSTRUCTION
	XML_NODE_DOCUMENT_TYPE
	XML_NODE_ELEMENT_START
	XML_NODE_ELEMENT_END
	XML_NODE_ATTRIBUTE
	XML_NODE_TEXT
	XML_NODE_CDATA
	XML_NODE_ENTITY_REFERENCE
	XML_NODE_WHITESPACE
	XML_NODE_COMMENT
)

// TOAuthChallengeMethod ENUM
type TOAuthChallengeMethod = int32

const (
	CmPlain TOAuthChallengeMethod = iota
	CmSHA256
)

// TSentinelStatus ENUM
type TSentinelStatus = int32

const (
	SsIdle TSentinelStatus = iota
	SsInitialDelay
	SsCheckingChildren
	SsClosing
)
