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

	cefTypes "github.com/energye/cef/109/types"
)

// MiscFunc  is static instance
var MiscFunc _MiscFuncClass

// _MiscFuncClass is class type defined by MiscFunc
type _MiscFuncClass uintptr

func (_MiscFuncClass) CefColorGetA(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(0, uintptr(color))
	return byte(r)
}

func (_MiscFuncClass) CefColorGetR(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(1, uintptr(color))
	return byte(r)
}

func (_MiscFuncClass) CefColorGetG(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(2, uintptr(color))
	return byte(r)
}

func (_MiscFuncClass) CefColorGetB(color cefTypes.TCefColor) byte {
	r := uCEFMiscFunctionsAPI().SysCallN(3, uintptr(color))
	return byte(r)
}

func (_MiscFuncClass) CefColorSetARGB(A byte, R byte, G byte, B byte) cefTypes.TCefColor {
	r := uCEFMiscFunctionsAPI().SysCallN(4, uintptr(A), uintptr(R), uintptr(G), uintptr(B))
	return cefTypes.TCefColor(r)
}

func (_MiscFuncClass) CefGetObject(ptr uintptr) lcl.IObject {
	r := uCEFMiscFunctionsAPI().SysCallN(5, uintptr(ptr))
	return lcl.AsObject(r)
}

func (_MiscFuncClass) CefRegisterExtension(name string, code string, handler IEngV8Handler) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(6, api.PasStr(name), api.PasStr(code), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefPostTask(threadId cefTypes.TCefThreadId, task IEngTask) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(7, uintptr(threadId), base.GetObjectUintptr(task))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefPostDelayedTask(threadId cefTypes.TCefThreadId, task IEngTask, delayMs int64) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(8, uintptr(threadId), base.GetObjectUintptr(task), uintptr(base.UnsafePointer(&delayMs)))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefCurrentlyOn(threadId cefTypes.TCefThreadId) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(9, uintptr(threadId))
	return api.GoBool(r)
}

