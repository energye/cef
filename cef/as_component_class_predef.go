//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/lcl/lcl"

// AsBitmap Convert a pointer object to an existing class object
func AsBitmap(obj interface{}) IBitmap {
	return lcl.AsBitmap(obj)
}

// AsStringList Convert a pointer object to an existing class object
func AsStringList(obj interface{}) IStringList {
	return lcl.AsStringList(obj)
}

// AsStrings Convert a pointer object to an existing class object
func AsStrings(obj interface{}) IStrings {
	return lcl.AsStrings(obj)
}

// AsCustomForm Convert a pointer object to an existing class object
func AsCustomForm(obj interface{}) ICustomForm {
	return lcl.AsCustomForm(obj)
}

// AsChromiumEvents Convert a pointer object to an existing class object
func AsChromiumEvents(obj interface{}) IChromiumEvents {
	return AsChromiumCore(obj)
}

// AsServerEvents Convert a pointer object to an existing class object
func AsServerEvents(obj interface{}) IServerEvents {
	return AsCEFServerComponent(obj)
}

// AsCEFUrlRequestClientEvents Convert a pointer object to an existing class object
func AsCEFUrlRequestClientEvents(obj interface{}) ICEFUrlRequestClientEvents {
	return AsCEFUrlRequestClientComponent(obj)
}

// AsCefViewDelegateEvents Convert a pointer object to an existing class object
func AsCefViewDelegateEvents(obj interface{}) ICefViewDelegateEvents {
	return AsCEFViewComponent(obj)
}
