//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

import (
	"github.com/energye/lcl/types"
)

type TCefWindowHandle = uintptr        // IFDEF : MSWINDOWS
type TCefCursorHandle = uintptr        // IFDEF : MSWINDOWS
type TCefEventHandle = uintptr         // IFDEF : MSWINDOWS
type TCefSharedTextureHandle = uintptr // IFDEF : MSWINDOWS
// TCefPlatformThreadId
//
//	Platform thread ID.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_thread_internal.h">CEF source file: /include/internal/cef_thread_internal.h (cef_platform_thread_id_t))</a>
type TCefPlatformThreadId = types.DWORD

// TCefPlatformThreadHandle
//
//	Platform thread handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_thread_internal.h">CEF source file: /include/internal/cef_thread_internal.h (cef_platform_thread_handle_t))</a>
type TCefPlatformThreadHandle = types.DWORD

// TCefTransitionType
//
//	Transition type for a request. Made up of one source value and 0 or more qualifiers.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t))</a>
type TCefTransitionType = types.Cardinal

// TCefColor
//
//	32-bit ARGB color value, not premultiplied. The color components are always in a known order. Equivalent to the SkColor type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_t))</a>
type TCefColor = types.Cardinal

// TCefErrorCode
//
//	Supported error code values.
//	Ranges:
//	0- 99 System related errors
//	100-199 Connection related errors
//	200-299 Certificate errors
//	300-399 HTTP errors
//	400-499 Cache errors
//	500-599 ?
//	600-699 FTP errors
//	700-799 Certificate manager errors
//	800-899 DNS resolver errors
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t))</a>
//	<a href="https://chromium.googlesource.com/chromium/src/+/master/net/base/net_error_list.h">For the complete list of error values see include/base/internal/cef_net_error_list.h which includes this Chromium source file /net/base/net_error_list.h)</a>
type TCefErrorCode = types.Integer

// TCefCertStatus
//
//	Supported certificate status code values. See net\cert\cert_status_flags.h for more information. CERT_STATUS_NONE is new in CEF because we use an enum while cert_status_flags.h uses a typedef and static const variables.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cert_status_t))</a>
type TCefCertStatus = types.Integer

// TCefSSLVersion
//
//	Supported SSL version values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_version_t))</a>
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:net/ssl/ssl_connection_status_flags.h">See net/ssl/ssl_connection_status_flags.h for more information.)</a>
type TCefSSLVersion = types.Integer

// TCefStringList
//
//	CEF string maps are a set of key/value string pairs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_list.h">CEF source file: /include/internal/cef_string_list.h (cef_string_list_t))</a>
type TCefStringList = types.Pointer

// TCefStringMap
//
//	CEF string maps are a set of key/value string pairs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_map.h">CEF source file: /include/internal/cef_string_map.h (cef_string_map_t))</a>
type TCefStringMap = types.Pointer

// TCefStringMultimap
//
//	CEF string multimaps are a set of key/value string pairs. More than one value can be assigned to a single key.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_multimap.h">CEF source file: /include/internal/cef_string_multimap.h (cef_string_multimap_t))</a>
type TCefStringMultimap = types.Pointer

// TCefStringMultimapHandle
//
//	CEF string multimaps are a set of key/value string pairs. More than one value can be assigned to a single key.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_multimap.h">CEF source file: /include/internal/cef_string_multimap.h (cef_string_multimap_t))</a>
type TCefStringMultimapHandle = types.Pointer

// TCefUriUnescapeRule
//
//	URI unescape rules passed to CefURIDecode().
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t))</a>
type TCefUriUnescapeRule = types.Integer

// TCefDomEventCategory
//
//	DOM event category flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t))</a>
type TCefDomEventCategory = types.Integer

// TCefEventFlags
//
//	Supported event bit flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t))</a>
type TCefEventFlags = types.Cardinal

// TCefDragOperations
//
//	"Verb" of a drag-and-drop operation as negotiated between the source and destination. These constants match their equivalents in WebCore's DragActions.h and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t))</a>
type TCefDragOperations = types.Cardinal
type TCefDragOperation = types.Cardinal

// TCefV8PropertyAttributes
//
//	V8 property attribute values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t))</a>
type TCefV8PropertyAttributes = types.Cardinal

// TCefUrlRequestFlags
//
//	Flags used to customize the behavior of CefURLRequest.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t))</a>
type TCefUrlRequestFlags = types.Cardinal

// TCefContextMenuTypeFlags
//
//	Supported context menu type flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t))</a>
type TCefContextMenuTypeFlags = types.Cardinal

