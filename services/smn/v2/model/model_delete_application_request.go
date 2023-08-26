package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteApplicationRequest struct {
	ApplicationUrn string `json:"application_urn"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
