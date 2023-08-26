package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DataPointInfo struct {
	Time *string `json:"time,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func (o DataPointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointInfo struct{}"
	}

	return strings.Join([]string{"DataPointInfo", string(data)}, " ")
}
