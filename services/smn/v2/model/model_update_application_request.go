package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateApplicationRequest struct {
	ApplicationUrn string `json:"application_urn"`

	Body *UpdateApplicationRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationRequest", string(data)}, " ")
}
