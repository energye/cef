//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// *********************************************************
// ********************** ATTENTION ! **********************
// *********************************************************
// **                                                     **
// **  MANY OF THE EVENTS IN CEF4DELPHI COMPONENTS LIKE   **
// **  TCHROMIUM, TFMXCHROMIUM OR TApplication ARE     **
// **  EXECUTED IN A CEF THREAD BY DEFAULT.               **
// **                                                     **
// **  WINDOWS CONTROLS MUST BE CREATED AND DESTROYED IN  **
// **  THE SAME THREAD TO AVOID ERRORS.                   **
// **  SOME OF THEM RECREATE THE HANDLERS IF THEY ARE     **
// **  MODIFIED AND CAN CAUSE THE SAME ERRORS.            **
// **                                                     **
// **  DON'T CREATE, MODIFY OR DESTROY WINDOWS CONTROLS   **
// **  INSIDE THE CEF4DELPHI EVENTS AND USE               **
// **  SYNCHRONIZATION OBJECTS TO PROTECT VARIABLES AND   **
// **  FIELDS IF THEY ARE ALSO USED IN THE MAIN THREAD.   **
// **                                                     **
// **  READ THIS FOR MORE INFORMATION :                   **
// **  https://www.briskbard.com/index.php?pageid=cef     **
// **                                                     **
// **  USE OUR FORUMS FOR MORE QUESTIONS :                **
// **  https://www.briskbard.com/forum/                   **
// **                                                     **
// *********************************************************
// *********************************************************

var globalApp ICefApplication

// SetGlobalCEFApp
//
//	设置全局 ICefApplication 实例，单例模式, 应在 NewCefApplication 后设置
func SetGlobalCEFApp(application ICefApplication) {
	if globalApp == nil {
		predefImportAPI().SysCallN(12, application.Instance())
		globalApp = application
	}
}

// DestroyGlobalCEFApp 销毁全局App实例
func DestroyGlobalCEFApp() {
	if globalApp != nil {
		predefImportAPI().SysCallN(13)
		globalApp.SetInstance(nil)
		globalApp = nil
	}
}

func GlobalCEFApp() ICefApplication {
	return globalApp
}

func GetSpecificVersion() SpecificVersion {
	r1 := predefImportAPI().SysCallN(14)
	return SpecificVersion(r1)
}
