package baidupcscmd

import (
	"fmt"
	"github.com/sersoong/BaiduPCS-Go/util"
)

// RunGetQuota 执行 获取当前用户空间配额信息, 并输出
func RunGetQuota() {
	quota, used, err := info.QuotaInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("总空间: %s, 已用空间: %s, 比率: %f%%\n",
		pcsutil.ConvertFileSize(quota),
		pcsutil.ConvertFileSize(used),
		100*float64(used)/float64(quota),
	)
}
