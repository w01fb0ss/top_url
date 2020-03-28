/*
 * Package: CreateUrl
 * File: create_url.go
 * Created Date: 2020/3/28 11:09 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 11:09 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	DATA_FILE = "./data.txt"
	r         *rand.Rand
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func main() {
	f, err := os.OpenFile(DATA_FILE, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	prefixList := []string{"baidu", "google", "yahoo", "microsoft", "csdn", "bilibili", "jianshu",
		"youtube", "reddit", "facebook", "bitcointalk", "github"}
	suffixList := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		suffixList = append(suffixList, i)
	}

	for i := 0; i < 10; i++ {
		for _, prefix := range prefixList {
			for _, suffix := range suffixList {
				str := "https://www." + prefix + ".com/" + strconv.Itoa(suffix) + RandString(10)
				x := rand.Intn(10000)
				for i := 0; i < x; i++ {
					_, err = f.WriteString(str + "\n")
					if err != nil {
						return
					}

				}
			}
		}
	}
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
