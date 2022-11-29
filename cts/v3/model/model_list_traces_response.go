package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTracesResponse struct {
	Traces *[]Traces `json:"traces,omitempty"`

	MetaData       *MetaData `json:"meta_data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTracesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTracesResponse struct{}"
	}

	return strings.Join([]string{"ListTracesResponse", string(data)}, " ")
}
