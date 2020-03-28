/*
 * Package: BKDRHash
 * File: bkdrHash_test.go
 * Created Date: 2020/3/28 4:35 下午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/28 4:35 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package BKDRHash

import (
	"testing"
)

func TestBKDRHash64(t *testing.T) {
	if BKDRHash64("https://www.baidu.com/0RKMRUTCBNP") == uint64(582072976742644383) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}
