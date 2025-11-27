//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// TBrowserUTF8KeyPressEvent ENUM
//
//	TBrowserKeyPressEvent = procedure(Sender: TObject; var Key: char; var AHandled: Boolean) of Object;
type TBrowserUTF8KeyPressEvent = int32

const (
	THiddenPropertyEditor TBrowserUTF8KeyPressEvent = iota
)

// TCefAlphaType ENUM
//
//	Describes how to interpret the alpha component of a pixel.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_alpha_type_t)</see>
type TCefAlphaType = int32

const (
	CEF_ALPHA_TYPE_OPAQUE TCefAlphaType = iota
	CEF_ALPHA_TYPE_PREMULTIPLIED
	CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// TCefAplicationStatus ENUM
//
//	Status of TCefAplicationCore.
type TCefAplicationStatus = int32

const (
	AsLoading TCefAplicationStatus = iota
	AsLoaded
	AsInitialized
	AsShuttingDown
	AsUnloaded
	AsErrorMissingFiles
	AsErrorDLLVersion
	AsErrorWindowsVersion
	AsErrorLoadingLibrary
	AsErrorInitializingLibrary
	AsErrorExecutingProcess
)

// TCefAutoplayPolicy ENUM
//
//	Autoplay policy types used by TCefApplicationCore.AutoplayPolicy. See the --autoplay-policy switch.
type TCefAutoplayPolicy = int32

const (
	AppDefault TCefAutoplayPolicy = iota
	AppDocumentUserActivationRequired
	AppNoUserGestureRequired
	AppUserGestureRequired
)

// TCefAxisAlignment ENUM
//
//	Specifies where along the axis the CefBoxLayout child views should be laid
//	out. Should be kept in sync with Chromium's views::LayoutAlignment type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_axis_alignment_t)</see>
type TCefAxisAlignment = int32

const (
	CEF_AXIS_ALIGNMENT_START TCefAxisAlignment = iota
	CEF_AXIS_ALIGNMENT_CENTER
	CEF_AXIS_ALIGNMENT_END
	CEF_AXIS_ALIGNMENT_STRETCH
)

// TCefBatterySaverModeState ENUM
//
//	Values used by the battery saver mode state preference.
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h">components/performance_manager/public/user_tuning/prefs.h</see>
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
//	Used by TCefBrowserNavigationTask to navigate in the right CEF thread.
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
//	Specifies the button display state.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_button_state_t)</see>
type TCefButtonState = int32

const (
	CEF_BUTTON_STATE_NORMAL TCefButtonState = iota
	CEF_BUTTON_STATE_HOVERED
	CEF_BUTTON_STATE_PRESSED
	CEF_BUTTON_STATE_DISABLED
)

// TCefChannelLayout ENUM
//
//	Enumerates the various representations of the ordering of audio channels.
//	Must be kept synchronized with media::ChannelLayout from Chromium.
//	See media\base\channel_layout.h
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_channel_layout_t)</see>
type TCefChannelLayout = int32

const (
	CEF_CHANNEL_LAYOUT_NONE                    TCefChannelLayout = iota
	CEF_CHANNEL_LAYOUT_UNSUPPORTED                               = 1
	CEF_CHANNEL_LAYOUT_MONO                                      = 2
	CEF_CHANNEL_LAYOUT_STEREO                                    = 3
	CEF_CHANNEL_LAYOUT_2_1                                       = 4
	CEF_CHANNEL_LAYOUT_SURROUND                                  = 5
	CEF_CHANNEL_LAYOUT_4_0                                       = 6
	CEF_CHANNEL_LAYOUT_2_2                                       = 7
	CEF_CHANNEL_LAYOUT_QUAD                                      = 8
	CEF_CHANNEL_LAYOUT_5_0                                       = 9
	CEF_CHANNEL_LAYOUT_5_1                                       = 10
	CEF_CHANNEL_LAYOUT_5_0_BACK                                  = 11
	CEF_CHANNEL_LAYOUT_5_1_BACK                                  = 12
	CEF_CHANNEL_LAYOUT_7_0                                       = 13
	CEF_CHANNEL_LAYOUT_7_1                                       = 14
	CEF_CHANNEL_LAYOUT_7_1_WIDE                                  = 15
	CEF_CHANNEL_LAYOUT_STEREO_DOWNMIX                            = 16
	CEF_CHANNEL_LAYOUT_2POINT1                                   = 17
	CEF_CHANNEL_LAYOUT_3_1                                       = 18
	CEF_CHANNEL_LAYOUT_4_1                                       = 19
	CEF_CHANNEL_LAYOUT_6_0                                       = 20
	CEF_CHANNEL_LAYOUT_6_0_FRONT                                 = 21
	CEF_CHANNEL_LAYOUT_HEXAGONAL                                 = 22
	CEF_CHANNEL_LAYOUT_6_1                                       = 23
	CEF_CHANNEL_LAYOUT_6_1_BACK                                  = 24
	CEF_CHANNEL_LAYOUT_6_1_FRONT                                 = 25
	CEF_CHANNEL_LAYOUT_7_0_FRONT                                 = 26
	CEF_CHANNEL_LAYOUT_7_1_WIDE_BACK                             = 27
	CEF_CHANNEL_LAYOUT_OCTAGONAL                                 = 28
	CEF_CHANNEL_LAYOUT_DISCRETE                                  = 29
	CEF_CHANNEL_LAYOUT_STEREO_AND_KEYBOARD_MIC                   = 30
	CEF_CHANNEL_LAYOUT_4_1_QUAD_SIDE                             = 31
	CEF_CHANNEL_LAYOUT_BITSTREAM                                 = 32
	CEF_CHANNEL_LAYOUT_5_1_4_DOWNMIX                             = 33
	CEF_CHANNEL_LAYOUT_1_1                                       = 34
	CEF_CHANNEL_LAYOUT_3_1_BACK                                  = 35
)

