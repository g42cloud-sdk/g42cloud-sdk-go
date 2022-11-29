package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePoolResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Pool           *Pool `json:"pool,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdatePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolResponse struct{}"
	}

	return strings.Join([]string{"UpdatePoolResponse", string(data)}, " ")
}
