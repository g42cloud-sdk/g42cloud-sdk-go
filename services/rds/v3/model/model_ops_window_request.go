package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OpsWindowRequest struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`
}

func (o OpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpsWindowRequest struct{}"
	}

	return strings.Join([]string{"OpsWindowRequest", string(data)}, " ")
}