func (_MiscFuncClass) FixCefTime(dt TCefTime) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(10, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefTimeToDateTime(dt TCefTime) (result types.TDateTime) {
	uCEFMiscFunctionsAPI().SysCallN(11, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) DateTimeToCefTime(dt types.TDateTime) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(12, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) DateTimeToCefBaseTime(dt types.TDateTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(13, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefTimeToDouble(dt TCefTime) (result float64) {
	uCEFMiscFunctionsAPI().SysCallN(14, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) DoubleToCefTime(dt float64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(15, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefTimeToUnixTime(dt TCefTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(16, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) UnixTimeToCefTime(dt int64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(17, uintptr(base.UnsafePointer(&dt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefTimeNow() (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(18, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) DoubleTimeNow() (result float64) {
	uCEFMiscFunctionsAPI().SysCallN(19, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefBaseTimeNow() (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(20, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CetTimeToCefBaseTime(ct TCefTime) (result int64) {
	uCEFMiscFunctionsAPI().SysCallN(21, uintptr(base.UnsafePointer(&ct)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CetTimeFromCefBaseTime(cbt int64) (result TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(22, uintptr(base.UnsafePointer(&cbt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CefBaseTimeToDateTime(cbt int64) (result types.TDateTime) {
	uCEFMiscFunctionsAPI().SysCallN(23, uintptr(base.UnsafePointer(&cbt)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) CustomPathIsRelative(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(24, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomPathCanonicalize(originalPath string, canonicalPath *string) bool {
	canonicalPathPtr := api.PasStr(*canonicalPath)
	r := uCEFMiscFunctionsAPI().SysCallN(25, api.PasStr(originalPath), uintptr(base.UnsafePointer(&canonicalPathPtr)))
	*canonicalPath = api.GoStr(canonicalPathPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomAbsolutePath(path string, mustExist bool) string {
	r := uCEFMiscFunctionsAPI().SysCallN(26, api.PasStr(path), api.PasBool(mustExist))
	return api.GoStr(r)
}

func (_MiscFuncClass) CustomPathIsURL(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(27, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomPathIsUNC(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(28, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetModulePath() string {
	r := uCEFMiscFunctionsAPI().SysCallN(29)
	return api.GoStr(r)
}

func (_MiscFuncClass) CefIsCertStatusError(status cefTypes.TCefCertStatus) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(30, uintptr(status))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefCrashReportingEnabled() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(31)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefRegisterSchemeHandlerFactory(schemeName string, domainName string, handler cefTypes.TCefResourceHandlerClass) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(32, api.PasStr(schemeName), api.PasStr(domainName), uintptr(handler))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefClearSchemeHandlerFactories() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(33)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefAddCrossOriginWhitelistEntry(sourceOrigin string, targetProtocol string, targetDomain string, allowTargetSubdomains bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(34, api.PasStr(sourceOrigin), api.PasStr(targetProtocol), api.PasStr(targetDomain), api.PasBool(allowTargetSubdomains))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefRemoveCrossOriginWhitelistEntry(sourceOrigin string, targetProtocol string, targetDomain string, allowTargetSubdomains bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(35, api.PasStr(sourceOrigin), api.PasStr(targetProtocol), api.PasStr(targetDomain), api.PasBool(allowTargetSubdomains))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefClearCrossOriginWhitelist() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(36)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetExtendedFileVersion(fileName string) (result uint64) {
	uCEFMiscFunctionsAPI().SysCallN(37, api.PasStr(fileName), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) SplitLongString(srcString string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(38, api.PasStr(srcString))
	return api.GoStr(r)
}

func (_MiscFuncClass) GetAbsoluteDirPath(srcPath string, rsltPath *string) bool {
	rsltPathPtr := api.PasStr(*rsltPath)
	r := uCEFMiscFunctionsAPI().SysCallN(39, api.PasStr(srcPath), uintptr(base.UnsafePointer(&rsltPathPtr)))
	*rsltPath = api.GoStr(rsltPathPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckSubprocessPath(subprocessPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(40, api.PasStr(subprocessPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckLocales(localesDirPath string, missingFiles *string, localesRequired string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(41, api.PasStr(localesDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)), api.PasStr(localesRequired))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckResources(resourcesDirPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(42, api.PasStr(resourcesDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckDLLs(frameworkDirPath string, missingFiles *string) bool {
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(43, api.PasStr(frameworkDirPath), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CheckDLLVersion(dLLFile string, major uint16, minor uint16, release uint16, build uint16) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(44, api.PasStr(dLLFile), uintptr(major), uintptr(minor), uintptr(release), uintptr(build))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetDLLHeaderMachine(dLLFile string, machine *int32) bool {
	machinePtr := uintptr(*machine)
	r := uCEFMiscFunctionsAPI().SysCallN(45, api.PasStr(dLLFile), uintptr(base.UnsafePointer(&machinePtr)))
	*machine = int32(machinePtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetFileTypeDescription(extension string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(46, api.PasStr(extension))
	return api.GoStr(r)
}

func (_MiscFuncClass) CheckFilesExist(list *lcl.IStringList, missingFiles *string) bool {
	listPtr := base.GetObjectUintptr(*list)
	missingFilesPtr := api.PasStr(*missingFiles)
	r := uCEFMiscFunctionsAPI().SysCallN(47, uintptr(base.UnsafePointer(&listPtr)), uintptr(base.UnsafePointer(&missingFilesPtr)))
	*list = lcl.AsStringList(listPtr)
	*missingFiles = api.GoStr(missingFilesPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) Is32BitProcess() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(48)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefFormatUrlForSecurityDisplay(originUrl string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(49, api.PasStr(originUrl))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefGetMimeType(extension string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(50, api.PasStr(extension))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefBase64Encode(data uintptr, dataSize cefTypes.NativeUInt) string {
	r := uCEFMiscFunctionsAPI().SysCallN(51, uintptr(data), uintptr(dataSize))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefBase64Decode(data string) (result ICefBinaryValue) {
	var resultPtr uintptr
	uCEFMiscFunctionsAPI().SysCallN(52, api.PasStr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (_MiscFuncClass) CefUriEncode(text string, usePlus bool) string {
	r := uCEFMiscFunctionsAPI().SysCallN(53, api.PasStr(text), api.PasBool(usePlus))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefUriDecode(text string, convertToUtf8 bool, unescapeRule cefTypes.TCefUriUnescapeRule) string {
	r := uCEFMiscFunctionsAPI().SysCallN(54, api.PasStr(text), api.PasBool(convertToUtf8), uintptr(unescapeRule))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefGetPath(pathKey cefTypes.TCefPathKey) string {
	r := uCEFMiscFunctionsAPI().SysCallN(55, uintptr(pathKey))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefIsRTL() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(56)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefCreateDirectory(fullPath string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(57, api.PasStr(fullPath))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefGetTempDirectory(outTempDir *string) bool {
	var tempDirPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(58, uintptr(base.UnsafePointer(&tempDirPtr)))
	*outTempDir = api.GoStr(tempDirPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefCreateNewTempDirectory(prefix string, outNewTempPath *string) bool {
	var newTempPathPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(59, api.PasStr(prefix), uintptr(base.UnsafePointer(&newTempPathPtr)))
	*outNewTempPath = api.GoStr(newTempPathPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefCreateTempDirectoryInDirectory(baseDir string, prefix string, outNewDir *string) bool {
	var newDirPtr uintptr
	r := uCEFMiscFunctionsAPI().SysCallN(60, api.PasStr(baseDir), api.PasStr(prefix), uintptr(base.UnsafePointer(&newDirPtr)))
	*outNewDir = api.GoStr(newDirPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CefDirectoryExists(path string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(61, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefDeleteFile(path string, recursive bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(62, api.PasStr(path), api.PasBool(recursive))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefZipDirectory(srcDir string, destFile string, includeHiddenFiles bool) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(63, api.PasStr(srcDir), api.PasStr(destFile), api.PasBool(includeHiddenFiles))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefIsKeyDown(wparam types.WParam) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(64, uintptr(wparam))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefIsKeyToggled(wparam types.WParam) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(65, uintptr(wparam))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetCefMouseModifiersToEventFlags() cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(66)
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetCefMouseModifiersWithWPARAM(awparam types.WParam) cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(67, uintptr(awparam))
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetCefKeyboardModifiers(wparam types.WParam, lparam types.LParam) cefTypes.TCefEventFlags {
	r := uCEFMiscFunctionsAPI().SysCallN(68, uintptr(wparam), uintptr(lparam))
	return cefTypes.TCefEventFlags(r)
}

func (_MiscFuncClass) GetWindowsMajorMinorVersion(wMajorVersion *types.DWORD, wMinorVersion *types.DWORD) bool {
	wMajorVersionPtr := uintptr(*wMajorVersion)
	wMinorVersionPtr := uintptr(*wMinorVersion)
	r := uCEFMiscFunctionsAPI().SysCallN(69, uintptr(base.UnsafePointer(&wMajorVersionPtr)), uintptr(base.UnsafePointer(&wMinorVersionPtr)))
	*wMajorVersion = types.DWORD(wMajorVersionPtr)
	*wMinorVersion = types.DWORD(wMinorVersionPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) RunningWindows10OrNewer() bool {
	r := uCEFMiscFunctionsAPI().SysCallN(70)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetDPIForHandle(handle types.HWND, dPI *types.UINT) bool {
	dPIPtr := uintptr(*dPI)
	r := uCEFMiscFunctionsAPI().SysCallN(71, uintptr(handle), uintptr(base.UnsafePointer(&dPIPtr)))
	*dPI = types.UINT(dPIPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) GetDefaultCEFUserAgent() string {
	r := uCEFMiscFunctionsAPI().SysCallN(72)
	return api.GoStr(r)
}

func (_MiscFuncClass) DeviceToLogicalWithIntDouble(value int32, deviceScaleFactor float64) int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(73, uintptr(value), uintptr(base.UnsafePointer(&deviceScaleFactor)))
	return int32(r)
}

func (_MiscFuncClass) DeviceToLogicalWithSingleDouble(value float32, deviceScaleFactor float64) (result float32) {
	uCEFMiscFunctionsAPI().SysCallN(74, uintptr(base.UnsafePointer(&value)), uintptr(base.UnsafePointer(&deviceScaleFactor)), uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) LogicalToDeviceWithIntDouble(value int32, deviceScaleFactor float64) int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(75, uintptr(value), uintptr(base.UnsafePointer(&deviceScaleFactor)))
	return int32(r)
}

func (_MiscFuncClass) GetScreenDPI() int32 {
	r := uCEFMiscFunctionsAPI().SysCallN(76)
	return int32(r)
}

func (_MiscFuncClass) GetDeviceScaleFactor() (result float32) {
	uCEFMiscFunctionsAPI().SysCallN(77, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) DeleteDirContents(directory string, excludeFiles lcl.IStringList) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(78, api.PasStr(directory), base.GetObjectUintptr(excludeFiles))
	return api.GoBool(r)
}

func (_MiscFuncClass) DeleteFileList(fileList lcl.IStringList) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(79, base.GetObjectUintptr(fileList))
	return api.GoBool(r)
}

func (_MiscFuncClass) MoveFileList(fileList lcl.IStringList, srcDirectory string, dstDirectory string) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(80, base.GetObjectUintptr(fileList), api.PasStr(srcDirectory), api.PasStr(dstDirectory))
	return api.GoBool(r)
}

func (_MiscFuncClass) CefGetDataURIWithStrX2(string string, mimeType string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(81, api.PasStr(string), api.PasStr(mimeType))
	return api.GoStr(r)
}

func (_MiscFuncClass) CefGetDataURIWithPointerIntStrX2(data uintptr, size int32, mimeType string, charset string) string {
	r := uCEFMiscFunctionsAPI().SysCallN(82, uintptr(data), uintptr(size), api.PasStr(mimeType), api.PasStr(charset))
	return api.GoStr(r)
}

func (_MiscFuncClass) ValidCefWindowHandle(handle cefTypes.TCefWindowHandle) bool {
	r := uCEFMiscFunctionsAPI().SysCallN(83, uintptr(handle))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetCommandLineSwitchValue(key string, value *string) bool {
	valuePtr := api.PasStr(*value)
	r := uCEFMiscFunctionsAPI().SysCallN(84, api.PasStr(key), uintptr(base.UnsafePointer(&valuePtr)))
	*value = api.GoStr(valuePtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) InitializeCefTime(time *TCefTime) {
	uCEFMiscFunctionsAPI().SysCallN(85, uintptr(base.UnsafePointer(time)))
}

func (_MiscFuncClass) WindowInfoAsChildWithWInfoWHandleRectStrDWORD(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, rect types.TRect, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(86, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), uintptr(base.UnsafePointer(&rect)), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsPopUpWithWInfoWHandleStrDWORD(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(87, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsWindowlessWithWInfoWHandleStrDWORD(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(88, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsChildWithWInfoWHandleRectStrBool(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, rect types.TRect, windowName string, hidden bool) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(89, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), uintptr(base.UnsafePointer(&rect)), api.PasStr(windowName), api.PasBool(hidden))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsPopUpWithWInfoWHandleStrBool(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string, hidden bool) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(90, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName), api.PasBool(hidden))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsWindowlessWithWInfoWHandleStrBool(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string, hidden bool) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(91, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName), api.PasBool(hidden))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsChildWithWInfoWHandleRectStr(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, rect types.TRect, windowName string) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(92, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), uintptr(base.UnsafePointer(&rect)), api.PasStr(windowName))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsPopUpWithWInfoWHandleStr(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(93, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) WindowInfoAsWindowlessWithWInfoWHandleStr(windowInfo *TCefWindowInfo, parent cefTypes.TCefWindowHandle, windowName string) {
	windowInfoPtr := windowInfo.ToPas()
	uCEFMiscFunctionsAPI().SysCallN(94, uintptr(base.UnsafePointer(windowInfoPtr)), uintptr(parent), api.PasStr(windowName))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) CefSetCrashKeyValue(key string, value string) {
	uCEFMiscFunctionsAPI().SysCallN(95, api.PasStr(key), api.PasStr(value))
}

func (_MiscFuncClass) CefLog(file string, line int32, severity int32, message string) {
	uCEFMiscFunctionsAPI().SysCallN(96, api.PasStr(file), uintptr(line), uintptr(severity), api.PasStr(message))
}

func (_MiscFuncClass) CefKeyEventLog(event TCefKeyEvent) {
	uCEFMiscFunctionsAPI().SysCallN(97, uintptr(base.UnsafePointer(&event)))
}

func (_MiscFuncClass) CefMouseEventLog(event TCefMouseEvent) {
	uCEFMiscFunctionsAPI().SysCallN(98, uintptr(base.UnsafePointer(&event)))
}

func (_MiscFuncClass) OutputDebugMessage(message string) {
	uCEFMiscFunctionsAPI().SysCallN(99, api.PasStr(message))
}

func (_MiscFuncClass) OutputLastErrorMessage() {
	uCEFMiscFunctionsAPI().SysCallN(100)
}

func (_MiscFuncClass) CefGetExtensionsForMimeType(mimeType string, extensions *lcl.IStringList) {
	extensionsPtr := base.GetObjectUintptr(*extensions)
	uCEFMiscFunctionsAPI().SysCallN(101, api.PasStr(mimeType), uintptr(base.UnsafePointer(&extensionsPtr)))
	*extensions = lcl.AsStringList(extensionsPtr)
}

func (_MiscFuncClass) CefLoadCRLSetsFile(path string) {
	uCEFMiscFunctionsAPI().SysCallN(102, api.PasStr(path))
}

func (_MiscFuncClass) CefCheckAltGrPressed(wparam types.WParam, event *TCefKeyEvent) {
	uCEFMiscFunctionsAPI().SysCallN(103, uintptr(wparam), uintptr(base.UnsafePointer(event)))
}

func (_MiscFuncClass) DropEffectToDragOperation(effect int32, allowedOps *cefTypes.TCefDragOperations) {
	allowedOpsPtr := uintptr(*allowedOps)
	uCEFMiscFunctionsAPI().SysCallN(104, uintptr(effect), uintptr(base.UnsafePointer(&allowedOpsPtr)))
	*allowedOps = cefTypes.TCefDragOperations(allowedOpsPtr)
}

func (_MiscFuncClass) DragOperationToDropEffect(dragOperations cefTypes.TCefDragOperations, effect *int32) {
	effectPtr := uintptr(*effect)
	uCEFMiscFunctionsAPI().SysCallN(105, uintptr(dragOperations), uintptr(base.UnsafePointer(&effectPtr)))
	*effect = int32(effectPtr)
}

func (_MiscFuncClass) DeviceToLogicalWithMouseEventDouble(event *TCefMouseEvent, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(106, uintptr(base.UnsafePointer(event)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) DeviceToLogicalWithTouchEventDouble(event *TCefTouchEvent, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(107, uintptr(base.UnsafePointer(event)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) DeviceToLogicalWithPointDouble(point *types.TPoint, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(108, uintptr(base.UnsafePointer(point)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) LogicalToDeviceWithRectDouble(rect *TCefRect, deviceScaleFactor float64) {
	uCEFMiscFunctionsAPI().SysCallN(109, uintptr(base.UnsafePointer(rect)), uintptr(base.UnsafePointer(&deviceScaleFactor)))
}

func (_MiscFuncClass) InitializeWindowHandle(handle *cefTypes.TCefWindowHandle) {
	handlePtr := uintptr(*handle)
	uCEFMiscFunctionsAPI().SysCallN(110, uintptr(base.UnsafePointer(&handlePtr)))
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
			/* 32 */ imports.NewTable("uCEFMiscFunctions_CefRegisterSchemeHandlerFactory", 0), // static function CefRegisterSchemeHandlerFactory
			/* 33 */ imports.NewTable("uCEFMiscFunctions_CefClearSchemeHandlerFactories", 0), // static function CefClearSchemeHandlerFactories
			/* 34 */ imports.NewTable("uCEFMiscFunctions_CefAddCrossOriginWhitelistEntry", 0), // static function CefAddCrossOriginWhitelistEntry
			/* 35 */ imports.NewTable("uCEFMiscFunctions_CefRemoveCrossOriginWhitelistEntry", 0), // static function CefRemoveCrossOriginWhitelistEntry
			/* 36 */ imports.NewTable("uCEFMiscFunctions_CefClearCrossOriginWhitelist", 0), // static function CefClearCrossOriginWhitelist
			/* 37 */ imports.NewTable("uCEFMiscFunctions_GetExtendedFileVersion", 0), // static function GetExtendedFileVersion
			/* 38 */ imports.NewTable("uCEFMiscFunctions_SplitLongString", 0), // static function SplitLongString
			/* 39 */ imports.NewTable("uCEFMiscFunctions_GetAbsoluteDirPath", 0), // static function GetAbsoluteDirPath
			/* 40 */ imports.NewTable("uCEFMiscFunctions_CheckSubprocessPath", 0), // static function CheckSubprocessPath
			/* 41 */ imports.NewTable("uCEFMiscFunctions_CheckLocales", 0), // static function CheckLocales
			/* 42 */ imports.NewTable("uCEFMiscFunctions_CheckResources", 0), // static function CheckResources
			/* 43 */ imports.NewTable("uCEFMiscFunctions_CheckDLLs", 0), // static function CheckDLLs
			/* 44 */ imports.NewTable("uCEFMiscFunctions_CheckDLLVersion", 0), // static function CheckDLLVersion
			/* 45 */ imports.NewTable("uCEFMiscFunctions_GetDLLHeaderMachine", 0), // static function GetDLLHeaderMachine
			/* 46 */ imports.NewTable("uCEFMiscFunctions_GetFileTypeDescription", 0), // static function GetFileTypeDescription
			/* 47 */ imports.NewTable("uCEFMiscFunctions_CheckFilesExist", 0), // static function CheckFilesExist
			/* 48 */ imports.NewTable("uCEFMiscFunctions_Is32BitProcess", 0), // static function Is32BitProcess
			/* 49 */ imports.NewTable("uCEFMiscFunctions_CefFormatUrlForSecurityDisplay", 0), // static function CefFormatUrlForSecurityDisplay
			/* 50 */ imports.NewTable("uCEFMiscFunctions_CefGetMimeType", 0), // static function CefGetMimeType
			/* 51 */ imports.NewTable("uCEFMiscFunctions_CefBase64Encode", 0), // static function CefBase64Encode
			/* 52 */ imports.NewTable("uCEFMiscFunctions_CefBase64Decode", 0), // static function CefBase64Decode
			/* 53 */ imports.NewTable("uCEFMiscFunctions_CefUriEncode", 0), // static function CefUriEncode
			/* 54 */ imports.NewTable("uCEFMiscFunctions_CefUriDecode", 0), // static function CefUriDecode
			/* 55 */ imports.NewTable("uCEFMiscFunctions_CefGetPath", 0), // static function CefGetPath
			/* 56 */ imports.NewTable("uCEFMiscFunctions_CefIsRTL", 0), // static function CefIsRTL
			/* 57 */ imports.NewTable("uCEFMiscFunctions_CefCreateDirectory", 0), // static function CefCreateDirectory
			/* 58 */ imports.NewTable("uCEFMiscFunctions_CefGetTempDirectory", 0), // static function CefGetTempDirectory
			/* 59 */ imports.NewTable("uCEFMiscFunctions_CefCreateNewTempDirectory", 0), // static function CefCreateNewTempDirectory
			/* 60 */ imports.NewTable("uCEFMiscFunctions_CefCreateTempDirectoryInDirectory", 0), // static function CefCreateTempDirectoryInDirectory
			/* 61 */ imports.NewTable("uCEFMiscFunctions_CefDirectoryExists", 0), // static function CefDirectoryExists
			/* 62 */ imports.NewTable("uCEFMiscFunctions_CefDeleteFile", 0), // static function CefDeleteFile
			/* 63 */ imports.NewTable("uCEFMiscFunctions_CefZipDirectory", 0), // static function CefZipDirectory
			/* 64 */ imports.NewTable("uCEFMiscFunctions_CefIsKeyDown", 0), // static function CefIsKeyDown
			/* 65 */ imports.NewTable("uCEFMiscFunctions_CefIsKeyToggled", 0), // static function CefIsKeyToggled
			/* 66 */ imports.NewTable("uCEFMiscFunctions_GetCefMouseModifiersToEventFlags", 0), // static function GetCefMouseModifiersToEventFlags
			/* 67 */ imports.NewTable("uCEFMiscFunctions_GetCefMouseModifiersWithWPARAM", 0), // static function GetCefMouseModifiersWithWPARAM
			/* 68 */ imports.NewTable("uCEFMiscFunctions_GetCefKeyboardModifiers", 0), // static function GetCefKeyboardModifiers
			/* 69 */ imports.NewTable("uCEFMiscFunctions_GetWindowsMajorMinorVersion", 0), // static function GetWindowsMajorMinorVersion
			/* 70 */ imports.NewTable("uCEFMiscFunctions_RunningWindows10OrNewer", 0), // static function RunningWindows10OrNewer
			/* 71 */ imports.NewTable("uCEFMiscFunctions_GetDPIForHandle", 0), // static function GetDPIForHandle
			/* 72 */ imports.NewTable("uCEFMiscFunctions_GetDefaultCEFUserAgent", 0), // static function GetDefaultCEFUserAgent
			/* 73 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithIntDouble", 0), // static function DeviceToLogicalWithIntDouble
			/* 74 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithSingleDouble", 0), // static function DeviceToLogicalWithSingleDouble
			/* 75 */ imports.NewTable("uCEFMiscFunctions_LogicalToDeviceWithIntDouble", 0), // static function LogicalToDeviceWithIntDouble
			/* 76 */ imports.NewTable("uCEFMiscFunctions_GetScreenDPI", 0), // static function GetScreenDPI
			/* 77 */ imports.NewTable("uCEFMiscFunctions_GetDeviceScaleFactor", 0), // static function GetDeviceScaleFactor
			/* 78 */ imports.NewTable("uCEFMiscFunctions_DeleteDirContents", 0), // static function DeleteDirContents
			/* 79 */ imports.NewTable("uCEFMiscFunctions_DeleteFileList", 0), // static function DeleteFileList
			/* 80 */ imports.NewTable("uCEFMiscFunctions_MoveFileList", 0), // static function MoveFileList
			/* 81 */ imports.NewTable("uCEFMiscFunctions_CefGetDataURIWithStrX2", 0), // static function CefGetDataURIWithStrX2
			/* 82 */ imports.NewTable("uCEFMiscFunctions_CefGetDataURIWithPointerIntStrX2", 0), // static function CefGetDataURIWithPointerIntStrX2
			/* 83 */ imports.NewTable("uCEFMiscFunctions_ValidCefWindowHandle", 0), // static function ValidCefWindowHandle
			/* 84 */ imports.NewTable("uCEFMiscFunctions_GetCommandLineSwitchValue", 0), // static function GetCommandLineSwitchValue
			/* 85 */ imports.NewTable("uCEFMiscFunctions_InitializeCefTime", 0), // static procedure InitializeCefTime
			/* 86 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsChildWithWInfoWHandleRectStrDWORD", 0), // static procedure WindowInfoAsChildWithWInfoWHandleRectStrDWORD
			/* 87 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsPopUpWithWInfoWHandleStrDWORD", 0), // static procedure WindowInfoAsPopUpWithWInfoWHandleStrDWORD
			/* 88 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsWindowlessWithWInfoWHandleStrDWORD", 0), // static procedure WindowInfoAsWindowlessWithWInfoWHandleStrDWORD
			/* 89 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsChildWithWInfoWHandleRectStrBool", 0), // static procedure WindowInfoAsChildWithWInfoWHandleRectStrBool
			/* 90 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsPopUpWithWInfoWHandleStrBool", 0), // static procedure WindowInfoAsPopUpWithWInfoWHandleStrBool
			/* 91 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsWindowlessWithWInfoWHandleStrBool", 0), // static procedure WindowInfoAsWindowlessWithWInfoWHandleStrBool
			/* 92 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsChildWithWInfoWHandleRectStr", 0), // static procedure WindowInfoAsChildWithWInfoWHandleRectStr
			/* 93 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsPopUpWithWInfoWHandleStr", 0), // static procedure WindowInfoAsPopUpWithWInfoWHandleStr
			/* 94 */ imports.NewTable("uCEFMiscFunctions_WindowInfoAsWindowlessWithWInfoWHandleStr", 0), // static procedure WindowInfoAsWindowlessWithWInfoWHandleStr
			/* 95 */ imports.NewTable("uCEFMiscFunctions_CefSetCrashKeyValue", 0), // static procedure CefSetCrashKeyValue
			/* 96 */ imports.NewTable("uCEFMiscFunctions_CefLog", 0), // static procedure CefLog
			/* 97 */ imports.NewTable("uCEFMiscFunctions_CefKeyEventLog", 0), // static procedure CefKeyEventLog
			/* 98 */ imports.NewTable("uCEFMiscFunctions_CefMouseEventLog", 0), // static procedure CefMouseEventLog
			/* 99 */ imports.NewTable("uCEFMiscFunctions_OutputDebugMessage", 0), // static procedure OutputDebugMessage
			/* 100 */ imports.NewTable("uCEFMiscFunctions_OutputLastErrorMessage", 0), // static procedure OutputLastErrorMessage
			/* 101 */ imports.NewTable("uCEFMiscFunctions_CefGetExtensionsForMimeType", 0), // static procedure CefGetExtensionsForMimeType
			/* 102 */ imports.NewTable("uCEFMiscFunctions_CefLoadCRLSetsFile", 0), // static procedure CefLoadCRLSetsFile
			/* 103 */ imports.NewTable("uCEFMiscFunctions_CefCheckAltGrPressed", 0), // static procedure CefCheckAltGrPressed
			/* 104 */ imports.NewTable("uCEFMiscFunctions_DropEffectToDragOperation", 0), // static procedure DropEffectToDragOperation
			/* 105 */ imports.NewTable("uCEFMiscFunctions_DragOperationToDropEffect", 0), // static procedure DragOperationToDropEffect
			/* 106 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithMouseEventDouble", 0), // static procedure DeviceToLogicalWithMouseEventDouble
			/* 107 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithTouchEventDouble", 0), // static procedure DeviceToLogicalWithTouchEventDouble
			/* 108 */ imports.NewTable("uCEFMiscFunctions_DeviceToLogicalWithPointDouble", 0), // static procedure DeviceToLogicalWithPointDouble
			/* 109 */ imports.NewTable("uCEFMiscFunctions_LogicalToDeviceWithRectDouble", 0), // static procedure LogicalToDeviceWithRectDouble
			/* 110 */ imports.NewTable("uCEFMiscFunctions_InitializeWindowHandle", 0), // static procedure InitializeWindowHandle
		}
	})
	return uCEFMiscFunctionsImport
}
