package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type LogtankItem struct {
	Id string `json:"id"`

	LogGroupId string `json:"log_group_id"`

	LogStreamId string `json:"log_stream_id"`

	CreateTime string `json:"create_time"`

	UpdateTime string `json:"update_time"`
}

func (o LogtankItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogtankItem struct{}"
	}

	return strings.Join([]string{"LogtankItem", string(data)}, " ")
}
