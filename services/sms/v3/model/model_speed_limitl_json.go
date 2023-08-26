package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SpeedLimitlJson struct {
	Start string `json:"start"`

	End string `json:"end"`

	Speed int32 `json:"speed"`
}

func (o SpeedLimitlJson) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimitlJson struct{}"
	}

	return strings.Join([]string{"SpeedLimitlJson", string(data)}, " ")
}
