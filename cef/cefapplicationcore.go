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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/types"
)

// IApplicationCoreEvents Parent: IObject
type IApplicationCoreEvents interface {
	IObject
}

// ICefApplicationCore Parent: IApplicationCoreEvents IInterfacedObject
type ICefApplicationCore interface {
	IApplicationCoreEvents
	IInterfacedObject
	// CheckCEFLibrary
	//  Used to check the CEF binaries manually.
	CheckCEFLibrary() bool // function
	// StartMainProcess
	//  Used to initialize CEF in the main browser process. In case CEF is
	//  configured to used the same executable for all processes then all
	//  processes must call this function. CEF can only be initialized once
	//  per process. This is a CEF feature and there's no workaround. This
	//  function returns immediately in when called in the main process and
	//  it blocks the execution when it's called from a CEF subprocess until
	//  that process ends.
	StartMainProcess() bool // function
	// StartSubProcess
	//  Used to initialize CEF in the subprocesses. This function can only be
	//  used when CEF is configured to use a different executable for the
	//  subprocesses. This function blocks the execution until the process ends.
	StartSubProcess() bool // function
	// ValidComponentID
	//  Returns true if a custom component ID is valid before executing a CEF task.
	ValidComponentID(componentID int32) bool // function
	// NextComponentID
	//  Returns the next component ID and adds this value to the valid ID list.
	NextComponentID() int32 // function
	// AddCustomCommandLine
	//  Used to add any command line switch that is not available as a
	//  TCEFApplicationCore property.
	AddCustomCommandLine(commandLine string, value string) // procedure
	// DoMessageLoopWork
	//  Perform a single iteration of CEF message loop processing. This function is
	//  provided for cases where the CEF message loop must be integrated into an
	//  existing application message loop. Use of this function is not recommended
	//  for most users; use either the RunMessageLoop function or
	//  TCefSettings.multi_threaded_message_loop if possible. When using this
	//  function care must be taken to balance performance against excessive CPU
	//  usage. It is recommended to enable the TCefSettings.external_message_pump
	//  option when using this function so that
	//  ICefBrowserProcessHandler.OnScheduleMessagePumpWork callbacks can
	//  facilitate the scheduling process. This function should only be called on
	//  the main application thread and only if cef_initialize() is called with a
	//  TCefSettings.multi_threaded_message_loop value of false (0). This function
	//  will not block.
	DoMessageLoopWork() // procedure
	// RunMessageLoop
	//  Run the CEF message loop. Use this function instead of an application-
	//  provided message loop to get the best balance between performance and CPU
	//  usage. This function should only be called on the main application thread
	//  and only if cef_initialize() is called with a
	//  TCefSettings.multi_threaded_message_loop value of false (0). This function
	//  will block until a quit message is received by the system.
	RunMessageLoop() // procedure
	// QuitMessageLoop
	//  Quit the CEF message loop that was started by calling
	//  RunMessageLoop. This function should only be called on the main
	//  application thread and only if RunMessageLoop was used.
	QuitMessageLoop() // procedure
	// UpdateDeviceScaleFactor
	//  Update the DeviceScaleFactor value with the current monitor scale.
	UpdateDeviceScaleFactor() // procedure
	// InitLibLocationFromArgs
	//  This procedure is only available in MacOS to read some configuration
	//  settings from the command line arguments.
	InitLibLocationFromArgs() // procedure
	// RemoveComponentID
	//  Removes a component ID from the valid ID list when a component is destroyed.
	RemoveComponentID(componentID int32) // procedure
	// NoSandbox
	//  Set to true (1) to disable the sandbox for sub-processes. See
	//  cef_sandbox_win.h for requirements to enable the sandbox on Windows. Also
	//  configurable using the "no-sandbox" command-line switch.
	NoSandbox() bool         // property NoSandbox Getter
	SetNoSandbox(value bool) // property NoSandbox Setter
	// BrowserSubprocessPath
	//  The path to a separate executable that will be launched for sub-processes.
	//  If this value is empty on Windows or Linux then the main process
	//  executable will be used. If this value is empty on macOS then a helper
	//  executable must exist at "Contents/Frameworks/<app>
	//  Helper.app/Contents/MacOS/<app> Helper" in the top-level app bundle. See
	//  the comments on CefExecuteProcess() for details. If this value is
	//  non-empty then it must be an absolute path. Also configurable using the
	//  "browser-subprocess-path" command-line switch.
	BrowserSubprocessPath() string         // property BrowserSubprocessPath Getter
	SetBrowserSubprocessPath(value string) // property BrowserSubprocessPath Setter
	// FrameworkDirPath
	//  The path to the CEF framework directory on macOS. If this value is empty
	//  then the framework must exist at "Contents/Frameworks/Chromium Embedded
	//  Framework.framework" in the top-level app bundle. If this value is
	//  non-empty then it must be an absolute path. Also configurable using the
	//  "framework-dir-path" command-line switch.
	FrameworkDirPath() string         // property FrameworkDirPath Getter
	SetFrameworkDirPath(value string) // property FrameworkDirPath Setter
	// MainBundlePath
	//  The path to the main bundle on macOS. If this value is empty then it
	//  defaults to the top-level app bundle. If this value is non-empty then it
	//  must be an absolute path. Also configurable using the "main-bundle-path"
	//  command-line switch.
	MainBundlePath() string         // property MainBundlePath Getter
	SetMainBundlePath(value string) // property MainBundlePath Setter
	// ChromeRuntime
	//  Set to true (1) to enable use of the Chrome runtime in CEF. This feature
	//  is considered experimental and is not recommended for most users at this
	//  time. See issue #2969 for details.
	ChromeRuntime() bool         // property ChromeRuntime Getter
	SetChromeRuntime(value bool) // property ChromeRuntime Setter
	// MultiThreadedMessageLoop
	//  Set to true (1) to have the browser process message loop run in a separate
	//  thread. If false (0) then the CefDoMessageLoopWork() function must be
	//  called from your application message loop. This option is only supported
	//  on Windows and Linux.
	MultiThreadedMessageLoop() bool         // property MultiThreadedMessageLoop Getter
	SetMultiThreadedMessageLoop(value bool) // property MultiThreadedMessageLoop Setter
	// ExternalMessagePump
	//  Set to true (1) to control browser process main (UI) thread message pump
	//  scheduling via the ICefBrowserProcessHandler.OnScheduleMessagePumpWork()
	//  callback. This option is recommended for use in combination with the
	//  CefDoMessageLoopWork() function in cases where the CEF message loop must
	//  be integrated into an existing application message loop (see additional
	//  comments and warnings on CefDoMessageLoopWork). Enabling this option is
	//  not recommended for most users; leave this option disabled and use either
	//  the CefRunMessageLoop() function or multi_threaded_message_loop if
	//  possible.
	ExternalMessagePump() bool         // property ExternalMessagePump Getter
	SetExternalMessagePump(value bool) // property ExternalMessagePump Setter
	// WindowlessRenderingEnabled
	//  Set to true (1) to enable windowless (off-screen) rendering support. Do
	//  not enable this value if the application does not use windowless rendering
	//  as it may reduce rendering performance on some systems.
	WindowlessRenderingEnabled() bool         // property WindowlessRenderingEnabled Getter
	SetWindowlessRenderingEnabled(value bool) // property WindowlessRenderingEnabled Setter
	// CommandLineArgsDisabled
	//  Set to true (1) to disable configuration of browser process features using
	//  standard CEF and Chromium command-line arguments. Configuration can still
	//  be specified using CEF data structures or via the
	//  ICefApp.OnBeforeCommandLineProcessing() method.
	CommandLineArgsDisabled() bool         // property CommandLineArgsDisabled Getter
	SetCommandLineArgsDisabled(value bool) // property CommandLineArgsDisabled Setter
	// Cache
	//  The directory where data for the global browser cache will be stored on
	//  disk. If this value is non-empty then it must be an absolute path that is
	//  either equal to or a child directory of TCefSettings.root_cache_path. If
	//  this value is empty then browsers will be created in "incognito mode"
	//  where in-memory caches are used for storage and no profile-specific data
	//  is persisted to disk (installation-specific data will still be persisted
	//  in root_cache_path). HTML5 databases such as localStorage will only
	//  persist across sessions if a cache path is specified. Can be overridden
	//  for individual ICefRequestContext instances via the
	//  ICefRequestContextSettings.cache_path value. When using the Chrome runtime
	//  any child directory value will be ignored and the "default" profile (also
	//  a child directory) will be used instead.
	Cache() string         // property Cache Getter
	SetCache(value string) // property Cache Setter
	// RootCache
	//  The root directory for installation-specific data and the parent directory
	//  for profile-specific data. All TCefSettings.cache_path and
	//  ICefRequestContextSettings.cache_path values must have this parent
	//  directory in common. If this value is empty and TCefSettings.cache_path is
	//  non-empty then it will default to the TCefSettings.cache_path value. Any
	//  non-empty value must be an absolute path. If both values are empty then
	//  the default platform-specific directory will be used
	//  ("~/.config/cef_user_data" directory on Linux, "~/Library/Application
	//  Support/CEF/User Data" directory on MacOS, "AppData\Local\CEF\User Data"
	//  directory under the user profile directory on Windows). Use of the default
	//  directory is not recommended in production applications (see below).
	//  Multiple application instances writing to the same root_cache_path
	//  directory could result in data corruption. A process singleton lock based
	//  on the root_cache_path value is therefore used to protect against this.
	//  This singleton behavior applies to all CEF-based applications using
	//  version 120 or newer. You should customize root_cache_path for your
	//  application and implement ICefBrowserProcessHandler.OnAlreadyRunningAppRelaunch,
	//  which will then be called on any app relaunch
	//  with the same root_cache_path value.
	//  Failure to set the root_cache_path value correctly may result in startup
	//  crashes or other unexpected behaviors (for example, the sandbox blocking
	//  read/write access to certain files).
	RootCache() string         // property RootCache Getter
	SetRootCache(value string) // property RootCache Setter
	// PersistSessionCookies
	//  To persist session cookies (cookies without an expiry date or validity
	//  interval) by default when using the global cookie manager set this value
	//  to true (1). Session cookies are generally intended to be transient and
	//  most Web browsers do not persist them. A |cache_path| value must also be
	//  specified to enable this feature. Also configurable using the
	//  "persist-session-cookies" command-line switch. Can be overridden for
	//  individual CefRequestContext instances via the
	//  TCefRequestContextSettings.persist_session_cookies value.
	PersistSessionCookies() bool         // property PersistSessionCookies Getter
	SetPersistSessionCookies(value bool) // property PersistSessionCookies Setter
	// PersistUserPreferences
	//  To persist user preferences as a JSON file in the cache path directory set
	//  this value to true (1). A |cache_path| value must also be specified
	//  to enable this feature. Also configurable using the
	//  "persist-user-preferences" command-line switch. Can be overridden for
	//  individual CefRequestContext instances via the
	//  TCefRequestContextSettings.persist_user_preferences value.
	PersistUserPreferences() bool         // property PersistUserPreferences Getter
	SetPersistUserPreferences(value bool) // property PersistUserPreferences Setter
	// UserAgent
	//  Value that will be returned as the User-Agent HTTP header. If empty the
	//  default User-Agent string will be used. Also configurable using the
	//  "user-agent" command-line switch.
	UserAgent() string         // property UserAgent Getter
	SetUserAgent(value string) // property UserAgent Setter
	// UserAgentProduct
	//  Value that will be inserted as the product portion of the default
	//  User-Agent string. If empty the Chromium product version will be used. If
	//  |userAgent| is specified this value will be ignored. Also configurable
	//  using the "user-agent-product" command-line switch.
	UserAgentProduct() string         // property UserAgentProduct Getter
	SetUserAgentProduct(value string) // property UserAgentProduct Setter
	// Locale
	//  The locale string that will be passed to WebKit. If empty the default
	//  locale of "en-US" will be used. This value is ignored on Linux where
	//  locale is determined using environment variable parsing with the
	//  precedence order: LANGUAGE, LC_ALL, LC_MESSAGES and LANG. Also
	//  configurable using the "lang" command-line switch.
	Locale() string         // property Locale Getter
	SetLocale(value string) // property Locale Setter
	// LogFile
	//  The directory and file name to use for the debug log. If empty a default
	//  log file name and location will be used. On Windows and Linux a
	//  "debug.log" file will be written in the main executable directory. On
	//  MacOS a "~/Library/Logs/[app name]_debug.log" file will be written where
	//  [app name] is the name of the main app executable. Also configurable using
	//  the "log-file" command-line switch.
	LogFile() string         // property LogFile Getter
	SetLogFile(value string) // property LogFile Setter
	// LogSeverity
	//  The log severity. Only messages of this severity level or higher will be
	//  logged. When set to DISABLE no messages will be written to the log file,
	//  but FATAL messages will still be output to stderr. Also configurable using
	//  the "log-severity" command-line switch with a value of "verbose", "info",
	//  "warning", "error", "fatal" or "disable".
	LogSeverity() cefTypes.TCefLogSeverity         // property LogSeverity Getter
	SetLogSeverity(value cefTypes.TCefLogSeverity) // property LogSeverity Setter
	// LogItems
	//  The log items prepended to each log line. If not set the default log items
	//  will be used. Also configurable using the "log-items" command-line switch
	//  with a value of "none" for no log items, or a comma-delimited list of
	//  values "pid", "tid", "timestamp" or "tickcount" for custom log items.
	LogItems() cefTypes.TCefLogItems         // property LogItems Getter
	SetLogItems(value cefTypes.TCefLogItems) // property LogItems Setter
	// JavaScriptFlags
	//  Custom flags that will be used when initializing the V8 JavaScript engine.
	//  The consequences of using custom flags may not be well tested. Also
	//  configurable using the "js-flags" command-line switch.
	JavaScriptFlags() string         // property JavaScriptFlags Getter
	SetJavaScriptFlags(value string) // property JavaScriptFlags Setter
	// ResourcesDirPath
	//  The fully qualified path for the resources directory. If this value is
	//  empty the *.pak files must be located in the module directory on
	//  Windows/Linux or the app bundle Resources directory on MacOS. If this
	//  value is non-empty then it must be an absolute path. Also configurable
	//  using the "resources-dir-path" command-line switch.
	ResourcesDirPath() string         // property ResourcesDirPath Getter
	SetResourcesDirPath(value string) // property ResourcesDirPath Setter
	// LocalesDirPath
	//  The fully qualified path for the locales directory. If this value is empty
	//  the locales directory must be located in the module directory. If this
	//  value is non-empty then it must be an absolute path. This value is ignored
	//  on MacOS where pack files are always loaded from the app bundle Resources
	//  directory. Also configurable using the "locales-dir-path" command-line
	//  switch.
	LocalesDirPath() string         // property LocalesDirPath Getter
	SetLocalesDirPath(value string) // property LocalesDirPath Setter
	// PackLoadingDisabled
	//  Set to true (1) to disable loading of pack files for resources and
	//  locales. A resource bundle handler must be provided for the browser and
	//  render processes via ICefApp.GetResourceBundleHandler() if loading of pack
	//  files is disabled. Also configurable using the "disable-pack-loading"
	//  command- line switch.
	PackLoadingDisabled() bool         // property PackLoadingDisabled Getter
	SetPackLoadingDisabled(value bool) // property PackLoadingDisabled Setter
	// RemoteDebuggingPort
	//  Set to a value between 1024 and 65535 to enable remote debugging on the
	//  specified port. Also configurable using the "remote-debugging-port"
	//  command-line switch. Specifying 0 via the command-line switch will result
	//  in the selection of an ephemeral port and the port number will be printed
	//  as part of the WebSocket endpoint URL to stderr. If a cache directory path
	//  is provided the port will also be written to the
	//  <cache-dir>/DevToolsActivePort file. Remote debugging can be accessed by
	//  loading the chrome://inspect page in Google Chrome. Port numbers 9222 and
	//  9229 are discoverable by default. Other port numbers may need to be
	//  configured via "Discover network targets" on the Devices tab.
	RemoteDebuggingPort() int32         // property RemoteDebuggingPort Getter
	SetRemoteDebuggingPort(value int32) // property RemoteDebuggingPort Setter
	// UncaughtExceptionStackSize
	//  The number of stack trace frames to capture for uncaught exceptions.
	//  Specify a positive value to enable the
	//  ICefRenderProcessHandler.OnUncaughtException() callback. Specify 0
	//  (default value) and OnUncaughtException() will not be called. Also
	//  configurable using the "uncaught-exception-stack-size" command-line
	//  switch.
	UncaughtExceptionStackSize() int32         // property UncaughtExceptionStackSize Getter
	SetUncaughtExceptionStackSize(value int32) // property UncaughtExceptionStackSize Setter
	// BackgroundColor
	//  Background color used for the browser before a document is loaded and when
	//  no document color is specified. The alpha component must be either fully
	//  opaque (0xFF) or fully transparent (0x00). If the alpha component is fully
	//  opaque then the RGB components will be used as the background color. If
	//  the alpha component is fully transparent for a windowed browser then the
	//  default value of opaque white be used. If the alpha component is fully
	//  transparent for a windowless (off-screen) browser then transparent
	//  painting will be enabled.
	BackgroundColor() cefTypes.TCefColor         // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor) // property BackgroundColor Setter
	// AcceptLanguageList
	//  Comma delimited ordered list of language codes without any whitespace that
	//  will be used in the "Accept-Language" HTTP request header and
	//  "navigator.language" JS attribute. Can be overridden for individual
	//  ICefRequestContext instances via the
	//  TCefRequestContextSettingsCefRequestContextSettings.accept_language_list value.
	AcceptLanguageList() string         // property AcceptLanguageList Getter
	SetAcceptLanguageList(value string) // property AcceptLanguageList Setter
	// CookieableSchemesList
	//  Comma delimited list of schemes supported by the associated
	//  ICefCookieManager. If |cookieable_schemes_exclude_defaults| is false (0)
	//  the default schemes ("http", "https", "ws" and "wss") will also be
	//  supported. Not specifying a |cookieable_schemes_list| value and setting
	//  |cookieable_schemes_exclude_defaults| to true (1) will disable all loading
	//  and saving of cookies. These settings will only impact the global
	//  ICefRequestContext. Individual ICefRequestContext instances can be
	//  configured via the TCefRequestContextSettings.cookieable_schemes_list and
	//  TCefRequestContextSettings.cookieable_schemes_exclude_defaults values.
	CookieableSchemesList() string         // property CookieableSchemesList Getter
	SetCookieableSchemesList(value string) // property CookieableSchemesList Setter
	// CookieableSchemesExcludeDefaults
	//  See the CookieableSchemesList property.
	CookieableSchemesExcludeDefaults() bool         // property CookieableSchemesExcludeDefaults Getter
	SetCookieableSchemesExcludeDefaults(value bool) // property CookieableSchemesExcludeDefaults Setter
	// ChromePolicyId
	//  Specify an ID to enable Chrome policy management via Platform and OS-user
	//  policies. On Windows, this is a registry key like
	//  "SOFTWARE\\Policies\\Google\\Chrome". On MacOS, this is a bundle ID like
	//  "com.google.Chrome". On Linux, this is an absolute directory path like
	//  "/etc/opt/chrome/policies". Only supported with the Chrome runtime. See
	//  https://support.google.com/chrome/a/answer/9037717 for details.
	//  Chrome Browser Cloud Management integration, when enabled via the
	//  "enable-chrome-browser-cloud-management" command-line flag, will also use
	//  the specified ID. See https://support.google.com/chrome/a/answer/9116814
	//  for details.
	ChromePolicyId() string         // property ChromePolicyId Getter
	SetChromePolicyId(value string) // property ChromePolicyId Setter
	// ChromeAppIconId
	//  Specify an ID for an ICON resource that can be loaded from the main
	//  executable and used when creating default Chrome windows such as DevTools
	//  and Task Manager. If unspecified the default Chromium ICON (IDR_MAINFRAME
	//  [101]) will be loaded from libcef.dll. Only supported with the Chrome
	//  runtime on Windows.
	ChromeAppIconId() int32         // property ChromeAppIconId Getter
	SetChromeAppIconId(value int32) // property ChromeAppIconId Setter
	// SingleProcess
	//  Runs the renderer and plugins in the same process as the browser.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --single-process</see>
	SingleProcess() bool         // property SingleProcess Getter
	SetSingleProcess(value bool) // property SingleProcess Setter
	// EnableMediaStream
	//  Enable media (WebRTC audio/video) streaming.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-media-stream</see>
	EnableMediaStream() bool         // property EnableMediaStream Getter
	SetEnableMediaStream(value bool) // property EnableMediaStream Setter
	// EnableSpeechInput
	//  Enable speech input (x-webkit-speech).
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-speech-input</see>
	EnableSpeechInput() bool         // property EnableSpeechInput Getter
	SetEnableSpeechInput(value bool) // property EnableSpeechInput Setter
	// UseFakeUIForMediaStream
	//  Bypass the media stream infobar by selecting the default device for media streams (e.g. WebRTC). Works with --use-fake-device-for-media-stream.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --single-process</see>
	UseFakeUIForMediaStream() bool         // property UseFakeUIForMediaStream Getter
	SetUseFakeUIForMediaStream(value bool) // property UseFakeUIForMediaStream Setter
	// EnableUsermediaScreenCapturing
	//  Enable screen capturing support for MediaStream API.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-usermedia-screen-capturing</see>
	EnableUsermediaScreenCapturing() bool         // property EnableUsermediaScreenCapturing Getter
	SetEnableUsermediaScreenCapturing(value bool) // property EnableUsermediaScreenCapturing Setter
	// EnableGPU
	//  Enable GPU hardware acceleration.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu-compositing</see>
	EnableGPU() bool         // property EnableGPU Getter
	SetEnableGPU(value bool) // property EnableGPU Setter
	// EnableFeatures
	//  List of feature names to enable.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-features</see>
	//  The list of features you can enable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	EnableFeatures() string         // property EnableFeatures Getter
	SetEnableFeatures(value string) // property EnableFeatures Setter
	// DisableFeatures
	//  List of feature names to disable.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-features</see>
	//  The list of features you can disable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	DisableFeatures() string         // property DisableFeatures Getter
	SetDisableFeatures(value string) // property DisableFeatures Setter
	// EnableBlinkFeatures
	//  Enable one or more Blink runtime-enabled features.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-blink-features</see>
	//  The list of Blink features you can enable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	EnableBlinkFeatures() string         // property EnableBlinkFeatures Getter
	SetEnableBlinkFeatures(value string) // property EnableBlinkFeatures Setter
	// DisableBlinkFeatures
	//  Disable one or more Blink runtime-enabled features.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-blink-features</see>
	//  The list of Blink features you can disable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	DisableBlinkFeatures() string         // property DisableBlinkFeatures Getter
	SetDisableBlinkFeatures(value string) // property DisableBlinkFeatures Setter
	// BlinkSettings
	//  Set blink settings. Format is <name>[=<value],<name>[=<value>],...
	//  The names are declared in Settings.json5. For boolean type, use "true", "false",
	//  or omit '=<value>' part to set to true. For enum type, use the int value of the
	//  enum value. Applied after other command line flags and prefs.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --blink-settings</see>
	//  The list of Blink settings you can disable is here:
	//  https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/settings.json5
	BlinkSettings() string         // property BlinkSettings Getter
	SetBlinkSettings(value string) // property BlinkSettings Setter
	// ForceFieldTrials
	//  This option can be used to force field trials when testing changes locally.
	//  The argument is a list of name and value pairs, separated by slashes.
	//  If a trial name is prefixed with an asterisk, that trial will start activated.
	//  For example, the following argument defines two trials, with the second one
	//  activated: "GoogleNow/Enable/*MaterialDesignNTP/Default/" This option can also
	//  be used by the browser process to send the list of trials to a non-browser
	//  process, using the same format. See FieldTrialList::CreateTrialsFromString()
	//  in field_trial.h for details.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrials</see>
	//  https://source.chromium.org/chromium/chromium/src/+/master:base/base_switches.cc
	ForceFieldTrials() string         // property ForceFieldTrials Getter
	SetForceFieldTrials(value string) // property ForceFieldTrials Setter
	// ForceFieldTrialParams
	//  This option can be used to force parameters of field trials when testing
	//  changes locally. The argument is a param list of (key, value) pairs prefixed
	//  by an associated (trial, group) pair. You specify the param list for multiple
	//  (trial, group) pairs with a comma separator.
	//  Example: "Trial1.Group1:k1/v1/k2/v2,Trial2.Group2:k3/v3/k4/v4"
	//  Trial names, groups names, parameter names, and value should all be URL
	//  escaped for all non-alphanumeric characters.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrial-params</see>
	//  https://source.chromium.org/chromium/chromium/src/+/master:components/variations/variations_switches.cc
	ForceFieldTrialParams() string         // property ForceFieldTrialParams Getter
	SetForceFieldTrialParams(value string) // property ForceFieldTrialParams Setter
	// SmoothScrolling
	//  On platforms that support it, enables smooth scroll animation.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-smooth-scrolling</see>
	SmoothScrolling() cefTypes.TCefState         // property SmoothScrolling Getter
	SetSmoothScrolling(value cefTypes.TCefState) // property SmoothScrolling Setter
	// MuteAudio
	//  Mutes audio sent to the audio device so it is not audible during automated testing.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --mute-audio</see>
	MuteAudio() bool         // property MuteAudio Getter
	SetMuteAudio(value bool) // property MuteAudio Setter
	// SitePerProcess
	//  Enforces a one-site-per-process security policy: Each renderer process, for its
	//  whole lifetime, is dedicated to rendering pages for just one site. Thus, pages
	//  from different sites are never in the same process. A renderer process's access
	//  rights are restricted based on its site.All cross-site navigations force process
	//  swaps. <iframe>s are rendered out-of-process whenever the src= is cross-site.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --site-per-process</see>
	//  More details here:
	//  https://www.chromium.org/developers/design-documents/site-isolation
	//  https://www.chromium.org/developers/design-documents/process-models
	SitePerProcess() bool         // property SitePerProcess Getter
	SetSitePerProcess(value bool) // property SitePerProcess Setter
	// DisableWebSecurity
	//  Don't enforce the same-origin policy.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-web-security</see>
	DisableWebSecurity() bool         // property DisableWebSecurity Getter
	SetDisableWebSecurity(value bool) // property DisableWebSecurity Setter
	// DisablePDFExtension
	//  Disable the PDF extension.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-pdf-extension</see>
	DisablePDFExtension() bool         // property DisablePDFExtension Getter
	SetDisablePDFExtension(value bool) // property DisablePDFExtension Setter
	// DisableSiteIsolationTrials
	//  Disables site isolation.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-site-isolation-trials</see>
	DisableSiteIsolationTrials() bool         // property DisableSiteIsolationTrials Getter
	SetDisableSiteIsolationTrials(value bool) // property DisableSiteIsolationTrials Setter
	// DisableChromeLoginPrompt
	//  Delegate all login requests to the client GetAuthCredentials
	//  callback when using the Chrome runtime.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-chrome-login-prompt</see>
	DisableChromeLoginPrompt() bool         // property DisableChromeLoginPrompt Getter
	SetDisableChromeLoginPrompt(value bool) // property DisableChromeLoginPrompt Setter
	// DisableExtensions
	//  Disable extensions.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-extensions</see>
	DisableExtensions() bool         // property DisableExtensions Getter
	SetDisableExtensions(value bool) // property DisableExtensions Setter
	// AutoplayPolicy
	//  Autoplay policy.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --autoplay-policy</see>
	AutoplayPolicy() cefTypes.TCefAutoplayPolicy         // property AutoplayPolicy Getter
	SetAutoplayPolicy(value cefTypes.TCefAutoplayPolicy) // property AutoplayPolicy Setter
	// DisableBackgroundNetworking
	//  Disable several subsystems which run network requests in the background.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-background-networking</see>
	DisableBackgroundNetworking() bool         // property DisableBackgroundNetworking Getter
	SetDisableBackgroundNetworking(value bool) // property DisableBackgroundNetworking Setter
	// MetricsRecordingOnly
	//  Enables the recording of metrics reports but disables reporting.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --metrics-recording-only</see>
	MetricsRecordingOnly() bool         // property MetricsRecordingOnly Getter
	SetMetricsRecordingOnly(value bool) // property MetricsRecordingOnly Setter
	// AllowFileAccessFromFiles
	//  By default, file:// URIs cannot read other file:// URIs. This is an override for developers who need the old behavior for testing.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-file-access-from-files</see>
	AllowFileAccessFromFiles() bool         // property AllowFileAccessFromFiles Getter
	SetAllowFileAccessFromFiles(value bool) // property AllowFileAccessFromFiles Setter
	// AllowRunningInsecureContent
	//  By default, an https page cannot run JavaScript, CSS or plugins from http URLs. This provides an override to get the old insecure behavior.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-running-insecure-content</see>
	AllowRunningInsecureContent() bool         // property AllowRunningInsecureContent Getter
	SetAllowRunningInsecureContent(value bool) // property AllowRunningInsecureContent Setter
	// EnablePrintPreview
	//  Enable print preview.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-print-preview</see>
	EnablePrintPreview() bool         // property EnablePrintPreview Getter
	SetEnablePrintPreview(value bool) // property EnablePrintPreview Setter
	// DefaultEncoding
	//  Default encoding.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --default-encoding</see>
	DefaultEncoding() string         // property DefaultEncoding Getter
	SetDefaultEncoding(value string) // property DefaultEncoding Setter
	// DisableJavascript
	//  Disable JavaScript.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript</see>
	DisableJavascript() bool         // property DisableJavascript Getter
	SetDisableJavascript(value bool) // property DisableJavascript Setter
	// DisableJavascriptCloseWindows
	//  Disable closing of windows via JavaScript.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-close-windows</see>
	DisableJavascriptCloseWindows() bool         // property DisableJavascriptCloseWindows Getter
	SetDisableJavascriptCloseWindows(value bool) // property DisableJavascriptCloseWindows Setter
	// DisableJavascriptAccessClipboard
	//  Disable clipboard access via JavaScript.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-access-clipboard</see>
	DisableJavascriptAccessClipboard() bool         // property DisableJavascriptAccessClipboard Getter
	SetDisableJavascriptAccessClipboard(value bool) // property DisableJavascriptAccessClipboard Setter
	// DisableJavascriptDomPaste
	//  Disable DOM paste via JavaScript execCommand("paste").
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-dom-paste</see>
	DisableJavascriptDomPaste() bool         // property DisableJavascriptDomPaste Getter
	SetDisableJavascriptDomPaste(value bool) // property DisableJavascriptDomPaste Setter
	// AllowUniversalAccessFromFileUrls
	//  Allow universal access from file URLs.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --allow-universal-access-from-files</see>
	AllowUniversalAccessFromFileUrls() bool         // property AllowUniversalAccessFromFileUrls Getter
	SetAllowUniversalAccessFromFileUrls(value bool) // property AllowUniversalAccessFromFileUrls Setter
	// DisableImageLoading
	//  Disable loading of images from the network. A cached image will still be rendered if requested.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-image-loading</see>
	DisableImageLoading() bool         // property DisableImageLoading Getter
	SetDisableImageLoading(value bool) // property DisableImageLoading Setter
	// ImageShrinkStandaloneToFit
	//  Shrink stand-alone images to fit.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --image-shrink-standalone-to-fit</see>
	ImageShrinkStandaloneToFit() bool         // property ImageShrinkStandaloneToFit Getter
	SetImageShrinkStandaloneToFit(value bool) // property ImageShrinkStandaloneToFit Setter
	// DisableTextAreaResize
	//  Disable resizing of text areas.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-text-area-resize</see>
	DisableTextAreaResize() bool         // property DisableTextAreaResize Getter
	SetDisableTextAreaResize(value bool) // property DisableTextAreaResize Setter
	// DisableTabToLinks
	//  Disable using the tab key to advance focus to links.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-tab-to-links</see>
	DisableTabToLinks() bool         // property DisableTabToLinks Getter
	SetDisableTabToLinks(value bool) // property DisableTabToLinks Setter
	// EnableProfanityFilter
	//  Enable the speech input profanity filter.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-profanity-filter</see>
	EnableProfanityFilter() bool         // property EnableProfanityFilter Getter
	SetEnableProfanityFilter(value bool) // property EnableProfanityFilter Setter
	// DisableSpellChecking
	//  Disable spell checking.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-spell-checking</see>
	DisableSpellChecking() bool         // property DisableSpellChecking Getter
	SetDisableSpellChecking(value bool) // property DisableSpellChecking Setter
	// OverrideSpellCheckLang
	//  Override the default spellchecking language which comes from locales.pak.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --override-spell-check-lang</see>
	OverrideSpellCheckLang() string         // property OverrideSpellCheckLang Getter
	SetOverrideSpellCheckLang(value string) // property OverrideSpellCheckLang Setter
	// TouchEvents
	//  Enable support for touch event feature detection.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --touch-events</see>
	TouchEvents() cefTypes.TCefState         // property TouchEvents Getter
	SetTouchEvents(value cefTypes.TCefState) // property TouchEvents Setter
	// DisableReadingFromCanvas
	//  Taints all <canvas> elements, regardless of origin.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-reading-from-canvas</see>
	DisableReadingFromCanvas() bool         // property DisableReadingFromCanvas Getter
	SetDisableReadingFromCanvas(value bool) // property DisableReadingFromCanvas Setter
	// HyperlinkAuditing
	//  Don't send hyperlink auditing pings.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-pings</see>
	HyperlinkAuditing() bool         // property HyperlinkAuditing Getter
	SetHyperlinkAuditing(value bool) // property HyperlinkAuditing Setter
	// DisableNewBrowserInfoTimeout
	//  Disable the timeout for delivering new browser info to the renderer process.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-new-browser-info-timeout</see>
	DisableNewBrowserInfoTimeout() bool         // property DisableNewBrowserInfoTimeout Getter
	SetDisableNewBrowserInfoTimeout(value bool) // property DisableNewBrowserInfoTimeout Setter
	// DevToolsProtocolLogFile
	//  File used for logging DevTools protocol messages.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --devtools-protocol-log-file</see>
	DevToolsProtocolLogFile() string         // property DevToolsProtocolLogFile Getter
	SetDevToolsProtocolLogFile(value string) // property DevToolsProtocolLogFile Setter
	// ForcedDeviceScaleFactor
	//  Overrides the device scale factor for the browser UI and the contents.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-device-scale-factor</see>
	ForcedDeviceScaleFactor() float32         // property ForcedDeviceScaleFactor Getter
	SetForcedDeviceScaleFactor(value float32) // property ForcedDeviceScaleFactor Setter
	// DisableZygote
	//  Disables the use of a zygote process for forking child processes. Instead, child processes will be forked and exec'd directly.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-zygote</see>
	DisableZygote() bool         // property DisableZygote Getter
	SetDisableZygote(value bool) // property DisableZygote Setter
	// UseMockKeyChain
	//  Uses mock keychain for testing purposes, which prevents blocking dialogs from causing timeouts.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --use-mock-keychain</see>
	UseMockKeyChain() bool         // property UseMockKeyChain Getter
	SetUseMockKeyChain(value bool) // property UseMockKeyChain Setter
	// DisableRequestHandlingForTesting
	//  Disable request handling in CEF to faciliate debugging of network-related issues.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-request-handling-for-testing</see>
	DisableRequestHandlingForTesting() bool         // property DisableRequestHandlingForTesting Getter
	SetDisableRequestHandlingForTesting(value bool) // property DisableRequestHandlingForTesting Setter
	// DisablePopupBlocking
	//  Disables pop-up blocking.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-popup-blocking</see>
	DisablePopupBlocking() bool         // property DisablePopupBlocking Getter
	SetDisablePopupBlocking(value bool) // property DisablePopupBlocking Setter
	// DisableBackForwardCache
	//  Disables the BackForwardCache feature.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-back-forward-cache</see>
	DisableBackForwardCache() bool         // property DisableBackForwardCache Getter
	SetDisableBackForwardCache(value bool) // property DisableBackForwardCache Setter
	// DisableComponentUpdate
	//  Disable the component updater. Widevine will not be downloaded or initialized.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-component-update</see>
	DisableComponentUpdate() bool         // property DisableComponentUpdate Getter
	SetDisableComponentUpdate(value bool) // property DisableComponentUpdate Setter
	// AllowInsecureLocalhost
	//  Enables TLS/SSL errors on localhost to be ignored (no interstitial, no blocking of requests).
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-insecure-localhost</see>
	AllowInsecureLocalhost() bool         // property AllowInsecureLocalhost Getter
	SetAllowInsecureLocalhost(value bool) // property AllowInsecureLocalhost Setter
	// KioskPrinting
	//  Enable automatically pressing the print button in print preview.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --kiosk-printing</see>
	KioskPrinting() bool         // property KioskPrinting Getter
	SetKioskPrinting(value bool) // property KioskPrinting Setter
	// TreatInsecureOriginAsSecure
	//  Treat given (insecure) origins as secure origins.
	//  Multiple origins can be supplied as a comma-separated list.
	//  For the definition of secure contexts, see https://w3c.github.io/webappsec-secure-contexts/
	//  and https://www.w3.org/TR/powerful-features/#is-origin-trustworthy
	//  Example: --unsafely-treat-insecure-origin-as-secure=http://a.test,http://b.test
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --unsafely-treat-insecure-origin-as-secure</see>
	TreatInsecureOriginAsSecure() string         // property TreatInsecureOriginAsSecure Getter
	SetTreatInsecureOriginAsSecure(value string) // property TreatInsecureOriginAsSecure Setter
	// NetLogEnabled
	//  Enables saving net log events to a file.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-net-log</see>
	NetLogEnabled() bool         // property NetLogEnabled Getter
	SetNetLogEnabled(value bool) // property NetLogEnabled Setter
	// NetLogFile
	//  File name used to log net events. If a value is given,
	//  it used as the path the the file, otherwise the file is named netlog.json
	//  and placed in the user data directory.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-net-log</see>
	NetLogFile() string         // property NetLogFile Getter
	SetNetLogFile(value string) // property NetLogFile Setter
	// NetLogCaptureMode
	//  Sets the granularity of events to capture in the network log.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --net-log-capture-mode</see>
	NetLogCaptureMode() cefTypes.TCefNetLogCaptureMode         // property NetLogCaptureMode Getter
	SetNetLogCaptureMode(value cefTypes.TCefNetLogCaptureMode) // property NetLogCaptureMode Setter
	// RemoteAllowOrigins
	//  Enables web socket connections from the specified origins only. '*' allows any origin.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-allow-origins</see>
	RemoteAllowOrigins() string         // property RemoteAllowOrigins Getter
	SetRemoteAllowOrigins(value string) // property RemoteAllowOrigins Setter
	// AutoAcceptCamAndMicCapture
	//  Bypasses the dialog prompting the user for permission to capture cameras and microphones.
	//  Useful in automatic tests of video-conferencing Web applications. This is nearly
	//  identical to kUseFakeUIForMediaStream, with the exception being that this flag does NOT
	//  affect screen-capture.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --auto-accept-camera-and-microphone-capture</see>
	AutoAcceptCamAndMicCapture() bool         // property AutoAcceptCamAndMicCapture Getter
	SetAutoAcceptCamAndMicCapture(value bool) // property AutoAcceptCamAndMicCapture Setter
	// UIColorMode
	//  Forces light or dark mode in UI for platforms that support it.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switches: --force-dark-mode --force-light-mode</see>
	UIColorMode() cefTypes.TCefUIColorMode         // property UIColorMode Getter
	SetUIColorMode(value cefTypes.TCefUIColorMode) // property UIColorMode Setter
	// DisableHangMonitor
	//  Suppresses hang monitor dialogs in renderer processes. This may allow slow unload handlers on a page to prevent the tab from closing, but the Task Manager can be used to terminate the offending process in this case.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-hang-monitor</see>
	DisableHangMonitor() bool         // property DisableHangMonitor Getter
	SetDisableHangMonitor(value bool) // property DisableHangMonitor Setter
	// HideCrashRestoreBubble
	//  Does not show the "Restore pages" popup bubble after incorrect shutdown.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --hide-crash-restore-bubble</see>
	HideCrashRestoreBubble() bool         // property HideCrashRestoreBubble Getter
	SetHideCrashRestoreBubble(value bool) // property HideCrashRestoreBubble Setter
	// IgnoreCertificateErrors
	//  Ignores certificate-related errors.
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:components/network_session_configurator/common/network_switch_list.h">Uses the following command line switch: --ignore-certificate-errors</see>
	IgnoreCertificateErrors() bool         // property IgnoreCertificateErrors Getter
	SetIgnoreCertificateErrors(value bool) // property IgnoreCertificateErrors Setter
	// WindowsSandboxInfo
	//  Pointer to the sandbox info. Currently unused in Delphi and Lazarus.
	WindowsSandboxInfo() uintptr         // property WindowsSandboxInfo Getter
	SetWindowsSandboxInfo(value uintptr) // property WindowsSandboxInfo Setter
	// ArgcCopy
	//  argc parameter copy used in Linux only.
	ArgcCopy() int32 // property argcCopy Getter
	// ArgvCopy
	//  argv parameter copy used in Linux only.
	ArgvCopy() types.PPAnsiChar // property argvCopy Getter
	// DeleteCache
	//  Used to delete all the cache files before CEF is initialized.
	DeleteCache() bool         // property DeleteCache Getter
	SetDeleteCache(value bool) // property DeleteCache Setter
	// DeleteCookies
	//  Used to delete all the cookies before CEF is initialized.
	DeleteCookies() bool         // property DeleteCookies Getter
	SetDeleteCookies(value bool) // property DeleteCookies Setter
	// CheckCEFFiles
	//  Checks if the CEF binaries are present and the DLL version.
	CheckCEFFiles() bool         // property CheckCEFFiles Getter
	SetCheckCEFFiles(value bool) // property CheckCEFFiles Setter
	// ShowMessageDlg
	//  Set to true when you need to use a showmessage dialog to show the error messages.
	ShowMessageDlg() bool         // property ShowMessageDlg Getter
	SetShowMessageDlg(value bool) // property ShowMessageDlg Setter
	// MissingBinariesException
	//  Raise an exception when the CEF binaries check fails.
	MissingBinariesException() bool         // property MissingBinariesException Getter
	SetMissingBinariesException(value bool) // property MissingBinariesException Setter
	// SetCurrentDir
	//  Used to set the current directory when the CEF libraries are loaded. This is required if the application is launched from a different application.
	SetCurrentDir() bool         // property SetCurrentDir Getter
	SetSetCurrentDir(value bool) // property SetCurrentDir Setter
	// GlobalContextInitialized
	//  Set to True when the global context is initialized and the application can start creating web browsers.
	GlobalContextInitialized() bool // property GlobalContextInitialized Getter
	// ChromeMajorVer
	//  Returns the major version information from Chromium.
	ChromeMajorVer() uint16 // property ChromeMajorVer Getter
	// ChromeMinorVer
	//  Returns the minor version information from Chromium.
	ChromeMinorVer() uint16 // property ChromeMinorVer Getter
	// ChromeRelease
	//  Returns the release version information from Chromium.
	ChromeRelease() uint16 // property ChromeRelease Getter
	// ChromeBuild
	//  Returns the build version information from Chromium.
	ChromeBuild() uint16 // property ChromeBuild Getter
	// ChromeVersion
	//  Returns the full version information from Chromium.
	ChromeVersion() string // property ChromeVersion Getter
	// LibCefVersion
	//  Complete libcef version information.
	LibCefVersion() string // property LibCefVersion Getter
	// LibCefPath
	//  Path to libcef.dll or libcef.so
	LibCefPath() string // property LibCefPath Getter
	// ChromeElfPath
	//  Returns the path to chrome_elf.dll.
	ChromeElfPath() string // property ChromeElfPath Getter
	// LibLoaded
	//  Set to true when TCEFApplicationCore has loaded the CEF libraries.
	LibLoaded() bool // property LibLoaded Getter
	// LogProcessInfo
	//  Add a debug log information line when the CEF libraries are loaded.
	LogProcessInfo() bool         // property LogProcessInfo Getter
	SetLogProcessInfo(value bool) // property LogProcessInfo Setter
	// ReRaiseExceptions
	//  Set to true to raise all exceptions.
	ReRaiseExceptions() bool         // property ReRaiseExceptions Getter
	SetReRaiseExceptions(value bool) // property ReRaiseExceptions Setter
	// DeviceScaleFactor
	//  Returns the device scale factor used in OSR mode.
	DeviceScaleFactor() float32 // property DeviceScaleFactor Getter
	// LocalesRequired
	//  List of locale files that will be checked with CheckCEFFiles.
	LocalesRequired() string         // property LocalesRequired Getter
	SetLocalesRequired(value string) // property LocalesRequired Setter
	// ProcessType
	//  CEF process type currently running.
	ProcessType() cefTypes.TCefProcessType // property ProcessType Getter
	// MustCreateResourceBundleHandler
	//  Force the creation of ICefResourceBundleHandler.
	MustCreateResourceBundleHandler() bool         // property MustCreateResourceBundleHandler Getter
	SetMustCreateResourceBundleHandler(value bool) // property MustCreateResourceBundleHandler Setter
	// MustCreateBrowserProcessHandler
	//  Force the creation of ICefBrowserProcessHandler.
	MustCreateBrowserProcessHandler() bool         // property MustCreateBrowserProcessHandler Getter
	SetMustCreateBrowserProcessHandler(value bool) // property MustCreateBrowserProcessHandler Setter
	// MustCreateRenderProcessHandler
	//  Force the creation of ICefRenderProcessHandler.
	MustCreateRenderProcessHandler() bool         // property MustCreateRenderProcessHandler Getter
	SetMustCreateRenderProcessHandler(value bool) // property MustCreateRenderProcessHandler Setter
	// MustCreateLoadHandler
	//  Force the creation of ICefLoadHandler.
	MustCreateLoadHandler() bool         // property MustCreateLoadHandler Getter
	SetMustCreateLoadHandler(value bool) // property MustCreateLoadHandler Setter
	// OsmodalLoop
	//  Set to true (1) before calling Windows APIs like TrackPopupMenu that enter a
	//  modal message loop. Set to false (0) after exiting the modal message loop.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_app_win.h">CEF source file: /include/internal/cef_app_win.h (cef_set_osmodal_loop)</see>
	SetOsmodalLoop(value bool) // property OsmodalLoop Setter
	// Status
	//  Returns the TCEFApplicationCore initialization status.
	Status() cefTypes.TCefAplicationStatus // property Status Getter
	// MissingLibFiles
	//  List of missing CEF library files.
	MissingLibFiles() string // property MissingLibFiles Getter
	// MustFreeLibrary
	//  Set to true to free the library handle when TCEFApplicationCore is destroyed.
	MustFreeLibrary() bool         // property MustFreeLibrary Getter
	SetMustFreeLibrary(value bool) // property MustFreeLibrary Setter
	// ChildProcessesCount
	//  Returns the number of CEF subprocesses running at that moment.
	ChildProcessesCount() int32 // property ChildProcessesCount Getter
	// UsedMemory
	//  Total used memory by all CEF processes.
	UsedMemory() uint64 // property UsedMemory Getter
	// TotalSystemMemory
	//  Total system memory in Windows.
	TotalSystemMemory() uint64 // property TotalSystemMemory Getter
	// AvailableSystemMemory
	//  Calculates the available memory in Windows.
	AvailableSystemMemory() uint64 // property AvailableSystemMemory Getter
	// SystemMemoryLoad
	//  Memory load in Windows.
	SystemMemoryLoad() uint32 // property SystemMemoryLoad Getter
	// ApiHashUniversal
	//  Calls cef_api_hash to get the universal hash.
	ApiHashUniversal() string // property ApiHashUniversal Getter
	// ApiHashPlatform
	//  Calls cef_api_hash to get the platform hash.
	ApiHashPlatform() string // property ApiHashPlatform Getter
	// ApiHashCommit
	//  Calls cef_api_hash to get the commit hash.
	ApiHashCommit() string // property ApiHashCommit Getter
	// ExitCode
	//  This property can optionally be read on the main application thread after
	//  CefInitialize to retrieve the initialization exit code. When CefInitialize
	//  returns true (1) the exit code will be 0 (CEF_RESULT_CODE_NORMAL_EXIT).
	//  Otherwise, see TCefResultCode for possible exit code values including
	//  browser process initialization errors and normal early exit conditions (such
	//  as CEF_RESULT_CODE_NORMAL_EXIT_PROCESS_NOTIFIED for process singleton
	//  relaunch behavior).
	ExitCode() cefTypes.TCefResultCode // property ExitCode Getter
	// LastErrorMessage
	//  Last error message that is usually shown when CEF finds a problem at initialization.
	LastErrorMessage() string // property LastErrorMessage Getter
	// XDisplay
	//  Return the singleton X11 display shared with Chromium. The display is not
	//  thread-safe and must only be accessed on the browser process UI thread.
	XDisplay() uintptr                                                               // property XDisplay Getter
	SetOnRegCustomSchemes(fn TOnRegisterCustomSchemesEvent)                          // property event
	SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferencesEvent)             // property event
	SetOnContextInitialized(fn TOnContextInitializedEvent)                           // property event
	SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunchEvent)               // property event
	SetOnAlreadyRunningAppRelaunch(fn TOnAlreadyRunningAppRelaunchEvent)             // property event
	SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWorkEvent)                 // property event
	SetOnGetDefaultClient(fn TOnGetDefaultClientEvent)                               // property event
	SetOnGetDefaultRequestContextHandler(fn TOnGetDefaultRequestContextHandlerEvent) // property event
	SetOnGetLocalizedString(fn TOnGetLocalizedStringEvent)                           // property event
	SetOnGetDataResource(fn TOnGetDataResourceEvent)                                 // property event
	SetOnGetDataResourceForScale(fn TOnGetDataResourceForScaleEvent)                 // property event
	SetOnWebKitInitialized(fn TOnWebKitInitializedEvent)                             // property event
	SetOnBrowserCreated(fn TOnBrowserCreatedEvent)                                   // property event
	SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent)                               // property event
	SetOnContextCreated(fn TOnContextCreatedEvent)                                   // property event
	SetOnContextReleased(fn TOnContextReleasedEvent)                                 // property event
	SetOnUncaughtException(fn TOnUncaughtExceptionEvent)                             // property event
	SetOnFocusedNodeChanged(fn TOnFocusedNodeChangedEvent)                           // property event
	SetOnProcessMessageReceived(fn TOnProcessMessageReceivedEvent)                   // property event
	SetOnLoadingStateChange(fn TOnRenderLoadingStateChange)                          // property event
	SetOnLoadStart(fn TOnRenderLoadStart)                                            // property event
	SetOnLoadEnd(fn TOnRenderLoadEnd)                                                // property event
	SetOnLoadError(fn TOnRenderLoadError)                                            // property event
	AsIntfApplicationCoreEvents() uintptr
}

