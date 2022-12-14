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

package retry

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"
)

// RandomJitter 全抖动指数退避 delay = random(0, Exponential)
type RandomJitter struct {
}

func (r *RandomJitter) ComputeDelayBeforeNextRetry() int32 {
	delay := utils.Min32(MaxDelay, BaseDelay*(utils.Pow32(3, 2)))
	return utils.RandInt32(0, delay)
}
