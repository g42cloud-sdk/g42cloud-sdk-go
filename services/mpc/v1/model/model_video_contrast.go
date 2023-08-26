package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoContrast struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`

	Contrast *string `json:"contrast,omitempty"`

	Brightness *string `json:"brightness,omitempty"`
}

func (o VideoContrast) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoContrast struct{}"
	}

	return strings.Join([]string{"VideoContrast", string(data)}, " ")
}
