/*
 * Package: FileToHeap
 * File: fileToHeap.go
 * Created Date: 2020/3/28 9:10 上午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 9:10 上午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package FileToHeap

import (
	"strconv"

	"github.com/w01fb0ss/top_url/FileCut"
	"github.com/w01fb0ss/top_url/MinHeap"
)

var (
	NUM_TOP  = 50
	CountMap map[string]int64
)

func init() {
	CountMap = make(map[string]int64)

}

// AddToCountMap 将所有的字符串添加到CountMap
func AddToCountMap(strs []string) {
	for _, str := range strs {
		if str == "" {
			continue
		}
		if _, ok := CountMap[str]; ok {
			CountMap[str]++
		} else {
			CountMap[str] = 1
		}
	}
}

// GetMinHeapFromTmp 获取tmp目录下文件的最小堆
func GetMinHeapFromTmp(path string) (*MinHeap.MinHeap, error) {
	err := FileCut.ReadFile(path, AddToCountMap)
	if err != nil {
		return nil, err
	}
	heap := MinHeap.NewMinHeap()
	for k, v := range CountMap {
		if heap.Length() < NUM_TOP {
			heap.Insert(&MinHeap.UrlObj{
				Count: v,
				Url:   k,
			})
			continue
		}
		min, _ := heap.Min()
		if min.Count <= v {
			heap.DeleteMin()
			heap.Insert(&MinHeap.UrlObj{
				Count: v,
				Url:   k,
			})
		}
	}
	return heap, nil
}

// MergeTwoHeap 合并两个最小堆，得到仅有一个NUM_TOP的最小堆
func MergeTwoHeap(heap, nHeap *MinHeap.MinHeap) *MinHeap.MinHeap {
	if nHeap == nil || nHeap.Length() == 0 {
		return heap
	}
	// 循环新堆每一个节点插入旧堆
	for nHeap.Length() != 0 {
		value, _ := nHeap.DeleteMin()
		if heap.Length() < NUM_TOP {
			heap.Insert(value)
			continue
		}
		min, _ := heap.Min()
		if min.Count >= value.Count {
			heap.DeleteMin()
			heap.Insert(value)
		}
	}

	return heap
}

// Reduce 每个最小堆合并成一个大堆
func Reduce() *MinHeap.MinHeap {
	heap := MinHeap.NewMinHeap()
	for i := 0; i < FileCut.NUM_FILE; i++ {
		// 对每一个切割的文件进行求最小堆
		NextHeap, err := GetMinHeapFromTmp(FileCut.TMP_DIR + strconv.Itoa(i) + ".txt")
		if err != nil {
			continue
		}

		heap = MergeTwoHeap(heap, NextHeap)
	}

	return heap
}
