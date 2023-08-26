package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigRequest struct {
}

func (o ShowConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigRequest", string(data)}, " ")
}
