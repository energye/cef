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

// ICefApplicationCore Parent: IObject
//
//	Parent class of TCefApplication used to simplify the CEF initialization and destruction.
type ICefApplicationCore interface {
	IObject
	// NoSandbox
	//  Set to true(1) to disable the sandbox for sub-processes. See
	//  cef_sandbox_win.h for requirements to enable the sandbox on Windows. Also
	//  configurable using the "no-sandbox" command-line switch.
	NoSandbox() bool // property
	// SetNoSandbox Set NoSandbox
	SetNoSandbox(AValue bool) // property
	// BrowserSubprocessPath
	//  The path to a separate executable that will be launched for sub-processes.
	//  If this value is empty on Windows or Linux then the main process
	//  executable will be used. If this value is empty on macOS then a helper
	//  executable must exist at "Contents/Frameworks/<app>
	//  Helper.app/Contents/MacOS/<app> Helper" in the top-level app bundle. See
	//  the comments on CefExecuteProcess() for details. If this value is
	//  non-empty then it must be an absolute path. Also configurable using the
	//  "browser-subprocess-path" command-line switch.
	BrowserSubprocessPath() string // property
	// SetBrowserSubprocessPath Set BrowserSubprocessPath
	SetBrowserSubprocessPath(AValue string) // property
	// FrameworkDirPath
	//  The path to the CEF framework directory on macOS. If this value is empty
	//  then the framework must exist at "Contents/Frameworks/Chromium Embedded
	//  Framework.framework" in the top-level app bundle. If this value is
	//  non-empty then it must be an absolute path. Also configurable using the
	//  "framework-dir-path" command-line switch.
	FrameworkDirPath() string // property
	// SetFrameworkDirPath Set FrameworkDirPath
	SetFrameworkDirPath(AValue string) // property
	// MainBundlePath
	//  The path to the main bundle on macOS. If this value is empty then it
	//  defaults to the top-level app bundle. If this value is non-empty then it
	//  must be an absolute path. Also configurable using the "main-bundle-path"
	//  command-line switch.
	MainBundlePath() string // property
	// SetMainBundlePath Set MainBundlePath
	SetMainBundlePath(AValue string) // property
	// ChromeRuntime
	//  Set to true(1) to enable use of the Chrome runtime in CEF. This feature
	//  is considered experimental and is not recommended for most users at this
	//  time. See issue #2969 for details.
	ChromeRuntime() bool // property
	// SetChromeRuntime Set ChromeRuntime
	SetChromeRuntime(AValue bool) // property
	// MultiThreadedMessageLoop
	//  Set to true(1) to have the browser process message loop run in a separate
	//  thread. If false(0) then the CefDoMessageLoopWork() function must be
	//  called from your application message loop. This option is only supported
	//  on Windows and Linux.
	MultiThreadedMessageLoop() bool // property
	// SetMultiThreadedMessageLoop Set MultiThreadedMessageLoop
	SetMultiThreadedMessageLoop(AValue bool) // property
	// ExternalMessagePump
	//  Set to true(1) to control browser process main(UI) thread message pump
	//  scheduling via the ICefBrowserProcessHandler.OnScheduleMessagePumpWork()
	//  callback. This option is recommended for use in combination with the
	//  CefDoMessageLoopWork() function in cases where the CEF message loop must
	//  be integrated into an existing application message loop(see additional
	//  comments and warnings on CefDoMessageLoopWork). Enabling this option is
	//  not recommended for most users; leave this option disabled and use either
	//  the CefRunMessageLoop() function or multi_threaded_message_loop if
	//  possible.
	ExternalMessagePump() bool // property
	// SetExternalMessagePump Set ExternalMessagePump
	SetExternalMessagePump(AValue bool) // property
	// WindowlessRenderingEnabled
	//  Set to true(1) to enable windowless(off-screen) rendering support. Do
	//  not enable this value if the application does not use windowless rendering
	//  as it may reduce rendering performance on some systems.
	WindowlessRenderingEnabled() bool // property
	// SetWindowlessRenderingEnabled Set WindowlessRenderingEnabled
	SetWindowlessRenderingEnabled(AValue bool) // property
	// CommandLineArgsDisabled
	//  Set to true(1) to disable configuration of browser process features using
	//  standard CEF and Chromium command-line arguments. Configuration can still
	//  be specified using CEF data structures or via the
	//  ICefApp.OnBeforeCommandLineProcessing() method.
	CommandLineArgsDisabled() bool // property
	// SetCommandLineArgsDisabled Set CommandLineArgsDisabled
	SetCommandLineArgsDisabled(AValue bool) // property
	// Cache
	//  The location where data for the global browser cache will be stored on
	//  disk. If this value is non-empty then it must be an absolute path that is
	//  either equal to or a child directory of TCefSettings.root_cache_path. If
	//  this value is empty then browsers will be created in "incognito mode"
	//  where in-memory caches are used for storage and no data is persisted to
	//  disk. HTML5 databases such as localStorage will only persist across
	//  sessions if a cache path is specified. Can be overridden for individual
	//  CefRequestContext instances via the TCefRequestContextSettings.cache_path
	//  value. When using the Chrome runtime the "default" profile will be used if
	//  |cache_path| and |root_cache_path| have the same value.
	Cache() string // property
	// SetCache Set Cache
	SetCache(AValue string) // property
	// RootCache
	//  The root directory that all TCefSettings.cache_path and
	//  TCefRequestContextSettings.cache_path values must have in common. If this
	//  value is empty and TCefSettings.cache_path is non-empty then it will
	//  default to the TCefSettings.cache_path value. If both values are empty
	//  then the default platform-specific directory will be used
	//  ("~/.config/cef_user_data" directory on Linux, "~/Library/Application
	//  Support/CEF/User Data" directory on MacOS, "AppData\Local\CEF\User Data"
	//  directory under the user profile directory on Windows). If this value is
	//  non-empty then it must be an absolute path. Failure to set this value
	//  correctly may result in the sandbox blocking read/write access to certain
	//  files.
	RootCache() string // property
	// SetRootCache Set RootCache
	SetRootCache(AValue string) // property
	// PersistSessionCookies
	//  To persist session cookies(cookies without an expiry date or validity
	//  interval) by default when using the global cookie manager set this value
	//  to true(1). Session cookies are generally intended to be transient and
	//  most Web browsers do not persist them. A |cache_path| value must also be
	//  specified to enable this feature. Also configurable using the
	//  "persist-session-cookies" command-line switch. Can be overridden for
	//  individual CefRequestContext instances via the
	//  TCefRequestContextSettings.persist_session_cookies value.
	PersistSessionCookies() bool // property
	// SetPersistSessionCookies Set PersistSessionCookies
	SetPersistSessionCookies(AValue bool) // property
	// PersistUserPreferences
	//  To persist user preferences as a JSON file in the cache path directory set
	//  this value to true(1). A |cache_path| value must also be specified
	//  to enable this feature. Also configurable using the
	//  "persist-user-preferences" command-line switch. Can be overridden for
	//  individual CefRequestContext instances via the
	//  TCefRequestContextSettings.persist_user_preferences value.
	PersistUserPreferences() bool // property
	// SetPersistUserPreferences Set PersistUserPreferences
	SetPersistUserPreferences(AValue bool) // property
	// UserAgent
	//  Value that will be returned as the User-Agent HTTP header. If empty the
	//  default User-Agent string will be used. Also configurable using the
	//  "user-agent" command-line switch.
	UserAgent() string // property
	// SetUserAgent Set UserAgent
	SetUserAgent(AValue string) // property
	// UserAgentProduct
	//  Value that will be inserted as the product portion of the default
	//  User-Agent string. If empty the Chromium product version will be used. If
	//  |userAgent| is specified this value will be ignored. Also configurable
	//  using the "user-agent-product" command-line switch.
	UserAgentProduct() string // property
	// SetUserAgentProduct Set UserAgentProduct
	SetUserAgentProduct(AValue string) // property
	// Locale
	//  The locale string that will be passed to WebKit. If empty the default
	//  locale of "en-US" will be used. This value is ignored on Linux where
	//  locale is determined using environment variable parsing with the
	//  precedence order: LANGUAGE, LC_ALL, LC_MESSAGES and LANG. Also
	//  configurable using the "lang" command-line switch.
	Locale() string // property
	// SetLocale Set Locale
	SetLocale(AValue string) // property
	// LogFile
	//  The directory and file name to use for the debug log. If empty a default
	//  log file name and location will be used. On Windows and Linux a
	//  "debug.log" file will be written in the main executable directory. On
	//  MacOS a "~/Library/Logs/[app name]_debug.log" file will be written where
	//  [app name] is the name of the main app executable. Also configurable using
	//  the "log-file" command-line switch.
	LogFile() string // property
	// SetLogFile Set LogFile
	SetLogFile(AValue string) // property
	// LogSeverity
	//  The log severity. Only messages of this severity level or higher will be
	//  logged. When set to DISABLE no messages will be written to the log file,
	//  but FATAL messages will still be output to stderr. Also configurable using
	//  the "log-severity" command-line switch with a value of "verbose", "info",
	//  "warning", "error", "fatal" or "disable".
	LogSeverity() TCefLogSeverity // property
	// SetLogSeverity Set LogSeverity
	SetLogSeverity(AValue TCefLogSeverity) // property
	// LogItems
	//  The log items prepended to each log line. If not set the default log items
	//  will be used. Also configurable using the "log-items" command-line switch
	//  with a value of "none" for no log items, or a comma-delimited list of
	//  values "pid", "tid", "timestamp" or "tickcount" for custom log items.
	LogItems() TCefLogItems // property
	// SetLogItems Set LogItems
	SetLogItems(AValue TCefLogItems) // property
	// JavaScriptFlags
	//  Custom flags that will be used when initializing the V8 JavaScript engine.
	//  The consequences of using custom flags may not be well tested. Also
	//  configurable using the "js-flags" command-line switch.
	JavaScriptFlags() string // property
	// SetJavaScriptFlags Set JavaScriptFlags
	SetJavaScriptFlags(AValue string) // property
	// ResourcesDirPath
	//  The fully qualified path for the resources directory. If this value is
	//  empty the *.pak files must be located in the module directory on
	//  Windows/Linux or the app bundle Resources directory on MacOS. If this
	//  value is non-empty then it must be an absolute path. Also configurable
	//  using the "resources-dir-path" command-line switch.
	ResourcesDirPath() string // property
	// SetResourcesDirPath Set ResourcesDirPath
	SetResourcesDirPath(AValue string) // property
	// LocalesDirPath
	//  The fully qualified path for the locales directory. If this value is empty
	//  the locales directory must be located in the module directory. If this
	//  value is non-empty then it must be an absolute path. This value is ignored
	//  on MacOS where pack files are always loaded from the app bundle Resources
	//  directory. Also configurable using the "locales-dir-path" command-line
	//  switch.
	LocalesDirPath() string // property
	// SetLocalesDirPath Set LocalesDirPath
	SetLocalesDirPath(AValue string) // property
	// PackLoadingDisabled
	//  Set to true(1) to disable loading of pack files for resources and
	//  locales. A resource bundle handler must be provided for the browser and
	//  render processes via ICefApp.GetResourceBundleHandler() if loading of pack
	//  files is disabled. Also configurable using the "disable-pack-loading"
	//  command- line switch.
	PackLoadingDisabled() bool // property
	// SetPackLoadingDisabled Set PackLoadingDisabled
	SetPackLoadingDisabled(AValue bool) // property
	// RemoteDebuggingPort
	//  Set to a value between 1024 and 65535 to enable remote debugging on the
	//  specified port. Also configurable using the "remote-debugging-port"
	//  command-line switch. Remote debugging can be accessed by loading the
	//  chrome://inspect page in Google Chrome. Port numbers 9222 and 9229 are
	//  discoverable by default. Other port numbers may need to be configured via
	//  "Discover network targets" on the Devices tab.
	RemoteDebuggingPort() int32 // property
	// SetRemoteDebuggingPort Set RemoteDebuggingPort
	SetRemoteDebuggingPort(AValue int32) // property
	// UncaughtExceptionStackSize
	//  The number of stack trace frames to capture for uncaught exceptions.
	//  Specify a positive value to enable the
	//  ICefRenderProcessHandler.OnUncaughtException() callback. Specify 0
	//  (default value) and OnUncaughtException() will not be called. Also
	//  configurable using the "uncaught-exception-stack-size" command-line
	//  switch.
	UncaughtExceptionStackSize() int32 // property
	// SetUncaughtExceptionStackSize Set UncaughtExceptionStackSize
	SetUncaughtExceptionStackSize(AValue int32) // property
	// BackgroundColor
	//  Background color used for the browser before a document is loaded and when
	//  no document color is specified. The alpha component must be either fully
	//  opaque(0xFF) or fully transparent(0x00). If the alpha component is fully
	//  opaque then the RGB components will be used as the background color. If
	//  the alpha component is fully transparent for a windowed browser then the
	//  default value of opaque white be used. If the alpha component is fully
	//  transparent for a windowless(off-screen) browser then transparent
	//  painting will be enabled.
	BackgroundColor() TCefColor // property
	// SetBackgroundColor Set BackgroundColor
	SetBackgroundColor(AValue TCefColor) // property
	// AcceptLanguageList
	//  Comma delimited ordered list of language codes without any whitespace that
	//  will be used in the "Accept-Language" HTTP request header and
	//  "navigator.language" JS attribute. Can be overridden for individual
	//  ICefRequestContext instances via the
	//  TCefRequestContextSettingsCefRequestContextSettings.accept_language_list value.
	AcceptLanguageList() string // property
	// SetAcceptLanguageList Set AcceptLanguageList
	SetAcceptLanguageList(AValue string) // property
	// CookieableSchemesList
	//  Comma delimited list of schemes supported by the associated
	//  ICefCookieManager. If |cookieable_schemes_exclude_defaults| is false(0)
	//  the default schemes("http", "https", "ws" and "wss") will also be
	//  supported. Not specifying a |cookieable_schemes_list| value and setting
	//  |cookieable_schemes_exclude_defaults| to true(1) will disable all loading
	//  and saving of cookies. These settings will only impact the global
	//  ICefRequestContext. Individual ICefRequestContext instances can be
	//  configured via the TCefRequestContextSettings.cookieable_schemes_list and
	//  TCefRequestContextSettings.cookieable_schemes_exclude_defaults values.
	CookieableSchemesList() string // property
	// SetCookieableSchemesList Set CookieableSchemesList
	SetCookieableSchemesList(AValue string) // property
	// CookieableSchemesExcludeDefaults
	//  See the CookieableSchemesList property.
	CookieableSchemesExcludeDefaults() bool // property
	// SetCookieableSchemesExcludeDefaults Set CookieableSchemesExcludeDefaults
	SetCookieableSchemesExcludeDefaults(AValue bool) // property
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
	ChromePolicyId() string // property
	// SetChromePolicyId Set ChromePolicyId
	SetChromePolicyId(AValue string) // property
	// SingleProcess
	//  Runs the renderer and plugins in the same process as the browser.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --single-process</a>
	SingleProcess() bool // property
	// SetSingleProcess Set SingleProcess
	SetSingleProcess(AValue bool) // property
	// EnableMediaStream
	//  Enable media(WebRTC audio/video) streaming.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-media-stream</a>
	EnableMediaStream() bool // property
	// SetEnableMediaStream Set EnableMediaStream
	SetEnableMediaStream(AValue bool) // property
	// EnableSpeechInput
	//  Enable speech input(x-webkit-speech).
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-speech-input</a>
	EnableSpeechInput() bool // property
	// SetEnableSpeechInput Set EnableSpeechInput
	SetEnableSpeechInput(AValue bool) // property
	// UseFakeUIForMediaStream
	//  Bypass the media stream infobar by selecting the default device for media streams(e.g. WebRTC). Works with --use-fake-device-for-media-stream.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --single-process</a>
	UseFakeUIForMediaStream() bool // property
	// SetUseFakeUIForMediaStream Set UseFakeUIForMediaStream
	SetUseFakeUIForMediaStream(AValue bool) // property
	// EnableUsermediaScreenCapturing
	//  Enable screen capturing support for MediaStream API.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-usermedia-screen-capturing</a>
	EnableUsermediaScreenCapturing() bool // property
	// SetEnableUsermediaScreenCapturing Set EnableUsermediaScreenCapturing
	SetEnableUsermediaScreenCapturing(AValue bool) // property
	// EnableGPU
	//  Enable GPU hardware acceleration.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu-compositing</a>
	EnableGPU() bool // property
	// SetEnableGPU Set EnableGPU
	SetEnableGPU(AValue bool) // property
	// EnableFeatures
	//  List of feature names to enable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-features</a>
	//  The list of features you can enable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	EnableFeatures() string // property
	// SetEnableFeatures Set EnableFeatures
	SetEnableFeatures(AValue string) // property
	// DisableFeatures
	//  List of feature names to disable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-features</a>
	//  The list of features you can disable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	DisableFeatures() string // property
	// SetDisableFeatures Set DisableFeatures
	SetDisableFeatures(AValue string) // property
	// EnableBlinkFeatures
	//  Enable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-blink-features</a>
	//  The list of Blink features you can enable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	EnableBlinkFeatures() string // property
	// SetEnableBlinkFeatures Set EnableBlinkFeatures
	SetEnableBlinkFeatures(AValue string) // property
	// DisableBlinkFeatures
	//  Disable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-blink-features</a>
	//  The list of Blink features you can disable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	DisableBlinkFeatures() string // property
	// SetDisableBlinkFeatures Set DisableBlinkFeatures
	SetDisableBlinkFeatures(AValue string) // property
	// BlinkSettings
	//  Set blink settings. Format is <name>[=<value],<name>[=<value>],...
	//  The names are declared in Settings.json5. For boolean type, use "true", "false",
	//  or omit '=<value>' part to set to true. For enum type, use the int value of the
	//  enum value. Applied after other command line flags and prefs.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --blink-settings</a>
	//  The list of Blink settings you can disable is here:
	//  https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/settings.json5
	BlinkSettings() string // property
	// SetBlinkSettings Set BlinkSettings
	SetBlinkSettings(AValue string) // property
	// ForceFieldTrials
	//  This option can be used to force field trials when testing changes locally.
	//  The argument is a list of name and value pairs, separated by slashes.
	//  If a trial name is prefixed with an asterisk, that trial will start activated.
	//  For example, the following argument defines two trials, with the second one
	//  activated: "GoogleNow/Enable/*MaterialDesignNTP/Default/" This option can also
	//  be used by the browser process to send the list of trials to a non-browser
	//  process, using the same format. See FieldTrialList::CreateTrialsFromString()
	//  in field_trial.h for details.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrials</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:base/base_switches.cc
	ForceFieldTrials() string // property
	// SetForceFieldTrials Set ForceFieldTrials
	SetForceFieldTrials(AValue string) // property
	// ForceFieldTrialParams
	//  This option can be used to force parameters of field trials when testing
	//  changes locally. The argument is a param list of(key, value) pairs prefixed
	//  by an associated(trial, group) pair. You specify the param list for multiple
	//  (trial, group) pairs with a comma separator.
	//  Example: "Trial1.Group1:k1/v1/k2/v2,Trial2.Group2:k3/v3/k4/v4"
	//  Trial names, groups names, parameter names, and value should all be URL
	//  escaped for all non-alphanumeric characters.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrial-params</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:components/variations/variations_switches.cc
	ForceFieldTrialParams() string // property
	// SetForceFieldTrialParams Set ForceFieldTrialParams
	SetForceFieldTrialParams(AValue string) // property
	// SmoothScrolling
	//  On platforms that support it, enables smooth scroll animation.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-smooth-scrolling</a>
	SmoothScrolling() TCefState // property
	// SetSmoothScrolling Set SmoothScrolling
	SetSmoothScrolling(AValue TCefState) // property
	// MuteAudio
	//  Mutes audio sent to the audio device so it is not audible during automated testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --mute-audio</a>
	MuteAudio() bool // property
	// SetMuteAudio Set MuteAudio
	SetMuteAudio(AValue bool) // property
	// SitePerProcess
	//  Enforces a one-site-per-process security policy: Each renderer process, for its
	//  whole lifetime, is dedicated to rendering pages for just one site. Thus, pages
	//  from different sites are never in the same process. A renderer process's access
	//  rights are restricted based on its site.All cross-site navigations force process
	//  swaps. <iframe>s are rendered out-of-process whenever the src= is cross-site.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --site-per-process</a>
	//  More details here:
	//  https://www.chromium.org/developers/design-documents/site-isolation
	//  https://www.chromium.org/developers/design-documents/process-models
	SitePerProcess() bool // property
	// SetSitePerProcess Set SitePerProcess
	SetSitePerProcess(AValue bool) // property
	// DisableWebSecurity
	//  Don't enforce the same-origin policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-web-security</a>
	DisableWebSecurity() bool // property
	// SetDisableWebSecurity Set DisableWebSecurity
	SetDisableWebSecurity(AValue bool) // property
	// DisablePDFExtension
	//  Disable the PDF extension.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-pdf-extension</a>
	DisablePDFExtension() bool // property
	// SetDisablePDFExtension Set DisablePDFExtension
	SetDisablePDFExtension(AValue bool) // property
	// DisableSiteIsolationTrials
	//  Disables site isolation.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-site-isolation-trials</a>
	DisableSiteIsolationTrials() bool // property
	// SetDisableSiteIsolationTrials Set DisableSiteIsolationTrials
	SetDisableSiteIsolationTrials(AValue bool) // property
	// DisableChromeLoginPrompt
	//  Delegate all login requests to the client GetAuthCredentials
	//  callback when using the Chrome runtime.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-chrome-login-prompt</a>
	DisableChromeLoginPrompt() bool // property
	// SetDisableChromeLoginPrompt Set DisableChromeLoginPrompt
	SetDisableChromeLoginPrompt(AValue bool) // property
	// DisableExtensions
	//  Disable extensions.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-extensions</a>
	DisableExtensions() bool // property
	// SetDisableExtensions Set DisableExtensions
	SetDisableExtensions(AValue bool) // property
	// AutoplayPolicy
	//  Autoplay policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --autoplay-policy</a>
	AutoplayPolicy() TCefAutoplayPolicy // property
	// SetAutoplayPolicy Set AutoplayPolicy
	SetAutoplayPolicy(AValue TCefAutoplayPolicy) // property
	// DisableBackgroundNetworking
	//  Disable several subsystems which run network requests in the background.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-background-networking</a>
	DisableBackgroundNetworking() bool // property
	// SetDisableBackgroundNetworking Set DisableBackgroundNetworking
	SetDisableBackgroundNetworking(AValue bool) // property
	// MetricsRecordingOnly
	//  Enables the recording of metrics reports but disables reporting.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --metrics-recording-only</a>
	MetricsRecordingOnly() bool // property
	// SetMetricsRecordingOnly Set MetricsRecordingOnly
	SetMetricsRecordingOnly(AValue bool) // property
	// AllowFileAccessFromFiles
	//  By default, file:// URIs cannot read other file:// URIs. This is an override for developers who need the old behavior for testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-file-access-from-files</a>
	AllowFileAccessFromFiles() bool // property
	// SetAllowFileAccessFromFiles Set AllowFileAccessFromFiles
	SetAllowFileAccessFromFiles(AValue bool) // property
	// AllowRunningInsecureContent
	//  By default, an https page cannot run JavaScript, CSS or plugins from http URLs. This provides an override to get the old insecure behavior.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-running-insecure-content</a>
	AllowRunningInsecureContent() bool // property
	// SetAllowRunningInsecureContent Set AllowRunningInsecureContent
	SetAllowRunningInsecureContent(AValue bool) // property
	// EnablePrintPreview
	//  Enable print preview.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-print-preview</a>
	EnablePrintPreview() bool // property
	// SetEnablePrintPreview Set EnablePrintPreview
	SetEnablePrintPreview(AValue bool) // property
	// DisableJavascript
	//  Default encoding.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --default-encoding</a>
	//  Disable JavaScript.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript</a>
	DisableJavascript() bool // property
	// SetDisableJavascript Set DisableJavascript
	SetDisableJavascript(AValue bool) // property
	// DisableJavascriptCloseWindows
	//  Disable closing of windows via JavaScript.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-close-windows</a>
	DisableJavascriptCloseWindows() bool // property
	// SetDisableJavascriptCloseWindows Set DisableJavascriptCloseWindows
	SetDisableJavascriptCloseWindows(AValue bool) // property
	// DisableJavascriptAccessClipboard
	//  Disable clipboard access via JavaScript.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-access-clipboard</a>
	DisableJavascriptAccessClipboard() bool // property
	// SetDisableJavascriptAccessClipboard Set DisableJavascriptAccessClipboard
	SetDisableJavascriptAccessClipboard(AValue bool) // property
	// DisableJavascriptDomPaste
	//  Disable DOM paste via JavaScript execCommand("paste").
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-javascript-dom-paste</a>
	DisableJavascriptDomPaste() bool // property
	// SetDisableJavascriptDomPaste Set DisableJavascriptDomPaste
	SetDisableJavascriptDomPaste(AValue bool) // property
	// AllowUniversalAccessFromFileUrls
	//  Allow universal access from file URLs.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --allow-universal-access-from-files</a>
	AllowUniversalAccessFromFileUrls() bool // property
	// SetAllowUniversalAccessFromFileUrls Set AllowUniversalAccessFromFileUrls
	SetAllowUniversalAccessFromFileUrls(AValue bool) // property
	// DisableImageLoading
	//  Disable loading of images from the network. A cached image will still be rendered if requested.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-image-loading</a>
	DisableImageLoading() bool // property
	// SetDisableImageLoading Set DisableImageLoading
	SetDisableImageLoading(AValue bool) // property
	// ImageShrinkStandaloneToFit
	//  Shrink stand-alone images to fit.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --image-shrink-standalone-to-fit</a>
	ImageShrinkStandaloneToFit() bool // property
	// SetImageShrinkStandaloneToFit Set ImageShrinkStandaloneToFit
	SetImageShrinkStandaloneToFit(AValue bool) // property
	// DisableTextAreaResize
	//  Disable resizing of text areas.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-text-area-resize</a>
	DisableTextAreaResize() bool // property
	// SetDisableTextAreaResize Set DisableTextAreaResize
	SetDisableTextAreaResize(AValue bool) // property
	// DisableTabToLinks
	//  Disable using the tab key to advance focus to links.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-tab-to-links</a>
	DisableTabToLinks() bool // property
	// SetDisableTabToLinks Set DisableTabToLinks
	SetDisableTabToLinks(AValue bool) // property
	// EnableProfanityFilter
	//  Enable the speech input profanity filter.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --enable-profanity-filter</a>
	EnableProfanityFilter() bool // property
	// SetEnableProfanityFilter Set EnableProfanityFilter
	SetEnableProfanityFilter(AValue bool) // property
	// DisableSpellChecking
	//  Disable spell checking.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-spell-checking</a>
	DisableSpellChecking() bool // property
	// SetDisableSpellChecking Set DisableSpellChecking
	SetDisableSpellChecking(AValue bool) // property
	// OverrideSpellCheckLang
	//  Override the default spellchecking language which comes from locales.pak.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --override-spell-check-lang</a>
	OverrideSpellCheckLang() string // property
	// SetOverrideSpellCheckLang Set OverrideSpellCheckLang
	SetOverrideSpellCheckLang(AValue string) // property
	// TouchEvents
	//  Enable support for touch event feature detection.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --touch-events</a>
	TouchEvents() TCefState // property
	// SetTouchEvents Set TouchEvents
	SetTouchEvents(AValue TCefState) // property
	// DisableReadingFromCanvas
	//  Taints all <canvas> elements, regardless of origin.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-reading-from-canvas</a>
	DisableReadingFromCanvas() bool // property
	// SetDisableReadingFromCanvas Set DisableReadingFromCanvas
	SetDisableReadingFromCanvas(AValue bool) // property
	// HyperlinkAuditing
	//  Don't send hyperlink auditing pings.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-pings</a>
	HyperlinkAuditing() bool // property
	// SetHyperlinkAuditing Set HyperlinkAuditing
	SetHyperlinkAuditing(AValue bool) // property
	// DisableNewBrowserInfoTimeout
	//  Disable the timeout for delivering new browser info to the renderer process.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-new-browser-info-timeout</a>
	DisableNewBrowserInfoTimeout() bool // property
	// SetDisableNewBrowserInfoTimeout Set DisableNewBrowserInfoTimeout
	SetDisableNewBrowserInfoTimeout(AValue bool) // property
	// DevToolsProtocolLogFile
	//  File used for logging DevTools protocol messages.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --devtools-protocol-log-file</a>
	DevToolsProtocolLogFile() string // property
	// SetDevToolsProtocolLogFile Set DevToolsProtocolLogFile
	SetDevToolsProtocolLogFile(AValue string) // property
	// ForcedDeviceScaleFactor
	//  Overrides the device scale factor for the browser UI and the contents.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-device-scale-factor</a>
	ForcedDeviceScaleFactor() (resultFloat32 float32) // property
	// SetForcedDeviceScaleFactor Set ForcedDeviceScaleFactor
	SetForcedDeviceScaleFactor(AValue float32) // property
	// DisableZygote
	//  Disables the use of a zygote process for forking child processes. Instead, child processes will be forked and exec'd directly.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-zygote</a>
	DisableZygote() bool // property
	// SetDisableZygote Set DisableZygote
	SetDisableZygote(AValue bool) // property
	// UseMockKeyChain
	//  Uses mock keychain for testing purposes, which prevents blocking dialogs from causing timeouts.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --use-mock-keychain</a>
	UseMockKeyChain() bool // property
	// SetUseMockKeyChain Set UseMockKeyChain
	SetUseMockKeyChain(AValue bool) // property
	// DisableRequestHandlingForTesting
	//  Disable request handling in CEF to faciliate debugging of network-related issues.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --disable-request-handling-for-testing</a>
	DisableRequestHandlingForTesting() bool // property
	// SetDisableRequestHandlingForTesting Set DisableRequestHandlingForTesting
	SetDisableRequestHandlingForTesting(AValue bool) // property
	// DisablePopupBlocking
	//  Disables pop-up blocking.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-popup-blocking</a>
	DisablePopupBlocking() bool // property
	// SetDisablePopupBlocking Set DisablePopupBlocking
	SetDisablePopupBlocking(AValue bool) // property
	// DisableBackForwardCache
	//  Disables the BackForwardCache feature.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-back-forward-cache</a>
	DisableBackForwardCache() bool // property
	// SetDisableBackForwardCache Set DisableBackForwardCache
	SetDisableBackForwardCache(AValue bool) // property
	// DisableComponentUpdate
	//  Disable the component updater. Widevine will not be downloaded or initialized.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-component-update</a>
	DisableComponentUpdate() bool // property
	// SetDisableComponentUpdate Set DisableComponentUpdate
	SetDisableComponentUpdate(AValue bool) // property
	// AllowInsecureLocalhost
	//  Enables TLS/SSL errors on localhost to be ignored(no interstitial, no blocking of requests).
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-insecure-localhost</a>
	AllowInsecureLocalhost() bool // property
	// SetAllowInsecureLocalhost Set AllowInsecureLocalhost
	SetAllowInsecureLocalhost(AValue bool) // property
	// KioskPrinting
	//  Enable automatically pressing the print button in print preview.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --kiosk-printing</a>
	KioskPrinting() bool // property
	// SetKioskPrinting Set KioskPrinting
	SetKioskPrinting(AValue bool) // property
	// TreatInsecureOriginAsSecure
	//  Treat given(insecure) origins as secure origins.
	//  Multiple origins can be supplied as a comma-separated list.
	//  For the definition of secure contexts, see https://w3c.github.io/webappsec-secure-contexts/
	//  and https://www.w3.org/TR/powerful-features/#is-origin-trustworthy
	//  Example: --unsafely-treat-insecure-origin-as-secure=http://a.test,http://b.test
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --unsafely-treat-insecure-origin-as-secure</a>
	TreatInsecureOriginAsSecure() string // property
	// SetTreatInsecureOriginAsSecure Set TreatInsecureOriginAsSecure
	SetTreatInsecureOriginAsSecure(AValue string) // property
	// NetLogEnabled
	//  Enables saving net log events to a file.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-net-log</a>
	NetLogEnabled() bool // property
	// SetNetLogEnabled Set NetLogEnabled
	SetNetLogEnabled(AValue bool) // property
	// NetLogFile
	//  File name used to log net events. If a value is given,
	//  it used as the path the the file, otherwise the file is named netlog.json
	//  and placed in the user data directory.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-net-log</a>
	NetLogFile() string // property
	// SetNetLogFile Set NetLogFile
	SetNetLogFile(AValue string) // property
	// NetLogCaptureMode
	//  Sets the granularity of events to capture in the network log.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --net-log-capture-mode</a>
	NetLogCaptureMode() TCefNetLogCaptureMode // property
	// SetNetLogCaptureMode Set NetLogCaptureMode
	SetNetLogCaptureMode(AValue TCefNetLogCaptureMode) // property
	// RemoteAllowOrigins
	//  Enables web socket connections from the specified origins only. '*' allows any origin.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-allow-origins</a>
	RemoteAllowOrigins() string // property
	// SetRemoteAllowOrigins Set RemoteAllowOrigins
	SetRemoteAllowOrigins(AValue string) // property
	// AutoAcceptCamAndMicCapture
	//  Bypasses the dialog prompting the user for permission to capture cameras and microphones.
	//  Useful in automatic tests of video-conferencing Web applications. This is nearly
	//  identical to kUseFakeUIForMediaStream, with the exception being that this flag does NOT
	//  affect screen-capture.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --auto-accept-camera-and-microphone-capture</a>
	AutoAcceptCamAndMicCapture() bool // property
	// SetAutoAcceptCamAndMicCapture Set AutoAcceptCamAndMicCapture
	SetAutoAcceptCamAndMicCapture(AValue bool) // property
	// UIColorMode
	//  Forces light or dark mode in UI for platforms that support it.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switches: --force-dark-mode --force-light-mode</a>
	UIColorMode() TCefUIColorMode // property
	// SetUIColorMode Set UIColorMode
	SetUIColorMode(AValue TCefUIColorMode) // property
	// IgnoreCertificateErrors
	//  Ignores certificate-related errors.
	//  <a href="https://source.chromium.org/chromium/chromium/src/+/main:components/network_session_configurator/common/network_switch_list.h">Uses the following command line switch: --ignore-certificate-errors</a>
	IgnoreCertificateErrors() bool // property
	// SetIgnoreCertificateErrors Set IgnoreCertificateErrors
	SetIgnoreCertificateErrors(AValue bool) // property
	// WindowsSandboxInfo
	//  Pointer to the sandbox info. Currently unused in Delphi and Lazarus.
	WindowsSandboxInfo() uintptr // property
	// SetWindowsSandboxInfo Set WindowsSandboxInfo
	SetWindowsSandboxInfo(AValue uintptr) // property
	// ArgcCopy
	//  argc parameter copy used in Linux only.
	ArgcCopy() int32 // property
	// ArgvCopy
	//  argv parameter copy used in Linux only.
	ArgvCopy() PString // property
	// DeleteCache
	//  Used to delete all the cache files before CEF is initialized.
	DeleteCache() bool // property
	// SetDeleteCache Set DeleteCache
	SetDeleteCache(AValue bool) // property
	// DeleteCookies
	//  Used to delete all the cookies before CEF is initialized.
	DeleteCookies() bool // property
	// SetDeleteCookies Set DeleteCookies
	SetDeleteCookies(AValue bool) // property
	// CheckCEFFiles
	//  Checks if the CEF binaries are present and the DLL version.
	CheckCEFFiles() bool // property
	// SetCheckCEFFiles Set CheckCEFFiles
	SetCheckCEFFiles(AValue bool) // property
	// ShowMessageDlg
	//  Set to true when you need to use a showmessage dialog to show the error messages.
	ShowMessageDlg() bool // property
	// SetShowMessageDlg Set ShowMessageDlg
	SetShowMessageDlg(AValue bool) // property
	// MissingBinariesException
	//  Raise an exception when the CEF binaries check fails.
	MissingBinariesException() bool // property
	// SetMissingBinariesException Set MissingBinariesException
	SetMissingBinariesException(AValue bool) // property
	// SetCurrentDir
	//  Used to set the current directory when the CEF libraries are loaded. This is required if the application is launched from a different application.
	SetCurrentDir() bool // property
	// SetSetCurrentDir Set SetCurrentDir
	SetSetCurrentDir(AValue bool) // property
	// GlobalContextInitialized
	//  Set to True when the global context is initialized and the application can start creating web browsers.
	GlobalContextInitialized() bool // property
	// ChromeMajorVer
	//  Returns the major version information from Chromium.
	ChromeMajorVer() uint16 // property
	// ChromeMinorVer
	//  Returns the minor version information from Chromium.
	ChromeMinorVer() uint16 // property
	// ChromeRelease
	//  Returns the release version information from Chromium.
	ChromeRelease() uint16 // property
	// ChromeBuild
	//  Returns the build version information from Chromium.
	ChromeBuild() uint16 // property
	// ChromeVersion
	//  Returns the full version information from Chromium.
	ChromeVersion() string // property
	// LibCefVersion
	//  Complete libcef version information.
	LibCefVersion() string // property
	// LibCefPath
	//  Path to libcef.dll or libcef.so
	LibCefPath() string // property
	// ChromeElfPath
	//  Returns the path to chrome_elf.dll.
	ChromeElfPath() string // property
	// LibLoaded
	//  Set to true when TCEFApplicationCore has loaded the CEF libraries.
	LibLoaded() bool // property
	// LogProcessInfo
	//  Add a debug log information line when the CEF libraries are loaded.
	LogProcessInfo() bool // property
	// SetLogProcessInfo Set LogProcessInfo
	SetLogProcessInfo(AValue bool) // property
	// ReRaiseExceptions
	//  Set to true to raise all exceptions.
	ReRaiseExceptions() bool // property
	// SetReRaiseExceptions Set ReRaiseExceptions
	SetReRaiseExceptions(AValue bool) // property
	// DeviceScaleFactor
	//  Returns the device scale factor used in OSR mode.
	DeviceScaleFactor() (resultFloat32 float32) // property
	// LocalesRequired
	//  List of locale files that will be checked with CheckCEFFiles.
	LocalesRequired() string // property
	// SetLocalesRequired Set LocalesRequired
	SetLocalesRequired(AValue string) // property
	// ProcessType
	//  CEF process type currently running.
	ProcessType() TCefProcessType // property
	// MustCreateResourceBundleHandler
	//  Force the creation of ICefResourceBundleHandler.
	MustCreateResourceBundleHandler() bool // property
	// SetMustCreateResourceBundleHandler Set MustCreateResourceBundleHandler
	SetMustCreateResourceBundleHandler(AValue bool) // property
	// MustCreateBrowserProcessHandler
	//  Force the creation of ICefBrowserProcessHandler.
	MustCreateBrowserProcessHandler() bool // property
	// SetMustCreateBrowserProcessHandler Set MustCreateBrowserProcessHandler
	SetMustCreateBrowserProcessHandler(AValue bool) // property
	// MustCreateRenderProcessHandler
	//  Force the creation of ICefRenderProcessHandler.
	MustCreateRenderProcessHandler() bool // property
	// SetMustCreateRenderProcessHandler Set MustCreateRenderProcessHandler
	SetMustCreateRenderProcessHandler(AValue bool) // property
	// MustCreateLoadHandler
	//  Force the creation of ICefLoadHandler.
	MustCreateLoadHandler() bool // property
	// SetMustCreateLoadHandler Set MustCreateLoadHandler
	SetMustCreateLoadHandler(AValue bool) // property
	// OsmodalLoop
	//  Set to true(1) before calling Windows APIs like TrackPopupMenu that enter a
	//  modal message loop. Set to false(0) after exiting the modal message loop.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_app_win.h">CEF source file: /include/internal/cef_app_win.h(cef_set_osmodal_loop)</a>
	// SetOsmodalLoop Set OsmodalLoop
	SetOsmodalLoop(AValue bool) // property
	// Status
	//  Returns the TCEFApplicationCore initialization status.
	Status() TCefAplicationStatus // property
	// MissingLibFiles
	//  List of missing CEF library files.
	MissingLibFiles() string // property
	// MustFreeLibrary
	//  Set to true to free the library handle when TCEFApplicationCore is destroyed.
	MustFreeLibrary() bool // property
	// SetMustFreeLibrary Set MustFreeLibrary
	SetMustFreeLibrary(AValue bool) // property
	// ChildProcessesCount
	//  Returns the number of CEF subprocesses running at that moment.
	ChildProcessesCount() int32 // property
	// UsedMemory
	//  Total used memory by all CEF processes.
	UsedMemory() (resultUint64 uint64) // property
	// TotalSystemMemory
	//  Total system memory in Windows.
	TotalSystemMemory() (resultUint64 uint64) // property
	// AvailableSystemMemory
	//  Calculates the available memory in Windows.
	AvailableSystemMemory() (resultUint64 uint64) // property
	// SystemMemoryLoad
	//  Memory load in Windows.
	SystemMemoryLoad() uint32 // property
	// ApiHashUniversal
	//  Calls cef_api_hash to get the universal hash.
	ApiHashUniversal() string // property
	// ApiHashPlatform
	//  Calls cef_api_hash to get the platform hash.
	ApiHashPlatform() string // property
	// ApiHashCommit
	//  Calls cef_api_hash to get the commit hash.
	ApiHashCommit() string // property
	// LastErrorMessage
	//  Last error message that is usually shown when CEF finds a problem at initialization.
	LastErrorMessage() string // property
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
	// AddCustomCommandLine
	//  Used to add any command line switch that is not available as a
	//  TCEFApplicationCore property.
	AddCustomCommandLine(aCommandLine string, aValue string) // procedure
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
	//  TCefSettings.multi_threaded_message_loop value of false(0). This function
	//  will not block.
	DoMessageLoopWork() // procedure
	// RunMessageLoop
	//  Run the CEF message loop. Use this function instead of an application-
	//  provided message loop to get the best balance between performance and CPU
	//  usage. This function should only be called on the main application thread
	//  and only if cef_initialize() is called with a
	//  TCefSettings.multi_threaded_message_loop value of false(0). This function
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
	// SetOnRegCustomSchemes
	//  Provides an opportunity to register custom schemes. Do not keep a
	//  reference to the |registrar| object. This function is called on the main
	//  thread for each process and the registered schemes should be the same
	//  across all processes.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_app_capi.h">CEF source file: /include/capi/cef_app_capi.h(cef_app_t)</a>
	SetOnRegCustomSchemes(fn TOnRegisterCustomSchemes) // property event
	// SetOnRegisterCustomPreferences
	//  Provides an opportunity to register custom preferences prior to global and
	//  request context initialization.
	//  If |type| is CEF_PREFERENCES_TYPE_GLOBAL the registered preferences can be
	//  accessed via ICefPreferenceManager.GetGlobalPreferences after
	//  OnContextInitialized is called. Global preferences are registered a single
	//  time at application startup. See related TCefSettings.cache_path and
	//  TCefSettings.persist_user_preferences configuration.
	//  If |type| is CEF_PREFERENCES_TYPE_REQUEST_CONTEXT the preferences can be
	//  accessed via the ICefRequestContext after
	//  ICefRequestContextHandler.OnRequestContextInitialized is called.
	//  Request context preferences are registered each time a new
	//  ICefRequestContext is created. It is intended but not required that all
	//  request contexts have the same registered preferences. See related
	//  TCefRequestContextSettings.cache_path and
	//  TCefRequestContextSettings.persist_user_preferences configuration.
	//  Do not keep a reference to the |registrar| object. This function is called
	//  on the browser process UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_process_handler_capi.h">CEF source file: /include/capi/cef_browser_process_handler_capi.h(cef_browser_process_handler_t)</a>
	SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferences) // property event
	// SetOnContextInitialized
	//  Called on the browser process UI thread immediately after the CEF context
	//  has been initialized.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_process_handler_capi.h">CEF source file: /include/capi/cef_browser_process_handler_capi.h(cef_browser_process_handler_t)</a>
	SetOnContextInitialized(fn TOnContextInitialized) // property event
	// SetOnBeforeChildProcessLaunch
	//  Called before a child process is launched. Will be called on the browser
	//  process UI thread when launching a render process and on the browser
	//  process IO thread when launching a GPU process. Provides an opportunity to
	//  modify the child process command line. Do not keep a reference to
	//  |command_line| outside of this function.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_process_handler_capi.h">CEF source file: /include/capi/cef_browser_process_handler_capi.h(cef_browser_process_handler_t)</a>
	SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunch) // property event
	// SetOnScheduleMessagePumpWork
	//  Called from any thread when work has been scheduled for the browser
	//  process main(UI) thread. This callback is used in combination with
	//  TCefSettings.external_message_pump and GlobalCEFApp.DoMessageLoopWork in
	//  cases where the CEF message loop must be integrated into an existing
	//  application message loop(see additional comments and warnings on
	//  GlobalCEFApp.DoMessageLoopWork). This callback should schedule a
	//  GlobalCEFApp.DoMessageLoopWork call to happen on the main(UI) thread.
	//  |delay_ms| is the requested delay in milliseconds. If |delay_ms| is <= 0
	//  then the call should happen reasonably soon. If |delay_ms| is > 0 then the
	//  call should be scheduled to happen after the specified delay and any
	//  currently pending scheduled call should be cancelled.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_process_handler_capi.h">CEF source file: /include/capi/cef_browser_process_handler_capi.h(cef_browser_process_handler_t)</a>
	SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWork) // property event
	// SetOnGetDefaultClient
	//  Return the default client for use with a newly created browser window. If
	//  null is returned the browser will be unmanaged(no callbacks will be
	//  executed for that browser) and application shutdown will be blocked until
	//  the browser window is closed manually. This function is currently only
	//  used with the chrome runtime.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_process_handler_capi.h">CEF source file: /include/capi/cef_browser_process_handler_capi.h(cef_browser_process_handler_t)</a>
	SetOnGetDefaultClient(fn TOnGetDefaultClient) // property event
	// SetOnGetLocalizedString
	//  Called to retrieve a localized translation for the specified |string_id|.
	//  To provide the translation set |string| to the translation string and
	//  return true(1). To use the default translation return false(0). Include
	//  cef_pack_strings.h for a listing of valid string ID values.
	//  This event may be called on multiple threads.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_bundle_handler_capi.h">CEF source file: /include/capi/cef_resource_bundle_handler_capi.h(cef_resource_bundle_handler_t)</a>
	SetOnGetLocalizedString(fn TOnGetLocalizedString) // property event
	// SetOnGetDataResource
	//  Called to retrieve data for the specified scale independent |resource_id|.
	//  To provide the resource data set |data| and |data_size| to the data
	//  pointer and size respectively and return true(1). To use the default
	//  resource data return false(0). The resource data will not be copied and
	//  must remain resident in memory. Include cef_pack_resources.h for a listing
	//  of valid resource ID values.
	//  This event may be called on multiple threads.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_bundle_handler_capi.h">CEF source file: /include/capi/cef_resource_bundle_handler_capi.h(cef_resource_bundle_handler_t)</a>
	SetOnGetDataResource(fn TOnGetDataResource) // property event
	// SetOnGetDataResourceForScale
	//  Called to retrieve data for the specified |resource_id| nearest the scale
	//  factor |scale_factor|. To provide the resource data set |data| and
	//  |data_size| to the data pointer and size respectively and return true(1).
	//  To use the default resource data return false(0). The resource data will
	//  not be copied and must remain resident in memory. Include
	//  cef_pack_resources.h for a listing of valid resource ID values.
	//  This event may be called on multiple threads.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_bundle_handler_capi.h">CEF source file: /include/capi/cef_resource_bundle_handler_capi.h(cef_resource_bundle_handler_t)</a>
	SetOnGetDataResourceForScale(fn TOnGetDataResourceForScale) // property event
	// SetOnWebKitInitialized
	//  Called after WebKit has been initialized.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnWebKitInitialized(fn TOnWebKitInitialized) // property event
	// SetOnBrowserCreated
	//  Called after a browser has been created. When browsing cross-origin a new
	//  browser will be created before the old browser with the same identifier is
	//  destroyed. |extra_info| is an optional read-only value originating from
	//  cef_browser_host_create_browser(),
	//  cef_browser_host_create_browser_sync(),
	//  ICefLifeSpanHandler.OnBeforePopup or
	//  cef_browser_view_create().
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnBrowserCreated(fn TOnBrowserCreated) // property event
	// SetOnBrowserDestroyed
	//  Called before a browser is destroyed.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnBrowserDestroyed(fn TOnBrowserDestroyed) // property event
	// SetOnContextCreated
	//  Called immediately after the V8 context for a frame has been created. To
	//  retrieve the JavaScript 'window' object use the
	//  ICefv8context.GetGlobal function. V8 handles can only be accessed
	//  from the thread on which they are created. A task runner for posting tasks
	//  on the associated thread can be retrieved via the
	//  ICefv8context.GetTaskRunner() function.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnContextCreated(fn TOnContextCreated) // property event
	// SetOnContextReleased
	//  Called immediately before the V8 context for a frame is released. No
	//  references to the context should be kept after this function is called.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnContextReleased(fn TOnContextReleased) // property event
	// SetOnUncaughtException
	//  Called for global uncaught exceptions in a frame. Execution of this
	//  callback is disabled by default. To enable set
	//  TCefSettings.uncaught_exception_stack_size > 0.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnUncaughtException(fn TOnUncaughtException) // property event
	// SetOnFocusedNodeChanged
	//  Called when a new node in the the browser gets focus. The |node| value may
	//  be NULL if no specific node has gained focus. The node object passed to
	//  this function represents a snapshot of the DOM at the time this function
	//  is executed. DOM objects are only valid for the scope of this function. Do
	//  not keep references to or attempt to access any DOM objects outside the
	//  scope of this function.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnFocusedNodeChanged(fn TOnFocusedNodeChanged) // property event
	// SetOnProcessMessageReceived
	//  Called when a new message is received from a different process. Return
	//  true(1) if the message was handled or false(0) otherwise. It is safe to
	//  keep a reference to |message| outside of this callback.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h(cef_render_process_handler_t)</a>
	SetOnProcessMessageReceived(fn TOnProcessMessageReceived) // property event
	// SetOnLoadingStateChange
	//  Called when the loading state has changed. This callback will be executed
	//  twice -- once when loading is initiated either programmatically or by user
	//  action, and once when loading is terminated due to completion,
	//  cancellation of failure. It will be called before any calls to OnLoadStart
	//  and after all calls to OnLoadError and/or OnLoadEnd.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadingStateChange(fn TOnRenderLoadingStateChange) // property event
	// SetOnLoadStart
	//  Called after a navigation has been committed and before the browser begins
	//  loading contents in the frame. The |frame| value will never be NULL --
	//  call the IsMain() function to check if this frame is the main frame.
	//  |transition_type| provides information about the source of the navigation
	//  and an accurate value is only available in the browser process. Multiple
	//  frames may be loading at the same time. Sub-frames may start or continue
	//  loading after the main frame load has ended. This function will not be
	//  called for same page navigations(fragments, history state, etc.) or for
	//  navigations that fail or are canceled before commit. For notification of
	//  overall browser load status use OnLoadingStateChange instead.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadStart(fn TOnRenderLoadStart) // property event
	// SetOnLoadEnd
	//  Called when the browser is done loading a frame. The |frame| value will
	//  never be NULL -- call the IsMain() function to check if this frame is the
	//  main frame. Multiple frames may be loading at the same time. Sub-frames
	//  may start or continue loading after the main frame load has ended. This
	//  function will not be called for same page navigations(fragments, history
	//  state, etc.) or for navigations that fail or are canceled before commit.
	//  For notification of overall browser load status use OnLoadingStateChange
	//  instead.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadEnd(fn TOnRenderLoadEnd) // property event
	// SetOnLoadError
	//  Called when a navigation fails or is canceled. This function may be called
	//  by itself if before commit or in combination with OnLoadStart/OnLoadEnd if
	//  after commit. |errorCode| is the error code number, |errorText| is the
	//  error text and |failedUrl| is the URL that failed to load. See
	//  net\base\net_error_list.h for complete descriptions of the error codes.
	//  This event will be called on the render process main thread(TID_RENDERER)
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadError(fn TOnRenderLoadError) // property event
}

