package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoSuperresolution struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`

	Scale *string `json:"scale,omitempty"`
}

func (o VideoSuperresolution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSuperresolution struct{}"
	}

	return strings.Join([]string{"VideoSuperresolution", string(data)}, " ")
}
