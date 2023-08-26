package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowsSpeedLimitsResponse struct {
	SpeedLimit     *[]SpeedLimitlJson `json:"speed_limit,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowsSpeedLimitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowsSpeedLimitsResponse struct{}"
	}

	return strings.Join([]string{"ShowsSpeedLimitsResponse", string(data)}, " ")
}
