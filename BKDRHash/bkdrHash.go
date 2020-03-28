/*
 * Package: main
 * File: bkdrhash.go
 * Created Date: 2020/3/26 4:18 下午
 * Author: alex
 * Contact: <xiaoshenlong1910@gmail.com>
 *
 * Last Modified: 2020/3/26 4:18 下午
 *
 * Copyright (c) 2020 XXX Company
 * license.
 */

package BKDRHash


func BKDRHash64(str string) uint64 {
	list := []byte(str)
	var seed uint64 = 131
	var hash uint64 = 0
	for i := 0; i < len(list); i++ {
		hash = hash*seed + uint64(list[i])
	}
	return hash & 0x7FFFFFFFFFFFFFFF
}
