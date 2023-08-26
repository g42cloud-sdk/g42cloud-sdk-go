package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateDnsNameResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDnsNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDnsNameResponse struct{}"
	}

	return strings.Join([]string{"CreateDnsNameResponse", string(data)}, " ")
}
