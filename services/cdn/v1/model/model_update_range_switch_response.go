package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateRangeSwitchResponse struct {
	OriginRange    *OriginRangeBody `json:"origin_range,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateRangeSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRangeSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateRangeSwitchResponse", string(data)}, " ")
}
