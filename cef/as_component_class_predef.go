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
