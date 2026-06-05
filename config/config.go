//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package config

var (
	GConfig = &Config{}
)

type Config struct {
	Chromium TChromium `json:"chromium"`
}

type TChromium struct {
	Dir     string `json:"dir"`
	Version string `json:"version"`
}
