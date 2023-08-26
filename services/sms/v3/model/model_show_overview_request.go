package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowOverviewRequest struct {
}

func (o ShowOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowOverviewRequest", string(data)}, " ")
}