// TCefChromePageActionIconType ENUM
//
//	Chrome page action icon types. Should be kept in sync with Chromium's
//	PageActionIconType type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_page_action_icon_type_t)</see>
type TCefChromePageActionIconType = int32

const (
	CEF_CPAIT_BOOKMARK_STAR TCefChromePageActionIconType = iota
	CEF_CPAIT_CLICK_TO_CALL
	CEF_CPAIT_COOKIE_CONTROLS
	CEF_CPAIT_FILE_SYSTEM_ACCESS
	CEF_CPAIT_FIND
	CEF_CPAIT_MEMORY_SAVER
	CEF_CPAIT_INTENT_PICKER
	CEF_CPAIT_LOCAL_CARD_MIGRATION
	CEF_CPAIT_MANAGE_PASSWORDS
	CEF_CPAIT_PAYMENTS_OFFER_NOTIFICATION
	CEF_CPAIT_PRICE_TRACKING
	CEF_CPAIT_PWA_INSTALL
	CEF_CPAIT_QR_CODE_GENERATOR_DEPRECATED
	CEF_CPAIT_READER_MODE_DEPRECATED
	CEF_CPAIT_SAVE_AUTOFILL_ADDRESS
	CEF_CPAIT_SAVE_CARD
	CEF_CPAIT_SEND_TAB_TO_SELF_DEPRECATED
	CEF_CPAIT_SHARING_HUB
	CEF_CPAIT_SIDE_SEARCH
	CEF_CPAIT_SMS_REMOTE_FETCHER
	CEF_CPAIT_TRANSLATE
	CEF_CPAIT_VIRTUAL_CARD_ENROLL
	CEF_CPAIT_VIRTUAL_CARD_MANUAL_FALLBACK
	CEF_CPAIT_ZOOM
	CEF_CPAIT_SAVE_IBAN
	CEF_CPAIT_MANDATORY_REAUTH
	CEF_CPAIT_PRICE_INSIGHTS
	CEF_CPAIT_PRICE_READ_ANYTHING
	CEF_CPAIT_PRODUCT_SPECIFICATIONS
	CEF_CPAIT_LENS_OVERLAY
)

// TCefChromeToolbarButtonType ENUM
//
//	Chrome toolbar button types. Should be kept in sync with CEF's internal
//	ToolbarButtonType type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_chrome_toolbar_button_type_t)</see>
type TCefChromeToolbarButtonType = int32

const (
	CEF_CTBT_CAST TCefChromeToolbarButtonType = iota
	CEF_CTBT_DOWNLOAD
	CEF_CTBT_SEND_TAB_TO_SELF
	CEF_CTBT_SIDE_PANEL
)

// TCefClearDataStorageTypes ENUM
//
//	Storage types used by the Storage.clearDataForOrigin DevTools method in TChromiumCore.ClearDataForOrigin.
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
//	Action taken after the TChromium.Onclose event.
//	cbaCancel : stop closing the browser.
//	cbaClose : continue closing the browser.
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
//	Print job color mode values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_model_t)</see>
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
//	Describes how to interpret the components of a pixel.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_color.h">CEF source file: /include/internal/cef_types_color.h (cef_color_type_t)</see>
type TCefColorType = int32

const (
	CEF_COLOR_TYPE_RGBA_8888 TCefColorType = iota
	CEF_COLOR_TYPE_BGRA_8888
)

// TCefColorVariant ENUM
//
//	Specifies the color variants supported by
//	ICefRequestContext.SetChromeThemeColor.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_variant_t)</see>
type TCefColorVariant = int32

const (
	CEF_COLOR_VARIANT_SYSTEM TCefColorVariant = iota
	CEF_COLOR_VARIANT_LIGHT
	CEF_COLOR_VARIANT_DARK
	CEF_COLOR_VARIANT_TONAL_SPOT
	CEF_COLOR_VARIANT_NEUTRAL
	CEF_COLOR_VARIANT_VIBRANT
	CEF_COLOR_VARIANT_EXPRESSIVE
)

// TCefCOMInitMode ENUM
//
//	Flags used to customize the behavior of CefURLRequest.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_com_init_mode_t)</see>
type TCefCOMInitMode = int32

const (
	COM_INIT_MODE_NONE TCefCOMInitMode = iota
	COM_INIT_MODE_STA
	COM_INIT_MODE_MTA
)

// TCefCompositionUnderlineStyle ENUM
//
//	Composition underline style.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_composition_underline_style_t)</see>
type TCefCompositionUnderlineStyle = int32

const (
	CEF_CUS_SOLID TCefCompositionUnderlineStyle = iota
	CEF_CUS_DOT
	CEF_CUS_DASH
	CEF_CUS_NONE
)

// TCefContentSettingTypes ENUM
//
//	Supported content setting types. Some types are platform-specific or only
//	supported with the Chrome runtime. Should be kept in sync with Chromium's
//	ContentSettingsType type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_content_settings.h">CEF source file: /include/internal/cef_types_content_settings.h (cef_content_setting_types_t)</see>
type TCefContentSettingTypes = int32

