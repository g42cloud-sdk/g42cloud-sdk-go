package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankRequest struct {
	Body *CreateLogtankRequestBody `json:"body,omitempty"`
}

func (o CreateLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequest struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequest", string(data)}, " ")
}
