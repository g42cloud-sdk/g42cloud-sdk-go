package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OutputWatermarkPara struct {
	TimeDuration *int32 `json:"time_duration,omitempty"`
}

func (o OutputWatermarkPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputWatermarkPara struct{}"
	}

	return strings.Join([]string{"OutputWatermarkPara", string(data)}, " ")
}
