/*
 * Package: FileToHeap
 * File: fileToHeap_test.go
 * Created Date: 2020/3/28 4:36 下午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 4:36 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package FileToHeap

import (
	"testing"
)

func TestReduce(t *testing.T) {

	heap := Reduce()
	if heap.Length() != 0 {
		urlObj, err := heap.DeleteMin()
		if err != nil {
			t.Log("Failed")
			return
		}
		if urlObj.Count != 3859 {
			t.Log("Failed")
			return
		}
		t.Log("Success")
	}
}
