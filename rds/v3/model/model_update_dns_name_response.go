package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
