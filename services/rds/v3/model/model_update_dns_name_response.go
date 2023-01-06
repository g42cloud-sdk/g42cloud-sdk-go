package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDnsNameResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDnsNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDnsNameResponse", string(data)}, " ")
}
