package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePoolRequest struct {
	PoolId string `json:"pool_id"`

	Body *UpdatePoolRequestBody `json:"body,omitempty"`
}

func (o UpdatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequest", string(data)}, " ")
}
