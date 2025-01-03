package cef

import (
	"github.com/energye/cef/cef/initialize"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tools"
	"os"
)

// Init 全局初始化, 需手动调用的函数
//
//	参数:
//	   libs 内置到应用程序的类库
//	   resources 内置到应用程序的资源文件
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	initialize.Initialize(libs, resources)
	// macos command line
	if tools.IsDarwin() {
		argsList := lcl.NewStringList()
		for _, v := range os.Args {
			argsList.Add(v)
		}
		// 手动设置进程命令参数
		SetCommandLine(argsList)
		argsList.Free()
	}
	api.SetCEFEventCallback(eventCallback)
	api.SetCEFRemoveEventCallback(removeEventCallback)
}
