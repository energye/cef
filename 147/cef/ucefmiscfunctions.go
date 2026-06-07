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

	cefTypes "github.com/energye/cef/147/types"
)

// MiscFunc  is static instance
var MiscFunc _MiscFuncClass

// _MiscFuncClass is class type defined by MiscFunc
type _MiscFuncClass uintptr

// CefColorGetA
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
func (_MiscFuncClass) CefColorGetA(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(0, uintptr(color))
	return byte(r)
}

// CefColorGetR
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
//	<summary>Return the red byte from a cef_color_t value.</summary>
func (_MiscFuncClass) CefColorGetR(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(1, uintptr(color))
	return byte(r)
}

// CefColorGetG
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
//	<summary>Return the red byte from a cef_color_t value.</summary>
//	<summary>Return the green byte from a cef_color_t value.</summary>
func (_MiscFuncClass) CefColorGetG(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(2, uintptr(color))
	return byte(r)
}

// CefColorGetB
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
//	<summary>Return the red byte from a cef_color_t value.</summary>
//	<summary>Return the green byte from a cef_color_t value.</summary>
//	<summary>Return the blue byte from a cef_color_t value.</summary>
func (_MiscFuncClass) CefColorGetB(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(3, uintptr(color))
	return byte(r)
}

// CefColorSetARGB
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
//	<summary>Return the red byte from a cef_color_t value.</summary>
//	<summary>Return the green byte from a cef_color_t value.</summary>
//	<summary>Return the blue byte from a cef_color_t value.</summary>
//	<summary>Return an cef_color_t value with the specified byte component values.</summary>
func (_MiscFuncClass) CefColorSetARGB(A byte, R byte, G byte, B byte) cefTypes.TCefColor {
	r := uCEFMiscFunctionsAPI().SysCallN(4, uintptr(A), uintptr(R), uintptr(G), uintptr(B))
	return cefTypes.TCefColor(r)
}

// CefGetObject
//
//	<summary>Return the alpha byte from a cef_color_t value.</summary>
//	<summary>Return the red byte from a cef_color_t value.</summary>
//	<summary>Return the green byte from a cef_color_t value.</summary>
//	<summary>Return the blue byte from a cef_color_t value.</summary>
//	<summary>Return an cef_color_t value with the specified byte component values.</summary>
//	<summary>Return an int64_t value with the specified low and high int32_t component values.</summary>
//	<summary>Return the low int32_t value from an int64_t value.</summary>
//	<summary>Return the high int32_t value from an int64_t value.</summary>
func (_MiscFuncClass) CefGetObject(ptr uintptr) lcl.IObject {
	r := uCEFMiscFunctionsAPI().SysCallN(5, uintptr(ptr))
	return lcl.AsObject(r)
}

// CefRegisterExtension
//
//	<summary>Converts ustring to TCefString.</summary>
//	<summary>Converts PCefString to ustring.</summary>
//	Register a new V8 extension with the specified JavaScript extension code and
//	handler. Functions implemented by the handler are prototyped using the
//	keyword 'native'. The calling of a native function is restricted to the
//	scope in which the prototype of the native function is defined. This
//	function may only be called on the render process main thread.
//
//	Example JavaScript extension code: <code>
//	// create the 'example' global object if it doesn't already exist.
//	if (!example)
//	example = {};
//	// create the 'example.test' global object if it doesn't already exist.
//	if (!example.test)
//	example.test = {};
//	(function() {
//	// Define the function 'example.test.myfunction'.
//	example.test.myfunction = function() {
//	// Call CefV8Handler::Execute() with the function name 'MyFunction'
//	// and no arguments.
//	native function MyFunction();
//	return MyFunction();
//	};
//	// Define the getter function for parameter 'example.test.myparam'.
//	example.test.__defineGetter__('myparam', function() {
//	// Call CefV8Handler::Execute() with the function name 'GetMyParam'
//	// and no arguments.
//	native function GetMyParam();
//	return GetMyParam();
//	});
//	// Define the setter function for parameter 'example.test.myparam'.
//	example.test.__defineSetter__('myparam', function(b) {
//	// Call CefV8Handler::Execute() with the function name 'SetMyParam'
//	// and a single argument.
//	native function SetMyParam();
//	if(b) SetMyParam(b);
//	});
//
//	// Extension definitions can also contain normal JavaScript variables
//	// and functions.
//	var myint = 0;
//	example.test.increment = function() {
//	myint += 1;
//	return myint;
//	};
//	})();
//	</code>
//
//	Example usage in the page: <code>
//	// Call the function.
//	example.test.myfunction();
//	// Set the parameter.
//	example.test.myparam = value;
//	// Get the parameter.
//	value = example.test.myparam;
//	// Call another function.
//	example.test.increment();
//	</code>
func (_MiscFuncClass) CefRegisterExtension(name string, code string, handler IEngV8Handler) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(6, api.PasStr(name), api.PasStr(code), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

// CefPostTask
//
//	<summary>Converts ustring to TCefString.</summary>
//	<summary>Converts PCefString to ustring.</summary>
//	Register a new V8 extension with the specified JavaScript extension code and
//	handler. Functions implemented by the handler are prototyped using the
//	keyword 'native'. The calling of a native function is restricted to the
//	scope in which the prototype of the native function is defined. This
//	function may only be called on the render process main thread.
//
//	Example JavaScript extension code: <code>
//	// create the 'example' global object if it doesn't already exist.
//	if (!example)
//	example = {};
//	// create the 'example.test' global object if it doesn't already exist.
//	if (!example.test)
//	example.test = {};
//	(function() {
//	// Define the function 'example.test.myfunction'.
//	example.test.myfunction = function() {
//	// Call CefV8Handler::Execute() with the function name 'MyFunction'
//	// and no arguments.
//	native function MyFunction();
//	return MyFunction();
//	};
//	// Define the getter function for parameter 'example.test.myparam'.
//	example.test.__defineGetter__('myparam', function() {
//	// Call CefV8Handler::Execute() with the function name 'GetMyParam'
//	// and no arguments.
//	native function GetMyParam();
//	return GetMyParam();
//	});
//	// Define the setter function for parameter 'example.test.myparam'.
//	example.test.__defineSetter__('myparam', function(b) {
//	// Call CefV8Handler::Execute() with the function name 'SetMyParam'
//	// and a single argument.
//	native function SetMyParam();
//	if(b) SetMyParam(b);
//	});
//
//	// Extension definitions can also contain normal JavaScript variables
//	// and functions.
//	var myint = 0;
//	example.test.increment = function() {
//	myint += 1;
//	return myint;
//	};
//	})();
//	</code>
//
//	Example usage in the page: <code>
//	// Call the function.
//	example.test.myfunction();
//	// Set the parameter.
//	example.test.myparam = value;
//	// Get the parameter.
//	value = example.test.myparam;
//	// Call another function.
//	example.test.increment();
//	</code>
//	Post a task for execution on the specified thread. Equivalent to using
//	TCefTaskRunnerRef.GetForThread(threadId).PostTask(task).
func (_MiscFuncClass) CefPostTask(threadId cefTypes.TCefThreadId, task IEngTask) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(7, uintptr(threadId), base.GetObjectUintptr(task))
	return api.GoBool(r)
}

// CefPostDelayedTask
//
//	<summary>Converts ustring to TCefString.</summary>
//	<summary>Converts PCefString to ustring.</summary>
//	Register a new V8 extension with the specified JavaScript extension code and
//	handler. Functions implemented by the handler are prototyped using the
//	keyword 'native'. The calling of a native function is restricted to the
//	scope in which the prototype of the native function is defined. This
//	function may only be called on the render process main thread.
//
//	Example JavaScript extension code: <code>
//	// create the 'example' global object if it doesn't already exist.
//	if (!example)
//	example = {};
//	// create the 'example.test' global object if it doesn't already exist.
//	if (!example.test)
//	example.test = {};
//	(function() {
//	// Define the function 'example.test.myfunction'.
//	example.test.myfunction = function() {
//	// Call CefV8Handler::Execute() with the function name 'MyFunction'
//	// and no arguments.
//	native function MyFunction();
//	return MyFunction();
//	};
//	// Define the getter function for parameter 'example.test.myparam'.
//	example.test.__defineGetter__('myparam', function() {
//	// Call CefV8Handler::Execute() with the function name 'GetMyParam'
//	// and no arguments.
//	native function GetMyParam();
//	return GetMyParam();
//	});
//	// Define the setter function for parameter 'example.test.myparam'.
//	example.test.__defineSetter__('myparam', function(b) {
//	// Call CefV8Handler::Execute() with the function name 'SetMyParam'
//	// and a single argument.
//	native function SetMyParam();
//	if(b) SetMyParam(b);
//	});
//
//	// Extension definitions can also contain normal JavaScript variables
//	// and functions.
//	var myint = 0;
//	example.test.increment = function() {
//	myint += 1;
//	return myint;
//	};
//	})();
//	</code>
//
//	Example usage in the page: <code>
//	// Call the function.
//	example.test.myfunction();
//	// Set the parameter.
//	example.test.myparam = value;
//	// Get the parameter.
//	value = example.test.myparam;
//	// Call another function.
//	example.test.increment();
//	</code>
//	Post a task for execution on the specified thread. Equivalent to using
//	TCefTaskRunnerRef.GetForThread(threadId).PostTask(task).
//	Post a task for delayed execution on the specified thread. Equivalent to
//	using TCefTaskRunnerRef.GetForThread(threadId).PostDelayedTask(task,
//	delay_ms).
func (_MiscFuncClass) CefPostDelayedTask(threadId cefTypes.TCefThreadId, task IEngTask, delayMs int64) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(8, uintptr(threadId), base.GetObjectUintptr(task), uintptr(base.UnsafePointer(&delayMs)))
	return api.GoBool(r)
}

