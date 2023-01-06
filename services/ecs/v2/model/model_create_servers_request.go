package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateServersRequest struct {
	Body *CreateServersRequestBody `json:"body,omitempty"`
}

func (o CreateServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServersRequest struct{}"
	}

	return strings.Join([]string{"CreateServersRequest", string(data)}, " ")
}
