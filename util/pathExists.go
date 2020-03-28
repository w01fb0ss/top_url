/*
 * Package: util
 * File: pathExists.go
 * Created Date: 2020/3/28 10:12 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 10:12 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package util

import (
	"os"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

