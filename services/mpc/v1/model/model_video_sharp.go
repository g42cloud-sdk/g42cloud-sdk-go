package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoSharp struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`

	Amount *string `json:"amount,omitempty"`
}

func (o VideoSharp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSharp struct{}"
	}

	return strings.Join([]string{"VideoSharp", string(data)}, " ")
}