type TCefApplicationCore struct {
	TInterfacedObject
}

func (m *TCefApplicationCore) CheckCEFLibrary() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) StartMainProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) StartSubProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) ValidComponentID(componentID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(4, m.Instance(), uintptr(componentID))
	return api.GoBool(r)
}

func (m *TCefApplicationCore) NextComponentID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) AddCustomCommandLine(commandLine string, value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(6, m.Instance(), api.PasStr(commandLine), api.PasStr(value))
}

func (m *TCefApplicationCore) DoMessageLoopWork() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(7, m.Instance())
}

func (m *TCefApplicationCore) RunMessageLoop() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(8, m.Instance())
}

func (m *TCefApplicationCore) QuitMessageLoop() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(9, m.Instance())
}

func (m *TCefApplicationCore) UpdateDeviceScaleFactor() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(10, m.Instance())
}

func (m *TCefApplicationCore) InitLibLocationFromArgs() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(11, m.Instance())
}

func (m *TCefApplicationCore) RemoveComponentID(componentID int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(12, m.Instance(), uintptr(componentID))
}

func (m *TCefApplicationCore) NoSandbox() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetNoSandbox(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) BrowserSubprocessPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(14, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetBrowserSubprocessPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(14, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) FrameworkDirPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(15, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetFrameworkDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) MainBundlePath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(16, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetMainBundlePath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ChromeRuntime() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetChromeRuntime(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MultiThreadedMessageLoop() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMultiThreadedMessageLoop(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ExternalMessagePump() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetExternalMessagePump(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) WindowlessRenderingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetWindowlessRenderingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) CommandLineArgsDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCommandLineArgsDisabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) Cache() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(22, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetCache(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(22, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) RootCache() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetRootCache(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) PersistSessionCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPersistSessionCookies(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) PersistUserPreferences() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPersistUserPreferences(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UserAgent() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetUserAgent(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) UserAgentProduct() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(27, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetUserAgentProduct(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(27, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) Locale() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetLocale(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LogFile() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(29, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(29, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LogSeverity() cefTypes.TCefLogSeverity {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(30, 0, m.Instance())
	return cefTypes.TCefLogSeverity(r)
}

func (m *TCefApplicationCore) SetLogSeverity(value cefTypes.TCefLogSeverity) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) LogItems() cefTypes.TCefLogItems {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(31, 0, m.Instance())
	return cefTypes.TCefLogItems(r)
}

func (m *TCefApplicationCore) SetLogItems(value cefTypes.TCefLogItems) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) JavaScriptFlags() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(32, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetJavaScriptFlags(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(32, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ResourcesDirPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(33, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetResourcesDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(33, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LocalesDirPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(34, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetLocalesDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(34, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) PackLoadingDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(35, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPackLoadingDisabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(35, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) RemoteDebuggingPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(36, 0, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) SetRemoteDebuggingPort(value int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) UncaughtExceptionStackSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(37, 0, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) SetUncaughtExceptionStackSize(value int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) BackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(38, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefApplicationCore) SetBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) AcceptLanguageList() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(39, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetAcceptLanguageList(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(39, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) CookieableSchemesList() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(40, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetCookieableSchemesList(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(40, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) CookieableSchemesExcludeDefaults() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(41, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCookieableSchemesExcludeDefaults(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(41, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ChromePolicyId() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(42, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetChromePolicyId(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(42, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ChromeAppIconId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(43, 0, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) SetChromeAppIconId(value int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) SingleProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(44, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSingleProcess(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(44, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableMediaStream() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(45, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableMediaStream(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(45, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableSpeechInput() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(46, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableSpeechInput(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(46, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UseFakeUIForMediaStream() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(47, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetUseFakeUIForMediaStream(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(47, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableUsermediaScreenCapturing() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(48, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableUsermediaScreenCapturing(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(48, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableGPU() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(49, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableGPU(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(49, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(50, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetEnableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(50, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(51, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetDisableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(51, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) EnableBlinkFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(52, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetEnableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(52, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableBlinkFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(53, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetDisableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(53, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) BlinkSettings() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(54, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetBlinkSettings(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(54, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForceFieldTrials() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(55, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetForceFieldTrials(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(55, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForceFieldTrialParams() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(56, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetForceFieldTrialParams(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(56, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) SmoothScrolling() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(57, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TCefApplicationCore) SetSmoothScrolling(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) MuteAudio() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(58, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMuteAudio(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(58, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SitePerProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(59, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSitePerProcess(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(59, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableWebSecurity() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(60, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableWebSecurity(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(60, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisablePDFExtension() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(61, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisablePDFExtension(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(61, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableSiteIsolationTrials() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableSiteIsolationTrials(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableChromeLoginPrompt() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(63, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableChromeLoginPrompt(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(63, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableExtensions() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableExtensions(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AutoplayPolicy() cefTypes.TCefAutoplayPolicy {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(65, 0, m.Instance())
	return cefTypes.TCefAutoplayPolicy(r)
}

func (m *TCefApplicationCore) SetAutoplayPolicy(value cefTypes.TCefAutoplayPolicy) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(65, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) DisableBackgroundNetworking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableBackgroundNetworking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MetricsRecordingOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(67, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMetricsRecordingOnly(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(67, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowFileAccessFromFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(68, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowFileAccessFromFiles(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(68, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowRunningInsecureContent() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(69, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowRunningInsecureContent(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(69, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnablePrintPreview() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(70, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnablePrintPreview(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(70, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DefaultEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(71, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(71, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableJavascript() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(72, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascript(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(72, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptCloseWindows() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(73, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptCloseWindows(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(73, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptAccessClipboard() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(74, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptAccessClipboard(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(74, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptDomPaste() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(75, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptDomPaste(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(75, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowUniversalAccessFromFileUrls() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(76, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowUniversalAccessFromFileUrls(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(76, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableImageLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(77, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableImageLoading(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(77, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ImageShrinkStandaloneToFit() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(78, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetImageShrinkStandaloneToFit(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(78, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableTextAreaResize() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(79, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableTextAreaResize(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(79, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableTabToLinks() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(80, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableTabToLinks(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(80, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableProfanityFilter() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(81, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableProfanityFilter(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(81, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableSpellChecking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(82, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableSpellChecking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(82, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) OverrideSpellCheckLang() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(83, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetOverrideSpellCheckLang(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(83, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) TouchEvents() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(84, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TCefApplicationCore) SetTouchEvents(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(84, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) DisableReadingFromCanvas() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(85, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableReadingFromCanvas(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(85, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) HyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(86, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetHyperlinkAuditing(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(86, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableNewBrowserInfoTimeout() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(87, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableNewBrowserInfoTimeout(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(87, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DevToolsProtocolLogFile() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(88, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetDevToolsProtocolLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(88, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForcedDeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(89, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) SetForcedDeviceScaleFactor(value float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(89, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCefApplicationCore) DisableZygote() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(90, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableZygote(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(90, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UseMockKeyChain() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(91, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetUseMockKeyChain(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(91, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableRequestHandlingForTesting() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(92, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableRequestHandlingForTesting(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(92, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisablePopupBlocking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(93, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisablePopupBlocking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(93, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableBackForwardCache() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(94, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableBackForwardCache(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(94, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableComponentUpdate() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(95, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableComponentUpdate(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(95, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowInsecureLocalhost() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(96, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowInsecureLocalhost(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(96, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) KioskPrinting() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(97, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetKioskPrinting(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(97, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) TreatInsecureOriginAsSecure() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(98, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetTreatInsecureOriginAsSecure(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(98, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) NetLogEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(99, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetNetLogEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(99, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) NetLogFile() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(100, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetNetLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(100, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) NetLogCaptureMode() cefTypes.TCefNetLogCaptureMode {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(101, 0, m.Instance())
	return cefTypes.TCefNetLogCaptureMode(r)
}

func (m *TCefApplicationCore) SetNetLogCaptureMode(value cefTypes.TCefNetLogCaptureMode) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(101, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) RemoteAllowOrigins() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(102, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetRemoteAllowOrigins(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(102, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) AutoAcceptCamAndMicCapture() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(103, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAutoAcceptCamAndMicCapture(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(103, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UIColorMode() cefTypes.TCefUIColorMode {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(104, 0, m.Instance())
	return cefTypes.TCefUIColorMode(r)
}

func (m *TCefApplicationCore) SetUIColorMode(value cefTypes.TCefUIColorMode) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(104, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) DisableHangMonitor() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(105, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableHangMonitor(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(105, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) HideCrashRestoreBubble() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(106, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetHideCrashRestoreBubble(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(106, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) IgnoreCertificateErrors() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(107, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetIgnoreCertificateErrors(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(107, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) WindowsSandboxInfo() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(108, 0, m.Instance())
	return uintptr(r)
}

func (m *TCefApplicationCore) SetWindowsSandboxInfo(value uintptr) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(108, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) ArgcCopy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(109, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) ArgvCopy() types.PPAnsiChar {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(110, m.Instance())
	return types.PPAnsiChar(r)
}

func (m *TCefApplicationCore) DeleteCache() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(111, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDeleteCache(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(111, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DeleteCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(112, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDeleteCookies(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(112, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) CheckCEFFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(113, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCheckCEFFiles(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(113, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ShowMessageDlg() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(114, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetShowMessageDlg(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(114, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MissingBinariesException() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(115, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMissingBinariesException(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(115, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SetCurrentDir() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(116, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSetCurrentDir(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(116, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) GlobalContextInitialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(117, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) ChromeMajorVer() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(118, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeMinorVer() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(119, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeRelease() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(120, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeBuild() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(121, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(122, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) LibCefVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(123, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) LibCefPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(124, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) ChromeElfPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(125, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) LibLoaded() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(126, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) LogProcessInfo() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(127, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetLogProcessInfo(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(127, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ReRaiseExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(128, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetReRaiseExceptions(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(128, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(129, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) LocalesRequired() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(130, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) SetLocalesRequired(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(130, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ProcessType() cefTypes.TCefProcessType {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(131, m.Instance())
	return cefTypes.TCefProcessType(r)
}

func (m *TCefApplicationCore) MustCreateResourceBundleHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(132, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateResourceBundleHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(132, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateBrowserProcessHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(133, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateBrowserProcessHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(133, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateRenderProcessHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(134, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateRenderProcessHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(134, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateLoadHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(135, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateLoadHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(135, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SetOsmodalLoop(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(136, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) Status() cefTypes.TCefAplicationStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(137, m.Instance())
	return cefTypes.TCefAplicationStatus(r)
}

func (m *TCefApplicationCore) MissingLibFiles() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(138, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) MustFreeLibrary() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(139, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustFreeLibrary(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(139, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ChildProcessesCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(140, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) UsedMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(141, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) TotalSystemMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(142, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) AvailableSystemMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(143, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) SystemMemoryLoad() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(144, m.Instance())
	return uint32(r)
}

func (m *TCefApplicationCore) ApiHashUniversal() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(145, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) ApiHashPlatform() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(146, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) ApiHashCommit() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(147, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) ExitCode() cefTypes.TCefResultCode {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(148, m.Instance())
	return cefTypes.TCefResultCode(r)
}

func (m *TCefApplicationCore) LastErrorMessage() string {
	if !m.IsValid() {
		return ""
	}
	r := cefApplicationCoreAPI().SysCallN(149, m.Instance())
	return api.GoStr(r)
}

func (m *TCefApplicationCore) XDisplay() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(150, m.Instance())
	return uintptr(r)
}

func (m *TCefApplicationCore) SetOnRegCustomSchemes(fn TOnRegisterCustomSchemesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRegisterCustomSchemesEvent(fn)
	base.SetEvent(m, 151, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferencesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRegisterCustomPreferencesEvent(fn)
	base.SetEvent(m, 152, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextInitialized(fn TOnContextInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextInitializedEvent(fn)
	base.SetEvent(m, 153, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeChildProcessLaunchEvent(fn)
	base.SetEvent(m, 154, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnAlreadyRunningAppRelaunch(fn TOnAlreadyRunningAppRelaunchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAlreadyRunningAppRelaunchEvent(fn)
	base.SetEvent(m, 155, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWorkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScheduleMessagePumpWorkEvent(fn)
	base.SetEvent(m, 156, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDefaultClient(fn TOnGetDefaultClientEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDefaultClientEvent(fn)
	base.SetEvent(m, 157, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDefaultRequestContextHandler(fn TOnGetDefaultRequestContextHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDefaultRequestContextHandlerEvent(fn)
	base.SetEvent(m, 158, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetLocalizedString(fn TOnGetLocalizedStringEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetLocalizedStringEvent(fn)
	base.SetEvent(m, 159, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDataResource(fn TOnGetDataResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDataResourceEvent(fn)
	base.SetEvent(m, 160, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDataResourceForScale(fn TOnGetDataResourceForScaleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDataResourceForScaleEvent(fn)
	base.SetEvent(m, 161, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnWebKitInitialized(fn TOnWebKitInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebKitInitializedEvent(fn)
	base.SetEvent(m, 162, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBrowserCreated(fn TOnBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserCreatedEvent(fn)
	base.SetEvent(m, 163, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserDestroyedEvent(fn)
	base.SetEvent(m, 164, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextCreated(fn TOnContextCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextCreatedEvent(fn)
	base.SetEvent(m, 165, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextReleased(fn TOnContextReleasedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextReleasedEvent(fn)
	base.SetEvent(m, 166, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnUncaughtException(fn TOnUncaughtExceptionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUncaughtExceptionEvent(fn)
	base.SetEvent(m, 167, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnFocusedNodeChanged(fn TOnFocusedNodeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusedNodeChangedEvent(fn)
	base.SetEvent(m, 168, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessMessageReceivedEvent(fn)
	base.SetEvent(m, 169, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadingStateChange(fn TOnRenderLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadingStateChange(fn)
	base.SetEvent(m, 170, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadStart(fn TOnRenderLoadStart) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadStart(fn)
	base.SetEvent(m, 171, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadEnd(fn TOnRenderLoadEnd) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadEnd(fn)
	base.SetEvent(m, 172, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadError(fn TOnRenderLoadError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadError(fn)
	base.SetEvent(m, 173, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) AsIntfApplicationCoreEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewApplicationCore class constructor
func NewApplicationCore() ICefApplicationCore {
	var applicationCoreEventsPtr uintptr // IApplicationCoreEvents
	r := cefApplicationCoreAPI().SysCallN(0, uintptr(base.UnsafePointer(&applicationCoreEventsPtr)))
	ret := AsCefApplicationCore(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, applicationCoreEventsPtr)
	}
	return ret
}

var (
	cefApplicationCoreOnce   base.Once
	cefApplicationCoreImport *imports.Imports = nil
)

func cefApplicationCoreAPI() *imports.Imports {
	cefApplicationCoreOnce.Do(func() {
		cefApplicationCoreImport = api.NewDefaultImports()
		cefApplicationCoreImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefApplicationCore_Create", 0), // constructor NewApplicationCore
			/* 1 */ imports.NewTable("TCefApplicationCore_CheckCEFLibrary", 0), // function CheckCEFLibrary
			/* 2 */ imports.NewTable("TCefApplicationCore_StartMainProcess", 0), // function StartMainProcess
			/* 3 */ imports.NewTable("TCefApplicationCore_StartSubProcess", 0), // function StartSubProcess
			/* 4 */ imports.NewTable("TCefApplicationCore_ValidComponentID", 0), // function ValidComponentID
			/* 5 */ imports.NewTable("TCefApplicationCore_NextComponentID", 0), // function NextComponentID
			/* 6 */ imports.NewTable("TCefApplicationCore_AddCustomCommandLine", 0), // procedure AddCustomCommandLine
			/* 7 */ imports.NewTable("TCefApplicationCore_DoMessageLoopWork", 0), // procedure DoMessageLoopWork
			/* 8 */ imports.NewTable("TCefApplicationCore_RunMessageLoop", 0), // procedure RunMessageLoop
			/* 9 */ imports.NewTable("TCefApplicationCore_QuitMessageLoop", 0), // procedure QuitMessageLoop
			/* 10 */ imports.NewTable("TCefApplicationCore_UpdateDeviceScaleFactor", 0), // procedure UpdateDeviceScaleFactor
			/* 11 */ imports.NewTable("TCefApplicationCore_InitLibLocationFromArgs", 0), // procedure InitLibLocationFromArgs
			/* 12 */ imports.NewTable("TCefApplicationCore_RemoveComponentID", 0), // procedure RemoveComponentID
			/* 13 */ imports.NewTable("TCefApplicationCore_NoSandbox", 0), // property NoSandbox
			/* 14 */ imports.NewTable("TCefApplicationCore_BrowserSubprocessPath", 0), // property BrowserSubprocessPath
			/* 15 */ imports.NewTable("TCefApplicationCore_FrameworkDirPath", 0), // property FrameworkDirPath
			/* 16 */ imports.NewTable("TCefApplicationCore_MainBundlePath", 0), // property MainBundlePath
			/* 17 */ imports.NewTable("TCefApplicationCore_ChromeRuntime", 0), // property ChromeRuntime
			/* 18 */ imports.NewTable("TCefApplicationCore_MultiThreadedMessageLoop", 0), // property MultiThreadedMessageLoop
			/* 19 */ imports.NewTable("TCefApplicationCore_ExternalMessagePump", 0), // property ExternalMessagePump
			/* 20 */ imports.NewTable("TCefApplicationCore_WindowlessRenderingEnabled", 0), // property WindowlessRenderingEnabled
			/* 21 */ imports.NewTable("TCefApplicationCore_CommandLineArgsDisabled", 0), // property CommandLineArgsDisabled
			/* 22 */ imports.NewTable("TCefApplicationCore_Cache", 0), // property Cache
			/* 23 */ imports.NewTable("TCefApplicationCore_RootCache", 0), // property RootCache
			/* 24 */ imports.NewTable("TCefApplicationCore_PersistSessionCookies", 0), // property PersistSessionCookies
			/* 25 */ imports.NewTable("TCefApplicationCore_PersistUserPreferences", 0), // property PersistUserPreferences
			/* 26 */ imports.NewTable("TCefApplicationCore_UserAgent", 0), // property UserAgent
			/* 27 */ imports.NewTable("TCefApplicationCore_UserAgentProduct", 0), // property UserAgentProduct
			/* 28 */ imports.NewTable("TCefApplicationCore_Locale", 0), // property Locale
			/* 29 */ imports.NewTable("TCefApplicationCore_LogFile", 0), // property LogFile
			/* 30 */ imports.NewTable("TCefApplicationCore_LogSeverity", 0), // property LogSeverity
			/* 31 */ imports.NewTable("TCefApplicationCore_LogItems", 0), // property LogItems
			/* 32 */ imports.NewTable("TCefApplicationCore_JavaScriptFlags", 0), // property JavaScriptFlags
			/* 33 */ imports.NewTable("TCefApplicationCore_ResourcesDirPath", 0), // property ResourcesDirPath
			/* 34 */ imports.NewTable("TCefApplicationCore_LocalesDirPath", 0), // property LocalesDirPath
			/* 35 */ imports.NewTable("TCefApplicationCore_PackLoadingDisabled", 0), // property PackLoadingDisabled
			/* 36 */ imports.NewTable("TCefApplicationCore_RemoteDebuggingPort", 0), // property RemoteDebuggingPort
			/* 37 */ imports.NewTable("TCefApplicationCore_UncaughtExceptionStackSize", 0), // property UncaughtExceptionStackSize
			/* 38 */ imports.NewTable("TCefApplicationCore_BackgroundColor", 0), // property BackgroundColor
			/* 39 */ imports.NewTable("TCefApplicationCore_AcceptLanguageList", 0), // property AcceptLanguageList
			/* 40 */ imports.NewTable("TCefApplicationCore_CookieableSchemesList", 0), // property CookieableSchemesList
			/* 41 */ imports.NewTable("TCefApplicationCore_CookieableSchemesExcludeDefaults", 0), // property CookieableSchemesExcludeDefaults
			/* 42 */ imports.NewTable("TCefApplicationCore_ChromePolicyId", 0), // property ChromePolicyId
			/* 43 */ imports.NewTable("TCefApplicationCore_ChromeAppIconId", 0), // property ChromeAppIconId
			/* 44 */ imports.NewTable("TCefApplicationCore_SingleProcess", 0), // property SingleProcess
			/* 45 */ imports.NewTable("TCefApplicationCore_EnableMediaStream", 0), // property EnableMediaStream
			/* 46 */ imports.NewTable("TCefApplicationCore_EnableSpeechInput", 0), // property EnableSpeechInput
			/* 47 */ imports.NewTable("TCefApplicationCore_UseFakeUIForMediaStream", 0), // property UseFakeUIForMediaStream
			/* 48 */ imports.NewTable("TCefApplicationCore_EnableUsermediaScreenCapturing", 0), // property EnableUsermediaScreenCapturing
			/* 49 */ imports.NewTable("TCefApplicationCore_EnableGPU", 0), // property EnableGPU
			/* 50 */ imports.NewTable("TCefApplicationCore_EnableFeatures", 0), // property EnableFeatures
			/* 51 */ imports.NewTable("TCefApplicationCore_DisableFeatures", 0), // property DisableFeatures
			/* 52 */ imports.NewTable("TCefApplicationCore_EnableBlinkFeatures", 0), // property EnableBlinkFeatures
			/* 53 */ imports.NewTable("TCefApplicationCore_DisableBlinkFeatures", 0), // property DisableBlinkFeatures
			/* 54 */ imports.NewTable("TCefApplicationCore_BlinkSettings", 0), // property BlinkSettings
			/* 55 */ imports.NewTable("TCefApplicationCore_ForceFieldTrials", 0), // property ForceFieldTrials
			/* 56 */ imports.NewTable("TCefApplicationCore_ForceFieldTrialParams", 0), // property ForceFieldTrialParams
			/* 57 */ imports.NewTable("TCefApplicationCore_SmoothScrolling", 0), // property SmoothScrolling
			/* 58 */ imports.NewTable("TCefApplicationCore_MuteAudio", 0), // property MuteAudio
			/* 59 */ imports.NewTable("TCefApplicationCore_SitePerProcess", 0), // property SitePerProcess
			/* 60 */ imports.NewTable("TCefApplicationCore_DisableWebSecurity", 0), // property DisableWebSecurity
			/* 61 */ imports.NewTable("TCefApplicationCore_DisablePDFExtension", 0), // property DisablePDFExtension
			/* 62 */ imports.NewTable("TCefApplicationCore_DisableSiteIsolationTrials", 0), // property DisableSiteIsolationTrials
			/* 63 */ imports.NewTable("TCefApplicationCore_DisableChromeLoginPrompt", 0), // property DisableChromeLoginPrompt
			/* 64 */ imports.NewTable("TCefApplicationCore_DisableExtensions", 0), // property DisableExtensions
			/* 65 */ imports.NewTable("TCefApplicationCore_AutoplayPolicy", 0), // property AutoplayPolicy
			/* 66 */ imports.NewTable("TCefApplicationCore_DisableBackgroundNetworking", 0), // property DisableBackgroundNetworking
			/* 67 */ imports.NewTable("TCefApplicationCore_MetricsRecordingOnly", 0), // property MetricsRecordingOnly
			/* 68 */ imports.NewTable("TCefApplicationCore_AllowFileAccessFromFiles", 0), // property AllowFileAccessFromFiles
			/* 69 */ imports.NewTable("TCefApplicationCore_AllowRunningInsecureContent", 0), // property AllowRunningInsecureContent
			/* 70 */ imports.NewTable("TCefApplicationCore_EnablePrintPreview", 0), // property EnablePrintPreview
			/* 71 */ imports.NewTable("TCefApplicationCore_DefaultEncoding", 0), // property DefaultEncoding
			/* 72 */ imports.NewTable("TCefApplicationCore_DisableJavascript", 0), // property DisableJavascript
			/* 73 */ imports.NewTable("TCefApplicationCore_DisableJavascriptCloseWindows", 0), // property DisableJavascriptCloseWindows
			/* 74 */ imports.NewTable("TCefApplicationCore_DisableJavascriptAccessClipboard", 0), // property DisableJavascriptAccessClipboard
			/* 75 */ imports.NewTable("TCefApplicationCore_DisableJavascriptDomPaste", 0), // property DisableJavascriptDomPaste
			/* 76 */ imports.NewTable("TCefApplicationCore_AllowUniversalAccessFromFileUrls", 0), // property AllowUniversalAccessFromFileUrls
			/* 77 */ imports.NewTable("TCefApplicationCore_DisableImageLoading", 0), // property DisableImageLoading
			/* 78 */ imports.NewTable("TCefApplicationCore_ImageShrinkStandaloneToFit", 0), // property ImageShrinkStandaloneToFit
			/* 79 */ imports.NewTable("TCefApplicationCore_DisableTextAreaResize", 0), // property DisableTextAreaResize
			/* 80 */ imports.NewTable("TCefApplicationCore_DisableTabToLinks", 0), // property DisableTabToLinks
			/* 81 */ imports.NewTable("TCefApplicationCore_EnableProfanityFilter", 0), // property EnableProfanityFilter
			/* 82 */ imports.NewTable("TCefApplicationCore_DisableSpellChecking", 0), // property DisableSpellChecking
			/* 83 */ imports.NewTable("TCefApplicationCore_OverrideSpellCheckLang", 0), // property OverrideSpellCheckLang
			/* 84 */ imports.NewTable("TCefApplicationCore_TouchEvents", 0), // property TouchEvents
			/* 85 */ imports.NewTable("TCefApplicationCore_DisableReadingFromCanvas", 0), // property DisableReadingFromCanvas
			/* 86 */ imports.NewTable("TCefApplicationCore_HyperlinkAuditing", 0), // property HyperlinkAuditing
			/* 87 */ imports.NewTable("TCefApplicationCore_DisableNewBrowserInfoTimeout", 0), // property DisableNewBrowserInfoTimeout
			/* 88 */ imports.NewTable("TCefApplicationCore_DevToolsProtocolLogFile", 0), // property DevToolsProtocolLogFile
			/* 89 */ imports.NewTable("TCefApplicationCore_ForcedDeviceScaleFactor", 0), // property ForcedDeviceScaleFactor
			/* 90 */ imports.NewTable("TCefApplicationCore_DisableZygote", 0), // property DisableZygote
			/* 91 */ imports.NewTable("TCefApplicationCore_UseMockKeyChain", 0), // property UseMockKeyChain
			/* 92 */ imports.NewTable("TCefApplicationCore_DisableRequestHandlingForTesting", 0), // property DisableRequestHandlingForTesting
			/* 93 */ imports.NewTable("TCefApplicationCore_DisablePopupBlocking", 0), // property DisablePopupBlocking
			/* 94 */ imports.NewTable("TCefApplicationCore_DisableBackForwardCache", 0), // property DisableBackForwardCache
			/* 95 */ imports.NewTable("TCefApplicationCore_DisableComponentUpdate", 0), // property DisableComponentUpdate
			/* 96 */ imports.NewTable("TCefApplicationCore_AllowInsecureLocalhost", 0), // property AllowInsecureLocalhost
			/* 97 */ imports.NewTable("TCefApplicationCore_KioskPrinting", 0), // property KioskPrinting
			/* 98 */ imports.NewTable("TCefApplicationCore_TreatInsecureOriginAsSecure", 0), // property TreatInsecureOriginAsSecure
			/* 99 */ imports.NewTable("TCefApplicationCore_NetLogEnabled", 0), // property NetLogEnabled
			/* 100 */ imports.NewTable("TCefApplicationCore_NetLogFile", 0), // property NetLogFile
			/* 101 */ imports.NewTable("TCefApplicationCore_NetLogCaptureMode", 0), // property NetLogCaptureMode
			/* 102 */ imports.NewTable("TCefApplicationCore_RemoteAllowOrigins", 0), // property RemoteAllowOrigins
			/* 103 */ imports.NewTable("TCefApplicationCore_AutoAcceptCamAndMicCapture", 0), // property AutoAcceptCamAndMicCapture
			/* 104 */ imports.NewTable("TCefApplicationCore_UIColorMode", 0), // property UIColorMode
			/* 105 */ imports.NewTable("TCefApplicationCore_DisableHangMonitor", 0), // property DisableHangMonitor
			/* 106 */ imports.NewTable("TCefApplicationCore_HideCrashRestoreBubble", 0), // property HideCrashRestoreBubble
			/* 107 */ imports.NewTable("TCefApplicationCore_IgnoreCertificateErrors", 0), // property IgnoreCertificateErrors
			/* 108 */ imports.NewTable("TCefApplicationCore_WindowsSandboxInfo", 0), // property WindowsSandboxInfo
			/* 109 */ imports.NewTable("TCefApplicationCore_argcCopy", 0), // property ArgcCopy
			/* 110 */ imports.NewTable("TCefApplicationCore_argvCopy", 0), // property ArgvCopy
			/* 111 */ imports.NewTable("TCefApplicationCore_DeleteCache", 0), // property DeleteCache
			/* 112 */ imports.NewTable("TCefApplicationCore_DeleteCookies", 0), // property DeleteCookies
			/* 113 */ imports.NewTable("TCefApplicationCore_CheckCEFFiles", 0), // property CheckCEFFiles
			/* 114 */ imports.NewTable("TCefApplicationCore_ShowMessageDlg", 0), // property ShowMessageDlg
			/* 115 */ imports.NewTable("TCefApplicationCore_MissingBinariesException", 0), // property MissingBinariesException
			/* 116 */ imports.NewTable("TCefApplicationCore_SetCurrentDir", 0), // property SetCurrentDir
			/* 117 */ imports.NewTable("TCefApplicationCore_GlobalContextInitialized", 0), // property GlobalContextInitialized
			/* 118 */ imports.NewTable("TCefApplicationCore_ChromeMajorVer", 0), // property ChromeMajorVer
			/* 119 */ imports.NewTable("TCefApplicationCore_ChromeMinorVer", 0), // property ChromeMinorVer
			/* 120 */ imports.NewTable("TCefApplicationCore_ChromeRelease", 0), // property ChromeRelease
			/* 121 */ imports.NewTable("TCefApplicationCore_ChromeBuild", 0), // property ChromeBuild
			/* 122 */ imports.NewTable("TCefApplicationCore_ChromeVersion", 0), // property ChromeVersion
			/* 123 */ imports.NewTable("TCefApplicationCore_LibCefVersion", 0), // property LibCefVersion
			/* 124 */ imports.NewTable("TCefApplicationCore_LibCefPath", 0), // property LibCefPath
			/* 125 */ imports.NewTable("TCefApplicationCore_ChromeElfPath", 0), // property ChromeElfPath
			/* 126 */ imports.NewTable("TCefApplicationCore_LibLoaded", 0), // property LibLoaded
			/* 127 */ imports.NewTable("TCefApplicationCore_LogProcessInfo", 0), // property LogProcessInfo
			/* 128 */ imports.NewTable("TCefApplicationCore_ReRaiseExceptions", 0), // property ReRaiseExceptions
			/* 129 */ imports.NewTable("TCefApplicationCore_DeviceScaleFactor", 0), // property DeviceScaleFactor
			/* 130 */ imports.NewTable("TCefApplicationCore_LocalesRequired", 0), // property LocalesRequired
			/* 131 */ imports.NewTable("TCefApplicationCore_ProcessType", 0), // property ProcessType
			/* 132 */ imports.NewTable("TCefApplicationCore_MustCreateResourceBundleHandler", 0), // property MustCreateResourceBundleHandler
			/* 133 */ imports.NewTable("TCefApplicationCore_MustCreateBrowserProcessHandler", 0), // property MustCreateBrowserProcessHandler
			/* 134 */ imports.NewTable("TCefApplicationCore_MustCreateRenderProcessHandler", 0), // property MustCreateRenderProcessHandler
			/* 135 */ imports.NewTable("TCefApplicationCore_MustCreateLoadHandler", 0), // property MustCreateLoadHandler
			/* 136 */ imports.NewTable("TCefApplicationCore_OsmodalLoop", 0), // property OsmodalLoop
			/* 137 */ imports.NewTable("TCefApplicationCore_Status", 0), // property Status
			/* 138 */ imports.NewTable("TCefApplicationCore_MissingLibFiles", 0), // property MissingLibFiles
			/* 139 */ imports.NewTable("TCefApplicationCore_MustFreeLibrary", 0), // property MustFreeLibrary
			/* 140 */ imports.NewTable("TCefApplicationCore_ChildProcessesCount", 0), // property ChildProcessesCount
			/* 141 */ imports.NewTable("TCefApplicationCore_UsedMemory", 0), // property UsedMemory
			/* 142 */ imports.NewTable("TCefApplicationCore_TotalSystemMemory", 0), // property TotalSystemMemory
			/* 143 */ imports.NewTable("TCefApplicationCore_AvailableSystemMemory", 0), // property AvailableSystemMemory
			/* 144 */ imports.NewTable("TCefApplicationCore_SystemMemoryLoad", 0), // property SystemMemoryLoad
			/* 145 */ imports.NewTable("TCefApplicationCore_ApiHashUniversal", 0), // property ApiHashUniversal
			/* 146 */ imports.NewTable("TCefApplicationCore_ApiHashPlatform", 0), // property ApiHashPlatform
			/* 147 */ imports.NewTable("TCefApplicationCore_ApiHashCommit", 0), // property ApiHashCommit
			/* 148 */ imports.NewTable("TCefApplicationCore_ExitCode", 0), // property ExitCode
			/* 149 */ imports.NewTable("TCefApplicationCore_LastErrorMessage", 0), // property LastErrorMessage
			/* 150 */ imports.NewTable("TCefApplicationCore_XDisplay", 0), // property XDisplay
			/* 151 */ imports.NewTable("TCefApplicationCore_OnRegCustomSchemes", 0), // event OnRegCustomSchemes
			/* 152 */ imports.NewTable("TCefApplicationCore_OnRegisterCustomPreferences", 0), // event OnRegisterCustomPreferences
			/* 153 */ imports.NewTable("TCefApplicationCore_OnContextInitialized", 0), // event OnContextInitialized
			/* 154 */ imports.NewTable("TCefApplicationCore_OnBeforeChildProcessLaunch", 0), // event OnBeforeChildProcessLaunch
			/* 155 */ imports.NewTable("TCefApplicationCore_OnAlreadyRunningAppRelaunch", 0), // event OnAlreadyRunningAppRelaunch
			/* 156 */ imports.NewTable("TCefApplicationCore_OnScheduleMessagePumpWork", 0), // event OnScheduleMessagePumpWork
			/* 157 */ imports.NewTable("TCefApplicationCore_OnGetDefaultClient", 0), // event OnGetDefaultClient
			/* 158 */ imports.NewTable("TCefApplicationCore_OnGetDefaultRequestContextHandler", 0), // event OnGetDefaultRequestContextHandler
			/* 159 */ imports.NewTable("TCefApplicationCore_OnGetLocalizedString", 0), // event OnGetLocalizedString
			/* 160 */ imports.NewTable("TCefApplicationCore_OnGetDataResource", 0), // event OnGetDataResource
			/* 161 */ imports.NewTable("TCefApplicationCore_OnGetDataResourceForScale", 0), // event OnGetDataResourceForScale
			/* 162 */ imports.NewTable("TCefApplicationCore_OnWebKitInitialized", 0), // event OnWebKitInitialized
			/* 163 */ imports.NewTable("TCefApplicationCore_OnBrowserCreated", 0), // event OnBrowserCreated
			/* 164 */ imports.NewTable("TCefApplicationCore_OnBrowserDestroyed", 0), // event OnBrowserDestroyed
			/* 165 */ imports.NewTable("TCefApplicationCore_OnContextCreated", 0), // event OnContextCreated
			/* 166 */ imports.NewTable("TCefApplicationCore_OnContextReleased", 0), // event OnContextReleased
			/* 167 */ imports.NewTable("TCefApplicationCore_OnUncaughtException", 0), // event OnUncaughtException
			/* 168 */ imports.NewTable("TCefApplicationCore_OnFocusedNodeChanged", 0), // event OnFocusedNodeChanged
			/* 169 */ imports.NewTable("TCefApplicationCore_OnProcessMessageReceived", 0), // event OnProcessMessageReceived
			/* 170 */ imports.NewTable("TCefApplicationCore_OnLoadingStateChange", 0), // event OnLoadingStateChange
			/* 171 */ imports.NewTable("TCefApplicationCore_OnLoadStart", 0), // event OnLoadStart
			/* 172 */ imports.NewTable("TCefApplicationCore_OnLoadEnd", 0), // event OnLoadEnd
			/* 173 */ imports.NewTable("TCefApplicationCore_OnLoadError", 0), // event OnLoadError
		}
	})
	return cefApplicationCoreImport
}
