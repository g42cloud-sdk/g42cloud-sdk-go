package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRestoreInstanceResponse struct {
	Instance *CreateInstanceRespItem `json:"instance,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRestoreInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateRestoreInstanceResponse", string(data)}, " ")
}
