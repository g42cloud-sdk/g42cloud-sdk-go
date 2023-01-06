package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationRequest struct {
	Body *CreateApplicationRequestBody `json:"body,omitempty"`
}

func (o CreateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequest", string(data)}, " ")
}
