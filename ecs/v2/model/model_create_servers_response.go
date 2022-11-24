package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServersResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ServerIds      *[]string `json:"serverIds,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServersResponse struct{}"
	}

	return strings.Join([]string{"CreateServersResponse", string(data)}, " ")
}