const (
	CEF_CONTENT_SETTING_TYPE_COOKIES TCefContentSettingTypes = iota
	CEF_CONTENT_SETTING_TYPE_IMAGES
	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT
	CEF_CONTENT_SETTING_TYPE_POPUPS
	CEF_CONTENT_SETTING_TYPE_GEOLOCATION
	CEF_CONTENT_SETTING_TYPE_NOTIFICATIONS
	CEF_CONTENT_SETTING_TYPE_AUTO_SELECT_CERTIFICATE
	CEF_CONTENT_SETTING_TYPE_MIXEDSCRIPT
	CEF_CONTENT_SETTING_TYPE_MEDIASTREAM_MIC
	CEF_CONTENT_SETTING_TYPE_MEDIASTREAM_CAMERA
	CEF_CONTENT_SETTING_TYPE_PROTOCOL_HANDLERS
	CEF_CONTENT_SETTING_TYPE_DEPRECATED_PPAPI_BROKER
	CEF_CONTENT_SETTING_TYPE_AUTOMATIC_DOWNLOADS
	CEF_CONTENT_SETTING_TYPE_MIDI_SYSEX
	CEF_CONTENT_SETTING_TYPE_SSL_CERT_DECISIONS
	CEF_CONTENT_SETTING_TYPE_PROTECTED_MEDIA_IDENTIFIER
	CEF_CONTENT_SETTING_TYPE_APP_BANNER
	CEF_CONTENT_SETTING_TYPE_SITE_ENGAGEMENT
	CEF_CONTENT_SETTING_TYPE_DURABLE_STORAGE
	CEF_CONTENT_SETTING_TYPE_USB_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_GUARD
	CEF_CONTENT_SETTING_TYPE_BACKGROUND_SYNC
	CEF_CONTENT_SETTING_TYPE_AUTOPLAY
	CEF_CONTENT_SETTING_TYPE_IMPORTANT_SITE_INFO
	CEF_CONTENT_SETTING_TYPE_PERMISSION_AUTOBLOCKER_DATA
	CEF_CONTENT_SETTING_TYPE_ADS
	CEF_CONTENT_SETTING_TYPE_ADS_DATA
	CEF_CONTENT_SETTING_TYPE_MIDI
	CEF_CONTENT_SETTING_TYPE_PASSWORD_PROTECTION
	CEF_CONTENT_SETTING_TYPE_MEDIA_ENGAGEMENT
	CEF_CONTENT_SETTING_TYPE_SOUND
	CEF_CONTENT_SETTING_TYPE_CLIENT_HINTS
	CEF_CONTENT_SETTING_TYPE_SENSORS
	CEF_CONTENT_SETTING_TYPE_ACCESSIBILITY_EVENTS
	CEF_CONTENT_SETTING_TYPE_PAYMENT_HANDLER
	CEF_CONTENT_SETTING_TYPE_USB_GUARD
	CEF_CONTENT_SETTING_TYPE_BACKGROUND_FETCH
	CEF_CONTENT_SETTING_TYPE_INTENT_PICKER_DISPLAY
	CEF_CONTENT_SETTING_TYPE_IDLE_DETECTION
	CEF_CONTENT_SETTING_TYPE_SERIAL_GUARD
	CEF_CONTENT_SETTING_TYPE_SERIAL_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_PERIODIC_BACKGROUND_SYNC
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_SCANNING
	CEF_CONTENT_SETTING_TYPE_HID_GUARD
	CEF_CONTENT_SETTING_TYPE_HID_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SCREEN
	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SYSTEM
	CEF_CONTENT_SETTING_TYPE_LEGACY_COOKIE_ACCESS
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_WRITE_GUARD
	CEF_CONTENT_SETTING_TYPE_NFC
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_READ_WRITE
	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_SANITIZED_WRITE
	CEF_CONTENT_SETTING_TYPE_SAFE_BROWSING_URL_CHECK_DATA
	CEF_CONTENT_SETTING_TYPE_VR
	CEF_CONTENT_SETTING_TYPE_AR
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_READ_GUARD
	CEF_CONTENT_SETTING_TYPE_STORAGE_ACCESS
	CEF_CONTENT_SETTING_TYPE_CAMERA_PAN_TILT_ZOOM
	CEF_CONTENT_SETTING_TYPE_WINDOW_MANAGEMENT
	CEF_CONTENT_SETTING_TYPE_INSECURE_PRIVATE_NETWORK
	CEF_CONTENT_SETTING_TYPE_LOCAL_FONTS
	CEF_CONTENT_SETTING_TYPE_PERMISSION_AUTOREVOCATION_DATA
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_LAST_PICKED_DIRECTORY
	CEF_CONTENT_SETTING_TYPE_DISPLAY_CAPTURE
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_ACCESS_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_SHARING
	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT_JIT
	CEF_CONTENT_SETTING_TYPE_HTTP_ALLOWED
	CEF_CONTENT_SETTING_TYPE_FORMFILL_METADATA
	CEF_CONTENT_SETTING_TYPE_DEPRECATED_FEDERATED_IDENTITY_ACTIVE_SESSION
	CEF_CONTENT_SETTING_TYPE_AUTO_DARK_WEB_CONTENT
	CEF_CONTENT_SETTING_TYPE_REQUEST_DESKTOP_SITE
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_API
	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_INTERACTIONS
	CEF_CONTENT_SETTING_TYPE_REDUCED_ACCEPT_LANGUAGE
	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_PERMISSION_REVIEW
	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_GUARD
	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_SIGNIN_STATUS
	CEF_CONTENT_SETTING_TYPE_REVOKED_UNUSED_SITE_PERMISSIONS
	CEF_CONTENT_SETTING_TYPE_TOP_LEVEL_STORAGE_ACCESS
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_AUTO_REAUTHN_PERMISSION
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_REGISTRATION
	CEF_CONTENT_SETTING_TYPE_ANTI_ABUSE
	CEF_CONTENT_SETTING_TYPE_THIRD_PARTY_STORAGE_PARTITIONING
	CEF_CONTENT_SETTING_TYPE_HTTPS_ENFORCED
	CEF_CONTENT_SETTING_TYPE_ALL_SCREEN_CAPTURE
	CEF_CONTENT_SETTING_TYPE_COOKIE_CONTROLS_METADATA
	CEF_CONTENT_SETTING_TYPE_TPCD_HEURISTICS_GRANTS
	CEF_CONTENT_SETTING_TYPE_TPCD_METADATA_GRANTS
	CEF_CONTENT_SETTING_TYPE_TPCD_TRIAL
	CEF_CONTENT_SETTING_TYPE_TOP_LEVEL_TPCD_TRIAL
	CEF_CONTENT_SETTING_TOP_LEVEL_TPCD_ORIGIN_TRIAL
	CEF_CONTENT_SETTING_TYPE_AUTO_PICTURE_IN_PICTURE
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_ACCESS_EXTENDED_PERMISSION
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_ACCESS_RESTORE_PERMISSION
	CEF_CONTENT_SETTING_TYPE_CAPTURED_SURFACE_CONTROL
	CEF_CONTENT_SETTING_TYPE_SMART_CARD_GUARD
	CEF_CONTENT_SETTING_TYPE_SMART_CARD_DATA
	CEF_CONTENT_SETTING_TYPE_WEB_PRINTING
	CEF_CONTENT_SETTING_TYPE_AUTOMATIC_FULLSCREEN
	CEF_CONTENT_SETTING_TYPE_SUB_APP_INSTALLATION_PROMPTS
	CEF_CONTENT_SETTING_TYPE_SPEAKER_SELECTION
	CEF_CONTENT_SETTING_TYPE_DIRECT_SOCKETS
	CEF_CONTENT_SETTING_TYPE_KEYBOARD_LOCK
	CEF_CONTENT_SETTING_TYPE_POINTER_LOCK
	CEF_CONTENT_SETTING_TYPE_REVOKED_ABUSIVE_NOTIFICATION_PERMISSIONS
	CEF_CONTENT_SETTING_TYPE_TRACKING_PROTECTION
)

