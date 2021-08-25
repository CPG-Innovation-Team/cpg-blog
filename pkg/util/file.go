package util

import "os"

//检查文件或目录是否存在
// 存在则返回true，否则返回false

func Exist(fn string) bool {
	_, err := os.Stat(fn)
	return err == nil || os.IsExist(err)
}
