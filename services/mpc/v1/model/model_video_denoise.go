package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoDenoise struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`
}

func (o VideoDenoise) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoDenoise struct{}"
	}

	return strings.Join([]string{"VideoDenoise", string(data)}, " ")
}
