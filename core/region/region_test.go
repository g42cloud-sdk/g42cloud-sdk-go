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

package region

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	serviceName = "Service"
	regionId    = "region-id-1"
)

var endpoint = fmt.Sprintf("https://%s.%s.myg42cloud.com", strings.ToLower(serviceName), regionId)

func TestNewRegion(t *testing.T) {
	reg := NewRegion(regionId, endpoint)
	assert.Equal(t, &Region{
		Id:        regionId,
		Endpoints: []string{endpoint},
	}, reg)
}

func TestRegion_WithEndpointsOverride(t *testing.T) {
	testEndpoints := []string{"test"}
	reg := NewRegion(regionId, endpoint)
	assert.Equal(t, endpoint, reg.Endpoints[0])
	reg.WithEndpointsOverride(testEndpoints)
	assert.Equal(t, testEndpoints, reg.Endpoints)
}
