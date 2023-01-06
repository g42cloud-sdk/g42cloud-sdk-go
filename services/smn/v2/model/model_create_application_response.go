package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	ApplicationUrn *string `json:"application_urn,omitempty"`

	ApplicationId  *string `json:"application_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
