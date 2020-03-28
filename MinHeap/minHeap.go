/*
 * Package: main
 * File: minHeap.go
 * Created Date: 2020/3/26 8:50 下午
 * Author: alex
 * Contact: <xiaomhenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/26 8:50 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package MinHeap

import (
	"fmt"
)

type MinHeap struct {
	ArrU []*UrlObj
}

type UrlObj struct {
	Count int64
	Url   string
}

// NewMinHeap 最小堆的初始化
func NewMinHeap() *MinHeap {
	first := &UrlObj{
		Count: -1,
		Url:   "",
	}
	h := &MinHeap{ArrU: []*UrlObj{first}}
	return h
}

// Length 返回最小堆的总长度，为实际总长度减一
func (H *MinHeap) Length() int {
	return len(H.ArrU) - 1
}

// Min 返回最小堆的堆顶，即返回它的最小值
func (H *MinHeap) Min() (*UrlObj, error) {
	if len(H.ArrU) > 1 {
		return H.ArrU[1], nil
	}
	return nil, fmt.Errorf("heap is empty")
}

// Insert 插入元素
func (H *MinHeap) Insert(v *UrlObj) {
	// 将元素插到最后
	H.ArrU = append(H.ArrU, v)

	// 将最末节点与它的母节点(即i/2)相比较
	i := len(H.ArrU) - 1
	for ; (H.ArrU[i/2]).Count > v.Count; i /= 2 {
		H.ArrU[i] = H.ArrU[i/2]
	}

	H.ArrU[i] = v
}

// DeleteMin 删除最小堆堆顶并返回
func (H *MinHeap) DeleteMin() (*UrlObj, error) {
	if len(H.ArrU) <= 1 {
		return nil, fmt.Errorf("MinHeap is empty")
	}

	// 替换顶节点并将所有节点上浮
	minElement := H.ArrU[1]
	lastElement := H.ArrU[H.Length()]
	var i, child int
	for i = 1; i*2 < len(H.ArrU); i = child {
		child = i * 2
		if child < H.Length() && H.ArrU[child+1].Count < H.ArrU[child].Count {
			child++
		}
		if lastElement.Count > H.ArrU[child].Count {
			H.ArrU[i] = H.ArrU[child]
		} else {
			break
		}
	}
	H.ArrU[i] = lastElement
	// 删除最末尾的节点
	H.ArrU = H.ArrU[:len(H.ArrU)-1]
	return minElement, nil
}
