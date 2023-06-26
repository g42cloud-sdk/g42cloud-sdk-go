package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ConcatInfo struct {
	Inputs *[]ObsObjInfo `json:"inputs,omitempty"`
}

func (o ConcatInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConcatInfo struct{}"
	}

	return strings.Join([]string{"ConcatInfo", string(data)}, " ")
}
