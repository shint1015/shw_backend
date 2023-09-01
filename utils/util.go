package utils

import "os"

// FileExistCheck ファイルが存在するかどうかチェック
func FileExistCheck(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