// TCefContentSettingValues ENUM
//
//	Supported content setting values. Should be kept in sync with Chromium's
//	ContentSetting type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_content_settings.h">CEF source file: /include/internal/cef_types_content_settings.h (cef_content_setting_values_t)</see>
type TCefContentSettingValues = int32

const (
	CEF_CONTENT_SETTING_VALUE_DEFAULT TCefContentSettingValues = iota
	CEF_CONTENT_SETTING_VALUE_ALLOW
	CEF_CONTENT_SETTING_VALUE_BLOCK
	CEF_CONTENT_SETTING_VALUE_ASK
	CEF_CONTENT_SETTING_VALUE_SESSION_ONLY
	CEF_CONTENT_SETTING_VALUE_DETECT_IMPORTANT_CONTENT
	CEF_CONTENT_SETTING_VALUE_NUM_VALUES
)

// TCefContextMenuMediaType ENUM
//
//	Supported context menu media types. These constants match their equivalents
//	in Chromium's ContextMenuDataMediaType and should not be renumbered.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_media_type_t)</see>
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
//	Cookie same site values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_same_site_t)</see>
type TCefCookieSameSite = int32

const (
	CEF_COOKIE_SAME_SITE_UNSPECIFIED TCefCookieSameSite = iota
	CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	CEF_COOKIE_SAME_SITE_LAX_MODE
	CEF_COOKIE_SAME_SITE_STRICT_MODE
)

// TCefCursorType ENUM
//
//	Cursor type values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cursor_type_t)</see>
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
//	Used by TCEFFileDialogInfo.
type TCEFDialogType = int32

const (
	DtOpen TCEFDialogType = iota
	DtOpenMultiple
	DtOpenFolder
	DtSave
)

// TCefDomDocumentType ENUM
//
//	DOM document types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_document_type_t)</see>
type TCefDomDocumentType = int32

const (
	DOM_DOCUMENT_TYPE_UNKNOWN TCefDomDocumentType = iota
	DOM_DOCUMENT_TYPE_HTML
	DOM_DOCUMENT_TYPE_XHTML
	DOM_DOCUMENT_TYPE_PLUGIN
)

// TCefDomEventPhase ENUM
//
//	DOM event processing phases.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_phase_t)</see>
type TCefDomEventPhase = int32

const (
	DOM_EVENT_PHASE_UNKNOWN TCefDomEventPhase = iota
	DOM_EVENT_PHASE_CAPTURING
	DOM_EVENT_PHASE_AT_TARGET
	DOM_EVENT_PHASE_BUBBLING
)

// TCefDomFormControlType ENUM
//
//	DOM form control types. Should be kept in sync with Chromium's
//	blink::mojom::FormControlType type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_form_control_type_t)</see>
type TCefDomFormControlType = int32

