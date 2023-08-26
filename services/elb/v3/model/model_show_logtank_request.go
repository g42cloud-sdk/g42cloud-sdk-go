package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowLogtankRequest struct {
	LogtankId string `json:"logtank_id"`
}

func (o ShowLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankRequest struct{}"
	}

	return strings.Join([]string{"ShowLogtankRequest", string(data)}, " ")
}
