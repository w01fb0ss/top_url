/*
 * Package: MinHeap
 * File: minHeap_test.go.go
 * Created Date: 2020/3/28 3:44 下午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 3:44 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package MinHeap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.Length() == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	heap.Insert(&UrlObj{1000, "https://www.microsoft.com/0JTZDRZDFLQ"})
	heap.Insert(&UrlObj{960, "https://www.baidu.com/0WSDZOKPNFW"})
	heap.Insert(&UrlObj{1200, "https://www.baidu.com/0ZHBLJFYRIA"})
	heap.Insert(&UrlObj{870, "https://www.baidu.com/0DFWVAAVRCI"})
	heap.Insert(&UrlObj{1500, "https://www.baidu.com/0RKMRUTCBNP"})

	if heap.Length() != 5 {
		t.Error("Insert 5 to Heap Failed")
	} else {
		t.Log("Insert 5 to Heap Pass")
	}

	v, _ := heap.Min()
	if v.Count != 870 {
		t.Error("Top is min Failed")
	} else {
		t.Log("Top is min Pass")
	}

	v, _ = heap.DeleteMin()
	if v.Count != 870 {
		t.Error("Delete Top min Failed")
	} else {
		t.Log("Delete Top min Pass")
	}

	v, _ = heap.DeleteMin()
	if v.Count != 960 {
		t.Error("Delete Top min Failed")
	} else {
		t.Log("Delete Top min Pass")
	}
}
