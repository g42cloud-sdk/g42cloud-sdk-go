package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SpeedLimit struct {
	SpeedLimit []SpeedLimitlJson `json:"speed_limit"`
}

func (o SpeedLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimit struct{}"
	}

	return strings.Join([]string{"SpeedLimit", string(data)}, " ")
}
