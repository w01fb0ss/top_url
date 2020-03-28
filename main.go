/*
 * Package: top_url
 * File: main.go
 * Created Date: 2020/3/24 10:59 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/24 10:59 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package main

import (
	"fmt"

	"github.com/w01fb0ss/top_url/FileCut"
	"github.com/w01fb0ss/top_url/FileToHeap"
	"github.com/w01fb0ss/top_url/ResultToFile"
)

var (
    DATA_FILE_NAME = "./data.txt"
)




func main() {
	// 分片
	err := FileCut.ReadFile(DATA_FILE_NAME, FileCut.WriteToTmp)
	if err != nil {
		fmt.Println(err)
	}

	// 最小堆求top 50
	heap := FileToHeap.Reduce()

	// 结果保存
	ResultToFile.ResultToFile(heap)
}