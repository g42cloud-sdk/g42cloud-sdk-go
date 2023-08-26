package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateLogtankRequest struct {
	LogtankId string `json:"logtank_id"`

	Body *UpdateLogtankRequestBody `json:"body,omitempty"`
}

func (o UpdateLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogtankRequest", string(data)}, " ")
}
