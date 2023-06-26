package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoDeblock struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`
}

func (o VideoDeblock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoDeblock struct{}"
	}

	return strings.Join([]string{"VideoDeblock", string(data)}, " ")
}
