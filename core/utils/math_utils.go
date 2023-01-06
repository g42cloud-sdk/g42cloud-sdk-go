// Copyright 2022 G42 Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package utils

import (
	"crypto/rand"
	"math/big"
)

var reader = rand.Reader

func Max32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func Min32(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}

func RandInt32(min, max int32) int32 {
	if min >= max || min == 0 || max == 0 {
		return max
	}

	b, err := rand.Int(reader, big.NewInt(int64(max)))
	if err != nil {
		return max
	}
	return int32(b.Int64()) + min
}

func Pow32(x, y int32) int32 {
	var ans int32 = 1
	for y != 0 {
		ans *= x
		y--
	}
	return ans
}
