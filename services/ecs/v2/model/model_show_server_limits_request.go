package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowServerLimitsRequest struct {
}

func (o ShowServerLimitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerLimitsRequest struct{}"
	}

	return strings.Join([]string{"ShowServerLimitsRequest", string(data)}, " ")
}
