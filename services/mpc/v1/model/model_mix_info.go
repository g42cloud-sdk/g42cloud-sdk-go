package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MixInfo struct {
	Inputs *[]InputSetting `json:"inputs,omitempty"`

	Layout *MixInfoLayout `json:"layout,omitempty"`
}

func (o MixInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MixInfo struct{}"
	}

	return strings.Join([]string{"MixInfo", string(data)}, " ")
}
