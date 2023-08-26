package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CompressRequest struct {
	CompressSwitch *int32 `json:"compress_switch,omitempty"`
}

func (o CompressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompressRequest struct{}"
	}

	return strings.Join([]string{"CompressRequest", string(data)}, " ")
}
