package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeManagement struct {
	ServerGroupReference *string `json:"serverGroupReference,omitempty"`
}

func (o NodeManagement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeManagement struct{}"
	}

	return strings.Join([]string{"NodeManagement", string(data)}, " ")
}
