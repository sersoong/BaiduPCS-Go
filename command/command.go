package baidupcscmd

import (
	"github.com/sersoong/BaiduPCS-Go/baidupcs"
	"github.com/sersoong/BaiduPCS-Go/config"
	"os"
)

var (
	info = new(baidupcs.PCSApi)
)

func init() {
	ReloadInfo()
}

// ReloadInfo 重载配置
func ReloadInfo() {
	pcsconfig.Reload()
	info = baidupcs.NewPCS(pcsconfig.ActiveBaiduUser.BDUSS)
}

// ReloadIfInConsole 程序在 Console 模式下才回重载配置
func ReloadIfInConsole() {
	if len(os.Args) == 1 {
		ReloadInfo()
	}
}
