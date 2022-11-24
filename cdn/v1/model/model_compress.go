package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Compress struct {
	Status string `json:"status"`

	Type *string `json:"type,omitempty"`
}

func (o Compress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Compress struct{}"
	}

	return strings.Join([]string{"Compress", string(data)}, " ")
}