// TCefApplicationCore Parent: TObject
//
//	Parent class of TCefApplication used to simplify the CEF initialization and destruction.
type TCefApplicationCore struct {
	TObject
	regCustomSchemesPtr          uintptr
	registerCustomPreferencesPtr uintptr
	contextInitializedPtr        uintptr
	beforeChildProcessLaunchPtr  uintptr
	scheduleMessagePumpWorkPtr   uintptr
	getDefaultClientPtr          uintptr
	getLocalizedStringPtr        uintptr
	getDataResourcePtr           uintptr
	getDataResourceForScalePtr   uintptr
	webKitInitializedPtr         uintptr
	browserCreatedPtr            uintptr
	browserDestroyedPtr          uintptr
	contextCreatedPtr            uintptr
	contextReleasedPtr           uintptr
	uncaughtExceptionPtr         uintptr
	focusedNodeChangedPtr        uintptr
	processMessageReceivedPtr    uintptr
	loadingStateChangePtr        uintptr
	loadStartPtr                 uintptr
	loadEndPtr                   uintptr
	loadErrorPtr                 uintptr
}

func NewCefApplicationCore() ICefApplicationCore {
	r1 := applicationCoreImportAPI().SysCallN(32)
	return AsCefApplicationCore(r1)
}

func (m *TCefApplicationCore) NoSandbox() bool {
	r1 := applicationCoreImportAPI().SysCallN(106, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetNoSandbox(AValue bool) {
	applicationCoreImportAPI().SysCallN(106, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) BrowserSubprocessPath() string {
	r1 := applicationCoreImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetBrowserSubprocessPath(AValue string) {
	applicationCoreImportAPI().SysCallN(16, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) FrameworkDirPath() string {
	r1 := applicationCoreImportAPI().SysCallN(73, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetFrameworkDirPath(AValue string) {
	applicationCoreImportAPI().SysCallN(73, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) MainBundlePath() string {
	r1 := applicationCoreImportAPI().SysCallN(92, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetMainBundlePath(AValue string) {
	applicationCoreImportAPI().SysCallN(92, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ChromeRuntime() bool {
	r1 := applicationCoreImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetChromeRuntime(AValue bool) {
	applicationCoreImportAPI().SysCallN(27, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MultiThreadedMessageLoop() bool {
	r1 := applicationCoreImportAPI().SysCallN(96, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMultiThreadedMessageLoop(AValue bool) {
	applicationCoreImportAPI().SysCallN(96, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ExternalMessagePump() bool {
	r1 := applicationCoreImportAPI().SysCallN(69, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetExternalMessagePump(AValue bool) {
	applicationCoreImportAPI().SysCallN(69, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) WindowlessRenderingEnabled() bool {
	r1 := applicationCoreImportAPI().SysCallN(161, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetWindowlessRenderingEnabled(AValue bool) {
	applicationCoreImportAPI().SysCallN(161, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) CommandLineArgsDisabled() bool {
	r1 := applicationCoreImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetCommandLineArgsDisabled(AValue bool) {
	applicationCoreImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) Cache() string {
	r1 := applicationCoreImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetCache(AValue string) {
	applicationCoreImportAPI().SysCallN(17, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) RootCache() string {
	r1 := applicationCoreImportAPI().SysCallN(118, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetRootCache(AValue string) {
	applicationCoreImportAPI().SysCallN(118, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) PersistSessionCookies() bool {
	r1 := applicationCoreImportAPI().SysCallN(110, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetPersistSessionCookies(AValue bool) {
	applicationCoreImportAPI().SysCallN(110, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) PersistUserPreferences() bool {
	r1 := applicationCoreImportAPI().SysCallN(111, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetPersistUserPreferences(AValue bool) {
	applicationCoreImportAPI().SysCallN(111, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) UserAgent() string {
	r1 := applicationCoreImportAPI().SysCallN(159, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetUserAgent(AValue string) {
	applicationCoreImportAPI().SysCallN(159, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) UserAgentProduct() string {
	r1 := applicationCoreImportAPI().SysCallN(160, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetUserAgentProduct(AValue string) {
	applicationCoreImportAPI().SysCallN(160, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) Locale() string {
	r1 := applicationCoreImportAPI().SysCallN(85, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetLocale(AValue string) {
	applicationCoreImportAPI().SysCallN(85, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) LogFile() string {
	r1 := applicationCoreImportAPI().SysCallN(88, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetLogFile(AValue string) {
	applicationCoreImportAPI().SysCallN(88, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) LogSeverity() TCefLogSeverity {
	r1 := applicationCoreImportAPI().SysCallN(91, 0, m.Instance(), 0)
	return TCefLogSeverity(r1)
}

func (m *TCefApplicationCore) SetLogSeverity(AValue TCefLogSeverity) {
	applicationCoreImportAPI().SysCallN(91, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) LogItems() TCefLogItems {
	r1 := applicationCoreImportAPI().SysCallN(89, 0, m.Instance(), 0)
	return TCefLogItems(r1)
}

func (m *TCefApplicationCore) SetLogItems(AValue TCefLogItems) {
	applicationCoreImportAPI().SysCallN(89, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) JavaScriptFlags() string {
	r1 := applicationCoreImportAPI().SysCallN(79, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetJavaScriptFlags(AValue string) {
	applicationCoreImportAPI().SysCallN(79, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ResourcesDirPath() string {
	r1 := applicationCoreImportAPI().SysCallN(117, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetResourcesDirPath(AValue string) {
	applicationCoreImportAPI().SysCallN(117, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) LocalesDirPath() string {
	r1 := applicationCoreImportAPI().SysCallN(86, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetLocalesDirPath(AValue string) {
	applicationCoreImportAPI().SysCallN(86, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) PackLoadingDisabled() bool {
	r1 := applicationCoreImportAPI().SysCallN(109, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetPackLoadingDisabled(AValue bool) {
	applicationCoreImportAPI().SysCallN(109, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) RemoteDebuggingPort() int32 {
	r1 := applicationCoreImportAPI().SysCallN(116, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCefApplicationCore) SetRemoteDebuggingPort(AValue int32) {
	applicationCoreImportAPI().SysCallN(116, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) UncaughtExceptionStackSize() int32 {
	r1 := applicationCoreImportAPI().SysCallN(154, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCefApplicationCore) SetUncaughtExceptionStackSize(AValue int32) {
	applicationCoreImportAPI().SysCallN(154, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) BackgroundColor() TCefColor {
	r1 := applicationCoreImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCefApplicationCore) SetBackgroundColor(AValue TCefColor) {
	applicationCoreImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) AcceptLanguageList() string {
	r1 := applicationCoreImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetAcceptLanguageList(AValue string) {
	applicationCoreImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) CookieableSchemesList() string {
	r1 := applicationCoreImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetCookieableSchemesList(AValue string) {
	applicationCoreImportAPI().SysCallN(31, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) CookieableSchemesExcludeDefaults() bool {
	r1 := applicationCoreImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetCookieableSchemesExcludeDefaults(AValue bool) {
	applicationCoreImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ChromePolicyId() string {
	r1 := applicationCoreImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetChromePolicyId(AValue string) {
	applicationCoreImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) SingleProcess() bool {
	r1 := applicationCoreImportAPI().SysCallN(143, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetSingleProcess(AValue bool) {
	applicationCoreImportAPI().SysCallN(143, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableMediaStream() bool {
	r1 := applicationCoreImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnableMediaStream(AValue bool) {
	applicationCoreImportAPI().SysCallN(64, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableSpeechInput() bool {
	r1 := applicationCoreImportAPI().SysCallN(67, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnableSpeechInput(AValue bool) {
	applicationCoreImportAPI().SysCallN(67, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) UseFakeUIForMediaStream() bool {
	r1 := applicationCoreImportAPI().SysCallN(156, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetUseFakeUIForMediaStream(AValue bool) {
	applicationCoreImportAPI().SysCallN(156, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableUsermediaScreenCapturing() bool {
	r1 := applicationCoreImportAPI().SysCallN(68, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnableUsermediaScreenCapturing(AValue bool) {
	applicationCoreImportAPI().SysCallN(68, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableGPU() bool {
	r1 := applicationCoreImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnableGPU(AValue bool) {
	applicationCoreImportAPI().SysCallN(63, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableFeatures() string {
	r1 := applicationCoreImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetEnableFeatures(AValue string) {
	applicationCoreImportAPI().SysCallN(62, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) DisableFeatures() string {
	r1 := applicationCoreImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetDisableFeatures(AValue string) {
	applicationCoreImportAPI().SysCallN(43, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) EnableBlinkFeatures() string {
	r1 := applicationCoreImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetEnableBlinkFeatures(AValue string) {
	applicationCoreImportAPI().SysCallN(61, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) DisableBlinkFeatures() string {
	r1 := applicationCoreImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetDisableBlinkFeatures(AValue string) {
	applicationCoreImportAPI().SysCallN(39, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) BlinkSettings() string {
	r1 := applicationCoreImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetBlinkSettings(AValue string) {
	applicationCoreImportAPI().SysCallN(15, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ForceFieldTrials() string {
	r1 := applicationCoreImportAPI().SysCallN(71, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetForceFieldTrials(AValue string) {
	applicationCoreImportAPI().SysCallN(71, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ForceFieldTrialParams() string {
	r1 := applicationCoreImportAPI().SysCallN(70, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetForceFieldTrialParams(AValue string) {
	applicationCoreImportAPI().SysCallN(70, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) SmoothScrolling() TCefState {
	r1 := applicationCoreImportAPI().SysCallN(145, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TCefApplicationCore) SetSmoothScrolling(AValue TCefState) {
	applicationCoreImportAPI().SysCallN(145, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) MuteAudio() bool {
	r1 := applicationCoreImportAPI().SysCallN(102, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMuteAudio(AValue bool) {
	applicationCoreImportAPI().SysCallN(102, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) SitePerProcess() bool {
	r1 := applicationCoreImportAPI().SysCallN(144, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetSitePerProcess(AValue bool) {
	applicationCoreImportAPI().SysCallN(144, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableWebSecurity() bool {
	r1 := applicationCoreImportAPI().SysCallN(58, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableWebSecurity(AValue bool) {
	applicationCoreImportAPI().SysCallN(58, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisablePDFExtension() bool {
	r1 := applicationCoreImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisablePDFExtension(AValue bool) {
	applicationCoreImportAPI().SysCallN(50, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableSiteIsolationTrials() bool {
	r1 := applicationCoreImportAPI().SysCallN(54, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableSiteIsolationTrials(AValue bool) {
	applicationCoreImportAPI().SysCallN(54, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableChromeLoginPrompt() bool {
	r1 := applicationCoreImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableChromeLoginPrompt(AValue bool) {
	applicationCoreImportAPI().SysCallN(40, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableExtensions() bool {
	r1 := applicationCoreImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableExtensions(AValue bool) {
	applicationCoreImportAPI().SysCallN(42, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) AutoplayPolicy() TCefAutoplayPolicy {
	r1 := applicationCoreImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TCefAutoplayPolicy(r1)
}

func (m *TCefApplicationCore) SetAutoplayPolicy(AValue TCefAutoplayPolicy) {
	applicationCoreImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) DisableBackgroundNetworking() bool {
	r1 := applicationCoreImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableBackgroundNetworking(AValue bool) {
	applicationCoreImportAPI().SysCallN(38, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MetricsRecordingOnly() bool {
	r1 := applicationCoreImportAPI().SysCallN(93, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMetricsRecordingOnly(AValue bool) {
	applicationCoreImportAPI().SysCallN(93, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) AllowFileAccessFromFiles() bool {
	r1 := applicationCoreImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetAllowFileAccessFromFiles(AValue bool) {
	applicationCoreImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) AllowRunningInsecureContent() bool {
	r1 := applicationCoreImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetAllowRunningInsecureContent(AValue bool) {
	applicationCoreImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnablePrintPreview() bool {
	r1 := applicationCoreImportAPI().SysCallN(65, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnablePrintPreview(AValue bool) {
	applicationCoreImportAPI().SysCallN(65, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableJavascript() bool {
	r1 := applicationCoreImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableJavascript(AValue bool) {
	applicationCoreImportAPI().SysCallN(45, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableJavascriptCloseWindows() bool {
	r1 := applicationCoreImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableJavascriptCloseWindows(AValue bool) {
	applicationCoreImportAPI().SysCallN(47, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableJavascriptAccessClipboard() bool {
	r1 := applicationCoreImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableJavascriptAccessClipboard(AValue bool) {
	applicationCoreImportAPI().SysCallN(46, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableJavascriptDomPaste() bool {
	r1 := applicationCoreImportAPI().SysCallN(48, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableJavascriptDomPaste(AValue bool) {
	applicationCoreImportAPI().SysCallN(48, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) AllowUniversalAccessFromFileUrls() bool {
	r1 := applicationCoreImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetAllowUniversalAccessFromFileUrls(AValue bool) {
	applicationCoreImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableImageLoading() bool {
	r1 := applicationCoreImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableImageLoading(AValue bool) {
	applicationCoreImportAPI().SysCallN(44, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ImageShrinkStandaloneToFit() bool {
	r1 := applicationCoreImportAPI().SysCallN(77, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetImageShrinkStandaloneToFit(AValue bool) {
	applicationCoreImportAPI().SysCallN(77, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableTextAreaResize() bool {
	r1 := applicationCoreImportAPI().SysCallN(57, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableTextAreaResize(AValue bool) {
	applicationCoreImportAPI().SysCallN(57, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableTabToLinks() bool {
	r1 := applicationCoreImportAPI().SysCallN(56, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableTabToLinks(AValue bool) {
	applicationCoreImportAPI().SysCallN(56, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) EnableProfanityFilter() bool {
	r1 := applicationCoreImportAPI().SysCallN(66, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetEnableProfanityFilter(AValue bool) {
	applicationCoreImportAPI().SysCallN(66, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableSpellChecking() bool {
	r1 := applicationCoreImportAPI().SysCallN(55, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableSpellChecking(AValue bool) {
	applicationCoreImportAPI().SysCallN(55, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) OverrideSpellCheckLang() string {
	r1 := applicationCoreImportAPI().SysCallN(108, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetOverrideSpellCheckLang(AValue string) {
	applicationCoreImportAPI().SysCallN(108, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) TouchEvents() TCefState {
	r1 := applicationCoreImportAPI().SysCallN(151, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TCefApplicationCore) SetTouchEvents(AValue TCefState) {
	applicationCoreImportAPI().SysCallN(151, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) DisableReadingFromCanvas() bool {
	r1 := applicationCoreImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableReadingFromCanvas(AValue bool) {
	applicationCoreImportAPI().SysCallN(52, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) HyperlinkAuditing() bool {
	r1 := applicationCoreImportAPI().SysCallN(75, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetHyperlinkAuditing(AValue bool) {
	applicationCoreImportAPI().SysCallN(75, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableNewBrowserInfoTimeout() bool {
	r1 := applicationCoreImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableNewBrowserInfoTimeout(AValue bool) {
	applicationCoreImportAPI().SysCallN(49, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DevToolsProtocolLogFile() string {
	r1 := applicationCoreImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetDevToolsProtocolLogFile(AValue string) {
	applicationCoreImportAPI().SysCallN(35, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ForcedDeviceScaleFactor() (resultFloat32 float32) {
	applicationCoreImportAPI().SysCallN(72, 0, m.Instance(), uintptr(unsafePointer(&resultFloat32)), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TCefApplicationCore) SetForcedDeviceScaleFactor(AValue float32) {
	applicationCoreImportAPI().SysCallN(72, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCefApplicationCore) DisableZygote() bool {
	r1 := applicationCoreImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableZygote(AValue bool) {
	applicationCoreImportAPI().SysCallN(59, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) UseMockKeyChain() bool {
	r1 := applicationCoreImportAPI().SysCallN(157, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetUseMockKeyChain(AValue bool) {
	applicationCoreImportAPI().SysCallN(157, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableRequestHandlingForTesting() bool {
	r1 := applicationCoreImportAPI().SysCallN(53, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableRequestHandlingForTesting(AValue bool) {
	applicationCoreImportAPI().SysCallN(53, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisablePopupBlocking() bool {
	r1 := applicationCoreImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisablePopupBlocking(AValue bool) {
	applicationCoreImportAPI().SysCallN(51, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableBackForwardCache() bool {
	r1 := applicationCoreImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableBackForwardCache(AValue bool) {
	applicationCoreImportAPI().SysCallN(37, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DisableComponentUpdate() bool {
	r1 := applicationCoreImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDisableComponentUpdate(AValue bool) {
	applicationCoreImportAPI().SysCallN(41, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) AllowInsecureLocalhost() bool {
	r1 := applicationCoreImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetAllowInsecureLocalhost(AValue bool) {
	applicationCoreImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) KioskPrinting() bool {
	r1 := applicationCoreImportAPI().SysCallN(80, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetKioskPrinting(AValue bool) {
	applicationCoreImportAPI().SysCallN(80, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) TreatInsecureOriginAsSecure() string {
	r1 := applicationCoreImportAPI().SysCallN(152, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetTreatInsecureOriginAsSecure(AValue string) {
	applicationCoreImportAPI().SysCallN(152, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) NetLogEnabled() bool {
	r1 := applicationCoreImportAPI().SysCallN(104, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetNetLogEnabled(AValue bool) {
	applicationCoreImportAPI().SysCallN(104, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) NetLogFile() string {
	r1 := applicationCoreImportAPI().SysCallN(105, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetNetLogFile(AValue string) {
	applicationCoreImportAPI().SysCallN(105, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) NetLogCaptureMode() TCefNetLogCaptureMode {
	r1 := applicationCoreImportAPI().SysCallN(103, 0, m.Instance(), 0)
	return TCefNetLogCaptureMode(r1)
}

func (m *TCefApplicationCore) SetNetLogCaptureMode(AValue TCefNetLogCaptureMode) {
	applicationCoreImportAPI().SysCallN(103, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) RemoteAllowOrigins() string {
	r1 := applicationCoreImportAPI().SysCallN(115, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetRemoteAllowOrigins(AValue string) {
	applicationCoreImportAPI().SysCallN(115, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) AutoAcceptCamAndMicCapture() bool {
	r1 := applicationCoreImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetAutoAcceptCamAndMicCapture(AValue bool) {
	applicationCoreImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) UIColorMode() TCefUIColorMode {
	r1 := applicationCoreImportAPI().SysCallN(153, 0, m.Instance(), 0)
	return TCefUIColorMode(r1)
}

func (m *TCefApplicationCore) SetUIColorMode(AValue TCefUIColorMode) {
	applicationCoreImportAPI().SysCallN(153, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) IgnoreCertificateErrors() bool {
	r1 := applicationCoreImportAPI().SysCallN(76, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetIgnoreCertificateErrors(AValue bool) {
	applicationCoreImportAPI().SysCallN(76, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) WindowsSandboxInfo() uintptr {
	r1 := applicationCoreImportAPI().SysCallN(162, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TCefApplicationCore) SetWindowsSandboxInfo(AValue uintptr) {
	applicationCoreImportAPI().SysCallN(162, 1, m.Instance(), uintptr(AValue))
}

func (m *TCefApplicationCore) ArgcCopy() int32 {
	r1 := applicationCoreImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCefApplicationCore) ArgvCopy() PString {
	r1 := applicationCoreImportAPI().SysCallN(10, m.Instance())
	return PString(r1)
}

func (m *TCefApplicationCore) DeleteCache() bool {
	r1 := applicationCoreImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDeleteCache(AValue bool) {
	applicationCoreImportAPI().SysCallN(33, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DeleteCookies() bool {
	r1 := applicationCoreImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetDeleteCookies(AValue bool) {
	applicationCoreImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) CheckCEFFiles() bool {
	r1 := applicationCoreImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetCheckCEFFiles(AValue bool) {
	applicationCoreImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ShowMessageDlg() bool {
	r1 := applicationCoreImportAPI().SysCallN(142, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetShowMessageDlg(AValue bool) {
	applicationCoreImportAPI().SysCallN(142, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MissingBinariesException() bool {
	r1 := applicationCoreImportAPI().SysCallN(94, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMissingBinariesException(AValue bool) {
	applicationCoreImportAPI().SysCallN(94, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) SetCurrentDir() bool {
	r1 := applicationCoreImportAPI().SysCallN(120, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetSetCurrentDir(AValue bool) {
	applicationCoreImportAPI().SysCallN(120, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) GlobalContextInitialized() bool {
	r1 := applicationCoreImportAPI().SysCallN(74, m.Instance())
	return GoBool(r1)
}

func (m *TCefApplicationCore) ChromeMajorVer() uint16 {
	r1 := applicationCoreImportAPI().SysCallN(23, m.Instance())
	return uint16(r1)
}

func (m *TCefApplicationCore) ChromeMinorVer() uint16 {
	r1 := applicationCoreImportAPI().SysCallN(24, m.Instance())
	return uint16(r1)
}

func (m *TCefApplicationCore) ChromeRelease() uint16 {
	r1 := applicationCoreImportAPI().SysCallN(26, m.Instance())
	return uint16(r1)
}

func (m *TCefApplicationCore) ChromeBuild() uint16 {
	r1 := applicationCoreImportAPI().SysCallN(21, m.Instance())
	return uint16(r1)
}

func (m *TCefApplicationCore) ChromeVersion() string {
	r1 := applicationCoreImportAPI().SysCallN(28, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) LibCefVersion() string {
	r1 := applicationCoreImportAPI().SysCallN(83, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) LibCefPath() string {
	r1 := applicationCoreImportAPI().SysCallN(82, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) ChromeElfPath() string {
	r1 := applicationCoreImportAPI().SysCallN(22, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) LibLoaded() bool {
	r1 := applicationCoreImportAPI().SysCallN(84, m.Instance())
	return GoBool(r1)
}

func (m *TCefApplicationCore) LogProcessInfo() bool {
	r1 := applicationCoreImportAPI().SysCallN(90, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetLogProcessInfo(AValue bool) {
	applicationCoreImportAPI().SysCallN(90, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ReRaiseExceptions() bool {
	r1 := applicationCoreImportAPI().SysCallN(114, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetReRaiseExceptions(AValue bool) {
	applicationCoreImportAPI().SysCallN(114, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) DeviceScaleFactor() (resultFloat32 float32) {
	applicationCoreImportAPI().SysCallN(36, m.Instance(), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TCefApplicationCore) LocalesRequired() string {
	r1 := applicationCoreImportAPI().SysCallN(87, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCefApplicationCore) SetLocalesRequired(AValue string) {
	applicationCoreImportAPI().SysCallN(87, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCefApplicationCore) ProcessType() TCefProcessType {
	r1 := applicationCoreImportAPI().SysCallN(112, m.Instance())
	return TCefProcessType(r1)
}

func (m *TCefApplicationCore) MustCreateResourceBundleHandler() bool {
	r1 := applicationCoreImportAPI().SysCallN(100, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMustCreateResourceBundleHandler(AValue bool) {
	applicationCoreImportAPI().SysCallN(100, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MustCreateBrowserProcessHandler() bool {
	r1 := applicationCoreImportAPI().SysCallN(97, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMustCreateBrowserProcessHandler(AValue bool) {
	applicationCoreImportAPI().SysCallN(97, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MustCreateRenderProcessHandler() bool {
	r1 := applicationCoreImportAPI().SysCallN(99, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMustCreateRenderProcessHandler(AValue bool) {
	applicationCoreImportAPI().SysCallN(99, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) MustCreateLoadHandler() bool {
	r1 := applicationCoreImportAPI().SysCallN(98, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMustCreateLoadHandler(AValue bool) {
	applicationCoreImportAPI().SysCallN(98, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) SetOsmodalLoop(AValue bool) {
	applicationCoreImportAPI().SysCallN(107, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) Status() TCefAplicationStatus {
	r1 := applicationCoreImportAPI().SysCallN(148, m.Instance())
	return TCefAplicationStatus(r1)
}

func (m *TCefApplicationCore) MissingLibFiles() string {
	r1 := applicationCoreImportAPI().SysCallN(95, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) MustFreeLibrary() bool {
	r1 := applicationCoreImportAPI().SysCallN(101, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplicationCore) SetMustFreeLibrary(AValue bool) {
	applicationCoreImportAPI().SysCallN(101, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplicationCore) ChildProcessesCount() int32 {
	r1 := applicationCoreImportAPI().SysCallN(20, m.Instance())
	return int32(r1)
}

func (m *TCefApplicationCore) UsedMemory() (resultUint64 uint64) {
	applicationCoreImportAPI().SysCallN(158, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCefApplicationCore) TotalSystemMemory() (resultUint64 uint64) {
	applicationCoreImportAPI().SysCallN(150, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCefApplicationCore) AvailableSystemMemory() (resultUint64 uint64) {
	applicationCoreImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCefApplicationCore) SystemMemoryLoad() uint32 {
	r1 := applicationCoreImportAPI().SysCallN(149, m.Instance())
	return uint32(r1)
}

func (m *TCefApplicationCore) ApiHashUniversal() string {
	r1 := applicationCoreImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) ApiHashPlatform() string {
	r1 := applicationCoreImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) ApiHashCommit() string {
	r1 := applicationCoreImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) LastErrorMessage() string {
	r1 := applicationCoreImportAPI().SysCallN(81, m.Instance())
	return GoStr(r1)
}

func (m *TCefApplicationCore) CheckCEFLibrary() bool {
	r1 := applicationCoreImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCefApplicationCore) StartMainProcess() bool {
	r1 := applicationCoreImportAPI().SysCallN(146, m.Instance())
	return GoBool(r1)
}

func (m *TCefApplicationCore) StartSubProcess() bool {
	r1 := applicationCoreImportAPI().SysCallN(147, m.Instance())
	return GoBool(r1)
}

func (m *TCefApplicationCore) AddCustomCommandLine(aCommandLine string, aValue string) {
	applicationCoreImportAPI().SysCallN(1, m.Instance(), PascalStr(aCommandLine), PascalStr(aValue))
}

func (m *TCefApplicationCore) DoMessageLoopWork() {
	applicationCoreImportAPI().SysCallN(60, m.Instance())
}

func (m *TCefApplicationCore) RunMessageLoop() {
	applicationCoreImportAPI().SysCallN(119, m.Instance())
}

func (m *TCefApplicationCore) QuitMessageLoop() {
	applicationCoreImportAPI().SysCallN(113, m.Instance())
}

func (m *TCefApplicationCore) UpdateDeviceScaleFactor() {
	applicationCoreImportAPI().SysCallN(155, m.Instance())
}

func (m *TCefApplicationCore) InitLibLocationFromArgs() {
	applicationCoreImportAPI().SysCallN(78, m.Instance())
}

func (m *TCefApplicationCore) SetOnRegCustomSchemes(fn TOnRegisterCustomSchemes) {
	if m.regCustomSchemesPtr != 0 {
		RemoveEventElement(m.regCustomSchemesPtr)
	}
	m.regCustomSchemesPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(137, m.Instance(), m.regCustomSchemesPtr)
}

func (m *TCefApplicationCore) SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferences) {
	if m.registerCustomPreferencesPtr != 0 {
		RemoveEventElement(m.registerCustomPreferencesPtr)
	}
	m.registerCustomPreferencesPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(138, m.Instance(), m.registerCustomPreferencesPtr)
}

func (m *TCefApplicationCore) SetOnContextInitialized(fn TOnContextInitialized) {
	if m.contextInitializedPtr != 0 {
		RemoveEventElement(m.contextInitializedPtr)
	}
	m.contextInitializedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(125, m.Instance(), m.contextInitializedPtr)
}

func (m *TCefApplicationCore) SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunch) {
	if m.beforeChildProcessLaunchPtr != 0 {
		RemoveEventElement(m.beforeChildProcessLaunchPtr)
	}
	m.beforeChildProcessLaunchPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(121, m.Instance(), m.beforeChildProcessLaunchPtr)
}

func (m *TCefApplicationCore) SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWork) {
	if m.scheduleMessagePumpWorkPtr != 0 {
		RemoveEventElement(m.scheduleMessagePumpWorkPtr)
	}
	m.scheduleMessagePumpWorkPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(139, m.Instance(), m.scheduleMessagePumpWorkPtr)
}

func (m *TCefApplicationCore) SetOnGetDefaultClient(fn TOnGetDefaultClient) {
	if m.getDefaultClientPtr != 0 {
		RemoveEventElement(m.getDefaultClientPtr)
	}
	m.getDefaultClientPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(130, m.Instance(), m.getDefaultClientPtr)
}

func (m *TCefApplicationCore) SetOnGetLocalizedString(fn TOnGetLocalizedString) {
	if m.getLocalizedStringPtr != 0 {
		RemoveEventElement(m.getLocalizedStringPtr)
	}
	m.getLocalizedStringPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(131, m.Instance(), m.getLocalizedStringPtr)
}

func (m *TCefApplicationCore) SetOnGetDataResource(fn TOnGetDataResource) {
	if m.getDataResourcePtr != 0 {
		RemoveEventElement(m.getDataResourcePtr)
	}
	m.getDataResourcePtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(128, m.Instance(), m.getDataResourcePtr)
}

func (m *TCefApplicationCore) SetOnGetDataResourceForScale(fn TOnGetDataResourceForScale) {
	if m.getDataResourceForScalePtr != 0 {
		RemoveEventElement(m.getDataResourceForScalePtr)
	}
	m.getDataResourceForScalePtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(129, m.Instance(), m.getDataResourceForScalePtr)
}

func (m *TCefApplicationCore) SetOnWebKitInitialized(fn TOnWebKitInitialized) {
	if m.webKitInitializedPtr != 0 {
		RemoveEventElement(m.webKitInitializedPtr)
	}
	m.webKitInitializedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(141, m.Instance(), m.webKitInitializedPtr)
}

func (m *TCefApplicationCore) SetOnBrowserCreated(fn TOnBrowserCreated) {
	if m.browserCreatedPtr != 0 {
		RemoveEventElement(m.browserCreatedPtr)
	}
	m.browserCreatedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(122, m.Instance(), m.browserCreatedPtr)
}

func (m *TCefApplicationCore) SetOnBrowserDestroyed(fn TOnBrowserDestroyed) {
	if m.browserDestroyedPtr != 0 {
		RemoveEventElement(m.browserDestroyedPtr)
	}
	m.browserDestroyedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(123, m.Instance(), m.browserDestroyedPtr)
}

func (m *TCefApplicationCore) SetOnContextCreated(fn TOnContextCreated) {
	if m.contextCreatedPtr != 0 {
		RemoveEventElement(m.contextCreatedPtr)
	}
	m.contextCreatedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(124, m.Instance(), m.contextCreatedPtr)
}

func (m *TCefApplicationCore) SetOnContextReleased(fn TOnContextReleased) {
	if m.contextReleasedPtr != 0 {
		RemoveEventElement(m.contextReleasedPtr)
	}
	m.contextReleasedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(126, m.Instance(), m.contextReleasedPtr)
}

func (m *TCefApplicationCore) SetOnUncaughtException(fn TOnUncaughtException) {
	if m.uncaughtExceptionPtr != 0 {
		RemoveEventElement(m.uncaughtExceptionPtr)
	}
	m.uncaughtExceptionPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(140, m.Instance(), m.uncaughtExceptionPtr)
}

func (m *TCefApplicationCore) SetOnFocusedNodeChanged(fn TOnFocusedNodeChanged) {
	if m.focusedNodeChangedPtr != 0 {
		RemoveEventElement(m.focusedNodeChangedPtr)
	}
	m.focusedNodeChangedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(127, m.Instance(), m.focusedNodeChangedPtr)
}

func (m *TCefApplicationCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceived) {
	if m.processMessageReceivedPtr != 0 {
		RemoveEventElement(m.processMessageReceivedPtr)
	}
	m.processMessageReceivedPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(136, m.Instance(), m.processMessageReceivedPtr)
}

func (m *TCefApplicationCore) SetOnLoadingStateChange(fn TOnRenderLoadingStateChange) {
	if m.loadingStateChangePtr != 0 {
		RemoveEventElement(m.loadingStateChangePtr)
	}
	m.loadingStateChangePtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(135, m.Instance(), m.loadingStateChangePtr)
}

func (m *TCefApplicationCore) SetOnLoadStart(fn TOnRenderLoadStart) {
	if m.loadStartPtr != 0 {
		RemoveEventElement(m.loadStartPtr)
	}
	m.loadStartPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(134, m.Instance(), m.loadStartPtr)
}

func (m *TCefApplicationCore) SetOnLoadEnd(fn TOnRenderLoadEnd) {
	if m.loadEndPtr != 0 {
		RemoveEventElement(m.loadEndPtr)
	}
	m.loadEndPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(132, m.Instance(), m.loadEndPtr)
}

func (m *TCefApplicationCore) SetOnLoadError(fn TOnRenderLoadError) {
	if m.loadErrorPtr != 0 {
		RemoveEventElement(m.loadErrorPtr)
	}
	m.loadErrorPtr = MakeEventDataPtr(fn)
	applicationCoreImportAPI().SysCallN(133, m.Instance(), m.loadErrorPtr)
}

var (
	applicationCoreImport       *imports.Imports = nil
	applicationCoreImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefApplicationCore_AcceptLanguageList", 0),
		/*1*/ imports.NewTable("CefApplicationCore_AddCustomCommandLine", 0),
		/*2*/ imports.NewTable("CefApplicationCore_AllowFileAccessFromFiles", 0),
		/*3*/ imports.NewTable("CefApplicationCore_AllowInsecureLocalhost", 0),
		/*4*/ imports.NewTable("CefApplicationCore_AllowRunningInsecureContent", 0),
		/*5*/ imports.NewTable("CefApplicationCore_AllowUniversalAccessFromFileUrls", 0),
		/*6*/ imports.NewTable("CefApplicationCore_ApiHashCommit", 0),
		/*7*/ imports.NewTable("CefApplicationCore_ApiHashPlatform", 0),
		/*8*/ imports.NewTable("CefApplicationCore_ApiHashUniversal", 0),
		/*9*/ imports.NewTable("CefApplicationCore_ArgcCopy", 0),
		/*10*/ imports.NewTable("CefApplicationCore_ArgvCopy", 0),
		/*11*/ imports.NewTable("CefApplicationCore_AutoAcceptCamAndMicCapture", 0),
		/*12*/ imports.NewTable("CefApplicationCore_AutoplayPolicy", 0),
		/*13*/ imports.NewTable("CefApplicationCore_AvailableSystemMemory", 0),
		/*14*/ imports.NewTable("CefApplicationCore_BackgroundColor", 0),
		/*15*/ imports.NewTable("CefApplicationCore_BlinkSettings", 0),
		/*16*/ imports.NewTable("CefApplicationCore_BrowserSubprocessPath", 0),
		/*17*/ imports.NewTable("CefApplicationCore_Cache", 0),
		/*18*/ imports.NewTable("CefApplicationCore_CheckCEFFiles", 0),
		/*19*/ imports.NewTable("CefApplicationCore_CheckCEFLibrary", 0),
		/*20*/ imports.NewTable("CefApplicationCore_ChildProcessesCount", 0),
		/*21*/ imports.NewTable("CefApplicationCore_ChromeBuild", 0),
		/*22*/ imports.NewTable("CefApplicationCore_ChromeElfPath", 0),
		/*23*/ imports.NewTable("CefApplicationCore_ChromeMajorVer", 0),
		/*24*/ imports.NewTable("CefApplicationCore_ChromeMinorVer", 0),
		/*25*/ imports.NewTable("CefApplicationCore_ChromePolicyId", 0),
		/*26*/ imports.NewTable("CefApplicationCore_ChromeRelease", 0),
		/*27*/ imports.NewTable("CefApplicationCore_ChromeRuntime", 0),
		/*28*/ imports.NewTable("CefApplicationCore_ChromeVersion", 0),
		/*29*/ imports.NewTable("CefApplicationCore_CommandLineArgsDisabled", 0),
		/*30*/ imports.NewTable("CefApplicationCore_CookieableSchemesExcludeDefaults", 0),
		/*31*/ imports.NewTable("CefApplicationCore_CookieableSchemesList", 0),
		/*32*/ imports.NewTable("CefApplicationCore_Create", 0),
		/*33*/ imports.NewTable("CefApplicationCore_DeleteCache", 0),
		/*34*/ imports.NewTable("CefApplicationCore_DeleteCookies", 0),
		/*35*/ imports.NewTable("CefApplicationCore_DevToolsProtocolLogFile", 0),
		/*36*/ imports.NewTable("CefApplicationCore_DeviceScaleFactor", 0),
		/*37*/ imports.NewTable("CefApplicationCore_DisableBackForwardCache", 0),
		/*38*/ imports.NewTable("CefApplicationCore_DisableBackgroundNetworking", 0),
		/*39*/ imports.NewTable("CefApplicationCore_DisableBlinkFeatures", 0),
		/*40*/ imports.NewTable("CefApplicationCore_DisableChromeLoginPrompt", 0),
		/*41*/ imports.NewTable("CefApplicationCore_DisableComponentUpdate", 0),
		/*42*/ imports.NewTable("CefApplicationCore_DisableExtensions", 0),
		/*43*/ imports.NewTable("CefApplicationCore_DisableFeatures", 0),
		/*44*/ imports.NewTable("CefApplicationCore_DisableImageLoading", 0),
		/*45*/ imports.NewTable("CefApplicationCore_DisableJavascript", 0),
		/*46*/ imports.NewTable("CefApplicationCore_DisableJavascriptAccessClipboard", 0),
		/*47*/ imports.NewTable("CefApplicationCore_DisableJavascriptCloseWindows", 0),
		/*48*/ imports.NewTable("CefApplicationCore_DisableJavascriptDomPaste", 0),
		/*49*/ imports.NewTable("CefApplicationCore_DisableNewBrowserInfoTimeout", 0),
		/*50*/ imports.NewTable("CefApplicationCore_DisablePDFExtension", 0),
		/*51*/ imports.NewTable("CefApplicationCore_DisablePopupBlocking", 0),
		/*52*/ imports.NewTable("CefApplicationCore_DisableReadingFromCanvas", 0),
		/*53*/ imports.NewTable("CefApplicationCore_DisableRequestHandlingForTesting", 0),
		/*54*/ imports.NewTable("CefApplicationCore_DisableSiteIsolationTrials", 0),
		/*55*/ imports.NewTable("CefApplicationCore_DisableSpellChecking", 0),
		/*56*/ imports.NewTable("CefApplicationCore_DisableTabToLinks", 0),
		/*57*/ imports.NewTable("CefApplicationCore_DisableTextAreaResize", 0),
		/*58*/ imports.NewTable("CefApplicationCore_DisableWebSecurity", 0),
		/*59*/ imports.NewTable("CefApplicationCore_DisableZygote", 0),
		/*60*/ imports.NewTable("CefApplicationCore_DoMessageLoopWork", 0),
		/*61*/ imports.NewTable("CefApplicationCore_EnableBlinkFeatures", 0),
		/*62*/ imports.NewTable("CefApplicationCore_EnableFeatures", 0),
		/*63*/ imports.NewTable("CefApplicationCore_EnableGPU", 0),
		/*64*/ imports.NewTable("CefApplicationCore_EnableMediaStream", 0),
		/*65*/ imports.NewTable("CefApplicationCore_EnablePrintPreview", 0),
		/*66*/ imports.NewTable("CefApplicationCore_EnableProfanityFilter", 0),
		/*67*/ imports.NewTable("CefApplicationCore_EnableSpeechInput", 0),
		/*68*/ imports.NewTable("CefApplicationCore_EnableUsermediaScreenCapturing", 0),
		/*69*/ imports.NewTable("CefApplicationCore_ExternalMessagePump", 0),
		/*70*/ imports.NewTable("CefApplicationCore_ForceFieldTrialParams", 0),
		/*71*/ imports.NewTable("CefApplicationCore_ForceFieldTrials", 0),
		/*72*/ imports.NewTable("CefApplicationCore_ForcedDeviceScaleFactor", 0),
		/*73*/ imports.NewTable("CefApplicationCore_FrameworkDirPath", 0),
		/*74*/ imports.NewTable("CefApplicationCore_GlobalContextInitialized", 0),
		/*75*/ imports.NewTable("CefApplicationCore_HyperlinkAuditing", 0),
		/*76*/ imports.NewTable("CefApplicationCore_IgnoreCertificateErrors", 0),
		/*77*/ imports.NewTable("CefApplicationCore_ImageShrinkStandaloneToFit", 0),
		/*78*/ imports.NewTable("CefApplicationCore_InitLibLocationFromArgs", 0),
		/*79*/ imports.NewTable("CefApplicationCore_JavaScriptFlags", 0),
		/*80*/ imports.NewTable("CefApplicationCore_KioskPrinting", 0),
		/*81*/ imports.NewTable("CefApplicationCore_LastErrorMessage", 0),
		/*82*/ imports.NewTable("CefApplicationCore_LibCefPath", 0),
		/*83*/ imports.NewTable("CefApplicationCore_LibCefVersion", 0),
		/*84*/ imports.NewTable("CefApplicationCore_LibLoaded", 0),
		/*85*/ imports.NewTable("CefApplicationCore_Locale", 0),
		/*86*/ imports.NewTable("CefApplicationCore_LocalesDirPath", 0),
		/*87*/ imports.NewTable("CefApplicationCore_LocalesRequired", 0),
		/*88*/ imports.NewTable("CefApplicationCore_LogFile", 0),
		/*89*/ imports.NewTable("CefApplicationCore_LogItems", 0),
		/*90*/ imports.NewTable("CefApplicationCore_LogProcessInfo", 0),
		/*91*/ imports.NewTable("CefApplicationCore_LogSeverity", 0),
		/*92*/ imports.NewTable("CefApplicationCore_MainBundlePath", 0),
		/*93*/ imports.NewTable("CefApplicationCore_MetricsRecordingOnly", 0),
		/*94*/ imports.NewTable("CefApplicationCore_MissingBinariesException", 0),
		/*95*/ imports.NewTable("CefApplicationCore_MissingLibFiles", 0),
		/*96*/ imports.NewTable("CefApplicationCore_MultiThreadedMessageLoop", 0),
		/*97*/ imports.NewTable("CefApplicationCore_MustCreateBrowserProcessHandler", 0),
		/*98*/ imports.NewTable("CefApplicationCore_MustCreateLoadHandler", 0),
		/*99*/ imports.NewTable("CefApplicationCore_MustCreateRenderProcessHandler", 0),
		/*100*/ imports.NewTable("CefApplicationCore_MustCreateResourceBundleHandler", 0),
		/*101*/ imports.NewTable("CefApplicationCore_MustFreeLibrary", 0),
		/*102*/ imports.NewTable("CefApplicationCore_MuteAudio", 0),
		/*103*/ imports.NewTable("CefApplicationCore_NetLogCaptureMode", 0),
		/*104*/ imports.NewTable("CefApplicationCore_NetLogEnabled", 0),
		/*105*/ imports.NewTable("CefApplicationCore_NetLogFile", 0),
		/*106*/ imports.NewTable("CefApplicationCore_NoSandbox", 0),
		/*107*/ imports.NewTable("CefApplicationCore_OsmodalLoop", 0),
		/*108*/ imports.NewTable("CefApplicationCore_OverrideSpellCheckLang", 0),
		/*109*/ imports.NewTable("CefApplicationCore_PackLoadingDisabled", 0),
		/*110*/ imports.NewTable("CefApplicationCore_PersistSessionCookies", 0),
		/*111*/ imports.NewTable("CefApplicationCore_PersistUserPreferences", 0),
		/*112*/ imports.NewTable("CefApplicationCore_ProcessType", 0),
		/*113*/ imports.NewTable("CefApplicationCore_QuitMessageLoop", 0),
		/*114*/ imports.NewTable("CefApplicationCore_ReRaiseExceptions", 0),
		/*115*/ imports.NewTable("CefApplicationCore_RemoteAllowOrigins", 0),
		/*116*/ imports.NewTable("CefApplicationCore_RemoteDebuggingPort", 0),
		/*117*/ imports.NewTable("CefApplicationCore_ResourcesDirPath", 0),
		/*118*/ imports.NewTable("CefApplicationCore_RootCache", 0),
		/*119*/ imports.NewTable("CefApplicationCore_RunMessageLoop", 0),
		/*120*/ imports.NewTable("CefApplicationCore_SetCurrentDir", 0),
		/*121*/ imports.NewTable("CefApplicationCore_SetOnBeforeChildProcessLaunch", 0),
		/*122*/ imports.NewTable("CefApplicationCore_SetOnBrowserCreated", 0),
		/*123*/ imports.NewTable("CefApplicationCore_SetOnBrowserDestroyed", 0),
		/*124*/ imports.NewTable("CefApplicationCore_SetOnContextCreated", 0),
		/*125*/ imports.NewTable("CefApplicationCore_SetOnContextInitialized", 0),
		/*126*/ imports.NewTable("CefApplicationCore_SetOnContextReleased", 0),
		/*127*/ imports.NewTable("CefApplicationCore_SetOnFocusedNodeChanged", 0),
		/*128*/ imports.NewTable("CefApplicationCore_SetOnGetDataResource", 0),
		/*129*/ imports.NewTable("CefApplicationCore_SetOnGetDataResourceForScale", 0),
		/*130*/ imports.NewTable("CefApplicationCore_SetOnGetDefaultClient", 0),
		/*131*/ imports.NewTable("CefApplicationCore_SetOnGetLocalizedString", 0),
		/*132*/ imports.NewTable("CefApplicationCore_SetOnLoadEnd", 0),
		/*133*/ imports.NewTable("CefApplicationCore_SetOnLoadError", 0),
		/*134*/ imports.NewTable("CefApplicationCore_SetOnLoadStart", 0),
		/*135*/ imports.NewTable("CefApplicationCore_SetOnLoadingStateChange", 0),
		/*136*/ imports.NewTable("CefApplicationCore_SetOnProcessMessageReceived", 0),
		/*137*/ imports.NewTable("CefApplicationCore_SetOnRegCustomSchemes", 0),
		/*138*/ imports.NewTable("CefApplicationCore_SetOnRegisterCustomPreferences", 0),
		/*139*/ imports.NewTable("CefApplicationCore_SetOnScheduleMessagePumpWork", 0),
		/*140*/ imports.NewTable("CefApplicationCore_SetOnUncaughtException", 0),
		/*141*/ imports.NewTable("CefApplicationCore_SetOnWebKitInitialized", 0),
		/*142*/ imports.NewTable("CefApplicationCore_ShowMessageDlg", 0),
		/*143*/ imports.NewTable("CefApplicationCore_SingleProcess", 0),
		/*144*/ imports.NewTable("CefApplicationCore_SitePerProcess", 0),
		/*145*/ imports.NewTable("CefApplicationCore_SmoothScrolling", 0),
		/*146*/ imports.NewTable("CefApplicationCore_StartMainProcess", 0),
		/*147*/ imports.NewTable("CefApplicationCore_StartSubProcess", 0),
		/*148*/ imports.NewTable("CefApplicationCore_Status", 0),
		/*149*/ imports.NewTable("CefApplicationCore_SystemMemoryLoad", 0),
		/*150*/ imports.NewTable("CefApplicationCore_TotalSystemMemory", 0),
		/*151*/ imports.NewTable("CefApplicationCore_TouchEvents", 0),
		/*152*/ imports.NewTable("CefApplicationCore_TreatInsecureOriginAsSecure", 0),
		/*153*/ imports.NewTable("CefApplicationCore_UIColorMode", 0),
		/*154*/ imports.NewTable("CefApplicationCore_UncaughtExceptionStackSize", 0),
		/*155*/ imports.NewTable("CefApplicationCore_UpdateDeviceScaleFactor", 0),
		/*156*/ imports.NewTable("CefApplicationCore_UseFakeUIForMediaStream", 0),
		/*157*/ imports.NewTable("CefApplicationCore_UseMockKeyChain", 0),
		/*158*/ imports.NewTable("CefApplicationCore_UsedMemory", 0),
		/*159*/ imports.NewTable("CefApplicationCore_UserAgent", 0),
		/*160*/ imports.NewTable("CefApplicationCore_UserAgentProduct", 0),
		/*161*/ imports.NewTable("CefApplicationCore_WindowlessRenderingEnabled", 0),
		/*162*/ imports.NewTable("CefApplicationCore_WindowsSandboxInfo", 0),
	}
)

func applicationCoreImportAPI() *imports.Imports {
	if applicationCoreImport == nil {
		applicationCoreImport = NewDefaultImports()
		applicationCoreImport.SetImportTable(applicationCoreImportTables)
		applicationCoreImportTables = nil
	}
	return applicationCoreImport
}
