package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRebootServersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRebootServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersResponse struct{}"
	}

	return strings.Join([]string{"BatchRebootServersResponse", string(data)}, " ")
}
