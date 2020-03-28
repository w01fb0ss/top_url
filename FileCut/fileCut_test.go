/*
 * Package: FileCut
 * File: fileCut_test.go
 * Created Date: 2020/3/28 4:35 下午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 4:35 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package FileCut

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestReadFile(t *testing.T) {
	path := "../data.txt"
	tmpPath := "./tmp"
	err := ReadFile(path, WriteToTmp)
	if err != nil {
		t.Error("Read file is wrong")
	}

	// 随机打开一个文件
	x := rand.Intn(100)
	tmpPath = tmpPath + "/" + strconv.Itoa(x) + ".txt"
	f, err := os.Open(tmpPath)
	if err != nil {
		t.Error("Cannot open tmp file")
	} else {
		t.Log("Open tmp file success")
	}
	defer f.Close()

}
