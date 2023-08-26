package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateCommandResultRequest struct {
	ServerId string `json:"server_id"`

	Body *CommandBody `json:"body,omitempty"`
}

func (o UpdateCommandResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateCommandResultRequest", string(data)}, " ")
}