// CefCurrentlyOn
//
//	<summary>Converts ustring to TCefString.</summary>
//	<summary>Converts PCefString to ustring.</summary>
//	Register a new V8 extension with the specified JavaScript extension code and
//	handler. Functions implemented by the handler are prototyped using the
//	keyword 'native'. The calling of a native function is restricted to the
//	scope in which the prototype of the native function is defined. This
//	function may only be called on the render process main thread.
//
//	Example JavaScript extension code: <code>
//	// create the 'example' global object if it doesn't already exist.
//	if (!example)
//	example = {};
//	// create the 'example.test' global object if it doesn't already exist.
//	if (!example.test)
//	example.test = {};
//	(function() {
//	// Define the function 'example.test.myfunction'.
//	example.test.myfunction = function() {
//	// Call CefV8Handler::Execute() with the function name 'MyFunction'
//	// and no arguments.
//	native function MyFunction();
//	return MyFunction();
//	};
//	// Define the getter function for parameter 'example.test.myparam'.
//	example.test.__defineGetter__('myparam', function() {
//	// Call CefV8Handler::Execute() with the function name 'GetMyParam'
//	// and no arguments.
//	native function GetMyParam();
//	return GetMyParam();
//	});
//	// Define the setter function for parameter 'example.test.myparam'.
//	example.test.__defineSetter__('myparam', function(b) {
//	// Call CefV8Handler::Execute() with the function name 'SetMyParam'
//	// and a single argument.
//	native function SetMyParam();
//	if(b) SetMyParam(b);
//	});
//
//	// Extension definitions can also contain normal JavaScript variables
//	// and functions.
//	var myint = 0;
//	example.test.increment = function() {
//	myint += 1;
//	return myint;
//	};
//	})();
//	</code>
//
//	Example usage in the page: <code>
//	// Call the function.
//	example.test.myfunction();
//	// Set the parameter.
//	example.test.myparam = value;
//	// Get the parameter.
//	value = example.test.myparam;
//	// Call another function.
//	example.test.increment();
//	</code>
//	Post a task for execution on the specified thread. Equivalent to using
//	TCefTaskRunnerRef.GetForThread(threadId).PostTask(task).
//	Post a task for delayed execution on the specified thread. Equivalent to
//	using TCefTaskRunnerRef.GetForThread(threadId).PostDelayedTask(task,
//	delay_ms).
//	Returns true (1) if called on the specified thread. Equivalent to using
//	TCefTaskRunnerRef.GetForThread(threadId).BelongsToCurrentThread().
func (_MiscFuncClass) CefCurrentlyOn(threadId cefTypes.TCefThreadId) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(9, uintptr(threadId))
	return api.GoBool(r)
}

// FixCefTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
func (_MiscFuncClass) FixCefTime(dt TCefTime) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(10, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CefTimeToDateTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
func (_MiscFuncClass) CefTimeToDateTime(dt TCefTime) (result types.TDateTime) {
	uCEFMiscFunctionsAPI().SysCallN(11, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// DateTimeToCefTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
func (_MiscFuncClass) DateTimeToCefTime(dt types.TDateTime) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(12, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// DateTimeToCefBaseTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
func (_MiscFuncClass) DateTimeToCefBaseTime(dt types.TDateTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(13, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CefTimeToDouble
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
func (_MiscFuncClass) CefTimeToDouble(dt TCefTime) (result float64) {
	uCEFMiscFunctionsAPI().SysCallN(14, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// DoubleToCefTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
func (_MiscFuncClass) DoubleToCefTime(dt float64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(15, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CefTimeToUnixTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
func (_MiscFuncClass) CefTimeToUnixTime(dt TCefTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(16, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// UnixTimeToCefTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
func (_MiscFuncClass) UnixTimeToCefTime(dt int64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(17, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CefTimeNow
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
func (_MiscFuncClass) CefTimeNow() (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(18, uintptr(base.UnsafePointer(&result)))
	return
}

// DoubleTimeNow
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
func (_MiscFuncClass) DoubleTimeNow() (result float64) {
	uCEFMiscFunctionsAPI().SysCallN(19, uintptr(base.UnsafePointer(&result)))
	return
}

// CefBaseTimeNow
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the delta in milliseconds between two time values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_delta)</see>
//	Retrieve the current system time in a TCefBaseTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_now)</see>
func (_MiscFuncClass) CefBaseTimeNow() (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(20, uintptr(base.UnsafePointer(&result)))
	return
}

// CetTimeToCefBaseTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the delta in milliseconds between two time values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_delta)</see>
//	Retrieve the current system time in a TCefBaseTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_now)</see>
//	Converts TCefTime to TCefBaseTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_basetime)</see>
func (_MiscFuncClass) CetTimeToCefBaseTime(ct TCefTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(21, uintptr(base.UnsafePointer(&ct)), uintptr(base.UnsafePointer(&result)))
	return
}

// CetTimeFromCefBaseTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the delta in milliseconds between two time values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_delta)</see>
//	Retrieve the current system time in a TCefBaseTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_now)</see>
//	Converts TCefTime to TCefBaseTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_basetime)</see>
//	Converts TCefBaseTime to TCefTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_basetime)</see>
func (_MiscFuncClass) CetTimeFromCefBaseTime(cbt int64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(22, uintptr(base.UnsafePointer(&cbt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CefBaseTimeToDateTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the delta in milliseconds between two time values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_delta)</see>
//	Retrieve the current system time in a TCefBaseTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_now)</see>
//	Converts TCefTime to TCefBaseTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_basetime)</see>
//	Converts TCefBaseTime to TCefTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_basetime)</see>
//	Converts TCefBaseTime to TDateTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_basetime)</see>
func (_MiscFuncClass) CefBaseTimeToDateTime(cbt int64) (result types.TDateTime) {
	uCEFMiscFunctionsAPI().SysCallN(23, uintptr(base.UnsafePointer(&cbt)), uintptr(base.UnsafePointer(&result)))
	return
}

// CustomPathIsRelative
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
func (_MiscFuncClass) CustomPathIsRelative(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(24, api.PasStr(path))
	return api.GoBool(r)
}

// CustomPathCanonicalize
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
func (_MiscFuncClass) CustomPathCanonicalize(originalPath string, canonicalPath *string) bool {
	canonicalPathPtr := api.PasStr(*canonicalPath)
	r := uCEFMiscFunctionsAPI().SysCallN(25, api.PasStr(originalPath), uintptr(base.UnsafePointer(&canonicalPathPtr)))
	*canonicalPath = api.GoStr(canonicalPathPtr)
	return api.GoBool(r)
}

// CustomAbsolutePath
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
func (_MiscFuncClass) CustomAbsolutePath(path string, mustExist bool) string {
	r := uCEFMiscFunctionsAPI().SysCallN(26, api.PasStr(path), api.PasBool(mustExist))
	return api.GoStr(r)
}

// CustomPathIsURL
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
func (_MiscFuncClass) CustomPathIsURL(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(27, api.PasStr(path))
	return api.GoBool(r)
}

// CustomPathIsUNC
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
func (_MiscFuncClass) CustomPathIsUNC(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(28, api.PasStr(path))
	return api.GoBool(r)
}

// GetModulePath
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
func (_MiscFuncClass) GetModulePath() string {
	r := uCEFMiscFunctionsAPI().SysCallN(29)
	return api.GoStr(r)
}

// CefIsCertStatusError
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
func (_MiscFuncClass) CefIsCertStatusError(status cefTypes.TCefCertStatus) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(30, uintptr(status))
	return api.GoBool(r)
}

// CefCrashReportingEnabled
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
func (_MiscFuncClass) CefCrashReportingEnabled() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(31)
	return api.GoBool(r)
}

// CefGetMinLogLevel
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) CefGetMinLogLevel() int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(32)
	return int32(r)
}

// CefGetLogSeverityName
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) CefGetLogSeverityName(severity int32) string {
	r := uCEFMiscFunctionsAPI().SysCallN(33, uintptr(severity))
	return api.GoStr(r)
}

// CefRegisterSchemeHandlerFactory
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Register a scheme handler factory with the global request context. An NULL
//	|DomainName| value for a standard scheme will cause the factory to match
//	all domain names. The |DomainName| value will be ignored for non-standard
//	schemes. If |SchemeName| is a built-in scheme and no handler is returned by
//	|factory| then the built-in scheme handler factory will be called. If
//	|SchemeName| is a custom scheme then you must also implement the
//	ICefApp.OnRegisterCustomSchemes function in all processes. This
//	function may be called multiple times to change or remove the factory that
//	matches the specified |SchemeName| and optional |DomainName|. Returns
//	false (0) if an error occurs. This function may be called on any thread in
//	the browser process. Using this function is equivalent to calling cef_reques
//	t_context_t::cef_request_context_get_global_context()->register_scheme_handl
//	er_factory().
func (_MiscFuncClass) CefRegisterSchemeHandlerFactory(schemeName string, domainName string, handler cefTypes.TCefResourceHandlerClass) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(34, api.PasStr(schemeName), api.PasStr(domainName), uintptr(handler))
	return api.GoBool(r)
}

// CefClearSchemeHandlerFactories
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Register a scheme handler factory with the global request context. An NULL
//	|DomainName| value for a standard scheme will cause the factory to match
//	all domain names. The |DomainName| value will be ignored for non-standard
//	schemes. If |SchemeName| is a built-in scheme and no handler is returned by
//	|factory| then the built-in scheme handler factory will be called. If
//	|SchemeName| is a custom scheme then you must also implement the
//	ICefApp.OnRegisterCustomSchemes function in all processes. This
//	function may be called multiple times to change or remove the factory that
//	matches the specified |SchemeName| and optional |DomainName|. Returns
//	false (0) if an error occurs. This function may be called on any thread in
//	the browser process. Using this function is equivalent to calling cef_reques
//	t_context_t::cef_request_context_get_global_context()->register_scheme_handl
//	er_factory().
//	Clear all scheme handler factories registered with the global request
//	context. Returns false (0) on error. This function may be called on any
//	thread in the browser process. Using this function is equivalent to calling
//	cef_request_context_t::cef_request_context_get_global_context()->clear_schem
//	e_handler_factories().
func (_MiscFuncClass) CefClearSchemeHandlerFactories() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(35)
	return api.GoBool(r)
}

// CefAddCrossOriginWhitelistEntry
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Register a scheme handler factory with the global request context. An NULL
//	|DomainName| value for a standard scheme will cause the factory to match
//	all domain names. The |DomainName| value will be ignored for non-standard
//	schemes. If |SchemeName| is a built-in scheme and no handler is returned by
//	|factory| then the built-in scheme handler factory will be called. If
//	|SchemeName| is a custom scheme then you must also implement the
//	ICefApp.OnRegisterCustomSchemes function in all processes. This
//	function may be called multiple times to change or remove the factory that
//	matches the specified |SchemeName| and optional |DomainName|. Returns
//	false (0) if an error occurs. This function may be called on any thread in
//	the browser process. Using this function is equivalent to calling cef_reques
//	t_context_t::cef_request_context_get_global_context()->register_scheme_handl
//	er_factory().
//	Clear all scheme handler factories registered with the global request
//	context. Returns false (0) on error. This function may be called on any
//	thread in the browser process. Using this function is equivalent to calling
//	cef_request_context_t::cef_request_context_get_global_context()->clear_schem
//	e_handler_factories().
//	Add an entry to the cross-origin access whitelist.
//	The same-origin policy restricts how scripts hosted from different origins
//	(scheme + domain + port) can communicate. By default, scripts can only
//	access resources with the same origin. Scripts hosted on the HTTP and HTTPS
//	schemes (but no other schemes) can use the "Access-Control-Allow-Origin"
//	header to allow cross-origin requests. For example,
//	https://source.example.com can make XMLHttpRequest requests on
//	http://target.example.com if the http://target.example.com request returns
//	an "Access-Control-Allow-Origin: https://source.example.com" response
//	header.
//	Scripts in separate frames or iframes and hosted from the same protocol and
//	domain suffix can execute cross-origin JavaScript if both pages set the
//	document.domain value to the same domain suffix. For example,
//	scheme://foo.example.com and scheme://bar.example.com can communicate using
//	JavaScript if both domains set document.domain="example.com".
//	This function is used to allow access to origins that would otherwise
//	violate the same-origin policy. Scripts hosted underneath the fully
//	qualified |source_origin| URL (like http://www.example.com) will be allowed
//	access to all resources hosted on the specified |target_protocol| and
//	|target_domain|. If |target_domain| is non-NULL and
//	|allow_target_subdomains| is false (0) only exact domain matches will be
//	allowed. If |target_domain| contains a top- level domain component (like
//	"example.com") and |allow_target_subdomains| is true (1) sub-domain matches
//	will be allowed. If |target_domain| is NULL and |allow_target_subdomains| if
//	true (1) all domains and IP addresses will be allowed.
//	This function cannot be used to bypass the restrictions on local or display
//	isolated schemes. See the comments on CefRegisterCustomScheme for more
//	information.
//	This function may be called on any thread. Returns false (0) if
//	|source_origin| is invalid or the whitelist cannot be accessed.
func (_MiscFuncClass) CefAddCrossOriginWhitelistEntry(sourceOrigin string, targetProtocol string, targetDomain string, allowTargetSubdomains bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(36, api.PasStr(sourceOrigin), api.PasStr(targetProtocol), api.PasStr(targetDomain), api.PasBool(allowTargetSubdomains))
	return api.GoBool(r)
}

// CefRemoveCrossOriginWhitelistEntry
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Register a scheme handler factory with the global request context. An NULL
//	|DomainName| value for a standard scheme will cause the factory to match
//	all domain names. The |DomainName| value will be ignored for non-standard
//	schemes. If |SchemeName| is a built-in scheme and no handler is returned by
//	|factory| then the built-in scheme handler factory will be called. If
//	|SchemeName| is a custom scheme then you must also implement the
//	ICefApp.OnRegisterCustomSchemes function in all processes. This
//	function may be called multiple times to change or remove the factory that
//	matches the specified |SchemeName| and optional |DomainName|. Returns
//	false (0) if an error occurs. This function may be called on any thread in
//	the browser process. Using this function is equivalent to calling cef_reques
//	t_context_t::cef_request_context_get_global_context()->register_scheme_handl
//	er_factory().
//	Clear all scheme handler factories registered with the global request
//	context. Returns false (0) on error. This function may be called on any
//	thread in the browser process. Using this function is equivalent to calling
//	cef_request_context_t::cef_request_context_get_global_context()->clear_schem
//	e_handler_factories().
//	Add an entry to the cross-origin access whitelist.
//	The same-origin policy restricts how scripts hosted from different origins
//	(scheme + domain + port) can communicate. By default, scripts can only
//	access resources with the same origin. Scripts hosted on the HTTP and HTTPS
//	schemes (but no other schemes) can use the "Access-Control-Allow-Origin"
//	header to allow cross-origin requests. For example,
//	https://source.example.com can make XMLHttpRequest requests on
//	http://target.example.com if the http://target.example.com request returns
//	an "Access-Control-Allow-Origin: https://source.example.com" response
//	header.
//	Scripts in separate frames or iframes and hosted from the same protocol and
//	domain suffix can execute cross-origin JavaScript if both pages set the
//	document.domain value to the same domain suffix. For example,
//	scheme://foo.example.com and scheme://bar.example.com can communicate using
//	JavaScript if both domains set document.domain="example.com".
//	This function is used to allow access to origins that would otherwise
//	violate the same-origin policy. Scripts hosted underneath the fully
//	qualified |source_origin| URL (like http://www.example.com) will be allowed
//	access to all resources hosted on the specified |target_protocol| and
//	|target_domain|. If |target_domain| is non-NULL and
//	|allow_target_subdomains| is false (0) only exact domain matches will be
//	allowed. If |target_domain| contains a top- level domain component (like
//	"example.com") and |allow_target_subdomains| is true (1) sub-domain matches
//	will be allowed. If |target_domain| is NULL and |allow_target_subdomains| if
//	true (1) all domains and IP addresses will be allowed.
//	This function cannot be used to bypass the restrictions on local or display
//	isolated schemes. See the comments on CefRegisterCustomScheme for more
//	information.
//	This function may be called on any thread. Returns false (0) if
//	|source_origin| is invalid or the whitelist cannot be accessed.
//	Remove an entry from the cross-origin access whitelist. Returns false (0) if
//	|source_origin| is invalid or the whitelist cannot be accessed.
func (_MiscFuncClass) CefRemoveCrossOriginWhitelistEntry(sourceOrigin string, targetProtocol string, targetDomain string, allowTargetSubdomains bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(37, api.PasStr(sourceOrigin), api.PasStr(targetProtocol), api.PasStr(targetDomain), api.PasBool(allowTargetSubdomains))
	return api.GoBool(r)
}

// CefClearCrossOriginWhitelist
//
//	Gets the current log verbose level (LogSeverity).
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_min_log_level)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Gets the current vlog level for the given file.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_get_vlog_level)</see>
//	Gets the log severity name.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
//	Register a scheme handler factory with the global request context. An NULL
//	|DomainName| value for a standard scheme will cause the factory to match
//	all domain names. The |DomainName| value will be ignored for non-standard
//	schemes. If |SchemeName| is a built-in scheme and no handler is returned by
//	|factory| then the built-in scheme handler factory will be called. If
//	|SchemeName| is a custom scheme then you must also implement the
//	ICefApp.OnRegisterCustomSchemes function in all processes. This
//	function may be called multiple times to change or remove the factory that
//	matches the specified |SchemeName| and optional |DomainName|. Returns
//	false (0) if an error occurs. This function may be called on any thread in
//	the browser process. Using this function is equivalent to calling cef_reques
//	t_context_t::cef_request_context_get_global_context()->register_scheme_handl
//	er_factory().
//	Clear all scheme handler factories registered with the global request
//	context. Returns false (0) on error. This function may be called on any
//	thread in the browser process. Using this function is equivalent to calling
//	cef_request_context_t::cef_request_context_get_global_context()->clear_schem
//	e_handler_factories().
//	Add an entry to the cross-origin access whitelist.
//	The same-origin policy restricts how scripts hosted from different origins
//	(scheme + domain + port) can communicate. By default, scripts can only
//	access resources with the same origin. Scripts hosted on the HTTP and HTTPS
//	schemes (but no other schemes) can use the "Access-Control-Allow-Origin"
//	header to allow cross-origin requests. For example,
//	https://source.example.com can make XMLHttpRequest requests on
//	http://target.example.com if the http://target.example.com request returns
//	an "Access-Control-Allow-Origin: https://source.example.com" response
//	header.
//	Scripts in separate frames or iframes and hosted from the same protocol and
//	domain suffix can execute cross-origin JavaScript if both pages set the
//	document.domain value to the same domain suffix. For example,
//	scheme://foo.example.com and scheme://bar.example.com can communicate using
//	JavaScript if both domains set document.domain="example.com".
//	This function is used to allow access to origins that would otherwise
//	violate the same-origin policy. Scripts hosted underneath the fully
//	qualified |source_origin| URL (like http://www.example.com) will be allowed
//	access to all resources hosted on the specified |target_protocol| and
//	|target_domain|. If |target_domain| is non-NULL and
//	|allow_target_subdomains| is false (0) only exact domain matches will be
//	allowed. If |target_domain| contains a top- level domain component (like
//	"example.com") and |allow_target_subdomains| is true (1) sub-domain matches
//	will be allowed. If |target_domain| is NULL and |allow_target_subdomains| if
//	true (1) all domains and IP addresses will be allowed.
//	This function cannot be used to bypass the restrictions on local or display
//	isolated schemes. See the comments on CefRegisterCustomScheme for more
//	information.
//	This function may be called on any thread. Returns false (0) if
//	|source_origin| is invalid or the whitelist cannot be accessed.
//	Remove an entry from the cross-origin access whitelist. Returns false (0) if
//	|source_origin| is invalid or the whitelist cannot be accessed.
//	Remove all entries from the cross-origin access whitelist. Returns false (0)
//	if the whitelist cannot be accessed.
func (_MiscFuncClass) CefClearCrossOriginWhitelist() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(38)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetExtendedFileVersion(fileName string) (result uint64) {
	uCEFMiscFunctionsAPI().SysCallN(39, api.PasStr(fileName), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) GetRegistryWindowsVersion(major *uint32, minor *uint32) bool {
	majorPtr := uintptr(*major)
	minorPtr := uintptr(*minor)
	r := uCEFMiscFunctionsAPI().SysCallN(40, uintptr(base.UnsafePointer(&majorPtr)), uintptr(base.UnsafePointer(&minorPtr)))
	*major = uint32(majorPtr)
	*minor = uint32(minorPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetRealWindowsVersion(major *uint32, minor *uint32) bool {
	majorPtr := uintptr(*major)
	minorPtr := uintptr(*minor)
	r := uCEFMiscFunctionsAPI().SysCallN(41, uintptr(base.UnsafePointer(&majorPtr)), uintptr(base.UnsafePointer(&minorPtr)))
	*major = uint32(majorPtr)
	*minor = uint32(minorPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckRealWindowsVersion(major uint32, minor uint32) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(42, uintptr(major), uintptr(minor))
	return api.GoBool(r)
}

func (_MiscFuncClass) SplitLongString(srcString string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(43, api.PasStr(srcString))
	return api.GoStr(r)
}

func (_MiscFuncClass) GetAbsoluteDirPath(srcPath string, rsltPath *string) bool {
	rsltPathPtr := api.PasStr(*rsltPath)
	r := uCEFMiscFunctionsAPI().SysCallN(44, api.PasStr(srcPath), uintptr(base.UnsafePointer(&rsltPathPtr)))
	*rsltPath = api.GoStr(rsltPathPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckSubprocessPath(subprocessPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(45, api.PasStr(subprocessPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckLocales(localesDirPath string, missingFiles *string, localesRequired string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(46, api.PasStr(localesDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)), api.PasStr(localesRequired))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckResources(resourcesDirPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(47, api.PasStr(resourcesDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckDLLs(frameworkDirPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(48, api.PasStr(frameworkDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckDLLVersion(dLLFile string, major uint16, minor uint16, release uint16, build uint16) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(49, api.PasStr(dLLFile), uintptr(major), uintptr(minor), uintptr(release), uintptr(build))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetDLLHeaderMachine(dLLFile string, machine *int32) bool {
	machinePtr := uintptr(*machine)
	r := uCEFMiscFunctionsAPI().SysCallN(50, api.PasStr(dLLFile), uintptr(base.UnsafePointer(&machinePtr)))
	*machine = int32(machinePtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetFileTypeDescription(extension string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(51, api.PasStr(extension))
	return api.GoStr(r)
}

func (_MiscFuncClass) CheckFilesExist(list *lcl.IStringList, missingFiles *string) bool {
	listPtr := base.GetObjectUintptr(*list)
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(52, uintptr(base.UnsafePointer(&listPtr)), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*list = lcl.AsStringList(listPtr)
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) Is32BitProcess() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(53)
	return api.GoBool(r)
}

// CefFormatUrlForSecurityDisplay
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
func (_MiscFuncClass) CefFormatUrlForSecurityDisplay(originUrl string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(54, api.PasStr(originUrl))
	return api.GoStr(r)
}

// CefGetMimeType
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
func (_MiscFuncClass) CefGetMimeType(extension string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(55, api.PasStr(extension))
	return api.GoStr(r)
}

// CefBase64Encode
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
func (_MiscFuncClass) CefBase64Encode(data uintptr, dataSize cefTypes.NativeUInt) string {
	r := uCEFMiscFunctionsAPI().SysCallN(56, uintptr(data), uintptr(dataSize))
	return api.GoStr(r)
}

// CefBase64Decode
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
func (_MiscFuncClass) CefBase64Decode(data string) (result ICefBinaryValue) {
	var resultPtr uintptr
	uCEFMiscFunctionsAPI().SysCallN(57, api.PasStr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

// CefUriEncode
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
func (_MiscFuncClass) CefUriEncode(text string, usePlus bool) string {
	r := uCEFMiscFunctionsAPI().SysCallN(58, api.PasStr(text), api.PasBool(usePlus))
	return api.GoStr(r)
}

// CefUriDecode
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
func (_MiscFuncClass) CefUriDecode(text string, convertToUtf8 bool, unescapeRule cefTypes.TCefUriUnescapeRule) string {
	r := uCEFMiscFunctionsAPI().SysCallN(59, api.PasStr(text), api.PasBool(convertToUtf8), uintptr(unescapeRule))
	return api.GoStr(r)
}

// CefGetPath
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
func (_MiscFuncClass) CefGetPath(pathKey cefTypes.TCefPathKey) string {
	r := uCEFMiscFunctionsAPI().SysCallN(60, uintptr(pathKey))
	return api.GoStr(r)
}

// CefIsRTL
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
func (_MiscFuncClass) CefIsRTL() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(61)
	return api.GoBool(r)
}

// CefCreateDirectory
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
func (_MiscFuncClass) CefCreateDirectory(fullPath string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(62, api.PasStr(fullPath))
	return api.GoBool(r)
}

// CefGetTempDirectory
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
func (_MiscFuncClass) CefGetTempDirectory(outTempDir *string) bool {
	var tempDirPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(63, uintptr(base.UnsafePointer(&tempDirPtr)))
	*outTempDir = api.GoStr(tempDirPtr)
	return api.GoBool(r)
}

// CefCreateNewTempDirectory
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
func (_MiscFuncClass) CefCreateNewTempDirectory(prefix string, outNewTempPath *string) bool {
	var newTempPathPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(64, api.PasStr(prefix), uintptr(base.UnsafePointer(&newTempPathPtr)))
	*outNewTempPath = api.GoStr(newTempPathPtr)
	return api.GoBool(r)
}

// CefCreateTempDirectoryInDirectory
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
func (_MiscFuncClass) CefCreateTempDirectoryInDirectory(baseDir string, prefix string, outNewDir *string) bool {
	var newDirPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(65, api.PasStr(baseDir), api.PasStr(prefix), uintptr(base.UnsafePointer(&newDirPtr)))
	*outNewDir = api.GoStr(newDirPtr)
	return api.GoBool(r)
}

// CefDirectoryExists
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
//	Returns true (1) if the given path exists and is a directory. Calling this
//	function on the browser process UI or IO threads is not allowed.
func (_MiscFuncClass) CefDirectoryExists(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(66, api.PasStr(path))
	return api.GoBool(r)
}

// CefDeleteFile
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
//	Returns true (1) if the given path exists and is a directory. Calling this
//	function on the browser process UI or IO threads is not allowed.
//	Deletes the given path whether it's a file or a directory. If |path| is a
//	directory all contents will be deleted. If |recursive| is true (1) any sub-
//	directories and their contents will also be deleted (equivalent to executing
//	"rm -rf", so use with caution). On POSIX environments if |path| is a
//	symbolic link then only the symlink will be deleted. Returns true (1) on
//	successful deletion or if |path| does not exist. Calling this function on
//	the browser process UI or IO threads is not allowed.
func (_MiscFuncClass) CefDeleteFile(path string, recursive bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(67, api.PasStr(path), api.PasBool(recursive))
	return api.GoBool(r)
}

// CefZipDirectory
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
//	Returns true (1) if the given path exists and is a directory. Calling this
//	function on the browser process UI or IO threads is not allowed.
//	Deletes the given path whether it's a file or a directory. If |path| is a
//	directory all contents will be deleted. If |recursive| is true (1) any sub-
//	directories and their contents will also be deleted (equivalent to executing
//	"rm -rf", so use with caution). On POSIX environments if |path| is a
//	symbolic link then only the symlink will be deleted. Returns true (1) on
//	successful deletion or if |path| does not exist. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Writes the contents of |srcDir| into a zip archive at |destFile|. If
//	|includeHiddenFiles| is true (1) files starting with "." will be included.
//	Returns true (1) on success. Calling this function on the browser process
//	UI or IO threads is not allowed.
func (_MiscFuncClass) CefZipDirectory(srcDir string, destFile string, includeHiddenFiles bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(68, api.PasStr(srcDir), api.PasStr(destFile), api.PasBool(includeHiddenFiles))
	return api.GoBool(r)
}

// GetDefaultCEFUserAgent
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
//	Returns true (1) if the given path exists and is a directory. Calling this
//	function on the browser process UI or IO threads is not allowed.
//	Deletes the given path whether it's a file or a directory. If |path| is a
//	directory all contents will be deleted. If |recursive| is true (1) any sub-
//	directories and their contents will also be deleted (equivalent to executing
//	"rm -rf", so use with caution). On POSIX environments if |path| is a
//	symbolic link then only the symlink will be deleted. Returns true (1) on
//	successful deletion or if |path| does not exist. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Writes the contents of |srcDir| into a zip archive at |destFile|. If
//	|includeHiddenFiles| is true (1) files starting with "." will be included.
//	Returns true (1) on success. Calling this function on the browser process
//	UI or IO threads is not allowed.
//	Loads the existing "Certificate Revocation Lists" file that is managed by
//	Google Chrome. This file can generally be found in Chrome's User Data
//	directory (e.g. "C:\Users\[User]\AppData\Local\Google\Chrome\User Data\" on
//	Windows) and is updated periodically by Chrome's component updater service.
//	Must be called in the browser process after the context has been
//	initialized. See https://dev.chromium.org/Home/chromium-security/crlsets for
//	background.
//	Return a user-agent string.
//	This function tries to replicate the BuildUserAgentFromOSAndProduct
//	function in Chromium but it's safer to call the 'Browser.getVersion'
//	DevTools method.
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:content/common/user_agent.cc">Chromium source file: content/common/user_agent.cc (BuildUserAgentFromOSAndProduct)</see>
//	<see href="https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion">See the Browser.getVersion article.</see>
func (_MiscFuncClass) GetDefaultCEFUserAgent() string {
	r := uCEFMiscFunctionsAPI().SysCallN(69)
	return api.GoStr(r)
}

func (_MiscFuncClass) CefIsKeyDown(wparam types.WParam) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(70, uintptr(wparam))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefIsKeyToggled(wparam types.WParam) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(71, uintptr(wparam))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetCefMouseModifiersToEventFlags() cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(72)
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetCefMouseModifiersWithWPARAM(awparam types.WParam) cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(73, uintptr(awparam))
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetCefKeyboardModifiers(wparam types.WParam, lparam types.LParam) cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(74, uintptr(wparam), uintptr(lparam))
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetWindowsMajorMinorVersion(wMajorVersion *types.DWORD, wMinorVersion *types.DWORD) bool {
	wMajorVersionPtr := uintptr(*wMajorVersion)
	wMinorVersionPtr := uintptr(*wMinorVersion)
	r := uCEFMiscFunctionsAPI().SysCallN(75, uintptr(base.UnsafePointer(&wMajorVersionPtr)), uintptr(base.UnsafePointer(&wMinorVersionPtr)))
	*wMajorVersion = types.DWORD(wMajorVersionPtr)
	*wMinorVersion = types.DWORD(wMinorVersionPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetIsWow64Process2(processMachine *uint16, nativeMachine *uint16) bool {
	processMachinePtr := uintptr(*processMachine)
	nativeMachinePtr := uintptr(*nativeMachine)
	r := uCEFMiscFunctionsAPI().SysCallN(76, uintptr(base.UnsafePointer(&processMachinePtr)), uintptr(base.UnsafePointer(&nativeMachinePtr)))
	*processMachine = uint16(processMachinePtr)
	*nativeMachine = uint16(nativeMachinePtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) IsWowProcess() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(77)
	return api.GoBool(r)
}

func (_MiscFuncClass) RunningWindows10OrNewer() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(78)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetDPIForHandle(handle types.HWND, dPI *types.UINT) bool {
	dPIPtr := uintptr(*dPI)
	r := uCEFMiscFunctionsAPI().SysCallN(79, uintptr(handle), uintptr(base.UnsafePointer(&dPIPtr)))
	*dPI = types.UINT(dPIPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) DeviceToLogicalWithIntDouble(value int32, deviceScaleFactor float64) int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(80, uintptr(value), uintptr(base.UnsafePointer(&deviceScaleFactor)))
	return int32(r)
}

func (_MiscFuncClass) DeviceToLogicalWithSingleDouble(value float32, deviceScaleFactor float64) (result float32) {
	uCEFMiscFunctionsAPI().SysCallN(81, uintptr(base.UnsafePointer(&value)), uintptr(base.UnsafePointer(&deviceScaleFactor)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) LogicalToDeviceWithIntDouble(value int32, deviceScaleFactor float64) int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(82, uintptr(value), uintptr(base.UnsafePointer(&deviceScaleFactor)))
	return int32(r)
}

func (_MiscFuncClass) GetScreenDPI() int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(83)
	return int32(r)
}

func (_MiscFuncClass) GetDeviceScaleFactor() (result float32) {
	uCEFMiscFunctionsAPI().SysCallN(84, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) TryRemoveDir(directory string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(85, api.PasStr(directory))
	return api.GoBool(r)
}

func (_MiscFuncClass) TryDeleteFile(fileName string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(86, api.PasStr(fileName))
	return api.GoBool(r)
}

func (_MiscFuncClass) TryRenameDir(oldName string, newName string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(87, api.PasStr(oldName), api.PasStr(newName))
	return api.GoBool(r)
}

func (_MiscFuncClass) TryRenameFile(oldName string, newName string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(88, api.PasStr(oldName), api.PasStr(newName))
	return api.GoBool(r)
}

func (_MiscFuncClass) DeleteDirContents(directory string, excludeFiles lcl.IStringList) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(89, api.PasStr(directory), base.GetObjectUintptr(excludeFiles))
	return api.GoBool(r)
}

func (_MiscFuncClass) DeleteFileList(fileList lcl.IStringList) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(90, base.GetObjectUintptr(fileList))
	return api.GoBool(r)
}

func (_MiscFuncClass) MoveFileList(fileList lcl.IStringList, srcDirectory string, dstDirectory string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(91, base.GetObjectUintptr(fileList), api.PasStr(srcDirectory), api.PasStr(dstDirectory))
	return api.GoBool(r)
}

// CefGetDataURIWithStrX2
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
func (_MiscFuncClass) CefGetDataURIWithStrX2(string string, mimeType string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(92, api.PasStr(string), api.PasStr(mimeType))
	return api.GoStr(r)
}

// CefGetDataURIWithPointerIntStrX2
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
//	Returns a URI with a DATA scheme encoding |aData| as a base64 string.
func (_MiscFuncClass) CefGetDataURIWithPointerIntStrX2(data uintptr, size int32, mimeType string, charset string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(93, uintptr(data), uintptr(size), api.PasStr(mimeType), api.PasStr(charset))
	return api.GoStr(r)
}

// ValidCefWindowHandle
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
//	Returns a URI with a DATA scheme encoding |aData| as a base64 string.
func (_MiscFuncClass) ValidCefWindowHandle(handle cefTypes.TCefWindowHandle) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(94, uintptr(handle))
	return api.GoBool(r)
}

// GetCommandLineSwitchValue
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
//	Returns a URI with a DATA scheme encoding |aData| as a base64 string.
//	Returns a command line switch value if it exists.
func (_MiscFuncClass) GetCommandLineSwitchValue(key string, value *string) bool {
	valuePtr := api.PasStr(*value)
	r := uCEFMiscFunctionsAPI().SysCallN(95, api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = api.GoStr(valuePtr)
	return api.GoBool(r)
}

// IsCEFSubprocess
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
//	Returns a URI with a DATA scheme encoding |aData| as a base64 string.
//	Returns a command line switch value if it exists.
//	Returns true if the command line switch has a "type" value.
func (_MiscFuncClass) IsCEFSubprocess() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(96)
	return api.GoBool(r)
}

// EditingCommandToString
//
//	Convert an editting command to string.
func (_MiscFuncClass) EditingCommandToString(editingCommand cefTypes.TCefEditingCommand) string {
	r := uCEFMiscFunctionsAPI().SysCallN(97, uintptr(editingCommand))
	return api.GoStr(r)
}

// CefResultCodeToString
//
//	Convert an editting command to string.
//	Convert the GlobalCEFApp.ExitCode value to a human readable message.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
//	<see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
func (_MiscFuncClass) CefResultCodeToString(exitCode cefTypes.TCefResultCode) string {
	r := uCEFMiscFunctionsAPI().SysCallN(98, uintptr(exitCode))
	return api.GoStr(r)
}

// InitializeCefTime
//
//	Returns a new TCefTime with a valid time in case the original has errors.
//	Converts a TCefTime value to TDateTime.
//	Converts a TDateTime value to TCefTime.
//	Converts a TDateTime value to TCefBaseTime.
//	Converts TCefTime to a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_doublet)</see>
//	Converts TCefTime from a double which is the number of seconds since
//	epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
//	means "not initialized".
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_doublet)</see>
//	Converts cef_time_t to time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_timet)</see>
//	Converts cef_time_t from time_t. time_t is almost always an integral value holding the number of seconds (not counting leap seconds) since 00:00, Jan 1 1970 UTC, corresponding to POSIX time.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_timet)</see>
//	Retrieve the current system time in a TCefTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the current system time in a double type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_now)</see>
//	Retrieve the delta in milliseconds between two time values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_delta)</see>
//	Retrieve the current system time in a TCefBaseTime type.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_now)</see>
//	Converts TCefTime to TCefBaseTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_to_basetime)</see>
//	Converts TCefBaseTime to TCefTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_basetime)</see>
//	Converts TCefBaseTime to TDateTime.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_from_basetime)</see>
//	Returns the time interval between now and from_ in milliseconds.
//	This funcion should only be used by TCEFTimerWorkScheduler.
//	Initialize a TCefTime variable.
func (_MiscFuncClass) InitializeCefTime(time *TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(99, uintptr(base.UnsafePointer(time)))
}

// CefSetCrashKeyValue
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
//	Sets or clears a specific key-value pair from the crash metadata.
func (_MiscFuncClass) CefSetCrashKeyValue(key string, value string) {
	uCEFMiscFunctionsAPI().SysCallN(100, api.PasStr(key), api.PasStr(value))
}

// CefLog
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
//	Sets or clears a specific key-value pair from the crash metadata.
//	Add a log message. See the LogSeverity defines for supported |severity|
//	values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_log)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) CefLog(file string, line int32, severity int32, message string) {
	uCEFMiscFunctionsAPI().SysCallN(101, api.PasStr(file), uintptr(line), uintptr(severity), api.PasStr(message))
}

// CefKeyEventLog
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
//	Sets or clears a specific key-value pair from the crash metadata.
//	Add a log message. See the LogSeverity defines for supported |severity|
//	values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_log)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) CefKeyEventLog(event TCefKeyEvent) {
	uCEFMiscFunctionsAPI().SysCallN(102, uintptr(base.UnsafePointer(&event)))
}

// CefMouseEventLog
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
//	Sets or clears a specific key-value pair from the crash metadata.
//	Add a log message. See the LogSeverity defines for supported |severity|
//	values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_log)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) CefMouseEventLog(event TCefMouseEvent) {
	uCEFMiscFunctionsAPI().SysCallN(103, uintptr(base.UnsafePointer(&event)))
}

// OutputDebugMessage
//
//	Returns true if aPath is a relative path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisrelativew">See the PathIsRelativeW article.</see>
//	Simplifies a path by removing navigation elements such as "." and ".." to produce a direct, well-formed path.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathcanonicalizew">See the PathCanonicalizeW article.</see>
//	Returns the absolute path version of aPath.
//	Tests aPath to determine if it conforms to a valid URL format.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisurlw">See the PathIsURLW article.</see>
//	Determines if aPath is a valid Universal Naming Convention (UNC) path, as opposed to a path based on a drive letter.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-pathisuncw">See the PathIsUNCW article.</see>
//	Retrieves the fully qualified path for the current module.
//	<see href="https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew">See the GetModuleFileNameW article.</see>
//	Returns true (1) if the certificate status represents an error.
//	Crash reporting is configured using an INI-style config file named
//	"crash_reporter.cfg". On Windows and Linux this file must be placed next to
//	the main application executable. On macOS this file must be placed in the
//	top-level app bundle Resources directory (e.g.
//	"<appname>.app/Contents/Resources"). File contents are as follows:
//
//	<pre>
//	# Comments start with a hash character and must be on their own line.
//
//	[Config]
//	ProductName=<Value of the "prod" crash key; defaults to "cef">
//	ProductVersion=<Value of the "ver" crash key; defaults to the CEF version>
//	AppName=<Windows only; App-specific folder name component for storing crash
//	information; default to "CEF">
//	ExternalHandler=<Windows only; Name of the external handler exe to use
//	instead of re-launching the main exe; default to empty>
//	BrowserCrashForwardingEnabled=<macOS only; True if browser process crashes
//	should be forwarded to the system crash
//	reporter; default to false>
//	ServerURL=<crash server URL; default to empty>
//	RateLimitEnabled=<True if uploads should be rate limited; default to true>
//	MaxUploadsPerDay=<Max uploads per 24 hours, used if rate limit is enabled;
//	default to 5>
//	MaxDatabaseSizeInMb=<Total crash report disk usage greater than this value
//	will cause older reports to be deleted; default to 20>
//	MaxDatabaseAgeInDays=<Crash reports older than this value will be deleted;
//	default to 5>
//
//	[CrashKeys]
//	my_key1=<small|medium|large>
//	my_key2=<small|medium|large>
//	</pre>
//
//	<b>Config section:</b>
//
//	If "ProductName" and/or "ProductVersion" are set then the specified values
//	will be included in the crash dump metadata. On macOS if these values are
//	set to NULL then they will be retrieved from the Info.plist file using the
//	"CFBundleName" and "CFBundleShortVersionString" keys respectively.
//
//	If "AppName" is set on Windows then crash report information (metrics,
//	database and dumps) will be stored locally on disk under the
//	"C:\Users\[CurrentUser]\AppData\Local\[AppName]\User Data" folder. On other
//	platforms the cef_settings_t.root_cache_path value will be used.
//
//	If "ExternalHandler" is set on Windows then the specified exe will be
//	launched as the crashpad-handler instead of re-launching the main process
//	exe. The value can be an absolute path or a path relative to the main exe
//	directory. On Linux the cef_settings_t.browser_subprocess_path value will be
//	used. On macOS the existing subprocess app bundle will be used.
//
//	If "BrowserCrashForwardingEnabled" is set to true (1) on macOS then browser
//	process crashes will be forwarded to the system crash reporter. This results
//	in the crash UI dialog being displayed to the user and crash reports being
//	logged under "~/Library/Logs/DiagnosticReports". Forwarding of crash reports
//	from non-browser processes and Debug builds is always disabled.
//
//	If "ServerURL" is set then crashes will be uploaded as a multi-part POST
//	request to the specified URL. Otherwise, reports will only be stored locally
//	on disk.
//
//	If "RateLimitEnabled" is set to true (1) then crash report uploads will be
//	rate limited as follows:
//	1. If "MaxUploadsPerDay" is set to a positive value then at most the
//	specified number of crashes will be uploaded in each 24 hour period.
//	2. If crash upload fails due to a network or server error then an
//	incremental backoff delay up to a maximum of 24 hours will be applied
//	for retries.
//	3. If a backoff delay is applied and "MaxUploadsPerDay" is > 1 then the
//	"MaxUploadsPerDay" value will be reduced to 1 until the client is
//	restarted. This helps to avoid an upload flood when the network or
//	server error is resolved.
//	Rate limiting is not supported on Linux.
//
//	If "MaxDatabaseSizeInMb" is set to a positive value then crash report
//	storage on disk will be limited to that size in megabytes. For example, on
//	Windows each dump is about 600KB so a "MaxDatabaseSizeInMb" value of 20
//	equates to about 34 crash reports stored on disk. Not supported on Linux.
//
//	If "MaxDatabaseAgeInDays" is set to a positive value then crash reports
//	older than the specified age in days will be deleted. Not supported on
//	Linux.
//
//	<b>CrashKeys section:</b>
//
//	A maximum of 26 crash keys of each size can be specified for use by the
//	application. Crash key values will be truncated based on the specified size
//	(small = 64 bytes, medium = 256 bytes, large = 1024 bytes). The value of
//	crash keys can be set from any thread or process using the
//	CefSetCrashKeyValue function. These key/value pairs will be sent to the
//	crash server along with the crash dump file.
//	Sets or clears a specific key-value pair from the crash metadata.
//	Add a log message. See the LogSeverity defines for supported |severity|
//	values.
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (cef_log)</see>
//	<see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
func (_MiscFuncClass) OutputDebugMessage(message string) {
	uCEFMiscFunctionsAPI().SysCallN(104, api.PasStr(message))
}

func (_MiscFuncClass) OutputLastErrorMessage() {
	uCEFMiscFunctionsAPI().SysCallN(105)
}

// CefGetExtensionsForMimeType
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
func (_MiscFuncClass) CefGetExtensionsForMimeType(mimeType string, extensions *lcl.IStringList) {
	extensionsPtr := base.GetObjectUintptr(*extensions)
	uCEFMiscFunctionsAPI().SysCallN(106, api.PasStr(mimeType), uintptr(base.UnsafePointer(&extensionsPtr)))
	*extensions = lcl.AsStringList(extensionsPtr)
}

// CefLoadCRLSetsFile
//
//	Combines specified |base_url| and |relative_url| into a ustring.
//	Parse the specified |url| into its component parts. Returns false (0) if the
//	URL is invalid.
//	Creates a URL from the specified |parts|, which must contain a non-NULL spec
//	or a non-NULL host and path (at a minimum), but not both.
//	This is a convenience function for formatting a URL in a concise and human-
//	friendly way to help users make security-related decisions (or in other
//	circumstances when people need to distinguish sites, origins, or otherwise-
//	simplified URLs from each other). Internationalized domain names (IDN) may
//	be presented in Unicode if the conversion is considered safe. The returned
//	value will (a) omit the path for standard schemes, excepting file and
//	filesystem, and (b) omit the port if it is the default for the scheme. Do
//	not use this for URLs which will be parsed or sent to other applications.
//	Returns the mime type for the specified file extension or an NULL string if
//	unknown.
//	Get the extensions associated with the given mime type. This should be
//	passed in lower case. There could be multiple extensions for a given mime
//	type, like "html,htm" for "text/html", or "txt,text,html,..." for "text/*".
//	Any existing elements in the provided vector will not be erased.
//	Encodes |data| as a base64 string.
//	Decodes the base64 encoded string |data|. The returned value will be NULL if
//	the decoding fails.
//	Escapes characters in |text| which are unsuitable for use as a query
//	parameter value. Everything except alphanumerics and -_.!~*'() will be
//	converted to "%XX". If |use_plus| is true (1) spaces will change to "+". The
//	result is basically the same as encodeURIComponent in Javacript.
//	Unescapes |text| and returns the result. Unescaping consists of looking for
//	the exact pattern "%XX" where each X is a hex digit and converting to the
//	character with the numerical value of those digits (e.g. "i%20=%203%3b"
//	unescapes to "i = 3;"). If |convert_to_utf8| is true (1) this function will
//	attempt to interpret the initial decoded result as UTF-8. If the result is
//	convertable into UTF-8 it will be returned as converted. Otherwise the
//	initial decoded result will be returned. The |unescape_rule| parameter
//	supports further customization the decoding process.
//	Retrieve the path associated with the specified |aPathKey|.
//	Can be called on any thread in the browser process.
//	Returns true (1) if the application text direction is right-to-left.
//	Creates a directory and all parent directories if they don't already exist.
//	Returns true (1) on successful creation or if the directory already exists.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Get the temporary directory provided by the system.
//	WARNING: In general, you should use the temp directory variants below
//	instead of this function. Those variants will ensure that the proper
//	permissions are set so that other users on the system can't edit them while
//	they're open (which could lead to security issues).
//	Creates a new directory. On Windows if |prefix| is provided the new
//	directory name is in the format of "prefixyyyy". Returns true (1) on success
//	and sets |newTempPath| to the full path of the directory that was created.
//	The directory is only readable by the current user. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Creates a directory within another directory. Extra characters will be
//	appended to |prefix| to ensure that the new directory does not have the same
//	name as an existing directory. Returns true (1) on success and sets
//	|newDir| to the full path of the directory that was created. The directory
//	is only readable by the current user. Calling this function on the browser
//	process UI or IO threads is not allowed.
//	Returns true (1) if the given path exists and is a directory. Calling this
//	function on the browser process UI or IO threads is not allowed.
//	Deletes the given path whether it's a file or a directory. If |path| is a
//	directory all contents will be deleted. If |recursive| is true (1) any sub-
//	directories and their contents will also be deleted (equivalent to executing
//	"rm -rf", so use with caution). On POSIX environments if |path| is a
//	symbolic link then only the symlink will be deleted. Returns true (1) on
//	successful deletion or if |path| does not exist. Calling this function on
//	the browser process UI or IO threads is not allowed.
//	Writes the contents of |srcDir| into a zip archive at |destFile|. If
//	|includeHiddenFiles| is true (1) files starting with "." will be included.
//	Returns true (1) on success. Calling this function on the browser process
//	UI or IO threads is not allowed.
//	Loads the existing "Certificate Revocation Lists" file that is managed by
//	Google Chrome. This file can generally be found in Chrome's User Data
//	directory (e.g. "C:\Users\[User]\AppData\Local\Google\Chrome\User Data\" on
//	Windows) and is updated periodically by Chrome's component updater service.
//	Must be called in the browser process after the context has been
//	initialized. See https://dev.chromium.org/Home/chromium-security/crlsets for
//	background.
func (_MiscFuncClass) CefLoadCRLSetsFile(path string) {
	uCEFMiscFunctionsAPI().SysCallN(107, api.PasStr(path))
}

func (_MiscFuncClass) CefCheckAltGrPressed(wparam types.WParam, event *TCefKeyEvent) {
	uCEFMiscFunctionsAPI().SysCallN(108, uintptr(wparam), uintptr(base.UnsafePointer(event)))
}

func (_MiscFuncClass) DropEffectToDragOperation(effect int32, allowedOps *cefTypes.TCefDragOperations) {
	allowedOpsPtr := uintptr(*allowedOps)
	uCEFMiscFunctionsAPI().SysCallN(109, uintptr(effect), uintptr(base.UnsafePointer(&allowedOpsPtr)))
	*allowedOps = cefTypes.TCefDragOperations(allowedOpsPtr)
}

func (_MiscFuncClass) DragOperationToDropEffect(dragOperations cefTypes.TCefDragOperations, effect *int32) {
	effectPtr := uintptr(*effect)
	uCEFMiscFunctionsAPI().SysCallN(110, uintptr(dragOperations), uintptr(base.UnsafePointer(&effectPtr)))
	*effect = int32(effectPtr)
}

func (_MiscFuncClass) DeviceToLogicalWithMouseEventDouble(event *TCefMouseEvent, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(111, uintptr(base.UnsafePointer(event)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) DeviceToLogicalWithTouchEventDouble(event *TCefTouchEvent, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(112, uintptr(base.UnsafePointer(event)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) DeviceToLogicalWithPointDouble(point *types.TPoint, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(113, uintptr(base.UnsafePointer(point)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) LogicalToDeviceWithRectDouble(rect *TCefRect, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(114, uintptr(base.UnsafePointer(rect)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

// InitializeWindowHandle
//
//	Returns a URI with a DATA scheme using |aString| as the URI's data.
//	Returns a URI with a DATA scheme encoding |aData| as a base64 string.
func (_MiscFuncClass) InitializeWindowHandle(handle *cefTypes.TCefWindowHandle) {
	handlePtr := uintptr(*handle)
	uCEFMiscFunctionsAPI().SysCallN(115, uintptr(base.UnsafePointer(&handlePtr)))
	*handle = cefTypes.TCefWindowHandle(handlePtr)
}

var (
	uCEFMiscFunctionsOnce   base.Once
	uCEFMiscFunctionsImport *imports.Imports = nil
)

func uCEFMiscFunctionsAPI() *imports.Imports {
	uCEFMiscFunctionsOnce.Do(func() {
		uCEFMiscFunctionsImport = api.NewDefaultImports()
		uCEFMiscFunctionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("uCEFMiscFunctions_CefColorGetA", 0), // static function CefColorGetA
			/* 1 */ imports.NewTable("uCEFMiscFunctions_CefColorGetR", 0), // static function CefColorGetR
			/* 2 */ imports.NewTable("uCEFMiscFunctions_CefColorGetG", 0), // static function CefColorGetG
			/* 3 */ imports.NewTable("uCEFMiscFunctions_CefColorGetB", 0), // static function CefColorGetB
			/* 4 */ imports.NewTable("uCEFMiscFunctions_CefColorSetARGB", 0), // static function CefColorSetARGB
			/* 5 */ imports.NewTable("uCEFMiscFunctions_CefGetObject", 0), // static function CefGetObject
			/* 6 */ imports.NewTable("uCEFMiscFunctions_CefRegisterExtension", 0), // static function CefRegisterExtension
			/* 7 */ imports.NewTable("uCEFMiscFunctions_CefPostTask", 0), // static function CefPostTask
			/* 8 */ imports.NewTable("uCEFMiscFunctions_CefPostDelayedTask", 0), // static function CefPostDelayedTask
			/* 9 */ imports.NewTable("uCEFMiscFunctions_CefCurrentlyOn", 0), // static function CefCurrentlyOn
			/* 10 */ imports.NewTable("uCEFMiscFunctions_FixCefTime", 0), // static function FixCefTime
			/* 11 */ imports.NewTable("uCEFMiscFunctions_CefTimeToDateTime", 0), // static function CefTimeToDateTime
			/* 12 */ imports.NewTable("uCEFMiscFunctions_DateTimeToCefTime", 0), // static function DateTimeToCefTime
			/* 13 */ imports.NewTable("uCEFMiscFunctions_DateTimeToCefBaseTime", 0), // static function DateTimeToCefBaseTime
			/* 14 */ imports.NewTable("uCEFMiscFunctions_CefTimeToDouble", 0), // static function CefTimeToDouble
			/* 15 */ imports.NewTable("uCEFMiscFunctions_DoubleToCefTime", 0), // static function DoubleToCefTime
			/* 16 */ imports.NewTable("uCEFMiscFunctions_CefTimeToUnixTime", 0), // static function CefTimeToUnixTime
			/* 17 */ imports.NewTable("uCEFMiscFunctions_UnixTimeToCefTime", 0), // static function UnixTimeToCefTime
			/* 18 */ imports.NewTable("uCEFMiscFunctions_CefTimeNow", 0), // static function CefTimeNow
			/* 19 */ imports.NewTable("uCEFMiscFunctions_DoubleTimeNow", 0), // static function DoubleTimeNow
			/* 20 */ imports.NewTable("uCEFMiscFunctions_CefBaseTimeNow", 0), // static function CefBaseTimeNow
			/* 21 */ imports.NewTable("uCEFMiscFunctions_CetTimeToCefBaseTime", 0), // static function CetTimeToCefBaseTime
			/* 22 */ imports.NewTable("uCEFMiscFunctions_CetTimeFromCefBaseTime", 0), // static function CetTimeFromCefBaseTime
			/* 23 */ imports.NewTable("uCEFMiscFunctions_CefBaseTimeToDateTime", 0), // static function CefBaseTimeToDateTime
			/* 24 */ imports.NewTable("uCEFMiscFunctions_CustomPathIsRelative", 0), // static function CustomPathIsRelative
			/* 25 */ imports.NewTable("uCEFMiscFunctions_CustomPathCanonicalize", 0), // static function CustomPathCanonicalize
			/* 26 */ imports.NewTable("uCEFMiscFunctions_CustomAbsolutePath", 0), // static function CustomAbsolutePath
			/* 27 */ imports.NewTable("uCEFMiscFunctions_CustomPathIsURL", 0), // static function CustomPathIsURL
			/* 28 */ imports.NewTable("uCEFMiscFunctions_CustomPathIsUNC", 0), // static function CustomPathIsUNC
			/* 29 */ imports.NewTable("uCEFMiscFunctions_GetModulePath", 0), // static function GetModulePath
			/* 30 */ imports.NewTable("uCEFMiscFunctions_CefIsCertStatusError", 0), // static function CefIsCertStatusError
			/* 31 */ imports.NewTable("uCEFMiscFunctions_CefCrashReportingEnabled", 0), // static function CefCrashReportingEnabled
			/* 32 */ imports.NewTable("uCEFMiscFunctions_CefGetMinLogLevel", 0), // static function CefGetMinLogLevel
			/* 33 */ imports.NewTable("uCEFMiscFunctions_CefGetLogSeverityName", 0), // static function CefGetLogSeverityName
			/* 34 */ imports.NewTable("uCEFMiscFunctions_CefRegisterSchemeHandlerFactory", 0), // static function CefRegisterSchemeHandlerFactory
			/* 35 */ imports.NewTable("uCEFMiscFunctions_CefClearSchemeHandlerFactories", 0), // static function CefClearSchemeHandlerFactories
			/* 36 */ imports.NewTable("uCEFMiscFunctions_CefAddCrossOriginWhitelistEntry", 0), // static function CefAddCrossOriginWhitelistEntry
			/* 37 */ imports.NewTable("uCEFMiscFunctions_CefRemoveCrossOriginWhitelistEntry", 0), // static function CefRemoveCrossOriginWhitelistEntry
			/* 38 */ imports.NewTable("uCEFMiscFunctions_CefClearCrossOriginWhitelist", 0), // static function CefClearCrossOriginWhitelist
			/* 39 */ imports.NewTable("uCEFMiscFunctions_GetExtendedFileVersion", 0), // static function GetExtendedFileVersion
			/* 40 */ imports.NewTable("uCEFMiscFunctions_GetRegistryWindowsVersion", 0), // static function GetRegistryWindowsVersion
			/* 41 */ imports.NewTable("uCEFMiscFunctions_GetRealWindowsVersion", 0), // static function GetRealWindowsVersion
			/* 42 */ imports.NewTable("uCEFMiscFunctions_CheckRealWindowsVersion", 0), // static function CheckRealWindowsVersion
			/* 43 */ imports.NewTable("uCEFMiscFunctions_SplitLongString", 0), // static function SplitLongString
			/* 44 */ imports.NewTable("uCEFMiscFunctions_GetAbsoluteDirPath", 0), // static function GetAbsoluteDirPath
			/* 45 */ imports.NewTable("uCEFMiscFunctions_CheckSubprocessPath", 0), // static function CheckSubprocessPath
			/* 46 */ imports.NewTable("uCEFMiscFunctions_CheckLocales", 0), // static function CheckLocales
			/* 47 */ imports.NewTable("uCEFMiscFunctions_CheckResources", 0), // static function CheckResources
			/* 48 */ imports.NewTable("uCEFMiscFunctions_CheckDLLs", 0), // static function CheckDLLs
			/* 49 */ imports.NewTable("uCEFMiscFunctions_CheckDLLVersion", 0), // static function CheckDLLVersion
			/* 50 */ imports.NewTable("uCEFMiscFunctions_GetDLLHeaderMachine", 0), // static function GetDLLHeaderMachine
			/* 51 */ imports.NewTable("uCEFMiscFunctions_GetFileTypeDescription", 0), // static function GetFileTypeDescription
			/* 52 */ imports.NewTable("uCEFMiscFunctions_CheckFilesExist", 0), // static function CheckFilesExist
			/* 53 */ imports.NewTable("uCEFMiscFunctions_Is32BitProcess", 0), // static function Is32BitProcess
			/* 54 */ imports.NewTable("uCEFMiscFunctions_CefFormatUrlForSecurityDisplay", 0), // static function CefFormatUrlForSecurityDisplay
			/* 55 */ imports.NewTable("uCEFMiscFunctions_CefGetMimeType", 0), // static function CefGetMimeType
			/* 56 */ imports.NewTable("uCEFMiscFunctions_CefBase64Encode", 0), // static function CefBase64Encode
			/* 57 */ imports.NewTable("uCEFMiscFunctions_CefBase64Decode", 0), // static function CefBase64Decode
			/* 58 */ imports.NewTable("uCEFMiscFunctions_CefUriEncode", 0), // static function CefUriEncode
			/* 59 */ imports.NewTable("uCEFMiscFunctions_CefUriDecode", 0), // static function CefUriDecode
			/* 60 */ imports.NewTable("uCEFMiscFunctions_CefGetPath", 0), // static function CefGetPath
			/* 61 */ imports.NewTable("uCEFMiscFunctions_CefIsRTL", 0), // static function CefIsRTL
			/* 62 */ imports.NewTable("uCEFMiscFunctions_CefCreateDirectory", 0), // static function CefCreateDirectory
			/* 63 */ imports.NewTable("uCEFMiscFunctions_CefGetTempDirectory", 0), // static function CefGetTempDirectory
			/* 64 */ imports.NewTable("uCEFMiscFunctions_CefCreateNewTempDirectory", 0), // static function CefCreateNewTempDirectory
			/* 65 */ imports.NewTable("uCEFMiscFunctions_CefCreateTempDirectoryInDirectory", 0), // static function CefCreateTempDirectoryInDirectory
			/* 66 */ imports.NewTable("uCEFMiscFunctions_CefDirectoryExists", 0), // static function CefDirectoryExists
			/* 67 */ imports.NewTable("uCEFMiscFunctions_CefDeleteFile", 0), // static function CefDeleteFile
			/* 68 */ imports.NewTable("uCEFMiscFunctions_CefZipDirectory", 0), // static function CefZipDirectory
			/* 69 */ imports.NewTable("uCEFMiscFunctions_GetDefaultCEFUserAgent", 0), // static function GetDefaultCEFUserAgent
			/* 70 */ imports.NewTable("uCEFMiscFunctions_CefIsKeyDown", 0), // static function CefIsKeyDown
			/* 71 */ imports.NewTable("uCEFMiscFunctions_CefIsKeyToggled", 0), // static function CefIsKeyToggled
			/* 72 */ imports.NewTable("uCEFMiscFunctions_GetCefMouseModifiersToEventFlags", 0), // static function GetCefMouseModifiersToEventFlags
			/* 73 */ imports.NewTable("uCEFMiscFunctions_GetCefMouseModifiersWithWPARAM", 0), // static function GetCefMouseModifiersWithWPARAM
			/* 74 */ imports.NewTable("uCEFMiscFunctions_GetCefKeyboardModifiers", 0), // static function GetCefKeyboardModifiers
			/* 75 */ imports.NewTable("uCEFMiscFunctions_GetWindowsMajorMinorVersion", 0), // static function GetWindowsMajorMinorVersion
			/* 76 */ imports.NewTable("uCEFMiscFunctions_GetIsWow64Process2", 0), // static function GetIsWow64Process2
			/* 77 */ imports.NewTable("uCEFMiscFunctions_IsWowProcess", 0), // static function IsWowProcess
			/* 78 */ imports.NewTable("uCEFMiscFunctions_RunningWindows10OrNewer", 0), // static function RunningWindows10OrNewer
			/* 79 */ imports.NewTable("uCEFMiscFunctions_GetDPIForHandle", 0), // static function GetDPIForHandle
			/* 80 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithIntDouble", 0), // static function DeviceToLogicalWithIntDouble
			/* 81 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithSingleDouble", 0), // static function DeviceToLogicalWithSingleDouble
			/* 82 */ imports.NewTable("uCEFMiscFunctions_LogicalToDeviceWithIntDouble", 0), // static function LogicalToDeviceWithIntDouble
			/* 83 */ imports.NewTable("uCEFMiscFunctions_GetScreenDPI", 0), // static function GetScreenDPI
			/* 84 */ imports.NewTable("uCEFMiscFunctions_GetDeviceScaleFactor", 0), // static function GetDeviceScaleFactor
			/* 85 */ imports.NewTable("uCEFMiscFunctions_TryRemoveDir", 0), // static function TryRemoveDir
			/* 86 */ imports.NewTable("uCEFMiscFunctions_TryDeleteFile", 0), // static function TryDeleteFile
			/* 87 */ imports.NewTable("uCEFMiscFunctions_TryRenameDir", 0), // static function TryRenameDir
			/* 88 */ imports.NewTable("uCEFMiscFunctions_TryRenameFile", 0), // static function TryRenameFile
			/* 89 */ imports.NewTable("uCEFMiscFunctions_DeleteDirContents", 0), // static function DeleteDirContents
			/* 90 */ imports.NewTable("uCEFMiscFunctions_DeleteFileList", 0), // static function DeleteFileList
			/* 91 */ imports.NewTable("uCEFMiscFunctions_MoveFileList", 0), // static function MoveFileList
			/* 92 */ imports.NewTable("uCEFMiscFunctions_CefGetDataURIWithStrX2", 0), // static function CefGetDataURIWithStrX2
			/* 93 */ imports.NewTable("uCEFMiscFunctions_CefGetDataURIWithPointerIntStrX2", 0), // static function CefGetDataURIWithPointerIntStrX2
			/* 94 */ imports.NewTable("uCEFMiscFunctions_ValidCefWindowHandle", 0), // static function ValidCefWindowHandle
			/* 95 */ imports.NewTable("uCEFMiscFunctions_GetCommandLineSwitchValue", 0), // static function GetCommandLineSwitchValue
			/* 96 */ imports.NewTable("uCEFMiscFunctions_IsCEFSubprocess", 0), // static function IsCEFSubprocess
			/* 97 */ imports.NewTable("uCEFMiscFunctions_EditingCommandToString", 0), // static function EditingCommandToString
			/* 98 */ imports.NewTable("uCEFMiscFunctions_CefResultCodeToString", 0), // static function CefResultCodeToString
			/* 99 */ imports.NewTable("uCEFMiscFunctions_InitializeCefTime", 0), // static procedure InitializeCefTime
			/* 100 */ imports.NewTable("uCEFMiscFunctions_CefSetCrashKeyValue", 0), // static procedure CefSetCrashKeyValue
			/* 101 */ imports.NewTable("uCEFMiscFunctions_CefLog", 0), // static procedure CefLog
			/* 102 */ imports.NewTable("uCEFMiscFunctions_CefKeyEventLog", 0), // static procedure CefKeyEventLog
			/* 103 */ imports.NewTable("uCEFMiscFunctions_CefMouseEventLog", 0), // static procedure CefMouseEventLog
			/* 104 */ imports.NewTable("uCEFMiscFunctions_OutputDebugMessage", 0), // static procedure OutputDebugMessage
			/* 105 */ imports.NewTable("uCEFMiscFunctions_OutputLastErrorMessage", 0), // static procedure OutputLastErrorMessage
			/* 106 */ imports.NewTable("uCEFMiscFunctions_CefGetExtensionsForMimeType", 0), // static procedure CefGetExtensionsForMimeType
			/* 107 */ imports.NewTable("uCEFMiscFunctions_CefLoadCRLSetsFile", 0), // static procedure CefLoadCRLSetsFile
			/* 108 */ imports.NewTable("uCEFMiscFunctions_CefCheckAltGrPressed", 0), // static procedure CefCheckAltGrPressed
			/* 109 */ imports.NewTable("uCEFMiscFunctions_DropEffectToDragOperation", 0), // static procedure DropEffectToDragOperation
			/* 110 */ imports.NewTable("uCEFMiscFunctions_DragOperationToDropEffect", 0), // static procedure DragOperationToDropEffect
			/* 111 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithMouseEventDouble", 0), // static procedure DeviceToLogicalWithMouseEventDouble
			/* 112 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithTouchEventDouble", 0), // static procedure DeviceToLogicalWithTouchEventDouble
			/* 113 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithPointDouble", 0), // static procedure DeviceToLogicalWithPointDouble
			/* 114 */ imports.NewTable("uCEFMiscFunctions_LogicalToDeviceWithRectDouble", 0), // static procedure LogicalToDeviceWithRectDouble
			/* 115 */ imports.NewTable("uCEFMiscFunctions_InitializeWindowHandle", 0), // static procedure InitializeWindowHandle
		}
	})
	return uCEFMiscFunctionsImport
}