const (
	DOM_FORM_CONTROL_TYPE_UNSUPPORTED TCefDomFormControlType = iota
	DOM_FORM_CONTROL_TYPE_BUTTON_BUTTON
	DOM_FORM_CONTROL_TYPE_BUTTON_SUBMIT
	DOM_FORM_CONTROL_TYPE_BUTTON_RESET
	DOM_FORM_CONTROL_TYPE_BUTTON_SELECT_LIST
	DOM_FORM_CONTROL_TYPE_BUTTON_POPOVER
	DOM_FORM_CONTROL_TYPE_FIELDSET
	DOM_FORM_CONTROL_TYPE_INPUT_BUTTON
	DOM_FORM_CONTROL_TYPE_INPUT_CHECKBOX
	DOM_FORM_CONTROL_TYPE_INPUT_COLOR
	DOM_FORM_CONTROL_TYPE_INPUT_DATE
	DOM_FORM_CONTROL_TYPE_INPUT_DATETIME_LOCAL
	DOM_FORM_CONTROL_TYPE_INPUT_EMAIL
	DOM_FORM_CONTROL_TYPE_INPUT_FILE
	DOM_FORM_CONTROL_TYPE_INPUT_HIDDEN
	DOM_FORM_CONTROL_TYPE_INPUT_IMAGE
	DOM_FORM_CONTROL_TYPE_INPUT_MONTH
	DOM_FORM_CONTROL_TYPE_INPUT_NUMBER
	DOM_FORM_CONTROL_TYPE_INPUT_PASSWORD
	DOM_FORM_CONTROL_TYPE_INPUT_RADIO
	DOM_FORM_CONTROL_TYPE_INPUT_RANGE
	DOM_FORM_CONTROL_TYPE_INPUT_RESET
	DOM_FORM_CONTROL_TYPE_INPUT_SEARCH
	DOM_FORM_CONTROL_TYPE_INPUT_SUBMIT
	DOM_FORM_CONTROL_TYPE_INPUT_TELEPHONE
	DOM_FORM_CONTROL_TYPE_INPUT_TEXT
	DOM_FORM_CONTROL_TYPE_INPUT_TIME
	DOM_FORM_CONTROL_TYPE_INPUT_URL
	DOM_FORM_CONTROL_TYPE_INPUT_WEEK
	DOM_FORM_CONTROL_TYPE_OUTPUT
	DOM_FORM_CONTROL_TYPE_SELECT_ONE
	DOM_FORM_CONTROL_TYPE_SELECT_MULTIPLE
	DOM_FORM_CONTROL_TYPE_SELECT_LIST
	DOM_FORM_CONTROL_TYPE_TEXT_AREA
)

// TCefDomNodeType ENUM
//
//	DOM node types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_node_type_t)</see>
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

// TCefEditingCommand ENUM
//
//	Blink editing commands used by the "Input.dispatchKeyEvent" DevTools method.
//	<see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
//	<see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</see>
type TCefEditingCommand = int32

const (
	EcNone TCefEditingCommand = iota
	EcAlignCenter
	EcAlignJustified
	EcAlignLeft
	EcAlignRight
	EcBackColor
	EcBackwardDelete
	EcBold
	EcCopy
	EcCreateLink
	EcCut
	EcDefaultParagraphSeparator
	EcDelete
	EcDeleteBackward
	EcDeleteBackwardByDecomposingPreviousCharacter
	EcDeleteForward
	EcDeleteToBeginningOfLine
	EcDeleteToBeginningOfParagraph
	EcDeleteToEndOfLine
	EcDeleteToEndOfParagraph
	EcDeleteToMark
	EcDeleteWordBackward
	EcDeleteWordForward
	EcFindString
	EcFontName
	EcFontSize
	EcFontSizeDelta
	EcForeColor
	EcFormatBlock
	EcForwardDelete
	EcHiliteColor
	EcIgnoreSpelling
	EcIndent
	EcInsertBacktab
	EcInsertHorizontalRule
	EcInsertHTML
	EcInsertImage
	EcInsertLineBreak
	EcInsertNewline
	EcInsertNewlineInQuotedContent
	EcInsertOrderedList
	EcInsertParagraph
	EcInsertTab
	EcInsertText
	EcInsertUnorderedList
	EcItalic
	EcJustifyCenter
	EcJustifyFull
	EcJustifyLeft
	EcJustifyNone
	EcJustifyRight
	EcMakeTextWritingDirectionLeftToRight
	EcMakeTextWritingDirectionNatural
	EcMakeTextWritingDirectionRightToLeft
	EcMoveBackward
	EcMoveBackwardAndModifySelection
	EcMoveDown
	EcMoveDownAndModifySelection
	EcMoveForward
	EcMoveForwardAndModifySelection
	EcMoveLeft
	EcMoveLeftAndModifySelection
	EcMovePageDown
	EcMovePageDownAndModifySelection
	EcMovePageUp
	EcMovePageUpAndModifySelection
	EcMoveParagraphBackward
	EcMoveParagraphBackwardAndModifySelection
	EcMoveParagraphForward
	EcMoveParagraphForwardAndModifySelection
	EcMoveRight
	EcMoveRightAndModifySelection
	EcMoveToBeginningOfDocument
	EcMoveToBeginningOfDocumentAndModifySelection
	EcMoveToBeginningOfLine
	EcMoveToBeginningOfLineAndModifySelection
	EcMoveToBeginningOfParagraph
	EcMoveToBeginningOfParagraphAndModifySelection
	EcMoveToBeginningOfSentence
	EcMoveToBeginningOfSentenceAndModifySelection
	EcMoveToEndOfDocument
	EcMoveToEndOfDocumentAndModifySelection
	EcMoveToEndOfLine
	EcMoveToEndOfLineAndModifySelection
	EcMoveToEndOfParagraph
	EcMoveToEndOfParagraphAndModifySelection
	EcMoveToEndOfSentence
	EcMoveToEndOfSentenceAndModifySelection
	EcMoveToLeftEndOfLine
	EcMoveToLeftEndOfLineAndModifySelection
	EcMoveToRightEndOfLine
	EcMoveToRightEndOfLineAndModifySelection
	EcMoveUp
	EcMoveUpAndModifySelection
	EcMoveWordBackward
	EcMoveWordBackwardAndModifySelection
	EcMoveWordForward
	EcMoveWordForwardAndModifySelection
	EcMoveWordLeft
	EcMoveWordLeftAndModifySelection
	EcMoveWordRight
	EcMoveWordRightAndModifySelection
	EcOutdent
	EcOverWrite
	EcPaste
	EcPasteAndMatchStyle
	EcPasteGlobalSelection
	EcPrint
	EcRedo
	EcRemoveFormat
	EcScrollLineDown
	EcScrollLineUp
	EcScrollPageBackward
	EcScrollPageForward
	EcScrollToBeginningOfDocument
	EcScrollToEndOfDocument
	EcSelectAll
	EcSelectLine
	EcSelectParagraph
	EcSelectSentence
	EcSelectToMark
	EcSelectWord
	EcSetMark
	EcStrikethrough
	EcStyleWithCSS
	EcSubscript
	EcSuperscript
	EcSwapWithMark
	EcToggleBold
	EcToggleItalic
	EcToggleUnderline
	EcTranspose
	EcUnderline
	EcUndo
	EcUnlink
	EcUnscript
	EcUnselect
	EcUseCSS
	EcYank
	EcYankAndSelect
)

