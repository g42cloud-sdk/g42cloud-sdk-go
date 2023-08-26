package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateHealthMonitorRequest struct {
	Body *CreateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequest", string(data)}, " ")
}