// TCefContextMenuMediaStateFlags
//
//	Supported context menu media state bit flags. These constants match their equivalents in Chromium's ContextMenuData::MediaFlags and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_media_state_flags_t))</a>
type TCefContextMenuMediaStateFlags = types.Cardinal

// TCefContextMenuEditStateFlags
//
//	Supported context menu edit state bit flags. These constants match their equivalents in Chromium's ContextMenuDataEditFlags and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_edit_state_flags_t))</a>
type TCefContextMenuEditStateFlags = types.Cardinal

// TCefJsonWriterOptions
//
//	Options that can be passed to CefWriteJSON.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t))</a>
type TCefJsonWriterOptions = types.Cardinal

// TCefSSLContentStatus
//
//	Supported SSL content status flags. See content/public/common/ssl_status.h for more information.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_content_status_t))</a>
type TCefSSLContentStatus = types.Cardinal

// TCefLogSeverity
//
//	Log severity levels.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t))</a>
type TCefLogSeverity = types.Cardinal

// TCefFileDialogMode
//
//	Supported file dialog modes.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t))</a>
type TCefFileDialogMode = types.Cardinal

// TCefDuplexMode
//
//	Print job duplex mode values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_duplex_mode_t))</a>
type TCefDuplexMode = types.Integer

// TCefSchemeOptions
//
//	Configuration options for registering a custom scheme. These values are used when calling AddCustomScheme.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t))</a>
type TCefSchemeOptions = types.Integer

// TCefMediaRouterCreateResult
//
//	Result codes for ICefMediaRouter.CreateRoute. Should be kept in sync with Chromium's media_router::mojom::RouteRequestResultCode type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_create_result_t))</a>
type TCefMediaRouterCreateResult = types.Integer

// TCefCookiePriority
//
//	Cookie priority values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_priority_t))</a>
type TCefCookiePriority = types.Integer

// TCefTextFieldCommands
//
//	Represents commands available to TextField.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_field_commands_t))</a>
type TCefTextFieldCommands = types.Integer

// TCefChromeToolbarType
//
//	Chrome toolbar types.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_toolbar_type_t))</a>
type TCefChromeToolbarType = types.Integer

// TCefDockingMode
//
//	Docking modes supported by ICefWindow.AddOverlay.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_docking_mode_t))</a>
type TCefDockingMode = types.Integer

// TCefShowState
//
//	Show states supported by ICefWindowDelegate.GetInitialShowState.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t))</a>
type TCefShowState = types.Integer

// TCefQuickMenuEditStateFlags
//
//	Supported quick menu state bit flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t))</a>
type TCefQuickMenuEditStateFlags = types.Integer

// TCefTouchHandleStateFlags
//
//	Values indicating what state of the touch handle is set.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_handle_state_flags_t))</a>
type TCefTouchHandleStateFlags = types.Integer

// TCefMediaAccessPermissionTypes
//
//	Media access permissions used by OnRequestMediaAccessPermission.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t))</a>
type TCefMediaAccessPermissionTypes = types.Integer

// TCefPermissionRequestTypes
//
//	Permission types used with OnShowPermissionPrompt. Some types are platform-specific or only supported with the Chrome runtime. Should be kept in sync with Chromium's permissions::RequestType type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_types_t))</a>
type TCefPermissionRequestTypes = types.Integer

// TCefDownloadInterruptReason
//
//	Download interrupt reasons. Should be kept in sync with Chromium's download::DownloadInterruptReason type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t))</a>
type TCefDownloadInterruptReason = types.Integer

// TCefMenuId
//
//	Supported menu IDs. Non-English translations can be provided for the IDS_MENU_* strings in ICefResourceBundleHandler.GetLocalizedString().
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t))</a>
type TCefMenuId = types.Integer

// TCefLogItems
//
//	Log items prepended to each log line.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t))</a>
type TCefLogItems = types.Cardinal

// TCefResultCode
// Process result codes. This is not a comprehensive list, as result codes
// might also include platform-specific crash values (Posix signal or Windows
// hardware exception), or internal-only implementation values.
// </summary>
// <remarks>
// <para>See the uCEFConstants unit for all possible values.</para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see></para>
// <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see></para>
// <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see></para>
// <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see></para>
type TCefResultCode = types.Integer
type TCefBaseTime = int64
type NativeUInt = types.NativeUInt

// TCefResourceHandlerClass = class of TCefResourceHandlerOwn
type TCefResourceHandlerClass = uintptr

const WM_APP = 0x8000
