package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CompressResponse struct {
	CompressSwitch int32 `json:"compress_switch"`

	CompressRules *[]CompressRules `json:"compress_rules,omitempty"`
}

func (o CompressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompressResponse struct{}"
	}

	return strings.Join([]string{"CompressResponse", string(data)}, " ")
}