// TCefFocusSource ENUM
//
//	Focus sources.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_focus_source_t)</see>
type TCefFocusSource = int32

const (
	FOCUS_SOURCE_NAVIGATION TCefFocusSource = iota
	FOCUS_SOURCE_SYSTEM
)

// TCefGestureCommand ENUM
//
//	Specifies the gesture commands.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_gesture_command_t)</see>
type TCefGestureCommand = int32

const (
	CEF_GESTURE_COMMAND_BACK TCefGestureCommand = iota
	CEF_GESTURE_COMMAND_FORWARD
)

// TCefHighEfficiencyModeState ENUM
//
//	Values used by the high efficiency mode state preference.
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h">components/performance_manager/public/user_tuning/prefs.h</see>
type TCefHighEfficiencyModeState = int32

const (
	KDisabled       TCefHighEfficiencyModeState = iota
	KEnabled                                    = 1
	KEnabledOnTimer                             = 2
	KDefault                                    = 3
)

// TCefHorizontalAlignment ENUM
//
//	Specifies the horizontal text alignment mode.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_horizontal_alignment_t)</see>
type TCefHorizontalAlignment = int32

const (
	CEF_HORIZONTAL_ALIGNMENT_LEFT TCefHorizontalAlignment = iota
	CEF_HORIZONTAL_ALIGNMENT_CENTER
	CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// TCefJsDialogType ENUM
//
//	Supported JavaScript dialog types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_jsdialog_type_t)</see>
type TCefJsDialogType = int32

const (
	JSDIALOGTYPE_ALERT TCefJsDialogType = iota
	JSDIALOGTYPE_CONFIRM
	JSDIALOGTYPE_PROMPT
)

// TCefJsonParserOptions ENUM
//
//	Options that can be passed to CefParseJSON.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_parser_options_t)</see>
type TCefJsonParserOptions = int32

const (
	JSON_PARSER_RFC                   TCefJsonParserOptions = iota
	JSON_PARSER_ALLOW_TRAILING_COMMAS                       = 1 << 0
)

// TCefKeyEventType ENUM
//
//	Notification that a character was typed. Use this for text input. Key
//	down events may generate 0, 1, or more than one character event depending
//	on the key, locale, and operating system.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_key_event_type_t)</see>
type TCefKeyEventType = int32

const (
	KEYEVENT_RAWKEYDOWN TCefKeyEventType = iota
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

// TCefKeyLocation ENUM
//
//	Key location value used in the TChromiumCore.dispatchKeyEvent DevTools method.
//	<see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
type TCefKeyLocation = int32

const (
	CEF_KEYLOCATION_NONE TCefKeyLocation = iota
	CEF_KEYLOCATION_LEFT
	CEF_KEYLOCATION_RIGHT
)

// TCefMediaRouteConnectionState ENUM
//
//	Connection state for a MediaRoute object.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_connection_state_t)</see>
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
//	Icon types for a MediaSink object. Should be kept in sync with Chromium's
//	media_router::SinkIconType type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_sink_icon_type_t)</see>
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
//	Used by TCefMediaSinkInfo and TCefMediaSourceInfo.
type TCefMediaType = int32

const (
	MtCast TCefMediaType = iota
	MtDial
	MtUnknown
)

// TCefMenuAnchorPosition ENUM
//
//	Specifies how a menu will be anchored for non-RTL languages. The opposite
//	position will be used for RTL languages.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_anchor_position_t)</see>
type TCefMenuAnchorPosition = int32

const (
	CEF_MENU_ANCHOR_TOPLEFT TCefMenuAnchorPosition = iota
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

// TCefMenuColorType ENUM
//
//	Supported color types for menu items.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_color_type_t)</see>
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
//	Supported menu item types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_item_type_t)</see>
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
//	Flags used to customize the behavior of CefURLRequest.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_message_loop_type_t)</see>
type TCefMessageLoopType = int32

const (
	ML_TYPE_DEFAULT TCefMessageLoopType = iota
	ML_TYPE_UI
	ML_TYPE_IO
)

// TCefMouseButtonType ENUM
//
//	Mouse button types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_mouse_button_type_t)</see>
type TCefMouseButtonType = int32

const (
	MBT_LEFT TCefMouseButtonType = iota
	MBT_MIDDLE
	MBT_RIGHT
)

// TCefNavigationType ENUM
//
//	Navigation types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_navigation_type_t)</see>
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
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:content/browser/network_service_instance_impl.cc">network_service_instance_impl.cc</see>
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:net/log/net_log_capture_mode.h">net_log_capture_mode.h</see>
type TCefNetLogCaptureMode = int32

const (
	NlcmDefault TCefNetLogCaptureMode = iota
	NlcmIncludeSensitive
	NlcmEverything
)

// TCefPaintElementType ENUM
//
//	Paint element types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_paint_element_type_t)</see>
type TCefPaintElementType = int32

const (
	PET_VIEW TCefPaintElementType = iota
	PET_POPUP
)

// TCefPathKey ENUM
//
//	Process termination status values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_path_key_t)</see>
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
//	Margin type for PDF printing.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_pdf_print_margin_type_t)</see>
type TCefPdfPrintMarginType = int32

const (
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	PDF_PRINT_MARGIN_NONE
	PDF_PRINT_MARGIN_CUSTOM
)

// TCefPermissionRequestResult ENUM
//
//	Permission request results.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_result_t)</see>
type TCefPermissionRequestResult = int32

const (
	CEF_PERMISSION_RESULT_ACCEPT TCefPermissionRequestResult = iota
	CEF_PERMISSION_RESULT_DENY
	CEF_PERMISSION_RESULT_DISMISS
	CEF_PERMISSION_RESULT_IGNORE
)

// TCefPointerType ENUM
//
//	The device type that caused the event.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_pointer_type_t)</see>
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
//	Post data elements may represent either bytes or files.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_postdataelement_type_t)</see>
type TCefPostDataElementType = int32

const (
	PDE_TYPE_EMPTY TCefPostDataElementType = iota
	PDE_TYPE_BYTES
	PDE_TYPE_FILE
)

// TCefPreferencesType ENUM
//
//	Preferences type passed to
//	ICefBrowserProcessHandler.OnRegisterCustomPreferences.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_preferences_type_t)</see>
type TCefPreferencesType = int32

const (
	CEF_PREFERENCES_TYPE_GLOBAL TCefPreferencesType = iota
	CEF_PREFERENCES_TYPE_REQUEST_CONTEXT
)

// TCefProcessId ENUM
//
//	Existing process IDs.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_process_id_t)</see>
type TCefProcessId = int32

const (
	PID_BROWSER TCefProcessId = iota
	PID_RENDERER
)

// TCefProcessType ENUM
//
//	Sub-process types of Chromium.
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
//
//	Supported proxy schemes in Chromium.
type TCefProxyScheme = int32

const (
	PsHTTP TCefProxyScheme = iota
	PsSOCKS4
	PsSOCKS5
)

// TCefReferrerPolicy ENUM
//
//	Policy for how the Referrer HTTP header value will be sent during
//	navigation. If the `--no-referrers` command-line flag is specified then the
//	policy value will be ignored and the Referrer value will never be sent. Must
//	be kept synchronized with net::URLRequest::ReferrerPolicy from Chromium.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_referrer_policy_t)</see>
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
//	Resource type for a request. These constants match their equivalents in
//	Chromium's ResourceType and should not be renumbered.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resource_type_t)</see>
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
//	Return values for ICefResponseFilter.Filter().
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_response_filter_status_t)</see>
type TCefResponseFilterStatus = int32

const (
	RESPONSE_FILTER_NEED_MORE_DATA TCefResponseFilterStatus = iota
	RESPONSE_FILTER_DONE
	RESPONSE_FILTER_ERROR
)

// TCefReturnValue ENUM
//
//	Return value types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_return_value_t)</see>
type TCefReturnValue = int32

const (
	RV_CANCEL TCefReturnValue = iota
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

// TCefRuntimeStyle ENUM
//
//	CEF supports both a Chrome runtime (based on the Chrome UI layer) and an
//	Alloy runtime (based on the Chromium content layer). The Chrome runtime
//	provides the full Chrome UI and browser functionality whereas the Alloy
//	runtime provides less default browser functionality but adds additional
//	client callbacks and support for windowless (off-screen) rendering. For
//	additional comparative details on runtime types see
//	https://bitbucket.org/chromiumembedded/cef/wiki/Architecture.md#markdown-header-cef3
//
//	Each runtime is composed of a bootstrap component and a style component. The
//	bootstrap component is configured via CefSettings.chrome_runtime and cannot
//	be changed after CefInitialize. The style component is individually
//	configured for each window/browser at creation time and, in combination with
//	the Chrome bootstrap, different styles can be mixed during runtime.
//
//	Windowless rendering will always use Alloy style. Windowed rendering with a
//	default window or client-provided parent window can configure the style via
//	CefWindowInfo.runtime_style. Windowed rendering with the Views framework can
//	configure the style via CefWindowDelegate::GetWindowRuntimeStyle and
//	CefBrowserViewDelegate::GetBrowserRuntimeStyle. Alloy style Windows with the
//	Views framework can host only Alloy style BrowserViews but Chrome style
//	Windows can host both style BrowserViews. Additionally, a Chrome style
//	Window can host at most one Chrome style BrowserView but potentially
//	multiple Alloy style BrowserViews. See CefWindowInfo.runtime_style
//	documentation for any additional platform-specific limitations.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_runtime.h">CEF source file: /include/internal/cef_types_runtime.h (cef_runtime_style_t)</see>
type TCefRuntimeStyle = int32

const (
	CEF_RUNTIME_STYLE_DEFAULT TCefRuntimeStyle = iota
	CEF_RUNTIME_STYLE_CHROME
	CEF_RUNTIME_STYLE_ALLOY
)

// TCefScaleFactor ENUM
//
//	Supported UI scale factors for the platform. SCALE_FACTOR_NONE is used for
//	density independent resources such as string, html/js files or an image that
//	can be used for any scale factors (such as wallpapers).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scale_factor_t)</see>
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

// TCefSimulatedMouseButton ENUM
//
//	Mouse button in the TChromiumCore.SimulateMouseEvent function.
//	<see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see>
type TCefSimulatedMouseButton = int32

const (
	CEF_SIMULATEDMOUSEBUTTON_NONE TCefSimulatedMouseButton = iota
	CEF_SIMULATEDMOUSEBUTTON_LEFT
	CEF_SIMULATEDMOUSEBUTTON_MIDDLE
	CEF_SIMULATEDMOUSEBUTTON_RIGHT
	CEF_SIMULATEDMOUSEBUTTON_BACK
	CEF_SIMULATEDMOUSEBUTTON_FORWARD
)

// TCefSimulatedMouseEventType ENUM
//
//	Type of mouse event in the TChromiumCore.SimulateMouseEvent function.
//	<see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see>
type TCefSimulatedMouseEventType = int32

const (
	MousePressed TCefSimulatedMouseEventType = iota
	MouseReleased
	MouseMoved
	MouseWheel
)

// TCefSimulatedPointerType ENUM
//
//	Pointer type in the TChromiumCore.SimulateMouseEvent function.
//	<see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see>
type TCefSimulatedPointerType = int32

const (
	CEF_SIMULATEDPOINTERTYPE_MOUSE TCefSimulatedPointerType = iota
	CEF_SIMULATEDPOINTERTYPE_PEN
)

// TCefSimulatedTouchEventType ENUM
//
//	Type of touch event in the TChromiumCore.SimulateTouchEvent function.
//	<see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent">See the Input.dispatchTouchEvent DevTools method</see>
type TCefSimulatedTouchEventType = int32

const (
	TouchStart TCefSimulatedTouchEventType = iota
	TouchEnd
	TouchMove
	TouchCancel
)

// TCefState ENUM
//
//	Represents the state of a setting.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_state_t)</see>
type TCefState = int32

const (
	STATE_DEFAULT TCefState = iota
	STATE_ENABLED
	STATE_DISABLED
)

// TCefStorageType ENUM
//
//	Storage types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_storage_type_t)</see>
type TCefStorageType = int32

const (
	ST_LOCALSTORAGE TCefStorageType = iota
	ST_SESSIONSTORAGE
)

// TCefTerminationStatus ENUM
//
//	Process termination status values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_termination_status_t)</see>
type TCefTerminationStatus = int32

const (
	TS_ABNORMAL_TERMINATION TCefTerminationStatus = iota
	TS_PROCESS_WAS_KILLED
	TS_PROCESS_CRASHED
	TS_PROCESS_OOM
	TS_LAUNCH_FAILED
	TS_INTEGRITY_FAILURE
)

// TCefTestCertType ENUM
//
//	Specifies the gesture commands.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_test_cert_type_t)</see>
type TCefTestCertType = int32

const (
	CEF_TEST_CERT_OK_IP TCefTestCertType = iota
	CEF_TEST_CERT_OK_DOMAIN
	CEF_TEST_CERT_EXPIRED
)

// TCefTextInpuMode ENUM
//
//	Input mode of a virtual keyboard. These constants match their equivalents
//	in Chromium's text_input_mode.h and should not be renumbered.
//	See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_input_mode_t)</see>
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
//	Text style types. Should be kepy in sync with gfx::TextStyle.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_style_t)</see>
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
//	Existing thread IDs.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_thread_id_t)</see>
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
//	Thread priority values listed in increasing order of importance.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_thread_priority_t)</see>
type TCefThreadPriority = int32

const (
	TP_BACKGROUND TCefThreadPriority = iota
	TP_NORMAL
	TP_DISPLAY
	TP_REALTIME_AUDIO
)

// TCefTouchEeventType ENUM
//
//	Touch points states types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_event_type_t)</see>
type TCefTouchEeventType = int32

const (
	CEF_TET_RELEASED TCefTouchEeventType = iota
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

// TCefUIColorMode ENUM
//
//	Color mode in UI for platforms that support it.
type TCefUIColorMode = int32

const (
	UicmSystemDefault TCefUIColorMode = iota
	UicmForceLight
	UicmForceDark
)

// TCefUrlRequestStatus ENUM
//
//	Flags that represent CefURLRequest status.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_status_t)</see>
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
//	Supported value types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_value_type_t)</see>
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
//
//	WebRTC handling policy types used by TChromiumCore.WebRTCIPHandlingPolicy.
type TCefWebRTCHandlingPolicy = int32

const (
	HpDefault TCefWebRTCHandlingPolicy = iota
	HpDefaultPublicAndPrivateInterfaces
	HpDefaultPublicInterfaceOnly
	HpDisableNonProxiedUDP
)

// TCefWindowOpenDisposition ENUM
//
//	The manner in which a link click should be opened. These constants match
//	their equivalents in Chromium's window_open_disposition.h and should not be
//	renumbered.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_window_open_disposition_t)</see>
type TCefWindowOpenDisposition = int32

const (
	CEF_WOD_UNKNOWN TCefWindowOpenDisposition = iota
	CEF_WOD_CURRENT_TAB
	CEF_WOD_SINGLETON_TAB
	CEF_WOD_NEW_FOREGROUND_TAB
	CEF_WOD_NEW_BACKGROUND_TAB
	CEF_WOD_NEW_POPUP
	CEF_WOD_NEW_WINDOW
	CEF_WOD_SAVE_TO_DISK
	CEF_WOD_OFF_THE_RECORD
	CEF_WOD_IGNORE_ACTION
	CEF_WOD_SWITCH_TO_TAB
	CEF_WOD_NEW_PICTURE_IN_PICTURE
)

// TCefXmlEncodingType ENUM
//
//	Supported XML encoding types. The parser supports ASCII, ISO-8859-1, and
//	UTF16 (LE and BE) by default. All other types must be translated to UTF8
//	before being passed to the parser. If a BOM is detected and the correct
//	decoder is available then that decoder will be used automatically.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_xml_encoding_type_t)</see>
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
//	XML node types.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_xml_node_type_t)</see>
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

// TCefZoomCommand ENUM
//
//	Specifies the zoom commands supported by ICefBrowserHost.Zoom.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_zoom_command_t)</see>
type TCefZoomCommand = int32

const (
	CEF_ZOOM_COMMAND_OUT TCefZoomCommand = iota
	CEF_ZOOM_COMMAND_RESET
	CEF_ZOOM_COMMAND_IN
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

// TSimulatedCefKeyEventType ENUM
//
//	Event type used by TChromiumCore.SimulateKeyEvent
type TSimulatedCefKeyEventType = int32

const (
	KetKeyDown TSimulatedCefKeyEventType = iota
	KetKeyUp
	KetRawKeyDown
	KetChar
)
