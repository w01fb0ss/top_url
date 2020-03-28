/*
 * Package: ReadFile
 * File: readfile.go
 * Created Date: 2020/3/28 9:06 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 9:06 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package FileCut

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/w01fb0ss/top_url/BKDRHash"
	"github.com/w01fb0ss/top_url/util"
)

var (
	MAX_LINE = 10000
	NUM_FILE = 100
	TMP_DIR = "./tmp"
)

func init()  {
	exist, _ := util.PathExists(TMP_DIR)
	if exist {
		os.RemoveAll(TMP_DIR)
	}
	os.Mkdir(TMP_DIR, 0755)
}


// ReadFile 读取文件使用handle处理
func ReadFile(path string, handle func(strs []string)) error {
	// 读取文件
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// 逐行读取处理
	count := 0 // 行数
	buf := bufio.NewReader(f)
	strs := make([]string, MAX_LINE)

	for {
		line, _, err := buf.ReadLine()
		// 使用handle处理
		if count == MAX_LINE {
			handle(strs)
			strs = make([]string, MAX_LINE)
			count = 0
		}
		strs = append(strs, string(line))
		if err != nil {
			if err == io.EOF {
				handle(strs)
				strs = make([]string, MAX_LINE)
				count = 0
				return nil
			}
			return err
		}
		count ++
	}
}

// WriteToTmp 将分片写入文件中
func WriteToTmp(strs []string)  {
	fileMap := make(map[string][]string)
	// 将字符串写入map，key为文件名，value为该文件的内容
	for _, str := range strs {
		if str == "" {
			continue
		}
		subFilePath := TMP_DIR + "/" + strconv.Itoa(int(BKDRHash.BKDRHash64(str)) % NUM_FILE) + ".txt"
		if _, ok := fileMap[subFilePath]; ok {
			fileMap[subFilePath] = append(fileMap[subFilePath], str)
		} else {
			fileMap[subFilePath] = []string{str}
		}
	}

	// 遍历map，每一个map写入一个文件
	for k, strs := range fileMap {
		f, err := os.OpenFile(k, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, str := range strs {
				_, err = f.WriteString(str + "\n")
				if err != nil {
					continue
				}
			}
		}
		defer f.Close()
	}
}
