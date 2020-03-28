/*
 * Package: ResultToFile
 * File: resultToFile.go
 * Created Date: 2020/3/28 10:09 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 10:09 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package ResultToFile

import (
	"fmt"
	"os"
	"strconv"

	"github.com/w01fb0ss/top_url/MinHeap"
	"github.com/w01fb0ss/top_url/util"
)

var (
	RESULT_FILE = "./result.txt"
)

func init()  {
	exist, _ := util.PathExists(RESULT_FILE)
	if exist {
		os.Remove(RESULT_FILE)
	}
}

// ResultToFile 将结果写入RESULT_FILE
func ResultToFile(heap *MinHeap.MinHeap)  {
	f, err := os.OpenFile(RESULT_FILE, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	i := 0
	for heap.Length() != 0 {
		urlObj, _ := heap.DeleteMin()
		_, err = f.Write([]byte("Index:" + strconv.Itoa(i) + " --- Count: " + strconv.FormatInt(urlObj.Count, 10) + " --- Url: " + urlObj.Url + "\n"))
		if err != nil {
			fmt.Println(err.Error())
		}
		i ++
	}
}
