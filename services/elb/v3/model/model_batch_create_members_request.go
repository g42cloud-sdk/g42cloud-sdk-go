package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateMembersRequest struct {
	PoolId string `json:"pool_id"`

	Body *BatchCreateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchCreateMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersRequest", string(data)}, " ")
}
